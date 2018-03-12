package handler

import (
	"github.com/gin-gonic/gin"
	"simple_blockchain/core/block"
	"simple_blockchain/common/util"
	"simple_blockchain/ly"
	"simple_blockchain/handler/req_model"
)

// 获取所有区块
func GetBlocksHandler(c *gin.Context) {
	util.RenderGinJsonResult(c, &req_model.GetBlocksResp{
		BaseResp: *util.GetBaseResp(nil, "获取成功"),
		Blocks: ly.Bc.GetBlocks(),
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
