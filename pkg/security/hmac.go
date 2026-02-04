package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func (s *Security) HashToken(token string) string {
	mac := hmac.New(sha256.New, []byte(s.hmacCfg.TokenSecret))
	mac.Write([]byte(token))
	return hex.EncodeToString(mac.Sum(nil))
}
