package proveedor

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"time"
)

const (
	resource = "proveedor"
)

const (
	saveScriptMySQL = "INSERT INTO Proveedor(cuit, nombre, ubicacion, pagina_web, status)" +
		"VALUES(?, ?, ?, ?, ?) "
	selectScriptMySQL  = "SELECT idProveedor, cuit, nombre, ubicacion, pagina_web, status FROM Proveedor"
	updateScriptMySQL  = "UPDATE Proveedor SET cuit = ?, nombre = ?, ubicacion = ?, pagina_web = ?, status = ? WHERE idProveedor = ?"
	insertPriceUpdated = `
		INSERT INTO Proveedores_Insumos (
			idProveedor, idInsumo, precio, fecha, status
		) VALUES (
		  	:idProveedor, :idInsumo, :precio, :fecha, :status
		)
		 ON DUPLICATE KEY UPDATE precio = VALUES(precio), fecha = VALUES(fecha), status = VALUES(status)`
	changedStatusHistorialPrice = `
		UPDATE Proveedores_Insumos
		SET status = :status WHERE idProveedor = :idProveedor AND idInsumo = :idInsumo`
)

type proveedor struct {
	ID        int64  `db:"idProveedor"`
	Cuit      string `db:"cuit"`
	Nombre    string `db:"nombre"`
	Ubicacion string `db:"ubicacion"`
	PaginaWeb string `db:"pagina_web"`
	Status    string `db:"status"`
}

func newEntity(c entities.Proveedor) proveedor {
	return proveedor{
		ID:        c.ID,
		Cuit:      c.Cuit,
		Nombre:    c.Nombre,
		Ubicacion: c.Ubicacion,
		PaginaWeb: c.PaginaWeb,
		Status:    c.Status,
	}
}

func (dbItem proveedor) toEntity() *entities.Proveedor {
	return &entities.Proveedor{
		ID:        dbItem.ID,
		Cuit:      dbItem.Cuit,
		Nombre:    dbItem.Nombre,
		Ubicacion: dbItem.Ubicacion,
		PaginaWeb: dbItem.PaginaWeb,
		Status:    dbItem.Status,
	}
}

type precioInsumoProveedor struct {
	IdProveedor int64     `db:"idProveedor"`
	IdInsumo    int64     `db:"idInsumo"`
	Precio      *float64  `db:"precio"`
	Fecha       time.Time `db:"fecha"`
	Status      string    `db:"status"`
}

func newHistorialPrecioInsumo(idProveedor *int64, idInsumo *int64, precioUnitario *float64, fecha time.Time, status string) precioInsumoProveedor {
	return precioInsumoProveedor{
		IdProveedor: *idProveedor,
		IdInsumo:    *idInsumo,
		Precio:      precioUnitario,
		Fecha:       fecha,
		Status:      status,
	}
}
