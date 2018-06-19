package cpool

import (
	"fmt"
	"io"
)

type RefBase struct {
	ConstBase
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (rb *RefBase) ReadRefBaseIndexes(r io.Reader) {
	failErr(readUint16(r,&rb.ClassIndex))
	failErr(readUint16(r,&rb.NameAndTypeIndex))
}


type CONSTANT_Fieldref_info struct{ RefBase }

func (c *CONSTANT_Fieldref_info) String() string {
	return fmt.Sprintf("%s %s", c.Pool.Lookup(c.ClassIndex), c.Pool.Lookup(c.NameAndTypeIndex))
}

func ReadCONSTANT_Fieldref_info(r io.Reader) *CONSTANT_Fieldref_info {
	fr := &CONSTANT_Fieldref_info{}
	fr.Tag = CONSTANT_Fieldref
	fr.ReadRefBaseIndexes(r)
	return fr
}

type CONSTANT_Methodref_info struct{ RefBase }

func ReadCONSTANT_Methodref_info(r io.Reader) *CONSTANT_Methodref_info {
	mr := &CONSTANT_Methodref_info{}
	mr.Tag = CONSTANT_Methodref
	mr.ReadRefBaseIndexes(r)
	return mr
}

type CONSTANT_InterfaceMethodref_info struct{ RefBase }

func ReadCONSTANT_InterfaceMethodref_info(r io.Reader) *CONSTANT_InterfaceMethodref_info {
	imr := &CONSTANT_InterfaceMethodref_info{}
	imr.Tag = CONSTANT_InterfaceMethodref
	imr.ReadRefBaseIndexes(r)
	return imr
}

