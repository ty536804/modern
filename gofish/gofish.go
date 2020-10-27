package gofish

import "time"

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:81.0) Gecko/20100101 Firefox/81.0"
	Qps = 50
)

var rateLimiter = time.Tick(time.Second/Qps)

type GoFish struct {
	Request *Request
}

func NewGoFish() *GoFish  {
	return &GoFish{}
}

func (g *GoFish) Visit() error {
	return g.Request.Do()
}