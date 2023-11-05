package global

import (
	"github.com/qiuxiaori/go-blog/pkg/logger"
	"github.com/qiuxiaori/go-blog/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	Logger        *logger.Logger
	JWTSetting    *setting.JWTSettingS
)
