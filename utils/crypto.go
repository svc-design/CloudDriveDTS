package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(data []byte, passphrase string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	key := pbkdf2.Key([]byte(passphrase), salt, 100_000, 32, sha256.New)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}
	ciphertext := aesgcm.Seal(nonce, nonce, data, nil)
	result := append(salt, ciphertext...)
	return base64.StdEncoding.EncodeToString(result), nil
}

func Decrypt(enc string, passphrase string) ([]byte, error) {
	raw, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return nil, err
	}
	if len(raw) < 16 {
		return nil, errors.New("ciphertext too short")
	}
	salt := raw[:16]
	key := pbkdf2.Key([]byte(passphrase), salt, 100_000, 32, sha256.New)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesgcm.NonceSize()
	if len(raw) < 16+nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce := raw[16 : 16+nonceSize]
	ciphertext := raw[16+nonceSize:]
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}
