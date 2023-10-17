package request

type IpReq struct {
	IpType int    `form:"ip_type" binding:"required"`
	Ip     string `form:"ip" binding:"required"`
}
