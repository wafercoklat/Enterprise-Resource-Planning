package datasource

import "REVAMPS-PHP-GO/internal/domain/model"

type datasource struct {
}

func NewData() *datasource {
	return &datasource{}
}

func (ds *datasource) Get(id string) (model.Account, error) {
	InitDBConn()
	var acc model.Account
	err := DBClient.Get(&acc, "SELECT * FROM accounts WHERE Id = ?", id)
	if err != nil {
		panic(err)
	}

	defer DBClient.Close()

	return acc, nil
}
