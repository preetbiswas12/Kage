package network

import (
	"net/http"
	"time"
)

var transport = http.DefaultTransport.(*http.Transport).Clone()

func init() {
	transport.MaxIdleConns = 100
	transport.MaxIdleConnsPerHost = 100
	transport.MaxConnsPerHost = 200
	transport.IdleConnTimeout = 30 * time.Second
	// Increased from 30s to 2 minutes to handle slow servers during large downloads
	transport.ResponseHeaderTimeout = 2 * time.Minute
	transport.ExpectContinueTimeout = 30 * time.Second
}

// Client is the default HTTP client used for downloads.
// Timeout is set to 10 minutes to handle large file downloads in long-running operations.
var Client = &http.Client{
	Timeout:   10 * time.Minute,
	Transport: transport,
}
