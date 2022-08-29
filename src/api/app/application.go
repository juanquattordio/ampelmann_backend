package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/juanquattordio/ampelmann_backend/src/api/infrastructure/dependencies"
)

const port = ":8080"

func Start() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%serror al cargar archivo .env %s\n", "\033[31m", "\033[0m")
	}
	router := gin.Default()

	handlers := dependencies.Start()
	configureMappings(router, handlers)
	_ = router.Run(port)

}
