package hash

import (
	"testing"
)

// Shallow tests:

func TestMD4(t *testing.T) {
	digest, err := MD4([]byte(`testing MD4`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}

}

func TestMD5(t *testing.T) {
	digest, err := MD5([]byte(`testing MD5`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA1(t *testing.T) {
	digest, err := SHA1([]byte(`testing SHA1`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA224(t *testing.T) {
	digest, err := SHA224([]byte(`testing SHA224`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA256(t *testing.T) {
	digest, err := SHA256([]byte(`testing SHA256`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA384(t *testing.T) {
	digest, err := SHA384([]byte(`testing SHA384`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA512(t *testing.T) {
	digest, err := SHA512([]byte(`testing SHA512`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestMD5SHA1(t *testing.T) {
	digest, err := MD5SHA1([]byte(`testing MD5SHA1`), nil)
	if err != nil {
		t.Log(err)
	} else {
		t.Error("implemented?!", digest)
	}

}

func TestRIPEMD160(t *testing.T) {
	digest, err := RIPEMD160([]byte(`testing RIPEMD160`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA3_224(t *testing.T) {
	digest, err := SHA3_224([]byte(`testing SHA3_224`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA3_256(t *testing.T) {
	digest, err := SHA3_256([]byte(`testing SHA3_256`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA3_384(t *testing.T) {
	digest, err := SHA3_384([]byte(`testing SHA3_384`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA3_512(t *testing.T) {
	digest, err := SHA3_512([]byte(`testing SHA3_512`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestMSHA512_224(t *testing.T) {
	digest, err := SHA512_224([]byte(`testing SHA512_224`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestSHA512_256(t *testing.T) {
	digest, err := SHA512_256([]byte(`testing SHA512_256`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestBLAKE2s_256(t *testing.T) {
	digest, err := BLAKE2s_256([]byte(`testing BLAKE2s_256`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestBLAKE2b_256(t *testing.T) {
	digest, err := BLAKE2b_256([]byte(`testing BLAKE2b_256`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestBLAKE2b_384(t *testing.T) {
	digest, err := BLAKE2b_384([]byte(`testing BLAKE2b_384`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}

func TestBLAKE2b_512(t *testing.T) {
	digest, err := BLAKE2b_512([]byte(`testing BLAKE2b_512`), nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(digest)
	}
}
