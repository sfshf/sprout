package logger

import (
	"context"
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Logger ----------------------------------------------------------------------

type Logger struct {
	logger  zerolog.Logger
	timeLoc *time.Location
	version string
}

// SetWriters set the writers of the logger.
func NewLogger(writers ...io.Writer) *Logger {
	w := zerolog.MultiLevelWriter(writers...)
	return &Logger{
		logger:  zerolog.New(w),
		timeLoc: time.Local,
	}
}

// Set time field location.
func (a *Logger) SetTimeFieldLocation(locName string) error {
	loc, err := time.LoadLocation(locName)
	if err != nil {
		return errors.Wrap(err, "[logger] time field -- failed to set time field location")
	}
	a.timeLoc = loc
	return nil
}

// Set time field name.
func (a *Logger) SetTimeFieldName(name string) {
	zerolog.TimestampFieldName = name
}

// Set time field format.
func (a *Logger) SetTimeFieldFormat(layout string) {
	zerolog.TimeFieldFormat = layout
}

// SetVersion set version of the app that using this logger package.
func (a *Logger) SetVersion(v string) {
	a.version = v
}

// Info return logger event of info level created by default logger.
func (a *Logger) Info(ctx context.Context) *zerolog.Event {
	e := a.logger.Info().Time(zerolog.TimestampFieldName, time.Now().In(a.timeLoc))
	return e
}

// Error return logger event of error level created by default logger.
func (a *Logger) Error(ctx context.Context) *zerolog.Event {
	e := a.logger.Error().Time(zerolog.TimestampFieldName, time.Now().In(a.timeLoc))
	return e
}

// Mongo writer ----------------------------------------------------------------

// MongoWriter create a mongo writer to logger to.
func MongoWriter(coll *mongo.Collection) (io.Writer, error) {
	if coll == nil {
		return nil, errors.Errorf("[logger] mongo writer -- nil collection")
	}
	return &mongoWriter{
		coll: coll,
	}, nil
}

type mongoWriter struct {
	coll *mongo.Collection
}

// To implement io.Writer.
func (m *mongoWriter) Write(p []byte) (n int, err error) {
	fields := make(map[string]interface{})
	err = bson.UnmarshalExtJSON(p, false, fields)
	if err != nil {
		return -1, errors.Wrap(err, "[logger] mongo writer -- failed to unmarshal extended json")
	}
	_, err = m.coll.InsertOne(context.Background(), fields)
	if err != nil {
		return -1, errors.Wrap(err, "[logger] mongo writer -- failed to write to mongo")
	}
	return len(p), nil
}

// To implement zerolog.LevelWriter.
func (m *mongoWriter) WriteLevel(level *zerolog.Level, p []byte) (n int, err error) {
	return
}
