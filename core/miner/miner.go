package miner

import "simple_blockchain/core/block"

/*

矿工，持续找下一个区块，如果正在找时收到新区块被找到的消息，则开始找下一个
1. 矿工先到任务发送者那取最新的任务
2. 开始挖矿，挖矿成功或矿已经被别人挖到时，就开始获取新任务再继续挖矿
3. 停止挖矿

*/
type Miner struct {
	// 矿工的id
	Id int `json:"id"`
	// 用于接收新区块找到的消息（新区块找到了，如果index大于等于当前正在找的block则停止当前的查找，直接获取下一个任务）
	NewBlockFoundChan chan int
	StopMineChan      chan int
	// 从这里获取任务
	tm        TaskMaster
	lastBlock block.Block
}

// 新建一个矿工
func NewMiner(id int, tMaster TaskMaster) *Miner {
	return &Miner{
		Id:                id,
		NewBlockFoundChan: make(chan int),
		StopMineChan:      make(chan int),
		tm:                tMaster,
	}
}

// 启动挖矿
func (m *Miner) Run() {
	for {

	}
}

// 实施挖矿
func (m *Miner) mine() {

}
