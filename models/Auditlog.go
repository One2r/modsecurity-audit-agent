package models

type Auditlog struct {
	Transaction struct {
		ClientIP   string `json:"client_ip"`
		TimeStamp  string `json:"time_stamp"`
		ServerID   string `json:"server_id"`
		ClientPort int    `json:"client_port"`
		HostIP     string `json:"host_ip"`
		HostPort   int    `json:"host_port"`
		UniqueID   string `json:"unique_id"`
		Request    struct {
			Method      string  `json:"method"`
			HTTPVersion float64 `json:"http_version"`
			URI         string  `json:"uri"`
			Headers     struct {
				UserAgent      string `json:"User-Agent"`
				Accept         string `json:"Accept"`
				Host           string `json:"Host"`
				AcceptEncoding string `json:"Accept-Encoding"`
				Connection     string `json:"Connection"`
			} `json:"headers"`
		} `json:"request"`
		Response struct {
			Body     string `json:"body"`
			HTTPCode int    `json:"http_code"`
			Headers  struct {
				Server        string `json:"Server"`
				Date          string `json:"Date"`
				ContentLength string `json:"Content-Length"`
				ContentType   string `json:"Content-Type"`
				Connection    string `json:"Connection"`
			} `json:"headers"`
		} `json:"response"`
		Producer struct {
			Modsecurity    string   `json:"modsecurity"`
			Connector      string   `json:"connector"`
			SecrulesEngine string   `json:"secrules_engine"`
			Components     []string `json:"components"`
		} `json:"producer"`
		Messages []struct {
			Message string `json:"message"`
			Details struct {
				Match      string `json:"match"`
				Reference  string `json:"reference"`
				RuleID     string `json:"ruleId"`
				File       string `json:"file"`
				LineNumber string `json:"lineNumber"`
				Data       string `json:"data"`
				Severity   string `json:"severity"`
				Ver        string `json:"ver"`
				Rev        string `json:"rev"`
				Tags       []any  `json:"tags"`
				Maturity   string `json:"maturity"`
				Accuracy   string `json:"accuracy"`
			} `json:"details"`
		} `json:"messages"`
	} `json:"transaction"`
}
