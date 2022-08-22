package productos

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Se usan las structs de FileStore y MockStore que están en el repository_test

func TestService_GetAll(t *testing.T) {
	//arrange
	input := []Product{
		{
			Id:             1,
			Nombre:         "Cerveza",
			Color:          "Verde",
			Precio:         18,
			Stock:          10,
			Codigo:         "ASDS456",
			Publicado:      true,
			Fecha_creacion: "15/01/12",
		}, {
			Id:             2,
			Nombre:         "Vino",
			Color:          "Verde",
			Precio:         18,
			Stock:          10,
			Codigo:         "ASDS456",
			Publicado:      true,
			Fecha_creacion: "15/01/12",
		},
	}
	dataJson, _ := json.Marshal(input)
	dbMock := MockStore{
		Data:          dataJson,
		ReadWasCalled: false,
	}
	storeMock := FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMock)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, input, result)
	assert.Nil(t, err)

}
func TestService_Update(t *testing.T) {
	expected := Product{
		Id:             1,
		Nombre:         "After Update",
		Color:          "Verde",
		Precio:         18,
		Stock:          10,
		Codigo:         "ASDS456",
		Publicado:      true,
		Fecha_creacion: "15/01/12",
	}
	input := Product{
		Id:             1,
		Nombre:         "Before Update",
		Color:          "Verde",
		Precio:         18,
		Stock:          10,
		Codigo:         "ASDS456",
		Publicado:      true,
		Fecha_creacion: "15/01/12",
	}

	data, _ := json.Marshal([]Product{input})

	dbMock := MockStore{
		Data:          data,
		ReadWasCalled: false,
	}

	storeStub := FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	result, _ := myService.Update(1, "After Update", "Verde", "ASDS456", "15/01/12", 10, 18, true)
	fmt.Println("result ", result)
	assert.Equal(t, expected.Nombre, result.Nombre)
	// assert.Equal(t, expected.Type, result.Type)
	// assert.Equal(t, expected.Price, result.Price)
	assert.Equal(t, expected, result)
	assert.True(t, storeStub.Mock.ReadWasCalled)

}

func TestService_Delete(t *testing.T) {

	input := []Product{
		{Id: 1,
			Nombre:         "After Update",
			Color:          "Verde",
			Precio:         18,
			Stock:          10,
			Codigo:         "ASDS456",
			Publicado:      true,
			Fecha_creacion: "15/01/12",
		},
		{Id: 2,
			Nombre:         "After Update",
			Color:          "Verde",
			Precio:         18,
			Stock:          10,
			Codigo:         "ASDS456",
			Publicado:      true,
			Fecha_creacion: "15/01/12",
		},
		{Id: 3,
			Nombre:         "After Update",
			Color:          "Verde",
			Precio:         18,
			Stock:          10,
			Codigo:         "ASDS456",
			Publicado:      true,
			Fecha_creacion: "15/01/12",
		},
	}

	expected := Product{
		Id:             2,
		Nombre:         "After Update",
		Color:          "Verde",
		Precio:         18,
		Stock:          10,
		Codigo:         "ASDS456",
		Publicado:      true,
		Fecha_creacion: "15/01/12",
	}

	data, _ := json.Marshal(input)

	dbMock := MockStore{
		Data:          data,
		ReadWasCalled: false,
	}

	storeStub := FileStoreMock{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	result, _ := myService.Delete(2)

	fmt.Println("result ", result)
	assert.Equal(t, expected, result)
	assert.True(t, storeStub.Mock.ReadWasCalled)

}
