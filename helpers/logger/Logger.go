package logger

import (
	"fmt"
	"runtime"
	"strings"
)

func Log(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	paths := strings.Split(file, "/")
	f := paths[len(paths)-1]
	s := fmt.Sprintf(format, a...)
	fmt.Printf("gojenkins-core: %s:%d: %s\n", f, line, s)
}
