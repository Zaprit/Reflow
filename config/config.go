// Package config is the package that handles application config
package config

import (
	_ "embed" // This is required for a go embed (linter made me put this here)
	"fmt"
	"os"
	"path/filepath"

	"github.com/alyu/configparser"
)

//go:embed reflow.conf.sample
var defaultConfig string

// ConfigurationFile is the configuration file for reflow, corrected for file separator.
var configurationFile = filepath.FromSlash("./conf/reflow.conf")

// Conf is the global configuration for reflow
var Conf *configparser.Configuration

// RepoURL is the global URL of the mod repository, this is mainly for backwards compatibility with solder.
var RepoURL string

// LoadConfig loads the configuration from disk and sets Conf
func LoadConfig() {
	_, err := os.Stat("conf")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("conf", 0750)
			if err != nil {
				panic(err)
			}
		}

		info, _ := os.Stat("conf")

		if !info.IsDir() {
			err = os.Remove("conf")
			if err != nil {
				panic(err.Error())
			}

			err = os.Mkdir("conf", 0750)

			if err != nil {
				panic(err.Error())
			}
		}

		var conf, er2 = os.Create(filepath.FromSlash("./conf/reflow.conf"))

		if er2 != nil {
			panic(err.Error())
		}

		_, er3 := conf.WriteString(defaultConfig)

		if er3 != nil {
			panic(err.Error())
		}

		err = conf.Sync()

		if err != nil {
			panic(err.Error())
		}

		err = conf.Close()

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Config created, please edit reflow.conf to match your settings")
		os.Exit(0)
	}

	cfg, err := configparser.Read(configurationFile)
	if err != nil {
		panic(err.Error())
	} else {
		Conf = cfg
	}
}

// LoadRepoConfig loads the configuration for the solder compatible mod repository
func LoadRepoConfig() {
	section, err := Conf.Section("repo")
	if err != nil {
		panic("No repo config found, currently this is required for reflow to function")
	}

	RepoURL = section.ValueOf("repo_url")
}
