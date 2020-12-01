package api

import (
	"Week02/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser c
func GetUser(c *gin.Context) {

	id := c.Param("id")
	var u dao.User
	u.ID, _ = strconv.Atoi(id)
	err := u.GetUserByID()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": u})
}
