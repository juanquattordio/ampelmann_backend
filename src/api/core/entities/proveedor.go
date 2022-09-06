package entities

type Proveedor struct {
	ID        int64
	Cuit      string
	Nombre    string
	Ubicacion string
	Email     string
	Status    string
}

func NewProveedor(cuit string, nombre string, ubicacion string, email string, status string) *Proveedor {
	proveedor := &Proveedor{
		Cuit:      cuit,
		Nombre:    nombre,
		Ubicacion: ubicacion,
		Email:     email,
		Status:    status,
	}

	return proveedor
}
