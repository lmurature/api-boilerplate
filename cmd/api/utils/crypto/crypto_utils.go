package crypto_utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	return bytes, err
}

func CheckPasswordHash(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
