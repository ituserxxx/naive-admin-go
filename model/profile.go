package model

type Profile struct {
	ID       int    `json:"id"`
	Gender   int    `json:"gender"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	UserId   int    `gorm:"column:userId"`
	NickName string `gorm:"column:nickName"`
}

func (Profile) TableName() string {
	return "profile"
}
