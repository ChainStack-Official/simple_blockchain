package miner

import (
	"fmt"
	"github.com/ChainStack-Official/simple_blockchain/common/hash_util"
	"github.com/ChainStack-Official/simple_blockchain/core/block"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

/*

矿工，持续找下一个区块，如果正在找时收到新区块被找到的消息，则开始找下一个
1. 矿工先到任务发送者那取最新的任务
2. 开始挖矿，挖矿成功或矿已经被别人挖到时，就开始获取新任务再继续挖矿
3. 停止挖矿

*/
type Miner struct {
	// 矿工的id
	Id int
	// 用于接收新区块找到的消息（新区块找到了，如果index大于等于当前正在找的block则停止当前的查找，直接获取下一个任务）
	// chan中发过来的是block的index
	NewBlockFoundChan chan int
	// 接收停止挖矿的消息
	StopMineChan chan int
	// 提交挖到的新的块
	SubmitNewBlockChan chan block.Block
	// 从这里获取任务
	tm TaskMaster
	// 现在池中最后一个块
	lastBlock block.Block
	// 当前正在挖的这个块
	curMiningBlock block.Block
	// 缓存当前的难度值
	curDifficulty int
	// 挖矿从那个值开始（每个矿工从同一个起点开始挖是没有意义的）
	StartTryAt uint64
	// 记录算到哪个值了
	mineCounter uint64
}

// 新建一个矿工
func NewMiner(id int, tMaster TaskMaster) *Miner {
	m := &Miner{
		Id:                id,
		NewBlockFoundChan: make(chan int),
		StopMineChan:      make(chan int),
		tm:                tMaster,
		StartTryAt:        uint64(1) << uint(20*id),
	}
	m.refreshLatestTask()
	return m
}

const mineHeartBeat = 10 * time.Millisecond

// 启动挖矿
func (m *Miner) Run() {
	log.Info("启动挖矿", "miner id", m.Id, "start try at", m.StartTryAt)
	// 加timer，1. 是为了取外部事件，2. 是为了cpu不会被挖矿吃满，导致其它功能无法得到响应
	timer := time.NewTimer(mineHeartBeat)
	for {
		// 检查是否发现了新区块，
		select {
		// 已经产生了新区块，需要判断是否停止当前这个挖矿
		case newBlockIndex := <-m.NewBlockFoundChan:
			m.refreshLastBlockIfNeeded(newBlockIndex)
			// 收到停止挖矿的消息，停掉挖矿操作
		case <-m.StopMineChan:
			log.Info("停止挖矿", "miner id", m.Id)
			m.tm = nil
			return
		case <-timer.C:
		}
		// 刷新一下难度值，挖矿过程中难度值有可能改变
		m.RefreshCurDifficulty()
		// 执行挖矿
		m.mine()
		timer.Reset(mineHeartBeat)
	}
}

// 有新区块产生，因此需要刷新lastBlock
func (m *Miner) refreshLastBlockIfNeeded(newBlockIndex int) {
	// 如果新的块的索引小于当前正在挖的块则继续挖
	if newBlockIndex < m.lastBlock.Index+1 {
		return
	}
	// 如果新块的索引大于等于当前挖的这个块的索引，则说明别人已经找到了当前块的hash，不用再挖当前这个块了
	m.refreshLatestTask()
	log.Info("当前块已经被别人找到，重新获取最新的区块", "last block index", m.lastBlock.Index, "last block msg", m.lastBlock.Msg, "miner id", m.Id, "当前要计算的block的msg", m.curMiningBlock.Msg)
}

// 刷新最新的任务
func (m *Miner) refreshLatestTask() {
	m.lastBlock = m.tm.GetLastBlock()
	m.curMiningBlock = m.tm.GetLatestTask()
	m.mineCounter = m.StartTryAt
}

// 刷新难度值
func (m *Miner) RefreshCurDifficulty() {
	m.curDifficulty = m.tm.GetCurDifficulty()
}

// 一次间歇尝试多少次挖掘
const tryCountEveryHeartBeat = 10 ^ 4

// 实施挖矿
func (m *Miner) mine() {
	// 没有需要加入链中的区块
	if m.curMiningBlock.SubmitTimestamp == 0 {
		// 如果没有就要去刷新一下看下有没新的块
		m.refreshLatestTask()
		return
	}
	m.curMiningBlock.Index = m.lastBlock.Index + 1
	m.curMiningBlock.Timestamp = time.Now().Unix()
	m.curMiningBlock.PrevHash = m.lastBlock.Hash
	m.curMiningBlock.Difficulty = m.curDifficulty
	// 1. 找出nonce 2. 写上hash 3. 找到hash就要通知外边找到新的区块了
	max := m.mineCounter + tryCountEveryHeartBeat
	for ; m.mineCounter < max; m.mineCounter++ {
		// tmpNonce := hex.Dump([]byte(m.mineCounter))
		tmpNonce := fmt.Sprintf("%x", m.mineCounter)
		// 找到了
		if hash_util.IsValidMineNonce(tmpNonce, m.curDifficulty) {
			m.curMiningBlock.Nonce = tmpNonce
			m.curMiningBlock.Hash = m.curMiningBlock.HashForThisBlock()
			log.Info("挖到矿了", "miner id", m.Id, "b index", m.curMiningBlock.Index, "nonce", tmpNonce, "nonce hash", hash_util.HashForBlock([]byte(tmpNonce)), "difficulty", m.curDifficulty, "msg", m.curMiningBlock.Msg)
			// 提交该block
			m.tm.SubmitNewBlock(m.curMiningBlock)
			// 刷新任务，并结束本次计算
			m.refreshLatestTask()
			break
		}
	}
}
