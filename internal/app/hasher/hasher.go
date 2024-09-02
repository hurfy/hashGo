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

// InitializeAlgorithm : initialize hash algorithm
func InitializeAlgorithm(format string) hash.Hash {
	algorithms := map[string]func() hash.Hash{
		"sha1":   sha1.New,
		"sha256": sha256.New,
		"sha512": sha512.New,
	}

	if algo, ok := algorithms[format]; !ok {
		return md5.New()
	} else {
		return algo()
	}
}

// GenerateHash : ...
func GenerateHash(path string, format hash.Hash) string {
	// read the file
	data, err := os.Open(path)
	if err != nil {
		return ""
	}

	// generate hash-sum in bytes
	if _, err := io.Copy(format, data); err != nil {
		return ""
	}
	hashInBytes := format.Sum(nil)

	return hex.EncodeToString(hashInBytes)
}
