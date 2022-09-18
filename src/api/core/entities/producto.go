package entities

type Producto struct {
	id          int64
	descripcion string
}

func NewProducto(descripcion string) *Producto {
	producto := &Producto{
		descripcion: descripcion,
	}

	return producto
}
func (p *Producto) GetDescripcion() string {
	return p.descripcion
}
