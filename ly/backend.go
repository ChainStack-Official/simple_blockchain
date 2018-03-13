package ly

import (
	"errors"
	"fmt"
	"simple_blockchain/core/blockchain"
	"simple_blockchain/core/miner"
	"sync"

	"github.com/ethereum/go-ethereum/log"
)

var Bc *blockchain.Blockchain

var mineLock sync.Mutex
var miners []*miner.Miner

func init() {
	Bc = blockchain.NewBlockchain()
}

// 启动挖矿
func StartMine(minerCount int) error {
	if minerCount <= 0 {
		return errors.New("矿工个数必须大于0")
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
