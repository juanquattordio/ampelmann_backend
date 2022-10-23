package deposito

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.Deposito {
	repo := Repository{
		db: db,
	}
	return &repo
}

var LastIdDeposito int64

func (r *Repository) Save(deposito entities.Deposito) error {
	stmt, err := r.db.Prepare(saveScriptMySQL)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(deposito.Descripcion, deposito.Status)
	if err != nil {
		return err
	}
	insertedId, _ := result.LastInsertId()
	LastIdDeposito = insertedId
	return nil
}
func (r *Repository) GetLastID() (int64, error) {
	return LastIdDeposito, nil
}
func (r *Repository) Search(id *int64, descripcion *string) (*entities.Deposito, error) {
	whereConditions, args := buildSearchWhere(id, descripcion)
	searchScript := selectScriptMySQL + whereConditions
	dbDeposito := new(deposito)

	err := r.db.Get(dbDeposito, searchScript, args...)

	if err != nil {
		return nil, err
	}

	depositoResult := dbDeposito.toEntity()
	return depositoResult, nil
}

func buildSearchWhere(id *int64, descripcion *string) (query string, args []interface{}) {
	if id != nil {
		query += " AND id_deposito = ?"
		args = append(args, id)
	}
	if descripcion != nil && *descripcion != "" {
		query += " AND descripcion = ?"
		args = append(args, descripcion)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}

func (r *Repository) Update(deposito *entities.Deposito) error {
	stmt, err := r.db.Prepare(updateScriptMySQL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(deposito.Descripcion, deposito.Status, deposito.ID)
	if err != nil {
		return err
	}

	return nil
}
