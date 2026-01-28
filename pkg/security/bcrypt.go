package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	return string(bytes), err
}

func CheckPasswordHash(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
