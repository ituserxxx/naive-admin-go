package inout

type LoginReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}
type AuthPwReq struct {
	NewPassword string `form:"newPassword" binding:"required"`
	OldPassword string `form:"oldPassword" binding:"required"`
}
type PatchUserReq struct {
	Id       int     `json:"id"  binding:"required"`
	Enable   *bool   `json:"enable,omitempty"`
	RoleIds  *[]int  `json:"roleIds,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}
type PatchProfileUserReq struct {
	Id      int    `json:"id"  binding:"required"`
	Gender  int    `json:"gender"`
	NickName string `json:"nickName"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
type EnableRoleReq struct {
	Enable bool `json:"enable" binding:"required"`
	Id     int  `json:"id"`
}

type AddUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Enable   bool   `json:"enable" binding:"required"`
	RoleIds  []int  `json:"roleIds" binding:"required"`
}

type AddRoleReq struct {
	Code          string `json:"code" binding:"required"`
	Enable        bool   `json:"enable"`
	Name          string `json:"name" binding:"required"`
	PermissionIds []int  `json:"permissionIds"`
}
type PatchRoleReq struct {
	Id            int     `json:"id"  binding:"required"`
	Code          *string `json:"code,omitempty"`
	Enable        *bool   `json:"enable,omitempty"`
	Name          *string `json:"name,omitempty"`
	PermissionIds *[]int  `json:"permissionIds,omitempty"`
}

type PatchRoleOpeateUserReq struct {
	Id      int   `json:"id" `
	UserIds []int `json:"userIds"`
}

type AddPermissionReq struct {
	Type      string `json:"type" binding:"required"`
	ParentId  *int   `json:"parentId"`
	Name      string `json:"name" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Path      string `json:"path"`
	Icon      string `json:"icon"`
	Layout    string `json:"layout"`
	Component string `json:"component"`
	Show      bool   `json:"show"`
	Enable    bool   `json:"enable"`
	KeepAlive bool   `json:"keepAlive"`
	Order     int    `json:"order"`
}

type PatchPermissionReq struct {
	Id        int    `json:"id"  binding:"required"`
	Type      string `json:"type" binding:"required"`
	ParentId  *int   `json:"parentId"`
	Name      string `json:"name" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Path      string `json:"path"`
	Icon      string `json:"icon"`
	Layout    string `json:"layout"`
	Component string `json:"component"`
	Show      int   `json:"show"`
	Enable    int   `json:"enable"`
	KeepAlive int   `json:"keepAlive"`
	Order     int    `json:"order"`
}
