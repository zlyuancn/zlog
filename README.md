
# zlog
> 日志模块

## 获得zlog
` go get -u github.com/zlyuancn/zlog `

## 导入zlog
```go
import "github.com/zlyuancn/zlog"
```

## 实例

```go
log := zlog.NewLogger()
log.Info("hello")
```

## 将日志输出到文件

```go
log := zlog.NewLogger(
    zlog.WithWriteToFile(true),
)
log.Info("hello")
```

## 其他选项

```
WithLevel(level Level) // 设置日志级别, 默认 info, 可选: debug, info, warn, error, dpanic, panic, fatal
WithWriteToFile(write bool) // 设置是否输出到文件, 默认 false
WithFileName(name string) // 设置日志文件名, 默认 zlog (不包含路径, 末尾会自动附加 .log 后缀)
WithAppendPid(append bool) // 设置是否在日志文件名后附加进程pid号, 默认 false
WithFilePath(path string) // 设置日志文件路径, 默认 ./log
WithFileMaxSize(size int) // 设置日志文件每个日志最大尺寸, 单位M, 默认 128
WithFileMaxBackupsNum(num int) // 设置日志文件最多保存多少个备份, 默认 3
WithFileMaxDurableTime(day int) // 设置日志文件最多保存多长时间(天), 默认 7
WithShowInitInfo(show bool) // 设置是否输出日志初始化完成信息, 默认 true
WithDevelopmentMode(on bool) // 设置是否使用开发者模式, 在开发者模式下 DPanic 会打印一条消息后让程序感到恐慌
WithShowFileAndLinenum(show bool) // 设置是否输出文件路径以及行号
```
