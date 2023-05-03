package checker

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// read json file from file path
func ReadJsonFile(filePath string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	return result, nil
}
