package configurable

import (
	"fmt"
	"os"
	"strings"

	"github.com/coreos/etcd/client"
	"github.com/fatih/color"
	"github.com/mcuadros/go-defaults"
	"github.com/mcuadros/go-etcd-hydrator"
	"github.com/src-d/envconfig"
)

/*
Interface to allow InitConfig to properly configure the config struct using
environment variables, default values and ETCD.
*/
type Configurable interface {
	EtcdConfigurable() bool
	EtcdServers() string
	EtcdEnvironment() string
}

// Initializes the configuration
func InitConfig(config Configurable) {
	defer defaults.SetDefaults(config)

	envconfig.MustProcess("config", config)

	/*
	   TODO: Leaving this as close as it was before.
	   Should we create a way for extending configuration?
	   Can we delay it until we know more about this issue?
	   Will etcd survive?
	*/
	if config.EtcdConfigurable() {
		if len(config.EtcdServers()) == 0 {
			printWarning()
			return
		}

		if isRunningTests() {
			panic("etcd servers cannot be used when running tests.")
		}

		etcdClient, err := client.New(client.Config{
			Endpoints: strings.Split(config.EtcdServers(), ","),
			Transport: client.DefaultTransport,
		})

		if err != nil {
			panic(err)
		}

		h := hydrator.NewHydrator(etcdClient)
		h.Folder = config.EtcdEnvironment()
		h.Hydrate(config)
	}
}

func printWarning() {
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Fprintf(os.Stderr, cyan("Running on dev mode, no etcd servers given.\n"))
}

func isRunningTests() bool {
	//this is a tricky way to check if the code is running under test: we check
	//if the binary running name ends of ".test"
	return strings.HasSuffix(os.Args[0], ".test")
}
