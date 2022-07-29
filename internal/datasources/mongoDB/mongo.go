package datasource

import (
<<<<<<< Updated upstream
	"REVAMPS-CMI-APPS/internal/entity/domain"
	"REVAMPS-CMI-APPS/internal/entity/ports"
=======
	models "REVAMPS-CMI-APPS/internal/domain/model"
>>>>>>> Stashed changes
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

// collection get from JSON
// database get from config
type datasource struct {
	portdBase ports.PortDataSource
	// id        string `json:"id,omitempty" bson:"id"`
	// table     string `json:"table,omitempty" bson:"table"`
}

var table = "accounts"

func NewData(portdBase_ ports.PortDataSource) *datasource {
	return &datasource{
		portdBase: portdBase_,
	}
}

<<<<<<< Updated upstream
func (ds *datasource) Retreive(_id int) (domain.Account, error) {
=======
func (ds *datasource) Get(id string) (models.Account, error) {
>>>>>>> Stashed changes
	// Open Connection
	client, ctx, cancel, err := OpenConn()
	if err != nil {
		panic(err)
	}
	// Close Connection
	defer CloseConn(client, ctx, cancel)

	// Select Database and Collection
	collectSetup := client.Database(os.Getenv("DATABASE")).Collection(table)

	// Check if Retreive All or by Id
	if _id != 0 {
		// Retreive Data By ID
		episode := domain.Account{}
		cursor, err := collectSetup.Find(ctx, bson.D{{Key: "_id", Value: _id}})
		if err != nil {
			panic(err)
		}

		//here must be error
		if err := cursor.All(ctx, &episode); err != nil {
			return episode, errors.New("fail to get value from collection")
		}
		return episode, nil
	}
	if _id == 0 {
		// Retreive Data All
		episode := domain.Account{}
		cursor, err := collectSetup.Find(ctx, bson.D{{}})
		if err != nil {
			panic(err)
		}

		if err := cursor.All(ctx, &episode); err != nil {
			return episode, errors.New("fail to get value from collection")
		}
		return episode, nil
	}

	return domain.Account{}, errors.New("reterive collection Fail")
}

// 	Store(domain.Account) error
// 	Update(domain.Account) error
// 	Delete(_id string) error
