// Package database handles the database singleton
package database

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/alyu/configparser"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Zaprit/Reflow/config"
)

var lock = &sync.Mutex{}

// GormInstance This could probably be improved upon, but it works. and I can not be bothered
type GormInstance struct {
	Instance gorm.DB
}

// Internal to this class, is the underlying singleton instance of the DB
var singleInstance *GormInstance

// GetDBInstance Get the singleton DB instance
func GetDBInstance() *GormInstance {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			configFile, err := configparser.Read(config.ConfigurationFile)
			if err != nil {
				panic("Missing reflow.conf file.\nMaybe you forgot to make a config file from the example?")
			}

			section, err := configFile.Section("database")

			if err != nil {
				panic("Missing database config section in reflow.conf")
			}

			var (
				driver  = section.ValueOf("driver")
				host    = section.ValueOf("hostname")
				port, _ = strconv.ParseInt(section.ValueOf("port"), 10, 16)
				dbname  = section.ValueOf("database")
				user    = section.ValueOf("username")
				pass    = section.ValueOf("password")
			)

			if driver == "postgres" {
				dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + fmt.Sprint(port) + " sslmode=disable"
				postgresDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
				singleInstance = &GormInstance{Instance: *postgresDB}

				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
			}
		}
	}

	return singleInstance
}