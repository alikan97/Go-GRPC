package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	hashedPass := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))
	return hashedPass, nil
}

func ComparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(storedPassword, ".")
	salt, err := hex.DecodeString(pwsalt[1])

	if err != nil {
		return false, fmt.Errorf("unable to verify user password")
	}

	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)

	if err != nil {
		return false, fmt.Errorf("failed to decrypt password %w", err)
	}

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
