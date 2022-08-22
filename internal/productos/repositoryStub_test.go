package productos

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func (s StubStore) Read(data interface{}) error {

	productos := []Product{
		{
			Id:             1,
			Nombre:         "Product 1",
			Color:          "Verde",
			Precio:         15,
			Stock:          10,
			Codigo:         "ASDS456",
			Publicado:      true,
			Fecha_creacion: "15/01/12",
		},
		{
			Id:             2,
			Nombre:         "Product 2",
			Color:          "Rojo",
			Precio:         15,
			Stock:          10,
			Codigo:         "ASDS456",
			Publicado:      true,
			Fecha_creacion: "15/01/12",
		},
	}

	productosEncoded, _ := json.Marshal(productos)

	err := json.Unmarshal(productosEncoded, &data)

	if err != nil {
		return err
	}

	return nil
}

func (s StubStore) Write(data interface{}) error {
	return nil
}

func TestRepo_GetAll(t *testing.T) {

	//arrange
	s := StubStore{}
	repoProductsTest := NewRepository(s)
	service := NewService(repoProductsTest)

	expected := []Product{
		{
			Id:             1,
			Nombre:         "Product 1",
			Color:          "Verde",
			Precio:         15,
			Stock:          10,
			Codigo:         "ASDS456",
			Publicado:      true,
			Fecha_creacion: "15/01/12",
		},
		{
			Id:             2,
			Nombre:         "Product 2",
			Color:          "Rojo",
			Precio:         15,
			Stock:          10,
			Codigo:         "ASDS456",
			Publicado:      true,
			Fecha_creacion: "15/01/12",
		},
	}

	//act
	resultado, _ := service.GetAll()

	//assert
	assert.Equal(t, expected, resultado)

}
