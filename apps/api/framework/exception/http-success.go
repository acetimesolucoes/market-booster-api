package http_exception

type HttpSuccess[Type any] struct {
	Data  Type  `json:"data"`
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
