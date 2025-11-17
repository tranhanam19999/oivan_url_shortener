package dto

type EchoHttpErrorResp struct {
	Internal error       `json:"-"` // Stores the error returned by an external dependency
	Message  interface{} `json:"message"`
	Code     int         `json:"-"`
}
