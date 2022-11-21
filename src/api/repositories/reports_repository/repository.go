package reports_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/cliente"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.ReportsProvider {
	repo := Repository{
		db: db,
	}
	return &repo
}

func (r *Repository) GetStockInsumosDesactivados() ([]entities.Insumo, error) {
	var dbInsumos []articulo

	err := r.db.Select(&dbInsumos, getStockInsumosDesactivados)

	if err != nil {
		return nil, err
	}

	insumoResult := toEntitiesInsumos(dbInsumos)
	return insumoResult, nil
}

func (r *Repository) GetStockProductosDesactivados() ([]entities.ProductoFinal, error) {
	var dbInsumos []articulo

	err := r.db.Select(&dbInsumos, getStockProductosDesactivados)

	if err != nil {
		return nil, err
	}

	insumoResult := toEntitiesProductos(dbInsumos)
	return insumoResult, nil
}

func (r *Repository) GetClientesDesactivados() ([]entities.Cliente, error) {
	var dbCliente []cliente.Cliente

	err := r.db.Select(&dbCliente, getClientesDesactivados)

	if err != nil {
		return nil, err
	}

	clienteResult := toEntitiesClientes(dbCliente)
	return clienteResult, nil
}

func (r *Repository) GetFacturacionBetweenDates(dateTo, dateFrom time.Time) (float64, error) {
	//stmt, err := r.db.Prepare(getFacturacionTotalBetweenDates)
	//result, err := stmt.Exec(dateTo, dateFrom)
	//if err != nil {
	//	return 0, errors.NewInternalServer("Fallo al crear documento")
	//}

	row := r.db.QueryRow(getFacturacionTotalBetweenDates, dateTo, dateFrom)
	result := new(float64)
	err := row.Scan(&result)

	return *result, err
}
