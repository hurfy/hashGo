package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hashGo/internal/types"
	"io"
)

// initializeAlgorithm : initialize hash algorithm
func initializeAlgorithm(format string) hash.Hash {
	var hashAlgorithm hash.Hash

	switch format {
	case "sha1":
		hashAlgorithm = sha1.New()

	case "sha256":
		hashAlgorithm = sha256.New()

	case "sha512":
		hashAlgorithm = sha512.New()

	default:
		hashAlgorithm = md5.New()
	}

	return hashAlgorithm
}

// GenerateHash : ...
func GenerateHash(file types.File, format string) (string, error) {
	var hashAlgo = initializeAlgorithm(format)

	// read the file
	data, err := file.Read()
	if err != nil {
		return "", err
	}

	// generate hash-sum in bytes
	if _, err := io.Copy(hashAlgo, data); err != nil {
		return "", err
	}
	hashInBytes := hashAlgo.Sum(nil)

	return hex.EncodeToString(hashInBytes), nil
}
