package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

// Config ...
type Config interface {
	url() string
}

// DatabaseConfig ...
type DatabaseConfig struct {
	username string
	password string
	host     string
	port     int
	database string
	debug    bool
}

// ServerConfig ...
type ServerConfig struct {
	host  string
	port  int
	debug bool
}

var database *DatabaseConfig
var server *ServerConfig

func init() {
	database = &DatabaseConfig{}
	database.username = gonv.GetStringEnv("USERNAMEDB", "root")
	database.password = gonv.GetStringEnv("PASSWORD", "")
	database.host = gonv.GetStringEnv("HOST", "localhost")
	database.port = gonv.GetIntEnv("PORTDB", 3306)
	database.database = gonv.GetStringEnv("DATABASE", "project_go_web")
	database.debug = gonv.GetBoolEnv("DEBUG", true)

	server = &ServerConfig{}
	server.host = gonv.GetStringEnv("HOST", "localhost")
	server.port = gonv.GetIntEnv("PORT", 3000)
	server.debug = gonv.GetBoolEnv("DEBUG", true)
}

// DirTemplate ...
func DirTemplate() string {
	return "templates/**/*.html"
}

// DirTemplateError ...
func DirTemplateError() string {
	return "templates/error.html"
}

// Debug ...
func Debug() bool {
	return server.debug
}

// URLDatabase ...
func URLDatabase() string {
	return database.url()
}

// URLServer ...
func URLServer() string {
	return server.url()
}

// ServerPort ...
func ServerPort() int {
	return server.port
}

// <username>:<password>@tcp(<host>:<port>)/<database>
func (databaseConfig *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", databaseConfig.username, databaseConfig.password, databaseConfig.host, databaseConfig.port, databaseConfig.database)
}

func (serverConfig *ServerConfig) url() string {
	return fmt.Sprintf("%s:%d", serverConfig.host, serverConfig.port)
}
