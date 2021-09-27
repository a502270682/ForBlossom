package proto

type PingReq struct {
}

type PingRsp struct {
	Success string `json:"success"`
}

type AccessTokenRsp struct {
	Token string `json:"token"`
	TTL   int    `json:"ttl"`
}
