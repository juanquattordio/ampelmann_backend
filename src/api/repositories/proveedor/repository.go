package proveedor

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.Proveedor {
	repo := Repository{
		db: db,
	}
	return &repo
}

var LastIdProveedor int64

func (r *Repository) Save(proveedor entities.Proveedor) error {

	stmt, err := r.db.Prepare(saveScriptMySQL)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(proveedor.Cuit, proveedor.Nombre, proveedor.Ubicacion, proveedor.PaginaWeb, proveedor.Status)
	if err != nil {
		return err
	}
	insertedId, _ := result.LastInsertId()
	LastIdProveedor = insertedId
	return nil
}
func (r *Repository) GetLastID() (int64, error) {
	return LastIdProveedor, nil
}
func (r *Repository) Search(id *int64, cuit *string) (*entities.Proveedor, error) {
	whereConditions, args := buildSearchWhere(id, cuit)
	searchScript := selectScriptMySQL + whereConditions
	dbProveedor := new(proveedor)

	err := r.db.Get(dbProveedor, searchScript, args...)

	if err != nil {
		return nil, err
	}

	proveedorResult := dbProveedor.toEntity()
	return proveedorResult, nil
}

func buildSearchWhere(id *int64, cuit *string) (query string, args []interface{}) {
	if id != nil {
		query += " AND idProveedor = ?"
		args = append(args, id)
	}
	if cuit != nil && *cuit != "" {
		query += " AND cuit = ?"
		args = append(args, cuit)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}

func (r *Repository) Update(proveedor *entities.Proveedor) error {

	stmt, err := r.db.Prepare(updateScriptMySQL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(proveedor.Cuit, proveedor.Nombre, proveedor.Ubicacion, proveedor.PaginaWeb, proveedor.Status, proveedor.ID)
	if err != nil {
		return err
	}

	return nil
}
