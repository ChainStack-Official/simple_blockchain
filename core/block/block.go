package block

import "simple_blockchain/common/hash_util"

// 块
type Block struct {
	// 第几块
	Index     int `json:"index"`
	// 块生成的时间戳
	Timestamp string `json:"timestamp"`
	// 当前块的内容
	Msg string `json:"msg"`
	// 当前块的hash
	Hash      string `json:"hash"`
	// 上一个块的hash
	PrevHash  string `json:"prev_hash"`
}

// 获得当前区块的hash
func (b *Block) HashForThisBlock() (string) {
	return hash_util.HashForBlock([]byte(b.getContentForHash()))
}

// 获取该块中用于hash的内容
func (b *Block) getContentForHash() string {
	return string(b.Index) + b.Timestamp + b.Msg + b.PrevHash
}

// 判断区块是否正确
//func (b *Block) IsValid() (error) {
//	switch {
//	case b.Timestamp == "":
//		return bcerr.GetError(bcerr.BlockTimestampIsBlankErr)
//	}
//	return nil
//}
