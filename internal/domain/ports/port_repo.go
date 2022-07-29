package ports

import models "REVAMPS-PHP-GO/internal/domain/model"

type PortRepo interface {
	Get(id string) (models.Account, error)
}
