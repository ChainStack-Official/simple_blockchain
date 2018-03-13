package main

import (
	"os"
	"simple_blockchain/handler"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化全局日志
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(true)))
}

func main() {
	log.Info("启动程序")

	// http服务
	router := gin.Default()
	// 提交一个msg（交易）
	router.POST("/msg", handler.NewMsgHandler)
	// 新建区块
	router.POST("/block", handler.NewBlockHandler)
	// 查看当前链上的所有块
	router.GET("/block", handler.GetBlocksHandler)
	// 修改难度值
	router.GET("/mine_conf", handler.ChangeDifficultyHandler)
	// 启动挖矿
	router.GET("/start_mine", handler.StartMineHandler)
	// 停止挖矿
	router.GET("/stop_mine", handler.StopMineHandler)
	router.Run(":3000")
}
