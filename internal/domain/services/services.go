package services

import (
	moduls "REVAMP-PHP-GO/internal/domain/modules"
	"REVAMP-PHP-GO/internal/domain/ports"
)

type Services struct {
	port ports.PortRepo
}

func New(port ports.PortRepo) *Services {
	return &Services{
		port: port,
	}
}

// Account Module
func (s *Services) AccountFindByID(id string) *interface{} {
	data, err := moduls.FindById(id, s.port)
	if err != nil {
		panic(err)
	}
	return &data
}

func (s *Services) AccountList() *interface{} {
	data, err := moduls.List(s.port)
	if err != nil {
		panic(err)
	}
	return &data
}

func (s *Services) AccountCreate(mdl interface{}) error {
	_, err := moduls.Create(mdl, s.port)
	return err
}

func (s *Services) AccountUpdate(id string, mdl interface{}) error {
	_, err := moduls.Update(id, mdl, s.port)
	return err
}

func (s *Services) AccountDelete(id string) error {
	_, err := moduls.Delete(id, s.port)
	return err
}

func (s *Services) UserFindByID(id string) *interface{} {
	data, err := moduls.UserFindById(id, s.port)
	if err != nil {
		panic(err)
	}
	return &data
}

// User Module
func (s *Services) UserList() *interface{} {
	data, err := moduls.UserList(s.port)
	if err != nil {
		panic(err)
	}
	return &data
}

func (s *Services) UserCreate(mdl interface{}) error {
	_, err := moduls.UserCreate(mdl, s.port)
	return err
}

func (s *Services) UserUpdate(id string, mdl interface{}) error {
	_, err := moduls.UserUpdate(id, mdl, s.port)
	return err
}

func (s *Services) UserDelete(id string) error {
	_, err := moduls.UserDelete(id, s.port)
	return err
}
