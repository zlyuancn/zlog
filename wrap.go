/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/4/17
   Description :
-------------------------------------------------
*/

package zlog

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logWrap struct {
	log            *zap.Logger
	fields         []zap.Field
	callerMinLevel zapcore.Level
}

var _ Logfer = (*logWrap)(nil)

func newLogWrap(log *zap.Logger, callerMinLevel zapcore.Level, fields ...zap.Field) *logWrap {
	l := &logWrap{
		log:            log,
		fields:         append([]zap.Field{}, fields...),
		callerMinLevel: callerMinLevel,
	}
	return l
}

func (m *logWrap) Core() zapcore.Core {
	return m.log.Core()
}

func (m *logWrap) print(level Level, format string, v []interface{}) {
	msg, fields := m.makeBody(format, v)
	zapLevel := parserLogLevel(level)
	if ce := m.log.Check(zapLevel, msg); ce != nil {
		if zapLevel < m.callerMinLevel {
			ce.Caller.Defined = false
		}
		ce.Write(fields...)
	}
}
func (m *logWrap) Log(level Level, v ...interface{}) {
	m.print(level, "", v)
}
func (m *logWrap) Debug(v ...interface{}) {
	m.print(DebugLevel, "", v)
}
func (m *logWrap) Info(v ...interface{}) {
	m.print(InfoLevel, "", v)
}
func (m *logWrap) Warn(v ...interface{}) {
	m.print(WarnLevel, "", v)
}
func (m *logWrap) Error(v ...interface{}) {
	m.print(ErrorLevel, "", v)
}
func (m *logWrap) DPanic(v ...interface{}) {
	m.print(DPanicLevel, "", v)
}
func (m *logWrap) Panic(v ...interface{}) {
	m.print(PanicLevel, "", v)
}
func (m *logWrap) Fatal(v ...interface{}) {
	m.print(FatalLevel, "", v)
}

func (m *logWrap) Logf(level Level, format string, v ...interface{}) {
	m.print(level, format, v)
}
func (m *logWrap) Debugf(format string, v ...interface{}) {
	m.print(DebugLevel, format, v)
}
func (m *logWrap) Infof(format string, v ...interface{}) {
	m.print(InfoLevel, format, v)
}
func (m *logWrap) Warnf(format string, v ...interface{}) {
	m.print(WarnLevel, format, v)
}
func (m *logWrap) Errorf(format string, v ...interface{}) {
	m.print(ErrorLevel, format, v)
}
func (m *logWrap) DPanicf(format string, v ...interface{}) {
	m.print(DPanicLevel, format, v)
}
func (m *logWrap) Panicf(format string, v ...interface{}) {
	m.print(PanicLevel, format, v)
}
func (m *logWrap) Fatalf(format string, v ...interface{}) {
	m.print(FatalLevel, format, v)
}

func (m *logWrap) makeBody(format string, v []interface{}) (string, []zap.Field) {
	args := make([]interface{}, 0, len(v))
	fields := append([]zap.Field{}, m.fields...)
	for _, value := range v {
		switch val := value.(type) {
		case zap.Field:
			fields = append(fields, val)
		case *zap.Field:
			fields = append(fields, *val)
		default:
			args = append(args, value)
		}
	}
	if format != "" {
		return fmt.Sprintf(format, args...), fields
	}
	return fmt.Sprint(args...), fields
}

// 包装添加一些ZapField, 这会创建一个Logfer副本
func WrapZapFields(l Logfer, fields ...zap.Field) (Logfer, bool) {
	if a, ok := l.(*logWrap); ok {
		fields = append(append([]zap.Field{}, a.fields...), fields...)
		return newLogWrap(a.log, a.callerMinLevel, fields...), true
	}
	return nil, false
}

// 包装添加一些ZapField, 这会创建一个Loger副本
func WrapZapFieldsWithLoger(l Loger, fields ...zap.Field) (Loger, bool) {
	if a, ok := l.(*logWrap); ok {
		fields = append(append([]zap.Field{}, a.fields...), fields...)
		return newLogWrap(a.log, a.callerMinLevel, fields...), true
	}
	return nil, false
}
