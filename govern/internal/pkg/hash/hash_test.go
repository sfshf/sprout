package hash

import (
	"testing"
)

func TestMD5(t *testing.T) {
	digest1, err := MD5([]byte(`testing MD5`), nil)
	if err != nil {
		t.Error(err)
	}
	digest2 := MD5StringIgnorePrefixAndError(`testing MD5`)
	if string(digest1) == digest2 {
		t.Log("EQUAL !!!")
	}
}

func TestSHA256(t *testing.T) {
	digest1, err := SHA256([]byte(`testing SHA256`), nil)
	if err != nil {
		t.Error(err)
	}
	digest2 := SHA256StringIgnorePrefixAndError(`testing SHA256`)
	if string(digest1) == digest2 {
		t.Log("EQUAL !!!")
	}
}
