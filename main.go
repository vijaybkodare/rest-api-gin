package main

import "github.com/gin-gonic/gin"
import "net/http"
import "os"
import "fmt"
import "encoding/json"
import "io/ioutil"
import "strconv"

type Employee struct {
    ID int
	FirstName string
    LastName string
    AddressId int
}

type EmployeeAddress struct{
	EmployeeId int
	Addres string
}

var appEnv string

func main() {
	appEnv = os.Getenv("APP_ENV")
	fmt.Printf("========APP_ENV: %s ========\n", appEnv)
	
	router := gin.Default()
	router.GET("/employees", getEmployeeList)
	router.GET("/addresses", getAllAddress)
	router.GET("/employeeaddresses/:id", getEmployeeAddress)

	port := "8080"
	if appEnv == "prod" {
		port = "8081"
	}
	router.Run("localhost:" + port)
}

func getEmployeeList(c *gin.Context) {
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
}