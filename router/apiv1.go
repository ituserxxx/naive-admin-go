package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"naive-admin-go/api"
	"naive-admin-go/middleware"
)

func Init(r *gin.Engine) {
	// 使用 cookie 存储会话数据
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("captch"))))
	r.Use(middleware.Cors())

	auth := r.Group("/auth")
	auth.POST("/login", api.Auth.Login)
	auth.GET("/captcha", api.Auth.Captcha)
	r.Use(middleware.Jwt())
	user := r.Group("/user")
	user.GET("/detail", api.User.Detail)
}
