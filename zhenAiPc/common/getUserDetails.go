package common

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_struct "goDemo/zhenAiPc/struct"
)

func getUserDetails(href string, userInfo _struct.UserXyy) {
	dom, _ := GetHttp(href)

	var house, age, xinzuo, height, weight, occupation, income, work, education string
	dom.Find(".purple-btns .purple").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		fmt.Println(i, text)
		switch i {
		case 0:
			house = text
		case 1:
			age = text
		case 2:
			xinzuo = text
		case 3:
			height = text
		case 4:
			weight = text
		case 5:
			occupation = text
		case 6:
			income = text
		case 7:
			work = text
		case 8:
			education = text
		}

	})
	user := _struct.UserXyy{
		Name:       userInfo.Name,
		Href:       href,
		Photo:      userInfo.Photo,
		House:      house,
		Age:        age,
		Xinzuo:     xinzuo,
		Height:     height,
		Weight:     weight,
		Occupation: occupation,
		Income:     income,
		Work:       work,
		Education:  education,
	}
	Mgr.AddUser(&user)

	//House      string //结婚
	//Age        string
	//age := dom.Find(".purple-btns .purple").Eq(1).Text()
	//Xinzuo     string //星座
	//Height     string
	//Weight     string
	//Occupation string //工作地
	//Income     string //收入
	//Work       string //职位
	//Education  string //教育学历
	//Gender     string //性别
	//Marriage   string //结婚
	//Car        string //车

}
