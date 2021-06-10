package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/akhilbidhuri/shopHere/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createToken() string {
	t, _ := time.Now().MarshalText()
	auxilaryUUID, _ := uuid.New().MarshalText()
	uuid, _ := uuid.FromBytes(append(t, auxilaryUUID...))
	return uuid.String()
}

func (a *App) addUser(c *gin.Context) {
	user := models.User{}
	if err := parseJsonFromReq(c.Request, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": INVALID_BODY,
		})
		return
	}
	err := user.Create(a.storage.DB)
	if err != nil {
		log.Println(" Create User Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": FAILED,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"user":    user.ToMap(),
	})
}

func (a *App) login(c *gin.Context) {
	user := models.User{}
	if err := parseJsonFromReq(c.Request, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": INVALID_BODY,
		})
		return
	}
	err := user.GetUserByUname(a.storage.DB, user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not present",
		})
		return
	}
	if err = models.VerifyPassword(user.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user details",
		})
		return
	}
	token := createToken()
	user.Token = token
	if err = user.Update(a.storage.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": FAILED,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"user":    user.ToMap(),
	})
}

func (a *App) listUser(c *gin.Context) {
	user := models.User{}
	users, err := user.GetAllUsers(a.storage.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get User list",
		})
		return
	}
	var usersList []map[string]interface{}
	for _, user := range *users {
		usersList = append(usersList, user.ToMap())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": SUCCESS,
		"users":   usersList,
	})
}
