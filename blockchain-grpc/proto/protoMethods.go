package proto

import (
"crypto/sha256"
"encoding/hex"
)

func (b *Block) SetHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}
