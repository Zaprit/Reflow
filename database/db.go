// Package database handles the database singleton
package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Zaprit/Reflow/config"
)

var lock = &sync.Mutex{}

// Internal to this class, is the underlying singleton instance of the DB
var singleton *gorm.DB

// GetDBInstance Get the singleton DB instance
func GetDBInstance() *gorm.DB {
	lock.Lock()
	defer lock.Unlock()

	// Check if it's been initialized
	if singleton != nil {
		return singleton
	}

	section, err := config.Conf.Section("database")

	if err != nil {
		panic("Missing database config section in reflow.conf")
	}

	CreateDBInstance(&DBConfig{
		section.ValueOf("driver"),
		section.ValueOf("hostname"),
		section.ValueOf("port"),
		section.ValueOf("database"),
		section.ValueOf("username"),
		section.ValueOf("password"),
	})

	return singleton
}

func CreateDBInstance(dbConfig *DBConfig) *gorm.DB {

	// For tests to work mostly
	if singleton != nil {
		return singleton
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		LoggerConfig,
	)

	switch dbConfig.Driver {
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.DBName, dbConfig.Port)
		postgresDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		singleton = postgresDB

	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
			dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		singleton = db

	case "sqlite":
		var db *gorm.DB

		var err error

		if dbConfig.DBName != "" {
			db, err = gorm.Open(sqlite.Open(dbConfig.DBName), &gorm.Config{Logger: newLogger})
		} else {
			db, err = gorm.Open(sqlite.Open("reflow.db"), &gorm.Config{Logger: newLogger})
		}

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		singleton = db
	default:
		panic("No DB driver specified")
	}

	return singleton
}
