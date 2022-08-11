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
	data := account.FindById(id, s.port)
	return &data
}

func (s *Services) AccountList() *interface{} {
	data := account.List(s.port)
	return &data
}

func (s *Services) AccountCreate(mdl interface{}) error {
	data := account.Create(mdl, s.port)
	return data
}

func (s *Services) AccountUpdate(id string, mdl interface{}) error {
	data := account.Update(id, mdl, s.port)
	return data
}

func (s *Services) AccountDelete(id string) error {
	data := account.Delete(id, s.port)
	return data
}
