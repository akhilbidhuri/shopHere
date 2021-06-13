package controller

import (
	"net/http"
	"strconv"

	"github.com/akhilbidhuri/shopHere/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func isCartEmpty(cartId int, db *gorm.DB) bool {
	cartItem := models.CartItem{Cart_id: cartId}
	cartItems, _ := cartItem.GetItemsForCart(db)
	return len(*cartItems) == 0
}

func (a *App) createOrder(c *gin.Context) {
	if err := validateToken(c.Request, a.storage.DB); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": INVALID_TOKEN,
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
	if isCartEmpty(cartID, a.storage.DB) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cart is Empty",
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
	updatedUser, err := cart.GetNewCart(a.storage.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": FAILED,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"order":   order.ToMap(),
		"user":    updatedUser.ToMap(),
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
		"orders":  ordersList,
	})
}

func (a *App) listOrdersForUser(c *gin.Context) {
	if err := validateToken(c.Request, a.storage.DB); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": INVALID_TOKEN,
		})
		return
	}
	userName := c.Param("userName")
	user := models.User{}
	if err := user.GetUserByUname(a.storage.DB, userName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": INVALID_REQUEST,
		})
		return
	}
	order := models.Order{User_id: user.Id}
	orders, err := order.GetOrderByUID(a.storage.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": INVALID_REQUEST,
		})
		return
	}
	var ordersList []map[string]interface{}
	for _, order = range *orders {
		ordersList = append(ordersList, order.ToMap())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"orders":  ordersList,
	})
}
