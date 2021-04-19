package validate

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
)

var (
	_once sync.Once

	_validator *defaultImpl
	_validate  *validator.Validate

	_uni      *ut.UniversalTranslator
	_errTrans ut.Translator
)

// Validate is a default interface for validating data.
type Validate interface {
	Struct(s interface{}) error
	StructCtx(ctx context.Context, s interface{}) error
	Var(field interface{}, tag string) error
	VarCtx(ctx context.Context, field interface{}, tag string) error
}

// WrappedValidator is a default interface for approaching the internal instances.
type WrappedValidator interface {
	Engine() Validate
	UniversalTranslator() interface{}
}

// IsWrapped used to switch to WrappedValidator interface, used to get the internal instance.
func IsWrapped(vali Validate) (WrappedValidator, bool) {
	if v, is := vali.(*defaultImpl); is {
		return v, is
	} else {
		return nil, is
	}
}

// IsValidate used to switch to Validate interface.
func IsValidate(wrapped WrappedValidator) (Validate, bool) {
	if v, is := wrapped.(*defaultImpl); is {
		return v, is
	} else {
		return nil, is
	}
}

type defaultImpl struct{}

// Validator return the default package validator.
func Validator() Validate {
	lazyinit()
	return _validator
}

func lazyinit() {
	_once.Do(func() {
		_validator = new(defaultImpl)
		_validate = validator.New()

		_uni = ut.New(en.New(), zh.New())
		_enTrans, _ := _uni.GetTranslator("en")
		_zhTrans, _ := _uni.GetTranslator("zh")
		en_translations.RegisterDefaultTranslations(_validate, _enTrans)
		zh_translations.RegisterDefaultTranslations(_validate, _zhTrans)

		_errTrans = _zhTrans

		_validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fmt.Sprintf("%q", fld.Tag.Get("comment"))
		})
	})
}

func translateValidationErrors(err error) error {
	if _errTrans == nil {
		return err
	}
	if errs, is := err.(validator.ValidationErrors); is {
		var b strings.Builder
		for _, e := range errs {
			if _, we := b.WriteString(e.Translate(_errTrans)); we != nil {
				return errs
			}
			if _, we := b.WriteRune('\n'); we != nil {
				return errs
			}
		}
		// Here, you can custom to type your parameter error.
		return errors.New(b.String())
	}
	return err
}

func (*defaultImpl) Struct(s interface{}) error {
	return translateValidationErrors(_validate.Struct(s))
}

func (*defaultImpl) StructCtx(ctx context.Context, s interface{}) error {
	return translateValidationErrors(_validate.StructCtx(ctx, s))
}

func (*defaultImpl) Var(field interface{}, tag string) error {
	return translateValidationErrors(_validate.Var(field, tag))
}

func (*defaultImpl) VarCtx(ctx context.Context, field interface{}, tag string) error {
	return translateValidationErrors(_validate.VarCtx(ctx, field, tag))
}

// If you want to register custom validations, suggest to use Engine to get the inner validate instance.
func (*defaultImpl) Engine() Validate {
	return _validate
}

func (*defaultImpl) UniversalTranslator() interface{} {
	return _uni
}
