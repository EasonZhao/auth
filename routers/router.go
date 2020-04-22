package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/securityin/auth/middleware/captcha"
	"github.com/securityin/auth/middleware/jwt"
	"github.com/securityin/auth/pkg/setting"
	"github.com/securityin/auth/routers/api"
	v1 "github.com/securityin/auth/routers/api/v1"
)

// InitRouter 返回 框架的实例 包含中间件 配置
func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	// 启动 gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 根目录
	root := r.Group("")
	{
		root.POST("/captcha", api.GetCaptcha)
		root.GET("/captcha/:captchaId", api.GetCaptchaImage)
		// 用户账号
		auth := root.Group("/auth")

		auth.Use(captcha.Captcha())
		{
			auth.POST("/register", api.Register)
			auth.POST("/login", api.Login)
			auth.POST("/code", api.SendCode)
			auth.POST("/phonelogin", api.PhoneLogin)
		}

		// 用户账号
		oauth := root.Group("/oauth")
		{
			oauth.GET("/github", api.LoginGithub)
			oauth.GET("/github/callback", api.CallBackGithub)
		}

		// apiv1
		apiv1 := root.Group("/api/v1")

		apiv1.Use(jwt.JWT())
		{
			apiv1.GET("/test", v1.TestAuth)

			// 用户
			user := apiv1.Group("/user")
			{
				user.POST("/getUserInfo", api.GetUserInfo)
			}
		}

	}

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
