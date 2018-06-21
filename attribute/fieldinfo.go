package attribute

import (
	"fmt"

	"github.com/mlctrez/javaclassparser/aflag"
	"github.com/mlctrez/javaclassparser/cpool"
	"github.com/mlctrez/javaclassparser/ioutil"
)

// FieldInfo describes a field on a java class
// See https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.5
type FieldInfo struct {
	baseAttribute
	AccessFlags     aflag.FieldAccessFlags
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []interface{}
}

func (c *FieldInfo) String() string {
	af := c.AccessFlags
	name := c.Pool.Lookup(c.NameIndex)
	descriptor := c.Pool.Lookup(c.DescriptorIndex)
	attributes := c.Attributes
	return fmt.Sprintf("%s %s %s %s", af, name, descriptor, attributes)
}

// ReadFieldInfo creates new field_info from the provided reader
func ReadFieldInfo(r cpool.PoolReader) (fi *FieldInfo, err error) {
	fi = &FieldInfo{}
	fi.Pool = r.ConstantPool

	if fi.AccessFlags, err = aflag.ReadFieldAccessFlags(r); err != nil {
		return
	}
	if err = ioutil.ReadUint16(r, &fi.NameIndex); err != nil {
		return
	}
	if err = ioutil.ReadUint16(r, &fi.DescriptorIndex); err != nil {
		return
	}

	var count uint16
	if err = ioutil.ReadUint16(r, &count); err != nil {
		return
	}
	fi.Attributes = make([]interface{}, count)
	var i uint16
	for i = 0; i < count; i++ {
		if fi.Attributes[i], err = ReadAttributeInfo(r); err != nil {
			return
		}
	}
	return
}

