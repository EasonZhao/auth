package app

import "github.com/securityin/auth/pkg/logging"

// MarkError 将错误 存入日志
func MarkError(v ...interface{}) {
	logging.GetLogger().Error(v...)
	return
}
