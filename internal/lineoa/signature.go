package lineoa

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func VerifySignature(channelSecret string, body []byte, xLineSignature string) bool {
	mac := hmac.New(sha256.New, []byte(channelSecret))
	mac.Write(body)
	expected := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(xLineSignature))
}
