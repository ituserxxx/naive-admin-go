package api

import (
	"github.com/gin-gonic/gin"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
)

var Permissions = &permissions{}

type permissions struct {
}

func (permissions) List(c *gin.Context) {
	var onePermissList = make( []model.Permission,0)
	db.Dao.Model(model.Permission{}).Where("parentId is NULL").Order("`order` Asc").Find(&onePermissList)
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

func (permissions) ListPage(c *gin.Context) {
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