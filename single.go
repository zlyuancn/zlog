/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/4/17
   Description :
-------------------------------------------------
*/

package zlog

var defaultLog = func() *logWrap {
    conf := DefaultConfig
    conf.ShowInitInfo = false
    return New(conf)
}()

func Log(level Level, v ...interface{}) {
    defaultLog.print(level, "", v)
}
func Debug(v ...interface{}) {
    defaultLog.print(DebugLevel, "", v)
}
func Info(v ...interface{}) {
    defaultLog.print(InfoLevel, "", v)
}
func Warn(v ...interface{}) {
    defaultLog.print(WarnLevel, "", v)
}
func Error(v ...interface{}) {
    defaultLog.print(ErrorLevel, "", v)
}
func DPanic(v ...interface{}) {
    defaultLog.print(DPanicLevel, "", v)
}
func Panic(v ...interface{}) {
    defaultLog.print(PanicLevel, "", v)
}
func Fatal(v ...interface{}) {
    defaultLog.print(FatalLevel, "", v)
}

func Logf(level Level, format string, v ...interface{}) {
    defaultLog.print(level, format, v)
}
func Debugf(format string, v ...interface{}) {
    defaultLog.print(DebugLevel, format, v)
}
func Infof(format string, v ...interface{}) {
    defaultLog.print(InfoLevel, format, v)
}
func Warnf(format string, v ...interface{}) {
    defaultLog.print(WarnLevel, format, v)
}
func Errorf(format string, v ...interface{}) {
    defaultLog.print(ErrorLevel, format, v)
}
func DPanicf(format string, v ...interface{}) {
    defaultLog.print(DPanicLevel, format, v)
}
func Panicf(format string, v ...interface{}) {
    defaultLog.print(PanicLevel, format, v)
}
func Fatalf(format string, v ...interface{}) {
    defaultLog.print(FatalLevel, format, v)
}
