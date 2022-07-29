package service

import (
	models "REVAMPS-PHP-GO/internal/domain/model"
	"REVAMPS-PHP-GO/internal/domain/ports"
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
