package employee

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

func getAll() []model.Employee {
	return common.LoadData[model.Employee]("employee")
}

func get(id int) model.Employee {
	models := getAll()
	for _, e := range models {
		if e.Id == id {
			return e
		}
	}
	return model.Employee{}
}
