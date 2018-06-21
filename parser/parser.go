package parser

import (
	"fmt"
	"io"

	"github.com/mlctrez/javaclassparser/aflag"
	"github.com/mlctrez/javaclassparser/attribute"
	"github.com/mlctrez/javaclassparser/cpool"
	"github.com/mlctrez/javaclassparser/ioutil"
)

// It's all at the following link
// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html

// New is the main entry point for parsing java byte code
// The reader is expected to point at a class file byte stream
func New(r io.Reader) (jcp *Class, err error) {

	jcp = &Class{}
	if err = jcp.readID(r); err != nil {
		return
	}
	if err = jcp.readMajorMinor(r); err != nil {
		return
	}
	if jcp.pool, err = cpool.Read(r); err != nil {
		return
	}

	poolReader := cpool.PoolReader{Reader: r, ConstantPool: jcp.pool}

	if err = jcp.readClassInfo(poolReader); err != nil {
		return
	}
	if err = jcp.readInterfaces(poolReader); err != nil {
		return
	}
	if err = jcp.readFields(poolReader); err != nil {
		return
	}
	if err = jcp.readMethods(poolReader); err != nil {
		return
	}
	if err = jcp.readAttributes(poolReader); err != nil {
		return
	}

	return
}

func (jcp *Class) readID(r io.Reader) (err error) {
	if err = ioutil.ReadUint32(r, &jcp.magic); err != nil {
		return
	}
	if 0xCAFEBABE != jcp.magic {
		return fmt.Errorf("incorrect magic header %X", jcp.magic)
	}
	return nil
}

func (jcp *Class) readMajorMinor(r io.Reader) (err error) {
	if err = ioutil.ReadUint16(r, &jcp.minor); err != nil {
		return
	}
	err = ioutil.ReadUint16(r, &jcp.major)
	return
}

func (jcp *Class) readClassInfo(r io.Reader) (err error) {
	jcp.accessFlags, err = aflag.ReadClassAccessFlags(r)

	if err = ioutil.ReadUint16(r, &jcp.classNameIndex); err != nil {
		return
	}
	if err = ioutil.ReadUint16(r, &jcp.superClassNameIndex); err != nil {
		return
	}

	// calculated once here to avoid multiple sprintfs when using it to sort
	jcp.Name = fmt.Sprintf("%s", jcp.pool.Lookup(jcp.classNameIndex))
	jcp.SuperClass = fmt.Sprintf("%s", jcp.pool.Lookup(jcp.superClassNameIndex))

	err = (&jcp.accessFlags).Validate()
	return
}

func (jcp *Class) readInterfaces(r io.Reader) (err error) {
	var interfaceCount uint16
	if err = ioutil.ReadUint16(r, &interfaceCount); err != nil {
		return
	}
	jcp.interfaces = make([]uint16, interfaceCount)

	var idx uint16
	for idx = 0; idx < interfaceCount; idx++ {
		if err = ioutil.ReadUint16(r, &jcp.interfaces[idx]); err != nil {
			return
		}
	}
	return
}

func (jcp *Class) readFields(r cpool.PoolReader) (err error) {
	var count uint16
	if err = ioutil.ReadUint16(r, &count); err != nil {
		return
	}
	jcp.fields = make([]*attribute.FieldInfo, count)
	var i uint16
	for i = 0; i < count; i++ {
		if jcp.fields[i], err = attribute.ReadFieldInfo(r); err != nil {
			return
		}
	}
	return
}

func (jcp *Class) readMethods(r cpool.PoolReader) (err error) {
	var count uint16
	if err = ioutil.ReadUint16(r, &count); err != nil {
		return
	}

	jcp.methods = make([]*attribute.MethodInfo, count)

	var i uint16
	for i = 0; i < count; i++ {
		if jcp.methods[i], err = attribute.ReadMethodInfo(r); err != nil {
			return
		}
	}
	return
}

func (jcp *Class) readAttributes(r cpool.PoolReader) (err error) {
	var count uint16
	if err = ioutil.ReadUint16(r, &count); err != nil {
		return
	}
	jcp.attributes = make([]interface{}, count)

	var i uint16
	for i = 0; i < count; i++ {
		if jcp.attributes[i], err = attribute.ReadAttributeInfo(r); err != nil {
			return
		}
	}
	return
}

type Class struct {
	magic               uint32
	major               uint16
	minor               uint16
	pool                cpool.ConstantPool
	accessFlags         aflag.ClassAccessFlags
	classNameIndex      uint16
	superClassNameIndex uint16
	interfaces          []uint16
	fields              []*attribute.FieldInfo
	methods             []*attribute.MethodInfo
	attributes          []interface{}

	// calculated fields
	Name       string
	SuperClass string
}

func (jcp *Class) Visit(f func(interface{})) {
	jcp.pool.Visit(f)
}
