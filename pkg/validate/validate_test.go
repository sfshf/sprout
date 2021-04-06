package validate

import (
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestStruct(t *testing.T) {
	form := struct {
		IDCard string `json:"id_card" validate:"len=18" comment:"身份证号码"`
		Name   string `json:"name" validate:"max=20" comment:"身份证姓名"`
		Phone  string `json:"phone" validate:"len=11" comment:"联系号码"`
	}{
		IDCard: "234515196705169875",
		Name:   "王大大",
		Phone:  "16987845874",
	}
	err := Validator().Struct(form)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("OK")
	}
	form.Name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	err = Validator().Struct(form)
	if err != nil {
		t.Log("YES: ", err)
	} else {
		t.Error("failed to validate")
	}
}

func TestStructCtx(t *testing.T) {
	form := struct {
		IDCard string `json:"id_card" validate:"len=18" comment:"身份证号码"`
		Name   string `json:"name" validate:"max=20" comment:"身份证姓名"`
		Phone  string `json:"phone" validate:"len=11" comment:"联系号码"`
	}{
		IDCard: "234515196705169875",
		Name:   "王大大",
		Phone:  "16987845874",
	}
	err := Validator().StructCtx(context.TODO(), form)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("OK")
	}
	form.Name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	form.Phone = "8748547"
	err = Validator().StructCtx(context.TODO(), form)
	if err != nil {
		t.Log("YES: ", err)
	} else {
		t.Error("failed to validate")
	}
}

func TestVar(t *testing.T) {
	form := struct {
		IDCard string `json:"id_card" validate:"len=18" comment:"身份证号码"`
		Name   string `json:"name" validate:"max=20" comment:"身份证姓名"`
		Phone  string `json:"phone" validate:"len=11" comment:"联系号码"`
	}{
		IDCard: "234515196705169875",
		Name:   "王大大",
		Phone:  "16987845874",
	}
	err := Validator().Var(form.IDCard, "len=10")
	if err != nil {
		t.Log("YES: ", err)
	} else {
		t.Error("failed to validate")
	}
}

func TestVarCtx(t *testing.T) {
	form := struct {
		IDCard string `json:"id_card" validate:"len=18" comment:"身份证号码"`
		Name   string `json:"name" validate:"max=20" comment:"身份证姓名"`
		Phone  string `json:"phone" validate:"len=11" comment:"联系号码"`
	}{
		IDCard: "234515196705169875",
		Name:   "王大大",
		Phone:  "16987845874",
	}
	err := Validator().VarCtx(context.TODO(), form.Name, "max=2")
	if err != nil {
		t.Log("YES: ", err)
	} else {
		t.Error("failed to validate")
	}
}

func TestIsWrappedAndIsValidate(t *testing.T) {
	if wrapped, is := IsWrapped(Validator()); is {
		t.Log("Is wrapped!")
		if _, is := IsValidate(wrapped); is {
			t.Log("Is Validate!")
		} else {
			t.Error("Now Validate!")
		}
		if _, is = wrapped.Engine().(*validator.Validate); is {
			t.Log("Is *validator.Validate")
		} else {
			t.Error("Not *validator.Validate")
		}
	} else {
		t.Error("Not wrapped!")
	}
}
