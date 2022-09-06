package providers

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Cliente interface {
	Save(cliente entities.Cliente) error
	GetLastID() (int64, error)
	Search(cliente search_cliente.Request) (entities.Cliente, error)
}
