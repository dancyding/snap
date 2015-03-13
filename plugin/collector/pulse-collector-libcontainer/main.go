package main

import (
	"os"

	// Import the pulse plugin library
	"github.com/intelsdilabs/pulse/control/plugin"
	// Import our collector plugin implementation
)

func main() {
	// Three things provided:
	//   the definition of the plugin metadata
	//   the implementation satfiying plugin.CollectorPlugin
	//   the collector configuration policy satifying plugin.ConfigRules

	// Define default policy
	policy := libcontainer.ConfigPolicy()

	// Define metadata about Plugin
	meta := libcontainer.Meta()

	// Start a collector
	//plugin.StartCollector(meta, new(facter.Facter), policy, os.Args[0], os.Args[1])
	plugin.Start(meta, new(libcontainer.Libcontainer), policy, os.Args[1])
}
