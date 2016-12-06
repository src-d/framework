package configurable

var (
	Config *BasicConfiguration = &BasicConfiguration{}
)

type BasicConfiguration struct {
	Environment string `envconfig:"ENVIRONMENT"`
	Etcd        struct {
		Servers string `envconfig:"ETCD_SERVERS",required:"true"`
	}
}

func (c BasicConfiguration) EtcdConfigurable() bool {
	return true
}

func (c BasicConfiguration) EtcdServers() string {
	return c.Etcd.Servers
}

func (c BasicConfiguration) EtcdEnvironment() string {
	return c.Environment
}

func init() {
	InitConfig(Config)
}
