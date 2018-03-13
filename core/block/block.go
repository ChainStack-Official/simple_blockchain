package block

import (
	"simple_blockchain/common/hash_util"
	"strconv"
)

// 块
type Block struct {
	// 第几块
	Index int `json:"index"`
	// 块生成的时间戳
	Timestamp int64 `json:"timestamp"`
	// 当前块的内容
	Msg string `json:"msg"`
	// 当前块的hash
	Hash string `json:"hash"`
	// 上一个块的hash
	PrevHash string `json:"prev_hash"`
	// 挖矿时的难度
	Difficulty int `json:"difficulty"`
	// 挖出来时的随机数
	Nonce string `json:"nonce"`
}

// 获得当前区块的hash
func (b *Block) HashForThisBlock() string {
	return hash_util.HashForBlock(b.getContentForHash())
}

// 获取该块中用于hash的内容
func (b *Block) getContentForHash() []byte {
	indexStr := strconv.FormatInt(int64(b.Index), 10)
	timeStr := strconv.FormatInt(b.Timestamp, 10)
	return []byte(indexStr + timeStr + b.Msg + b.PrevHash)
}

// 判断区块是否正确
//func (b *Block) IsValid() (error) {
//	switch {
//	case b.Timestamp == "":
//		return bcerr.GetError(bcerr.BlockTimestampIsBlankErr)
//	}
//	return nil
//}
