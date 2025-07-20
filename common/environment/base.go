package environment

import (
	"github.com/v4fly/v4ray-core/v0/common/environment/filesystemcap"
	"github.com/v4fly/v4ray-core/v0/features/extension/storage"
	"github.com/v4fly/v4ray-core/v0/transport/internet"
	"github.com/v4fly/v4ray-core/v0/transport/internet/tagged"
)

type BaseEnvironmentCapabilitySet interface {
	FeaturesLookupCapabilitySet
	LogCapabilitySet
}

type BaseEnvironment interface {
	BaseEnvironmentCapabilitySet
	doNotImpl()
}

type SystemNetworkCapabilitySet interface {
	Dialer() internet.SystemDialer
	Listener() internet.SystemListener
}

type InstanceNetworkCapabilitySet interface {
	OutboundDialer() tagged.DialFunc
}

type FeaturesLookupCapabilitySet interface {
	RequireFeatures() interface{}
}

type LogCapabilitySet interface {
	RecordLog() interface{}
}

type FileSystemCapabilitySet interface {
	filesystemcap.FileSystemCapabilitySet
}

type PersistentStorageCapabilitySet interface {
	PersistentStorage() storage.ScopedPersistentStorage
}
type TransientStorageCapabilitySet interface {
	TransientStorage() storage.ScopedTransientStorage
}
