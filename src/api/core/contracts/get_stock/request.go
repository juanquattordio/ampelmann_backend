package get_stock

import "github.com/juanquattordio/ampelmann_backend/src/api/core/errors"

type Request struct {
	IdInsumo   *int64 `form:"id_insumo" json:"id_insumo"`
	IdDeposito *int64 `form:"id_deposito" json:"id_deposito"`
}

func (command *Request) Validate() error {
	if command.IdInsumo == nil && command.IdDeposito == nil {
		return errors.NewInternalServer("Parametros de búsqueda no válidos")
	}
	return nil
}
