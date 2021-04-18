package model

import (
	"github.com/sfshf/sprout/pkg/crypto/hash"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

func UpperStringPtr(s string) *string {
	u := strings.ToUpper(s)
	return &u
}

func StringPtr(s string) *string {
	return &s
}

func NewPasswdSalt() string {
	key := PasswdSalt + time.Now().String()
	return hash.MD5StringIgnorePrefixAndError(key)
}

func PasswdPtr(passwd string, salt string) *string {
	data := hash.MD5StringIgnorePrefixAndError(salt + passwd)
	return &data
}

// DatetimePtr get a pointer from a millisecond-unit timestamp.
func DatetimePtr(ts int64) *primitive.DateTime {
	dt := primitive.NewDateTimeFromTime(time.Unix(0, ts*1e6))
	return &dt
}

func NewDatetime(t time.Time) *primitive.DateTime {
	dt := primitive.NewDateTimeFromTime(t)
	return &dt
}
