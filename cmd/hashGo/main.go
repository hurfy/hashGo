package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"hashGo/internal/app/hasher"
)

var (
	path, format string
	subDirs      bool
)

func configureFlags() {
	// TODO: Change default values in the future
	pflag.StringVarP(&path, "path", "p", "./Makefile", "Path to directory")
	pflag.StringVarP(&format, "format", "f", "md5", "Output format (md5, sha1, sha256, sha512)")
	pflag.BoolVarP(&subDirs, "subdirs", "s", false, "Include subdirectories")

	pflag.Parse()
}

func main() {
	// Configure flags
	configureFlags()

	// Temporary function
	hash, err := hasher.GenerateHash(path, format)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Hash: %s\n", hash)

	// Temporary delay
	fmt.Println("\nSuccess! Press Enter to exit: ")
	fmt.Scanln()
}
