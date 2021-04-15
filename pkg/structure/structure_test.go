package structure

import (
	"testing"
)

// Shallow tests:

func TestCopy(t *testing.T) {
	type UserSchema struct {
		Name    string
		Age     uint
		Address string
	}
	type UserModel struct {
		ID       uint
		Name     string
		Age      uint
		WhereNow string
	}
	from := UserModel{
		ID:       1,
		Name:     "Gavin",
		Age:      28,
		WhereNow: "at home",
	}
	to := &UserSchema{}
	if err := Copy(to, from); err != nil {
		t.Error(err)
	} else {
		t.Log(to)
	}

}
