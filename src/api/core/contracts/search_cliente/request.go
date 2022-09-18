package search_cliente

type Request struct {
	Id   *int64  `form:"id_cliente" json:"id_cliente"`
	Cuit *string `form:"cuit" json:"cuit"`
}

func (command Request) Validate() error {
	if command.Id == nil && command.Cuit == nil {
		return nil
	}
	return nil
}
