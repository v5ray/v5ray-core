package dataurlfetcher

import (
	"context"
	"strings"

	"github.com/vincent-petithory/dataurl"

	"github.com/v4fly/v4ray-core/v0/app/subscription"
	"github.com/v4fly/v4ray-core/v0/app/subscription/documentfetcher"
	"github.com/v4fly/v4ray-core/v0/common"
)

//go:generate go run github.com/v4fly/v4ray-core/v0/common/errors/errorgen

func newDataURLFetcher() *dataURLFetcher {
	return &dataURLFetcher{}
}

func init() {
	common.Must(documentfetcher.RegisterFetcher("dataurl", newDataURLFetcher()))
}

type dataURLFetcher struct{}

func (d *dataURLFetcher) DownloadDocument(ctx context.Context, source *subscription.ImportSource, opts ...documentfetcher.FetcherOptions) ([]byte, error) {
	dataURL, err := dataurl.DecodeString(source.Url)
	if err != nil {
		return nil, newError("unable to decode dataURL").Base(err)
	}
	if dataURL.MediaType.Type != "application" {
		return nil, newError("unsupported media type: ", dataURL.MediaType.Type)
	}
	if !strings.HasPrefix(dataURL.MediaType.Subtype, "vnd.v2ray.subscription") {
		return nil, newError("unsupported media subtype: ", dataURL.MediaType.Subtype)
	}
	return []byte(source.Url), nil
}
