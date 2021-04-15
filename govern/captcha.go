package main

import (
	b64Captcha "github.com/mojocn/base64Captcha"
	"github.com/sfshf/sprout/govern/config"
)

func NewPictureCaptcha() *b64Captcha.Captcha {
	c := config.C.PicCaptcha
	driver := b64Captcha.NewDriverDigit(c.Height, c.Width, c.Length, c.MaxSkew, c.DotCount)
	var store b64Captcha.Store
	if c.RedisStore {
		// TODO Redis store.
	} else {
		store = b64Captcha.NewMemoryStore(c.Threshold, c.Expiration)
	}
	return b64Captcha.NewCaptcha(driver, store)
}
