package types

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResult(data interface{}) R {
	return R{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}
