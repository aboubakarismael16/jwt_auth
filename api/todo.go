package api

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/utils"
	"net/http"
)

type Todo struct {
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
}

func CreateTodo(c *gin.Context) {
	var td Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	//Extract the access token metadata
	metadata, err := utils.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	userid, err := utils.FetchAuth(metadata)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	td.UserID = userid
	//you can proceed to save the Todo to a database
	//but we will just return it to the caller:

	c.JSON(http.StatusCreated, td)
}
