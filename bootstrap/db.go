package bootstrap

import (
	"fmt"
	"gin-demo/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

func InitDB() *gorm.DB {
	switch global.App.Config.Datasource.Driver {
	case "mysql":
		return initMysql()
	default:
		return initMysql()
	}
}

// 自定义gorm writer
func getGormWriter() logger.Writer {

	var writer io.Writer
	// 启用了文件记录日志
	if global.App.Config.Datasource.EnableFileLogWriter {
		writer = getLogWriter()
	} else {
		writer = os.Stdout
	}

	return log.New(writer, "\r\n", log.LstdFlags)
}

func initMysql() *gorm.DB {

	datasource := global.App.Config.Datasource

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		datasource.UserName,
		datasource.Password,
		datasource.Host,
		datasource.Port,
		datasource.Database,
		datasource.Charset,
	)

	config := mysql.Config{
		DSN: dsn,
	}

	if db, err := gorm.Open(mysql.New(config), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   gormLogger(),
	}); err != nil {
		global.App.Log.Error("mysql connect error", zap.Any("err", err))
		return nil
	} else {
		sqlDb, _ := db.DB()
		sqlDb.SetMaxIdleConns(datasource.MaxIdleConns)
		sqlDb.SetMaxOpenConns(datasource.MaxOpenConns)
		return db
	}
}

func gormLogger() logger.Interface {

	dbConfig := global.App.Config.Datasource

	var logMode logger.LogLevel

	switch dbConfig.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond, // 慢sql阈值
		Colorful:                  true,
		IgnoreRecordNotFoundError: false, // 忽略记录未找到错误
		LogLevel:                  logMode,
	})
}
