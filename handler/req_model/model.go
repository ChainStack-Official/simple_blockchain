package req_model

import (
	"simple_blockchain/common/http_base"
	"simple_blockchain/core/block"
)

type GetBlocksResp struct {
	http_base.BaseResp
	Blocks []block.Block `json:"blocks"`
}
