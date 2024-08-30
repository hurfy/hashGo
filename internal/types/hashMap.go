package types

import (
	"encoding/json"
	"os"
)

type HashMap map[string]string

// Extend : Extends using key:value pairs from the specified map
func (hm *HashMap) Extend(_map map[string]string) {
	for k, v := range _map {
		(*hm)[k] = v
	}
}

// Append : Adds using k and v
func (hm *HashMap) Append(k, v string) {
	(*hm)[k] = v
}

// serializeToJson : ...
func (hm *HashMap) serializeToJson() ([]byte, error) {
	jsonString, err := json.Marshal(*hm)
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}

// SaveAsJson : ...
func (hm *HashMap) SaveAsJson(path *string) error {
	// serialize
	jsonString, err := hm.serializeToJson()
	if err != nil {
		return err
	}

	// save
	if err = os.WriteFile(*path, jsonString, os.ModePerm); err != nil {
		return err
	}

	return nil
}
