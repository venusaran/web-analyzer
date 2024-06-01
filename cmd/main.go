package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	controller "github.com/venusaran/web-analyzer/api/rest/controller/analyzer"
	"github.com/venusaran/web-analyzer/api/rest/router"
	"github.com/venusaran/web-analyzer/pkg/util"
)

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
