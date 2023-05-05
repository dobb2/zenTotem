package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func SignHMAC(message string, key string) string {

	byteMessage := []byte(message)
	byteKey := []byte(key)
	h := hmac.New(sha256.New, byteKey)
	h.Write(byteMessage)
	dst := h.Sum(nil)

	return fmt.Sprintf("%x", dst)
}
