package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
	//MongoType Type = "mongo"
)

type FileStore struct {
	FileName string
	//Mock     *Mock
}

type Mock struct {
	Data          []byte
	Err           error
	ReadWasCalled bool
}

// func (fs *FileStore) AddMock(mock *Mock) {
// 	fs.Mock = mock
// }
// func (fs *FileStore) ClearMock() {
// 	fs.Mock = nil
// }

// type MongoStore struct {
// 	FileName string
// }

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{FileName: fileName}

		//case MongoType:	cuando implemente la lectura/escritura de BD Mongo
		//return &MongoStore{fileName}
	}
	return nil
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Rompe en el write")
	}
	errWrite := os.WriteFile(fs.FileName, fileData, 0644)
	if errWrite != nil {
		return err
	}
	return nil
}

// func (fs *FileStore) Write(data interface{}) error {
// 	if fs.Mock != nil {
// 		if fs.Mock.Err != nil {
// 			return fs.Mock.Err
// 		}
// 		encodedData, _ := json.Marshal(data)
// 		fs.Mock.Data = encodedData
// 		return nil
// 	}
// 	fileData, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		return err
// 	}
// 	f, err := os.OpenFile(fs.FileName, os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = f.Write(fileData)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
func (fs *FileStore) Read(data interface{}) error {
	// if fs.Mock != nil {
	// 	if fs.Mock.Err != nil {
	// 		return fs.Mock.Err
	// 	}
	// 	fs.Mock.ReadWasCalled = true
	// 	fmt.Println("Data con &", &data, " - Data de mock", data)
	// 	return json.Unmarshal(fs.Mock.Data, data) // si no funciona, quitar &
	// }

	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)
}
