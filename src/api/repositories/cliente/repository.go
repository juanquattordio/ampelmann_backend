package cliente

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.Cliente {
	repo := repository{
		db: db,
	}
	return &repo
}

var LastIdCliente int64

func (r repository) Save(cliente entities.Cliente) error {

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
func (r repository) GetLastID() (int64, error) {
	return LastIdCliente, nil
}
func (r repository) Search(request search_cliente.Request) (entities.Cliente, error) {
	whereConditions, args := buildSearchWhere(request)
	searchScript := selectScriptMySQL + whereConditions
	dbCliente := new(cliente)

	err := r.db.Get(dbCliente, searchScript, args...)

	if err != nil {
		return entities.Cliente{}, err
	}

	clienteResult := dbCliente.toEntity()
	return clienteResult, nil
}

func buildSearchWhere(filter search_cliente.Request) (query string, args []interface{}) {
	if filter.Id != nil {
		query += " AND idCliente = ?"
		args = append(args, filter.Id)
	}
	if *filter.Cuit != "" {
		query += " AND cuit = ?"
		args = append(args, filter.Cuit)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}
