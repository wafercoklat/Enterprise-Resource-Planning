// It will be out to client (gate of service)
package ports

import models "REVAMPS-CMI-APPS/internal/entity/model"

type PortService interface {
	// AddData(string) error
	// CreateTransaction() error
	// DeleteTransaction() error
	// UpdateTransaction() error
	// ViewAllTransaction() error
	Retreive(id string) (models.Account, error)
}
