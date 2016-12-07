package configurable

import (
	"github.com/mcuadros/go-defaults"
	"github.com/src-d/envconfig"
)

// Interface to allow InitConfig to properly configure the config struct using
// environment variables, default values and ETCD.
type Configurable interface {
	Init()
}

// Initializes the configuration
func InitConfig(config Configurable) {
	defer defaults.SetDefaults(config)

	envconfig.MustProcess("config", config)

	config.Init()
}
