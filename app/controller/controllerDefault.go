package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/okxBRC20/app/utils/err"
	"github.com/kanelinweihao/okxBRC20/app/utils/pack"
	"github.com/kanelinweihao/okxBRC20/app/utils/result"
	"io/fs"
	"net/http"
)

func Home(c *gin.Context) {
	projectName := "okxBRC20"
	c.String(http.StatusOK, projectName)
	return
}

func Ping(c *gin.Context) {
	data := "Pong"
	resultSuccess := result.GetResultSuccess(data)
	c.JSON(http.StatusOK, resultSuccess)
	return
}

func Favicon(c *gin.Context) {
	fsEmbedFavicon := pack.GetFSEmbedFavicon()
	fsFavicon, errFS := fs.Sub(fsEmbedFavicon, "resource/ico")
	err.ErrCheck(errFS)
	handlerFavicon := http.FileServer(http.FS(fsFavicon))
	handlerFavicon.ServeHTTP(c.Writer, c.Request)
	return
}
