package blackhole_test

import (
	"context"
	"testing"

	"github.com/v4fly/v4ray-core/v0/common"
	"github.com/v4fly/v4ray-core/v0/common/buf"
	"github.com/v4fly/v4ray-core/v0/common/serial"
	"github.com/v4fly/v4ray-core/v0/proxy/blackhole"
	"github.com/v4fly/v4ray-core/v0/transport"
	"github.com/v4fly/v4ray-core/v0/transport/pipe"
)

func TestBlackHoleHTTPResponse(t *testing.T) {
	handler, err := blackhole.New(context.Background(), &blackhole.Config{
		Response: serial.ToTypedMessage(&blackhole.HTTPResponse{}),
	})
	common.Must(err)

	reader, writer := pipe.New(pipe.WithoutSizeLimit())

	readerError := make(chan error)
	var mb buf.MultiBuffer
	go func() {
		b, e := reader.ReadMultiBuffer()
		mb = b
		readerError <- e
	}()

	link := transport.Link{
		Reader: reader,
		Writer: writer,
	}
	common.Must(handler.Process(context.Background(), &link, nil))
	common.Must(<-readerError)
	if mb.IsEmpty() {
		t.Error("expect http response, but nothing")
	}
}
