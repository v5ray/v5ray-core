package drain

import "io"

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
