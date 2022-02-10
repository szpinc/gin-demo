package validate

import "github.com/go-playground/validator/v10"

type ValidatorMessages map[string]string

type Validator interface {
	GetMessage() ValidatorMessages
}

// GetErrorMsg 获取错误信息
func GetErrorMsg(request interface{}, err error) string {
	if validErrors, isValidError := err.(validator.ValidationErrors); isValidError {
		_, isValidator := request.(Validator)

		for _, e := range validErrors {
			if !isValidator {
				return e.Error()
			}
			// 若request结构体实现了Validator接口即可实现自定义错误信息
			if message, ok := request.(Validator).GetMessage()[e.Field()+"."+e.Tag()]; ok {
				return message
			}
		}
	}
	return "Parameter error"
}
