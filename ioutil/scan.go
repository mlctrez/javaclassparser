package ioutil

import (
	"bufio"
	"io"
)

func ScanLines(reader io.ReadCloser, eachLine func(line string)) {
	defer reader.Close()
	sc := bufio.NewScanner(reader)
	for sc.Scan() {
		eachLine(sc.Text())
	}
}

func MustOpen(r io.ReadCloser, err error) io.ReadCloser {
	if err != nil {
		panic(err)
	}
	return r
}
