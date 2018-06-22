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

func (rb *RefBase) ReadRefBaseIndexes(r io.Reader) (err error) {
	if err = ioutil.ReadUint16(r, &rb.ClassIndex); err != nil {
		return
	}
	err = ioutil.ReadUint16(r, &rb.NameAndTypeIndex)
	return
}

func (rb *RefBase) ClassName() string {
	return fmt.Sprintf("%s", rb.Pool.Lookup(rb.ClassIndex))
}

type ConstantFieldrefInfo struct{ RefBase }

func (c *ConstantFieldrefInfo) String() string {
	return fmt.Sprintf("%s %s", c.Pool.Lookup(c.ClassIndex), c.Pool.Lookup(c.NameAndTypeIndex))
}

func ReadConstantFieldrefInfo(r PoolReader) (fr *ConstantFieldrefInfo, err error) {
	fr = &ConstantFieldrefInfo{}
	fr.Pool = r.ConstantPool
	fr.Tag = ConstantFieldref
	fr.Type = "CONSTANT_Fieldref_info"
	err = fr.ReadRefBaseIndexes(r)
	return
}

type ConstantMethodrefInfo struct{ RefBase }

func (mr *ConstantMethodrefInfo) String() string {
	return fmt.Sprintf("%s %s", mr.Pool.Lookup(mr.ClassIndex), mr.Pool.Lookup(mr.NameAndTypeIndex))
}

func ReadConstantMethodrefInfo(r PoolReader) (mr *ConstantMethodrefInfo, err error) {
	mr = &ConstantMethodrefInfo{}
	mr.Pool = r.ConstantPool
	mr.Tag = ConstantMethodref
	mr.Type = "CONSTANT_Methodref_info"
	err = mr.ReadRefBaseIndexes(r)
	return
}

type ConstantInterfaceMethodrefInfo struct{ RefBase }

func (imr *ConstantInterfaceMethodrefInfo) String() string {
	return fmt.Sprintf("%s %s", imr.Pool.Lookup(imr.ClassIndex), imr.Pool.Lookup(imr.NameAndTypeIndex))
}

func ReadConstantInterfaceMethodrefInfo(r PoolReader) (imr *ConstantInterfaceMethodrefInfo, err error) {
	imr = &ConstantInterfaceMethodrefInfo{}
	imr.Pool = r.ConstantPool
	imr.Tag = ConstantInterfaceMethodref
	imr.Type = "CONSTANT_InterfaceMethodref_info"
	err = imr.ReadRefBaseIndexes(r)
	return
}
