package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	klogv2 "k8s.io/klog/v2"

	"blog_service/internal/routers"
)

const HTTPPort = 8080

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", HTTPPort),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
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
