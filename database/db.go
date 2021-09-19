// Package database handles the database singleton
package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Zaprit/Reflow/config"
)

var lock = &sync.Mutex{}

// Internal to this class, is the underlying singleton instance of the DB
var singleInstance *gorm.DB

// GetDBInstance Get the singleton DB instance
func GetDBInstance() *gorm.DB {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			section, err := config.Conf.Section("database")

			if err != nil {
				panic("Missing database config section in reflow.conf")
			}

			var (
				driver = section.ValueOf("driver")
				host   = section.ValueOf("hostname")
				port   = section.ValueOf("port")
				dbname = section.ValueOf("database")
				user   = section.ValueOf("username")
				pass   = section.ValueOf("password")
			)

			switch driver {
			case "postgres":
				dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
					host, user, pass, dbname, port)
				postgresDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
				}

				singleInstance = postgresDB

			case "mysql":
				dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
					user, pass, host, port, dbname)
				db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
				}

				singleInstance = db

			case "sqlite":
				db, err := gorm.Open(sqlite.Open("reflow.db"), &gorm.Config{})

				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
				}

				singleInstance = db
			default:
				panic("No DB driver specified")
			}
		}
	}

	return singleInstance
}
