package types

import (
	"encoding/json"
	"os"
)

type HashMap map[string]string

// serializeToJson : ...
func (hm *HashMap) serializeToJson() ([]byte, error) {
	jsonString, err := json.Marshal(*hm)
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}

// SaveAsJson : ...
func (hm *HashMap) SaveAsJson(path string) error {
	// serialize
	jsonString, err := hm.serializeToJson()
	if err != nil {
		return err
	}

	// save
	if err = os.WriteFile(path, jsonString, os.ModePerm); err != nil {
		return err
	}

	return nil
}
