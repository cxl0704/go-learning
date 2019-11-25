package common

import (
	"runtime"
)

//获取当前调用函数的名字
func PrintFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}