package http_exception

import "github.com/acetimesolutions/marketbooster/domain"

type HttpSuccess[Enterprises domain.Enterprises] struct {
	Data  Enterprises `json:"data"`
	Page  int64       `json:"page"`
	Limit int64       `json:"limit"`
}
