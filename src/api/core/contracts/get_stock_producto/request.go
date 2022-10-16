package get_stock_producto

import "github.com/juanquattordio/ampelmann_backend/src/api/core/errors"

type Request struct {
	IdProducto *int64 `form:"id_producto" json:"id_producto"`
	IdDeposito *int64 `form:"id_deposito" json:"id_deposito"`
}

func (command *Request) Validate() error {
	if command.IdProducto == nil && command.IdDeposito == nil {
		return errors.NewInternalServer("Parametros de búsqueda no válidos")
	}
	return nil
}
