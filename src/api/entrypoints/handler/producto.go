package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go_web/C2-tt-Estructuras/internal/productos"
	"github.com/go_web/C2-tt-Estructuras/src/api/config/web"
)

type request struct {
	Nombre         string  `json:"nombre"`
	Color          string  `json:"color"`
	Precio         float64 `json:"precio"`
	Stock          int     `json:"stock"`
	Codigo         string  `json:"codigo"`
	Publicado      *bool   `json:"publicado"`
	Fecha_creacion string  `json:"fecha_creacion"`
}

type Product struct {
	service productos.Service
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, "error de GetAll"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (prod *Product) GetFilterProductos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type RequestFilter struct {
			//	Id             *int     `json:"id"`
			Nombre *string `form:"nombre" json:"nombre"`
			Color  *string `json:"color" binding:"required"`
			// Precio         *float64 `json:"precio"`
			// Stock          *int     `json:"stock"`
			// Codigo         *string  `json:"codigo"`
			// Publicado      *bool    `json:"publicado"`
			// Fecha_creacion *string  `json:"fecha_creacion"`
		}
		productos, err := prod.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprint("Error en getAll de Filter: ", err.Error()))
		}

		var req RequestFilter
		if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			fmt.Println(&req)
			ctx.JSON(403, web.NewResponse(403, nil, err.Error())) //"Error JUAN en el ShouldBindJSON"))
			return
		}

		productosFiltrados := productos
		//productosFiltrados = nil
		for _, producto := range productos {
			fmt.Println("&req.Nombre", req.Nombre, " - ", *req.Nombre, "producto.Nombre", producto.Nombre, "& - ", &producto.Nombre)
			fmt.Println("&req.Color", req.Color, " - ", *req.Color, "producto.Color", producto.Color, "& - ", &producto.Color)
			if &req.Nombre == nil || producto.Nombre == *req.Nombre &&
				// 		//req.Id == nil || producto.Id == *req.Id { //&& // ver esta diferencia con las demas

				(req.Color == nil || producto.Color == *req.Color) {
				// 		// (req.Precio == nil || &producto.Precio == req.Precio) &&
				// 		// (req.Stock == nil || &producto.Stock == req.Stock) &&
				// 		// (req.Publicado == nil || &producto.Publicado == req.Publicado) &&
				// 		// (req.Fecha_creacion == nil || &producto.Fecha_creacion == req.Fecha_creacion)

				productosFiltrados = append(productosFiltrados, producto)
			}
		}
		if len(productosFiltrados) == 0 {
			ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("Ninguna transaccion cumple el criterio %s", *req.Nombre), ""))
			return
		}
		fmt.Println(productosFiltrados)
		ctx.JSON(200, web.NewResponse(200, productosFiltrados, ""))
	}

}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		// Esto es delegando la validación al ShoulBind
		type requestPost struct {
			Nombre         *string  `form:"nombre" json:"nombre" binding:"required"`
			Color          *string  `form:"color" json:"color" binding:"required"`
			Precio         *float64 `form:"precio" json:"precio" binding:"required"`
			Stock          *int     `form:"stock" json:"stock" binding:"required"`
			Codigo         *string  `form:"codigo" json:"codigo" binding:"required"`
			Publicado      *bool    `form:"publicado" json:"publicado" binding:"required"`
			Fecha_creacion *string  `form:"fecha_creacion" json:"fecha_creacion" binding:"required"`
		}
		var req requestPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		p, err := c.service.Store(*req.Nombre, *req.Color, *req.Codigo, *req.Fecha_creacion, *req.Stock, *req.Precio, *req.Publicado)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "Error al ejecutar Store"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))

		// Esto comentado es haciendo la validación con IF.
		// var req request
		// if err := ctx.ShouldBindJSON(&req); err != nil {
		// 	ctx.JSON(404, web.NewResponse(404, nil, "Error en el ShouldBindJSON"))
		// 	return
		// }

		// if req.Nombre == "" {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
		// 	return
		// }
		// if req.Color == "" {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El color del producto es requerido"))
		// 	return
		// }
		// if req.Precio == 0 {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
		// 	return
		// }
		// if req.Stock == 0 {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El stock del producto es requerido"))
		// 	return
		// }
		// if req.Codigo == "" {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El codigo del producto es requerido"))
		// 	return
		// }
		// if req.Publicado == nil { // ver que acá verifico que el campo no sea nulo, así acepta el false.
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El campo publicado es requerido"))
		// 	return
		// }
		// if req.Fecha_creacion == "" {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El campo fecha de creeacion del producto es requerido"))
		// 	return
		// }

		// p, err := c.service.Store(req.Nombre, req.Color, req.Codigo, req.Fecha_creacion, req.Stock, req.Precio, *req.Publicado)
		// if err != nil {
		// 	ctx.JSON(404, web.NewResponse(404, nil, "Error al ejecutar Store"))
		// 	return
		// }
		// ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}
		type requestPost struct {
			Nombre         *string  `form:"nombre" json:"nombre" binding:"required"`
			Color          *string  `form:"color" json:"color" binding:"required"`
			Precio         *float64 `form:"precio" json:"precio" binding:"required"`
			Stock          *int     `form:"stock" json:"stock" binding:"required"`
			Codigo         *string  `form:"codigo" json:"codigo" binding:"required"`
			Publicado      *bool    `form:"publicado" json:"publicado" binding:"required"`
			Fecha_creacion *string  `form:"fecha_creacion" json:"fecha_creacion" binding:"required"`
		}
		var req requestPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		p, err := c.service.Update(int(id), *req.Nombre, *req.Color, *req.Codigo, *req.Fecha_creacion, *req.Stock, *req.Precio, *req.Publicado)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "No se encontró el ID")) // error si no se encuentra el ID
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))

		// esto lo comento porque delego la verificación al ShouldBind
		// var req request
		// if err := ctx.ShouldBindJSON(&req); err != nil {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "Error en el ShouldBindJSON"))
		// 	return
		// }
		// if req.Nombre == "" {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
		// 	return
		// }
		// if req.Color == "" {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El color del producto es requerido"))
		// 	return
		// }
		// if req.Precio == 0 {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
		// 	return
		// }
		// if req.Stock == 0 {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El stock del producto es requerido"))
		// 	return
		// }
		// if req.Codigo == "" {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El codigo del producto es requerido"))
		// 	return
		// }
		// if req.Publicado == nil { // ver que acá verifico que el campo no sea nulo, así acepta el false.
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El campo publicado es requerido"))
		// 	return
		// }
		// if req.Fecha_creacion == "" {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "El campo fecha de creeacion del producto es requerido"))
		// 	return
		// }

		// id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		// if err != nil {
		// 	ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))
		// 	return
		// }

		// p, err := c.service.Update(int(id), req.Nombre, req.Color, req.Codigo, req.Fecha_creacion, req.Stock, req.Precio, *req.Publicado)
		// if err != nil {
		// 	ctx.JSON(404, web.NewResponse(404, nil, "No se encontró el ID")) // error si no se encuentra el ID
		// 	return
		// }

		// ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id no valido"))
			return
		}

		p, err := c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "No se encontró el ID")) // error si no se encuentra el ID
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
		//ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d - %s ha sido eliminado", p.Id, p.Nombre)})
	}
}

