package cliente

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "cliente"
)

const saveScriptMySQL = `
	INSERT INTO Cliente (nombre, ubicacion, paginaWeb
	) VALUES (
		:nombre, :ubicacion, :paginaWeb
	) ON DUPLICATE KEY UPDATE
		nombre = VALUES(nombre),
		ubicacion = VALUES(ubicacion),
		paginaWeb = VALUES(paginaWeb)`

type cliente struct {
	ID        int64  `gorm:"id"`
	Nombre    string `gorm:"nombre"`
	Ubicacion string `gorm:"ubicacion"`
	PaginaWeb string `gorm:"pagina_web"`
}

func newEntity(c entities.Cliente) cliente {
	return cliente{
		ID:        c.ID,
		Nombre:    c.Nombre,
		Ubicacion: c.Ubicacion,
		PaginaWeb: c.PaginaWeb,
	}
}

func (dbItem cliente) GetSaveScript() string {
	return saveScriptMySQL
}
