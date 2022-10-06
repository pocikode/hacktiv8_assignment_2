package router

import (
	"assignment2/gin/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	order *controller.OrderHandler
}

func (r *Router) CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/orders", r.order.GetOrdersHandler)
	router.POST("/orders", r.order.CreateOrderHandler)
	router.PUT("/orders/:id", r.order.UpdateOrderHandler)
	router.DELETE("/orders/:id", r.order.DeleteOrderHandler)

	return router
}

func NewRouter(order *controller.OrderHandler) *Router {
	return &Router{
		order: order,
	}
}
