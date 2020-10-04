package config

import "os"

type Config struct {
	DB     *Database
	Redis  *Redis
	Server *Server
	Site   *Site
}

func NewConfig() *Config {
	return &Config{
		DB: &Database{
			DBHost: Host{
				Master: getEnv("DB_HOST", "localhost"),
				Slave: []string{
					getEnv("DB_SLAVE_HOST_1", "localhost"),
					getEnv("DB_SLAVE_HOST_2", "localhost"),
				},
			},
			DBName:     getEnv("DB_NAME", "homestead"),
			DBUser:     getEnv("DB_USER", "homestead"),
			DBPassword: getEnv("DB_PASSWORD", "homestead"),
		},
		Redis: &Redis{
			DBHost:     getEnv("REDIS_HOST", "127.0.0.1"),
			DBPort:     getEnv("REDIS_PORT", "6379"),
			DBPassword: getEnv("REDIS_PASSWORD", ""),
			DBIndex:    getEnv("REDIS_DB", "0"),
		},
		Server: &Server{
			Env:      getEnv("APP_ENV", "development"),
			BindAddr: getEnv("APP_ADDR", "127.0.0.1"),
		},
		Site: &Site{
			ManageUrl: getEnv("MANAGE_URL", "127.0.0.1"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
