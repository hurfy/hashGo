package main

import (
	"flag"
	"fmt"
	"hashGo/internal/app/files"
	"hashGo/internal/types"
)

var (
	inputPath, outputPath, format string
	subDirs, show, save           bool
)

// configureFlags : configures and parses flags
func configureFlags() {
	flag.StringVar(&inputPath, "p", "./", "Path to directory")
	flag.StringVar(&format, "f", "md5", "Output format (md5, sha1, sha256, sha512)")
	flag.BoolVar(&show, "o", false, "Show output")
	flag.BoolVar(&subDirs, "sd", false, "Include subdirectories")
	flag.BoolVar(&save, "s", false, "Save output to file")
	flag.StringVar(&outputPath, "op", "./output.json", "Path to output file")

	flag.Parse()
}

// hashFiles : basic hashing function, responsible for hashing, output to console, saving data to file
func hashFiles() error {
	var hashes = make(types.HashMap)

	err := files.HashDirectory(&hashes, &inputPath, &subDirs, &format)
	if err != nil {
		return err
	}

	// Show output
	if show {
		for k, v := range hashes {
			fmt.Printf("%s: %s\n", k, v)
		}
	}

	// Save to json
	if save {
		if err := hashes.SaveAsJson(&outputPath); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// Configure flags
	configureFlags()

	// Hash files
	if err := hashFiles(); err != nil {
		fmt.Println(err, "\nPress Enter to exit...")
	} else {
		fmt.Println("Done!\nPress Enter to exit...")
	}

	// End
	fmt.Scanln()
}
