package api

import (
	"github.com/gin-gonic/gin"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
)

var User = &user{}

type user struct {
}

func (user) Detail(c *gin.Context) {
	var data = &inout.UserDetailRes{}
	var uid, _ = c.Get("uid")
	db.Dao.Model(model.User{}).Where("id=?", uid).Find(&data)
	db.Dao.Model(model.Profile{}).Where("userId=?", uid).Find(&data.Profile)
	urolIdList := db.Dao.Model(model.UserRolesRole{}).Where("userId=?", uid).Select("roleId")
	db.Dao.Model(model.Role{}).Where("id IN (?)", urolIdList).Find(&data.Roles)
	if len(data.Roles) > 0 {
		data.CurrentRole = data.Roles[0]
	}
	Resp.Succ(c, data)
}

func (user) List(c *gin.Context) {
	var data = inout.UserListRes{
		PageData: make([]inout.UserListItem,0),
	}
	var gender = c.DefaultQuery("gender","")
	var enable = c.DefaultQuery("enable","")
	var username = c.DefaultQuery("username","")
	var pageNoReq = c.DefaultQuery("pageNo","1")
	var pageSizeReq = c.DefaultQuery("pageSize","10")
	pageNo,_ := strconv.Atoi(pageNoReq)
	pageSize,_ := strconv.Atoi(pageSizeReq)
	var proLists []model.Profile
	orm := db.Dao.Model(model.Profile{})
	if gender != ""{
		orm = orm.Where("gender=?",gender)
	}
	if enable != ""{
		orm = orm.Where("userId in(?)",db.Dao.Model(model.User{}).Where("enable=?",enable).Select("id"))
	}
	if username != ""{
		orm = orm.Where("nickName like ?","%"+username+"%")
	}

	var total int64
	orm.Count(&total)
	orm.Offset((pageNo-1)*pageSize).Limit(pageSize).Select("*").Find(&proLists)
	for _, datum := range proLists{
		 var u model.User
		 db.Dao.Model(model.User{}).Where("id=?",datum.UserId).First(&u)
		var rols []*model.Role
		db.Dao.Model(model.Role{}).Where("id IN (?)", db.Dao.Model(model.UserRolesRole{}).Where("userId=?",datum.UserId).Select("roleId")).Find(&rols)
		data.PageData = append(data.PageData,inout.UserListItem{
			User:    u,
			Profile: datum,
			Roles:   rols,
		})
	}
	Resp.Succ(c,data)
}