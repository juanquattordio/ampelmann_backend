package dependencies

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/config/db"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_factura_compra"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/get_stock"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/movimiento_depositos"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_historial_precios_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints/handlers/api"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/administracion/documento"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/stock"
)

type HandlerContainer struct {
	CreateCliente            entrypoints.Handler
	SearchCliente            entrypoints.Handler
	UpdateCliente            entrypoints.Handler
	CreateProveedor          entrypoints.Handler
	SearchProveedor          entrypoints.Handler
	UpdateProveedor          entrypoints.Handler
	UpdateHistorial          entrypoints.Handler
	CreateInsumo             entrypoints.Handler
	SearchInsumo             entrypoints.Handler
	UpdateInsumo             entrypoints.Handler
	GetStock                 entrypoints.Handler
	CreateDeposito           entrypoints.Handler
	UpdateDeposito           entrypoints.Handler
	CreateMovimientoDeposito entrypoints.Handler
	CreateFacturaCompra      entrypoints.Handler
}

func Start() *HandlerContainer {

	// Database
	DB := db.StorageDB

	// Repositories
	clienteRepository := cliente.NewRepository(DB)
	insumoRepository := insumo.NewRepository(DB)
	proveedorRepository := proveedor.NewRepository(DB, insumoRepository)
	depositoRepository := deposito.NewRepository(DB)
	documentosRepository := documento.NewRepository(DB)
	stockRepository := stock.NewRepository(DB, documentosRepository)

	// Use Cases
	createClienteUseCase := &create_cliente.Implementation{
		ClienteProvider: clienteRepository,
	}
	searchClienteUseCase := &search_cliente.Implementation{
		ClienteProvider: clienteRepository,
	}
	updateClienteUseCase := &update_cliente.Implementation{
		ClienteProvider: clienteRepository,
	}
	createProveedorUseCase := &create_proveedor.Implementation{
		ProveedorProvider: proveedorRepository,
	}
	searchProveedorUseCase := &search_proveedor.Implementation{
		ProveedorProvider: proveedorRepository,
	}
	updateProveedorUseCase := &update_proveedor.Implementation{
		ProveedorProvider: proveedorRepository,
	}
	updateHistorialPreciosProveedorUseCase := &update_historial_precios_proveedor.Implementation{
		ProveedorProvider: proveedorRepository,
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
	getStockUseCase := &get_stock.Implementation{
		InsumoProvider:   insumoRepository,
		DepositoProvider: depositoRepository,
		StockProvider:    stockRepository,
	}
	createDepositoUseCase := &create_deposito.Implementation{
		DepositoProvider: depositoRepository,
	}
	updateDepositoUseCase := &update_deposito.Implementation{
		DepositoProvider: depositoRepository,
		StockProvider:    stockRepository,
	}
	movimientoDepositoUseCase := &movimiento_depositos.Implementation{
		DepositoProvider: depositoRepository,
		StockProvider:    stockRepository,
	}
	createFacturaCompraUseCase := &create_factura_compra.Implementation{
		ProveedorProvider: proveedorRepository,
		DocumentoProvider: documentosRepository,
		InsumoProvider:    insumoRepository,
		StockProvider:     stockRepository,
	}

	// API handlers
	handlers := HandlerContainer{}
	handlers.CreateCliente = &api.CreateCliente{
		CreateClienteUseCase: createClienteUseCase,
	}
	handlers.SearchCliente = &api.SearchCliente{
		SearchClienteUseCase: searchClienteUseCase,
	}
	handlers.UpdateCliente = &api.UpdateCliente{
		UpdateClienteUseCase: updateClienteUseCase,
	}
	handlers.CreateProveedor = &api.CreateProveedor{
		CreateProveedorUseCase: createProveedorUseCase,
	}
	handlers.SearchProveedor = &api.SearchProveedor{
		SearchProveedorUseCase: searchProveedorUseCase,
	}
	handlers.UpdateProveedor = &api.UpdateProveedor{
		UpdateProveedorUseCase: updateProveedorUseCase,
	}
	handlers.UpdateHistorial = &api.UpdateHistorialPreciosProveedor{
		UpdateHistorialPreciosProveedor: updateHistorialPreciosProveedorUseCase,
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
	handlers.GetStock = &api.GetStock{
		GetStockUseCase: getStockUseCase,
	}
	handlers.CreateDeposito = &api.CreateDeposito{
		CreateDepositoUseCase: createDepositoUseCase,
	}
	handlers.UpdateDeposito = &api.UpdateDeposito{
		UpdateDepositoUseCase: updateDepositoUseCase,
	}
	handlers.CreateMovimientoDeposito = &api.CreateMovimientoDeposito{
		CreateMovimientoDepositoUseCase: movimientoDepositoUseCase,
	}
	handlers.CreateFacturaCompra = &api.CreateFacturaCompra{
		CreateFacturaCompraUseCase: createFacturaCompraUseCase,
	}

	return &handlers
}
