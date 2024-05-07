package connection

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func GetHTTPDataFast(URL string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	req.SetRequestURI(URL)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("unexpected status code ::: %d from URL: %s", resp.StatusCode(), URL)
	}

	body := resp.Body()

	return body, nil
}
