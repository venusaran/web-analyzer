package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controller "github.com/venusaran/web-analyzer/api/rest/controller/analyzer"
)

type Router struct {
	taskController controller.TaskController
}

func NewRouter(taskController controller.TaskController) *Router {
	return &Router{
		taskController: taskController,
	}
}

func (router Router) SetUpGinEngine(r *gin.Engine) *gin.Engine {
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", " Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))

	v1 := r.Group("/v1/analyzer")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"status":  http.StatusOK,
					"message": "I'm OK!",
					"data":    nil,
				},
			)
		})

		baseGroup := v1.Group("")
		baseGroup.POST("", router.taskController.ExecuteWebScrapingTask)
	}

	return r
}
