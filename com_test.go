package main

import (
	"crypto/md5"
	"fmt"
	"naive-admin-go/config"
	"naive-admin-go/db"
	"naive-admin-go/model"
	"testing"
)

func TestMd(t *testing.T) {
	a := fmt.Sprintf("%s", md5.Sum([]byte("123456")))
	println(11111111)
	b := fmt.Sprintf("%s", md5.Sum([]byte("123456")))
	fmt.Printf("%v", a == b)
}
func TestName(t *testing.T) {
	config.Init()
	db.Init()
	var u *model.User
	db.Dao.Model(model.User{}).First(&u)
	u.Password = fmt.Sprintf("%x", md5.Sum([]byte("123456")))
	db.Dao.Save(&u)
	var u2 *model.User
	mm := fmt.Sprintf("%x", md5.Sum([]byte("123456")))
	db.Dao.Model(model.User{}).Where("password=?", mm).First(&u2)
	//fmt.Printf("\nmm  ->%x",mm)
	fmt.Printf("\n111 chis ->%#v", u2.ID)
	//fmt.Printf("111 chis ->%v",fmt.Sprintf("%s", md5.Sum([]byte("123456")))==u2.Password)
}
