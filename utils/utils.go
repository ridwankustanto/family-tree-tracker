package utils

import (
	"strings"
	"time"
)

func ForeverSleep(d time.Duration, f func(int) error) {
	for i := 0; ; i++ {
		err := f(i)
		if err == nil {
			return
		}
		time.Sleep(d)
	}
}

func FormatUUID(s string) string {
	return strings.ReplaceAll(s, "-", "")
}
