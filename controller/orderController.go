package controller

import (
	"net/http"
	"strconv"

	"github.com/akhilbidhuri/shopHere/models"
	"github.com/gin-gonic/gin"
)

func (a *App) createOrder(c *gin.Context) {
	if err := validateToken(c.Request, a.storage.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": INVALID_REQUEST,
		})
		return
	}
	cartID, err := strconv.Atoi(c.Param("cartID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": INVALID_BODY,
		})
		return
	}
	cart := models.Cart{}
	if err = cart.GetCartByID(a.storage.DB, cartID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": INVALID_BODY,
		})
		return
	}
	order := models.Order{}
	order.Cart_id = cart.Id
	order.User_id = cart.User_id
	if err = order.Create(a.storage.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": FAILED,
		})
		return
	}
	cart.Is_purchased = true
	if err = cart.Update(a.storage.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": FAILED,
		})
		return
	}
	//Assign new Empty cart to user once order is created
	cart.GetNewCart(a.storage.DB)
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"order":   order.ToMap(),
	})
}

func (a *App) listOrders(c *gin.Context) {
	order := models.Order{}
	orders, err := order.GetAllOrders(a.storage.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get Order list",
		})
		return
	}
	var ordersList []map[string]interface{}
	for _, order := range *orders {
		ordersList = append(ordersList, order.ToMap())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"users":   ordersList,
	})
}
