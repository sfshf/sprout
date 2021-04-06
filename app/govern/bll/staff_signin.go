package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"github.com/sfshf/sprout/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetPicCaptchaResp struct {
	PicCaptchaId   string `json:"picCaptchaId"`
	PicCaptchaB64s string `json:"picCaptchaB64s"`
}

type SigninReq struct {
	Account          string `json:"account" binding:"gte=2,lte=32"`
	Password         string `json:"password" binding:"gte=8,lte=32"`
	PicCaptchaId     string `json:"picCaptchaId" binding:"required"`
	PicCaptchaAnswer string `json:"picCaptchaAnswer" binding:"required"`
	Timestamp        int64  `json:"timestamp" binding:"gte=0"`
}

type SigninResp struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (a *Staff) GeneratePicCaptchaIdAndB64s(ctx context.Context, id string) (*GetPicCaptchaResp, error) {
	_ = a.picCaptcha.Store.Get(id, true)
	id, b64s, err := a.picCaptcha.Generate()
	if err != nil {
		return nil, err
	}
	return &GetPicCaptchaResp{
		PicCaptchaId:   id,
		PicCaptchaB64s: b64s,
	}, nil
}

func (a *Staff) GetPicCaptchaAnswer(ctx context.Context, id string) string {
	return a.picCaptcha.Store.Get(id, false)
}

func (a *Staff) VerifyPictureCaptcha(ctx context.Context, id string, answer string) bool {
	if id == "" || answer == "" {
		return false
	}
	return a.picCaptcha.Verify(id, answer, true)
}

func (a *Staff) VerifyAccountAndPassword(ctx context.Context, account, password string) (*model.Staff, error) {
	staff, err := a.staffRepo.FindOneByAccount(ctx, account)
	if err != nil {
		return nil, err
	}
	if passwd := model.PasswdPtr(password, *staff.PasswordSalt); *passwd != *staff.Password {
		return nil, schema.ErrInvalidAccountOrPassword
	}
	return staff, nil
}

func (a *Staff) SignIn(ctx context.Context, id *primitive.ObjectID, ip *string, ts *primitive.DateTime) (*SigninResp, error) {
	token, expiresAt, err := a.auther.GenerateToken(id.Hex())
	if err != nil {
		return nil, err
	}
	tokenPtr := model.StringPtr(token)
	if err := a.staffRepo.SignIn(ctx, id, tokenPtr, ip, ts); err != nil {
		return nil, err
	}
	return &SigninResp{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}
