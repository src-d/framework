package configurable

import (
	"github.com/mcuadros/go-defaults"
	"github.com/src-d/envconfig"
)

// Configurable allows InitConfig to properly configure the config struct using
// environment variables and default values.
type Configurable interface {
	Init()
}

// Initializes the configuration
func InitConfig(config Configurable) {
	defer defaults.SetDefaults(config)

	envconfig.MustProcess("config", config)

	config.Init()
}
