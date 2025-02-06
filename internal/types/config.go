package types

import (
	"flag"
	"hash"
	"hashGo/internal/app/hasher"
	"strings"
)

type Config struct {
	InputPath  string
	OutputFile string
	SubDirs    bool
	ExcDirs    []string
	Format     string
}

// initHashAlgo : ...
func (c *Config) InitHashAlgo() hash.Hash {
	return hasher.InitializeAlgorithm(c.Format)
}

// configureFlags : configures and parses flags
func ConfigureFlags() *Config {
	var config = new(Config)

	flag.StringVar(&config.InputPath, "p", "./", "Root directory")
	flag.StringVar(&config.OutputFile, "o", "", "Output file name")
	flag.StringVar(&config.Format, "f", "md5", "Hash format[md5, sha1, sha256, sha512]")
	flag.BoolVar(&config.SubDirs, "s", false, "Include subdirectories")
	flag.Func("e", "Exclude directories (semicolon-separated)", func(s string) error {
		config.ExcDirs = strings.Split(s, ";")
		return nil
	})

	flag.Parse()

	return config
}
