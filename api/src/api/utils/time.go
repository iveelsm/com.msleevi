package utils

import "time"

func NanoToMillis(elasped time.Duration) (result int64) {
	result = elasped.Nanoseconds() / 1000000
	return
}
