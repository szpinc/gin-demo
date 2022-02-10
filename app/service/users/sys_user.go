package users

import (
	"gin-demo/app/models"
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
