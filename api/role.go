package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
)

var Role = &role{}

type role struct {
}

func (role) PermissionsTree(c *gin.Context) {
	var uid, _ = c.Get("uid")

	var adminRole int64
	db.Dao.Model(model.UserRolesRole{}).Where("userId=? and roleId=1", uid).Count(&adminRole)
	orm := db.Dao.Model(model.Permission{}).Where("parentId is NULL").Order("`order` Asc")

	if adminRole == 0 {
		uroleIdList := db.Dao.Model(model.UserRolesRole{}).Where("userId=?", uid).Select("roleId")
		rpermisId := db.Dao.Model(model.RolePermissionsPermission{}).Where("roleId in(?)", uroleIdList).Select("permissionId")
		orm = orm.Where("id in(?)", rpermisId)
	}

	var onePermissList []model.Permission
	orm.Find(&onePermissList)

	for i, perm := range onePermissList {
		var twoPerissList []model.Permission
		db.Dao.Model(model.Permission{}).Where("parentId = ?", perm.ID).Order("`order` Asc").Find(&twoPerissList)
		for i2, perm2 := range twoPerissList {
			var twoPerissList2 []model.Permission
			db.Dao.Model(model.Permission{}).Where("parentId = ?", perm2.ID).Order("`order` Asc").Find(&twoPerissList2)
			twoPerissList[i2].Children = twoPerissList2
		}
		onePermissList[i].Children = twoPerissList
	}
	Resp.Succ(c, onePermissList)
}

func (role) List(c *gin.Context) {
	var data = &inout.RoleListRes{}

	db.Dao.Model(model.Role{}).Find(&data)

	Resp.Succ(c, data)
}
func (role) ListPage(c *gin.Context) {
	var data = &inout.RoleListPageRes{}
	var name = c.DefaultQuery("name", "")
	var enable = c.DefaultQuery("enable", "")
	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)
	orm := db.Dao.Model(model.Role{})
	if name != "" {
		orm = orm.Where("name like ?", "%"+name+"%")
	}
	if enable != "" {
		ena := false
		if enable == "1" {
			ena = true
		}
		orm = orm.Where("enable = ?", ena)
	}
	var total int64
	orm.Count(&total)

	orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&data.PageData)
	for i, datum := range data.PageData {
		var perIdList []int64
		db.Dao.Model(model.RolePermissionsPermission{}).Where("roleId=?", datum.ID).Select("permissionId").Find(&perIdList)
		data.PageData[i].PermissionIds = perIdList
	}
	Resp.Succ(c, data)
}
func (role) Update(c *gin.Context) {
	var params inout.PatchRoleReq
	err := c.BindJSON(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	orm := db.Dao.Model(model.Role{}).Where("id=?", params.Id)
	if params.Name != nil {
		orm.Update("name", *params.Name)
	}
	if params.Enable != nil {
		orm.Update("enable", *params.Enable)
	}
	if params.Code != nil {
		orm.Update("code", *params.Code)
	}
	if params.PermissionIds != nil {
		db.Dao.Where("roleId=?", params.Id).Delete(model.RolePermissionsPermission{})
		if len(*params.PermissionIds) > 0 {
			for _, i2 := range *params.PermissionIds {
				db.Dao.Model(model.RolePermissionsPermission{}).Create(&model.RolePermissionsPermission{
					PermissionId: i2,
					RoleId:       params.Id,
				})
			}
		}
	}
	Resp.Succ(c, err)
}

func (role) Add(c *gin.Context) {
	var params inout.AddRoleReq
	err := c.Bind(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	err = db.Dao.Transaction(func(tx *gorm.DB) error {
		var record = model.Role{
			Code:   params.Code,
			Name:   params.Name,
			Enable: params.Enable,
		}
		err = tx.Create(&record).Error
		if err != nil {
			return err
		}

		for _, id := range params.PermissionIds {
			tx.Create(&model.RolePermissionsPermission{
				RoleId:       record.ID,
				PermissionId: id,
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

func (role) Delete(c *gin.Context) {
	uid := c.Param("id")
	err := db.Dao.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", uid).Delete(&model.Role{})
		tx.Where("roleId =?", uid).Delete(&model.UserRolesRole{})
		tx.Where("roleId =?", uid).Delete(&model.RolePermissionsPermission{})
		return nil
	})
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	Resp.Succ(c, "")
}
func (role) AddUser(c *gin.Context) {
	var params inout.PatchRoleOpeateUserReq
	err := c.BindJSON(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	uid, _ := strconv.Atoi(c.Param("id"))
	params.Id = uid
	db.Dao.Where("userId in (?) and roleId = ?", params.UserIds, params.Id).Delete(model.UserRolesRole{})
	for _, id := range params.UserIds {
		db.Dao.Model(model.UserRolesRole{}).Create(model.UserRolesRole{
			UserId: id,
			RoleId: params.Id,
		})
	}
	Resp.Succ(c, "")
}
func (role) RemoveUser(c *gin.Context) {
	var params inout.PatchRoleOpeateUserReq
	err := c.BindJSON(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	uid, _ := strconv.Atoi(c.Param("id"))
	params.Id = uid
	db.Dao.Where("userId in (?) and roleId = ?", params.UserIds, params.Id).Delete(model.UserRolesRole{})
	Resp.Succ(c, "")
}
