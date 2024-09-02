package main

import (
	"flag"
	"fmt"
	"hash"
	"hashGo/internal/app/files"
	"hashGo/internal/app/hasher"
)

type Config struct {
	inputPath  string
	outputPath string
	subDirs    bool
	format     string
}

// initHashAlgo : ...
func (c *Config) initHashAlgo() hash.Hash {
	return hasher.InitializeAlgorithm(c.format)
}

// configureFlags : configures and parses flags
func configureFlags() *Config {
	var config = new(Config)

	flag.StringVar(&config.inputPath, "p", "./", "Root directory")
	flag.StringVar(&config.outputPath, "o", "", "Path to output file")
	flag.StringVar(&config.format, "f", "md5", "Hash format[md5, sha1, sha256, sha512]")
	flag.BoolVar(&config.subDirs, "s", false, "Include subdirectories")
	flag.Parse()

	return config
}

// hashFiles : basic hashing function, responsible for hashing, output to console, saving data to file
func hashFiles(config *Config) error {
	hashes, err := files.HashDirectory(config.inputPath, config.subDirs, config.initHashAlgo())
	if err != nil {
		return err
	}

	// if the path to output file is not specified, we will print result
	if config.outputPath != "" {
		if err := hashes.SaveAsJson(config.outputPath); err != nil {
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
	var config = configureFlags()

	if err := hashFiles(config); err != nil {
		panic(err.Error())
	}
}
