package search_insumo

type Request struct {
	Id     *int64  `form:"id_insumo" json:"id_insumo"`
	Nombre *string `form:"nombre" json:"nombre"`
}

func (command Request) Validate() error {
	if command.Id == nil && command.Nombre == nil {
		return nil
	}
	return nil
}
