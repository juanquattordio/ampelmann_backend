package entities

type Cliente struct {
	ID        int64
	Cuit      string
	Nombre    string
	Ubicacion string
	Email     string
	Status    string
}

func NewCliente(cuit string, nombre string, ubicacion string, email string, status string) *Cliente {
	cliente := &Cliente{
		Cuit:      cuit,
		Nombre:    nombre,
		Ubicacion: ubicacion,
		Email:     email,
		Status:    status,
	}

	return cliente
}
