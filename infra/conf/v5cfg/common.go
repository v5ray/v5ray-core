package v5cfg

import (
	"context"
	"encoding/json"

	"github.com/golang/protobuf/proto"

	"github.com/v4fly/v4ray-core/v0/common/environment/envctx"
	"github.com/v4fly/v4ray-core/v0/common/environment/envimpl"
	"github.com/v4fly/v4ray-core/v0/common/registry"
)

func loadHeterogeneousConfigFromRawJSON(interfaceType, name string, rawJSON json.RawMessage) (proto.Message, error) {
	fsdef := envimpl.NewDefaultFileSystemDefaultImpl()
	ctx := envctx.ContextWithEnvironment(context.TODO(), fsdef)
	if len(rawJSON) == 0 {
		rawJSON = []byte("{}")
	}
	return registry.LoadImplementationByAlias(ctx, interfaceType, name, []byte(rawJSON))
}

// LoadHeterogeneousConfigFromRawJSON private API
func LoadHeterogeneousConfigFromRawJSON(ctx context.Context, interfaceType, name string, rawJSON json.RawMessage) (proto.Message, error) {
	return loadHeterogeneousConfigFromRawJSON(interfaceType, name, rawJSON)
}
