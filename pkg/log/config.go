package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func initLogrus(options *Options) (err error) {
	logger := logrus.StandardLogger()
	if options.LogGrpc {
		ReplaceGrpcLogger(logrus.NewEntry(logger))
	}

	var formatter logrus.Formatter
	fieldMap := logrus.FieldMap{
		logrus.FieldKeyTime:  "date",
		logrus.FieldKeyLevel: "level",
		logrus.FieldKeyMsg:   "message",
		logrus.FieldKeyFunc:  "caller",
	}

	var ws []io.Writer
	var hasStdout bool
	for _, output := range options.OutputPaths {
		if output == "stdout" {
			ws = append(ws, os.Stdout)
			hasStdout = true
		} else {
			iow := &lumberjack.Logger{
				Filename:   output,
				MaxSize:    options.RotationMaxSize,
				MaxBackups: options.RotationMaxBackups,
				MaxAge:     options.RotationMaxAge,
			}
			ws = append(ws, iow)
		}
	}
	multiWriter := io.MultiWriter(ws...)
	logger.SetOutput(multiWriter)

	if hasStdout {
		formatter = &logrus.TextFormatter{
			FieldMap: fieldMap,
		}
	} else {
		formatter = &logrus.JSONFormatter{
			FieldMap: fieldMap,
		}
	}

	switch options.LogLevel {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	}

	logger.SetFormatter(formatter)

	logger.AddHook(&contextHook{})

	return nil
}

func Configure(opts *Options) (err error) {
	defaultOpt := DefaultOptions()
	if opts.RotationMaxAge == 0 {
		opts.RotationMaxAge = defaultOpt.RotationMaxAge
	}
	if opts.RotationMaxBackups == 0 {
		opts.RotationMaxBackups = defaultOpt.RotationMaxBackups
	}

	if opts.RotationMaxSize == 0 {
		opts.RotationMaxSize = defaultOpt.RotationMaxSize
	}

	if opts.LogLevel == "" {
		opts.LogLevel = defaultOpt.LogLevel
	}

	return initLogrus(opts)
}
