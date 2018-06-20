package cpool

import (
	"fmt"

	"github.com/mlctrez/javaclassparser/ioutil"
)

type ConstantClassInfo struct {
	ConstBase
	NameIndex uint16
}

func (c *ConstantClassInfo) String() string {
	return fmt.Sprintf("%s", c.Pool.Lookup(c.NameIndex))
}

func ReadConstantClassInfo(r PoolReader) *ConstantClassInfo {
	c := &ConstantClassInfo{}
	c.Pool = r.ConstantPool
	c.Tag = ConstantClass
	c.Type = "CONSTANT_Class_info"
	failErr(ioutil.ReadUint16(r, &c.NameIndex))
	return c
}
