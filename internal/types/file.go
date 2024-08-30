package types

import (
	"os"
)

type File string

// Read : ...
func (f *File) Read() (*os.File, error) {
	file, err := os.Open(string(*f))
	if err != nil {
		return nil, err
	}
	//defer file.Close()

	return file, nil
}

// IsDirectory : ...
func (f *File) IsDirectory() (bool, error) {
	info, err := os.Stat(string(*f))

	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}

// ReadDirectory : ...
func (f *File) ReadDirectory(dir os.File) ([]os.FileInfo, error) {
	files, err := dir.Readdir(0)
	if len(files) == 0 || err != nil {
		return nil, err
	}

	return files, nil
}
