package productos

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/go_web/C2-tt-Estructuras/pkg/store"
)

type Product struct {
	Id             int     `json:"id"`
	Nombre         string  `json:"nombre"`
	Color          string  `json:"color"`
	Precio         float64 `json:"precio"`
	Stock          int     `json:"stock"`
	Codigo         string  `json:"codigo"`
	Publicado      bool    `json:"publicado"`
	Fecha_creacion string  `json:"fecha_creacion"`
}
type Product_Warehouse struct {
	Id                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	Color             string  `json:"color"`
	Precio            float64 `json:"precio"`
	Stock             int     `json:"stock"`
	Codigo            string  `json:"codigo"`
	Publicado         bool    `json:"publicado"`
	Fecha_creacion    string  `json:"fecha_creacion"`
	Warehouse         string  `json:"warehouse"`
	Warehouse_address string  `json:"warehouse_address"`
}

// var lastID int ya no lo hacemos global sino que se obtiene de la BD.

type Repository interface { // está en mayuscula para poder exportarla a los otros archivos.
	GetAll() ([]Product, error)
	Store(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	LastID() (int, error)
	Update(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	Delete(id int) (Product, error)
	UpdateNameAndPrice(id int, nombre string, precio float64) (Product, error)
}

// type Repository interface { // está en mayuscula para poder exportarla a los otros archivos.

// }
type repository struct { // esta es la instancia que se hace privada para desacoplar las capas.
	db store.Store
}

// Estas funciones es usando un json para la persistencia de datos

func (r *repository) GetAll() ([]Product, error) {
	// 	var ps []Product
	// 	err := r.db.Read(&ps)
	// 	if err != nil {
	// 		return []Product{}, err
	// 	}
	return nil, nil
	// 	return ps, nil
}

func (r *repository) LastID() (int, error) {
	// 	var ps []Product
	// 	if err := r.db.Read(&ps); err != nil {
	// 		return 0, err
	// 	}
	// 	if len(ps) == 0 {
	// 		return 0, nil
	// 	}
	// 	return ps[len(ps)-1].Id, nil
	// 	//return lastID, nil
	return 0, nil
}

func (r *repository) Store(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {

	// 	var ps []Product
	// 	err := r.db.Read(&ps)
	// 	if err != nil {
	// 		return Product{}, err
	// 	}
	// 	p := Product{id, nombre, color, precio, stock, codigo, publicado, fecha_creacion}

	// 	ps = append(ps, p)
	// 	if err := r.db.Write(ps); err != nil { //	Cuando ejecutamos el método Read, no recibimos
	// 		//ni controlamos el error en caso que no pueda obtener productos porque el archivo no exista.
	// 		//El método Write se encargará de crearlo.
	// 		return Product{}, fmt.Errorf("Rompe en el write del respository")
	// 	}
	// 	//lastID = p.Id
	// 	return p, nil
	return Product{}, nil
}
func (r *repository) Update(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {
	// 	p := Product{id, nombre, color, precio, stock, codigo, publicado, fecha_creacion}
	// 	var ps []Product
	// 	err := r.db.Read(&ps)
	// 	if err != nil {
	// 		return Product{}, err
	// 	}
	// 	if len(ps) > 0 {
	// 		updated := false
	// 		for i := range ps {
	// 			if ps[i].Id == id {
	// 				p.Id = id
	// 				ps[i] = p
	// 				updated = true
	// 			}
	// 		}

	// 		if !updated {
	// 			return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	// 		} else {
	// 			if err := r.db.Write(ps); err != nil {
	// 				return Product{}, fmt.Errorf("Rompe en el write del respository")
	// 			}
	// 		}
	// 	} else {
	// 		return Product{}, fmt.Errorf("No hay elementos para actualizar")
	// 	}
	// 	return p, nil
	return Product{}, nil
}

func (r *repository) UpdateNameAndPrice(id int, nombre string, precio float64) (Product, error) {
	// 	var p Product
	// 	var ps []Product
	// 	err := r.db.Read(&ps)
	// 	if err != nil {
	// 		return Product{}, err
	// 	}
	// 	if len(ps) > 0 {
	// 		updated := false
	// 		for i := range ps {
	// 			if ps[i].Id == id {
	// 				p = ps[i]
	// 				p.Id = id
	// 				p.Nombre = nombre
	// 				p.Precio = precio
	// 				ps[i] = p
	// 				updated = true
	// 			}
	// 		}
	// 		if !updated {
	// 			return Product{}, fmt.Errorf("Producto %d - %s no encontrado", id, nombre)
	// 		} else {
	// 			if err := r.db.Write(ps); err != nil {
	// 				return Product{}, fmt.Errorf("Rompe en el write del respository")
	// 			}
	// 		}
	// 	} else {
	// 		return Product{}, fmt.Errorf("No hay elementos para actualizar")
	// 	}
	// 	return p, nil
	return Product{}, nil
}

func (r *repository) Delete(id int) (Product, error) {
	// 	p := Product{}
	// 	var index int
	// 	var ps []Product
	// 	err := r.db.Read(&ps)
	// 	if err != nil {
	// 		return Product{}, err
	// 	}
	// 	if len(ps) > 0 {
	// 		deleted := false
	// 		for i := range ps {
	// 			if ps[i].Id == id {
	// 				p = ps[i]
	// 				index = i
	// 				deleted = true
	// 			}
	// 		}
	// 		if !deleted {
	// 			return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	// 		}
	// 	} else {
	// 		return Product{}, fmt.Errorf("No hay elementos para eliminar")
	// 	}
	// 	ps = append(ps[:index], ps[index+1:]...) // aca pisa el slice de Productos extrayendo el eliminado
	// 	if err := r.db.Write(ps); err != nil {
	// 		return Product{}, fmt.Errorf("Rompe en el write del respository")
	// 	}
	// 	return p, nil
	return Product{}, nil
}

// Esta función es usando un json para la persistencia de datos
func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

type repositoryMySQL struct { // esta es la instancia que se hace privada para desacoplar las capas.
	db *sql.DB
}

type RepositoryMySQL interface { // está en mayuscula para poder exportarla a los otros archivos.
	GetAll() ([]Product, error)
	GetByName(name string) (Product, error)
	GetById(id int) (Product, error)
	Store(nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	Update(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	UpdateWithContext(ctx context.Context, id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error)
	GetFullData() ([]Product_Warehouse, error)
}

// Constantes de querys consulta a BD
const (
	qSelectAll    = "select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from products"
	qSelectByName = "select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from products where nombre = ?"
	qSelectById   = "select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from products where id = ?"
	qInsertOne    = "INSERT INTO products(nombre, color, precio, stock, codigo, publicado, fecha_creacion)" +
		"VALUES( ?, ?, ?, ?, ?, ?, ?)"
	qUpdateOneById = "UPDATE products SET nombre = ?, color = ?, precio = ?, stock = ?, codigo = ?, publicado = ?, fecha_creacion = ?" +
		" WHERE id = ?"
	qIJoinProductsWarehouseById = "select p.id, p.nombre, p.color, p.precio, p.stock, p.codigo, p.publicado, p.fecha_creacion, w.name, w.adress FROM products p" +
		" INNER JOIN warehouses w ON p.id_warehouse = w.id"
)

func (r *repositoryMySQL) GetAll() ([]Product, error) {
	var productos []Product
	rows, err := r.db.Query(qSelectAll)
	if err != nil {
		log.Println(err)
		return productos, nil
	}
	var product Product
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.Fecha_creacion); err != nil {
			log.Fatal(err.Error())
			return nil, err
		} else {
			// Ver cómo lo implementa Nacho en la diapositiva, no usa el Else
			productos = append(productos, product)
		}
	}
	return productos, nil
}

func (r *repositoryMySQL) GetByName(name string) (Product, error) {
	var product Product
	rows, err := r.db.Query(qSelectByName, name)
	if err != nil {
		log.Println(err)
		return product, nil
	}
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.Fecha_creacion); err != nil {
			log.Println(err.Error())
			return product, nil
		}
	}
	return product, nil
}
func (r *repositoryMySQL) GetById(id int) (Product, error) {
	var product Product
	rows, err := r.db.Query(qSelectById, id)
	if err != nil {
		log.Println(err)
		return product, nil
	}
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.Fecha_creacion); err != nil {
			log.Println(err.Error())
			return product, nil
		}
	}
	return product, nil
}

