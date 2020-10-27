package main

import (
	"fmt"
	"gushiwen/gofish"
	"gushiwen/handle"
)

func main()  {
	url := "https://so.gushiwen.cn/authors/"

	h := handle.AuthorHandle{}

	fish := gofish.NewGoFish()
	request, err := gofish.NewRequest("GET",url,gofish.UserAgent,&h,nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	fish.Request = request
	fish.Visit()
}