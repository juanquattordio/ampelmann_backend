package dependencies

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/config/db"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints/handlers/api"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/cliente"
)

type HandlerContainer struct {
	CreateCliente entrypoints.Handler
}

func Start() *HandlerContainer {

	// Database
	db := db.StorageDB
	//db, err := database.Connect()

	// Repositories
	clienteRepository := cliente.NewRepository(db)

	// API providers
	//clienteProvider := cliente.New{
	//	db: db,
	//}

	// Use Cases
	createClienteUseCase := create_cliente.NewUseCase(clienteRepository)

	// API handlers
	handlers := HandlerContainer{}
	handlers.CreateCliente = api.NewCreateCliente(createClienteUseCase)

	return &handlers
}
