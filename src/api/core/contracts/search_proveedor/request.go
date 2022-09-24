package search_proveedor

import "github.com/juanquattordio/ampelmann_backend/src/api/core/errors"

type Request struct {
	Id   *int64  `form:"id_proveedor" json:"id_proveedor"`
	Cuit *string `form:"cuit" json:"cuit"`
}

func (command *Request) Validate() error {
	if command.Id == nil && *command.Cuit == "" {
		return errors.NewInternalServer("Parametros de búsqueda no válidos")
	}
	return nil
}
