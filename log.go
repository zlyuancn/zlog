/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2019/8/30
   Description :
-------------------------------------------------
*/

package zlog

import (
    "fmt"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
)

func NewLogger(opts ...Option) *zap.Logger {
    conf := newConfig(opts...)

    var encoderConfig = makeEncoderConfig()  // 编码器配置
    var ws = makeWriteSyncers(conf)          // 输出合成器
    var level = parserLogLevel(conf, "info") // 日志级别

    // zap核心
    core := zapcore.NewCore(
        zapcore.NewConsoleEncoder(encoderConfig),
        zapcore.NewMultiWriteSyncer(ws...),
        level,
    )

    // 构建zap选项
    zapopts := []zap.Option{}
    if conf.DevelopmentMode {
        zapopts = append(zapopts, zap.Development())
    }
    if conf.ShowFileAndLinenum {
        zapopts = append(zapopts, zap.AddCaller())
    }

    // 创建zap工具
    logger := zap.New(core, zapopts...) // 构造日志

    if conf.ShowInitInfo {
        logger.Info("zlog 初始化成功")
    }
    return logger
}

// 构建编码器配置
func makeEncoderConfig() zapcore.EncoderConfig {
    return zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "linenum",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
        EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
        EncodeDuration: zapcore.SecondsDurationEncoder, //
        EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
        EncodeName:     zapcore.FullNameEncoder,
    }
}

// 构建输出合成器
func makeWriteSyncers(conf *config) []zapcore.WriteSyncer {
    ws := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}
    if conf.WriteToFile {
        // 创建文件夹
        err := os.MkdirAll(conf.Path, 666)
        if err != nil {
            panic(err)
        }

        //构建lumberjack的hook
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
    return ws
}

// 解析日志等级
func parserLogLevel(conf *config, defaultLevel string) zapcore.Level {
    var l = new(zapcore.Level)
    if err := l.Set(string(conf.Level)); err != nil {
        _ = l.Set(defaultLevel)
    }
    return *l
}
