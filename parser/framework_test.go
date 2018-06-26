package parser

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	c := &Config{Archive:"sample.zip"}
	c.LogElapsed=false
	for n := 0; n < b.N; n++ {
		Scan(c, func(work *Work) {})
	}
}
