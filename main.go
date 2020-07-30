package main

import (
	"blog_service/internal/model"
	"blog_service/pkg/logger"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"

	"blog_service/global"
	"blog_service/internal/routers"
	"blog_service/pkg/setting"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init setup setting error: %v", err)
	}
	if err := setupLogger(); err != nil {
		log.Fatalf("init setup logger error: %v", err)
	}
	/*	if err := setupDBEngine(); err != nil {
		log.Fatalf("init setup db engine error: %v", err)
	}*/
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + "/" + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	if err = setting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

// @title blog service demo
// @version 1.0
// @description blog service demo
// @termOfService https://github.com/wangxf1987/blog_service
func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()
	s := &http.Server{
		Addr:              fmt.Sprintf(":%s", global.ServerSetting.HttpPort),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       global.ServerSetting.ReadTimeout,
		ReadHeaderTimeout: global.ServerSetting.WriteTimeout,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start http server error: %v", err)
		os.Exit(1)
	}
}
