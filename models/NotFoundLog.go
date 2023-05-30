package models

type NotFoundLog struct {
	ClientIP  string `json:"client_ip"`
	TimeStamp string `json:"time_stamp"`
	Request   struct {
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
}
