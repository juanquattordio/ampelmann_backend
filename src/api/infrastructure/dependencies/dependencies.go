package dependencies

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/config/db"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_batch"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_factura_compra"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_factura_venta"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/delete_receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/get_stock"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/get_stock_producto"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/movimiento_depositos"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/reports"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_historial_precios_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints/handlers/api"
	"github.com/juanquattordio/ampelmann_backend/src/api/entrypoints/handlers/reports_handler"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/administracion/documento"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/batch"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/reports_repository"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/stock"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/stock_producto"
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
	GetStockInsumo           entrypoints.Handler
	CreateProductoFinal      entrypoints.Handler
	SearchProductoFinal      entrypoints.Handler
	UpdateProductoFinal      entrypoints.Handler
	GetStockProductoFinal    entrypoints.Handler
	CreateDeposito           entrypoints.Handler
	UpdateDeposito           entrypoints.Handler
	CreateMovimientoDeposito entrypoints.Handler
	CreateFacturaCompra      entrypoints.Handler
	CreateFacturaVenta       entrypoints.Handler
	CreateReceta             entrypoints.Handler
	UpdateReceta             entrypoints.Handler
	DeleteReceta             entrypoints.Handler
	CreateBatch              entrypoints.Handler
	InsumosReports           entrypoints.Handler
	ProductosReports         entrypoints.Handler
	ClientesReports          entrypoints.Handler
	VentasReports            entrypoints.Handler
}

func Start() *HandlerContainer {

	// Database
	DB := db.StorageDB

	// Repositories
	clienteRepository := cliente.NewRepository(DB)
	insumoRepository := insumo.NewRepository(DB)
	productoFinalRepository := producto_final.NewRepository(DB)
	proveedorRepository := proveedor.NewRepository(DB, insumoRepository)
	depositoRepository := deposito.NewRepository(DB)
	documentosRepository := documento.NewRepository(DB)
	recetaRepository := receta.NewRepository(DB)
	stockRepository := stock.NewRepository(DB, documentosRepository)
	stockProductoRepository := stock_producto.NewRepository(DB)
	batchRepository := batch.NewRepository(DB)
	reportsRepository := reports_repository.NewRepository(DB)

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
	createProductoFinalUseCase := &create_producto_final.Implementation{
		ProductoFinalProvider: productoFinalRepository,
	}
	searchProductoFinalUseCase := &search_producto_final.Implementation{
		ProductoFinalProvider: productoFinalRepository,
	}
	updateProductoFinalUseCase := &update_producto_final.Implementation{
		ProductoFinalProvider: productoFinalRepository,
	}
	getStockProductoUseCase := &get_stock_producto.Implementation{
		ProductoProvider: productoFinalRepository,
		DepositoProvider: depositoRepository,
		StockProvider:    stockProductoRepository,
	}
	createDepositoUseCase := &create_deposito.Implementation{
		DepositoProvider: depositoRepository,
	}
	updateDepositoUseCase := &update_deposito.Implementation{
		DepositoProvider:      depositoRepository,
		StockProvider:         stockRepository,
		StockProductoProvider: stockProductoRepository,
	}
	movimientoDepositoUseCase := &movimiento_depositos.Implementation{
		DepositoProvider:      depositoRepository,
		StockProvider:         stockRepository,
		StockProductoProvider: stockProductoRepository,
	}
	createFacturaCompraUseCase := &create_factura_compra.Implementation{
		ProveedorProvider: proveedorRepository,
		DocumentoProvider: documentosRepository,
		InsumoProvider:    insumoRepository,
		StockProvider:     stockRepository,
	}
	createFacturaVentaUseCase := &create_factura_venta.Implementation{
		ClienteProvider:           clienteRepository,
		DocumentoProvider:         documentosRepository,
		ProductoProvider:          productoFinalRepository,
		MovimientoDepositoUseCase: movimientoDepositoUseCase,
	}
	createRecetaUseCase := &create_receta.Implementation{
		ProductoProvider: productoFinalRepository,
		InsumoProvider:   insumoRepository,
		RecetaProvider:   recetaRepository,
	}
	updateRecetaUseCase := &update_receta.Implementation{
		RecetaProvider:   recetaRepository,
		ProductoProvider: productoFinalRepository,
		InsumoProvider:   insumoRepository,
	}
	deleteRecetaUseCase := &delete_receta.Implementation{
		RecetaProvider: recetaRepository,
	}
	createBatchUseCase := &create_batch.Implementation{
		BatchProvider:            batchRepository,
		RecetaProvider:           recetaRepository,
		MovimientoInsumosUseCase: movimientoDepositoUseCase,
	}
	reportsUseCase := &reports.Implementation{
		ReportsProvider: reportsRepository,
	}
	//productosReportsUseCase := &reports.Implementation{
	//	ReportsProvider: reportsRepository,
	//}
	//clientesReportsUseCase := &reports.Implementation{
	//	ReportsProvider: reportsRepository,
	//}
	//facturacionReportsUseCase := &reports.Implementation{
	//	ReportsProvider: reportsRepository,
	//}

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
	handlers.GetStockInsumo = &api.GetStockInsumo{
		GetStockUseCase: getStockUseCase,
	}
	handlers.CreateProductoFinal = &api.CreateProductoFinal{
		CreateProductoFinalUseCase: createProductoFinalUseCase,
	}
	handlers.SearchProductoFinal = &api.SearchProductoFinal{
		SearchProductoFinalUseCase: searchProductoFinalUseCase,
	}
	handlers.UpdateProductoFinal = &api.UpdateProductoFinal{
		UpdateProductoFinalUseCase: updateProductoFinalUseCase,
	}
	handlers.GetStockProductoFinal = &api.GetStockProducto{
		GetStockProductoUseCase: getStockProductoUseCase,
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
	handlers.CreateFacturaVenta = &api.CreateFacturaVenta{
		CreateFacturaVentaUseCase: createFacturaVentaUseCase,
	}
	handlers.CreateReceta = &api.CreateReceta{
		CreateRecetaUseCase: createRecetaUseCase,
	}
	handlers.UpdateReceta = &api.UpdateReceta{
		UpdateRecetaUseCase: updateRecetaUseCase,
	}
	handlers.DeleteReceta = &api.DeleteReceta{
		DeleteRecetaUseCase: deleteRecetaUseCase,
	}
	handlers.CreateBatch = &api.CreateBatch{
		CreateBatchUseCase: createBatchUseCase,
	}
	handlers.InsumosReports = &reports_handler.InsumosReports{
		ReportsUseCase: reportsUseCase,
	}
	handlers.ProductosReports = &reports_handler.ProductosReports{
		ReportsUseCase: reportsUseCase,
	}
	handlers.ClientesReports = &reports_handler.ClientesReports{
		ReportsUseCase: reportsUseCase,
	}
	handlers.VentasReports = &reports_handler.FacturacionReports{
		ReportsUseCase: reportsUseCase,
	}

	return &handlers
}
