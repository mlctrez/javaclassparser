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

type baseAttribute struct {
	Pool cpool.ConstantPool
}

// SourceFileAttribute represents the original java source file name
// See https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.10
type SourceFileAttribute struct {
	baseAttribute
	AttributeNameIndex uint16
	AttributeLength    uint64
	SourceFileIndex    uint16
}

func (s *SourceFileAttribute) String() string {
	return fmt.Sprintf("%s %s", TypeName(s), s.Pool.Lookup(s.SourceFileIndex))
}

type ConstantValueAttribute struct {
	baseAttribute
	ConstantValueIndex uint16
}

func (s *ConstantValueAttribute) String() string {
	return fmt.Sprintf("%s %v", TypeName(s), s.Pool.Lookup(s.ConstantValueIndex))
}

type CodeAttribute struct {
	baseAttribute
	MaxStack       uint16
	MaxLocals      uint16
	Code           []*bytecode.ByteCode
	ExceptionTable []*CodeAttributeExceptionTable
	Attributes     []interface{}
}

type CodeAttributeExceptionTable struct {
	baseAttribute
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func (c *CodeAttributeExceptionTable) String() string {
	return fmt.Sprintf("Start %X End %X Handler %X CatchType %X", c.StartPc, c.EndPc, c.HandlerPc, c.CatchType)
}

func ReadCodeAttributeExceptionTable(r io.Reader) (et *CodeAttributeExceptionTable, err error) {
	et = &CodeAttributeExceptionTable{}
	if err = ioutil.ReadUint16(r, &et.StartPc); err != nil {
		return
	}
	if err = ioutil.ReadUint16(r, &et.EndPc); err != nil {
		return
	}
	if err = ioutil.ReadUint16(r, &et.HandlerPc); err != nil {
		return
	}
	err = ioutil.ReadUint16(r, &et.CatchType)
	return
}

type ExceptionsAttribute struct {
	baseAttribute
	Exceptions []uint16
}

func (s *ExceptionsAttribute) String() string {
	el := make([]interface{}, len(s.Exceptions))
	for i, e := range s.Exceptions {
		el[i] = s.Pool.Lookup(e)
	}
	return fmt.Sprintf("%v %v", TypeName(s), el)
}

func ReadExceptionsAttribute(r cpool.PoolReader) (ea *ExceptionsAttribute, err error) {
	ea = &ExceptionsAttribute{}
	ea.Pool = r.ConstantPool
	var numExceptions uint16
	if err = ioutil.ReadUint16(r, &numExceptions); err != nil {
		return
	}
	ea.Exceptions = make([]uint16, numExceptions)
	err = binary.Read(r, binary.BigEndian, &ea.Exceptions)
	return
}

func ReadCodeAttribute(r cpool.PoolReader) (ca *CodeAttribute, err error) {
	ca = &CodeAttribute{}
	if err = ioutil.ReadUint16(r, &ca.MaxStack); err != nil {
		return
	}
	if err = ioutil.ReadUint16(r, &ca.MaxLocals); err != nil {
		return
	}
	if ca.Code, err = bytecode.Read(r); err != nil {
		return
	}

	var exceptionTableLength uint16
	if err = ioutil.ReadUint16(r, &exceptionTableLength); err != nil {
		return
	}
	ca.ExceptionTable = make([]*CodeAttributeExceptionTable, exceptionTableLength)
	var i uint16
	for i = 0; i < exceptionTableLength; i++ {
		if ca.ExceptionTable[i], err = ReadCodeAttributeExceptionTable(r); err != nil {
			return
		}
	}
	var attributesLength uint16
	ca.Attributes = make([]interface{}, attributesLength)
	for i = 0; i < attributesLength; i++ {
		if ca.Attributes[i], err = ReadAttributeInfo(r); err != nil {
			return
		}
	}
	return
}

func ReadAttributeInfo(r cpool.PoolReader) (ai interface{}, err error) {

	var attributeNameIndex uint16
	var attributeLength uint32

	if err = ioutil.ReadUint16(r, &attributeNameIndex); err != nil {
		return
	}
	if err = ioutil.ReadUint32(r, &attributeLength); err != nil {
		return
	}

	info := make([]uint8, attributeLength)
	if _, err = io.ReadFull(r, info); err != nil {
		return
	}

	var attributeName *cpool.ConstantUtf8Info
	var ok bool
	if attributeName, ok = r.ConstantPool.Lookup(attributeNameIndex).(*cpool.ConstantUtf8Info); !ok {
		err = fmt.Errorf("invalid attributeNameIndex %X", attributeNameIndex)
		return
	}
	lr := cpool.PoolReader{Reader: bytes.NewReader(info), ConstantPool: r.ConstantPool}

	switch attributeName.Value {
	case "SourceFile":
		s := &SourceFileAttribute{}
		s.Pool = r.ConstantPool
		err = ioutil.ReadUint16(lr, &s.SourceFileIndex)
		return s, err
	case "ConstantValue":
		c := &ConstantValueAttribute{}
		c.Pool = r.ConstantPool
		err = ioutil.ReadUint16(lr, &c.ConstantValueIndex)
		return c, err
	case "Code":
		return ReadCodeAttribute(lr)
	case "Exceptions":
		return ReadExceptionsAttribute(lr)
	case "EnclosingMethod":
		// TODO: finish the remainder of these known ones from the spec
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
		// TODO: are these useful, or should they just be logged and ignored
	case "Scala":
	case "ScalaSig":
	case "org.aspectj.weaver.PointcutDeclaration":
	case "org.aspectj.weaver.MethodDeclarationLineNumber":
	case "org.aspectj.weaver.AjSynthetic":
	case "org.aspectj.weaver.WeaverVersion":
	case "org.aspectj.weaver.WeaverState":
	case "org.aspectj.weaver.Advice":
	case "org.aspectj.weaver.Aspect":
	case "org.aspectj.weaver.EffectiveSignature":
	case "org.aspectj.weaver.SourceContext":
	case "org.aspectj.weaver.Privileged":
	case "org.aspectj.weaver.TypeMunger":
	case "org.aspectj.weaver.Declare":

	default:
		fmt.Println("unhandled attributeName", attributeName.Value)
	}
	return
}

func TypeName(i interface{}) (name string) {
	name = reflect.TypeOf(i).String()
	np := strings.Split(name, ".")
	return np[len(np)-1]
}

type MethodInfo struct {
	baseAttribute
	AccessFlags     aflag.MethodAccessFlags
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []interface{}
}

func (c *MethodInfo) String() string {
	return fmt.Sprintf("%s %s %s", c.AccessFlags, c.Pool.Lookup(c.NameIndex), c.Pool.Lookup(c.DescriptorIndex))
}

func ReadMethodInfo(r cpool.PoolReader) (mi *MethodInfo, err error) {

	mi = &MethodInfo{}
	mi.Pool = r.ConstantPool

	if mi.AccessFlags, err = aflag.ReadMethodAccessFlags(r); err != nil {
		return
	}
	if err = ioutil.ReadUint16(r, &mi.NameIndex); err != nil {
		return
	}
	if err = ioutil.ReadUint16(r, &mi.DescriptorIndex); err != nil {
		return
	}

	var count uint16
	if err = ioutil.ReadUint16(r, &count); err != nil {
		return
	}

	mi.Attributes = make([]interface{}, count)

	var i uint16
	for i = 0; i < count; i++ {
		if mi.Attributes[i], err = ReadAttributeInfo(r); err != nil {
			return
		}
	}
	return
}
