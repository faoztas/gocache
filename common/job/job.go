package job

import (
	"fmt"
	"time"

	"github.com/faoztas/gocache/common/log"
	"github.com/faoztas/gocache/utils/errors"
)

func Run(fn func(), delay time.Duration) {
	go func() {
		defer retry(fn, delay)
		fn()
	}()
}

func retry(fn func(), delay time.Duration) {
	if r := recover(); r != nil {
		var err error
		switch x := r.(type) {
		case string:
			err = fmt.Errorf(x)
		case error:
			err = x
		default:
			err = fmt.Errorf(errors.UnknownPanic)
		}
		log.Debug(fmt.Errorf(errors.RoutineError, err))
	}
	log.Info(fmt.Errorf(errors.TaskMsg, delay.Seconds()))
	time.Sleep(delay)
	Run(fn, delay)
}
