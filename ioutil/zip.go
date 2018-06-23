package ioutil

import (
	"archive/zip"
	"bytes"
	"io"
)

func MustOpenZipReader(name string) *zip.ReadCloser {
	closer, err := zip.OpenReader(name)
	if err != nil {
		panic(err)
	}
	return closer
}

func ReadZipFile(f *zip.File) (bb *bytes.Buffer, err error) {
	var closer io.ReadCloser
	if closer, err = f.Open(); err != nil {
		return
	}
	defer closer.Close()
	bb = &bytes.Buffer{}
	_, err = io.Copy(bb, closer)
	return
}
