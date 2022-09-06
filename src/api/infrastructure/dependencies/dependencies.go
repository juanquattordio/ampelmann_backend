package dependencies

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/config/db"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints/handlers/api"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/cliente"
)

type HandlerContainer struct {
	CreateCliente entrypoints.Handler
	SearchCliente entrypoints.Handler
}

func Start() *HandlerContainer {

	// Database
	db := db.StorageDB

	// Repositories
	clienteRepository := cliente.NewRepository(db)

	// Use Cases
	createClienteUseCase := &create_cliente.Implementation{
		ClienteProvider: clienteRepository,
	}
	searchClienteUseCase := &search_cliente.Implementation{
		ClienteProvider: clienteRepository,
	}

	// API handlers
	handlers := HandlerContainer{}
	handlers.CreateCliente = &api.CreateCliente{
		CreateClienteUseCase: createClienteUseCase,
	}
	handlers.SearchCliente = &api.SearchCliente{
		SearchClienteUseCase: searchClienteUseCase,
	}

	return &handlers
}
