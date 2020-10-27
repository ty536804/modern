package gofish

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Url string
	Method string
	Headers *http.Header
	Body io.Reader
	Handle Handel
	Client http.Client
}

func (r *Request) Do() error {
	request, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return  err
	}

	request.Header = *r.Headers

	resp, err := r.Client.Do(request)
	if err !=nil {
		return  err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("errorstatus code:%d",resp.StatusCode)
	}

	r.Handle.Worker(resp.Body, r.Url)
	defer resp.Body.Close()

	return nil
}

// 发送网络请求
//func (r *Request) NewRequest(method, Url, userAgent string,handle Handel, body io.Reader) (*Request, error) {
func NewRequest(method, Url, userAgent string,handle Handel, body io.Reader) (*Request, error) {
	//验证url是否合法
	_,err := url.Parse(Url)
	if err != nil {
		return nil,err
	}

	// 添加头信息
	hdr := http.Header{}
	if userAgent != "" {
		hdr.Add("User-Agent", userAgent)
	} else {
		hdr.Add("User-Agent", UserAgent)
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	return &Request{
		Url:Url,
		Method:method,
		Headers:&hdr,
		Handle:handle,
		Body:body,
		Client:client,
	},nil
}