package datasource

import (
	models "REVAMPS-CMI-APPS/internal/entity/model"
	"database/sql"
	"encoding/json"
	"errors"
)

type datasource struct {
	datas map[string][]byte
}

func NewData() *datasource {
	return &datasource{
		datas: map[string][]byte{},
	}
}

func (ds *datasource) Retreive(id string) (models.Account, error) {
	if data, ok := ds.datas[id]; ok {
		rows := DBClient.QueryRow("Select * from customers where id = ?", id)
		defer DBClient.Close()

		account := models.Account{}

		if err := rows.Scan(&account); err != sql.ErrNoRows {
			panic(err)
		}

		err := json.Unmarshal(data, &account)

		if err != nil {
			return models.Account{}, errors.New("fail to get value from datas")
		}

		return account, nil
	}

	return models.Account{}, errors.New("data not found")
}
