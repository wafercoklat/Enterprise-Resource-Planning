// It will be out to client (gate of service)
package ports

import models "REVAMPS-CMI-APPS/internal/domain/model"

type PortService interface {
	ViewDataById(id string) (models.Account, error)
	// AddData(string) error
	// CreateTransaction() error
	// DeleteTransaction() error
	// UpdateTransaction() error
	// ViewAllTransaction() error
}
