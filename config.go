/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/8/30
   Description :
-------------------------------------------------
*/

package zlog

import (
    "go.uber.org/zap/zapcore"
)

type Level string

const (
    DebugLevel  = "debug"  // 开发用, 生产模式下不应该是debug
    InfoLevel   = "info"   // 默认级别, 用于告知程序运行情况
    WarnLevel   = "warn"   // 比信息更重要, 但不需要单独的人工检查
    ErrorLevel  = "error"  // 高优先级的, 如果应用程序运行正常, 就不应该生成任何错误级别的日志
    DPanicLevel = "dpanic" // 严重的错误, 在开发者模式下日志记录器在写完消息后程序会感到恐慌
    PanicLevel  = "panic"  // 记录一条消息, 然后记录一条消息, 然后程序会感到恐慌
    FatalLevel  = "fatal"  // 记录一条消息, 然后调用 os.Exit(1)
)

var levelMapping = map[Level]zapcore.Level{
    DebugLevel:  zapcore.DebugLevel,
    InfoLevel:   zapcore.InfoLevel,
    WarnLevel:   zapcore.WarnLevel,
    ErrorLevel:  zapcore.ErrorLevel,
    DPanicLevel: zapcore.DPanicLevel,
    PanicLevel:  zapcore.PanicLevel,
    FatalLevel:  zapcore.FatalLevel,
}

type LogConfig struct {
    Level              string // 日志等级, debug, info, warn, error, dpanic, panic, fatal
    JsonEncoder        bool   // 启用json编码器, 输出的每一行日志转为json格式
    WriteToStream      bool   // 输出到屏幕
    WriteToFile        bool   // 日志是否输出到文件
    Name               string // 日志文件名, 末尾会自动附加 .log 后缀
    AppendPid          bool   // 是否在日志文件名后附加进程号
    Path               string // 默认日志存放路径
    FileMaxSize        int    // 每个日志最大尺寸,单位M
    FileMaxBackupsNum  int    // 日志文件最多保存多少个备份
    FileMaxDurableTime int    // 文件最多保存多长时间,单位天
    TimeFormat         string // 时间显示格式
    IsTerminal         bool   // 是否为控制台模式(控制台会打印彩色日志等级)
    ShowInitInfo       bool   // 显示初始化信息
    DevelopmentMode    bool   // 开发者模式, 在开发者模式下日志记录器在写完消息后程序会感到恐慌
    ShowFileAndLinenum bool   // 显示文件路径和行号
    CallerSkip         int    // 程序跳转次数
}

var DefaultConfig = LogConfig{
    Level:              "debug",
    JsonEncoder:        false,
    WriteToStream:      true,
    WriteToFile:        false,
    Name:               "zlog",
    AppendPid:          false,
    Path:               "./log",
    FileMaxSize:        32,
    FileMaxBackupsNum:  3,
    FileMaxDurableTime: 7,
    TimeFormat:         "2006-01-02 15:04:05",
    IsTerminal:         true,
    ShowInitInfo:       true,
    DevelopmentMode:    true,
    ShowFileAndLinenum: false,
    CallerSkip:         0,
}
