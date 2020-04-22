package captcha

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/securityin/auth/pkg/app"
	"github.com/securityin/auth/pkg/e"
)

// Captcha 验证码中间件
func Captcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.GetGin(c)
		var code int
		var data interface{}
		code = e.SUCCESS

		if id, r := c.GetQuery("id"); !r {
			code = e.ErrorUserCaptchaMissing
		} else if digits, r := c.GetQuery("digits"); !r {
			code = e.ErrorUserCaptchaMissing
		} else {
			if !captcha.VerifyString(id, digits) {
				code = e.ErrorUserCaptcha
			}
		}

		if code != e.SUCCESS {
			appG.Response(http.StatusUnauthorized, code, data)
			// 拦截
			c.Abort()
			return
		}
		c.Next()
	}

}
