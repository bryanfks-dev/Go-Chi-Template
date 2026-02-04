package security

import "golang.org/x/crypto/bcrypt"

func (s *Security) HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), s.bcryptCfg.Cost)
	return string(bytes), err
}

func (s *Security) CheckPasswordHash(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
