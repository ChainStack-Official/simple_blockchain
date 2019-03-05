package block

import (
	"github.com/ChainStack-Official/simple_blockchain/common/hash_util"
)

// 块
type Block struct {
	// 第几块
	Index int `json:"index"`
	// 块生成的时间戳
	Timestamp int64 `json:"timestamp"`
	// 提交的时候的时间戳
	SubmitTimestamp int64 `json:"tx_submit_timestamp"`
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

// 获取该块中用于hash的内容，TODO：用于hash的内容其实很重要
func (b *Block) getContentForHash() []byte {
	//TODO：验证是否双重花费，与链中交易的设计很有关，现在这里无法确认同一个msg被记了两次，还是确实是两个msg只是内容相同而已，这里可以通过提交时间戳来判断，但是如果提交时间戳没有经过用户签名，矿工还是可以修改该时间戳，因此用户签名的功能很重要

	//如果是用hash来验证块是否已存在链上，那么就不能用下边两个参数（当前项目中这两个参数在不同矿工那里的值是不同的）
	//indexStr := strconv.FormatInt(int64(b.Index), 10)
	//timeStr := strconv.FormatInt(b.Timestamp, 10)
	return []byte(b.Msg + b.PrevHash)
}

// 判断区块是否正确
//func (b *Block) IsValid() (error) {
//	switch {
//	case b.Timestamp == "":
//		return bcerr.GetError(bcerr.BlockTimestampIsBlankErr)
//	}
//	return nil
//}
