package javaclassparser

import (
	"fmt"
	"io"

	"github.com/mlctrez/javaclassparser/aflag"
	"github.com/mlctrez/javaclassparser/attribute"
	"github.com/mlctrez/javaclassparser/cpool"
	"github.com/mlctrez/javaclassparser/ioutil"
)

// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html

func failErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (jcp *ClassParser) readID(r io.Reader) error {
	failErr(ioutil.ReadUint32(r, &jcp.magic))
	if 0xCAFEBABE != jcp.magic {
		return fmt.Errorf("incorrect magic header %X", jcp.magic)
	}
	return nil
}

func (jcp *ClassParser) readMajorMinor(r io.Reader) {
	failErr(ioutil.ReadUint16(r, &jcp.minor))
	failErr(ioutil.ReadUint16(r, &jcp.major))
}

func (jcp *ClassParser) readClassInfo(r io.Reader) {
	jcp.accessFlags = aflag.ReadClassAccessFlags(r)

	failErr(ioutil.ReadUint16(r, &jcp.classNameIndex))
	failErr(ioutil.ReadUint16(r, &jcp.superClassNameIndex))

	//err := (&jcp.accessFlags).ValidateClass()
	//failErr(err)
}

func (jcp *ClassParser) readInterfaces(r io.Reader) {
	var interfaceCount uint16
	failErr(ioutil.ReadUint16(r, &interfaceCount))
	jcp.interfaces = make([]uint16, interfaceCount)
	var idx uint16

	for idx = 0; idx < interfaceCount; idx++ {
		// TODO: optimize this
		failErr(ioutil.ReadUint16(r, &jcp.interfaces[idx]))
	}
	return
}

func (jcp *ClassParser) readFields(r cpool.PoolReader) {
	var count uint16
	failErr(ioutil.ReadUint16(r, &count))

	jcp.fields = make([]*attribute.FieldInfo, count)
	var i uint16
	for i = 0; i < count; i++ {
		jcp.fields[i] = attribute.ReadFieldInfo(r)
	}
	return
}

func (jcp *ClassParser) readMethods(r cpool.PoolReader) {
	var count uint16
	failErr(ioutil.ReadUint16(r, &count))

	jcp.methods = make([]*attribute.MethodInfo, count)

	var i uint16
	for i = 0; i < count; i++ {
		jcp.methods[i] = attribute.ReadMethodInfo(r)
	}
	return
}

func (jcp *ClassParser) readAttributes(r cpool.PoolReader) {
	var count uint16
	failErr(ioutil.ReadUint16(r, &count))
	jcp.attributes = make([]interface{}, count)

	var i uint16
	for i = 0; i < count; i++ {
		jcp.attributes[i] = attribute.ReadAttributeInfo(r)
	}
	return
}

type ClassParser struct {
	Path  string
	Class string

	magic               uint32
	major               uint16
	minor               uint16
	constantPool        cpool.ConstantPool
	accessFlags         aflag.ClassAccessFlags
	classNameIndex      uint16
	superClassNameIndex uint16
	interfaces          []uint16
	fields              []*attribute.FieldInfo
	methods             []*attribute.MethodInfo
	attributes          []interface{}
}

func (jcp *ClassParser) Lookup(index uint16) interface{} {
	return jcp.constantPool.Lookup(index)
}

func (jcp *ClassParser) Visit(f func(interface{}))  {
	jcp.constantPool.Visit(f)
}


func (jcp *ClassParser) Parse(r io.Reader) (err error) {
	failErr(jcp.readID(r))
	jcp.readMajorMinor(r)
	if jcp.constantPool, err = cpool.Read(r); err != nil {
		return
	}

	pr := cpool.PoolReader{Reader: r, ConstantPool: jcp}
	jcp.readClassInfo(pr)
	jcp.readInterfaces(pr)
	jcp.readFields(pr)
	jcp.readMethods(pr)
	jcp.readAttributes(pr)
	return nil
}

func (jcp *ClassParser) SummarizeOut() {
	for i, f := range jcp.methods {
		_ = i
		for i := 0; i < len(f.Attributes); i++ {
			attr := f.Attributes[i]
			if code, ok := attr.(*attribute.CodeAttribute); ok {
				methodName := jcp.Lookup(f.NameIndex)
				className := jcp.Lookup(jcp.classNameIndex)
				//fmt.Println("class",reflect.TypeOf(className))
				//fmt.Println("method",reflect.TypeOf(methodName))

				if len(code.Code) > 400 {
					fmt.Printf("%05d %s.%s\n", len(code.Code), className, methodName)
				}
			}
		}
	}
}

func (jcp *ClassParser) SuperClass() string {
	return fmt.Sprintf("%s", jcp.Lookup(jcp.superClassNameIndex))
}

func (jcp *ClassParser) ClassName() string {
	return fmt.Sprintf("%s", jcp.Lookup(jcp.classNameIndex))
}

func (jcp *ClassParser) DebugOut() {

	fmt.Println("**************", jcp.Path, jcp.Class)

	jcp.constantPool.DebugOut()

	fmt.Print("access, className, superClass = ")
	fmt.Println(jcp.accessFlags, jcp.Lookup(jcp.classNameIndex), jcp.Lookup(jcp.superClassNameIndex))

	for i, itf := range jcp.interfaces {
		fmt.Printf("interface %3d %s\n", i, jcp.Lookup(itf))
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
					fmt.Printf(" Code %04X %s\n", instruction.Offset, instruction.StringWithIndex(jcp))
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
