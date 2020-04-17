/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/8/30
   Description :
-------------------------------------------------
*/

package zlog

import (
    "fmt"
    "os"
    "strings"
    "time"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
)

var DefaultLogger Loger = defaultLog

func New(conf LogConfig, opts ...zap.Option) Loger {
    var encoder = makeEncoder(&conf) // 编码器配置
    var ws = makeWriteSyncer(&conf)  // 输出合成器
    var level = makeLevel(&conf)     // 日志级别

    core := zapcore.NewCore(encoder, ws, level)
    opts = makeOpts(&conf, opts...)
    log := newLogWrap(zap.New(core, opts...))

    if conf.ShowInitInfo {
        log.Info("zlog 初始化成功")
    }
    return log
}

func makeEncoder(conf *LogConfig) zapcore.Encoder {
    cfg := zapcore.EncoderConfig{
        MessageKey:    "msg",
        LevelKey:      "level",
        TimeKey:       "time",
        NameKey:       "logger",
        CallerKey:     "linenum",
        StacktraceKey: "stacktrace",
        LineEnding:    zapcore.DefaultLineEnding,   // 末尾自动加上\n
        EncodeLevel:   zapcore.CapitalLevelEncoder, // 大写level
        EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
            enc.AppendString(t.Format(conf.TimeFormat))
        },
        EncodeDuration: zapcore.NanosDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
        EncodeName:     zapcore.FullNameEncoder,
    }
    if conf.IsTerminal {
        cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // 大写彩色level
    }
    if conf.MillisDuration {
        cfg.EncodeDuration = zapcore.MillisDurationEncoder
    }
    if conf.JsonEncoder {
        return zapcore.NewJSONEncoder(cfg)
    }
    return zapcore.NewConsoleEncoder(cfg)
}

func makeWriteSyncer(conf *LogConfig) zapcore.WriteSyncer {
    var ws []zapcore.WriteSyncer
    if conf.WriteToStream {
        ws = append(ws, zapcore.AddSync(os.Stdout))
    }

    if conf.WriteToFile {
        // 创建文件夹
        err := os.MkdirAll(conf.Path, 666)
        if err != nil {
            fmt.Printf("无法创建日志目录: <%s>: %s\n", conf.Path, err)
            os.Exit(1)
        }

        // 构建lumberjack的hook
        name := conf.Name
        if conf.AppendPid {
            name = fmt.Sprintf("%s_%d", name, os.Getpid())
        }
        lumberjackHook := &lumberjack.Logger{
            Filename:   fmt.Sprintf("%s/%s.log", conf.Path, name), // 日志文件路径
            MaxSize:    conf.FileMaxSize,                          // 每个日志文件保存的最大尺寸 单位：M
            MaxBackups: conf.FileMaxBackupsNum,                    // 日志文件最多保存多少个备份
            MaxAge:     conf.FileMaxDurableTime,                   // 文件最多保存多少天
            Compress:   false,                                     // 是否压缩
        }
        ws = append(ws, zapcore.AddSync(lumberjackHook))
    }
    return zapcore.NewMultiWriteSyncer(ws...)
}

func makeLevel(conf *LogConfig) zapcore.Level {
    level := Level(strings.ToLower(conf.Level))
    return parserLogLevel(level)
}

func makeOpts(conf *LogConfig, opts ...zap.Option) []zap.Option {
    const callerSkipOffset = 2

    opts = append([]zap.Option{}, opts...)
    if conf.DevelopmentMode {
        opts = append(opts, zap.Development())
    }
    if conf.ShowFileAndLinenum {
        opts = append(opts, zap.AddCaller())
    }

    opts = append(opts, zap.AddCallerSkip(conf.CallerSkip+callerSkipOffset))
    return opts
}

func parserLogLevel(level Level) zapcore.Level {
    l, ok := levelMapping[level]
    if ok {
        return l
    }

    return zapcore.InfoLevel
}
