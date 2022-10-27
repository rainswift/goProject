package utils

import (
	"github.com/gin-gonic/gin"
	"server/models"
	"strconv"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {

	limit := 20
	page := 1
	sort := "created_at asc"
	gender := ""
	age := ""
	height := ""
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		case "gender":
			gender = queryValue
			break
		case "age":
			age = queryValue
			break
		case "height":
			height = queryValue
			break
		}
	}
	return models.Pagination{
		Limit:  limit,
		Page:   page,
		Sort:   sort,
		Gender: gender,
		Age:    age,
		Height: height,
	}

}
