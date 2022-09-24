package proveedor

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

const (
	resource = "proveedor"
)

const (
	saveScriptMySQL = "INSERT INTO Proveedor(cuit, nombre, ubicacion, pagina_web, status)" +
		"VALUES(?, ?, ?, ?, ?) "
	selectScriptMySQL = "SELECT idProveedor, cuit, nombre, ubicacion, pagina_web, status FROM Proveedor"
	updateScriptMySQL = "UPDATE Proveedor SET cuit = ?, nombre = ?, ubicacion = ?, pagina_web = ?, status = ? WHERE idProveedor = ?"
)

type proveedor struct {
	ID        int64  `db:"idProveedor"`
	Cuit      string `db:"cuit"`
	Nombre    string `db:"nombre"`
	Ubicacion string `db:"ubicacion"`
	PaginaWeb string `db:"pagina_web"`
	Status    string `db:"status"`
}

func newEntity(c entities.Proveedor) proveedor {
	return proveedor{
		ID:        c.ID,
		Cuit:      c.Cuit,
		Nombre:    c.Nombre,
		Ubicacion: c.Ubicacion,
		PaginaWeb: c.PaginaWeb,
		Status:    c.Status,
	}
}

func (dbItem proveedor) toEntity() *entities.Proveedor {
	return &entities.Proveedor{
		ID:        dbItem.ID,
		Cuit:      dbItem.Cuit,
		Nombre:    dbItem.Nombre,
		Ubicacion: dbItem.Ubicacion,
		PaginaWeb: dbItem.PaginaWeb,
		Status:    dbItem.Status,
	}
}
