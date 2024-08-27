package json

import (
	"encoding/json"
	"os"
)

// SaveJSON : serialization of map[string]string to JSON and saving it to a file
func SaveJSON(path *string, data map[string]string) error {
	// serialize
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// save
	err = os.WriteFile(*path, jsonString, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
