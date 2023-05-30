# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -ldflags "-s -w"
GOCLEAN=$(GOCMD) clean
BINARY_NAME=modsecurity-audit-agent
BINARY_DARWIN=$(BINARY_NAME).darwin
BINARY_UNIX=$(BINARY_NAME).unix
BINARY_ARM=$(BINARY_NAME).arm64
BINARY_WIN=$(BINARY_NAME).exe

build:
		$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/

clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_DARWIN)
		rm -f $(BINARY_UNIX)
		rm -f $(BINARY_ARM)

# Cross compilation
build-darwin:
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_DARWIN) -v ./cmd/ && upx -9 $(BINARY_DARWIN)
build-unix:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/ && upx -9 $(BINARY_UNIX)
build-arm64:
		CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BINARY_ARM) -v ./cmd/ && upx -9 $(BINARY_ARM)
build-win:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WIN) -v ./cmd/ && upx -9 $(BINARY_WIN)
