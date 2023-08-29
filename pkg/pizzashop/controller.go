package pizzashop

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	Db *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := handler{
		Db: db,
	}
	r.Use(FontCacheMiddleware())
	r.Static("/fonts", "./fonts")

	//pizzas
	routes := r.Group("/pizzas")
	routes.GET("/", h.getPizzas)

	//toppings
	routes = r.Group("/toppings")
	routes.GET("/", h.getToppings)
	routes.DELETE("/:id", h.deleteTopping)

	//crusts
	routes = r.Group("/crusts")
	routes.GET("/", h.getCrust)

	//sizes
	routes = r.Group("/sizes")
	routes.GET("/", h.getSizes)

	//customers
	routes = r.Group("/customer")
	routes.GET("/:id", h.getCustomerById)
	routes.POST("/", h.createCustomer)

	routes.GET("/:id/orders", h.getCustomerOrder)

	//orders
	routes = r.Group("/orders")
	routes.POST("/", h.createPizzaOrder)

	routes = r.Group("/checkPrice")
	routes.POST("/", h.getPrice)
}
