package main

import (
	"encoding/json"
	"fmt"

	"github.com/Zaprit/Reflow/database"

	"github.com/Zaprit/Reflow/solderAPI"

	"github.com/Zaprit/Reflow/config"
)

func main() {
	APIName, _ := json.Marshal(config.DefaultInfo)
	fmt.Printf("Reflow %s API: \"%s\"\n", config.DefaultInfo.Version, APIName)

	config.LoadConfig()
	database.InitDB()

	fmt.Printf("Starting Reflow server on %s:%d\n", config.ConfigData.Server.Listen, config.ConfigData.Server.Port)
	solderAPI.StartServer(fmt.Sprintf("%s:%d", config.ConfigData.Server.Listen, config.ConfigData.Server.Port))
}
