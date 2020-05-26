package utils

import (
	"time"
)

// RetryFunc is a function that should be retried
type RetryFunc func(attempt int) error

// Retry allows you to retry a function up to maxAttempts.
// If funcToRetry returns an error, this will sleep up to
// the delay duration and retry the function.
func RetryWithLinearDelay(maxAttempts int, delay time.Duration, funcToRetry RetryFunc) error {
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
