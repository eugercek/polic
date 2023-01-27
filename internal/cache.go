package internal

import (
	"os"
	"path/filepath"
	"time"
)

var CacheFile = filepath.Join(os.TempDir(), "polic.cache")

func CacheOk() bool {
	finfo, err := os.Stat(CacheFile)
	if err != nil {
		return false
	}

	if time.Now().Sub(finfo.ModTime()).Hours() > 12 {
		return false
	}

	return true
}

func GetCache() ([]byte, error) {
	bs, err := os.ReadFile(CacheFile)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func FillCache(bs []byte) error {
	if err := os.WriteFile(CacheFile, bs, os.ModePerm); err != nil {
		return err
	}

	return nil
}
