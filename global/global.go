package global

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/prclin/minimal-tiktok/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Configuration *config.Configuration
	Logger        *zap.SugaredLogger
	Datasource    *gorm.DB
	OSSClient     *oss.Client
)
