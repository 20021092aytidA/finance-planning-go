package subscriptions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s SubscriptionHandler) View(c *gin.Context) {
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

	errGet, subs := SubscriptionService{}.Get(&query)
	if errGet != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "failed fetching subscriptions!",
			"description": errGet.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success fetching subscriptions!",
		"data":    subs,
	})
}

func (s SubscriptionHandler) Create(c *gin.Context) {
	var postBody PostModel
	if errBody := c.ShouldBind(&postBody); errBody != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "missing body!",
			"description": errBody.Error(),
		})
		return
	}

	errPost := SubscriptionService{}.Post(&postBody)
	if errPost != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "failed creating subscription!",
			"description": errPost.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "success creating subscription!",
		"data":    postBody,
	})
}

func (s SubscriptionHandler) Drop(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"message":     "missing param!",
			"description": "param needed",
		})
		return
	}

	errDel := SubscriptionService{}.Delete(id)
	if errDel != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      http.StatusInternalServerError,
			"message":     "failed deleting subscription!",
			"description": errDel.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success deleting subscription",
	})
}
