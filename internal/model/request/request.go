package request

type IpReq struct {
	IpType int    `form:"ip_type" binding:"required"`
	Ip     string `form:"ip" binding:"required"`
}

type WorkwxMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

type DingtalkMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
}

type MdTplMessage struct {
	ClientIp  string
	Host      string
	URI       string
	TimeStamp string
	RuleId    string
	Match     string
	HTTPCode  int
	IsBanIp   string
}
