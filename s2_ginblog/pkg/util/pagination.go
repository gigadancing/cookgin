package util

import (
	"cookgin/s2_ginblog/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0
	if page, _ := com.StrTo(c.Query("page")).Int(); page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
