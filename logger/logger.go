package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"taskman4/configs"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Debug *log.Logger
)

func Init() {
	fileInfo, err := os.OpenFile(configs.AppSettings.AppParams.LogInfo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}

	fileError, err := os.OpenFile(configs.AppSettings.AppParams.LogError, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	
	fileDebug, err := os.OpenFile(configs.AppSettings.AppParams.LogDebug, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}

	Info = log.New(fileInfo, "", log.Ldate|log.Lmicroseconds)
	Error = log.New(fileError, "", log.Ldate|log.Lmicroseconds)
	Debug = log.New(fileDebug, "", log.Ldate|log.Lmicroseconds)

	lumberLogInfo := &lumberjack.Logger{
		Filename:   configs.AppSettings.AppParams.LogInfo,
		MaxSize:    configs.AppSettings.AppParams.LogMaxSize, // megabytes
		MaxBackups: configs.AppSettings.AppParams.LogMaxBackups,
		MaxAge:     configs.AppSettings.AppParams.LogMaxAge,   //days
		Compress:   configs.AppSettings.AppParams.LogCompress, // disabled by default
		LocalTime:  true,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   configs.AppSettings.AppParams.LogError,
		MaxSize:    configs.AppSettings.AppParams.LogMaxSize, // megabytes
		MaxBackups: configs.AppSettings.AppParams.LogMaxBackups,
		MaxAge:     configs.AppSettings.AppParams.LogMaxAge,   //days
		Compress:   configs.AppSettings.AppParams.LogCompress, // disabled by default
		LocalTime:  true,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   configs.AppSettings.AppParams.LogDebug,
		MaxSize:    configs.AppSettings.AppParams.LogMaxSize, // megabytes
		MaxBackups: configs.AppSettings.AppParams.LogMaxBackups,
		MaxAge:     configs.AppSettings.AppParams.LogMaxAge,   //days
		Compress:   configs.AppSettings.AppParams.LogCompress, // disabled by default
		LocalTime:  true,
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	Info.SetOutput(gin.DefaultWriter)
	Error.SetOutput(lumberLogError)
	Debug.SetOutput(lumberLogDebug)
}
