package util

import (
	"fmt"

	"github.com/imroc/log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Debugf("%s: %s", err, msg)
		panic(fmt.Sprintf("%s: %s", err, msg))
	}
}

func CheckErrors(err error, msg string) {
	if err != nil {
		log.Debugf("%s: %s", err, msg)
	}
}
