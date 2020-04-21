package users

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/securityin/auth/models"
	"github.com/securityin/auth/pkg/e"
	"github.com/securityin/auth/pkg/logging"
	"github.com/securityin/auth/pkg/util"
)

// Oauth 第三方登录认证
type Oauth struct {
	ID               uint `gorm:"primary_key"`
	UserID           uint
	OauthType        uint   `json:"oauth_type"`
	OauthID          string `json:"oauth_id" gorm:"unique"`
	OauthAccessToken string `json:"access_token"`
	OauthExpires     string `json:"expires_at"`
	Status           uint   `json:"status"`
}

// AddUserOauth 添加用户账号 与 初始化个人信息
func AddUserOauth(userOatuh map[string]interface{}) util.Error {

	tx := models.DB.Begin()

	// 首先创建 user
	userID, err := addUser(tx)
	if err != nil {
		tx.Rollback()
		return util.ErrNewCode(e.ErrorUserInfoCreate)
	}
	fmt.Println(userOatuh)

	oauthInfo := Oauth{
		UserID:           userID,
		OauthID:          userOatuh["oauth_id"].(string),
		OauthType:        userOatuh["oauth_type"].(uint),
		OauthAccessToken: userOatuh["access_token"].(string),
		OauthExpires:     userOatuh["expires"].(string),
	}
	if err := tx.Create(&oauthInfo).Error; err != nil {
		tx.Rollback()
		logging.GetLogger().Error(err)
		return util.ErrNewCode(e.ErrorUserLoginCreate)
	}
	tx.Commit()
	return nil
}

// ExistUserOauth 判断用户账号是否存在
func ExistUserOauth(maps map[string]interface{}) (bool, util.Error) {
	var user Oauth
	err := models.DB.Select("id").Where(maps).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.GetLogger().Error(err)
		return false, util.ErrNewCode(e.ErrorUserLoginEmpty)
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
