package datasource

import (
	models "REVAMPS-CMI-APPS/internal/entity/model"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

// collection get from JSON
type datasource struct {
	// key map[string][]byte
}

func NewData() *datasource {
	return &datasource{
		// key: map[string][]byte{},
	}
}

func (ds *datasource) Retreive(id string) (models.Account, error) {
	// Open Connection
	client, ctx, cancel, err := OpenConn()
	if err != nil {
		panic(err)
	}
	// Close Connection
	defer CloseConn(client, ctx, cancel)

	// Select Database and Collection
	collectSetup := client.Database(os.Getenv("DATABASE")).Collection("accounts")

	episode := models.Account{}
	cursor, err := collectSetup.Find(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		panic(err)
	}

	//here must be error
	if err := cursor.All(ctx, &episode); err != nil {
		return episode, errors.New("fail to get value from collection")
	}
	return episode, nil

	// Check if Retreive All or by Id
	// if ds.id != 0 {
	// Retreive Data By ID

	// }
	// if ds.id == 0 {
	// 	// Retreive Data All
	// 	episode := domain.Account{}
	// 	cursor, err := collectSetup.Find(ctx, bson.D{{}})
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if err := cursor.All(ctx, &episode); err != nil {
	// 		return episode, errors.New("fail to get value from collection")
	// 	}
	// 	return episode, nil
	// }

	// return models.Account{}, errors.New("reterive collection Fail")
}

// 	Store(domain.Account) error
// 	Update(domain.Account) error
// 	Delete(_id string) error
