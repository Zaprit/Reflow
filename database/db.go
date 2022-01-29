// Package database handles the database singleton
package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Zaprit/Reflow/models"

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

	CreateDBInstance(&DBConfig{
		config.Conf.Database.Driver,
		config.Conf.Database.Hostname,
		config.Conf.Database.Port,
		config.Conf.Database.Database,
		config.Conf.Database.Username,
		config.Conf.Database.Username,
	})

	return singleton
}

// CreateDBInstance functions almost identically to GetDBInstance but lets you specify configuration
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
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable application_name=Reflow",
			dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.DBName, dbConfig.Port)
		postgresDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}

		singleton = postgresDB

	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
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

// InitDB Initializes and migrates the DB ready for use by the server
func InitDB() {
	err := GetDBInstance().AutoMigrate(
		&models.Mod{}, &models.ModVersion{},
		&models.APIKey{}, &models.Modpack{},
		&models.ModpackBuild{}, &models.BuildModversion{})
	if err != nil {
		panic("Failed to migrate tables")
	}
}
