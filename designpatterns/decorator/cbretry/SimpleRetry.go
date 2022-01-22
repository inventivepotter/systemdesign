package decorator

import (
	"errors"
	"fmt"
	"strconv"
)

// SimpleRetryRequest is used to implement simple retry upto given number
type SimpleRetryRequest struct {
	Fetcher
	maxRetries int
}

func (srr SimpleRetryRequest) Fetch(args Args) (Response, error) {
	res, err := srr.Fetcher.Fetch(args)
	if err != nil {
		if srr.maxRetries > 0 {
			srr.maxRetries--
			fmt.Println("failed, retry count is " + strconv.Itoa(srr.maxRetries))
			return srr.Fetch(args)
		}
		fmt.Println("failed, max retry count reached")
		return nil, errors.New("failed, max retry count reached")
	}
	fmt.Println("request successful")
	return res, nil
}

func NewSRR(f Fetcher, maxRetries int) SimpleRetryRequest {
	return SimpleRetryRequest{
		Fetcher:    f,
		maxRetries: maxRetries,
	}
}
