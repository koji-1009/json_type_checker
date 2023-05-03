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

// read json file froma file path to struct of CheckerRule
func ReadCheckerRuleJsonFile(filePath string) (*CheckerRule, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result *CheckerRule
	json.Unmarshal([]byte(byteValue), &result)
	return result, nil
}

// CheckerRule is a struct of checker rule
type CheckerRule struct {
	FileName     string            `json:"file_name"`
	TopLevelType string            `json:"top_level_type"`
	Types        []CheckerRuleType `json:"types"`
}

type CheckerRuleType struct {
	TypeName  string          `json:"type_name"`
	ValueType json.RawMessage `json:"value_type"`
}
