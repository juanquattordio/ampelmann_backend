package providers

type Stock interface {
	GetStock(idInsumo *int64, idDeposito *int64) (float64, error)
	//Update(insumo *entities.Insumo) error
}
