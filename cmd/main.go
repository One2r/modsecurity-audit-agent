package main

import (
	"io"
	"os"
	"time"
	"net/http"
	"os/signal"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"modsecurity-auditlog-agent/internal/router"
	//"modsecurity-auditlog-agent/internal/util"
)

func main() {
	appPath, _ := os.Getwd()
	configDir := appPath + "/configs"

	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		os.Exit(0)
	}

	local := time.FixedZone("CST", 8*60*60)
	time.Local = local

	logDir := appPath + "/storage/logs"

	internalLogFile := logDir + "/internal.log"
	writeSyncer, _ := os.OpenFile(internalLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	logLevel, _ := zapcore.ParseLevel(viper.GetString("app.log-level"))
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	if !viper.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DisableConsoleColor()

	ginLogFile := logDir + "/gin.log"
	f, _ := os.OpenFile(ginLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s\" %d %s \"%s\" %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	router.SetRouter(r)

	//util.Init()

	srv := &http.Server{
		Addr:    ":" + viper.GetString("app.port"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatal("Server Shutdown err :", err)
	}
}
