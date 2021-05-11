package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/app/govern/schema"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GetPicCaptchaResp struct {
	PicCaptchaId   string `json:"picCaptchaId"`
	PicCaptchaB64s string `json:"picCaptchaB64s"`
}

type SignInReq struct {
	Account          string `json:"account" binding:"gte=2,lte=32"`
	Password         string `json:"password" binding:"gte=8,lte=32"`
	PicCaptchaId     string `json:"picCaptchaId" binding:"required"`
	PicCaptchaAnswer string `json:"picCaptchaAnswer" binding:"required"`
	Timestamp        int64  `json:"timestamp" binding:"gte=0"`
}

type SignInResp struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (a *Staff) GeneratePicCaptchaIdAndB64s(ctx context.Context, obsoleteId string) (*GetPicCaptchaResp, error) {
	_ = a.picCaptcha.Store.Get(obsoleteId, true)
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

func (a *Staff) SignIn(ctx context.Context, objId *primitive.ObjectID, ip *string, ts *primitive.DateTime) (*SignInResp, error) {
	token, expiresAt, err := a.auther.GenerateToken(objId.Hex())
	if err != nil {
		return nil, err
	}
	obj := &model.Staff{
		ID:             objId,
		SignInToken:    model.StringPtr(token),
		LastSignInIp:   ip,
		LastSignInTime: ts,
	}
	if err := a.staffRepo.UpdateOneByID(ctx, obj); err != nil {
		return nil, err
	}
	if !a.redisCache.Set(ctx, ginx.RedisKeyPrefix+objId.Hex(), token, time.Unix(0, expiresAt*1e6).Sub(time.Now())) {
		return nil, schema.ErrFailure
	}
	return &SignInResp{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}
