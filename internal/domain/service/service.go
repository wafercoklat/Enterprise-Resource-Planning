package service

import (
	models "REVAMPS-CMI-APPS/internal/domain/model"
	"REVAMPS-CMI-APPS/internal/domain/ports"
	"fmt"
)

type Service struct {
	PortRepo ports.PortRepo
}

func New(portRepo ports.PortRepo) *Service {
	return &Service{
		PortRepo: portRepo,
	}
}

func (s *Service) ViewDataById(id string) (models.Account, error) {
	data, err := s.PortRepo.Get(id)
	if err != nil {
		fmt.Printf("Terdapat kesalahan. Err: %s", err)
	}
	return data, nil

}
