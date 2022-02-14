package main

import (
	"encoding/json"
	"fmt"

	"github.com/Zaprit/Reflow/database"

	"github.com/Zaprit/Reflow/solderapi"

	"github.com/Zaprit/Reflow/config"
)

func main() {
	APIName, _ := json.Marshal(config.DefaultInfo)
	fmt.Printf("Reflow %s API: \"%s\"\n", config.DefaultInfo.Version, APIName)

	config.LoadDefaultConfig()
	database.InitDB()

	fmt.Printf("Starting Reflow server on %s:%d\n", config.Conf.Server.Listen, config.Conf.Server.Port)
	solderapi.StartServer(fmt.Sprintf("%s:%d", config.Conf.Server.Listen, config.Conf.Server.Port))
}
