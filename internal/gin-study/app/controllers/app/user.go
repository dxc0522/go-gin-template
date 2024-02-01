package app

import (
	"github.com/go-gin-template/internal/gin-study/app/common/request"
	"github.com/go-gin-template/internal/gin-study/app/common/response"
	"github.com/go-gin-template/internal/gin-study/app/services"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
