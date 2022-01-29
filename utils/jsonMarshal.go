package utils

import (
	"encoding/json"

	"github.com/Zaprit/Reflow/config"
)

func Marshal(v interface{}) []byte {

	var out []byte

	var err error

	if config.ConfigData.Server.Debug {
		out, err = json.MarshalIndent(v, "", "\t")
	} else {
		out, err = json.Marshal(v)
	}
	if err != nil {
		panic(err.Error())
	}
	return out
}
