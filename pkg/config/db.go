package config

type Host struct {
	Master string
	Slave  []string
}

type Database struct {
	DBHost           Host
	DBName           string
	DBUser           string
	DBPassword       string
	DBMaxConnections string
}

type RedisDriver string

type Redis struct {
	Driver RedisDriver

	DBHost     string
	DBPort     string
	DBPassword string
	DBIndex    string

	SentinelDBHosts    string
	SentinelDBPassword string
	SentinelDBService  string
}
