package utils

import (
	"encoding/json"
	"github.com/Zaprit/Reflow/config"
)

func Marshal(v interface{}) []byte {

	serverConf, err := config.Conf.Section("server")

	if err != nil {
		panic(err.Error())
	}

	debug := serverConf.ValueOf("debug")

	var out []byte

	if debug == "true" || debug == "yes" {
		out, err = json.MarshalIndent(v, "", "\t")
	} else {
		out, err = json.Marshal(v)
	}
	if err != nil {
		panic(err.Error())
	}
	return out
}