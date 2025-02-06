package main

import (
	"fmt"
	"hashGo/internal/app/files"
	"hashGo/internal/types"
	"strings"
)

// hashFiles : basic hashing function, responsible for hashing, output to console, saving data to file
func hashFiles(config *types.Config) error {
	hashes, err := files.HashDirectory(config)
	if err != nil {
		return err
	}

	// if the path to output file is not specified, we will print result
	if config.OutputFile != "" {

		// add file extension
		if !strings.HasSuffix(config.OutputFile, ".json") {
			config.OutputFile += ".json"
		}

		if err := hashes.SaveAsJson(config.OutputFile); err != nil {
			return err
		}
	} else {
		for k, v := range hashes {
			fmt.Printf("%s: %s\n", k, v)
		}
	}

	return nil
}

func main() {
	var config = types.ConfigureFlags()

	if err := hashFiles(config); err != nil {
		panic(err.Error())
	}
}
