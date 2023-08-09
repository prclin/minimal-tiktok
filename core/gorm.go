package core

import (
	"github.com/prclin/minimal-tiktok/config"
	"github.com/prclin/minimal-tiktok/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

/*
GormLogger 包裹zap的SugaredLogger作为Gorm的Logger
*/
type GormLogger struct {
	Logger *zap.SugaredLogger
}

func (gl GormLogger) Printf(s string, i ...interface{}) {
	gl.Logger.Infof(strings.ReplaceAll(s, "\u001b", " "), i...)
}

/*
initGormMysql 初始化mysql连接
*/
func initGormMysql() {
	mysqlConfig := config.DefaultConfiguration.Datasource.Mysql

	//将zap的WriteSyncer作为Gorm的日志输出位置
	gormLogger := logger.New(GormLogger{Logger: global.Logger}, logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		//必须关闭，否则日志输出会乱码
		Colorful: false,
	})

	datasource, err := gorm.Open(mysql.Open(mysqlConfig.DSN), &gorm.Config{
		Logger:         gormLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(err)
	}
	mysqlDb, err := datasource.DB()
	mysqlDb.SetMaxIdleConns(mysqlConfig.MaxIdleConnections)
	mysqlDb.SetMaxOpenConns(mysqlConfig.MaxOpenConnections)
	mysqlDb.SetConnMaxLifetime(time.Duration(mysqlConfig.ConnectionMaxLifetime) * time.Second)
	if err = mysqlDb.Ping(); err != nil {
		panic(err)
	}
	global.Datasource = datasource
}
