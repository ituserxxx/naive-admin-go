package api

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"naive-admin-go/utils"
	"net/http"
)

var Auth = &auth{}

type auth struct {
}

func (auth) Captcha(c *gin.Context) {
	svg, code := utils.GenerateSVG(80, 40)
	session := sessions.Default(c)
	session.Set("captch", code)
	session.Save()
	// 设置 Content-Type 为 "image/svg+xml"
	c.Header("Content-Type", "image/svg+xml; charset=utf-8")
	// 返回验证码
	c.Data(http.StatusOK, "image/svg+xml", svg)
}

func (auth) Login(c *gin.Context) {
	var params inout.Login
	err := c.Bind(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	session := sessions.Default(c)
	if params.Captcha != session.Get("captch") {
		Resp.Err(c, 20001, "验证码不正确")
		return
	}
	var info *model.User
	db.Dao.Model(model.User{}).
		Where("username =? ", params.Username).
		Where("password=?", fmt.Sprintf("%x", md5.Sum([]byte(params.Password)))).
		Find(&info)
	if info.ID == 0 {
		Resp.Err(c, 20001, "账号或密码不正确")
		return

	}
	Resp.Succ(c, inout.LoginRes{
		AccessToken: utils.GenerateToken(info.ID),
	})

}


func (auth) Logout(c *gin.Context) {
	Resp.Succ(c, true)
}