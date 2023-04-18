package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/go-order-microservice/database"
	"github.com/markgerald/go-order-microservice/models"
	"net/http"
)

func GetAllItemsFromOrder(c *gin.Context) {
	var orderItems []models.OrderItem
	id := c.Params.ByName("id")
	database.DB.Find(&orderItems, "order_id = ?", id)
	c.JSON(200, orderItems)
}

func NewOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	database.DB.Create(&orderItem)
	c.JSON(http.StatusOK, orderItem)
}
