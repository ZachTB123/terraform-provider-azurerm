package utils

import (
	"time"
)

// RetryFunc is a function that should be retried
type RetryFunc func(attempt int) error

// Do executes the function using config
func Retry(maxAttempts int, delay time.Duration, funcToRetry RetryFunc, ) error {
	var err error

	for i := 0; i < maxAttempts; i++ {
		err = funcToRetry(i)

		if err == nil {
			break
		} else {
			time.Sleep(delay)
		}
	}

	return err
}