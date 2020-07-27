package main

import (
	"os"

	klogv2 "k8s.io/klog/v2"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	if err := r.Run(); err != nil {
		klogv2.Infof("start gin service error: %v", err)
		os.Exit(1)
	}
}
