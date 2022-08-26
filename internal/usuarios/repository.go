package usuarios

import (
	"fmt"
	"github.com/go_web/C2-tt-Estructuras/src/api/config/store"
)

// type UsuarioRepository struct {
// 	Id            int
// 	Nombre        string  `json:"Nombre"`
// 	Apellido      string  `json:"Apellido"`
// 	Email         string  `fjson:"Email"`
// 	Edad          int     `json:"Edad"`
// 	Altura        float64 `json:"Altura"`
// 	Activo        bool    `json:"Activo"`
// 	FechaCreacion string  `json:"FechaCreacion"`
// }

// var u []UsuarioRepository
// var lastID int
var ERROR_PRODUCT_NOT_FOUND = fmt.Errorf("El elemento buscado no se encontró")

// type Repository interface {
// 	GetAll() ([]UsuarioRepository, error)
// 	Store(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (UsuarioRepository, error)
// 	LastID() (int, error)
// 	Update(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (UsuarioRepository, error)
// 	Delete(id int) (UsuarioRepository, error)
// 	UpdateLastNameAndAge(id int, apellido string, edad int) (UsuarioRepository, error)
// }

// type repository struct{}

// func (r *repository) GetAll() ([]UsuarioRepository, error) {
// 	return u, nil
// }

// func (r *repository) LastID() (int, error) {
// 	return lastID, nil
// }

// func (r *repository) Store(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (UsuarioRepository, error) {
// 	aux := UsuarioRepository{id, Nombre, Apellido, Email, Edad, Altura, Activo, FechaCreacion}
// 	u = append(u, aux)
// 	lastID = aux.Id
// 	return aux, nil
// }

// func (r *repository) Update(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (UsuarioRepository, error) {
// 	usuario := UsuarioRepository{id, Nombre, Apellido, Email, Edad, Altura, Activo, FechaCreacion}
// 	if len(u) > 0 {
// 		updated := false
// 		for i := range u {
// 			if u[i].Id == id {
// 				usuario.Id = id
// 				u[i] = usuario
// 				updated = true
// 			}
// 		}
// 		if !updated {
// 			return UsuarioRepository{}, fmt.Errorf("Usuario %d no encontrado", id)
// 		}
// 	} else {
// 		return UsuarioRepository{}, fmt.Errorf("No hay elementos para actualizar")
// 	}
// 	return usuario, nil
// }

// func (r *repository) Delete(id int) (UsuarioRepository, error) {
// 	usuario := UsuarioRepository{}
// 	var index int
// 	if len(u) > 0 {
// 		deleted := false
// 		for i := range u {
// 			if u[i].Id == id {
// 				usuario = u[i]
// 				index = i
// 				deleted = true
// 			}
// 		}
// 		if !deleted {
// 			return UsuarioRepository{}, fmt.Errorf("Usuario %d no encontrado", id)
// 		}
// 	} else {
// 		return UsuarioRepository{}, fmt.Errorf("No hay elementos para eliminar")
// 	}
// 	u = append(u[:index], u[index+1:]...) // aca pisa el slice de Productos extrayendo el eliminado
// 	return usuario, nil
// }

// func (r *repository) UpdateLastNameAndAge(id int, apellido string, edad int) (UsuarioRepository, error) {
// 	//var usuario Usuario
// 	var index int
// 	if len(u) > 0 {
// 		updated := false
// 		for i := range u {
// 			if u[i].Id == id {
// 				index = i
// 				//usuario = u[i]
// 				//usuario.Id = id
// 				u[i].Apellido = apellido
// 				//usuario.Apellido = apellido
// 				u[i].Edad = edad
// 				//usuario.Edad = edad
// 				//u[i] = usuario
// 				updated = true
// 			}
// 		}
// 		if !updated {
// 			return UsuarioRepository{}, ERROR_PRODUCT_NOT_FOUND
// 		}
// 	} else {
// 		return UsuarioRepository{}, fmt.Errorf("No hay elementos para actualizar")
// 	}
// 	return u[index], nil
// }

// func NewRepository() Repository {
// 	return &repository{}
// }

type Usuario struct {
	Id            int
	Nombre        string  `json:"Nombre"`
	Apellido      string  `json:"Apellido"`
	Email         string  `fjson:"Email"`
	Edad          int     `json:"Edad"`
	Altura        float64 `json:"Altura"`
	Activo        bool    `json:"Activo"`
	FechaCreacion string  `json:"FechaCreacion"`
}

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (Usuario, error)
	LastID() (int, error)
	Update(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (Usuario, error)
	Delete(id int) (Usuario, error)
	UpdateNameEdad(id int, name string, edad int) (Usuario, error)
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Usuario, error) {
	var u []Usuario
	err := r.db.Read(&u)
	if err != nil {
		return []Usuario{}, err
	}
	return u, nil
}
func (r *repository) LastID() (int, error) {
	u, err := r.GetAll()
	if err != nil {
		return 0, err
	}
	return u[len(u)-1].Id, nil
}

func (r *repository) Store(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (Usuario, error) {
	var u []Usuario
	aux := Usuario{id, Nombre, Apellido, Email, Edad, Altura, Activo, FechaCreacion}
	err := r.db.Read(&u)
	fmt.Println(u)
	if err != nil {
		return Usuario{}, err
	}
	u = append(u, aux)
	err = r.db.Write(u)
	if err != nil {
		return Usuario{}, err
	}
	return aux, nil
}

func (r *repository) Update(id int, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (Usuario, error) {
	var u []Usuario
	aux := Usuario{id, Nombre, Apellido, Email, Edad, Altura, Activo, FechaCreacion}
	updated := false
	err := r.db.Read(&u)
	if err != nil {
		return Usuario{}, err
	}
	for i := range u {
		if u[i].Id == id {
			aux.Id = id
			u[i] = aux
			updated = true
		}
	}
	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	} else {
		if err := r.db.Write(u); err != nil {
			return Usuario{}, fmt.Errorf("Rompe en el write del respository")
		}
	}
	return aux, nil
}

func (r *repository) UpdateNameEdad(id int, Nombre string, Edad int) (Usuario, error) {
	var u []Usuario
	var uAux Usuario
	updated := false
	for i := range u {
		if u[i].Id == id {
			u[i].Nombre = Nombre
			u[i].Edad = Edad
			updated = true
			uAux = u[i]
		}
	}
	if !updated {
		return Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	return uAux, nil

}

func (r *repository) Delete(id int) (Usuario, error) {
	var u []Usuario
	deleted := false
	var index int
	for i := range u {
		if u[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return Usuario{}, fmt.Errorf("Usuario %d no encontrado", id)
	}
	u = append(u[:index], u[index+1:]...)
	return u[index], nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
