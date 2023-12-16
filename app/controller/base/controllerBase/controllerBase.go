package controllerBase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/okxBRC20/app/utils/result"
	"net/http"
)

func Back(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
	fmt.Println()
	return
}

func Success(c *gin.Context, data interface{}) {
	resultSuccess := result.GetResultSuccess(data)
	c.JSON(http.StatusOK, resultSuccess)
	fmt.Println()
	return
}

func Fail(c *gin.Context, err error) {
	resultFail := result.GetResultFail(err)
	c.JSON(http.StatusBadRequest, resultFail)
	fmt.Println()
	return
}
