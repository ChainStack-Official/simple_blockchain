package req_model

import (
	"github.com/ChainStack-Official/simple_blockchain/common/http_base"
	"github.com/ChainStack-Official/simple_blockchain/core/block"
)

type GetBlocksResp struct {
	http_base.BaseResp
	Blocks []block.Block `json:"blocks"`
}
