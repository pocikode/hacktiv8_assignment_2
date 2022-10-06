package controller

import (
	"assignment2/gin/repository"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	repo repository.OrderRepository
}

func (o *OrderHandler) GetOrdersHandler(ctx *gin.Context) {
	orders, err := o.repo.GetOrders()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if len(*orders) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Data is empty.",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  orders,
		"total": len(*orders),
	})
}

func (o *OrderHandler) CreateOrderHandler(ctx *gin.Context) {
	var req repository.Order

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = o.repo.CreateOrder(&req)
	if err != nil {
		errMsg := err.Error()
		if err == sql.ErrConnDone {
			errMsg = "DB ERROR: " + errMsg
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errMsg,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created.",
	})
}

func (o *OrderHandler) UpdateOrderHandler(ctx *gin.Context) {
	var req repository.Order

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	orderId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Order ID harus numeric.",
		})
	}

	order, err := o.repo.GetOrderByID(orderId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Order tidak ditemukan.",
		})
	}

	order.CustomerName = req.CustomerName
	order.OrderedAt = req.OrderedAt
	order.Items = req.Items

	err = o.repo.UpdateOrder(order)
	if err != nil {
		errMsg := err.Error()
		if err == sql.ErrConnDone {
			errMsg = "DB ERROR: " + errMsg
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errMsg,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order updated.",
	})
}

func (o *OrderHandler) DeleteOrderHandler(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Order ID harus numeric.",
		})
	}

	order, err := o.repo.GetOrderByID(orderId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Order tidak ditemukan.",
		})
	}

	err = o.repo.DeleteOrder(order)
	if err != nil {
		errMsg := err.Error()
		if err == sql.ErrConnDone {
			errMsg = "DB ERROR: " + errMsg
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errMsg,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order deleted.",
	})
}

func NewOrderHandler(repo repository.OrderRepository) *OrderHandler {
	return &OrderHandler{
		repo: repo,
	}
}
