package insumo

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.Insumo {
	repo := Repository{
		db: db,
	}
	return &repo
}

var LastIdInsumo int64

func (r *Repository) Save(insumo entities.Insumo) error {

	stmt, err := r.db.Prepare(saveScriptMySQL)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(insumo.Nombre, insumo.Stock, insumo.Status)
	if err != nil {
		return err
	}
	insertedId, _ := result.LastInsertId()
	LastIdInsumo = insertedId
	return nil
}
func (r *Repository) GetLastID() (int64, error) {
	return LastIdInsumo, nil
}
func (r *Repository) Search(id *int64, nombre *string) (*entities.Insumo, error) {
	whereConditions, args := buildSearchWhere(id, nombre)
	searchScript := selectScriptMySQL + whereConditions
	dbInsumo := new(insumo)

	err := r.db.Get(dbInsumo, searchScript, args...)

	if err != nil {
		return nil, err
	}

	insumoResult := dbInsumo.toEntity()
	return insumoResult, nil
}
func (r *Repository) Update(insumo *entities.Insumo) error {

	stmt, err := r.db.Prepare(updateScriptMySQL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(insumo.Nombre, insumo.Stock, insumo.Status, insumo.IdInsumo)
	if err != nil {
		return err
	}

	return nil
}

func buildSearchWhere(id *int64, nombre *string) (query string, args []interface{}) {
	if id != nil {
		query += " AND idInsumo = ?"
		args = append(args, id)
	}
	if nombre != nil && *nombre != "" {
		query += " AND nombre = ?"
		args = append(args, nombre)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}
