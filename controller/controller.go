package controller

import (
	"log"
	"os"

	"github.com/akhilbidhuri/shopHere/models"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type App struct {
	router  *gin.Engine
	storage models.Storage
}

func (a *App) Intitialize() {
	a.router = gin.Default()
	if err := a.storage.Setup(); err != nil {
		log.Fatal("Failed to create db/storage, Error :", err)
	}
	a.router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
	}))
	a.initializeRoutes()
	log.Println("Starting Server on Port ", os.Getenv("PORT"), "...")
	a.router.Run(os.Getenv("PORT"))
}

func (a *App) initializeRoutes() {
	a.router.Use(static.Serve("/", static.LocalFile("./build", true)))
	a.router.POST("/user/create", a.addUser)
	a.router.POST("/user/login", a.login)
	a.router.GET("/user/list", a.listUser)
	a.router.POST("/item/create", a.addItem)
	a.router.GET("/item/list", a.listItems)
	a.router.POST("/cart/add", a.addItemToCart)
	a.router.GET("/cart/list", a.listCarts)
	a.router.GET("/cart/:cartID/list", a.listCartItems)
	a.router.GET("/cart/:cartID/complete", a.createOrder)
	a.router.GET("/order/list", a.listOrders)
	a.router.GET("/order/:userName/list", a.listOrdersForUser)
}
