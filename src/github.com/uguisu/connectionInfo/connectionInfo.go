package connectionInfo

import "fmt"

type IgetConnection interface {
	getConnectionString() string
}

// Database connection information
type ConnectionInfo struct {
	user       string
	pwd        string
	schema     string
	port       string
	serverName string
}

/**
 * Get Connection String
 */
func (cInfo ConnectionInfo) GetConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		cInfo.user,
		cInfo.pwd,
		cInfo.serverName,
		cInfo.port,
		cInfo.schema,
	)
}

/**
 * Create default connection information
 */
func Create() *ConnectionInfo {
	return &ConnectionInfo{
		user:       "root",
		pwd:        "root",
		serverName: "192.168.11.120",
		port:       "3306",
		schema:     "himysql",
	}
}
