package reports

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type ResponseClientesDesactivados struct {
	ClientesDesactivados []search_cliente.Response `json:"clientes_desactivados"`
}

func NewResponseStockClientesDesactivados(clientes []entities.Cliente) *[]search_cliente.Response {
	clientesList := make([]search_cliente.Response, len(clientes))
	for i := range clientes {
		clientesList[i].ID = fmt.Sprintf("%d", clientes[i].ID)
		clientesList[i].Nombre = clientes[i].Nombre
		clientesList[i].Cuit = clientes[i].Cuit
		clientesList[i].Email = clientes[i].Email
		clientesList[i].Status = clientes[i].Status
	}
	return &clientesList
}
