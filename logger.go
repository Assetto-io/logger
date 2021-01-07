package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"strconv"
	"strings"
)

const (
	DebugLvl = "debug"
	InfoLvl  = "info"
	ErrorLvl = "error"
)

type Logger struct {
	log *zap.Logger
}

type LogField struct {
	Key   string
	Value interface{}
}

func New(name string, lvl string) *Logger {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(setLevel(lvl)),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			NameKey:      "id",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	return &Logger{log: log.Named(name)}
}

func (l *Logger) Debug(msg string, tags ...LogField) {
	_, fn, ln, ok := runtime.Caller(1)
	if ok {
		tags = append(tags, Field("line", formatCaller(fn, ln)))
	}
	l.log.Debug(msg, toZapField(tags)...)
	l.log.Sync()
}

func (l *Logger) Info(msg string, tags ...LogField) {
	l.log.Info(msg, toZapField(tags)...)
	l.log.Sync()
}

func (l *Logger) Error(msg string, err error, tags ...LogField) {
	zapTags := toZapField(tags)
	zapTags = append(zapTags, zap.NamedError("error", err))

	_, fn, ln, ok := runtime.Caller(1)
	if ok {
		zapTags = append(zapTags, zap.Any("line", formatCaller(fn, ln)))
	}

	l.log.Error(msg, zapTags...)
	l.log.Sync()
}

func Field(key string, value interface{}) LogField {
	return LogField{Key: key, Value: value}
}

// ======================== Helper Functions ========================

func formatCaller(fn string, ln int) string {
	idx := strings.LastIndexByte(fn, '/')
	if idx == -1 {
		return fn
	}
	// Find the penultimate separator.
	idx = strings.LastIndexByte(fn[:idx], '/')
	if idx == -1 {
		return fn
	}
	var sb strings.Builder
	// Keep everything after the penultimate separator.
	sb.WriteString(fn[idx+1:])
	sb.WriteString(": ")
	sb.WriteString(strconv.Itoa(ln))

	return sb.String()
}

func toZapField(field []LogField) []zap.Field {
	l := len(field)
	if l == 0 {
		return nil
	}
	r := make([]zap.Field, l, l)
	for i, f := range field {
		r[i] = zap.Any(f.Key, f.Value)
	}
	return r
}

func setLevel(lvl string) zapcore.Level {
	if len(lvl) == 0 {
		return zap.InfoLevel
	}

	switch lvl {
	case DebugLvl:
		return zap.DebugLevel
	case InfoLvl:
		return zap.InfoLevel
	case ErrorLvl:
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
