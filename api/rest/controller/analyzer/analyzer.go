package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venusaran/web-analyzer/internal/service/scraper"
	interfaces "github.com/venusaran/web-analyzer/pkg/interfaces"
)

type TaskController interface {
	ExecuteWebScrapingTask(*gin.Context)
}

type taskController struct {
	srv scraper.ScraperService
}

func NewTaskController() TaskController {
	return &taskController{}
}

func (t taskController) ExecuteWebScrapingTask(c *gin.Context) {
	var payload interfaces.TargetURL
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	linkStats, data, err := t.srv.RetrieveData(payload.URL)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"link_stats": linkStats,
			"page_data":  data,
			"status":     http.StatusOK,
			"message":    "Data Extracted",
		},
	)
}
