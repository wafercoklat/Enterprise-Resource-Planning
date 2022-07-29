package service

import (
	"REVAMPS-CMI-APPS/internal/domain/model"
	"REVAMPS-CMI-APPS/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetNotFound(t *testing.T) {
	// Program Mock
	mockAcc_repo := new(mocks.PortRepo)
	// Check fungsi "GET" dengan parameter 1, dan hasil returnnya harus kosong
	mockAcc_repo.On("Get", "1").Return(model.Account{}, nil)

	mockAcc_serv := New(mockAcc_repo)
	result, err := mockAcc_serv.ViewDataById("1")
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "", result.AccountName)
}

func TestService_GetSuccess(t *testing.T) {
	data := model.Account{
		Id:                  1,
		AccountCode:         "ACC-001",
		ParentAccountID:     1,
		AccountName:         "Aktivasi",
		CurrencyID:          "IDR, USD",
		IsDebit:             1,
		AccountType:         1,
		LevelAcc:            1,
		IsDisabled:          1,
		RequireCostCenter:   1,
		AllowAllCostCenters: 1,
	}
	// Program Mock
	mockAcc_repo := new(mocks.PortRepo)
	// Check fungsi "GET" dengan parameter 1, dan hasil returnnya harus kosong
	mockAcc_repo.On("Get", "1").Return(data, nil)

	// Try
	mockAcc_serv := New(mockAcc_repo)
	result, err := mockAcc_serv.ViewDataById("1")

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "Aktivasi", result.AccountName)
}
