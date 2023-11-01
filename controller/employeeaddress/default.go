package employeeaddress

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"vkodare.com/rest-api-gin/common"
	"vkodare.com/rest-api-gin/model"
)

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getAll())
}

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.IndentedJSON(http.StatusOK, get(id))
}

func getAll() []model.EmployeeAddress {
	return common.LoadData[model.EmployeeAddress]("employeeaddress")
}

func get(id int) []model.EmployeeAddress {
	models := getAll()
	result := make([]model.EmployeeAddress, 0)
	for _, e := range models {
		if e.EmployeeId == id {
			result = append(result, e)
		}
	}
	return result
}
