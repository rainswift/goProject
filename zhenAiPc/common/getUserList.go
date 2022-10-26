package common

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_struct "goDemo/zhenAiPc/struct"
	"regexp"
)

const cityListRe = `：\S*`

func reg(s string) string {

	re := regexp.MustCompile(cityListRe)
	result1 := re.FindAllStringSubmatch(s, -1)
	var str = ""
	if len(result1) != 0 {
		str = result1[0][0]
		nameRune := []rune(str)
		str = string(nameRune[1:])
	}

	return str
}

func GetUserList(href string, page string) {
	fmt.Println(href, page)
	//url := "https://www.zhenai.com/zhenghun/beijing/"+string(page)
	var url bytes.Buffer
	url.WriteString(href)
	url.WriteString("/" + page)

	fmt.Printf("buffer.String(): %v\n", url.String())
	dom, _ := GetHttp(url.String())

	dom.Find(".list-item").Each(func(i int, s *goquery.Selection) {

		hrefs, _ := s.Find(".photo a").Attr("href")

		photo, _ := s.Find(".photo a img").Attr("src")
		name := s.Find("th").Text()

		//性别：女士	居住地：福建厦门
		//年龄：34	学   历：大专
		//婚况：离异	身   高：160
		sex := s.Find("td").Eq(0).Text()
		habitation := s.Find("td").Eq(1).Text()
		Age := s.Find("td").Eq(2).Text()
		Education := s.Find("td").Eq(3).Text()
		House := s.Find("td").Eq(4).Text()
		Height := s.Find("td").Eq(5).Text()

		user := _struct.UserXyy{
			Name:       name,
			Href:       hrefs,
			Photo:      photo,
			Gender:     reg(sex),
			Habitation: reg(habitation),
			Age:        reg(Age),
			Education:  reg(Education),
			House:      reg(House),
			Height:     reg(Height),
		}

		//fmt.Println(user)

		if name != "" {
			Mgr.AddUser(&user)
			//getUserDetails(href,user)
		}

	})
}
