package blockchain

import (
	"errors"
	"simple_blockchain/core/bcerr"
	"simple_blockchain/core/block"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// 区块链
type Blockchain struct {
	// 当前链上的所有块
	blocks []block.Block
	// TODO：全网的难度是怎么达成共识的？
	curDifficulty int

	lock   sync.RWMutex
	dfLock sync.Mutex
	//
	//BcErrChan chan error
}

// 获取当前的难度值
func (bc *Blockchain) GetCurDifficulty() int {
	bc.dfLock.Lock()
	defer bc.dfLock.Unlock()

	return bc.curDifficulty
}

// 改变难度值
func (bc *Blockchain) ChangeDifficulty(newDifficulty int) {
	bc.dfLock.Lock()
	defer bc.dfLock.Unlock()

	bc.curDifficulty = newDifficulty
	log.Info("修改难度值", "value", bc.curDifficulty)
}

func NewBlockchain() *Blockchain {
	genesis := block.Block{
		Index:     1,
		Timestamp: time.Now().Unix(),
		Msg:       "我是创世块",
	}
	genesis.Hash = genesis.HashForThisBlock()
	return &Blockchain{
		blocks:        []block.Block{genesis},
		curDifficulty: 1,
	}
}

func (bc *Blockchain) GetBlocks() []block.Block {
	return bc.blocks
}

// 收到新的块，curDifficulty：当前的难度值
func (bc *Blockchain) NewBlock(b block.Block) error {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	if err := bc.IsValidNewBlock(b); err != nil {
		return err
	}
	bc.blocks = append(bc.blocks, b)
	return nil
}

// 返回最近的一个块
func (bc *Blockchain) GetLatestTask() block.Block {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	return bc.blocks[len(bc.blocks)-1]
}

/*

1. 检查索引是否正确
2. 检查上个区块的hash是否正确
3. 检查新区块的hash是否正确

*/
func (bc *Blockchain) IsValidNewBlock(newB block.Block) error {
	bcLen := len(bc.blocks)
	if bcLen == 0 {
		return nil
	}
	lastBlock := bc.blocks[bcLen-1]
	if lastBlock.Index+1 != newB.Index {
		return bcerr.GetError(bcerr.NewBlockIndexWrongErr)
	} else if lastBlock.Hash != newB.PrevHash {
		return bcerr.GetError(bcerr.NewBlockPrevHashWrongErr)
	} else if newB.Difficulty != bc.curDifficulty {
		return errors.New("新区块的难度值与当前难度不匹配")
	} else if newB.HashForThisBlock() != newB.Hash {
		return bcerr.GetError(bcerr.NewBlockHashWrongErr)
	}
	return nil
}
