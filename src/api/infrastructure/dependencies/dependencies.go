package dependencies

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/config/db"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints/handlers/api"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/insumo"
)

type HandlerContainer struct {
	CreateCliente entrypoints.Handler
	SearchCliente entrypoints.Handler
	CreateInsumo  entrypoints.Handler
	SearchInsumo  entrypoints.Handler
	UpdateInsumo  entrypoints.Handler
}

func Start() *HandlerContainer {

	// Database
	db := db.StorageDB

	// Repositories
	clienteRepository := cliente.NewRepository(db)
	insumoRepository := insumo.NewRepository(db)

	// Use Cases
	createClienteUseCase := &create_cliente.Implementation{
		ClienteProvider: clienteRepository,
	}
	searchClienteUseCase := &search_cliente.Implementation{
		ClienteProvider: clienteRepository,
	}
	createInsumoUseCase := &create_insumo.Implementation{
		InsumoProvider: insumoRepository,
	}
	searchInsumoUseCase := &search_insumo.Implementation{
		InsumoProvider: insumoRepository,
	}
	updateInsumoUseCase := &update_insumo.Implementation{
		InsumoProvider: insumoRepository,
	}

	// API handlers
	handlers := HandlerContainer{}
	handlers.CreateCliente = &api.CreateCliente{
		CreateClienteUseCase: createClienteUseCase,
	}
	handlers.SearchCliente = &api.SearchCliente{
		SearchClienteUseCase: searchClienteUseCase,
	}
	handlers.CreateInsumo = &api.CreateInsumo{
		CreateInsumoUseCase: createInsumoUseCase,
	}
	handlers.SearchInsumo = &api.SearchInsumo{
		SearchInsumoUseCase: searchInsumoUseCase,
	}
	handlers.UpdateInsumo = &api.UpdateInsumo{
		UpdateInsumoUseCase: updateInsumoUseCase,
	}

	return &handlers
}
