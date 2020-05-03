package config

import "os"

// Config type
type Config struct {
	DB     *Database
	Server *Server
	Site   *Site
}

//Database config
type Database struct {
	DBName     string
	DBUser     string
	DBPassword string
}

//Server config
type Server struct {
	BindAddr string
	LogLevel string
}

type Site struct {
	ManageUrl string
}

//NewConfig create a new config instance
func NewConfig() *Config {
	return &Config{
		DB: &Database{
			DBName:     getEnv("DB_NAME", "stage"),
			DBUser:     getEnv("DB_USER", "stage"),
			DBPassword: getEnv("DB_PASS", "stage"),
		},
		Server: &Server{
			BindAddr: getEnv("BIND_ADDR", "stage"),
			LogLevel: getEnv("LOG_LEVEL", "stage"),
		},
		Site: &Site{
			ManageUrl: getEnv("MANAGE_URL", "http://manage.vpn.test"),
		},
	}
}

//Getting environment param helper
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
