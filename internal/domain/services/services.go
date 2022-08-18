package services

import (
	account "REVAMP-PHP-GO/internal/domain/modules"
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

func (s *Services) AccountFindByID(id string) *interface{} {
	data, err := account.FindById(id, s.port)
	if err != nil {
		panic(err)
	}
	return &data
}

func (s *Services) AccountList() *interface{} {
	data, err := account.List(s.port)
	if err != nil {
		panic(err)
	}
	return &data
}

func (s *Services) AccountCreate(mdl interface{}) error {
	_, err := account.Create(mdl, s.port)
	return err
}

func (s *Services) AccountUpdate(id string, mdl interface{}) error {
	_, err := account.Update(id, mdl, s.port)
	return err
}

func (s *Services) AccountDelete(id string) error {
	_, err := account.Delete(id, s.port)
	return err
}
