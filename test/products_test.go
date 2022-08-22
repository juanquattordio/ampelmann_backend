package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go_web/C2-tt-Estructuras/cmd/server/handler"
	"github.com/go_web/C2-tt-Estructuras/internal/productos"
	"github.com/go_web/C2-tt-Estructuras/pkg/store"
	"github.com/stretchr/testify/assert"
)

func copyFile(src string, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	tmpPath := fmt.Sprintf("tmp_%s", pathDB)
	err := copyFile(pathDB, tmpPath)
	if err != nil {
		panic(err)
	}
	db := store.New(store.FileType, tmpPath)
	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PATCH("/:id", p.UpdateNameAndPrice())
	pr.DELETE("/:id", p.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetProduct_OK(t *testing.T) {
	type producto struct {
		ID             int     `json:"id"`
		Nombre         string  `json:"nombre"`
		Color          string  `json:"color"`
		Precio         float64 `json:"precio"`
		Stock          int     `json:"stock"`
		Codigo         string  `json:"codigo"`
		Publicado      bool    `json:"publicado"`
		Fecha_creacion string  `json:"fecha_creacion"`
	}
	type Response struct {
		Code  string     `json:"code"`
		Data  []producto `json:"data,omitempty"`
		Error string     `json:"error,omitempty"`
	}

	// crear el Server y definir las Rutas
	r := createServer("products.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	//var objRes, en este caso es tipo Response y no []producto.
	var objRes Response

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//assert.True(t, len(objRes) > 0)
	assert.True(t, len(objRes.Data) > 0)
	fmt.Println(objRes.Data)
}

func TestUpdateNameAndPriceProduct(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("products.json")

	// crear Request del tipo PATCH y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPatch, "/products/5", `{
		"nombre": "gin",
		"precio": 12.9
	}`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
func TestDeleteProduct(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("products.json")

	// crear Request del tipo DELETE y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/products/5", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
func TestCreateProduct(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("products.json")

	// crear Request y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/products/", `{
			"nombre": "leche",
			"color": "verde",
			"precio": 7.6,
			"stock": 100,
			"codigo": "ABC123",
			"publicado": true,
			"fecha_creacion": "2022-04-01"
	 }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
