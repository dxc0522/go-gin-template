package common

import (
	"github.com/go-gin-template/internal/gin-study/app/common/request"
	"github.com/go-gin-template/internal/gin-study/app/common/response"
	"github.com/go-gin-template/internal/gin-study/app/services"

	"github.com/gin-gonic/gin"
)

func ImageUpload(c *gin.Context) {
	var form request.FileUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveFile(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}
