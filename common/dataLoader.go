package common

import "encoding/json"
import "io/ioutil"
import "os"
import . "vkodare.com/rest-api-gin/model"

func LoadData[T Employee | EmployeeAddress | Address](fileName string) []T {
	jsonFile, _ := os.Open("resource/" + AppEnv + "/" + fileName + ".json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var models []T
	json.Unmarshal(byteValue, &models)
	return models
}
