package parser

import (
	"strings"
)

// ExtractName removes the inner class and leading and trailing [L ;
func ExtractName(name string) string {
	cn := name
	cn = strings.Split(cn, "$")[0]
	cn = strings.TrimPrefix(cn, "[L")
	cn = strings.TrimSuffix(cn, ";")
	return cn
}

