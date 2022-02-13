package utils

import (
	"encoding/json"

	"github.com/Zaprit/Reflow/config"
)

func Marshal(v interface{}) ([]byte, error) {

	var out []byte

	var err error

	if config.Conf.Server.Debug {
		out, err = json.MarshalIndent(v, "", "\t")
	} else {
		out, err = json.Marshal(v)
	}

	if err != nil {
		return nil, err
	} else {
		return out, nil
	}
}
