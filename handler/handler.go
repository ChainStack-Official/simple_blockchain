package handler

import (
	"errors"
	"simple_blockchain/common/util"
	"simple_blockchain/core/block"
	"simple_blockchain/handler/req_model"
	"simple_blockchain/ly"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取所有区块
func GetBlocksHandler(c *gin.Context) {
	util.RenderGinJsonResult(c, &req_model.GetBlocksResp{
		BaseResp: *util.GetBaseResp(nil, "获取成功"),
		Blocks:   ly.Bc.GetBlocks(),
	})
}

// 新增区块
func NewBlockHandler(c *gin.Context) {
	var req block.Block
	if err := c.BindJSON(&req); err != nil {
		util.RenderGinJsonResult(c, util.GetBaseResp(err, ""))
		return
	}
	err := ly.Bc.NewBlock(req)
	util.RenderGinJsonResult(c, util.GetBaseResp(err, "新建成功"))
}

// 启动挖矿
func StartMineHandler(c *gin.Context) {
	mCountStr := c.Query("miner_count")
	mCount, _ := strconv.Atoi(mCountStr)
	if mCount == 0 {
		mCount = 1
	}
	err := ly.StartMine(1)
	util.RenderGinJsonResult(c, util.GetBaseResp(err, "开始挖矿了"))
}

// 停止挖矿
func StopMineHandler(c *gin.Context) {
	err := ly.StopMine()
	util.RenderGinJsonResult(c, util.GetBaseResp(err, "停止挖矿了"))
}

// 修改难度值
func ChangeDifficultyHandler(c *gin.Context) {
	difficultyStr := c.Query("difficulty")
	difficulty, _ := strconv.Atoi(difficultyStr)
	if difficulty <= 0 {
		util.RenderGinJsonResult(c, util.GetBaseResp(errors.New("难度值必须大于0"), ""))
		return
	}
	ly.Bc.ChangeDifficulty(difficulty)
	util.RenderGinJsonResult(c, util.GetBaseResp(nil, "修改成功"))
}
