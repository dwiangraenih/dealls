package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CheckPasswordHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		log.Println("error when compare hash and password")
		return false
	}
	return true
}

func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error when generate password")
		return "", err
	}
	return string(hash), nil
}