func (c *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id no valido"))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "Error en el ShouldBindJSON"))
			return
		}

		// cambiar estas validaciones usando punteros (cuando corresponda) para que admita que el precio sea 0.
		// o bien, que lea el puntero, y si es nil, que sólo actualice los campos que no sean nil. Entonces, si nos pasa
		// sólo el nombre y no el precio, conservo el precio anterior y sólo actualizo el nombre.
		// si necesito usar punteros en un determinado campo y en el req global no, puedo crearme una struct req dentro de la función
		// y hacer el requiered sólo de ese campo y es válido sólo dentro la función
		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre es requerido"))
			return
		}
		if req.Precio == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio es requerido"))
			return
		}

		p, err := c.service.UpdateNameAndPrice(int(id), req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "No se encontró el ID")) // error si no se encuentra el ID
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
		//ctx.JSON(200, gin.H{"data": fmt.Sprintf("El nombre del producto %d se actualizó a %s y su precio a $%.2f", p.Id, p.Nombre, p.Precio)})
	}
}

func NewProduct(p productos.Service) *Product {
	return &Product{
		service: p,
	}
}

// Handler para implementración de MySQL
type ProductMySQL struct {
	service productos.ServiceMySQL
}

func (c *ProductMySQL) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productos, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "Error de GetAll"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, productos, ""))
	}
}

