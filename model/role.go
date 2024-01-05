package model

type Role struct {
	ID     int
	Code   string
	Name   string
	Enable int
}

func (Role) TableName() string {
	return "role"
}
