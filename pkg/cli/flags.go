package cli

import (
	"flag"
)

// Command line parameters
var FlagConfigPath = flag.String("config-path", "./profiles", "Storage directory for configuration profiles")
var FlagProfile    = flag.String("profile"    , "default"   , "Name of the configuration profile. Configuration settings '<config-path>/application-<profile>.yaml'")
