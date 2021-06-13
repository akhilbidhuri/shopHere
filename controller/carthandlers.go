package controller

import (
	"net/http"
	"strconv"

	"github.com/akhilbidhuri/shopHere/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (a *App) addItemToCart(c *gin.Context) {
	if err := validateToken(c.Request, a.storage.DB); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": INVALID_TOKEN,
		})
		return
	}
	cartItem := models.CartItem{}
	if err := parseJsonFromReq(c.Request, &cartItem); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": INVALID_BODY,
		})
		return
	}
	err := cartItem.Create(a.storage.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": FAILED,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   SUCCESS,
		"cart_item": cartItem.ToMap(),
	})
}

func (a *App) listCarts(c *gin.Context) {
	cart := models.Cart{}
	carts, err := cart.GetAllCarts(a.storage.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get Cart list",
		})
		return
	}
	var cartsList []map[string]interface{}
	for _, cart := range *carts {
		cartsList = append(cartsList, cart.ToMap())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"carts":   cartsList,
	})
}

func (a *App) listCartItems(c *gin.Context) {
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
	cartItem := models.CartItem{Cart_id: cartID}
	cartItems, err := cartItem.GetItemsForCart(a.storage.DB)
	if err != nil {
		var msg string
		if err == gorm.ErrRecordNotFound {
			msg = "No Items Found in Cart"
		} else {
			msg = FAILED
		}
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
		return
	}
	var cartItemsList []map[string]interface{}
	for _, ci := range *cartItems {
		cartItemsList = append(cartItemsList, ci.ToMap())
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   SUCCESS,
		"cartItems": cartItemsList,
	})
}
