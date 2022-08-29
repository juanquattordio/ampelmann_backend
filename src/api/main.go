package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/juanquattordio/ampelmann_backend/src/api/app"
)

func main() {
	app.Start()

	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "8080"
	//}
	//
	//if err := run(port); err != nil {
	//	fmt.Printf("error running server", err)
	//}
}

//
//func run(port string) error {
//	//db.Connect()
//
//	r := gin.Default()
//
//	router := routes.NewRouter(r, db.StorageDB)
//	router.MapRoutes()
//
//	if err := r.Run(); err != nil {
//		return err
//	}
//	return nil
//}
