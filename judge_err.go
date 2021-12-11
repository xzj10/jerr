package jerr

import (
	"fmt"
	"runtime"
	"strings"
)

// 获取调用此方法 所在的文件、方法名、行号
func GetLogInfo(skip int) string {
	fn, file, line, ok := runtime.Caller(skip)
	if ok {
		fary := strings.Split(file, "/")
		file = fary[len(fary)-1]
		fnName := runtime.FuncForPC(fn).Name()
		str := fmt.Sprintf("错误发生在: [%v,%v(),%v]", file, fnName, line)
		return str
	}
	return "[]"
}

// 判断错误.
/*
	err: 要判断的错误
	isRecover: 是否需要捕获改错误, bool类型
	msg: 对错误打印的辅助描述信息
*/
func JudgeErr(err error, isRecover bool, msg ...interface{}) error {
	if isRecover {
		defer func() {
			recover()
		}()
	}
	if err != nil {
		errs := fmt.Sprintf("[ %v ]", err)
		fmt.Println(GetLogInfo(2), msg, errs)
		panic(err)
	}
	return err
}
