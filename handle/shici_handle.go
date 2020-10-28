package handle

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gushiwen/db"
	"gushiwen/gofish"
	"io"
	"strings"
)

const Base = "https://so.gushiwen.cn"

type AuthorHandle struct {

}

// @Summer 作者列表
func (a *AuthorHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Errorf("doc err:%s", err)
	}

	doc.Find(".sons").Find(".cont").Find("a").Each(func(i int, s *goquery.Selection) {
		link,_ := s.Attr("href")

		h := HomePageHandle{}
		fish := gofish.NewGoFish()
		request, err := gofish.NewRequest("GET",  Base+link, gofish.UserAgent, &h,nil)

		if err != nil {
			fmt.Println(err)
			return
		}

		fish.Request = request
		fish.Visit()
	})
}

// @Summer 作者主页
type HomePageHandle struct {

}

func (h *HomePageHandle) Worker(body io.Reader, url string)  {
	doc,err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Errorf("home page doc err:%s", err)
	}

	doc.Find(".sonspic").Find(".cont").Find("p").Find("a").Each(func(i int, s *goquery.Selection) {
		link,_ := s.Attr("href")

		h := PoetryList{}

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

// @Summer 诗词列表
type PoetryList struct {

}

func (p *PoetryList) Worker(body io.Reader, url string)  {
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