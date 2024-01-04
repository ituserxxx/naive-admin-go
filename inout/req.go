package inout

type Login struct {
	Username     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha string `form:"captcha" binding:"required"`
}
