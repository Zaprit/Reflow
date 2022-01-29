package database

import (
	"time"

	"gorm.io/gorm/logger"
)

// DBConfig is an object that contains some db config
type DBConfig struct {
	Driver string
	Host   string
	Port   int
	DBName string
	User   string
	Pass   string
}

// LoggerConfig is the logging config for Gorm
var LoggerConfig = logger.Config{
	SlowThreshold:             time.Second, // Slow SQL threshold
	LogLevel:                  logger.Warn, // Log level
	IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	Colorful:                  true,        // Disable color
}
