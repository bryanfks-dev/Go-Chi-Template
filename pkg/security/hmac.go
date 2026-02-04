package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

func (s *Security) HashToken(token string) string {
	algo := getHMACAlgorithm(s.hmacCfg.Algorithm)
	mac := hmac.New(algo, []byte(s.hmacCfg.TokenSecret))
	mac.Write([]byte(token))
	return hex.EncodeToString(mac.Sum(nil))
}

func (s *Security) CompareHashAndToken(hashedToken, token string) bool {
	algo := getHMACAlgorithm(s.hmacCfg.Algorithm)
	mac := hmac.New(algo, []byte(s.hmacCfg.TokenSecret))
	mac.Write([]byte(token))
	expectedMAC := mac.Sum(nil)
	hashedTokenBytes, err := hex.DecodeString(hashedToken)
	if err != nil {
		return false
	}
	return hmac.Equal(hashedTokenBytes, expectedMAC)
}

func getHMACAlgorithm(algo string) func() hash.Hash {
	val, ok := HMACAlgorithms[algo]
	if !ok {
		return sha256.New
	}
	return val
}
