package miner

import "simple_blockchain/core/block"

type TaskMaster interface {
	GetLatestTask() block.Block
	GetCurDifficulty() int
}
