package hash

import (
	"crypto/md5"
	_ "crypto/md5"
	"crypto/sha256"
	_ "crypto/sha256"
	"fmt"
)

// MD5 -----------------------------------------------------------

func MD5StringIgnorePrefixAndError(data string) string {
	s, _ := MD5String(data, "")
	return s
}

func MD5String(data, prefix string) (string, error) {
	byts, err := MD5([]byte(data), []byte(prefix))
	return fmt.Sprintf("%x", byts), err
}

func MD5(data, prefix []byte) ([]byte, error) {
	h := md5.New()
	_, err := h.Write(data)
	if err != nil {
		return nil, err
	}
	return h.Sum(prefix), nil
}

// SHA256 --------------------------------------------------------

func SHA256StringIgnorePrefixAndError(data string) string {
	s, _ := SHA256String(data, "")
	return s
}

func SHA256String(data, prefix string) (string, error) {
	byts, err := SHA256([]byte(data), []byte(prefix))
	return fmt.Sprintf("%x", byts), err
}

func SHA256(data, prefix []byte) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write(data)
	if err != nil {
		return nil, err
	}
	return h.Sum(prefix), nil
}
