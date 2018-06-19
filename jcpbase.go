package javaclassparser

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/mlctrez/javaclassparser/bytecode"
	"github.com/mlctrez/javaclassparser/cpool"
)

// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html

func failErr(err error) {
	if err != nil {
		panic(err)
	}
}

func read(r io.Reader, data interface{}) {
	failErr(binary.Read(r, binary.BigEndian, data))
}

//type ConstBase struct {
//	Pool cpool.ConstantPool
//	Tag  uint8
//}
//
//type RefBase struct {
//	ConstBase
//	ClassIndex       uint16
//	NameAndTypeIndex uint16
//}
//
//func (rb *RefBase) ReadRefBaseIndexes(r io.Reader) {
//	read(r, &rb.ClassIndex)
//	read(r, &rb.NameAndTypeIndex)
//}
//
//type CONSTANT_Class_info struct {
//	ConstBase
//	NameIndex uint16
//}
//
//func (c *CONSTANT_Class_info) String() string {
//	return fmt.Sprintf("%s", c.Pool.Lookup(c.NameIndex))
//}
//
//func (jcp *ClassParser) ReadCONSTANT_Class_info(r io.Reader) *CONSTANT_Class_info {
//	c := &CONSTANT_Class_info{}
//	c.Tag = CONSTANT_Class
//	c.Pool = jcp
//	read(r, &c.NameIndex)
//	return c
//}
//
//type CONSTANT_Fieldref_info struct{ RefBase }
//
//func (c *CONSTANT_Fieldref_info) String() string {
//	return fmt.Sprintf("%s %s", c.Pool.Lookup(c.ClassIndex), c.Pool.Lookup(c.NameAndTypeIndex))
//}
//
//func (jcp *ClassParser) ReadCONSTANT_Fieldref_info(r io.Reader) *CONSTANT_Fieldref_info {
//	fr := &CONSTANT_Fieldref_info{}
//	fr.Pool = jcp
//	fr.Tag = CONSTANT_Fieldref
//	fr.ReadRefBaseIndexes(r)
//	return fr
//}
//
//type CONSTANT_Methodref_info struct{ RefBase }
//
//func (jcp *ClassParser) ReadCONSTANT_Methodref_info(r io.Reader) *CONSTANT_Methodref_info {
//	mr := &CONSTANT_Methodref_info{}
//	mr.Pool = jcp
//	mr.Tag = CONSTANT_Methodref
//	mr.ReadRefBaseIndexes(r)
//	return mr
//}
//
//type CONSTANT_InterfaceMethodref_info struct{ RefBase }
//
//func (jcp *ClassParser) ReadCONSTANT_InterfaceMethodref_info(r io.Reader) *CONSTANT_InterfaceMethodref_info {
//	imr := &CONSTANT_InterfaceMethodref_info{}
//	imr.Pool = jcp
//	imr.Tag = CONSTANT_InterfaceMethodref
//	imr.ReadRefBaseIndexes(r)
//	return imr
//}
//
//type CONSTANT_String_info struct {
//	ConstBase
//	StringIndex uint16
//}
//
//func (c *CONSTANT_String_info) String() string {
//	return fmt.Sprintf("%s", c.Pool.Lookup(c.StringIndex))
//}
//
//func (jcp *ClassParser) ReadCONSTANT_String_info(r io.Reader) *CONSTANT_String_info {
//	cs := &CONSTANT_String_info{}
//	cs.Pool = jcp
//	cs.Tag = CONSTANT_String
//	read(r, &cs.StringIndex)
//	return cs
//}
//
//type CONSTANT_Integer_info struct {
//	ConstBase
//	Value int32
//}
//
//func (c *CONSTANT_Integer_info) String() string {
//	return fmt.Sprintf("%d", c.Value)
//}
//
//func (jcp *ClassParser) ReadCONSTANT_Integer_info(r io.Reader) *CONSTANT_Integer_info {
//	ci := &CONSTANT_Integer_info{}
//	ci.Pool = jcp
//	ci.Tag = CONSTANT_Integer
//	read(r, &ci.Value)
//	return ci
//}
//
//type CONSTANT_Float_info struct {
//	ConstBase
//	Value float32
//}
//
//func (c *CONSTANT_Float_info) String() string {
//	return fmt.Sprintf("%f", c.Value)
//}
//
//func (jcp *ClassParser) ReadCONSTANT_Float_info(r io.Reader) *CONSTANT_Float_info {
//	cf := &CONSTANT_Float_info{}
//	cf.Pool = jcp
//	cf.Tag = CONSTANT_Float
//	var floatBits uint32
//	read(r, &floatBits)
//	cf.Value = math.Float32frombits(floatBits)
//	return cf
//}
//
//type CONSTANT_Long_info struct {
//	ConstBase
//	Value int64
//}
//
//func (c *CONSTANT_Long_info) String() string {
//	return fmt.Sprintf("%d", c.Value)
//}
//
//func (jcp *ClassParser) ReadCONSTANT_Long_info(r io.Reader) *CONSTANT_Long_info {
//	cl := &CONSTANT_Long_info{}
//	cl.Pool = jcp
//	cl.Tag = CONSTANT_Long
//	read(r, &cl.Value)
//	return cl
//}
//
//type CONSTANT_Double_info struct {
//	ConstBase
//	Value float64
//}
//
//func (c *CONSTANT_Double_info) String() string {
//	return fmt.Sprintf("%f", c.Value)
//}
//
//func (jcp *ClassParser) ReadCONSTANT_Double_info(r io.Reader) *CONSTANT_Double_info {
//	cd := &CONSTANT_Double_info{}
//	cd.Pool = jcp
//	cd.Tag = CONSTANT_Double
//	read(r, &cd.Value)
//	return cd
//}
//
//type CONSTANT_NameAndType_info struct {
//	ConstBase
//	NameIndex       uint16
//	DescriptorIndex uint16
//}
//
//func (c *CONSTANT_NameAndType_info) String() string {
//	return fmt.Sprintf("%s %s", c.Pool.Lookup(c.NameIndex), c.Pool.Lookup(c.DescriptorIndex))
//}
//
//func (jcp *ClassParser) ReadCONSTANT_NameAndType_info(r io.Reader) *CONSTANT_NameAndType_info {
//	nat := &CONSTANT_NameAndType_info{}
//	nat.Pool = jcp
//	nat.Tag = CONSTANT_NameAndType
//	read(r, &nat.NameIndex)
//	read(r, &nat.DescriptorIndex)
//	return nat
//}
//
//type CONSTANT_Utf8_info struct {
//	ConstBase
//	Value string
//}
//
//func (c *CONSTANT_Utf8_info) String() string {
//	// TODO: this was %q but changed to %s
//	return fmt.Sprintf("%s", c.Value)
//}
//
//func (jcp *ClassParser) ReadCONSTANT_Utf8_info(r io.Reader) *CONSTANT_Utf8_info {
//	u := &CONSTANT_Utf8_info{}
//	u.Pool = jcp
//	u.Tag = CONSTANT_Utf8
//	var length uint16
//	read(r, &length)
//	buff := make([]uint8, length)
//	read, err := r.Read(buff)
//	if err != nil {
//		fmt.Println(err)
//	}
//	failErr(err)
//	if length != uint16(read) {
//		failErr(fmt.Errorf("incorrect length, expected %d but got %d", length, read))
//	}
//	u.Value = string(buff)
//	return u
//}
//
//type CONSTANT_MethodHandle_info struct {
//	ConstBase
//	ReferenceKind  uint8
//	ReferenceIndex uint16
//}
//
//func (jcp *ClassParser) ReadCONSTANT_MethodHandle_info(r io.Reader) *CONSTANT_MethodHandle_info {
//	mh := &CONSTANT_MethodHandle_info{}
//	mh.Pool = jcp
//	mh.Tag = CONSTANT_MethodHandle
//	read(r, &mh.ReferenceKind)
//	read(r, &mh.ReferenceIndex)
//	return mh
//}
//
//type CONSTANT_MethodType_info struct {
//	ConstBase
//	DescriptorIndex uint16
//}
//
//func (jcp *ClassParser) ReadCONSTANT_MethodType_info(r io.Reader) *CONSTANT_MethodType_info {
//	mt := &CONSTANT_MethodType_info{}
//	mt.Pool = jcp
//	mt.Tag = CONSTANT_MethodType
//	read(r, &mt.DescriptorIndex)
//	return mt
//}
//
//type CONSTANT_InvokeDynamic_info struct {
//	ConstBase
//	BoostrapMethodAttrIndex uint16
//	NameAndTypeIndex        uint16
//}
//
//func (jcp *ClassParser) ReadCONSTANT_InvokeDynamic_info(r io.Reader) *CONSTANT_InvokeDynamic_info {
//	cid := &CONSTANT_InvokeDynamic_info{}
//	cid.Pool = jcp
//	cid.Tag = CONSTANT_InvokeDynamic
//	read(r, &cid.BoostrapMethodAttrIndex)
//	read(r, &cid.NameAndTypeIndex)
//	return cid
//}

