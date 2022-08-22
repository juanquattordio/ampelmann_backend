package usuarios

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FileStoreMock struct {
	FileName string
	Mock     *MockStore
}

type MockStore struct {
	Data          []byte // si sé que sólo voy arecibir []Usuario, data podría ser de ese tipo directamente.
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

// func (fsmock *FileStoreMock) Read(data interface{}) error {
// 	if fsmock.Mock.Err != nil {
// 		return fsmock.Mock.Err
// 	}
// 	fsmock.Mock.ReadWasCalled = true
// 	err := json.Unmarshal(fsmock.Mock.Data, &data) // si no funciona, quitar &
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (fsmock *FileStoreMock) Write(data interface{}) error {
// 	if fsmock.Mock.Err != nil {
// 		return fsmock.Mock.Err
// 	}
// 	encodedData, _ := json.Marshal(data)
// 	fsmock.Mock.Data = encodedData
// 	return nil
// }

func TestRepo_Store(t *testing.T) {

	//arrange
	input := []Usuario{
		{
			Id:            10,
			Nombre:        "Fabio",
			Apellido:      "Quattordio",
			Email:         "juanmaria.quattordio@ml.com",
			Edad:          31,
			Altura:        1.79,
			Activo:        false,
			FechaCreacion: "2022-04-01"}}

	jsonData, _ := json.Marshal(input)
	mockUser := MockStore{
		ReadWasCalled: false,
		Data:          jsonData,
	}
	fileStoreMock := FileStoreMock{
		FileName: "",
		Mock:     &mockUser,
	}
	repoUserTest := NewRepository(&fileStoreMock)
	expected := Usuario{
		Id:            10,
		Nombre:        "Fabio",
		Apellido:      "Quattordio",
		Email:         "juanmaria.quattordio@ml.com",
		Edad:          31,
		Altura:        1.79,
		Activo:        false,
		FechaCreacion: "2022-04-01"}

	//act
	resultado, _ := repoUserTest.Store(input[0].Id, input[0].Nombre, input[0].Apellido, input[0].Email, input[0].Edad, input[0].Altura, input[0].Activo, input[0].FechaCreacion)

	//assert
	assert.Equal(t, expected, resultado)
	assert.True(t, mockUser.ReadWasCalled)
}
