package lib

import (
	"net/url"
)

type LocalRegisterRequest struct {
	Service string
	URL *url.URL
}

type LocalRegisterResponse struct {
	Port int64
}

type LocalDataRequest struct {
	Service string
	data interface{}
}
