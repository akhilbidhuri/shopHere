package controller

import (
	"net/http"

	"github.com/akhilbidhuri/shopHere/models"
	"github.com/gin-gonic/gin"
)

func (a *App) addItem(c *gin.Context) {
	item := models.Item{}
	if err := parseJsonFromReq(c.Request, &item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": INVALID_BODY,
		})
		return
	}
	err := item.Create(a.storage.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": FAILED,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"item":    item.ToMap(),
	})
}

func (a *App) listItems(c *gin.Context) {
	item := models.Item{}
	items, err := item.GetAllItems(a.storage.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get Item list",
		})
		return
	}
	var itemsList []map[string]interface{}
	for _, item := range *items {
		itemsList = append(itemsList, item.ToMap())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"users":   itemsList,
	})
}
