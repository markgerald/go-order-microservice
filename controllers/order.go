package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/go-order-microservice/database"
	"github.com/markgerald/go-order-microservice/message"
	"github.com/markgerald/go-order-microservice/models"
	"net/http"
	"strconv"
)

func NewOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	database.DB.Create(&order)
	c.JSON(http.StatusOK, order)
}

func EditOrder(c *gin.Context) {
	var order models.Order
	id := c.Params.ByName("id")
	database.DB.First(&order, id)
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&order).UpdateColumns(order)
	if order.IsPayed == true {
		message.Producer(strconv.FormatUint(uint64(order.ID), 10))
	}
	c.JSON(http.StatusOK, order)
}

func GetOrderById(c *gin.Context) {
	var order models.Order
	database.DB.Preload("OrderItems").First(&order, c.Params.ByName("id"))
	if order.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "No Order"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetAllOrdersByUserId(c *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders, "user_id = ?", c.Params.ByName("user_id"))
	c.JSON(200, orders)
}