func (r *repositoryMySQL) Store(nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {
	stmt, err := r.db.Prepare(qInsertOne)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	//var result sql.Result
	result, err := stmt.Exec(nombre, color, precio, stock, codigo, publicado, fecha_creacion) // retorna un sql.Result y un error
	if err != nil {
		return Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecución obtenemos el Id insertado
	fmt.Println("El id creado es ", insertedId)
	product := Product{int(insertedId), nombre, color, precio, stock, codigo, publicado, fecha_creacion}
	fmt.Println(product)
	return product, nil
}

func (r *repositoryMySQL) Update(id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {
	stmt, err := r.db.Prepare(qUpdateOneById)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria

	_, err = stmt.Exec(nombre, color, precio, stock, codigo, publicado, fecha_creacion, id) // retorna un sql.Result y un error
	if err != nil {
		return Product{}, err
	}
	p := Product{id, nombre, color, precio, stock, codigo, publicado, fecha_creacion}
	return p, nil
}

func (r *repositoryMySQL) UpdateWithContext(ctx context.Context, id int, nombre, color, codigo, fecha_creacion string, stock int, precio float64, publicado bool) (Product, error) {
	stmt, err := r.db.Prepare(qUpdateOneById)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria

	_, err = stmt.ExecContext(ctx, nombre, color, precio, stock, codigo, publicado, fecha_creacion, id) // retorna un sql.Result y un error
	if err != nil {
		return Product{}, err
	}
	p := Product{id, nombre, color, precio, stock, codigo, publicado, fecha_creacion}
	return p, nil
}

func (r *repositoryMySQL) GetFullData() ([]Product_Warehouse, error) {
	var productos []Product_Warehouse
	rows, err := r.db.Query(qIJoinProductsWarehouseById)
	if err != nil {
		log.Println(err)
		return productos, nil
	}
	var product Product_Warehouse
	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Nombre, &product.Color, &product.Precio, &product.Stock, &product.Codigo, &product.Publicado, &product.Fecha_creacion, &product.Warehouse, &product.Warehouse_address); err != nil {
			log.Fatal(err.Error())
			return nil, err
		} else {
			// Ver cómo lo implementa Nacho en la diapositiva, no usa el Else
			productos = append(productos, product)
		}
	}
	return productos, nil
}

func NewRepositoryMySQL(db *sql.DB) RepositoryMySQL {
	return &repositoryMySQL{
		db: db,
	}
}
