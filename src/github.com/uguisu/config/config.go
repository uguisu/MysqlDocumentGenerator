package config

import (
	"log"

	dbconnect "github.com/uguisu/dbconnect"
)

// connection information object
var _connectionInfo *dbconnect.ConnectionInfo

func GetConnectionInfoObject() *dbconnect.ConnectionInfo {
	return _connectionInfo
}

/**
 * Load database config
 */
func loadDatabaseConfig() <-chan bool {

	loadDatabaseConfigFinished := make(chan bool)

	go func() {

		log.Println("Execute loadDatabaseConfig() ...")

		// Create settings
		_connectionInfo = dbconnect.Create()

		loadDatabaseConfigFinished <- true
		close(loadDatabaseConfigFinished)
	}()

	return loadDatabaseConfigFinished
}

/**
 * Load Config
 */
func LoadConfig() <-chan bool {

	// Declare a channel for loading config
	configFinished := make(chan bool)

	// Load database config
	// TODO not a goroutin
	loadDatabaseConfigFinished := loadDatabaseConfig()

	// waiting settings method return
	go func() {

		<-loadDatabaseConfigFinished

		configFinished <- true
		close(configFinished)
	}()

	return configFinished
}
