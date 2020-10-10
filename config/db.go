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
	DBHost     string
	DBPort     string
	DBPassword string
	DBIndex    string
}
