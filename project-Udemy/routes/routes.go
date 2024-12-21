package routes

import (
	"github.com/Wong-bui/Udemy-project/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvents)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
