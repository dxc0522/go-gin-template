package middleware

import (
	"github.com/go-gin-template/internal/gin-study/app/common/response"
	"github.com/go-gin-template/internal/gin-study/global"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	write := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Log.Filename,
		MaxSize:    global.App.Config.Log.MaxSize,
		MaxBackups: global.App.Config.Log.MaxBackups,
		MaxAge:     global.App.Config.Log.MaxAge,
		Compress:   global.App.Config.Log.Compress,
	})
	return gin.RecoveryWithWriter(
		write,
		response.ServerError)
}