func (jcp *ClassParser) readID(r io.Reader) error {
	read(r, &jcp.magic)
	if 0xCAFEBABE != jcp.magic {
		return errors.New("incorrect magic header")
	}
	return nil
}

func (jcp *ClassParser) readMajorMinor(r io.Reader) {
	read(r, &jcp.minor)
	read(r, &jcp.major)
}
//
//func (jcp *ClassParser) readConstantPool(r io.Reader) {
//	var cpLen uint16
//	read(r, &cpLen)
//
//	jcp.constantPool = make([]interface{}, cpLen)
//
//	for i := 1; i < int(cpLen); i++ {
//		var tag uint8
//		read(r, &tag)
//
//		switch tag {
//		case CONSTANT_Class:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_Class_info(r)
//		case CONSTANT_Fieldref:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_Fieldref_info(r)
//		case CONSTANT_Methodref:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_Methodref_info(r)
//		case CONSTANT_InterfaceMethodref:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_InterfaceMethodref_info(r)
//		case CONSTANT_String:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_String_info(r)
//		case CONSTANT_Integer:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_Integer_info(r)
//		case CONSTANT_Float:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_Float_info(r)
//		case CONSTANT_Long:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_Long_info(r)
//		case CONSTANT_Double:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_Double_info(r)
//		case CONSTANT_NameAndType:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_NameAndType_info(r)
//		case CONSTANT_Utf8:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_Utf8_info(r)
//		case CONSTANT_MethodHandle:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_MethodHandle_info(r)
//		case CONSTANT_MethodType:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_MethodType_info(r)
//		case CONSTANT_InvokeDynamic:
//			jcp.constantPool[i] = jcp.ReadCONSTANT_InvokeDynamic_info(r)
//		default:
//			panic("unknown tag in constantPool : " + strconv.Itoa(int(tag)))
//		}
//
//		// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
//		/*
//		All 8-byte constants take up two entries in the constant_pool table of the class file.
//		If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
//		table at index n, then the next usable item in the pool is located at index n+2.
//		The constant_pool index n+1 must be valid but is considered unusable.
//		 */
//		switch tag {
//		case CONSTANT_Double, CONSTANT_Long:
//			i++
//		}
//	}
//}

