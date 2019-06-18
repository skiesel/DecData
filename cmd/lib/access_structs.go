package lib

import (
	"net"
)

type AccessRegisterRequest struct {
	ClientID string
	IP net.IP
}

type AccessRegisterResponse struct {

}

type AccessRequest struct {
	ClientID string
}

type AccessResponse struct {
	IP net.IP
}
