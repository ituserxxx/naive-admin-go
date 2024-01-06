package model

type Role struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Enable int    `json:"enable"`
}

func (Role) TableName() string {
	return "role"
}
