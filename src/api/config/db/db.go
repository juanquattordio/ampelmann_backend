package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

// Las funciones llamadas init() son ejecutadas cuando importo cada package. En este caso, como tengo el import en el main, cuando ejecuto main, se ejecuta el init
func init() {
	dataSource := "root:Sendero1659++@tcp(localhost:3306)/ampelmann_bd"
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")
}