func (jcp *ClassParser) readClassInfo(r io.Reader) {
	read(r, &jcp.accessFlags)
	read(r, &jcp.classNameIndex)
	read(r, &jcp.superClassNameIndex)

	err := (&jcp.accessFlags).ValidateClass()
	failErr(err)
	return
}

func (jcp *ClassParser) readInterfaces(r io.Reader) {
	var interfaceCount uint16
	read(r, &interfaceCount)
	jcp.interfaces = make([]uint16, interfaceCount)
	var idx uint16

	for idx = 0; idx < interfaceCount; idx++ {
		read(r, &jcp.interfaces[idx])
	}
	return
}

type AttributeBase struct {
	Pool cpool.ConstantPool
}

type field_info struct {
	AttributeBase
	AccessFlags     AccessFlag
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []interface{}
}

func (c *field_info) String() string {
	af := c.AccessFlags
	name := c.Pool.Lookup(c.NameIndex)
	descriptor := c.Pool.Lookup(c.DescriptorIndex)
	attributes := c.Attributes
	return fmt.Sprintf("%s %s %s %s", af, name, descriptor, attributes)
}

type SourceFile_attribute struct {
	AttributeBase
	AttributeNameIndex uint16
	AttributeLength    uint64
	SourceFileIndex    uint16
}

