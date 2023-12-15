package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanelinweihao/OKXBRC20/app/middleware"
	"github.com/kanelinweihao/OKXBRC20/app/router"
	"github.com/kanelinweihao/OKXBRC20/app/utils/err"
	"github.com/kanelinweihao/OKXBRC20/app/utils/times"
	"io"
	"os"
	"strings"
)

func Boot() {
	ready()
	exec()
	return
}

func ready() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	// gin.ForceConsoleColor()
	f := getFileOfGinLog()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	return
}

func getFileOfGinLog() (f io.Writer) {
	projectName := "OKXBRC20"
	projectNameLower := strings.ToLower(projectName)
	pathLog := fmt.Sprintf(
		"./tmp/logs/log_gin_%s.log",
		projectNameLower)
	f, errLogCreate := os.Create(pathLog)
	err.ErrCheck(errLogCreate)
	return f
}

func exec() {
	r := gin.Default()
	r.Use(middleware.FuncMiddleWareBase())
	// setStaticFile(r)
	setRouter(r)
	runRouter(r)
	return
}

func setStaticFile(r *gin.Engine) {
	r.StaticFile("/favicon.ico", "./resource/ico/favicon.ico")
	return
}

func setRouter(r *gin.Engine) {
	router.SetRouter(r)
	times.Sleep(100, "ms")
	times.ShowTimeAndMsg("Ready!")
	return
}

func runRouter(r *gin.Engine) {
	addr := ":8080"
	errRun := r.Run(addr)
	err.ErrCheck(errRun)
	return
}
