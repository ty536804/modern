package gofish

import "io"

//处理抓取页面，解析页面的功能

// 定义接口
type Handel interface {
	Worker(body io.Reader,url string)//
}


