package entities

type Cliente struct {
	ID        int64
	Cuit      string
	Nombre    string
	Ubicacion string
	PaginaWeb string
}

func NewCliente(cuit string, nombre string, ubicacion string, paginaWeb string) *Cliente {
	cliente := &Cliente{
		Cuit:      cuit,
		Nombre:    nombre,
		Ubicacion: ubicacion,
		PaginaWeb: paginaWeb,
	}

	return cliente
}
