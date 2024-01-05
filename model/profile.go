package model

type Profile struct {
	ID       int
	Gender   int
	Avatar   string
	Address  string
	Email    string
	UserId   int    `gorm:"column:userId"`
	NickName string `gorm:"column:nickName"`
}

func (Profile) TableName() string {
	return "profile"
}
