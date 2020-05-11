package common

import (
	"fmt"
	"runtime"
	"strings"
)

type ErrDesc struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func trimStringFromDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}

func (e *ErrDesc) FormatErrMsg(fun string, err error) *ErrDesc {
	pc, _, line, ok := runtime.Caller(1)
	if ok {
		fun := trimStringFromDot(runtime.FuncForPC(pc).Name())
		if err != nil {
			e.Msg = fmt.Sprintf("fun %s %d,%s,%s", fun, line, e.Msg, err)
		} else {
			e.Msg = fmt.Sprintf("fun %s %d ,%s", fun, line, e.Msg)
		}
	} else {
		if err != nil {
			e.Msg = fmt.Sprintf("fun %s,%s,%s", fun, e.Msg, err)
		} else {
			e.Msg = fmt.Sprintf("fun %s,%s", fun, e.Msg)
		}
	}
	return e
}

func (e *ErrDesc) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.Code, e.Msg)
}

func (e *ErrDesc) ToError() error {
	return fmt.Errorf("code:%d,msg:%s", e.Code, e.Msg)
}

var (
	ErrFoundMod      = &ErrDesc{Code: -1, Msg: "not found chain module"}
	ErrFoundFun      = &ErrDesc{Code: -2, Msg: "not found chain fun"}
	ErrUnknownDef    = &ErrDesc{Code: -3, Msg: "unknown define"}
	ErrPutState      = &ErrDesc{Code: -4, Msg: "failed to put state"}
	ErrGetState      = &ErrDesc{Code: -4, Msg: "failed to get state"}
	ErrSetEvent      = &ErrDesc{Code: -4, Msg: "failed set event"}
	ErrArgsNum       = &ErrDesc{Code: -5, Msg: "fun input args num error"}
	ErrJsonMarshal   = &ErrDesc{Code: -6, Msg: "failed bytes stream to marshal"}
	ErrJsonUnmarshal = &ErrDesc{Code: -6, Msg: "failed bytes stream to unmarshal"}
	ErrGetStateNil   = &ErrDesc{Code: -7, Msg: "failed get state value is null"}
	ErrStateKeyExist = &ErrDesc{Code: -8, Msg: "failed state key is exist,expect key is null"}
	ErrNewCompKey    = &ErrDesc{Code: -9, Msg: "failed create composite key"}
)
