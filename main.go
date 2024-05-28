package main

import (
	"Lineman_project/api"
	"Lineman_project/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	covidController controller.DataController = controller.New()
	covidService    api.ApiService            = api.New()
)

func getCovidSummary(c *gin.Context) {
	res := covidController.GetCovidSummary()

	c.JSON(http.StatusOK, res)

}

func handlerRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/covid/summary", getCovidSummary)
	return router
}

func main() {
	r := handlerRouter()

	r.Run(":8080")
}