func (s *SourceFile_attribute) String() string {
	return fmt.Sprintf("%s %s", TypeName(s), s.Pool.Lookup(s.SourceFileIndex))
}

type ConstantValue_attribute struct {
	AttributeBase
	ConstantValueIndex uint16
}

func (s *ConstantValue_attribute) String() string {
	return fmt.Sprintf("%s %v", TypeName(s), s.Pool.Lookup(s.ConstantValueIndex))
}

type Code_attribute struct {
	AttributeBase
	MaxStack       uint16
	MaxLocals      uint16
	Code           []*bytecode.ByteCode
	ExceptionTable []*Code_attribute_exception_table
	Attributes     []interface{}
}

type Code_attribute_exception_table struct {
	AttributeBase
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func (c *Code_attribute_exception_table) String() string {
	return fmt.Sprintf("Start %X End %X Handler %X CatchType %X", c.StartPc, c.EndPc, c.HandlerPc, c.CatchType)
}

func (jcp *ClassParser) ReadCodeAttributeExceptionTable(r io.Reader) *Code_attribute_exception_table {
	et := &Code_attribute_exception_table{}
	et.Pool = jcp
	read(r, &et.StartPc)
	read(r, &et.EndPc)
	read(r, &et.HandlerPc)
	read(r, &et.CatchType)
	return et
}

type Exceptions_attribute struct {
	AttributeBase
	Exceptions []uint16
}

func (s *Exceptions_attribute) String() string {
	el := make([]interface{}, len(s.Exceptions))
	for i, e := range s.Exceptions {
		el[i] = s.Pool.Lookup(e)
	}
	return fmt.Sprintf("%v %v", TypeName(s), el)
}

func (jcp *ClassParser) ReadCodeAttribute(r io.Reader) *Code_attribute {
	c := &Code_attribute{}
	c.Pool = jcp
	read(r, &c.MaxStack)
	read(r, &c.MaxLocals)

	codes, err := bytecode.Read(r)
	failErr(err)
	c.Code = codes

	var exceptionTableLength uint16
	read(r, &exceptionTableLength)
	c.ExceptionTable = make([]*Code_attribute_exception_table, exceptionTableLength)
	var i uint16
	for i = 0; i < exceptionTableLength; i++ {
		c.ExceptionTable[i] = jcp.ReadCodeAttributeExceptionTable(r)
	}
	var attributesLength uint16
	c.Attributes = make([]interface{}, attributesLength)
	for i = 0; i < attributesLength; i++ {
		c.Attributes[i] = jcp.ReadAttributeInfo(r)
	}
	return c
}

func (jcp *ClassParser) ReadAttributeInfo(r io.Reader) interface{} {

	var attributeNameIndex uint16
	var attributeLength uint32
	read(r, &attributeNameIndex)
	read(r, &attributeLength)
	info := make([]uint8, attributeLength)
	read(r, &info)

	lr := bytes.NewReader(info)

	var attributeName *cpool.CONSTANT_Utf8_info
	var ok bool
	if attributeName, ok = jcp.Lookup(attributeNameIndex).(*cpool.CONSTANT_Utf8_info); !ok {
		failErr(fmt.Errorf("invalid attributeNameIndex %X", attributeNameIndex))
	}

	switch attributeName.Value {
	case "SourceFile":
		s := &SourceFile_attribute{}
		s.Pool = jcp
		read(lr, &s.SourceFileIndex)
		return s
	case "ConstantValue":
		c := &ConstantValue_attribute{}
		c.Pool = jcp
		read(lr, &c.ConstantValueIndex)
		return c
	case "Code":
		return jcp.ReadCodeAttribute(bytes.NewReader(info))
	case "Exceptions":
		ea := &Exceptions_attribute{}
		ea.Pool = jcp
		var numExceptions uint16
		read(lr, &numExceptions)
		ea.Exceptions = make([]uint16, numExceptions)
		read(lr, &ea.Exceptions)
		return ea
	case "EnclosingMethod":
		// TODO: finish the remainder of these
	case "InnerClasses":
	case "BootstrapMethods":
	case "Signature":
	case "Synthetic":
	case "Deprecated":
	case "RuntimeVisibleAnnotations":
	case "AnnotationDefault":
	case "RuntimeInvisibleAnnotations":
	case "RuntimeInvisibleParameterAnnotations":
	case "RuntimeVisibleParameterAnnotations":
	case "SourceDebugExtension":
	case "Bridge":
	case "MethodParameters":

	default:
		if !strings.HasPrefix(attributeName.Value, "org.aspectj") {
			fmt.Println("unhandled attributeName", attributeName.Value)
		}
	}
	return nil
}

func (jcp *ClassParser) ReadFieldInfo(r io.Reader) *field_info {
	fi := &field_info{}
	fi.Pool = jcp
	read(r, &fi.AccessFlags)
	read(r, &fi.NameIndex)
	read(r, &fi.DescriptorIndex)

	var count uint16
	read(r, &count)
	fi.Attributes = make([]interface{}, count)
	var i uint16
	for i = 0; i < count; i++ {
		fi.Attributes[i] = jcp.ReadAttributeInfo(r)
	}
	return fi
}

func (jcp *ClassParser) readFields(r io.Reader) {
	var count uint16
	read(r, &count)

	jcp.fields = make([]*field_info, count)
	var i uint16
	for i = 0; i < count; i++ {
		jcp.fields[i] = jcp.ReadFieldInfo(r)
	}
	return
}

type MethodInfo struct {
	AttributeBase
	AccessFlags     AccessFlag
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []interface{}
}

func (c *MethodInfo) String() string {
	return fmt.Sprintf("%s %s %s", c.AccessFlags, c.Pool.Lookup(c.NameIndex), c.Pool.Lookup(c.DescriptorIndex))
}

func (jcp *ClassParser) ReadMethodInfo(r io.Reader) *MethodInfo {

	mi := &MethodInfo{}
	mi.Pool = jcp
	read(r, &mi.AccessFlags)
	read(r, &mi.NameIndex)
	read(r, &mi.DescriptorIndex)

	var count uint16
	read(r, &count)
	mi.Attributes = make([]interface{}, count)

	var i uint16
	for i = 0; i < count; i++ {
		mi.Attributes[i] = jcp.ReadAttributeInfo(r)
	}
	return mi
}

func (jcp *ClassParser) readMethods(r io.Reader) {
	var count uint16
	read(r, &count)
	jcp.methods = make([]*MethodInfo, count)

	var i uint16
	for i = 0; i < count; i++ {
		jcp.methods[i] = jcp.ReadMethodInfo(r)
	}
	return
}

func (jcp *ClassParser) readAttributes(r io.Reader) {
	var count uint16
	read(r, &count)
	jcp.attributes = make([]interface{}, count)

	var i uint16
	for i = 0; i < count; i++ {
		jcp.attributes[i] = jcp.ReadAttributeInfo(r)
	}
	return
}

func TypeName(i interface{}) (name string) {
	name = reflect.TypeOf(i).String()
	np := strings.Split(name, ".")
	return np[len(np)-1]
}

type ClassParser struct {
	magic               uint32
	major               uint16
	minor               uint16
	constantPool        cpool.ConstantPool
	accessFlags         AccessFlag
	classNameIndex      uint16
	superClassNameIndex uint16
	interfaces          []uint16
	fields              []*field_info
	methods             []*MethodInfo
	attributes          []interface{}
}

func (jcp *ClassParser) Lookup(index uint16) interface{} {
	return jcp.constantPool.Lookup(index)
}

func (jcp *ClassParser) Parse(r io.Reader) (err error) {
	failErr(jcp.readID(r))
	jcp.readMajorMinor(r)
	if jcp.constantPool,err = cpool.Read(r); err != nil {
		return
	}
	jcp.readClassInfo(r)
	jcp.readInterfaces(r)
	jcp.readFields(r)
	jcp.readMethods(r)
	jcp.readAttributes(r)
	return nil
}

func (jcp *ClassParser) SummarizeOut() {
	for i, f := range jcp.methods {
		_ = i
		for i := 0; i < len(f.Attributes); i++ {
			attr := f.Attributes[i]
			if code, ok := attr.(*Code_attribute); ok {
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

func (jcp *ClassParser) DebugOut() {

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
			if code, ok := attr.(*Code_attribute); ok {
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
