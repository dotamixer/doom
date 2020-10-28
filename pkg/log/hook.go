package log

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type contextHook struct{}

func (c contextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (c contextHook) Fire(entry *logrus.Entry) error {
	if entry.Context != nil {
		var traceID, requestID []string
		meta, ok := metadata.FromIncomingContext(entry.Context)
		if ok {
			traceID = meta.Get("x-b3-traceid")
			requestID = meta.Get("x-request-id")
		}

		if len(traceID) > 0 {
			entry.Data["trace-id"] = traceID[0]
		}

		if len(requestID) > 0 {
			entry.Data["request-id"] = requestID[0]
		}
	}

	return nil
}

var _ logrus.Hook = contextHook{}
