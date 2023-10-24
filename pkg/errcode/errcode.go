package errcode

import (
	"net/http"
	"fmt"
)

type Error struct {
	code int
	message string 
	// code int `json:"code"`
	// message string `json:"message"`
	// details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, message string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在，请更换一个", code))
	}
	codes[code] = message
	return &Error{code: code, message: message}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}