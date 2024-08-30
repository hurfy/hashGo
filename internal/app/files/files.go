package files

import (
	"hashGo/internal/app/hasher"
	"hashGo/internal/types"
	"path/filepath"
)

// HashDirectory : hashing of all files inside the directory, will be called recursively if subDirs = true
func HashDirectory(hashes *types.HashMap, path *string, subDirs *bool, format *string) error {
	var directory = types.File(*path)

	// is the path a directory?
	isDir, err := directory.IsDirectory()
	if !isDir || err != nil {
		return err
	}

	// try to open
	dir, err := directory.Read()
	if err != nil {
		return err
	}

	// try to read the directory
	files, err := directory.ReadDirectory(*dir)
	if err != nil {
		return err
	}

	// hash files
	for _, v := range files {
		var filePath = filepath.Join(*path, v.Name())

		if *subDirs && v.IsDir() {
			// recursive hashMap extension
			if err := HashDirectory(hashes, &filePath, subDirs, format); err != nil {
				return err
			}
		} else {
			// generate hash
			if hash, err := hasher.GenerateHash(types.File(filePath), *format); err != nil {
				continue
			} else {
				hashes.Append(filePath, hash)
			}
		}
	}

	return nil
}
