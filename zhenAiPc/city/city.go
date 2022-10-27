package city

import (
	"github.com/PuerkitoBio/goquery"
	"goProject/zhenAiPc/common"
	_struct "goProject/zhenAiPc/struct"
)

func ParseCity() _struct.CityList {
	dom, _ := common.GetHttp("http://www.zhenai.com/zhenghun")
	result := _struct.CityList{}
	dom.Find(".city-list a").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(i)
		href, _ := s.Attr("href")
		text := s.Text()
		//fmt.Println(text,href)
		result.CityList = append(result.CityList, _struct.CityListName{
			Name: text,
			Href: href,
		})
	})
	return result
}
