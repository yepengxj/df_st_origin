package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/openshift/origin/pkg/cmd/dockerregistry"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	// TODO convert to flags instead of a config file?
	configurationPath := ""
	if flag.NArg() > 0 {
		configurationPath = flag.Arg(0)
	}
	if configurationPath == "" {
<<<<<<< HEAD
		configurationPath = os.Getenv("REGISTRY_CONFIGURATION_PATH")
	}

	if configurationPath == "" {
		fmt.Println("configuration path unspecified")
		os.Exit(1)
	}
	// Prevent a warning about unrecognized environment variable
	os.Unsetenv("REGISTRY_CONFIGURATION_PATH")
=======
		fmt.Println("configuration path unspecified")
		os.Exit(1)
	}
>>>>>>> a967081138277d7e44651d9773a272ac8045916e

	configFile, err := os.Open(configurationPath)
	if err != nil {
		log.Fatalf("Unable to open configuration file: %s", err)
	}

	dockerregistry.Execute(configFile)
}
