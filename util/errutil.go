package util

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
)

// ErrorInfo 错误通用格式
type ErrorInfo struct {
	ErrMsg string
	ErrPos string
}

func (e *ErrorInfo) Error() string {
	return e.String()
}

func (e *ErrorInfo) String() string {
	if e == nil {
		return "[errno:-1||errmsg:nil]"
	}
	return fmt.Sprintf("[errmsg:%s||errpos:%s]", e.ErrMsg, e.ErrPos)
}

// Errorf 创建新的错误 格式化错误信息
func Errorf(format string, a ...interface{}) *ErrorInfo {
	var errpos string
	_, file, line, ok := runtime.Caller(2)
	if ok {
		errpos = path.Base(file) + ":" + strconv.Itoa(line)
	}
	return &ErrorInfo{fmt.Sprintf(format, a...), errpos}
}
