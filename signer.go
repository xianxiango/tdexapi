package tdex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Signer signs provided payloads.
type Signer interface {
	// Sign signs provided payload and returns encoded string sum.
	Sign(payload []byte, data []byte) string
}

// HmacSigner uses HMAC SHA256 for signing payloads.
type HmacSigner struct {
	Key []byte
}

// Sign signs provided payload and returns encoded string sum.
func (hs *HmacSigner) Sign(payload []byte, data []byte) string {
	mac := hmac.New(sha256.New, hs.Key)
	mac.Write(payload)
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}
