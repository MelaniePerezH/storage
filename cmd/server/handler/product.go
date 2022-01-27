package handler

import (
	"github.com/extlurosell/meli_bootcamp_go_w3-2/internal/domain"
	"github.com/extlurosell/meli_bootcamp_go_w3-2/internal/product"
	"github.com/extlurosell/meli_bootcamp_go_w3-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	productService product.Service
}

func NewProduct(productService product.Service) *Product {
	return &Product{
		productService: productService,
	}
}
func (r *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		buyers, err := r.productService.GetAll()
		if err != nil {
			web.Error(c, 400, "No entiendo")
			return
		}
		web.Success(c, 200, buyers)
	}
}

func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		found, err := p.productService.Get(c, name)
		if err != nil {
			web.Error(c, 404, product.ErrNotFound.Error())
			return
		}
		web.Success(c, 200, found)
	}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, err.Error())
			return
		}

		product, err := p.productService.Save(c, req.Name, req.Tipo, req.Count, req.Price)

		if err != nil {
			web.Error(c, 409, err.Error())
			return
		}
		web.Success(c, 201, product)

	}
}
func (e *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		emp, _ := e.productService.Get(c, name)

		if err := c.ShouldBindJSON(&emp); err != nil {
			web.Error(c, 404, "Not found")
			return
		}

		_, err := e.productService.Update(c, emp)
		if err != nil {
			web.Error(c, 404, "No se pudo actualizar")
			return
		}
		web.Success(c, 200, emp)
	}
}
