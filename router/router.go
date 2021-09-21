package router

import (
	"github.com/gin-gonic/gin"
	"github.com/glavanan/mailgun/handler"
)

//GetRouter return the router for the application
func GetRouter() *gin.Engine {
	router := gin.Default()
	router.PUT("events/:domain/delivered", handler.Delivered)
	router.PUT("events/:domain/bounced", handler.Bounced)
	router.GET("domains/:domain", handler.GetCatchAll)
	return router
}
