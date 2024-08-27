package files

import (
	"hashGo/internal/app/hasher"
	"os"
	"path/filepath"
)

// isDirectory : is the file a directory?
func isDirectory(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}

// extendMap : extends the map by adding a new key:value pair
func extendMap(dMap, nMap map[string]string) {
	for k, v := range nMap {
		(dMap)[k] = v
	}
}

// readDirectory : reading a directory, getting a list of files inside
func readDirectory(path string) ([]os.FileInfo, error) {
	// check if the path is a directory
	if _, err := isDirectory(path); err != nil {
		return nil, err
	}

	// try to open the directory
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// try to read file list from the directory
	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// HashDirectory : hashing of all files inside the directory, will be called recursively if subDirs = true
func HashDirectory(path string, subDirs *bool) map[string]string {
	var hashMap = make(map[string]string)

	// try to read the directory
	files, err := readDirectory(path)
	if err != nil {
		return hashMap
	}

	// hash files
	for _, v := range files {
		// create file path
		var filePath = filepath.Join(path, v.Name())

		if *subDirs && v.IsDir() {
			// recursive hashMap extension
			if _map := HashDirectory(filePath, subDirs); _map != nil {
				extendMap(hashMap, _map)
			}
		} else {
			// generate hash
			hash, err := hasher.GenerateHash(filePath, "md5")
			if err != nil {
				continue
			}

			hashMap[filePath] = hash
		}
	}

	return hashMap
}
