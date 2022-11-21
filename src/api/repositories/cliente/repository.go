package cliente

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.Cliente {
	repo := Repository{
		db: db,
	}
	return &repo
}

var LastIdCliente int64

func (r *Repository) Save(cliente entities.Cliente) error {

	stmt, err := r.db.Prepare(saveScriptMySQL)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(cliente.Cuit, cliente.Nombre, cliente.Ubicacion, cliente.Email, cliente.Status)
	if err != nil {
		return err
	}
	insertedId, _ := result.LastInsertId()
	LastIdCliente = insertedId
	return nil
}
func (r *Repository) GetLastID() (int64, error) {
	return LastIdCliente, nil
}
func (r *Repository) Search(id *int64, cuit *string) (*entities.Cliente, error) {
	whereConditions, args := buildSearchWhere(id, cuit)
	searchScript := selectScriptMySQL + whereConditions
	dbCliente := new(Cliente)

	err := r.db.Get(dbCliente, searchScript, args...)

	if err != nil {
		return nil, err
	}

	clienteResult := dbCliente.toEntity()
	return clienteResult, nil
}

func buildSearchWhere(id *int64, cuit *string) (query string, args []interface{}) {
	if id != nil {
		query += " AND idCliente = ?"
		args = append(args, id)
	}
	if cuit != nil && *cuit != "" {
		query += " AND cuit = ?"
		args = append(args, cuit)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}

func (r *Repository) Update(cliente *entities.Cliente) error {
	stmt, err := r.db.Prepare(updateScriptMySQL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cliente.Cuit, cliente.Nombre, cliente.Ubicacion, cliente.Email, cliente.Status, cliente.ID)
	if err != nil {
		return err
	}

	return nil
}
