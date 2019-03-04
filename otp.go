package otp

import (
	"errors"
	"math/rand"
)

func Encrypt(plain, key []byte) ([]byte, error) {
	if len(key) != len(plain) {
		return []byte{}, errors.New("The plaintext and the key should have the same size.")
	}

	cipher := []byte{}

	for index, value := range plain {
		cipher = append(cipher, value ^ key[index])
	}

	return cipher, nil
}

func Decrypt(cipher, key []byte) ([]byte, error) {
	if len(key) != len(cipher) {
		return []byte{}, errors.New("The plaintext and the key should have the same size.")
	}

	plain := []byte{}

	for index, value := range cipher {
		plain = append(plain, value ^ key[index])
	}

	return plain, nil
}

func GenerateKey(plaintext []byte) []byte {
	length := len(plaintext)

	key := make([]byte, length)
	rand.Read(key)

	return key
}