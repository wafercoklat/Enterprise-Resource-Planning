package account

import (
	domain "STACK-ERP/modules/account/domain"
	"STACK-ERP/port"
	"fmt"
)

type Services struct {
	portRepo port.PortRepo
}

func NewService(port port.PortRepo) *Services {
	return &Services{
		portRepo: port,
	}
}

func (s *Services) FindByID(id string) *domain.Account {
	data, err := s.portRepo.FindByID(id, &domain.Account{}, domain.Tbl_account)
	if err != nil {
		fmt.Println(err)
	}
	return data.(*domain.Account)
}

func (s *Services) List() *[]domain.Account {
	data, err := s.portRepo.List(&domain.Acclist, domain.Tbl_account)
	if err != nil {
		fmt.Println(err)
	}
	return data.(*[]domain.Account)
}

func (s *Services) Create(data domain.Account) error {
	_, err := s.portRepo.Create(data, domain.Tbl_account, domain.Acccolumn, domain.Accvalue)
	return err
}

func (s *Services) Update(id string, data domain.Account) error {
	_, err := s.portRepo.Update(data, id, domain.Tbl_account, "ID", domain.Acccolumnvalue)
	return err
}

func (s *Services) Delete(id string) error {
	_, err := s.portRepo.Delete(id, domain.Tbl_account, "ID")
	return err
}
