package models

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:用户手机号"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	Nickname string `json:"nickname" gorm:"not null;default:'';comment:用户昵称"`
}

func (u User) TableName() string {
	return "sys_user"
}
