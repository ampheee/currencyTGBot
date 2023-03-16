package middleware

import "time"

func ConnectWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			attempts--
			time.Sleep(delay)
			continue
		}
		return nil
	}
	return
}
