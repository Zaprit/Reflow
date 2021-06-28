package app

import (
	"fmt"
	"sync"

	"github.com/revel/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

type GormInstance struct {
	instance gorm.DB
}

var singleInstance *GormInstance

func getDBInstance() *GormInstance {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creting GormInstance Now")
			c, _ := config.ReadDefault("conf/reflow.cfg")
			var (
				driver, _ = c.String("Database", "driver")
				host, _   = c.String("Database", "hostname")
				port, _   = c.Int("Database", "port")
				dbname, _ = c.String("Database", "database")
				user, _   = c.String("Database", "username")
				pass, _   = c.String("Database", "password")
			)
			if driver == "postgres" {
				dsn := fmt.Sprintf("host=%s user=%s password=$s dbname=$s port=$d sslmode=disable", host, user, pass, dbname, port)
				pgdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
				singleInstance = &GormInstance{instance: *pgdb}
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
			}

		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}
