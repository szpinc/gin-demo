package bootstrap

import (
	"gin-demo/global"
	"gin-demo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	level   zapcore.Level // 日志等级
	options []zap.Option  // 配置项
)

func InitLog() *zap.Logger {
	// 创建目录
	createRootDir()
	// 设置日志等级
	setLevel()

	if global.App.Config.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}
	return zap.New(getZapCore(), options...)
}

func createRootDir() {
	if exists, _ := utils.PathExists(global.App.Config.Log.RootDir); !exists {
		_ = os.Mkdir(global.App.Config.Log.RootDir, os.ModePerm)
	}
}

func setLevel() {
	switch global.App.Config.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder
	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.Config.App.Env + "." + l.String())
	}

	if global.App.Config.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewCore(encoder, getLogWriter(), level)
}

func getLogWriter() zapcore.WriteSyncer {
	file := lumberjack.Logger{
		Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Log.FileName,
		MaxSize:    global.App.Config.Log.MaxSize,
		MaxBackups: global.App.Config.Log.MaxBackups,
		Compress:   global.App.Config.Log.Compress,
	}
	return zapcore.AddSync(&file)
}
