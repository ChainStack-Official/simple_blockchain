package miner

import "github.com/ChainStack-Official/simple_blockchain/core/block"

type TaskMaster interface {
	// 获取最近的链上的块
	GetLastBlock() block.Block
	// 获取最近的一个要挖的块
	GetLatestTask() block.Block
	// 获取难度值
	GetCurDifficulty() int
	// 提交新的挖掘结果
	SubmitNewBlock(b block.Block)
}
