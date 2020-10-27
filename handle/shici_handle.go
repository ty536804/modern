package handle

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gushiwen/db"
	"gushiwen/gofish"
	"io"
	"strings"
)

type AuthorHandle struct {
}

const Base = "https://so.gushiwen.cn"

func (a *AuthorHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Errorf("doc err:%s", err)
	}

	doc.Find(".sons").Find(".cont").Find("a").Each(func(i int, s *goquery.Selection) {
		author := s.Text()
		link,_ := s.Attr("href")
		fmt.Printf("author:%s, link:%s \n",author,link)


		h := PomeHomeHandle{}

		fish := gofish.NewGoFish()
		request, err := gofish.NewRequest("GET", Base+link, gofish.UserAgent, &h,nil)

		if err != nil {
			fmt.Println(err)
			return
		}

		fish.Request = request
		fish.Visit()
	})
}

type PomeHomeHandle struct {

}
func (p *PomeHomeHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Errorf("doc err:%s",err)
	}

	doc.Find(".sonspic").Find(".cont").Find("p").Find("a").Each(func(i int, s *goquery.Selection) {
		link,_ := s.Attr("href")
		fmt.Printf("作品主页=%s \n",Base+link)

		h := PomeHomeInfoHandle{}

		fish := gofish.NewGoFish()
		request, err := gofish.NewRequest("GET", Base+link, gofish.UserAgent, &h,nil)

		if err != nil {
			fmt.Println(err)
			return
		}

		fish.Request = request
		fish.Visit()
	})
}

type PomeHomeInfoHandle struct {

}
func (p *PomeHomeInfoHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Errorf("doc err:%s",err)
	}

	doc.Find(".cont").Each(func(i int, s *goquery.Selection) {
		author := ""
		dynsty := ""
		content := ""
		title := ""

		title = strings.TrimSpace(s.Find("p").Find("a").Find("b").Text())
		authorAndDynsty := strings.TrimSpace(s.Find(".source").Text())
		authorAndDynstySlice := strings.Split(authorAndDynsty,"：")
		if len(authorAndDynstySlice) ==2 {
			dynsty = authorAndDynstySlice[0]
			author = authorAndDynstySlice[1]
		}
		s.Find(".contson").Each(func(i int, s *goquery.Selection) {
			content = strings.TrimSpace(s.Text())
		})

		if author !="" &&  dynsty !="" && content !="" && title !="" {
			p := db.Poem{}
			p.Author = author
			p.Content = content
			p.Title = title
			p.Dynasty = dynsty
			p.Save()
		}
	})
}