package deposito

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

const (
	resource = "deposito"
)

const (
	saveScriptMySQL = "INSERT INTO Deposito(descripcion, status)" +
		"VALUES(?, ?) "
	selectScriptMySQL = "SELECT id_deposito, descripcion, status FROM Deposito"
	updateScriptMySQL = "UPDATE Deposito SET descripcion = ?, status = ? WHERE id_deposito = ?"
)

type deposito struct {
	ID          int64  `db:"id_deposito"`
	Descripcion string `db:"descripcion"`
	Status      string `db:"status"`
}

func newEntity(c entities.Deposito) deposito {
	return deposito{
		ID:          c.ID,
		Descripcion: c.Descripcion,
		Status:      c.Status,
	}
}

func (dbItem deposito) toEntity() *entities.Deposito {
	return &entities.Deposito{
		ID:          dbItem.ID,
		Descripcion: dbItem.Descripcion,
		Status:      dbItem.Status,
	}
}
