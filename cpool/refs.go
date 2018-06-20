package cpool

import (
	"fmt"
	"io"

	"github.com/mlctrez/javaclassparser/ioutil"
)

type RefBase struct {
	ConstBase
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (rb *RefBase) ReadRefBaseIndexes(r io.Reader) {
	failErr(ioutil.ReadUint16(r, &rb.ClassIndex))
	failErr(ioutil.ReadUint16(r, &rb.NameAndTypeIndex))
}

type ConstantFieldrefInfo struct{ RefBase }

func (c *ConstantFieldrefInfo) String() string {
	return fmt.Sprintf("%s %s", c.Pool.Lookup(c.ClassIndex), c.Pool.Lookup(c.NameAndTypeIndex))
}

func ReadConstantFieldrefInfo(r PoolReader) *ConstantFieldrefInfo {
	fr := &ConstantFieldrefInfo{}
	fr.Pool = r.ConstantPool
	fr.Tag = ConstantFieldref
	fr.Type = "CONSTANT_Fieldref_info"
	fr.ReadRefBaseIndexes(r)
	return fr
}

type ConstantMethodrefInfo struct{ RefBase }

func (mr *ConstantMethodrefInfo) String() string {
	return fmt.Sprintf("%s %s", mr.Pool.Lookup(mr.ClassIndex), mr.Pool.Lookup(mr.NameAndTypeIndex))
}

func ReadConstantMethodrefInfo(r PoolReader) *ConstantMethodrefInfo {
	mr := &ConstantMethodrefInfo{}
	mr.Pool = r.ConstantPool
	mr.Tag = ConstantMethodref
	mr.Type = "CONSTANT_Methodref_info"
	mr.ReadRefBaseIndexes(r)
	return mr
}

type ConstantInterfaceMethodrefInfo struct{ RefBase }

func (imr *ConstantInterfaceMethodrefInfo) String() string {
	return fmt.Sprintf("%s %s", imr.Pool.Lookup(imr.ClassIndex), imr.Pool.Lookup(imr.NameAndTypeIndex))
}

func ReadConstantInterfaceMethodrefInfo(r PoolReader) *ConstantInterfaceMethodrefInfo {
	imr := &ConstantInterfaceMethodrefInfo{}
	imr.Pool = r.ConstantPool
	imr.Tag = ConstantInterfaceMethodref
	imr.Type = "CONSTANT_InterfaceMethodref_info"
	imr.ReadRefBaseIndexes(r)
	return imr
}
