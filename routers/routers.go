package routers

import (
	"github.com/Abhishek-Mali-Simform/assessments/handlers"
	"github.com/gin-gonic/gin"
)

var Route = gin.Default()

func InitRoute() {
	Route.GET("/person/:person_id/info", handlers.RetrievePersonInfo)
	Route.POST("/person/create", handlers.CreatePersonInfo)

}
