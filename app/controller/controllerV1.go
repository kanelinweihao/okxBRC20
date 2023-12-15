package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/OKXBRC20/app/utils/result"
	"net/http"
)

func success(c *gin.Context, data interface{}) {
	resultSuccess := result.GetResultSuccess(data)
	c.JSON(http.StatusOK, resultSuccess)
	return
}

func Inscribe(c *gin.Context) {
	name := c.DefaultQuery("name", "Mr.Jin")
	data := fmt.Sprintf("hello %s", name)
	success(c, data)
	return
}

func Transaction(c *gin.Context) {
	name := c.Query("name")
	data := fmt.Sprintf("hello %s\n", name)
	success(c, data)
	return
}
