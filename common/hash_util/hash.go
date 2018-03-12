package hash_util

import (
	"crypto/sha256"
	"encoding/hex"
)

// 计算块的hash
func HashForBlock(content []byte) string {
	h := sha256.New()
	h.Write(content)
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}