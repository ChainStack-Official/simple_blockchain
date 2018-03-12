package blockchain

import (
	"simple_blockchain/core/block"
	"simple_blockchain/core/bcerr"
)

// 区块链
type Blockchain struct {
	// 当前链上的所有块
	Blocks []block.Block `json:"blocks"`
}

// 收到新的块
func (bc *Blockchain)NewBlock(b block.Block) error {
	if err := bc.IsValidNewBlock(b); err != nil {
		return err
	}
	bc.Blocks = append(bc.Blocks, b)
	return nil
}

/*

1. 检查索引是否正确
2. 检查上个区块的hash是否正确
3. 检查新区块的hash是否正确

*/
func (bc *Blockchain)IsValidNewBlock(newB block.Block) error {
	bcLen := len(bc.Blocks)
	if bcLen == 0 {
		return nil
	}
	lastBlock := bc.Blocks[bcLen - 1]
	if lastBlock.Index + 1 != newB.Index {
		return bcerr.GetError(bcerr.NewBlockIndexWrongErr)
	} else if lastBlock.Hash != newB.PrevHash {
		return bcerr.GetError(bcerr.NewBlockPrevHashWrongErr)
	} else if newB.HashForThisBlock() != newB.Hash {
		return bcerr.GetError(bcerr.NewBlockHashWrongErr)
	}
	return nil
}
