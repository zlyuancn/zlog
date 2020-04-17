/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/4/17
   Description :
-------------------------------------------------
*/

package zlog

import (
    "go.uber.org/zap"
)

type Loger interface {
    Log(level Level, msg string, fields ...zap.Field)
    Debug(msg string, fields ...zap.Field)
    Info(msg string, fields ...zap.Field)
    Warn(msg string, fields ...zap.Field)
    Error(msg string, fields ...zap.Field)
    DPanic(msg string, fields ...zap.Field)
    Panic(msg string, fields ...zap.Field)
    Fatal(msg string, fields ...zap.Field)
}

type logWrap struct {
    log *zap.Logger
}

var _ Loger = (*logWrap)(nil)

func newLogWrap(log *zap.Logger) *logWrap {
    l := &logWrap{
        log: log,
    }
    return l
}

func (m *logWrap) print(level Level, msg string, fields ...zap.Field) {
    if ce := m.log.Check(parserLogLevel(level), msg); ce != nil {
        ce.Write(fields...)
    }
}
func (m *logWrap) Log(level Level, msg string, fields ...zap.Field) {
    m.print(level, msg, fields...)
}
func (m *logWrap) Debug(msg string, fields ...zap.Field) {
    m.print(DebugLevel, msg, fields...)
}
func (m *logWrap) Info(msg string, fields ...zap.Field) {
    m.print(InfoLevel, msg, fields...)
}
func (m *logWrap) Warn(msg string, fields ...zap.Field) {
    m.print(WarnLevel, msg, fields...)
}
func (m *logWrap) Error(msg string, fields ...zap.Field) {
    m.print(ErrorLevel, msg, fields...)
}
func (m *logWrap) DPanic(msg string, fields ...zap.Field) {
    m.print(DPanicLevel, msg, fields...)
}
func (m *logWrap) Panic(msg string, fields ...zap.Field) {
    m.print(PanicLevel, msg, fields...)
}
func (m *logWrap) Fatal(msg string, fields ...zap.Field) {
    m.print(FatalLevel, msg, fields...)
}
