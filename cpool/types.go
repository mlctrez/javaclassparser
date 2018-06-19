package cpool

import (
	"fmt"
	"io"
	"math"
)

type CONSTANT_String_info struct {
	ConstBase
	StringIndex uint16
}

func (c *CONSTANT_String_info) String() string {
	return fmt.Sprintf("%s", c.Pool.Lookup(c.StringIndex))
}

func ReadCONSTANT_String_info(r io.Reader) *CONSTANT_String_info {
	cs := &CONSTANT_String_info{}
	cs.Tag = CONSTANT_String
	failErr(readUint16(r, &cs.StringIndex))
	return cs
}

type CONSTANT_Integer_info struct {
	ConstBase
	Value int32
}

func (c *CONSTANT_Integer_info) String() string {
	return fmt.Sprintf("%d", c.Value)
}

func ReadCONSTANT_Integer_info(r io.Reader) *CONSTANT_Integer_info {
	ci := &CONSTANT_Integer_info{}
	ci.Tag = CONSTANT_Integer
	failErr(readInt32(r, &ci.Value))
	return ci
}

type CONSTANT_Float_info struct {
	ConstBase
	Value float32
}

func (c *CONSTANT_Float_info) String() string {
	return fmt.Sprintf("%f", c.Value)
}

func ReadCONSTANT_Float_info(r io.Reader) *CONSTANT_Float_info {
	cf := &CONSTANT_Float_info{}
	cf.Tag = CONSTANT_Float
	var floatBits uint32
	failErr(readUint32(r, &floatBits))
	cf.Value = math.Float32frombits(floatBits)
	return cf
}

type CONSTANT_Long_info struct {
	ConstBase
	Value int64
}

func (c *CONSTANT_Long_info) String() string {
	return fmt.Sprintf("%d", c.Value)
}

func ReadCONSTANT_Long_info(r io.Reader) *CONSTANT_Long_info {
	cl := &CONSTANT_Long_info{}
	cl.Tag = CONSTANT_Long
	failErr(readInt64(r, &cl.Value))
	return cl
}

type CONSTANT_Double_info struct {
	ConstBase
	Value float64
}

func (c *CONSTANT_Double_info) String() string {
	return fmt.Sprintf("%f", c.Value)
}

func ReadCONSTANT_Double_info(r io.Reader) *CONSTANT_Double_info {
	cd := &CONSTANT_Double_info{}
	cd.Tag = CONSTANT_Double
	read(r, &cd.Value)
	return cd
}

type CONSTANT_NameAndType_info struct {
	ConstBase
	NameIndex       uint16
	DescriptorIndex uint16
}

func (c *CONSTANT_NameAndType_info) String() string {
	return fmt.Sprintf("%s %s", c.Pool.Lookup(c.NameIndex), c.Pool.Lookup(c.DescriptorIndex))
}

func ReadCONSTANT_NameAndType_info(r io.Reader) *CONSTANT_NameAndType_info {
	nat := &CONSTANT_NameAndType_info{}
	nat.Tag = CONSTANT_NameAndType
	failErr(readUint16(r, &nat.NameIndex))
	failErr(readUint16(r, &nat.DescriptorIndex))
	return nat
}

// TODO: this probably does not need to be anything other than a string
type CONSTANT_Utf8_info struct {
	ConstBase
	Value string
}

func (c *CONSTANT_Utf8_info) String() string {
	// TODO: this was %q but changed to %s
	return fmt.Sprintf("%s", c.Value)
}

func ReadCONSTANT_Utf8_info(r io.Reader) *CONSTANT_Utf8_info {
	u := &CONSTANT_Utf8_info{}
	u.Tag = CONSTANT_Utf8
	var length uint16
	failErr(readUint16(r, &length))
	buff := make([]uint8, length)
	read, err := io.ReadFull(r, buff)
	failErr(err)
	if length != uint16(read) {
		failErr(fmt.Errorf("incorrect length, expected %d but got %d", length, read))
	}
	u.Value = string(buff)
	return u
}
