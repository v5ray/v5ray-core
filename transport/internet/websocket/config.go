package websocket

import (
	"net/http"

	"github.com/v4fly/v4ray-core/v0/common"
	"github.com/v4fly/v4ray-core/v0/transport/internet"
)

const protocolName = "websocket"

func (c *Config) GetNormalizedPath() string {
	path := c.Path
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		return "/" + path
	}
	return path
}

func (c *Config) GetRequestHeader() http.Header {
	header := http.Header{}
	for _, h := range c.Header {
		header.Add(h.Key, h.Value)
	}
	return header
}

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
