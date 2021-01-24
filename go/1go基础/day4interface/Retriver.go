package day4interface

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriver struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r Retriver) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(response, true)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
