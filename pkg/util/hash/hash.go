package hash

import (
	"crypto"
	_ "crypto/md5"
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
	return hash(crypto.MD5, data, prefix)
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
	return hash(crypto.SHA256, data, prefix)
}

// unexported function.
func hash(h crypto.Hash, data, prefix []byte) ([]byte, error) {
	hash := h.New()
	hash.Reset()
	_, err := hash.Write(data)
	if err != nil {
		return nil, fmt.Errorf("[hash] failed to write data: %s", err)
	}
	return hash.Sum(prefix), nil
}
