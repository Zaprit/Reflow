// Package config is the package that handles application config
package config

import (
	_ "embed" // This is required for a go embed (linter made me put this here)
	"path/filepath"

	"github.com/BurntSushi/toml"
)

//go:embed reflow.toml.sample
var defaultConfig string

// ConfigurationFile is the configuration file for reflow, corrected for file separator.
var configurationFile = filepath.FromSlash("./conf/reflow.toml")

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Repo     RepoConfig
}

type DatabaseConfig struct {
	Driver   string
	Hostname string
	Port     int
	Database string
	Username string
	Password string
}

type ServerConfig struct {
	Listen  string
	Port    int
	TLS     bool
	TLSCert string
	Debug   bool
}

type RepoConfig struct {
	RepoURL string
}

// ConfigData the struct with the config in (duh)
var ConfigData Config

// LoadConfig loads the configuration from disk and sets Conf
func LoadConfig() {

	_, err := toml.DecodeFile(configurationFile, &ConfigData)
	if err != nil {
		panic(err.Error())
	}

}
