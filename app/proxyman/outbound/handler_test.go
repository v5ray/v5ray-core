package outbound_test

import (
	"context"
	"testing"
	_ "unsafe"

	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/v4fly/v4ray-core/v0"
	"github.com/v4fly/v4ray-core/v0/app/policy"
	. "github.com/v4fly/v4ray-core/v0/app/proxyman/outbound"
	"github.com/v4fly/v4ray-core/v0/app/stats"
	"github.com/v4fly/v4ray-core/v0/common/environment"
	"github.com/v4fly/v4ray-core/v0/common/environment/deferredpersistentstorage"
	"github.com/v4fly/v4ray-core/v0/common/environment/envctx"
	"github.com/v4fly/v4ray-core/v0/common/environment/filesystemimpl"
	"github.com/v4fly/v4ray-core/v0/common/environment/systemnetworkimpl"
	"github.com/v4fly/v4ray-core/v0/common/environment/transientstorageimpl"
	"github.com/v4fly/v4ray-core/v0/common/net"
	"github.com/v4fly/v4ray-core/v0/common/serial"
	"github.com/v4fly/v4ray-core/v0/features/outbound"
	"github.com/v4fly/v4ray-core/v0/proxy/freedom"
	"github.com/v4fly/v4ray-core/v0/transport/internet"
)

func TestInterfaces(t *testing.T) {
	_ = (outbound.Handler)(new(Handler))
	_ = (outbound.Manager)(new(Manager))
}

//go:linkname toContext github.com/v4fly/v4ray-core/v0.toContext
func toContext(ctx context.Context, v *core.Instance) context.Context

func TestOutboundWithoutStatCounter(t *testing.T) {
	config := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&stats.Config{}),
			serial.ToTypedMessage(&policy.Config{
				System: &policy.SystemPolicy{
					Stats: &policy.SystemPolicy_Stats{
						InboundUplink: true,
					},
				},
			}),
		},
	}

	v, _ := core.New(config)
	v.AddFeature((outbound.Manager)(new(Manager)))
	ctx := toContext(context.Background(), v)
	defaultNetworkImpl := systemnetworkimpl.NewSystemNetworkDefault()
	defaultFilesystemImpl := filesystemimpl.NewDefaultFileSystemDefaultImpl()
	deferredPersistentStorageImpl := deferredpersistentstorage.NewDeferredPersistentStorage(ctx)
	rootEnv := environment.NewRootEnvImpl(ctx,
		transientstorageimpl.NewScopedTransientStorageImpl(), defaultNetworkImpl.Dialer(), defaultNetworkImpl.Listener(),
		defaultFilesystemImpl, deferredPersistentStorageImpl)
	proxyEnvironment := rootEnv.ProxyEnvironment("o")
	ctx = envctx.ContextWithEnvironment(ctx, proxyEnvironment)
	h, _ := NewHandler(ctx, &core.OutboundHandlerConfig{
		Tag:           "tag",
		ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
	})
	conn, _ := h.(*Handler).Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), 13146))
	_, ok := conn.(*internet.StatCouterConnection)
	if ok {
		t.Errorf("Expected conn to not be StatCouterConnection")
	}
}

func TestOutboundWithStatCounter(t *testing.T) {
	config := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&stats.Config{}),
			serial.ToTypedMessage(&policy.Config{
				System: &policy.SystemPolicy{
					Stats: &policy.SystemPolicy_Stats{
						OutboundUplink:   true,
						OutboundDownlink: true,
					},
				},
			}),
		},
	}

	v, _ := core.New(config)
	v.AddFeature((outbound.Manager)(new(Manager)))
	ctx := toContext(context.Background(), v)
	defaultNetworkImpl := systemnetworkimpl.NewSystemNetworkDefault()
	defaultFilesystemImpl := filesystemimpl.NewDefaultFileSystemDefaultImpl()
	deferredPersistentStorageImpl := deferredpersistentstorage.NewDeferredPersistentStorage(ctx)
	rootEnv := environment.NewRootEnvImpl(ctx,
		transientstorageimpl.NewScopedTransientStorageImpl(), defaultNetworkImpl.Dialer(), defaultNetworkImpl.Listener(),
		defaultFilesystemImpl, deferredPersistentStorageImpl)
	proxyEnvironment := rootEnv.ProxyEnvironment("o")
	ctx = envctx.ContextWithEnvironment(ctx, proxyEnvironment)
	h, _ := NewHandler(ctx, &core.OutboundHandlerConfig{
		Tag:           "tag",
		ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
	})
	conn, _ := h.(*Handler).Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), 13146))
	_, ok := conn.(*internet.StatCouterConnection)
	if !ok {
		t.Errorf("Expected conn to be StatCouterConnection")
	}
}
