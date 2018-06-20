package attribute

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/mlctrez/javaclassparser/aflag"
	"github.com/mlctrez/javaclassparser/bytecode"
	"github.com/mlctrez/javaclassparser/cpool"
	"github.com/mlctrez/javaclassparser/ioutil"
)

func failErr(err error) {
	if err != nil {
		panic(err)
	}
}

type AttributeBase struct {
	Pool cpool.ConstantPool
}

type FieldInfo struct {
	AttributeBase
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

func ReadFieldInfo(r cpool.PoolReader) *FieldInfo {
	fi := &FieldInfo{}
	fi.Pool = r.ConstantPool

	fi.AccessFlags = aflag.ReadFieldAccessFlags(r)

	failErr(ioutil.ReadUint16(r, &fi.NameIndex))
	failErr(ioutil.ReadUint16(r, &fi.DescriptorIndex))

	var count uint16
	failErr(ioutil.ReadUint16(r, &count))
	fi.Attributes = make([]interface{}, count)
	var i uint16
	for i = 0; i < count; i++ {
		fi.Attributes[i] = ReadAttributeInfo(r)
	}
	return fi
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

func ReadCodeAttributeExceptionTable(r io.Reader) *Code_attribute_exception_table {
	et := &Code_attribute_exception_table{}
	failErr(ioutil.ReadUint16(r, &et.StartPc))
	failErr(ioutil.ReadUint16(r, &et.EndPc))
	failErr(ioutil.ReadUint16(r, &et.HandlerPc))
	failErr(ioutil.ReadUint16(r, &et.CatchType))
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

func ReadCodeAttribute(r cpool.PoolReader) *Code_attribute {
	c := &Code_attribute{}
	failErr(ioutil.ReadUint16(r, &c.MaxStack))
	failErr(ioutil.ReadUint16(r, &c.MaxLocals))

	codes, err := bytecode.Read(r)
	failErr(err)
	c.Code = codes

	var exceptionTableLength uint16
	failErr(ioutil.ReadUint16(r, &exceptionTableLength))
	c.ExceptionTable = make([]*Code_attribute_exception_table, exceptionTableLength)
	var i uint16
	for i = 0; i < exceptionTableLength; i++ {
		c.ExceptionTable[i] = ReadCodeAttributeExceptionTable(r)
	}
	var attributesLength uint16
	c.Attributes = make([]interface{}, attributesLength)
	for i = 0; i < attributesLength; i++ {
		c.Attributes[i] = ReadAttributeInfo(r)
	}
	return c
}

func ReadAttributeInfo(r cpool.PoolReader) interface{} {

	var attributeNameIndex uint16
	var attributeLength uint32
	failErr(ioutil.ReadUint16(r, &attributeNameIndex))
	failErr(ioutil.ReadUint32(r, &attributeLength))

	// TODO: optimize for common lengths
	info := make([]uint8, attributeLength)
	failErr(binary.Read(r, binary.BigEndian, &info))
	lr := bytes.NewReader(info)

	cp := r.ConstantPool

	var attributeName *cpool.ConstantUtf8Info
	var ok bool
	if attributeName, ok = cp.Lookup(attributeNameIndex).(*cpool.ConstantUtf8Info); !ok {
		failErr(fmt.Errorf("invalid attributeNameIndex %X", attributeNameIndex))
	}

	switch attributeName.Value {
	case "SourceFile":
		s := &SourceFile_attribute{}
		s.Pool = cp
		failErr(ioutil.ReadUint16(lr, &s.SourceFileIndex))
		return s
	case "ConstantValue":
		c := &ConstantValue_attribute{}
		c.Pool = cp
		failErr(ioutil.ReadUint16(lr, &c.ConstantValueIndex))
		return c
	case "Code":
		lpr := cpool.PoolReader{Reader: lr, ConstantPool: cp}
		return ReadCodeAttribute(lpr)
	case "Exceptions":
		ea := &Exceptions_attribute{}
		ea.Pool = cp
		var numExceptions uint16
		failErr(ioutil.ReadUint16(lr, &numExceptions))
		ea.Exceptions = make([]uint16, numExceptions)
		// TODO: optimize
		failErr(binary.Read(lr, binary.BigEndian, &ea.Exceptions))
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

	case "ScalaInlineInfo":
	case "Scala":
	case "ScalaSig":

	case "org.aspectj.weaver.PointcutDeclaration":
	case "org.aspectj.weaver.MethodDeclarationLineNumber":
	case "org.aspectj.weaver.AjSynthetic":
	case "org.aspectj.weaver.WeaverVersion":
	case "org.aspectj.weaver.WeaverState":
	case "org.aspectj.weaver.EffectiveSignature":

	default:
		//if !strings.HasPrefix(attributeName.Value, "org.aspectj") {
		fmt.Println("unhandled attributeName", attributeName.Value)
		//}
	}
	return nil
}

func TypeName(i interface{}) (name string) {
	name = reflect.TypeOf(i).String()
	np := strings.Split(name, ".")
	return np[len(np)-1]
}

type MethodInfo struct {
	AttributeBase
	AccessFlags     aflag.MethodAccessFlags
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []interface{}
}

func (c *MethodInfo) String() string {
	return fmt.Sprintf("%s %s %s", c.AccessFlags, c.Pool.Lookup(c.NameIndex), c.Pool.Lookup(c.DescriptorIndex))
}

func ReadMethodInfo(r cpool.PoolReader) *MethodInfo {

	mi := &MethodInfo{}
	mi.Pool = r.ConstantPool
	mi.AccessFlags = aflag.ReadMethodAccessFlags(r)
	failErr(ioutil.ReadUint16(r, &mi.NameIndex))
	failErr(ioutil.ReadUint16(r, &mi.DescriptorIndex))

	var count uint16
	failErr(ioutil.ReadUint16(r, &count))

	mi.Attributes = make([]interface{}, count)

	var i uint16
	for i = 0; i < count; i++ {
		mi.Attributes[i] = ReadAttributeInfo(r)
	}
	return mi
}
