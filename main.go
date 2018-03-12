package main

import (
	"github.com/ethereum/go-ethereum/log"
	"os"
	"github.com/gin-gonic/gin"
	"simple_blockchain/handler"
)

func init() {
	// 初始化全局日志
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(true)))
}

func main() {
	log.Info("启动程序")

	// http服务
	router := gin.Default()
	// 新建区块
	router.POST("/block", handler.NewBlockHandler)
	router.GET("/block", handler.GetBlocksHandler)
	router.Run(":3000")
}

