package hash_util

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// 计算块的hash
func HashForBlock(content []byte) string {
	h := sha256.New()
	h.Write(content)
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// 检查是否是合格的Nonce
func IsValidMineNonce(nonce string, difficulty int) bool {
	return IsValidMineHash(HashForBlock([]byte(nonce)), difficulty)
}

// 检查这个hash是否是合格的hash
func IsValidMineHash(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}
