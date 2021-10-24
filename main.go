package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Zaprit/Reflow/technicapi"

	"github.com/Zaprit/Reflow/config"
)

func main() {
	APIName, _ := json.Marshal(config.DefaultInfo)
	fmt.Printf("Reflow %s API: \"%s\"\n", config.DefaultInfo.Version, APIName)

	technicapi.InitConfig()
	technicapi.InitDB()

	serverConfig, err := config.Conf.Section("server")

	if err != nil {
		log.Panicf("invalid server configuration: %s", err.Error())
	}

	fmt.Printf("Starting server on port %s\n", serverConfig.ValueOf("port"))
	technicapi.StartServer(fmt.Sprintf(":%s", serverConfig.ValueOf("port")))
}
