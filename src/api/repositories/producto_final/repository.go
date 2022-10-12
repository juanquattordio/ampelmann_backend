package producto_final

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.ProductoFinal {
	repo := Repository{
		db: db,
	}
	return &repo
}

var LastIdProductoFinal int64

func (r *Repository) Save(productoFinal entities.ProductoFinal) (int64, error) {
	//newEntity(productoFinal)
	result, err := r.db.NamedExec(saveScriptMySQL, newEntity(productoFinal))
	if err != nil {
		return 0, err
	}
	insertedId, _ := result.LastInsertId()
	LastIdProductoFinal = insertedId
	return insertedId, nil
}

func (r *Repository) Search(id *int64, descripcion *string) (*entities.ProductoFinal, error) {
	whereConditions, args := buildSearchWhere(id, descripcion)
	searchScript := selectScriptMySQL + whereConditions
	dbProductoFinal := new(productoFinal)

	err := r.db.Get(dbProductoFinal, searchScript, args...)

	if err != nil {
		return nil, err
	}

	productoFinalResult := dbProductoFinal.toEntity()
	return productoFinalResult, nil
}
func (r *Repository) Update(productoFinal *entities.ProductoFinal) error {

	_, err := r.db.NamedExec(updateScriptMySQL, newEntity(*productoFinal))
	if err != nil {
		return err
	}

	return nil
}

func buildSearchWhere(id *int64, nombre *string) (query string, args []interface{}) {
	if id != nil {
		query += " AND id_producto_final = ?"
		args = append(args, id)
	}
	if nombre != nil && *nombre != "" {
		query += " AND descripcion = ?"
		args = append(args, nombre)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}
