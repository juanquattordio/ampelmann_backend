package producto_final

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "insumo"
)

const (
	saveScriptMySQL   = "INSERT INTO Producto_Final(descripcion, status) VALUES(:descripcion, :status)"
	selectScriptMySQL = "SELECT id_producto, descripcion, status FROM Producto_Final"
	updateScriptMySQL = "UPDATE Producto_Final SET descripcion = :descripcion, status = :status WHERE id_producto = :id_producto"
)

type productoFinal struct {
	ID          int64  `db:"id_producto"`
	Descripcion string `db:"descripcion"`
	Status      string `db:"status"`
}

func newEntity(i entities.ProductoFinal) productoFinal {
	return productoFinal{
		ID:          i.Id,
		Descripcion: i.Descripcion,
		Status:      i.Status,
	}
}

func (dbItem productoFinal) toEntity() *entities.ProductoFinal {
	return &entities.ProductoFinal{
		Id:          dbItem.ID,
		Descripcion: dbItem.Descripcion,
		Status:      dbItem.Status,
	}
}
