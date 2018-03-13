package blockchain

import (
	"errors"
	"simple_blockchain/common/hash_util"
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
	// 新的未处理的区块池（在挖矿赶不上新的交易产生时，就需要将未处理的交易放这里）
	// TODO：block本身还需要一个机制来保证矿工不会修改其信息（私钥签名在这起的作用）。如果不同矿工使用了同一个PrevHash对几个不同的合法交易记账，则可能导致分叉
	newBlocksPool []block.Block
	// TODO：全网的难度是怎么达成共识的？
	curDifficulty int

	lock              sync.RWMutex
	dfLock            sync.Mutex
	newBlocksPoolLock sync.Mutex

	// 用于发出新的区块提交的通知
	NewBlockCommittedChan chan int
}

// 添加新的交易到待记录的池子中
func (bc *Blockchain) AddBlockToNewBlocksPool(msg string) {
	bc.newBlocksPoolLock.Lock()
	defer bc.newBlocksPoolLock.Unlock()
	bc.newBlocksPool = append(bc.newBlocksPool, block.Block{Msg: msg, SubmitTimestamp: time.Now().Unix()})
}

// 获取链上的最后一个区块
func (bc *Blockchain) GetLastBlock() block.Block {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	return bc.blocks[len(bc.blocks)-1]

}

// 返回最近一个需要挖的区块任务
func (bc *Blockchain) GetLatestTask() (taskB block.Block) {
	bc.newBlocksPoolLock.Lock()
	defer bc.newBlocksPoolLock.Unlock()

	if len(bc.newBlocksPool) > 0 {
		return bc.newBlocksPool[0]
	}
	return
}

// 提交新的挖掘结果
func (bc *Blockchain) SubmitNewBlock(b block.Block) {
	if err := bc.NewBlock(b); err != nil {
		log.Warn("收到新的区块，但是将其加入链中时报错", "err", err.Error())
	}
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
		blocks:                []block.Block{genesis},
		curDifficulty:         1,
		NewBlockCommittedChan: make(chan int),
	}
}

func (bc *Blockchain) GetBlocks() []block.Block {
	return bc.blocks
}

// 收到新的块，curDifficulty：当前的难度值
func (bc *Blockchain) NewBlock(b block.Block) error {
	bc.lock.Lock()
	defer bc.lock.Unlock()
	bc.newBlocksPoolLock.Lock()
	defer bc.newBlocksPoolLock.Unlock()

	if err := bc.IsValidNewBlock(b); err != nil {
		return err
	}
	bc.blocks = append(bc.blocks, b)
	// 发出新区块确认的通知，TODO：这里用事件机制实现的话就不用一层一层传递了
	bc.NewBlockCommittedChan <- b.Index
	log.Info("确认新块", "b index", b.Index, "msg", b.Msg)
	// 检查这个块是否来自当前的交易池，如果是则要将其从池子中移除
	if len(bc.newBlocksPool) > 0 && bc.newBlocksPool[0].SubmitTimestamp == b.SubmitTimestamp {
		log.Info("从交易池中移除新增的块", "block index", b.Index, "block msg", b.Msg)
		bc.newBlocksPool = bc.newBlocksPool[1:]
	} else {
		log.Warn("新增的块不是来自于当前的交易池", "block msg", b.Msg)
	}
	return nil
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
		// 检查挖出的随机数是否正确
	} else if !hash_util.IsValidMineNonce(newB.Nonce, newB.Difficulty) {
		return errors.New("随机数不符合难度要求")
	} else if newB.HashForThisBlock() != newB.Hash {
		return bcerr.GetError(bcerr.NewBlockHashWrongErr)
	}
	return nil
}