func (c *ProductMySQL) GetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Query("name")
		p, err := c.service.GetByName(name)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "No se encontró el nombre")) // error si no se encuentra el ID
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *ProductMySQL) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id no valido"))
			return
		}
		p, err := c.service.GetById(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "No se encontró el id")) // error si no se encuentra el ID
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *ProductMySQL) Insertar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type requestPost struct {
			Nombre         *string  `form:"nombre" json:"nombre" binding:"required"`
			Color          *string  `form:"color" json:"color" binding:"required"`
			Precio         *float64 `form:"precio" json:"precio" binding:"required"`
			Stock          *int     `form:"stock" json:"stock" binding:"required"`
			Codigo         *string  `form:"codigo" json:"codigo" binding:"required"`
			Publicado      *bool    `form:"publicado" json:"publicado" binding:"required"`
			Fecha_creacion *string  `form:"fecha_creacion" json:"fecha_creacion" binding:"required"`
		}
		var req requestPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		p, err := c.service.Store(*req.Nombre, *req.Color, *req.Codigo, *req.Fecha_creacion, *req.Stock, *req.Precio, *req.Publicado)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "Error al ejecutar Store"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *ProductMySQL) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type requestPost struct {
			Nombre         *string  `form:"nombre" json:"nombre" binding:"required"`
			Color          *string  `form:"color" json:"color" binding:"required"`
			Precio         *float64 `form:"precio" json:"precio" binding:"required"`
			Stock          *int     `form:"stock" json:"stock" binding:"required"`
			Codigo         *string  `form:"codigo" json:"codigo" binding:"required"`
			Publicado      *bool    `form:"publicado" json:"publicado" binding:"required"`
			Fecha_creacion *string  `form:"fecha_creacion" json:"fecha_creacion" binding:"required"`
		}
		var req requestPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		p, err := c.service.Update(int(id), *req.Nombre, *req.Color, *req.Codigo, *req.Fecha_creacion, *req.Stock, *req.Precio, *req.Publicado)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "No se encontró el ID")) // error si no se encuentra el ID
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *ProductMySQL) UpdateWithContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type requestPost struct {
			Nombre         *string  `form:"nombre" json:"nombre" binding:"required"`
			Color          *string  `form:"color" json:"color" binding:"required"`
			Precio         *float64 `form:"precio" json:"precio" binding:"required"`
			Stock          *int     `form:"stock" json:"stock" binding:"required"`
			Codigo         *string  `form:"codigo" json:"codigo" binding:"required"`
			Publicado      *bool    `form:"publicado" json:"publicado" binding:"required"`
			Fecha_creacion *string  `form:"fecha_creacion" json:"fecha_creacion" binding:"required"`
		}
		var req requestPost
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Id inválido"))
			return
		}

		ctxBD, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		p, err := c.service.UpdateWithContext(ctxBD, int(id), *req.Nombre, *req.Color, *req.Codigo, *req.Fecha_creacion, *req.Stock, *req.Precio, *req.Publicado)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error())) // error si no se encuentra el ID
			//ctx.JSON(404, web.NewResponse(404, nil, "No se encontró el ID")) // error si no se encuentra el ID
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *ProductMySQL) GetFullData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productos, err := c.service.GetFullData()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "Error de GetAll"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, productos, ""))
	}
}

func NewProductMySQL(p productos.ServiceMySQL) *ProductMySQL {
	return &ProductMySQL{
		service: p,
	}
}
