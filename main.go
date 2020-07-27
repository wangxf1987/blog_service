package main

import (
	"blog_service/internal/model"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	klogv2 "k8s.io/klog/v2"

	"blog_service/global"
	"blog_service/internal/routers"
	"blog_service/pkg/setting"
)

func init() {
	if err := setupSetting(); err != nil {
		klogv2.Fatalf("init setup setting error: %v", err)
	}
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
		klogv2.Infof("start http server error: %v", err)
		os.Exit(1)
	}
}
