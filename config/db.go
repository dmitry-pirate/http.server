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

type Redis struct {
	Driver string

	DBHost     string
	DBPort     string
	DBPassword string
	DBIndex    string

	SentinelDBHosts    string
	SentinelDBPassword string
	SentinelDBService  string
}
