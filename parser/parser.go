package parser

import (
	"fmt"
	"io"

	"github.com/mlctrez/javaclassparser/aflag"
	"github.com/mlctrez/javaclassparser/attribute"
	"github.com/mlctrez/javaclassparser/cpool"
	"github.com/mlctrez/javaclassparser/ioutil"
)

// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html

// TODO: re-work error propagation

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
	Path  string
	Class string

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
}

func (jcp *Class) Visit(f func(interface{})) {
	jcp.pool.Visit(f)
}

func (jcp *Class) SummarizeOut() {
	for i, f := range jcp.methods {
		_ = i
		for i := 0; i < len(f.Attributes); i++ {
			attr := f.Attributes[i]
			if code, ok := attr.(*attribute.CodeAttribute); ok {
				methodName := jcp.pool.Lookup(f.NameIndex)
				className := jcp.pool.Lookup(jcp.classNameIndex)
				//fmt.Println("class",reflect.TypeOf(className))
				//fmt.Println("method",reflect.TypeOf(methodName))

				if len(code.Code) > 400 {
					fmt.Printf("%05d %s.%s\n", len(code.Code), className, methodName)
				}
			}
		}
	}
}

func (jcp *Class) ClassName() string {
	return fmt.Sprintf("%s", jcp.pool.Lookup(jcp.classNameIndex))
}

func (jcp *Class) DebugOut() {

	fmt.Println("**************", jcp.Path, jcp.Class)

	jcp.pool.DebugOut()

	fmt.Print("access, className, superClass = ")
	fmt.Println(jcp.accessFlags, jcp.pool.Lookup(jcp.classNameIndex), jcp.pool.Lookup(jcp.superClassNameIndex))

	for i, itf := range jcp.interfaces {
		fmt.Printf("interface %3d %s\n", i, jcp.pool.Lookup(itf))
	}

	fmt.Println("*** class fields")

	for i, f := range jcp.fields {
		fmt.Println(i, f)
	}

	fmt.Println("*** class methods")

	for i, f := range jcp.methods {
		fmt.Println(i, f)
		for i := 0; i < len(f.Attributes); i++ {
			attr := f.Attributes[i]
			if code, ok := attr.(*attribute.CodeAttribute); ok {
				for j := 0; j < len(code.Code); j++ {
					instruction := code.Code[j]
					fmt.Printf(" Code %04X %s\n", instruction.Offset, instruction.StringWithIndex(jcp.pool))
				}
				for j := 0; j < len(code.ExceptionTable); j++ {
					fmt.Printf(" ExceptionTable %+v\n", code.ExceptionTable[j])
				}
				for j := 0; j < len(code.Attributes); j++ {
					fmt.Printf(" Attributes %+v\n", code.Attributes[j])
				}
			} else {
				fmt.Printf(" attribute %d %+v\n", i, f.Attributes[i])
			}
		}
	}

	fmt.Println("*** class attributes")

	for i, f := range jcp.attributes {
		fmt.Println(i, f)
	}

}
