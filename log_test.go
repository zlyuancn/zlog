/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/4/17
   Description :
-------------------------------------------------
*/

package zlog

import (
    "testing"
)

func TestNew(t *testing.T) {
    conf := DefaultConfig
    l := New(conf)
    l.Debug("123")
    l.Info("123")
}
