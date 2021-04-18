package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/sfshf/sprout/pkg/crypto/hash"
	"io"
)

// https://tools.ietf.org/html/rfc5246
// https://tools.ietf.org/html/rfc5652

const (
	secret = `sprout:secret_key`
)

func secretKey(key []byte) ([]byte, error) {
	return hash.MD5([]byte(string(key)+secret), nil)
}

// PKCS5Padding/PKCS7Padding
func padding(plaintext []byte, size int) []byte {
	mod := len(plaintext) % size
	if mod > 0 {
		pad := size - mod
		pads := bytes.Repeat([]byte{byte(pad)}, pad)
		plaintext = append(plaintext, pads...)
	}
	return plaintext
}

func unpadding(plaintext []byte) []byte {
	l := len(plaintext)
	pad := plaintext[l-1]
	return plaintext[:l-int(pad)]
}

// AES16 + CBC
func AESCBCEncryptStringIgnoreError(plaintext, key string) string {
	ciphertext, _ := AESCBCEncryptString(plaintext, key)
	return ciphertext
}

func AESCBCEncryptString(plaintext, key string) (string, error) {
	ciphertext, err := AESCBCEncrypt([]byte(plaintext), []byte(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

func AESCBCEncrypt(plaintext, key []byte) ([]byte, error) {
	plaintext = padding(plaintext, aes.BlockSize)
	key, err := secretKey(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func AESCBCDecryptStringIgnoreError(ciphertext, key string) string {
	plaintext, _ := AESCBCDecryptString(ciphertext, key)
	return plaintext
}

func AESCBCDecryptString(ciphertext, key string) (string, error) {
	cipherbytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := AESCBCDecrypt(cipherbytes, []byte(key))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", plaintext), nil
}

func AESCBCDecrypt(ciphertext, key []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("cipher text too short")
	}
	key, err := secretKey(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("cipher text is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)
	plaintext := unpadding(ciphertext)
	return plaintext, nil
}

// AES16 + CFB

func AESCFBEncryptStringIgnoreError(plaintext, key string) string {
	ciphertext, _ := AESCFBEncryptString(plaintext, key)
	return ciphertext
}

func AESCFBEncryptString(plaintext, key string) (string, error) {
	ciphertext, err := AESCFBEncrypt([]byte(plaintext), []byte(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

func AESCFBEncrypt(plaintext, key []byte) ([]byte, error) {
	key, err := secretKey(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func AESCFBDecryptStringIgnoreError(ciphertext, key string) string {
	plaintext, _ := AESCFBDecryptString(ciphertext, key)
	return plaintext
}

func AESCFBDecryptString(ciphertext, key string) (string, error) {
	cipherbytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := AESCFBDecrypt(cipherbytes, []byte(key))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", plaintext), nil
}

func AESCFBDecrypt(ciphertext, key []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("cipher text too short")
	}
	key, err := secretKey(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

// AES16 + CTR

func AESCTREncryptStringIgnoreError(plaintext, key string) string {
	ciphertext, _ := AESCTREncryptString(plaintext, key)
	return ciphertext
}

func AESCTREncryptString(plaintext, key string) (string, error) {
	ciphertext, err := AESCTREncrypt([]byte(plaintext), []byte(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

func AESCTREncrypt(plaintext, key []byte) ([]byte, error) {
	key, err := secretKey(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func AESCTRDecryptStringIgnoreError(ciphertext, key string) string {
	plaintext, _ := AESCTRDecryptString(ciphertext, key)
	return plaintext
}

func AESCTRDecryptString(ciphertext, key string) (string, error) {
	cipherbytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := AESCTRDecrypt(cipherbytes, []byte(key))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", plaintext), nil
}

func AESCTRDecrypt(ciphertext, key []byte) ([]byte, error) {
	key, err := secretKey(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}

// AES16 + OFB

func AESOFBEncryptStringIgnoreError(plaintext, key string) string {
	ciphertext, _ := AESOFBEncryptString(plaintext, key)
	return ciphertext
}

func AESOFBEncryptString(plaintext, key string) (string, error) {
	ciphertext, err := AESOFBEncrypt([]byte(plaintext), []byte(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

func AESOFBEncrypt(plaintext, key []byte) ([]byte, error) {
	key, err := secretKey(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func AESOFBDecryptStringIgnoreError(ciphertext, key string) string {
	plaintext, _ := AESOFBDecryptString(ciphertext, key)
	return plaintext
}

func AESOFBDecryptString(ciphertext, key string) (string, error) {
	cipherbytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := AESOFBDecrypt(cipherbytes, []byte(key))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", plaintext), nil
}

func AESOFBDecrypt(ciphertext, key []byte) ([]byte, error) {
	key, err := secretKey(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}

// StreamReader: AES16 + OFB

// StreamWriter: AES16 + OFB

// AES16 + GCM
