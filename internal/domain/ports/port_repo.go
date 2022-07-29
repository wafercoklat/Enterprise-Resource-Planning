package ports

import models "REVAMPS-CMI-APPS/internal/domain/model"

type PortRepo interface {
	Get(id string) (models.Account, error)
}
