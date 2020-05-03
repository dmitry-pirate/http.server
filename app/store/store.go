package store

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"vpn_api/app/config"
)

//Store config
type Store struct {
	config     *config.Config
	connection *sqlx.DB
}

//New ...
func New(config *config.Config) (*Store, error) {
	connectionString := fmt.Sprintf("%s:%s@/%s?parseTime=true", config.DB.DBUser, config.DB.DBPassword, config.DB.DBName)
	connection, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return &Store{
		config:     config,
		connection: connection,
	}, nil
}

//Close ...
func (s *Store) Close() error {
	return s.connection.Close()
}

func (s *Store) GetConnection() *sqlx.DB {
	return s.connection
}

func (s *Store) GetConfig() *config.Config {
	return s.config
}
