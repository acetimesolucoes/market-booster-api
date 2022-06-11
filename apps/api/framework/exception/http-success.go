package http_exception

type HttpSuccess[T any] struct {
	Data  T     `json:"data"`
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
