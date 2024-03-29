package cliente

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

const (
	resource = "cliente"
)

const (
	saveScriptMySQL = "INSERT INTO Cliente(cuit, nombre, ubicacion, email, status)" +
		"VALUES(?, ?, ?, ?, ?) "
	selectScriptMySQL = "SELECT idCliente, cuit, nombre, ubicacion, email, status FROM Cliente"
	updateScriptMySQL = "UPDATE Cliente SET cuit = ?, nombre = ?, ubicacion = ?, email = ?, status = ? WHERE idCliente = ?"
)

type Cliente struct {
	ID        int64  `db:"idCliente"`
	Cuit      string `db:"cuit"`
	Nombre    string `db:"nombre"`
	Ubicacion string `db:"ubicacion"`
	Email     string `db:"email"`
	Status    string `db:"status"`
}

func newEntity(c entities.Cliente) Cliente {
	return Cliente{
		ID:        c.ID,
		Cuit:      c.Cuit,
		Nombre:    c.Nombre,
		Ubicacion: c.Ubicacion,
		Email:     c.Email,
		Status:    c.Status,
	}
}

func (dbItem Cliente) toEntity() *entities.Cliente {
	return &entities.Cliente{
		ID:        dbItem.ID,
		Cuit:      dbItem.Cuit,
		Nombre:    dbItem.Nombre,
		Ubicacion: dbItem.Ubicacion,
		Email:     dbItem.Email,
		Status:    dbItem.Status,
	}
}
