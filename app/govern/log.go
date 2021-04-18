package main

import (
	"github.com/sfshf/sprout/app/govern/config"
	"github.com/sfshf/sprout/pkg/logger"
	"github.com/sfshf/sprout/repo"
	"io"
	"os"
)

func NewLogger(repo *repo.AccessLog) (*logger.Logger, error) {
	c := config.C.Log
	var writers []io.Writer
	if !c.SkipStdout {
		writers = append(writers, os.Stderr)
	}
	if c.Log2Mongo {
		writer, err := logger.MongoWriter(repo.Collection())
		if err != nil {
			return nil, err
		}
		writers = append(writers, writer)
	}
	return logger.NewLogger(writers...), nil
}
