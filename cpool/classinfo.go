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

func ReadConstantClassInfo(r PoolReader, index uint16) (c *ConstantClassInfo, err error) {
	c = &ConstantClassInfo{}
	c.Index = index
	c.Pool = r.ConstantPool
	c.Tag = ConstantClass
	c.Type = "CONSTANT_Class_info"
	err = ioutil.ReadUint16(r, &c.NameIndex)
	return
}
