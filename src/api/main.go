package main

import (
	"fmt"
	"github.com/go_web/C2-tt-Estructuras/src/api/config/db"
	"github.com/go_web/C2-tt-Estructuras/src/api/entrypoints/handler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go_web/C2-tt-Estructuras/internal/productos"
	"github.com/joho/godotenv"
)

// func RequestIdMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("X-Request-Id", uuid.New().String())
// 		c.Next()
// 	}
// }
func GetDummyEndpoint(c *gin.Context) {
	resp := map[string]string{"hello": "world"}
	c.JSON(200, resp)
	// c.Next() ver qué pasa si lo agrego en un handler que es el último de una cadena
}
func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")
	// Pass on to the next-in-chain
	c.Next()
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")
	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("api_token")
		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}
		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%serror al cargar archivo .env %s\n", "\033[31m", "\033[0m")
	}
	db := db.StorageDB
	repoProducto := productos.NewRepositoryMySQL(db)
	serviceProducto := productos.NewServiceMySQL(repoProducto)
	handlerProducto := handler.NewProductMySQL(serviceProducto)

	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("/", handlerProducto.GetAll())
	pr.GET("/:id", handlerProducto.GetById())
	pr.GET("", handlerProducto.GetByName()) // ojo con esto porque la ruta no distingue entre /:id y /:name, son todos params strings, y rompe porque ya hay una ruta definida.
	pr.POST("/", handlerProducto.Insertar())
	//pr.PUT("/:id", handlerProducto.Update())
	pr.PUT("/:id", handlerProducto.UpdateWithContext())
	pr.GET("/full", handlerProducto.GetFullData())

	err = r.Run()
	if err != nil {
		log.Panic("El programa no se pudo ejecutar")
	}
}
