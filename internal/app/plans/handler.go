package plans

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p PlanHandler) View(c *gin.Context) {
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

	errGet, plans := PlanService{}.Get(&query)
	if errGet != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "failed fetching plans!",
			"description": errGet.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success fetching plans!",
		"data":    plans,
	})
}

func (p PlanHandler) Create(c *gin.Context) {
	var postBody PostModel
	if errBody := c.ShouldBind(&postBody); errBody != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "missing body!",
			"description": errBody.Error(),
		})
		return
	}

	errPost := PlanService{}.Post(&postBody)
	if errPost != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "failed creating plan!",
			"description": errPost.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "success creating plan!",
		"data":    postBody,
	})

}

func (p PlanHandler) Drop(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "missing param!",
			"description": "param is needed!",
		})
		return
	}

	errDel := PlanService{}.Delete(id)
	if errDel != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "failed deleting plan!",
			"description": errDel.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success deleting plan!",
	})
}
