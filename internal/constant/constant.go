package constant

const (
	IP_TYPE_WHITELIST int = 1
	IP_TYPE_BLACKLIST int = 2
)

const (
	RDS_KEY_IP_WHITELIST             string = "waf:ip:whitelist"
	RDS_KEY_IP_BLACKLIST             string = "waf:ip:blacklist"
	RDS_KEY_PREFIX_STATS_IP_SCAN_404 string = "waf:stats:scan404:"
)
