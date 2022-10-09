package repositories

import "github.com/jmoiron/sqlx"

type BaseRepository struct {
	DBClient *sqlx.DB
	Tx       *sqlx.Tx
}
