package configurable

var (
	// Initialized BasicConfiguration
	Config *BasicConfiguration = &BasicConfiguration{}
)

// Default configuration
type BasicConfiguration struct {
}

// Default, specific initialization for BasicConfiguration
func (c *BasicConfiguration) Init() {
}

func init() {
	InitConfig(Config)
}
