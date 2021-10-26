package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Zaprit/Reflow/static"

	"github.com/Zaprit/Reflow/database"

	"github.com/Zaprit/Reflow/technicapi"

	"github.com/Zaprit/Reflow/config"
)

func main() {
	APIName, _ := json.Marshal(config.DefaultInfo)
	fmt.Printf("Reflow %s API: \"%s\"\n", config.DefaultInfo.Version, APIName)

	config.LoadConfig()
	config.LoadRepoConfig()
	database.InitDB()
	static.InitStatic()

	serverConfig, err := config.Conf.Section("server")

	if err != nil {
		log.Panicf("invalid server configuration: %s", err.Error())
	}

	fmt.Printf("Starting server on port %s\n", serverConfig.ValueOf("port"))
	technicapi.StartServer(fmt.Sprintf("%s:%s", serverConfig.ValueOf("listen"), serverConfig.ValueOf("port")))
}
