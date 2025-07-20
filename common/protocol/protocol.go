package protocol

import (
	"errors"
)

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

var ErrProtoNeedMoreData = errors.New("protocol matches, but need more data to complete sniffing")
