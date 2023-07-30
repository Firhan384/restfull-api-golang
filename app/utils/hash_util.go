package utils

import "golang.org/x/crypto/bcrypt"

const SALT = 14

func TextHash(text string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(text), SALT)
	return string(hash)
}

func CompareHash(text, hash string) bool {
	isCompare := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return isCompare != nil
}
