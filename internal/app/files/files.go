package files

import (
	"fmt"
	"hash"
	"hashGo/internal/app/hasher"
	"hashGo/internal/types"
	"io/fs"
	"os"
)

// HashDirectory : ...
func HashDirectory(root string, subDirs bool, format hash.Hash) (types.HashMap, error) {
	var hashes = make(types.HashMap)

	if err := fs.WalkDir(
		os.DirFS(root),
		".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			// skip directories if required
			if !subDirs && path != "." && d.IsDir() {
				return fs.SkipDir
			}
			// generate hash if not a directory
			if !d.IsDir() {
				fullPath := fmt.Sprintf("%v/%v", root, path)
				hashes[fullPath] = hasher.GenerateHash(fullPath, format)
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return hashes, nil
}
