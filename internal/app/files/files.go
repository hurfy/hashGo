package files

import (
	"hashGo/internal/app/hasher"
	"hashGo/internal/app/utils"
	"hashGo/internal/types"
	"io/fs"
	"os"
	"path/filepath"
)

// HashDirectory : ...
func HashDirectory(config *types.Config) (types.HashMap, error) {
	var hashes = make(types.HashMap)

	if err := fs.WalkDir(
		os.DirFS(config.InputPath),
		".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// skip directories if required
			if !config.SubDirs && path != "." && d.IsDir() {
				return fs.SkipDir
			}

			// excluded
			if utils.Contains(config.ExcDirs, path) {
				return fs.SkipDir
			}

			// generate hash if not a directory
			if !d.IsDir() {
				fullPath := filepath.Join(config.InputPath, path)
				hashes[fullPath] = hasher.GenerateHash(fullPath, config.InitHashAlgo())
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return hashes, nil
}
