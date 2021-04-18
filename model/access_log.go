package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccessLog struct {
	ID              *primitive.ObjectID `bson:"_id,omitempty"`
	Level           *string             `bson:"level,omitempty"`
	Time            *time.Time          `bson:"time,omitempty"`
	ClientIp        *string             `bson:"clientIp,omitempty"`
	Proto           *string             `bson:"proto,omitempty"`
	Method          *string             `bson:"method,omitempty"`
	Path            *string             `bson:"path,omitempty"`
	Queries         *string             `bson:"queries,omitempty"`
	RequestHeaders  *string             `bson:"requestHeaders,omitempty"`
	RequestBody     *string             `bson:"requestBody,omitempty"`
	StatusCode      *string             `bson:"statusCode,omitempty"`
	ResponseHeaders *string             `bson:"responseHeaders,omitempty"`
	ResponseBody    *string             `bson:"responseBody,omitempty"`
	Latency         *string             `bson:"latency,omitempty"`
	TraceId         *string             `bson:"traceId,omitempty"`
	SessionId       *string             `bson:"sessionId,omitempty"`
	Tag             *string             `bson:"tag,omitempty"`
	Stack           *string             `bson:"stack,omitempty"`
}

/*
VersionKey = "version"
	TraceIDKey = "trace_id"
	UserIDKey  = "user_id"
	TagKey     = "tag"
	StackKey   = "stack"
*/
