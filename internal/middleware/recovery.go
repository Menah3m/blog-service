package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"menah3m/blog-service/global"
	"menah3m/blog-service/pkg/app"
	"menah3m/blog-service/pkg/email"
	"menah3m/blog-service/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出, 发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
				)
				if err != nil {
					global.Logger.Panicf("mail.SendMail err:%v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
