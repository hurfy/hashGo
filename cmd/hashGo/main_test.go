package main

import (
	"testing"
)

var config = Config{
	inputPath:  "./",
	outputPath: "",
	format:     "md5",
	subDirs:    true,
}

func TestOk(t *testing.T) {
	if err := hashFiles(&config); err != nil {
		t.Errorf("Test Failed")
	}
}
