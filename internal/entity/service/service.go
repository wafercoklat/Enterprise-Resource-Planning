package service

import (
	models "REVAMPS-CMI-APPS/internal/entity/model"
	"REVAMPS-CMI-APPS/internal/entity/ports"
)

type service struct {
	portDataSource ports.PortDataSource
}

func New(portDataSource ports.PortDataSource) *service {
	return &service{
		portDataSource: portDataSource,
	}
}

func (s *service) Retreive(id string) (models.Account, error) {
	data, err := s.portDataSource.Retreive(id)
	if err != nil {
		panic(err)
	}
	return data, nil

}

// func (s *service) VieyByIdTransaction(id int) (domain.Account, error) {
// 	return nil, nil
// }
