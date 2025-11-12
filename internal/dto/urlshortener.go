package dto

type EncodeURLReq struct {
	URL string
}

type DecodeURLReq struct {
	URL string
}

type EncodeURLResp struct {
	URL string `json:"url"`
}

type DecodeURLResp struct {
	URL string `json:"url"`
}
