package decorator

import (
	"fmt"
	"strconv"
	"time"
)

type ExponentialBackOff struct {
	Fetcher
	maxRetries     int
	currentTimeout int
	maxTimeout     int
}

func (ebo ExponentialBackOff) Fetch(args Args) (Response, error) {
	res, err := ebo.Fetcher.Fetch(args)
	if err != nil {
		if ebo.currentTimeout < ebo.maxTimeout && ebo.maxRetries > 0 {
			ebo.currentTimeout = ebo.currentTimeout * 2
			time.Sleep(time.Duration(ebo.currentTimeout) * time.Second)
			ebo.maxRetries--
			fmt.Println("failed, retry count is " + strconv.Itoa(ebo.maxRetries) + " and sleep time is " + strconv.Itoa(ebo.currentTimeout))
			return ebo.Fetch(args)
		} else {
			fmt.Println("failed, max retry count reached")
			return nil, err
		}
	}
	fmt.Println("request successful")
	return res, nil
}

func NewEBO(f Fetcher, maxRetries int, maxTimeout int) ExponentialBackOff {
	return ExponentialBackOff{
		Fetcher:        f,
		maxRetries:     maxRetries,
		maxTimeout:     maxTimeout,
		currentTimeout: 1,
	}
}
