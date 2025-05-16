package utils

import "time"

func ExponentialBackoff(attempt int) time.Duration {
	return time.Second * time.Duration(1<<uint(attempt))
}
