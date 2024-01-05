package inout

import (
	"naive-admin-go/model"
)

type LoginRes struct {
	AccessToken string `json:"accessToken"`
}

type UserDetailRes struct {
	model.User
}
