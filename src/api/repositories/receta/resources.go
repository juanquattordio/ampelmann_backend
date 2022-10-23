package receta

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "recetaHeader"
)

const (
	reinsertRecetaHeader     = "INSERT INTO Receta_Header (id_receta, descripcion_paso_paso, id_producto, litros_finales) VALUES(?,?,?,?)"
	insertRecetaHeader       = "INSERT INTO Receta_Header (descripcion_paso_paso, id_producto, litros_finales) VALUES(?,?,?)"
	insertRecetaLine         = "INSERT INTO Receta_Insumo (id_receta, id_insumo, cantidad_insumo, observaciones) VALUES(?,?,?,?)"
	getRecetaDetails         = "SELECT RH.id_receta, RH.descripcion_paso_paso, RH.id_producto, RH.litros_finales, RI.id_insumo, I.unidad_medida, RI.cantidad_insumo, RI.observaciones FROM Receta_Header RH INNER JOIN Receta_Insumo RI on RH.id_receta = RI.id_receta INNER JOIN Insumo I on RI.id_insumo = I.idInsumo WHERE RH.id_receta = ?"
	deleteRecetaHeader       = "DELETE FROM Receta_Header WHERE id_receta=?"
	deleteRecetaIngredientes = "DELETE FROM Receta_Insumo WHERE id_receta=?"
)

type recetaSearch struct {
	IdReceta      int64   `db:"id_receta"`
	PasoPaso      string  `db:"descripcion_paso_paso"`
	IdProducto    int64   `db:"id_producto"`
	Litros        float64 `db:"litros_finales"`
	IdInsumo      int64   `db:"id_insumo"`
	UnidadMedida  string  `db:"unidad_medida"`
	Cantidad      float64 `db:"cantidad_insumo"`
	Observaciones string  `db:"observaciones"`
}

func toEntity(insumos []recetaSearch) *entities.RecetaHeader {
	var ingredientes []entities.Ingredientes
	for _, insumo := range insumos {
		ingrediente := insumo.toEntity()
		ingredientes = append(ingredientes, *ingrediente)
	}
	return &entities.RecetaHeader{
		IdHeader:        insumos[0].IdReceta,
		PasoPaso:        insumos[0].PasoPaso,
		IdProductoFinal: &insumos[0].IdProducto,
		LitrosFinales:   insumos[0].Litros,
		Ingredientes:    ingredientes,
	}
}

func (dbItem recetaSearch) toEntity() *entities.Ingredientes {
	return &entities.Ingredientes{
		IdInsumo:      dbItem.IdInsumo,
		UnidadMedida:  dbItem.UnidadMedida,
		Cantidad:      dbItem.Cantidad,
		Observaciones: dbItem.Observaciones,
	}
}
