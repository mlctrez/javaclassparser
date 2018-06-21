package cpool

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/mlctrez/javaclassparser/ioutil"
)

type ConstantStringInfo struct {
	ConstBase
	StringIndex uint16
}

func (c *ConstantStringInfo) String() string {
	return fmt.Sprintf("%s", c.Pool.Lookup(c.StringIndex))
}

func ReadConstantStringInfo(r PoolReader) (cs *ConstantStringInfo, err error) {
	cs = &ConstantStringInfo{}
	cs.Pool = r.ConstantPool
	cs.Tag = ConstantString
	cs.Type = "CONSTANT_String_info"
	err = ioutil.ReadUint16(r, &cs.StringIndex)
	return
}

type ConstantIntegerInfo struct {
	ConstBase
	Value int32
}

func (c *ConstantIntegerInfo) String() string {
	return fmt.Sprintf("%d", c.Value)
}

func ReadConstantIntegerInfo(r PoolReader) (ci *ConstantIntegerInfo, err error) {
	ci = &ConstantIntegerInfo{}
	ci.Pool = r.ConstantPool
	ci.Tag = ConstantInteger
	ci.Type = "CONSTANT_Integer_info"
	err = ioutil.ReadInt32(r, &ci.Value)
	return
}

type ConstantFloatInfo struct {
	ConstBase
	Value float32
}

func (c *ConstantFloatInfo) String() string {
	return fmt.Sprintf("%f", c.Value)
}

func ReadConstantFloatInfo(r PoolReader) (cf *ConstantFloatInfo, err error) {
	cf = &ConstantFloatInfo{}
	cf.Pool = r.ConstantPool
	cf.Tag = ConstantFloat
	cf.Type = "CONSTANT_Float_info"
	var floatBits uint32
	err = ioutil.ReadUint32(r, &floatBits)
	if err != nil {
		return
	}
	cf.Value = math.Float32frombits(floatBits)
	return
}

type ConstantLongInfo struct {
	ConstBase
	Value int64
}

func (c *ConstantLongInfo) String() string {
	return fmt.Sprintf("%d", c.Value)
}

func ReadConstantLongInfo(r PoolReader) (cl *ConstantLongInfo, err error) {
	cl = &ConstantLongInfo{}
	cl.Pool = r.ConstantPool
	cl.Tag = ConstantLong
	cl.Type = "CONSTANT_Long_info"
	err = ioutil.ReadInt64(r, &cl.Value)
	return
}

type ConstantDoubleInfo struct {
	ConstBase
	Value float64
}

func (c *ConstantDoubleInfo) String() string {
	return fmt.Sprintf("%f", c.Value)
}

func ReadConstantDoubleInfo(r PoolReader) (cd *ConstantDoubleInfo, err error) {
	cd = &ConstantDoubleInfo{}
	cd.Pool = r.ConstantPool
	cd.Tag = ConstantDouble
	cd.Type = "CONSTANT_Double_info"
	err = binary.Read(r, binary.BigEndian, &cd.Value)
	return
}

type ConstantNameAndTypeInfo struct {
	ConstBase
	NameIndex       uint16
	DescriptorIndex uint16
}

func (c *ConstantNameAndTypeInfo) String() string {
	return fmt.Sprintf("%s %s", c.Pool.Lookup(c.NameIndex), c.Pool.Lookup(c.DescriptorIndex))
}

func ReadConstantNameAndTypeInfo(r PoolReader) (nat *ConstantNameAndTypeInfo, err error) {
	nat = &ConstantNameAndTypeInfo{}
	nat.Pool = r.ConstantPool
	nat.Tag = ConstantNameAndType
	nat.Type = "CONSTANT_NameAndType_info"
	if err = ioutil.ReadUint16(r, &nat.NameIndex); err != nil {
		return
	}
	err = ioutil.ReadUint16(r, &nat.DescriptorIndex)
	return
}

// TODO: this probably does not need to be anything other than a string
type ConstantUtf8Info struct {
	ConstBase
	Value string
}

func (c *ConstantUtf8Info) String() string {
	// TODO: this was %q but changed to %s
	return fmt.Sprintf("%s", c.Value)
}

func ReadConstantUtf8Info(r PoolReader) (u *ConstantUtf8Info, err error) {
	u = &ConstantUtf8Info{}
	u.Pool = r.ConstantPool
	u.Tag = ConstantUtf8
	u.Type = "CONSTANT_Utf8_info"
	var length uint16
	if err = ioutil.ReadUint16(r, &length); err != nil {
		return
	}
	buff := make([]uint8, length)
	if _, err := io.ReadFull(r, buff); err != nil {

	}
	u.Value = string(buff)
	return
}
