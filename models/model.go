package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/securityin/auth/pkg/logging"
	"github.com/securityin/auth/pkg/setting"
)

var (
	// DB 数据库
	DB *gorm.DB
)

// Model 基类
type Model struct {
	ID       uint      `gorm:"primary_key" json:"id"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"update_at"`
	// * 代表 NULL
	DeleteAt *time.Time `json:"delete_at"`
}

// Setup 初始化
func Setup() {
	var err error
	DB, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		logging.GetLogger().Fatalln(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.LogMode(setting.DatabaseSetting.LogMode)
	DB.SetLogger(logging.GetLogger())
}

// CloseDB 关闭
func CloseDB() {
	defer DB.Close()
}
