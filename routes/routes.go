package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/go-order-microservice/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controllers.Home)
	r.POST("/orders", controllers.NewOrder)
	r.PATCH("/orders/:id", controllers.EditOrder)
	r.GET("/orders/:id/items", controllers.GetAllItemsFromOrder)
	r.GET("/orders/:id", controllers.GetOrderById)
	r.POST("/order_items", controllers.NewOrderItem)
	r.GET("/users/:user_id/orders", controllers.GetAllOrdersByUserId)
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
