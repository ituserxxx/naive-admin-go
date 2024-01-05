package model

import (
	"time"
)

type User struct {
	ID          int
	Username    string
	Password    string
	Enable      int
	CreateTime  time.Time `gorm:"column:createTime"`
	UpdateTime  time.Time `gorm:"column:updateTime"`
	Profile     Profile   `json:"profile"`
	Roles       []Role    `json:"roles"`
	CurrentRole Role      `json:"currentRole"`
}

func (User) TableName() string {
	return "user"
}
