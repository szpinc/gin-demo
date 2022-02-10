package request

import "gin-demo/app/common/validate"

type UserRegister struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
}

func (u UserRegister) GetErrorMsg() validate.ValidatorMessages {
	return validate.ValidatorMessages{
		"UserName.required": "用户不能为空",
		"Password.required": "用户密码不能为空",
		"Nickname.required": "用户昵称不能为空",
	}
}
