package productos

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go_web/C2-tt-Estructuras/pkg/util"
	"github.com/stretchr/testify/assert"
)

type FileStoreMock struct {
	FileName string
	Mock     *MockStore
}

type MockStore struct {
	Data          []byte
	Err           error
	ReadWasCalled bool
}

func (fsmock *FileStoreMock) Read(data interface{}) error {
	if fsmock.Mock.Err != nil {
		return fsmock.Mock.Err
	}
	fsmock.Mock.ReadWasCalled = true
	fmt.Println("Data con &", &data, " - Data de mock", data)
	err := json.Unmarshal(fsmock.Mock.Data, &data) // si no funciona, quitar &
	if err != nil {
		return err
	}
	return nil
}
func (fsmock *FileStoreMock) Write(data interface{}) error {
	if fsmock.Mock.Err != nil {
		return fsmock.Mock.Err
	}
	encodedData, _ := json.Marshal(data)
	fsmock.Mock.Data = encodedData
	return nil
}

func Test_sqlRepository_Store(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := NewRepositoryMySQL(db)

	product := Product{
		Nombre:         "Vino",
		Color:          "Azul",
		Codigo:         "123",
		Fecha_creacion: "2020-11-12",
		Stock:          12,
		Precio:         15.3,
		Publicado:      true,
	}
	producto, err := repository.Store(product.Nombre, product.Color, product.Codigo, product.Fecha_creacion, product.Stock, product.Precio, product.Publicado)
	fmt.Println("Producto de test", producto)
	assert.NoError(t, err)
	assert.NotZero(t, producto)
	assert.NotEqual(t, "Naranjas", producto.Nombre)
	assert.Equal(t, product.Nombre, producto.Nombre)

}
func Test_sqlRepository_GetOne(t *testing.T) {
	db, err := util.InitDb()
	assert.NoError(t, err)
	repository := NewRepositoryMySQL(db)

	product := Product{
		Nombre:         "Asado",
		Color:          "Rojo",
		Codigo:         "4563213",
		Fecha_creacion: "2020-11-12",
		Stock:          12,
		Precio:         15.3,
		Publicado:      true,
	}
	productoStore, err := repository.Store(product.Nombre, product.Color, product.Codigo, product.Fecha_creacion, product.Stock, product.Precio, product.Publicado)

	productoGetById, errById := repository.GetById(productoStore.Id)

	productoGetByName, errByName := repository.GetByName(productoStore.Nombre)

	assert.NoError(t, errById)
	assert.NotZero(t, productoGetById)
	assert.NotEqual(t, "Vino", productoGetById.Nombre)
	assert.Equal(t, product.Nombre, productoGetById.Nombre)

	assert.NoError(t, errByName)
	assert.NotZero(t, productoGetByName)
	assert.NotEqual(t, "Vino", productoGetByName.Nombre)
	assert.Equal(t, product.Nombre, productoGetByName.Nombre)

	fmt.Println("Id de producto: ", productoGetById.Id, " y ", productoGetByName.Id)
}

func Test_sqlRepository_GetOne_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productId := 1
	columns := []string{"id", "nombre", "color", "precio", "stock", "codigo", "publicado", "fecha_creacion"}

	rows := sqlmock.NewRows(columns)

	rows.AddRow(3, "camisa", "rojo", 12, 10, "654sda", true, "22/05/12")

	mock.ExpectQuery("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from products").WithArgs(productId).WillReturnRows(rows)
	repo := NewRepositoryMySQL(db)

	p, err := repo.GetById(1)

	fmt.Println(*rows)
	assert.Equal(t, "camisa", p.Nombre)

}

// func Test_sqlRepository_GetOne_Mock_WithError(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()
// 	//productId := 1
// 	columns := []string{"id", "nombre", "precio", "stock", "codigo", "publicado", "fecha_creacion"}

// 	rows := sqlmock.NewRows(columns)

// 	rows.AddRow(3, "camisa", "12", 10, "654sda", true, "22/05/12")

// 	mock.ExpectQuery("select id, nombre, color, precio, stock, codigo, publicado, fecha_creacion from products").WillReturnRows(rows)
// 	repo := NewRepositoryMySQL(db)

// 	p, err := repo.GetById(2)

// 	fmt.Println(p)
// 	assert.Error(t, err)

// }

func Test_sqlRepository_Store_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"))
	mock.ExpectPrepare("INSERT INTO products")

	// ver que aca le digo qu elo guarde con el ID 3
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(3, 1))
	productId := 3
	repo := NewRepositoryMySQL(db)
	product := Product{
		Id:             productId,
		Nombre:         "Asado",
		Color:          "Rojo",
		Codigo:         "4563213",
		Fecha_creacion: "2020-11-12",
		Stock:          12,
		Precio:         15.3,
		Publicado:      true,
	}
	p, err := repo.Store(product.Nombre, product.Color, product.Codigo, product.Fecha_creacion, product.Stock, product.Precio, product.Publicado)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, product.Id, p.Id)
	fmt.Println("El Id creado es ", p.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}
