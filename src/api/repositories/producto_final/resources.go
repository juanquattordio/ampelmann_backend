package producto_final

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "insumo"
)

const (
	saveScriptMySQL   = "INSERT INTO Producto_Final(descripcion, unidad_medida, status) VALUES(:descripcion, :unidad_medida, :status)"
	selectScriptMySQL = "SELECT id_producto, descripcion, unidad_medida, status FROM Producto_Final"
	updateScriptMySQL = "UPDATE Producto_Final SET descripcion = :descripcion, unidad_medida = :unidad_medida, status = :status WHERE id_producto = :id_producto"
)

type productoFinal struct {
	ID          int64  `db:"id_producto"`
	Descripcion string `db:"descripcion"`
	Unidad      string `db:"unidad_medida"`
	Status      string `db:"status"`
}

func newEntity(i entities.ProductoFinal) productoFinal {
	return productoFinal{
		ID:          i.Id,
		Descripcion: i.Descripcion,
		Unidad:      i.Unidad,
		Status:      i.Status,
	}
}

func (dbItem productoFinal) toEntity() *entities.ProductoFinal {
	return &entities.ProductoFinal{
		Id:          dbItem.ID,
		Descripcion: dbItem.Descripcion,
		Unidad:      dbItem.Unidad,
		Status:      dbItem.Status,
	}
}
