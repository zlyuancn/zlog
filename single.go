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

var defaultLog = func() *logWrap {
    conf := DefaultConfig
    return New(conf).(*logWrap)
}()

func Log(level Level, msg string, fields ...zap.Field) {
    defaultLog.print(level, msg, fields...)
}
func Debug(msg string, fields ...zap.Field) {
    defaultLog.print(DebugLevel, msg, fields...)
}
func Info(msg string, fields ...zap.Field) {
    defaultLog.print(InfoLevel, msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {
    defaultLog.print(WarnLevel, msg, fields...)
}
func Error(msg string, fields ...zap.Field) {
    defaultLog.print(ErrorLevel, msg, fields...)
}
func DPanic(msg string, fields ...zap.Field) {
    defaultLog.print(DPanicLevel, msg, fields...)
}
func Panic(msg string, fields ...zap.Field) {
    defaultLog.print(PanicLevel, msg, fields...)
}
func Fatal(msg string, fields ...zap.Field) {
    defaultLog.print(FatalLevel, msg, fields...)
}
