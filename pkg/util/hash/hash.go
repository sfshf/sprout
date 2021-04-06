package hash

import (
	"crypto"
	"fmt"

	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"github.com/pkg/errors"
	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
)

// MD4
func MD4(data, prefix []byte) (string, error) {
	return hash(crypto.MD4, data, prefix)
}

// MD5
func MD5(data, prefix []byte) (string, error) {
	return hash(crypto.MD5, data, prefix)
}

// SHA1
func SHA1(data, prefix []byte) (string, error) {
	return hash(crypto.SHA1, data, prefix)
}

// SHA224
func SHA224(data, prefix []byte) (string, error) {
	return hash(crypto.SHA224, data, prefix)
}

// SHA256
func SHA256(data, prefix []byte) (string, error) {
	return hash(crypto.SHA256, data, prefix)
}

// SHA384
func SHA384(data, prefix []byte) (string, error) {
	return hash(crypto.SHA256, data, prefix)
}

// SHA512
func SHA512(data, prefix []byte) (string, error) {
	return hash(crypto.SHA512, data, prefix)
}

// MD5SHA1 -- no implementation; MD5+SHA1 used for TLS RSA
func MD5SHA1(data, prefix []byte) (string, error) {
	return "", errors.New("[hash] no implementation")
}

// RIPEMD160
func RIPEMD160(data, prefix []byte) (string, error) {
	return hash(crypto.RIPEMD160, data, prefix)
}

// SHA3_224
func SHA3_224(data, prefix []byte) (string, error) {
	return hash(crypto.SHA3_224, data, prefix)
}

// SHA3_256
func SHA3_256(data, prefix []byte) (string, error) {
	return hash(crypto.SHA3_256, data, prefix)
}

// SHA3_384
func SHA3_384(data, prefix []byte) (string, error) {
	return hash(crypto.SHA3_384, data, prefix)
}

// SHA3_512
func SHA3_512(data, prefix []byte) (string, error) {
	return hash(crypto.SHA3_512, data, prefix)
}

// SHA512_224
func SHA512_224(data, prefix []byte) (string, error) {
	return hash(crypto.SHA512_224, data, prefix)
}

// SHA512_256
func SHA512_256(data, prefix []byte) (string, error) {
	return hash(crypto.SHA512_256, data, prefix)
}

// BLAKE2s_256
func BLAKE2s_256(data, prefix []byte) (string, error) {
	return hash(crypto.BLAKE2s_256, data, prefix)
}

// BLAKE2b_256
func BLAKE2b_256(data, prefix []byte) (string, error) {
	return hash(crypto.BLAKE2b_256, data, prefix)
}

// BLAKE2b_384
func BLAKE2b_384(data, prefix []byte) (string, error) {
	return hash(crypto.BLAKE2b_384, data, prefix)
}

// BLAKE2b_512
func BLAKE2b_512(data, prefix []byte) (string, error) {
	return hash(crypto.BLAKE2b_512, data, prefix)
}

// unexported function.
func hash(h crypto.Hash, data, prefix []byte) (string, error) {
	hash := h.New()
	hash.Reset()
	_, err := hash.Write(data)
	if err != nil {
		return "", errors.Wrap(err, "[hash] failed to write data")
	}
	return fmt.Sprintf("%x", hash.Sum(prefix)), nil
}
