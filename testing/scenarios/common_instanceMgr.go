package scenarios

import (
	core "github.com/v4fly/v4ray-core/v0"
	"github.com/v4fly/v4ray-core/v0/app/instman"
	"github.com/v4fly/v4ray-core/v0/common"
	"github.com/v4fly/v4ray-core/v0/common/serial"
	"github.com/v4fly/v4ray-core/v0/features/extension"
)

func NewInstanceManagerInstanceConfig() *core.Config {
	config := &core.Config{}
	config.App = append(config.App, serial.ToTypedMessage(&instman.Config{}))
	return config
}

func NewInstanceManagerCoreInstance() (*core.Instance, extension.InstanceManagement) {
	coreConfig := NewInstanceManagerInstanceConfig()
	instance, err := core.New(coreConfig)
	if err != nil {
		panic(err)
	}
	common.Must(instance.Start())
	instanceMgr := instance.GetFeature(extension.InstanceManagementType())
	InstanceMgrIfce := instanceMgr.(extension.InstanceManagement)
	return instance, InstanceMgrIfce
}
