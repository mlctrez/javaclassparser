package ioutil

import (
	"archive/zip"
)

func MustOpenZipReader(name string) *zip.ReadCloser {
	closer, err := zip.OpenReader(name)
	if err != nil {
		panic(err)
	}
	return closer
}
