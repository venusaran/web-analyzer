package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controller "github.com/venusaran/web-analyzer/api/rest/controller/analyzer"
	"github.com/venusaran/web-analyzer/api/rest/router"
	_ "github.com/venusaran/web-analyzer/docs"
	"github.com/venusaran/web-analyzer/pkg/util"
)

// @title 	Web Analyzer API
// @version	1.0
// @description An API to analyze a web page and it's contents

// @host 	localhost:8080
// @BasePath /v1/analyzer
func main() {
	l := setUpListener()
	m := cmux.New(l)
	httpl := m.Match(cmux.Any())
	go serveHTTP(httpl, constructTaskController())
	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		panic(err)
	}
}

func serveHTTP(l net.Listener, taskController controller.TaskController) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// to serve static files
	r.Static("/static", "./static")

	// static route for landing page
	r.GET("/", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "static/index.html") })

	// add swagger route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	newRouter := router.NewRouter(taskController)
	newRouter.SetUpGinEngine(r)
	s := &http.Server{
		Handler: r,
	}
	fmt.Printf("Web Scraping Server Running@%v\n ", l.Addr())
	err := s.Serve(l)
	if err != nil {
		panic(err)
	}
}

func setUpListener() net.Listener {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", util.GetEnv("PORT", "8080")))
	if err != nil {
		panic(err)
	}
	return l
}

func constructTaskController() controller.TaskController {
	return controller.NewTaskController()
}
