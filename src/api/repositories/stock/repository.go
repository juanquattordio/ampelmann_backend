package stock

import (
	"database/sql"
	goErrors "errors"
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.Stock {
	repo := Repository{
		db: db,
	}
	return &repo
}

func (r *Repository) GetStockInsumo(idInsumo *int64, idDeposito *int64) (float64, error) {
	if idDeposito == nil {
		dbStock := new(stockInsumo)
		rows, err := r.db.Query(sumStockByInsumo, idInsumo)
		if err != nil {
			return 0, err
		}
		for rows.Next() {
			_ = rows.Scan(&dbStock.IdInsumo, &dbStock.Stock)
		}
		return dbStock.Stock, nil
	}

	// caso de stockInsumo por insumo y deposito
	whereConditions, args := buildSearchWhere(idInsumo, idDeposito)
	stockScript := getStockInsumoDeposito + whereConditions
	dbStock := new(stockInsumo)

	err := r.db.Get(dbStock, stockScript, args...)

	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return dbStock.Stock, nil

}

func (r *Repository) GetStockDeposito(idDeposito *int64) ([]entities.Insumo, error) {

	rows, err := r.db.Query(getStockByDeposito, idDeposito)
	if err != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	var insumos []entities.Insumo
	for rows.Next() {
		insumoDB := new(stockDeposito)
		_ = rows.Scan(&insumoDB.IdDeposito, &insumoDB.IdInsumo, &insumoDB.NombreInsumo, &insumoDB.Stock)
		insumo := insumoDB.toEntity()
		insumos = append(insumos, insumo)
	}

	return insumos, nil

}

func buildSearchWhere(idInsumo *int64, idDeposito *int64) (query string, args []interface{}) {
	if idInsumo != nil {
		query += " AND id_insumo = ?"
		args = append(args, idInsumo)
	}
	if idDeposito != nil {
		query += " AND id_deposito = ?"
		args = append(args, idDeposito)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}
