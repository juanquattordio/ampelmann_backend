package batch

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.Batch {
	repo := Repository{
		db: db,
	}
	return &repo
}

func (r *Repository) CreateBatch(batch *entities.Batch) error {
	stmt, err := r.db.Prepare(saveScriptMySQL)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(batch.IdReceta, batch.Fecha, batch.LitrosProducidos)
	if err != nil {
		return err
	}

	insertedId, _ := result.LastInsertId()
	batch.IdBatch = insertedId

	return nil
}
func (r *Repository) GetLastBacth() (int64, error) {
	row := r.db.QueryRow(lastBatchId)
	lastIdBatch := new(int64)
	err := row.Scan(&lastIdBatch)
	if err != nil {
		return 0, err
	}

	return *lastIdBatch, nil
}

func (r *Repository) DeleteBatch(tx *sqlx.Tx, idBatch int64) error {
	return nil
}

//func (r *Repository) Search(idBatch *int64, idReceta *int64) (*entities.Batch, error) {
//	whereConditions, args := buildSearchWhere(idBatch, idReceta)
//	searchScript := selectScriptMySQL + whereConditions
//	dbBatch := new(batch)
//
//	err := r.db.Get(dbBatch, searchScript, args...)
//
//	if err != nil {
//		return nil, err
//	}
//
//	batchResult := dbBatch.toEntity()
//	return batchResult, nil
//
//	if idBatch == nil {
//		dbBatch := []batch
//		rows, err := r.db.Query(sumStockByInsumo, idInsumo)
//		if err != nil {
//			return 0, err
//		}
//		for rows.Next() {
//			_ = rows.Scan(&dbStock.IdInsumo, &dbStock.Stock)
//		}
//		return dbStock.Stock, nil
//	}
//
//	// caso de stockInsumo por insumo y deposito
//	whereConditions, args := buildSearchWhere(idInsumo, idDeposito)
//	stockScript := getStockInsumoDeposito + whereConditions
//	dbStock := new(stockInsumo)
//
//	err := r.db.Get(dbStock, stockScript, args...)
//
//	if err != nil {
//		if goErrors.Is(err, sql.ErrNoRows) {
//			return 0, nil
//		}
//		return 0, err
//	}
//
//	return dbStock.Stock, nil
//
//}

func buildSearchWhere(idBatch *int64, idReceta *int64) (query string, args []interface{}) {
	if idBatch != nil {
		query += " AND id_batch = ?"
		args = append(args, idBatch)
	}
	if idReceta != nil {
		query += " AND id_receta = ?"
		args = append(args, idReceta)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}
