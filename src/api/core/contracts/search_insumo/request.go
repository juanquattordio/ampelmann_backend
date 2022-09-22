package search_insumo

import "github.com/juanquattordio/ampelmann_backend/src/api/core/errors"

type Request struct {
	Id     *int64  `form:"id_insumo" json:"id_insumo"`
	Nombre *string `form:"nombre" json:"nombre"`
}

func (command *Request) Validate() error {
	if command.Id == nil && *command.Nombre == "" {
		return errors.NewInternalServer("Parametros de búsqueda no válidos")
	}
	return nil
}
