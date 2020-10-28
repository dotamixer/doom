package resync

import "time"

type Options struct {
	retry    int
	timeOut  time.Duration //second
	timeLock time.Duration
}

var defaultOpts = defaultOptions()

func defaultOptions() *Options {
	return &Options{
		retry:    5,
		timeOut:  time.Second * 5,
		timeLock: time.Second * 5,
	}
}
