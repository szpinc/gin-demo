package users

import (
	"crypto/sha256"
	"fmt"
	"gin-demo/app/models"
	"gin-demo/app/models/request"
	"gin-demo/global"

	"gorm.io/gorm"
)

func GetUserInfo(userId uint) *models.User {
	result := models.User{}
	if tx := global.App.DB.Find(&result, models.User{
		BaseModel: models.BaseModel{
			ID: userId,
		},
	}); tx.Error != nil {
		switch tx.Error {
		case gorm.ErrRecordNotFound: // 记录不存在，返回nil
			return nil
		default: // 其他错误,则panic
			panic(tx.Error)
		}
	}
	return &result
}

func Register(userRegister request.UserRegister) *models.User {

	// 密码md5加密
	passwordEncoded := fmt.Sprintf("%v", sha256.Sum256([]byte(userRegister.Password)))

	u := models.User{
		Name:     userRegister.UserName,
		Password: passwordEncoded,
		Nickname: userRegister.Nickname,
	}
}
