package files

import (
	"hashGo/internal/app/hasher"
	"os"
	"path/filepath"
)

// isDirectory : ...
func isDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}

// extendMap : ...
func extendMap(dMap, nMap map[string]string) {
	for k, v := range nMap {
		(dMap)[k] = v
	}
}

// readDirectory : ...
func readDirectory(path string) ([]os.FileInfo, error) {
	// Check if the path is a directory
	if _, err := isDirectory(path); err != nil {
		return nil, err
	}

	// Try to open the directory
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// Try to read file list from the directory
	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// HashDirectory : ...
func HashDirectory(path string, subDirs *bool) map[string]string {
	var hashMap = make(map[string]string)

	// Try to read the directory
	files, err := readDirectory(path)
	if err != nil {
		return hashMap
	}

	// Hash files
	for _, v := range files {
		// Create file path
		var filePath = filepath.Join(path, v.Name())

		if *subDirs && v.IsDir() {
			// Recursive hashMap extension
			if _map := HashDirectory(filePath, subDirs); _map != nil {
				extendMap(hashMap, _map)
			}
		} else {
			// Generate hash
			hash, err := hasher.GenerateHash(filePath, "md5")
			if err != nil {
				continue
			}

			hashMap[filePath] = hash
		}
	}

	return hashMap
}
