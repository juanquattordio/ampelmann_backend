package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/juanquattordio/ampelmann_backend/src/api/app"
)

func main() {
	app.Start()
}
