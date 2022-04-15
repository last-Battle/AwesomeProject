package abc

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type HTTPRetriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *HTTPRetriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	} else {
		return string(result)
	}
}

// func (r *HTTPRetriever) Post(aaa string) string {
// 	return aaa
// }
