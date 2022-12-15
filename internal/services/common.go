package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"time"
)

const (
	salt     = "ghafbn6badgnevh54n"
	aesKey   = "207506173460746f263686e67656106120736737869732077264276f63726574"
	nonceLen = 12
)

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func encryptAES(data string) (string, error) {
	key, err := hex.DecodeString(aesKey)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, nonceLen)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherData := aesgcm.Seal(nil, nonce, []byte(data), nil)

	res := fmt.Sprintf("%s", append(nonce, cipherData...))

	return res, nil
}

func decryptAES(data []byte) (string, error) {
	key, err := hex.DecodeString(aesKey)
	if err != nil {
		return "", err
	}

	nonce, err := hex.DecodeString(fmt.Sprintf("%x", data[:nonceLen]))
	if err != nil {
		return "", err
	}

	cipherData, err := hex.DecodeString(fmt.Sprintf("%x", data[nonceLen:]))
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	decryptData, err := aesgcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", decryptData), err
}

func checkDate(day string) error {
	now := time.Now()
	now = now.Add(-(time.Hour * 24))

	workDay, err := time.Parse(dateFormat, day)
	if err != nil {
		return err
	}

	if now.After(workDay) {
		return errors.New("entered a date that has already passed")
	}
	return nil
}
