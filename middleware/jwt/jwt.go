package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/securityin/auth/pkg/app"
	"github.com/securityin/auth/pkg/e"
	"github.com/securityin/auth/pkg/util"
)

// JWT middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.GetGin(c)
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Request.Header.Get("jwtToken")
		if token == "" {
			code = e.ErrorInvalidParamsWithoutToken
		} else {
			claims, err := util.ParseToken2(token)
			if err != nil {
				code = e.ErrorAuthParseTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
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
