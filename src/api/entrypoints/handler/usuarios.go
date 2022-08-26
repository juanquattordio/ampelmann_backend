package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go_web/C2-tt-Estructuras/internal/usuarios"
)

type requestUsuario struct {
	Nombre        string  `json:"Nombre"`
	Apellido      string  `json:"Apellido"`
	Email         string  `fjson:"Email"`
	Edad          int     `json:"Edad"`
	Altura        float64 `json:"Altura"`
	Activo        *bool   `json:"Activo"`
	FechaCreacion string  `json:"FechaCreacion"`
}

type UsuarioHandler struct {
	service usuarios.Service
}

func (c *UsuarioHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *UsuarioHandler) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req requestUsuario
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		u, err := c.service.Store(req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, *req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, u)
	}
}

func (c *UsuarioHandler) Update() gin.HandlerFunc {
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

		var req requestUsuario
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El campo nombre es requerido"})
			return
		}

		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"error": "El campo apellido es requerido"})
			return
		}
		if req.Email == "" {
			ctx.JSON(400, gin.H{"error": "El campo email es requerido"})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, gin.H{"error": "El campo edad es requerido"})
			return
		}
		if req.Altura == 0 {
			ctx.JSON(400, gin.H{"error": "El campo altura es requerido"})
			return
		}
		if &req.Activo == nil {
			ctx.JSON(400, gin.H{"error": "El campo activo es requerido"})
			return
		}
		if req.FechaCreacion == "" {
			ctx.JSON(400, gin.H{"error": "El campo fecha de creacion es requerido"})
			return
		}

		usuario, err := c.service.Update(int(id), req.Nombre, req.Apellido, req.Email, req.Edad, req.Altura, *req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()}) // error si no se encuentra el ID
			return
		}
		ctx.JSON(200, usuario)
	}
}

func (c *UsuarioHandler) Delete() gin.HandlerFunc {
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

		usuario, err := c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()}) // error si no se encuentra el ID
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El Usuario %d - %s ha sido eliminado", usuario.Id, usuario.Nombre)})
	}
}

func (c *UsuarioHandler) UpdateLastNameAndAge() gin.HandlerFunc {
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

		var req requestUsuario
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Apellido == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del Usuarioo es requerido"})
			return
		}
		if req.Edad == 0 {
			ctx.JSON(400, gin.H{"error": "El campo edad es requerido"})
			return
		}

		usuario, err := c.service.UpdateNameEdad(int(id), req.Nombre, req.Edad)
		if err != nil {
			if errors.Is(err, usuarios.ERROR_PRODUCT_NOT_FOUND) {
				ctx.JSON(404, gin.H{"error": err.Error()}) // error si no se encuentra el ID, devuelve 404
				return
			}
			ctx.JSON(400, gin.H{"error": err.Error()}) // otro tipo de error, le asigna el 400
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El nombre del Usuario %d se actualizó a %s y su edad a %d", usuario.Id, usuario.Nombre, usuario.Edad)})
	}
}

func NewUser(u usuarios.Service) *UsuarioHandler {
	return &UsuarioHandler{
		service: u,
	}
}
