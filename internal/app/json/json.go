package json

import (
	"encoding/json"
	"os"
)

func SaveJSON(path *string, data map[string]string) error {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(*path, jsonString, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
