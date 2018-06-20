package main

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/mlctrez/javaclassparser/cpool"
)

func main() {

	var info interface{}
	info = cpool.ReadConstantStringInfo(bytes.NewBuffer([]byte{0, 0, 0, 0}))
	_ = info

	fmt.Println(reflect.TypeOf(info).String())
	if cb, ok := info.(cpool.ConstBase); ok {
		fmt.Println(cb, ok)
	}

}
