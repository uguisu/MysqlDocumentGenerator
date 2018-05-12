package config

import (
	"log"

	connectionInfo "github.com/uguisu/connectionInfo"
)

// connection information object
var _connectionInfo *connectionInfo.ConnectionInfo

func GetConnectionInfoObject() *connectionInfo.ConnectionInfo {
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
		_connectionInfo = connectionInfo.Create()

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
	loadDatabaseConfigFinished := loadDatabaseConfig()

	// waiting settings method return
	go func() {

		<-loadDatabaseConfigFinished

		configFinished <- true
		close(configFinished)
	}()

	return configFinished
}
