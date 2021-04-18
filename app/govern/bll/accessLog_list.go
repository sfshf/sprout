package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type AccessLogListReq struct {
	Level     string          `form:"level" binding:""`
	TimeBegin time.Time       `form:"timeBegin" binding:""`
	TimeEnd   time.Time       `form:"timeEnd" binding:""`
	ClientIp  string          `form:"clientIp" binding:""`
	Path      string          `form:"path" binding:""`
	TraceId   string          `form:"traceId" binding:""`
	SessionId string          `form:"sessionId" binding:""`
	Tag       string          `form:"tag" binding:""`
	OrderBy   *schema.OrderBy `form:"orderBy" binding:""`
	schema.PaginationReq
}

type AccessLogListElem struct {
	ID              string    `json:"id"`
	Level           string    `json:"level"`
	Time            time.Time `json:"time"`
	ClientIp        string    `json:"clientIp"`
	Proto           string    `json:"proto"`
	Method          string    `json:"method"`
	Path            string    `json:"path"`
	Queries         string    `json:"queries"`
	RequestHeaders  string    `json:"requestHeaders"`
	RequestBody     string    `json:"requestBody"`
	StatusCode      string    `json:"statusCode"`
	ResponseHeaders string    `json:"responseHeaders"`
	ResponseBody    string    `json:"responseBody"`
	Latency         string    `json:"latency"`
	TraceId         string    `json:"traceId"`
	SessionId       string    `json:"sessionId"`
	Tag             string    `json:"tag"`
	Stack           string    `json:"stack"`
}

type AccessLogListResp struct {
	schema.PaginationResp
}

func (a *AccessLog) List(ctx context.Context, arg *AccessLogListReq, sort bson.M) (*AccessLogListResp, error) {
	var and bson.A
	if arg.Level != "" {
		and = append(and, bson.M{"level": arg.Level})
	}
	if !arg.TimeBegin.IsZero() {
		and = append(and, bson.M{"time": bson.M{"$gte": arg.TimeBegin}})
	}
	if !arg.TimeEnd.IsZero() {
		and = append(and, bson.M{"time": bson.M{"$lt": arg.TimeEnd}})
	}
	if arg.ClientIp != "" {
		and = append(and, bson.M{"clientIp": arg.ClientIp})
	}
	if arg.Path != "" {
		and = append(and, bson.M{"path": arg.Path})
	}
	if arg.TraceId != "" {
		and = append(and, bson.M{"traceId": arg.TraceId})
	}
	if arg.SessionId != "" {
		and = append(and, bson.M{"sessionId": arg.SessionId})
	}
	if arg.Tag != "" {
		and = append(and, bson.M{"tag": arg.Tag})
	}
	var filter bson.M
	if len(and) > 0 {
		filter = bson.M{"$and": and}
	}
	total, err := a.accessLog.CountByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	opt := options.Find().SetSort(sort).SetSkip(arg.PerPage * (arg.Page - 1)).SetLimit(arg.PerPage)
	res, err := a.accessLog.FindManyByFilter(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	data := make([]AccessLogListElem, 0, len(res))
	for _, v := range res {
		elem := AccessLogListElem{
			ID:              v.ID.Hex(),
			Level:           *v.Level,
			Time:            *v.Time,
			ClientIp:        *v.ClientIp,
			Proto:           *v.Proto,
			Method:          *v.Method,
			Path:            *v.Path,
			RequestHeaders:  *v.RequestHeaders,
			StatusCode:      *v.StatusCode,
			ResponseHeaders: *v.ResponseHeaders,
			ResponseBody:    *v.ResponseBody,
			Latency:         *v.Latency,
		}
		if v.Queries != nil {
			elem.Queries = *v.Queries
		}
		if v.RequestBody != nil {
			elem.RequestBody = *v.RequestBody
		}
		if v.TraceId != nil {
			elem.TraceId = *v.TraceId
		}
		if v.SessionId != nil {
			elem.SessionId = *v.SessionId
		}
		if v.Tag != nil {
			elem.Tag = *v.Tag
		}
		if v.Stack != nil {
			elem.Stack = *v.Stack
		}
		data = append(data, elem)
	}
	return &AccessLogListResp{
		schema.PaginationResp{
			Data:  data,
			Total: total,
		},
	}, nil
}
