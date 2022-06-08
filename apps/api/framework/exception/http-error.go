package http_exception

type HttpError struct {
	Error   any    `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
