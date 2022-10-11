package proveedor

import (
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
	"time"
)

type Repository struct {
	db             *sqlx.DB
	insumoProvider providers.Insumo
}

func NewRepository(db *sqlx.DB, insumoProvider providers.Insumo) providers.Proveedor {
	repo := Repository{
		db:             db,
		insumoProvider: insumoProvider,
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
func (r *Repository) UpdateHistorialPrecioInsumo(idProveedor *int64, idInsumo *int64,
	precioUnitario *float64, fecha time.Time, status string) error {

	// valida que exista el proveedor
	proveedorExists, err := r.Search(idProveedor, nil)
	if proveedorExists == nil || goErrors.Is(err, sql.ErrNoRows) {
		return goErrors.New(fmt.Sprintf("id_proveedor %d not found", *idProveedor))
	}
	// valida que existan los insumos comprados
	productoExists, err := r.insumoProvider.Search(idInsumo, nil)
	if err != nil || productoExists == nil || goErrors.Is(err, sql.ErrNoRows) {
		return goErrors.New(fmt.Sprintf("id_insumo %d not found", idInsumo))
	}

	// valida que el status que se quiere insertar sea válido
	if status == "" {
		status = constants.Activo
	} else {
		if status != constants.Desactivo && status != constants.Activo {
			return goErrors.New("status no valido")
		}
	}
	historialToSave := newHistorialPrecioInsumo(idProveedor, idInsumo, precioUnitario, fecha, status)
	if precioUnitario != nil {
		_, err = r.db.NamedExec(insertPriceUpdated, historialToSave)
		if err != nil {
			return err
		}
	} else {
		_, err = r.db.NamedExec(changedStatusHistorialPrice, historialToSave)
		if err != nil {
			return err
		}
	}

	return nil
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
