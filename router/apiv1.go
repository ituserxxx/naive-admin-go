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

	r.POST("/auth/login", api.Auth.Login)
	r.GET("/auth/captcha", api.Auth.Captcha)
	r.POST("/auth/logout", api.Auth.Logout)

	r.Use(middleware.Jwt())

	r.GET("/user", api.User.List)
	r.GET("/user/detail", api.User.Detail)

	r.GET("/role", api.Role.List)
	r.GET("/role/page", api.Role.ListPage)
	r.GET("/role/permissions/tree", api.Role.PermissionsTree)

	r.GET("/permission/tree", api.Permissions.List)
	r.GET("/permission/menu/tree", api.Permissions.List)
}
