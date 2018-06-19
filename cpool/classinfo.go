package cpool

import (
	"fmt"
	"io"
)

type CONSTANT_Class_info struct {
	ConstBase
	NameIndex uint16
}

func (c *CONSTANT_Class_info) String() string {
	return fmt.Sprintf("%s", c.Pool.Lookup(c.NameIndex))
}

func ReadCONSTANT_Class_info(r io.Reader) *CONSTANT_Class_info {
	c := &CONSTANT_Class_info{}
	c.Tag = CONSTANT_Class
	failErr(readUint16(r, &c.NameIndex))
	return c
}
