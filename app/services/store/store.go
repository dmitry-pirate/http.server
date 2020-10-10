package store

import (
	"fmt"
	"github.com/basketforcode/http.server/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strconv"
)

//pool of connections to master and slave
type connectionsPool struct {
	master *sqlx.DB
	slave  []*sqlx.DB
}

type Store struct {
	config     *config.Config
	connection *connectionsPool
}

//connect to mysql db
func New(config *config.Config) (*Store, error) {
	store := Store{
		config: config,
		connection: &connectionsPool{
			master: nil,
			slave:  nil,
		},
	}

	var sc []*sqlx.DB
	for _, slave := range config.DB.DBHost.Slave {
		slc, err := store.openConnection(slave)
		if err != nil {
			continue
		}
		sc = append(sc, slc)
	}

	mc, err := store.openConnection(config.DB.DBHost.Master)
	if err != nil {
		return nil, err
	}
	store.connection.master = mc
	store.connection.slave = sc
	return &store, nil
}

//close opened connection
func (s *Store) Close() error {
	if err := s.connection.master.Close(); err != nil {
		return err
	}
	for _, slave := range s.connection.slave {
		if err := slave.Close(); err != nil {
			return err
		}
	}
	return nil
}

//get opened slave connection and return master connection if slave pool is empty
func (s *Store) SlaveConnection() *sqlx.DB {
	ls := len(s.connection.slave)
	if ls == 0 {
		return s.MasterConnection()
	}
	rI := rand.Intn(ls)
	return s.connection.slave[rI]
}

//get opened master connection
func (s *Store) MasterConnection() *sqlx.DB {
	return s.connection.master
}

//get config structure
func (s *Store) Config() *config.Config {
	return s.config
}

//connect to host
func (s *Store) openConnection(host string) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", s.config.DB.DBUser, s.config.DB.DBPassword, host, s.config.DB.DBName)
	c, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	maxConnections, err := strconv.Atoi(s.config.DB.DBMaxConnections)
	if err != nil {
		maxConnections = 50
	}
	c.SetMaxOpenConns(maxConnections)
	return c, nil
}
