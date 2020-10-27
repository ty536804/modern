package handle

import (
	"fmt"
	"strings"
)


func GetUrls(url string,size int) []string {
	urls := make([]string, 0)
	urlTpl := strings.Replace(url,"A1.aspx","A%d.aspx",1)

	for i:=1; i <= size; i++  {
		urls = append(urls,fmt.Sprintf(urlTpl,i))
		fmt.Println(fmt.Sprintf(urlTpl,i))
	}

	return urls
}