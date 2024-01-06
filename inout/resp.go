package inout

import (
	"naive-admin-go/model"
)

type LoginRes struct {
	AccessToken string `json:"accessToken"`
}

type UserDetailRes struct {
	model.User
	Profile     *model.Profile `json:"profile"`
	Roles       []*model.Role  `json:"roles" `
	CurrentRole *model.Role    `json:"currentRole"`
}
type PermissionsTreeItem struct {
	model.Permission
	Children []model.Permission `json:"children"`
}
type RoleListRes []*model.Role
type PermissionsTreeRes []PermissionsTreeItem

type UserListItem struct {
	model.User
	model.Profile
	Roles []*model.Role `json:"roles"`
}
type UserListRes struct {
	PageData []UserListItem `json:"pageData"`
	Total    int            `json:"total"`
}
type RoleListPageItem struct {
	model.Role
	PermissionIds []int64 `json:"permissionIds" gorm:"-"`
}
type RoleListPageRes struct {
	PageData []RoleListPageItem `json:"pageData"`
	Total    int            `json:"total"`
}