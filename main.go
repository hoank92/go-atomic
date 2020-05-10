package main

import (
	"github.com/gin-gonic/gin"
	"go-atomic/main/helpers"
	"go-atomic/main/routers"
	"io/ioutil"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	db := helpers.Init()
	helpers.InitData()
	defer db.Close()
	r := routers.SetupRouter()
	// running
	_ = r.Run()

}
