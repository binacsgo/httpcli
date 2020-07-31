package httpcli

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

const (
	defaultDialTimeout = 10 * time.Second
	keepAliveInterval  = 30 * time.Second
)

var (
	gTransport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
		ResponseHeaderTimeout: 2 * time.Minute,
		DisableCompression:    true,
		DisableKeepAlives:     false,
		IdleConnTimeout:       2 * time.Minute,
		MaxIdleConns:          4,
		MaxIdleConnsPerHost:   2,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			dialer := &net.Dialer{Timeout: defaultDialTimeout, KeepAlive: keepAliveInterval}
			return dialer.DialContext(ctx, network, addr)
		},
	}
	gTransportSkipVerify = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		ResponseHeaderTimeout: 2 * time.Minute,
		DisableCompression:    true,
		DisableKeepAlives:     false,
		IdleConnTimeout:       2 * time.Minute,
		MaxIdleConns:          4,
		MaxIdleConnsPerHost:   2,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			dialer := &net.Dialer{Timeout: defaultDialTimeout, KeepAlive: keepAliveInterval}
			return dialer.DialContext(ctx, network, addr)
		},
	}
	gHTTPClient = &http.Client{
		Transport: gTransport,
	}
	gHTTPClientSkipVerify = &http.Client{
		Transport: gTransportSkipVerify,
	}
)
