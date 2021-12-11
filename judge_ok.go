package jerr

import "fmt"

// 判断是否OK
func JudgeOk(ok, isRecover bool, msg ...interface{}) bool {
	if isRecover {
		defer func() {
			recover()
		}()
	}
	if !ok {
		err := GetLogInfo(2)
		fmt.Println(err, msg)
		panic(err)
	}
	return ok
}
