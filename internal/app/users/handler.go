package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u UserHandler) View(c *gin.Context) {
	var query AllowedQuery = AllowedQuery{
		PageQuery: PageQuery{
			Page:  1,
			Limit: 10,
		},
	}

	if errQuery := c.ShouldBindQuery(&query); errQuery != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "query not accepted",
			"description": errQuery.Error(),
		})
		return
	}

	errGet, users := UserService{}.Get(query)
	if errGet != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "failed fetching users!",
			"description": errGet.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success fetching users!",
		"data":    users,
	})
}
