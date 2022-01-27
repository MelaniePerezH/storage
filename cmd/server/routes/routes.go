package routes

import (
	"database/sql"

	"github.com/extlurosell/meli_bootcamp_go_w3-2/cmd/server/handler"
	"github.com/extlurosell/meli_bootcamp_go_w3-2/internal/product"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildProductRoutes()

}

func (r *router) setGroup() {
	r.rg = r.r.Group("")
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handler.NewProduct(service)

	productsGroup := r.rg.Group("/products")

	productsGroup.POST("/", handler.Create())
	productsGroup.GET("/:name", handler.Get())
	productsGroup.GET("/", handler.GetAll())
	productsGroup.PUT("/", handler.Update())
}
