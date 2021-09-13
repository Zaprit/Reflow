// Package config is the package that handles application config
package config

import "path/filepath"

// ConfigurationFile is the configuration file for reflow, corrected for file separator.
var ConfigurationFile = filepath.FromSlash("./conf/reflow.conf")
