package datasource

import (
	"REVAMPS-PHP-GO/internal/domain/model"
	"errors"

	"github.com/stretchr/testify/mock"
)

type Repo_Mock struct {
	Mock mock.Mock
}

func (m *Repo_Mock) Get(id string) (model.Account, error) {
	args := m.Mock.Called(id)
	if args.Get(0) == nil {
		return model.Account{}, errors.New("gagal mocking kosong")
	} else {
		data := args.Get(0).(model.Account)
		return data, nil
	}
}
