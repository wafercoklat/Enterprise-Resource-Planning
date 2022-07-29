package datasource
<<<<<<< Updated upstream
=======

import (
	models "REVAMPS-CMI-APPS/internal/domain/model"
)

type datasource struct {
}

func NewData() *datasource {
	return &datasource{}
}

func (ds *datasource) Get(id string) (models.Account, error) {
	InitDBConn()
	var acc models.Account
	err := DBClient.Get(&acc, "SELECT * FROM accounts WHERE Id = ?", id)
	if err != nil {
		panic(err)
	}

	defer DBClient.Close()

	return acc, nil
}
>>>>>>> Stashed changes
