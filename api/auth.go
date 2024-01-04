package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"naive-admin-go/inout"
	"naive-admin-go/utils"
	"naive-admin-go/utils/response"
	"net/http"
)

var Auth = &auth{}
type auth struct {

}


func (auth)Captcha(c *gin.Context)  {
	svg,code := utils.GenerateSVG(80,40)
	cookieStr := c.Request.Header.Get("cookie")
	session.Set(cookieStr, code)
	session.Save()
	// 设置 Content-Type 为 "image/svg+xml"
	c.Header("Content-Type", "image/svg+xml; charset=utf-8")
	// 返回验证码
	c.Data(http.StatusOK, "image/svg+xml",svg)
}

func (auth)Login(c *gin.Context)  {
	var params inout.Login
	err := c.Bind(&params)
	if err != nil {
		response.Err(c,20001,err.Error())
	}
	cookieStr := c.Request.Header.Get("cookie")
	session := sessions.Default(c)

	fmt.Printf("1111111%#v code= %#v", params,session.Get(cookieStr))
}