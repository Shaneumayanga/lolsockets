package lolsockets

import (
	"crypto/sha1"
	"encoding/base64"
)

const magicStr = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

func computeAcceptKey(challengeKey string) string {
	h := sha1.New()
	h.Write([]byte(challengeKey))
	h.Write([]byte(magicStr))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func isValidChallengeKey(s string) bool {
	if s == "" {
		return false
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	return err == nil && len(decoded) == 16
}

func isEmpty(str ...string) bool {
	for _, val := range str {
		if val == "" {
			return false
		}
	}
	return true
}
