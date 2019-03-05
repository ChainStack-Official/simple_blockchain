package ly

import (
	"errors"
	"fmt"
	"github.com/ChainStack-Official/simple_blockchain/core/blockchain"
	"github.com/ChainStack-Official/simple_blockchain/core/miner"
	"sync"

	"github.com/ethereum/go-ethereum/log"
)

var Bc *blockchain.Blockchain

var mineLock sync.Mutex
var miners []*miner.Miner

// TODO：service要做成实例，光是包无法满足某些功能
func init() {
	Bc = blockchain.NewBlockchain()
	// 监听新区块创建的事件，TODO：换成事件通知后这里就不需要做中转了
	go func() {
		for {
			select {
			case newBlockIndex := <-Bc.NewBlockCommittedChan:
				for _, m := range miners {
					m.NewBlockFoundChan <- newBlockIndex
				}
			}
		}
	}()
}

// 最多多少个矿工
const maxMinerCount = 4

// 启动挖矿
func StartMine(minerCount int) error {
	if minerCount <= 0 {
		return errors.New("矿工个数必须大于0")
	} else if minerCount > maxMinerCount {
		return fmt.Errorf("矿工个数不能大于%v个", maxMinerCount)
	}
	mineLock.Lock()
	defer mineLock.Unlock()
	if len(miners) > 0 {
		return fmt.Errorf("已经有矿工在挖矿了，当前矿工个数：%v", len(miners))
	}
	log.Info("启动挖矿", "miner count", minerCount)
	miners = []*miner.Miner{}
	for i := 0; i < minerCount; i++ {
		m := miner.NewMiner(i, Bc)
		miners = append(miners, m)
		// TODO：这里要加一个WaitGroup
		go func() {
			m.Run()
		}()
	}
	return nil
}

// 停止挖矿
func StopMine() error {
	mineLock.Lock()
	defer mineLock.Unlock()
	log.Info("停止挖矿")
	// 发出停止挖矿的通知
	for _, m := range miners {
		m.StopMineChan <- 1
	}
	miners = []*miner.Miner{}
	return nil
}
