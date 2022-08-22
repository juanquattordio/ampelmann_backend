package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go_web/C2-tt-Estructuras/internal/transactions"
)

type requestTransaccion struct {
	CodigoTransaccion string  `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
}

type Transaccion struct {
	service transactions.Service
}

func (c *Transaccion) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		t, err := c.service.GetAll()

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, t)
	}
}

func (c *Transaccion) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req requestTransaccion

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Store(req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)

		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, p)
	}
}

func (c *Transaccion) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req requestTransaccion
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.CodigoTransaccion == "" {
			ctx.JSON(400, gin.H{"error": "El codigo de transaccion es requerido"})
			return
		}
		if req.Moneda == "" {
			ctx.JSON(400, gin.H{"error": "La moneda es requerido"})
			return
		}
		if req.Monto == 0 {
			ctx.JSON(400, gin.H{"error": "El monto es requerido"})
			return
		}
		if req.Emisor == "" {
			ctx.JSON(400, gin.H{"error": "El emisor es requerido"})
			return
		}
		if req.Receptor == "" {
			ctx.JSON(400, gin.H{"error": "El receptor es requerido"})
			return
		}
		if req.FechaTransaccion == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de la transaccion es requerida"})
			return
		}

		t, err := c.service.Update(int(id), req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()}) // error si no se encuentra el ID
			return
		}
		ctx.JSON(200, t)
	}
}

func (c *Transaccion) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		t, err := c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()}) // error si no se encuentra el ID
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("La transaccion %d - codigo %s ha sido eliminado", t.Id, t.CodigoTransaccion)})
	}
}

func (c *Transaccion) UpdateCodeAndAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req requestTransaccion
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.CodigoTransaccion == "" {
			ctx.JSON(400, gin.H{"error": "El codigo de transaccion es requerido"})
			return
		}
		if req.Monto == 0 {
			ctx.JSON(400, gin.H{"error": "El monto es requerida"})
			return
		}

		t, err := c.service.UpdateCodeAndAmount(int(id), req.CodigoTransaccion, req.Monto)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()}) // error si no se encuentra el ID
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El codigo de la transaccion %d se actualizó a %s y su monto a $%.2f", t.Id, t.CodigoTransaccion, t.Monto)})
	}
}

func NewTransaccion(t transactions.Service) *Transaccion {
	return &Transaccion{
		service: t,
	}
}
