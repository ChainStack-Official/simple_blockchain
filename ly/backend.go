package ly

import "simple_blockchain/core/blockchain"

var Bc *blockchain.Blockchain

func init() {
	Bc = blockchain.NewBlockchain()
}

