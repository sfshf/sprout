package logger

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// Test like:
// $ go test . -v -count=2
// and:
// $ go test -bench . -count=2

func TestStdlogWithWriter(t *testing.T) {
	// Set to os.Stderr
	l := NewLogger(os.Stderr)

	l.Info(context.Background()).Str("foo", "bar").Msg("testing TestStdlogWithWriter")
}

func TestSetTimeFieldName(t *testing.T) {
	// Set to os.Stderr
	l := NewLogger(os.Stderr)

	l.SetTimeFieldName("timestamp")

	l.Info(context.Background()).Str("foo", "bar").Msg("testing TestSetTimeFieldName")
}

func TestSetTimeLocation(t *testing.T) {

	// Set to os.Stderr
	l := NewLogger(os.Stderr)

	l.SetTimeFieldLocation("Asia/Shanghai")

	l.Info(context.Background()).Str("foo", "bar").Msg("testing TestSetTimeLocation")
}

func TestSetTimeFieldForamt(t *testing.T) {
	// Set to os.Stderr
	l := NewLogger(os.Stderr)

	l.SetTimeFieldFormat("2006.01.02.15.04.05.999")

	l.Info(context.Background()).Str("foo", "bar").Msg("testing TestSetTimeFieldForamt")
}

func TestStdlogWithoutWriter(t *testing.T) {
	l := NewLogger()
	l.Info(context.Background()).Str("foo", "bar").Msg("testing TestStdlogWithoutWriter")
}

func TestSetVersion(t *testing.T) {
	l := NewLogger(os.Stderr)
	l.SetVersion("v0.0.1")
	l.Info(context.Background()).Str("foo", "bar").Msg("testing TestSetVersion")
}

func TestStdlogWithContexts(t *testing.T) {
	ctx := context.Background()
	ctx = NewUserIDContext(
		NewTraceIDContext(
			NewTagContext(ctx, "__test__"),
			"traceID_test",
		),
		"userID_test",
	)
	l := NewLogger(os.Stderr)
	l.Info(ctx).Str("foo", "bar").Msg("testing TestStdlogWithContexts")
}

func TestStdlogWithErrorStack(t *testing.T) {
	ctx := context.Background()
	ctx = NewStackContext(
		NewUserIDContext(
			NewTraceIDContext(
				NewTagContext(ctx, "__test__"),
				"traceID_test",
			),
			"userID_test",
		),
		errors.New("error_test"),
	)
	l := NewLogger(os.Stderr)
	l.Error(ctx).Str("foo", "bar").Msg("testing TestStdlogWithStack")
}

// func TestMongoWriter(t *testing.T) {
// 	ctx := context.TODO()
// 	ctx = NewUserIDContext(
// 		NewTraceIDContext(
// 			NewTagContext(ctx, "__test__"),
// 			"traceID_test",
// 		),
// 		"userID_test",
// 	)
//
// 	// Set mongo writer.
// 	cli, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = cli.Connect(context.TODO())
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	db := cli.Database("test")
// 	coll := db.Collection("logger")
// 	toMongo, err := MongoWriter(coll, 0)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	setStdlog(toMongo)
//
// 	Info(ctx).Str("foo", "bar").Msg("testing TestMongoWriter")
// }
//
// func TestStdlogWithMultiWriter(t *testing.T) {
// 	ctx := context.TODO()
// 	ctx = NewUserIDContext(
// 		NewTraceIDContext(
// 			NewTagContext(ctx, "__test__"),
// 			"traceID_test",
// 		),
// 		"userID_test",
// 	)
//
// 	// Set mongo writer.
// 	cli, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = cli.Connect(context.TODO())
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	db := cli.Database("test")
// 	coll := db.Collection("logger")
// 	toMongo, err := MongoWriter(coll, 0)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	setStdlog(os.Stderr, toMongo)
//
// 	Info(ctx).Str("foo", "bar").Msg("testing TestStdlogWithMultiWriter")
// }

func BenchmarkLogWithContexts(b *testing.B) {
	ctx := NewTagContext(
		NewUserIDContext(
			NewTraceIDContext(context.TODO(), "traceID_foo"),
			"userID_foo",
		),
		"tag_foo",
	)
	l := NewLogger(ioutil.Discard)
	l.SetVersion("v0.0.0")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info(ctx).Msg("benchmark")
		}
	})
}

func BenchmarkLogWithoutContexts(b *testing.B) {
	l := NewLogger(ioutil.Discard)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info(context.TODO()).
				Str("version", "v0.0.0").
				Str("trace_id", "traceID_foo").
				Str("user_id", "userID_foo").
				Str("tag", "tag_foo").
				Msg("benchmark")
		}
	})
}

func BenchmarkZeroLog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("version", "v0.0.0").
				Str("trace_id", "traceID_foo").
				Str("user_id", "userID_foo").
				Str("tag", "tag_foo").
				Msg("benchmark")
		}
	})
}

func BenchmarkZapLog(b *testing.B) {
	logger := zap.NewNop()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("benchmark",
				zap.Time("timestamp", time.Now()),
				zap.String("version", "v0.0.0"),
				zap.String("trace_id", "traceID_foo"),
				zap.String("user_id", "userID_foo"),
				zap.String("tag", "tag_foo"),
			)
		}
	})
}
