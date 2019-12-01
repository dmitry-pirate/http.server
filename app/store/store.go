package store

import "vpn_api/app/config"

//Store config
type Store struct {
	config *config.Config
}

//New ...
func New() *Store {
	return &Store{}
}

//Open ..
func (s *Store) Open() error {
	return nil
}

//Close ...
func (s *Store) Close() {

}
