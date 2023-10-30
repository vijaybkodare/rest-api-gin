package address

import "github.com/gin-gonic/gin"
import "net/http"
import "vkodare.com/rest-api-gin/common"
import "vkodare.com/rest-api-gin/model"
import "strconv"

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getAll())
}

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.IndentedJSON(http.StatusOK, get(id))
}

func getAll() []model.Address{
	return common.LoadData[model.Address]("address")
}

func get(id int) model.Address{
	models := getAll()
    for _, e := range models {
        if e.Id == id {
            return e
        }
    }
	return model.Address{}
}