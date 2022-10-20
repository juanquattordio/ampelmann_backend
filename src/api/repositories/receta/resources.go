package receta

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "recetaHeader"
)

const (
	reinsertRecetaHeader     = "INSERT INTO Receta_Header (id_receta, descripcion_paso_paso, id_producto, litros_finales) VALUES(?,?,?,?)"
	insertRecetaHeader       = "INSERT INTO Receta_Header (descripcion_paso_paso, id_producto, litros_finales) VALUES(?,?,?)"
	insertRecetaLine         = "INSERT INTO Receta_Insumo (id_receta, id_insumo, cantidad_insumo, observaciones) VALUES(?,?,?,?)"
	getRecetaDetails         = "SELECT RI.id_receta, descripcion_paso_paso, id_producto, litros_finales, RI.id_insumo, I.unidad_medida, RI.cantidad_insumo, RI.observaciones FROM Receta_Header INNER JOIN Receta_Insumo RI on Receta_Header.id_receta = RI.id_receta INNER JOIN Insumo I on RI.id_insumo = I.idInsumo WHERE RI.id_receta = ?"
	deleteRecetaHeader       = "DELETE FROM Receta_Header WHERE id_receta=?"
	deleteRecetaIngredientes = "DELETE FROM Receta_Insumo WHERE id_receta=?"
)

type recetaHeader struct {
	IdReceta   int64   `db:"id_receta"`
	PasoPaso   string  `db:"descripcion_paso_paso"`
	IdProducto int64   `db:"id_producto"`
	Litros     float64 `db:"litros_finales"`
}

func (dbItem recetaHeader) toEntity() *entities.RecetaHeader {
	return &entities.RecetaHeader{
		IdHeader:        dbItem.IdReceta,
		PasoPaso:        dbItem.PasoPaso,
		IdProductoFinal: &dbItem.IdProducto,
		LitrosFinales:   dbItem.Litros,
		Ingredientes:    nil,
	}
}
