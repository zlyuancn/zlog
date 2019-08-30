/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/8/30
   Description :
-------------------------------------------------
*/

package zlog

import "strings"

type Level string

const (
    DebugLevel  = "debug"  //开发用, 生产模式下不应该是debug
    InfoLevel   = "info"   //默认级别, 用于告知程序运行情况
    WarnLevel   = "warn"   //比信息更重要，但不需要单独的人工检查。
    ErrorLevel  = "error"  //高优先级的。如果应用程序运行正常，就不应该生成任何错误级别的日志。
    DPanicLevel = "dpanic" //特别重要的错误，在开发者模式下日志记录器在写完消息后程序会感到恐慌。
    Panic       = "panic"  //记录一条消息, 然后记录一条消息, 然后程序会感到恐慌。
    Fatal       = "fatal"  //记录一条消息, 然后调用 os.Exit(1)
)

type config struct {
    Level              Level  // 日志等级, debug, info, warn, error, dpanic, panic, fatal
    WriteToFile        bool   // 日志是否输出到文件
    Name               string // 日志文件名, 末尾会自动附加 .log 后缀
    AppendPid          bool   // 是否在日志文件名后附加进程号
    Path               string // 默认日志存放路径
    FileMaxSize        int    // 每个日志最大尺寸,单位M
    FileMaxBackupsNum  int    // 日志文件最多保存多少个备份
    FileMaxDurableTime int    // 文件最多保存多长时间,单位天
    ShowInitInfo       bool   // 显示初始化信息
    DevelopmentMode    bool   // 开发者模式
    ShowFileAndLinenum bool   // 显示文件路径和行号
}

type Option func(conf *config)

func newConfig(opts ...Option) *config {
    conf := &config{
        Level:              "info",
        WriteToFile:        false,
        Name:               "zlog",
        AppendPid:          false,
        Path:               "./log",
        FileMaxSize:        128,
        FileMaxBackupsNum:  3,
        FileMaxDurableTime: 7,
        ShowInitInfo:       true,
        DevelopmentMode:    true,
        ShowFileAndLinenum: false,
    }
    for _, o := range opts {
        o(conf)
    }
    return conf
}

// 设置日志级别
func WithLevel(level Level) Option {
    return func(conf *config) {
        conf.Level = level
    }
}

// 设置是否输出到文件
func WithWriteToFile(write bool) Option {
    return func(conf *config) {
        conf.WriteToFile = write
    }
}

// 设置日志文件名(不包含路径, 末尾会自动附加 .log 后缀)
func WithFileName(name string) Option {
    return func(conf *config) {
        conf.Name = name
    }
}

// 设置是否在日志文件名后附加进程pid号
func WithAppendPid(append bool) Option {
    return func(conf *config) {
        conf.AppendPid = append
    }
}

// 设置日志文件路径
func WithFilePath(path string) Option {
    return func(conf *config) {
        path = strings.TrimRight(path, "/")
        path = strings.TrimRight(path, "\\")
        conf.Path = path
    }
}

// 设置日志文件每个日志最大尺寸, 单位M
func WithFileMaxSize(size int) Option {
    return func(conf *config) {
        if size > 0 {
            conf.FileMaxSize = size
        }
    }
}

// 设置日志文件最多保存多少个备份
func WithFileMaxBackupsNum(num int) Option {
    return func(conf *config) {
        conf.FileMaxBackupsNum = num
    }
}

// 设置日志文件最多保存多长时间(天)
func WithFileMaxDurableTime(day int) Option {
    return func(conf *config) {
        conf.FileMaxDurableTime = day
    }
}

// 设置是否输出日志初始化完成信息
func WithShowInitInfo(show bool) Option {
    return func(conf *config) {
        conf.ShowInitInfo = show
    }
}

// 设置是否使用开发者模式, 在开发者模式下 DPanic 会打印一条消息后让程序感到恐慌, 非开发者模式会打印一条消息后继续执行
func WithDevelopmentMode(on bool) Option {
    return func(conf *config) {
        conf.DevelopmentMode = on
    }
}

// 设置是否输出文件路径以及行号
func WithShowFileAndLinenum(show bool) Option {
    return func(conf *config) {
        conf.ShowFileAndLinenum = show
    }
}
