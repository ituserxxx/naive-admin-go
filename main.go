package main

import (
	"github.com/gin-gonic/gin"
	"naive-admin-go/config"
	"naive-admin-go/db"
	"naive-admin-go/router"
	"time"
)

func main() {
	var Loc, _ = time.LoadLocation("Asia/Shanghai")
	time.Local = Loc
	app := gin.Default()
	config.Init()
	db.Init()
	router.Init(app)
	app.Run(":8800")
}
