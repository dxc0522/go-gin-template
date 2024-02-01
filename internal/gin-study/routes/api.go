package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gin-template/internal/gin-study/app/common/request"
	"github.com/go-gin-template/internal/gin-study/app/common/response"
	"github.com/go-gin-template/internal/gin-study/app/controllers/app"
	"github.com/go-gin-template/internal/gin-study/app/controllers/common"
	"github.com/go-gin-template/internal/gin-study/app/middleware"
	"github.com/go-gin-template/internal/gin-study/app/services"
	"github.com/go-gin-template/internal/gin-study/global"
	"github.com/go-gin-template/internal/gin-study/model"
	"log"
	"net/http"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		global.App.Log.Info("ping pong")
		log.Println("ping pong")
		c.String(http.StatusOK, "pong")
	})
	router.POST("/test", func(c *gin.Context) {
		var form request.ToDo
		if err := c.ShouldBindJSON(&form); err != nil {
			response.ValidateFail(c, request.GetErrorMsg(form, err))
			return
		}
		record := model.ToDo{
			Label: form.Label,
			Value: form.Value,
		}
		err := global.App.DB.Create(&record).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, record)
		}
	})
	router.POST("/user/register", func(c *gin.Context) {
		var form request.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/upload", common.ImageUpload)
	}
}
