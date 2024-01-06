package api

import (
	"github.com/gin-gonic/gin"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
)

var Role = &role{}

type role struct {
}

func (role) PermissionsTree(c *gin.Context) {
	var data = make([]inout.PermissionsTreeItem, 0)
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

	for _, perm := range onePermissList {
		var twoPerissList []model.Permission
		db.Dao.Model(model.Permission{}).Where("parentId = ?", perm.ID).Order("`order` Asc").Find(&twoPerissList)
		data = append(data, inout.PermissionsTreeItem{Permission: perm, Children: twoPerissList})
	}
	Resp.Succ(c, data)
}

func (role) List(c *gin.Context) {
	var data = &inout.RoleListRes{}

	db.Dao.Model(model.Role{}).Find(&data)

	Resp.Succ(c, data)
}
func (role) ListPage(c *gin.Context) {
	var data = &inout.RoleListPageRes{}
	var name = c.DefaultQuery("name","")
	var pageNoReq = c.DefaultQuery("pageNo","1")
	var pageSizeReq = c.DefaultQuery("pageSize","10")
	pageNo,_ := strconv.Atoi(pageNoReq)
	pageSize,_ := strconv.Atoi(pageSizeReq)
	orm := db.Dao.Model(model.Role{})
	if name != ""{
		orm = orm.Where("name like ?","%"+name+"%")
	}
	var total int64
	orm.Count(&total)

	orm.Offset((pageNo-1)*pageSize).Limit(pageSize).Find(&data.PageData)
	for i, datum := range data.PageData {
		var perIdList []int64
		db.Dao.Model(model.RolePermissionsPermission{}).Where("roleId=?",datum.ID).Select("permissionId").Find(&perIdList)
		data.PageData[i].PermissionIds = perIdList
	}
	Resp.Succ(c, data)
}