package main

import (
	"fmt"
	"net/http"

	"github.com/securityin/auth/models"
	"github.com/securityin/auth/pkg/gredis"
	"github.com/securityin/auth/pkg/logging"
	"github.com/securityin/auth/pkg/oauth"
	"github.com/securityin/auth/pkg/setting"
	"github.com/securityin/auth/routers"
)

func main() {
	setting.Setup()
	logging.Setup()
	models.Setup()
	err := gredis.Setup()
	if err != nil {
		logging.GetLogger().Fatalln(err)
	}
	oauth.Setup()
	fmt.Println(setting.ServerSetting)
	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.Port),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if err = server.ListenAndServe(); err != nil {
		logging.GetLogger().Fatalln(err)
	}
}
