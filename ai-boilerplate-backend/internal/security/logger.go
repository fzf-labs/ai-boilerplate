package security

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type redactingLogger struct {
	next log.Logger
}

func NewRedactingLogger(next log.Logger) log.Logger {
	if next == nil {
		return nil
	}
	return &redactingLogger{next: next}
}

func (l *redactingLogger) Log(level log.Level, keyvals ...any) error {
	filtered := make([]any, len(keyvals))
	for i, value := range keyvals {
		key := ""
		if i%2 == 1 {
			key = fmt.Sprint(keyvals[i-1])
		}
		filtered[i] = RedactValue(key, value)
	}
	return l.next.Log(level, filtered...)
}
