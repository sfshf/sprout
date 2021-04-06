package model

import (
	"github.com/sfshf/sprout/pkg/util/hash"
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
	data, _ := hash.MD5([]byte(key), nil)
	return data
}

func PasswdPtr(passwd string, salt string) *string {
	data, _ := hash.MD5([]byte(salt+passwd), nil)
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
