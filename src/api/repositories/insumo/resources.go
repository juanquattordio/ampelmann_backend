package insumo

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "insumo"
)

const (
	saveScriptMySQL = "INSERT INTO Insumo(nombre, unidad_medida, status)" +
		"VALUES(?, ?, ?) "
	selectScriptMySQL = "SELECT idInsumo, nombre, unidad_medida, status FROM Insumo"
	updateScriptMySQL = "UPDATE Insumo SET nombre = ?, unidad_medida = ?, status = ? WHERE idInsumo = ?"
)

type insumo struct {
	ID     int64  `db:"idInsumo"`
	Nombre string `db:"nombre"`
	Unidad string `db:"unidad_medida"`
	Status string `db:"status"`
}

func newEntity(i entities.Insumo) insumo {
	return insumo{
		ID:     i.IdInsumo,
		Nombre: i.Nombre,
		Unidad: i.Unidad,
		Status: i.Status,
	}
}

func (dbItem insumo) toEntity() *entities.Insumo {
	return &entities.Insumo{
		IdInsumo: dbItem.ID,
		Nombre:   dbItem.Nombre,
		Unidad:   dbItem.Unidad,
		Status:   dbItem.Status,
	}
}
