package main

import (
	"flag"
	"fmt"
	"hashGo/internal/app/files"
	"hashGo/internal/app/json"
)

var (
	inputPath, outputPath, format string
	subDirs, show, save           bool
)

func configureFlags() {
	flag.StringVar(&inputPath, "p", "./", "Path to directory")
	flag.StringVar(&format, "f", "md5", "Output format (md5, sha1, sha256, sha512)")
	flag.BoolVar(&show, "o", false, "Show output")
	flag.BoolVar(&subDirs, "sd", false, "Include subdirectories")
	flag.BoolVar(&save, "s", false, "Save output to file")
	flag.StringVar(&outputPath, "op", "./output.json", "Path to output file")

	flag.Parse()
}

// hashFiles : ...
func hashFiles() error {
	var res = files.HashDirectory(inputPath, &subDirs)

	// Show output
	if show {
		for k, v := range res {
			fmt.Printf("%s: %s\n", k, v)
		}
	}

	// Save to json
	if save {
		err := json.SaveJSON(&outputPath, res)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func main() {
	// Configure flags
	configureFlags()

	// Hash files
	err := hashFiles()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done!\nPress Enter to exit...")
	}

	// End
	fmt.Scanln()
}
