package main

import (
	"hashGo/internal/types"
	"testing"
)

var config = types.Config{
	InputPath:  "./",
	OutputFile: "",
	Format:     "md5",
	SubDirs:    true,
	ExcDirs:    []string{},
}

func TestOk(t *testing.T) {
	if err := hashFiles(&config); err != nil {
		t.Errorf("Test Failed")
	}
}
