// This will be a gate of  the data
package ports

import models "REVAMPS-CMI-APPS/internal/entity/model"

type PortDataSource interface {
	Retreive(id string) (models.Account, error) //Retreive By Id or Retreive All, Check if id null then determine if All or by id
	// Store(domain.Account) error
	// Update(domain.Account) error
	// Delete(_id string) error
}
