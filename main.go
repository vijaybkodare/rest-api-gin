package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"vkodare.com/rest-api-gin/common"
	"vkodare.com/rest-api-gin/controller/address"
	"vkodare.com/rest-api-gin/controller/employee"
	"vkodare.com/rest-api-gin/controller/employeeaddress"
)

func main() {
	if len(os.Getenv("APP_ENV")) > 0 {
		common.AppEnv = os.Getenv("APP_ENV")
	}
	fmt.Printf("========APP_ENV: %s ========\n", common.AppEnv)

	router := gin.Default()
	router.GET("/employee", employee.GetAll)
	router.GET("/address", address.GetAll)
	router.GET("/employeeaddress", employeeaddress.GetAll)

	router.GET("/employeeaddress/:id", employeeaddress.Get)
	router.GET("/employee/:id", employee.Get)
	router.GET("/address/:id", address.Get)

	port := "8080"
	if common.AppEnv == "prod" {
		port = "8081"
	}
	router.Run("localhost:" + port)
}
