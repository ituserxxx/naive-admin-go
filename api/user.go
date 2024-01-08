package api

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
	"time"
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
		PageData: make([]inout.UserListItem, 0),
	}
	var gender = c.DefaultQuery("gender", "")
	var enable = c.DefaultQuery("enable", "")
	var username = c.DefaultQuery("username", "")
	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)
	var profileList []model.Profile
	orm := db.Dao.Model(model.Profile{})
	if gender != "" {
		orm = orm.Where("gender=?", gender)
	}
	if enable != "" {
		orm = orm.Where("userId in(?)", db.Dao.Model(model.User{}).Where("enable=?", enable).Select("id"))
	}
	if username != "" {
		orm = orm.Where("nickName like ?", "%"+username+"%")
	}

	var total int64
	orm.Count(&total)
	orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&profileList)
	for _, datum := range profileList {
		var uinfo model.User
		db.Dao.Model(model.User{}).Where("id=?", datum.UserId).First(&uinfo)
		var rols []*model.Role
		db.Dao.Model(model.Role{}).Where("id IN (?)", db.Dao.Model(model.UserRolesRole{}).Where("userId=?", datum.UserId).Select("roleId")).Find(&rols)
		data.PageData = append(data.PageData, inout.UserListItem{
			ID:         uinfo.ID,
			Username:   uinfo.Username,
			Enable:     uinfo.Enable,
			CreateTime: uinfo.CreateTime,
			UpdateTime: uinfo.UpdateTime,
			Gender:     datum.Gender,
			Avatar:     datum.Avatar,
			Address:    datum.Address,
			Email:      datum.Email,
			Roles:      rols,
		})
	}
	Resp.Succ(c, data)
}

func (user) Profile(c *gin.Context) {
	var params inout.PatchProfileUserReq
	err := c.BindJSON(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	err = db.Dao.Model(model.Profile{}).Where("id=?", params.Id).Updates(model.Profile{
		Gender:   params.Gender,
		Address:  params.Address,
		Email:    params.Email,
		NickName: params.NickName,
	}).Error
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	Resp.Succ(c, err)
}
func (user) Update(c *gin.Context) {
	var params inout.PatchUserReq
	err := c.BindJSON(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	orm := db.Dao.Model(model.User{}).Where("id=?", params.Id)
	if params.Password != nil {
		orm.Update("password", fmt.Sprintf("%x", md5.Sum([]byte(*params.Password))))
	}
	if params.Enable != nil {
		orm.Update("enable", *params.Enable)
	}
	if params.Username != nil {
		orm.Update("username", *params.Username)
		db.Dao.Model(model.Profile{}).Where("userId=?", params.Id).Update("nickName", *params.Username)
	}
	if params.RoleIds != nil {
		db.Dao.Where("userId=?", params.Id).Delete(model.UserRolesRole{})
		if len(*params.RoleIds) > 0 {
			for _, i2 := range *params.RoleIds {
				db.Dao.Model(model.UserRolesRole{}).Create(&model.UserRolesRole{
					UserId: params.Id,
					RoleId: i2,
				})
			}
		}
	}

	Resp.Succ(c, err)
}

func (user) Add(c *gin.Context) {
	var params inout.AddUserReq
	err := c.Bind(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	err = db.Dao.Transaction(func(tx *gorm.DB) error {
		var newUser = model.User{
			Username:   params.Username,
			Password:   fmt.Sprintf("%x", md5.Sum([]byte(params.Password))),
			Enable:     params.Enable,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = tx.Create(&newUser).Error
		if err != nil {
			return err
		}
		tx.Create(&model.Profile{
			UserId:   newUser.ID,
			NickName: newUser.Username,
		})
		for _, id := range params.RoleIds {
			tx.Create(&model.UserRolesRole{
				UserId: newUser.ID,
				RoleId: id,
			})
		}
		return nil
	})
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	Resp.Succ(c, "")
}
func (user) Delete(c *gin.Context) {
	uid := c.Param("id")
	err := db.Dao.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", uid).Delete(&model.User{})
		tx.Where("userId =?", uid).Delete(&model.UserRolesRole{})
		tx.Where("userId =?", uid).Delete(&model.Profile{})
		return nil
	})
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	Resp.Succ(c, "")
}
