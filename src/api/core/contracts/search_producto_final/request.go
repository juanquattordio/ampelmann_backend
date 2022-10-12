package search_producto_final

import "github.com/juanquattordio/ampelmann_backend/src/api/core/errors"

type Request struct {
	Id          *int64  `form:"id_producto_final" json:"id_producto_final"`
	Descripcion *string `form:"descripcion" json:"descripcion"`
}

func (command *Request) Validate() error {
	if command.Id == nil && *command.Descripcion == "" {
		return errors.NewInternalServer("Parametros de búsqueda no válidos")
	}
	return nil
}
