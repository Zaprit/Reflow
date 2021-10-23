package database

import (
	"gorm.io/gorm/logger"
	"time"
)

// DBConfig is an object that contains some db config
type DBConfig struct {
	Driver string
	Host   string
	Port   string
	DBName string
	User   string
	Pass   string
}

var LoggerConfig = logger.Config{
SlowThreshold:             time.Second,   // Slow SQL threshold
LogLevel:                  logger.Silent, // Log level
IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
Colorful:                  false,         // Disable color
}