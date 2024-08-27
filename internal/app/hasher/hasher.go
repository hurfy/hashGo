package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

// initializeAlgorithm ...
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

// GenerateHash : generates a hash of the given file path using the specified format.
func GenerateHash(path, format string) (string, error) {
	var hashAlgo = initializeAlgorithm(format)

	// Read the file
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Generate hash-sum in bytes
	if _, err := io.Copy(hashAlgo, file); err != nil {
		return "", err
	}
	hashInBytes := hashAlgo.Sum(nil)

	return hex.EncodeToString(hashInBytes), nil
}
