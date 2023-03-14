package logger

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// once give opportunity to store logger from first call with multiply calls from another funcs
var once sync.Once
var log zerolog.Logger

func GetLogger() zerolog.Logger {
	once.Do(func() {
		c := color.New(color.FgRed)
		log = zerolog.New(zerolog.
			ConsoleWriter{
			TimeFormat: time.DateTime,
			Out:        os.Stderr,
			FormatCaller: func(i interface{}) string {
				return "|" + filepath.Base(fmt.Sprintf("%s|", i))
			},
			FormatErrFieldName: func(i interface{}) string {
				return c.Sprint(strings.ToUpper(fmt.Sprintf("[%s] -> ", i)))
			},
		}).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
	})
	return log
}
