package dto

type EncodeURLReq struct {
	URL string `json:"url" binding:"required,url"`
}

type DecodeURLReq struct {
	URL string `json:"url" binding:"required,url"`
}

type EncodeURLResp struct {
	URL string `json:"url"`
}

type DecodeURLResp struct {
	URL string `json:"url"`
}
