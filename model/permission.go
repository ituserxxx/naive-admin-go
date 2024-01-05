package model

type Permission struct {
	ID          int
	Name        string
	Code        string
	Type        string
	ParentId    int `gorm:"column:parentId"`
	Path        string
	Redirect    string
	Icon        string
	Component   string
	Layout      string
	KeepAlive   int `gorm:"column:keepAlive"`
	Method      string
	Description string
	Show        int
	Enable      int
	Order       int
}

func (Permission) TableName() string {
	return "permission"
}
