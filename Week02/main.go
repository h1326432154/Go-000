package main

import (
	"Week02/api"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	user := r.Group("user")
	{
		user.GET("/:id", api.GetUser)
	}
	r.Run(":" + "8000")
}
