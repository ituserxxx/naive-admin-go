package model

type UserRolesRole struct {
	UserId int `gorm:"column:userId"`
	RoleId int `gorm:"column:roleId"`
}

func (UserRolesRole) TableName() string {
	return "user_roles_role"
}
