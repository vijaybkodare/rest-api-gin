package main

import "github.com/gin-gonic/gin"
import "fmt"
import "os"
import "vkodare.com/rest-api-gin/employee"
import "vkodare.com/rest-api-gin/address"
import "vkodare.com/rest-api-gin/employeeaddress"
import "vkodare.com/rest-api-gin/common"

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

/*func getEmployeeList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, readEmployees())
}

func getEmployeeAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.IndentedJSON(http.StatusOK, readEmployeeAddresses(id))
}

func getAllAddress(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, readAddresses())
}

func readEmployees() []Employee {
	jsonFile, _ := os.Open("employees." + appEnv + ".json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var entities []Employee
	json.Unmarshal(byteValue, &entities)
	return entities
}

func readAddresses() []EmployeeAddress {
	jsonFile, _ := os.Open("employee.address." + appEnv + ".json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var entities []EmployeeAddress
	json.Unmarshal(byteValue, &entities)
	return entities
}

func readEmployeeAddresses(id int) []EmployeeAddress {
	entities := readAddresses()
	result := make([]EmployeeAddress, 0)
    for _, e := range entities {
        if e.EmployeeId == id {
            result = append(result, e)
        }
    }
	return result
}*/