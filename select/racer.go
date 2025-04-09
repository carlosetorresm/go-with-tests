package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a, nil
	}
	return b, nil
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
