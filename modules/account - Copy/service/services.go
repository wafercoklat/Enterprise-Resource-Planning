package account

import (
	domain "STACK-ERP/modules/estimation/domain"
	"STACK-ERP/port"
	"fmt"
	"strconv"
)

type Services struct {
	portRepo port.PortRepo
}

func NewService(port port.PortRepo) *Services {
	return &Services{
		portRepo: port,
	}
}

func (s *Services) FindByID(id string) *domain.Estimation {
	var data domain.Estimation

	// Get Header
	header, err := s.portRepo.FindByID(id, &domain.EstimationHeader{}, domain.Tbl_estimation["header"])
	if err != nil {
		fmt.Println(err)
	}
	data.EstHeader = header.(domain.EstimationHeader)

	if domain.Tbl_estimation["detail"] != "" {
		detail, err := s.portRepo.FindByID(id, &domain.EstimationDetails{}, domain.Tbl_estimation["detail"])
		if err != nil {
			fmt.Println(err)
		}
		data.EstDetails = detail.([]domain.EstimationDetails)
	}
	return &data
}

func (s *Services) EstimationList() *domain.Estimation {
	data, err := s.portRepo.List(&domain.EstList, domain.Tbl_estimation["header"])
	if err != nil {
		fmt.Println(err)
	}
	return data.(*domain.Estimation)
}

func (s *Services) Create(data domain.Estimation) error {
	// Insert data header
	headerID, err := s.portRepo.Create(data.EstHeader, domain.Tbl_estimation["header"], domain.EstColumn["header"], domain.EstValue["header"])

	// Insert data detail
	if domain.Tbl_estimation["detail"] != "" {
		// Set the ID parent
		for i := range data.EstDetails {
			data.EstDetails[i].EstimationID = int(headerID)
		}
		_, err := s.portRepo.Create(data.EstDetails, domain.Tbl_estimation["detail"], domain.EstColumn["detail"], domain.EstValue["detail"])

		return err
	}
	return err
}

func (s *Services) Update(id string, data domain.Estimation) error {
	// Update header
	_, err := s.portRepo.Update(data.EstHeader, id, domain.Tbl_estimation["header"], "ID", domain.EstColVal["header"])

	// Update detail
	for i := range data.EstDetails {
		_, err := s.portRepo.Update(data.EstDetails[i], strconv.Itoa(data.EstDetails[i].EstimationDetailID), domain.Tbl_estimation["detail"], "DetailID", domain.EstColVal["detail"])
		if err != nil {
			fmt.Println(err)
		}
	}

	return err
}

func (s *Services) Delete(id string) error {
	// Delete the header
	var err error
	_, err = s.portRepo.Delete(id, domain.Tbl_estimation["header"], "ID")
	if err != nil {
		fmt.Println("Fail to delete")
	}

	// Delete the detail
	if domain.Tbl_estimation["detail"] != "" {
		_, err = s.portRepo.Delete(id, domain.Tbl_estimation["detail"], "ID")
		if err != nil {
			fmt.Println("Fail to delete the detail")
		}
	}
	return err
}

func (s *Services) DetailDelete(id string) error {
	// Delete the detail only
	_, err := s.portRepo.Delete(id, domain.Tbl_estimation["detail"], "DetailID")
	if err != nil {
		fmt.Println("Fail to delete")
	}
	return err
}
