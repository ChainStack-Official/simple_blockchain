package handler

import (
	"errors"
	"github.com/ChainStack-Official/simple_blockchain/common/util"
	"github.com/ChainStack-Official/simple_blockchain/core/block"
	"github.com/ChainStack-Official/simple_blockchain/handler/req_model"
	"github.com/ChainStack-Official/simple_blockchain/ly"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

// TODO：handler可以考虑做成实例

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

// 新增一个内容记录到区块链中
func NewMsgHandler(c *gin.Context) {
	var req map[string]string
	if err := c.BindJSON(&req); err != nil {
		util.RenderGinJsonResult(c, util.GetBaseResp(err, ""))
		return
	}
	msg := req["content"]
	if msg == "" {
		util.RenderGinJsonResult(c, util.GetBaseResp(errors.New("内容不能为空"), ""))
		return
	}
	ly.Bc.AddBlockToNewBlocksPool(msg)
	util.RenderGinJsonResult(c, util.GetBaseResp(nil, "发起成功"))
}

// 启动挖矿
func StartMineHandler(c *gin.Context) {
	mCountStr := c.Query("miner_count")
	log.Debug("收到启动挖矿的请求", "count", mCountStr)
	mCount, _ := strconv.Atoi(mCountStr)
	if mCount == 0 {
		mCount = 1
	}
	err := ly.StartMine(mCount)
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
