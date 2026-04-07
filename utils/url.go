package utils

import (
	"github.com/valyala/fasthttp"
)

// // NOTE: GET
// url := "http://10.240.131.151:8080/ping"
// code, body, err := utils.HttpGet(url)
// if err != nil {
//   logger.Println(err.Error())
// } else {
//   logger.Println("Response Get:", code, body)
// }
//
// // NOTE: post
// url = "http://10.240.131.151:8080/post"
// jArr := []byte(`{"name":"fish"}`)
// code, body, err = utils.HttpPost(url, jArr)
// if err != nil {
//   logger.Println(err.Error())
// } else {
//   logger.Println("Response Post:", code, body)
// }

func HttpGet(url string) (int, string, error) {
	// NOTE: http get
	body := ""
	status, res, err := fasthttp.Get(nil, url)
	if err == nil {
		body = string(res)
	}

	return status, body, err
}

func HttpPost(url string, reqBody []byte) (int, string, error) {
	// NOTE: http post
	body := ""
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	req.SetBody(reqBody)

	var err error
	if err = fasthttp.Do(req, resp); err == nil {
		body = string(resp.Body())
	}

	return resp.StatusCode(), body, err
}
