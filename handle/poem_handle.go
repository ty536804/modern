package handle

import (
	"fmt"
	"strings"
)

// @Param 分页
func GetUrls(url string,size int) []string {
	urls := make([]string, 0)
	urlTpl := strings.Replace(url,"A1.aspx","A%d.aspx",1)

	for i:=1; i <= size; i++  {
		urls = append(urls,fmt.Sprintf(urlTpl,i))
	}

	return urls
}