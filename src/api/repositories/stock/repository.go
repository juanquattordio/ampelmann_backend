package stock

import (
	"database/sql"
	goErrors "errors"
	"github.com/jmoiron/sqlx"
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

func (r *Repository) GetStock(idInsumo *int64, idDeposito *int64) (float64, error) {
	if idDeposito == nil {
		dbStock := new(stock)
		rows, err := r.db.Query(sumStockByInsumo, idInsumo)
		if err != nil {
			return 0, err
		}
		for rows.Next() {
			_ = rows.Scan(&dbStock.IdInsumo, &dbStock.Stock)
		}
		return dbStock.Stock, nil

	} else if idInsumo == nil {
		// caso de stock por depósito
	}

	// caso de stock por insumo y deposito
	whereConditions, args := buildSearchWhere(idInsumo, idDeposito)
	stockScript := selectScriptMySQL + whereConditions
	dbStock := new(stock)

	err := r.db.Get(dbStock, stockScript, args...)

	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	//stockResult := dbStock.toEntity()
	return dbStock.Stock, nil

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
