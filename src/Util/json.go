package util

/*
File Type : API
Perpose : Have to functions to deal with Json Objects
*/
import (
	"encoding/json"
	"errors"
)

/*
todo : read the json file and put it into a data type
Arguments ---------------------
file_path - The file path of the json file which we want to read
JsonData - The struct where you want to store the data

Return ---------------------
Data, error
If you pass no object to store JsonData then it will return map[string]interface{}
Else if will return nil , nil and udpdat ethe data inside the Struct you provided
*/
func GetJson(JsonData *string, JsonObj ...any) (*map[string]interface{}, error) {

	if len(JsonObj) > 0 {
		// Create a map to store the JSON data
		err := json.Unmarshal([]byte(*JsonData), &JsonObj[0])
		if err != nil {
			return nil, errors.New("Error Unmarshal the file: " + err.Error())
		}
		return nil, err
	}
	var _jsonData map[string]interface{}
	// Unmarshal the JSON data into the map
	err := json.Unmarshal([]byte(*JsonData), &_jsonData)
	if err != nil {
		return nil, errors.New("Error Unmarshal the file: " + err.Error())
	}
	return &_jsonData, nil
}
