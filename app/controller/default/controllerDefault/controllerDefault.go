package controllerDefault

import (
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/okxBRC20/app/controller/base/controllerBase"
	"github.com/kanelinweihao/okxBRC20/app/utils/err"
	"github.com/kanelinweihao/okxBRC20/app/utils/pack"
	"io/fs"
	"net/http"
)

func Home(c *gin.Context) {
	data := "okxBRC20"
	controllerBase.Back(c, data)
	return
}

func Ping(c *gin.Context) {
	data := "Pong"
	controllerBase.Success(c, data)
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
