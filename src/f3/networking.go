package f3

import (
	"net"
	"net/http"
	"time"
)

// DefaultDialer is the default dialer that is being used when new Form 3 clients are created.
var DefaultDialer = &net.Dialer{}

// DefaultTransport is the default transport configuration to be used when new Form 3 clients are created.
var DefaultTransport = &http.Transport{
	DialContext:     DefaultDialer.DialContext,
	MaxIdleConns:    10,
	IdleConnTimeout: 30 * time.Second,
}
