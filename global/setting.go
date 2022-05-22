package global

import (
	"menah3m/blog-service/pkg/logger"
	"menah3m/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServiceSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)
