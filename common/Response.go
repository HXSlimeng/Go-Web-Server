package common

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageRes struct {
	Response
	Total int64 `json:"total"`
}