package Database

import (
	"fmt"
	"github.com/Zaprit/Reflow/src/Config"
	"sync"

	"github.com/revel/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// magic oogaly boogaly, I don't entirely know how this works, found it on the internet and it works
var lock = &sync.Mutex{}

//This could probably be improved upon but it works and I can not be bothered
type GormInstance struct {
	Instance gorm.DB
}



//Internal to this class, is the underlying singleton instance of the DB
var singleInstance *GormInstance

//Get the singleton DB instance
func GetDBInstance() *GormInstance {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			var c *config.Config
			var err error
			for _, confDir := range Config.ConfPaths {

				c, err = config.ReadDefault(confDir + "/db.conf")
				if err == nil {
					break
				}
			}
			if c == nil {
				panic("Missing reflow.conf file.\nMaybe you forgot to make a config file from the example?")
			}
			var (
				driver, _ = c.String("database", "driver")
				host, _   = c.String("database", "hostname")
				port, _   = c.Int("database", "port")
				dbname, _ = c.String("database", "database")
				user, _   = c.String("database", "username")
				pass, _   = c.String("database", "password")
			)

			if driver == "postgres" {

				dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + fmt.Sprint(port) + " sslmode=disable"
				pgdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
				singleInstance = &GormInstance{Instance: *pgdb}
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
			}

		}
	}
	return singleInstance
}