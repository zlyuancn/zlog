/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/4/17
   Description :
-------------------------------------------------
*/

package zlog

import (
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type extraLogger struct {
	zapcore.Core
	logs []*logWrap
}

// 添加额外的输出
func WithExtraLogger(logs ...*logWrap) zap.Option {
	return zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return &extraLogger{
			Core: core,
			logs: append(([]*logWrap)(nil), logs...),
		}
	})
}
func (c *extraLogger) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}
func (h *extraLogger) With(fields []zapcore.Field) zapcore.Core {
	return &extraLogger{
		Core: h.Core.With(fields),
		logs: h.logs,
	}
}
func (c *extraLogger) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	err := c.Core.Write(ent, fields)
	for i := range c.logs {
		err = multierr.Append(err, c.logs[i].Core().Write(ent, fields))
	}
	return err
}
