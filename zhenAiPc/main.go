package main

import (
	"goDemo/zhenAiPc/city"
	"goDemo/zhenAiPc/common"
	"strconv"
	"time"
)

func show(str string) {
	for i := 1; i <= 6; i++ {
		common.GetUserList(str, strconv.Itoa(i))
	}
}

func main() {

	cityList := city.ParseCity()

	for _, m := range cityList.CityList {
		time.Sleep(time.Millisecond * 100)
		go show(m.Href)
	}

}
