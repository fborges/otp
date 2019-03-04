package otp

import (
	"testing"
	"reflect"
)


func TestEncrypt(t *testing.T) {
	plaintext := []byte{1,2,3}
	key := []byte{82, 253, 252}

	cipher, _ := Encrypt(plaintext, key)
	expectedCipher := []byte{83, 255, 255}

	if !reflect.DeepEqual(cipher, expectedCipher) {
		t.Errorf("Expected: %v\nGot: %v", expectedCipher, cipher)
	}
}

func TestDecrypt(t *testing.T) {
	key := []byte{82, 253, 252}
	cipher := []byte{83, 255, 255}

	expectedPlain := []byte{1,2,3}
	plain, _ := Decrypt(cipher, key)

	if !reflect.DeepEqual(plain, expectedPlain) {
		t.Errorf("Expected: %v\nGot: %v", expectedPlain, plain)
	}
}

func TestGenerateKey(t *testing.T) {
	plain := []byte{3,138,43}	
	key := GenerateKey(plain)

	if len(key) != 3 {
		t.Errorf("Expected length: 3\nGot: %v", len(key))
	}
}