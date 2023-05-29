package article

import (
	"errors"
	"fmt"
	"github.com/sunmi-OS/gocore/v2/api/ecode"
)

const (
	ErrProductNotFound_Num = 10001
)

var (
	ErrMap = map[int]string{
		ErrProductNotFound_Num: "product not found",
	}
)

func ErrProductNotFound(msg ...string) *ecode.ErrorV2 {
	msgStr := ErrMap[ErrProductNotFound_Num]
	if len(msg) > 0 {
		msgStr = msg[0]
	}
	return ecode.NewV2(ErrProductNotFound_Num, msgStr)
}

func IsErrProductNotFound(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Status.Code == ErrProductNotFound_Num
	}
	return false
}

func IsErrProductNotFoundSame(err error) bool { //  这个方法待优化
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Status.Code == ErrProductNotFound_Num && se.Status.Reason == ErrMap[ErrProductNotFound_Num]
	}
	return false
}

func fun1() error {
	return ErrProductNotFound()
}

func fun2() error {
	err := ErrProductNotFound("hahaha")

	if IsErrProductNotFound(err) {
		fmt.Println("same error code")
	}

	if IsErrProductNotFoundSame(err) {
		fmt.Println("same error code and reason")
	}

	err = ErrProductNotFound()
	if IsErrProductNotFound(err) {
		fmt.Println("same error")
	}

}
