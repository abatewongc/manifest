package util

import (
	"log"
	"runtime"
)

func HandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, fn, line, _ := runtime.Caller(1)

		log.Fatalf("[error] in %s [%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	}
	return
}