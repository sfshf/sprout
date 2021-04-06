package log

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/pkg/errors"
	kinet "github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Context ---------------------------------------------------------------------

const (
	VersionKey = "version"
	TraceIDKey = "trace_id"
	UserIDKey  = "user_id"
	TagKey     = "tag"
	StackKey   = "stack"
)

var (
	version string
)

// SetVersion set version of the app that using this log package.
func SetVersion(v string) {
	version = v
}

type (
	traceIDKey struct{}
	userIDKey  struct{}
	tagKey     struct{}
	stackKey   struct{}
)

// NewTraceIDContext wrap a context with a traceID.
func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

// FromTraceIDContext get a traceID from a context, if exists.
func FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewUserIDContext wrap a context with a userID.
func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

// FromUserIDContext get a userID from a context, if exists.
func FromUserIDContext(ctx context.Context) string {
	v := ctx.Value(userIDKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewTagContext wrap a context with a tag.
func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

// FromTagContext get a tag from a context, if exists.
func FromTagContext(ctx context.Context) string {
	v := ctx.Value(tagKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewStackContext wrap a context with an error stack.
func NewStackContext(ctx context.Context, stack error) context.Context {
	return context.WithValue(ctx, stackKey{}, stack)
}

// FromStackContext get an error stack from a context, if exists.
func FromStackContext(ctx context.Context) error {
	v := ctx.Value(stackKey{})
	if v != nil {
		if s, ok := v.(error); ok {
			return s
		}
	}
	return nil
}

// Logger ----------------------------------------------------------------------

// Set time field name.
func SetTimeFieldName(name string) {
	_mutex.Lock()
	kinet.TimestampFieldName = name
	_mutex.Unlock()
}

// Set time field location.
func SetTimeFieldLocation(locName string) error {
	loc, err := time.LoadLocation(locName)
	if err != nil {
		return errors.Wrap(err, "[log] time field -- failed to set time field location")
	}
	_mutex.Lock()
	_timeFieldLocation = loc
	_mutex.Unlock()
	return nil
}

// Set time field format.
func SetTimeFieldFormat(layout string) {
	_mutex.Lock()
	kinet.TimeFieldFormat = layout
	_mutex.Unlock()
}

// New create a new logger.
func New(w io.Writer) kinet.Logger {
	return kinet.New(w)
}

var (
	_stdlog kinet.Logger
	_once   sync.Once

	_mutex             sync.Mutex
	_timeFieldLocation = time.Local
)

// SetStdlog set the writers of default logger.
func SetStdlog(writers ...io.Writer) {
	_once.Do(func() {
		setStdlog(writers...)
	})
}

// setStdlog convenient for testing.
func setStdlog(writers ...io.Writer) {
	w := kinet.MultiLevelWriter(writers...)
	_stdlog = kinet.New(w)
}

// Info return log event of info level created by default logger.
func Info(ctx context.Context) *kinet.Event {
	e := _stdlog.Info().Time(kinet.TimestampFieldName, time.Now().In(_timeFieldLocation))
	return withContext(ctx, e)
}

// Error return log event of error level created by default logger.
func Error(ctx context.Context) *kinet.Event {
	e := _stdlog.Error().Time(kinet.TimestampFieldName, time.Now().In(_timeFieldLocation))
	return withContext(ctx, e)
}

// withContext a helper function to set log contexts.
func withContext(ctx context.Context, e *kinet.Event) *kinet.Event {
	if ctx == nil {
		ctx = context.Background()
	}
	if version != "" {
		e = e.Str(VersionKey, version)
	}
	if v := FromTraceIDContext(ctx); v != "" {
		e = e.Str(TraceIDKey, v)
	}
	if v := FromUserIDContext(ctx); v != "" {
		e = e.Str(UserIDKey, v)
	}
	if v := FromTagContext(ctx); v != "" {
		e = e.Str(TagKey, v)
	}
	if v := FromStackContext(ctx); v != nil {
		e = e.Errs(StackKey, []error{v})
	}
	return e
}

// Mongo writer ----------------------------------------------------------------

// MongoWriter create a mongo writer to log to.
func MongoWriter(coll *mongo.Collection, buf uint) (io.Writer, error) {
	if coll == nil {
		return nil, errors.Errorf("[log] mongo writer -- nil collection")
	}
	var bufNum int
	if buf > 100 {
		bufNum = 100
	} else {
		bufNum = int(buf)
	}
	return &mongoWriter{
		coll:   coll,
		bufNum: bufNum,
		buf:    make([]interface{}, 0, buf),
	}, nil
}

type mongoWriter struct {
	coll   *mongo.Collection
	bufNum int
	buf    []interface{}
	mu     sync.Mutex
}

// To implement io.Writer.
func (m *mongoWriter) Write(p []byte) (n int, err error) {
	fields := make(map[string]interface{})
	err = bson.UnmarshalExtJSON(p, false, fields)
	if err != nil {
		return -1, errors.Wrap(err, "[log] mongo writer -- failed to unmarshal extended json")
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	m.buf = append(m.buf, bson.M(fields))
	if len(m.buf) >= m.bufNum {
		_, err = m.coll.InsertMany(context.TODO(), m.buf)
		if err != nil {
			return -1, errors.Wrap(err, "[log] mongo writer -- failed to write to mongo")
		}
		m.buf = m.buf[:0]
	}
	return len(p), nil
}

// To implement zerolog.LevelWriter.
func (m *mongoWriter) WriteLevel(level *kinet.Level, p []byte) (n int, err error) {
	return
}
