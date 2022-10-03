package router

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/api"
	"log"
)

func RouterInit() {
	var router = gin.Default()
	router.POST("/login", api.Login)
	router.POST("/todo", api.CreateTodo)
	router.POST("/logout", api.Logout)
	router.POST("/refresh", api.Refresh)
	log.Fatal(router.Run(":8080"))
}
