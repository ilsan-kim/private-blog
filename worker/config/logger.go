package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"log"
	"os"
	"runtime"
	"strings"
)

var Json zerolog.Logger

func init() {
	Json = zerolog.New(os.Stdout).With().Timestamp().Logger().Hook(shortCallerHook())
}

func shortCallerHook() zerolog.HookFunc {
	return func(e *zerolog.Event, level zerolog.Level, message string) {
		_, file, line, ok := runtime.Caller(7)
		if !ok {
			return
		}
		ss := strings.Split(file, "/")
		if len(ss) > 2 {
			file = strings.Join(ss[len(ss)-2:], "/")
		}
		e.Str("caller", fmt.Sprintf("%s:%d", file, line))
	}
}

// 표준 log 패키지를 json 형식으로 출력하도록 설정한다.
func UseJsonLogger() {
	log.SetFlags(0)
	log.SetOutput(Json)
}
