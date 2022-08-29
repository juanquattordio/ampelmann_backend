package cliente

import (
	"database/sql"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) providers.Cliente {
	repo := repository{
		db: db,
	}
	return &repo
}

var LastIdCliente int64

func (r repository) Save(cliente entities.Cliente) error {

	//query := saveScriptMySQL
	query := "INSERT INTO Cliente(cuit, nombre, ubicacion, paginaWeb)" +
		"VALUES(?, ?, ?, ?) "

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(cliente.Cuit, cliente.Nombre, cliente.Ubicacion, cliente.PaginaWeb)
	if err != nil {
		return err
	}
	insertedId, _ := result.LastInsertId()
	LastIdCliente = insertedId
	return nil
}
func (r repository) GetLastID() (int64, error) {
	return LastIdCliente, nil
}
