package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status  Status      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Status struct {
	Code      int  `json:"code"`
	IsSuccess bool `json:"is_success"`
}

func OnSuccess(c *gin.Context, code int, message string, data any) {
	responseStr := Response{
		Status: Status{
			Code:      code,
			IsSuccess: true,
		},
		Message: message,
		Data:    data,
	}

	c.JSON(code, responseStr)
}

func OnFailed(c *gin.Context, code int, message string, err error) {
	responseStr := Response{
		Status: Status{
			Code:      code,
			IsSuccess: false,
		},
		Message: message,
		Data:    err.Error(),
	}

	c.JSON(code, responseStr)
}
