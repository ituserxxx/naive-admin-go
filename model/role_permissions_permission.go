package model

type RolePermissionsPermission struct {
	RoleId       int `gorm:"column:roleId"`
	PermissionId int `gorm:"column:permissionId"`
}

func (RolePermissionsPermission) TableName() string {
	return "role_permissions_permission"
}
