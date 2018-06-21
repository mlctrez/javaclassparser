package cpool

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/mlctrez/javaclassparser/ioutil"
)

// ConstantPool breaks cyclic dependencies between packages
// TODO: this may not be needed anymore after recent refactoring
type ConstantPool interface {
	Lookup(index uint16) interface{}
	Visit(func(interface{}))
	DebugOut()
}

// PoolReader combines the current reader and a reference to the constant pool
type PoolReader struct {
	io.Reader
	ConstantPool
}

func Read(outerReader io.Reader) (cp ConstantPool, err error) {

	var cpLen uint16
	if err = ioutil.ReadUint16(outerReader, &cpLen); err != nil {
		return
	}

	pi := &poolImpl{}
	cp = pi

	pi.constantPool = make([]interface{}, cpLen)

	r := PoolReader{Reader: outerReader, ConstantPool: cp}

	ts := []byte{0}
	for i := 1; i < int(cpLen); i++ {
		if _, err := io.ReadFull(r, ts); err != nil {
			return nil, err
		}
		tag := ts[0]
		switch tag {
		case ConstantClass:
			pi.constantPool[i] = ReadConstantClassInfo(r)
		case ConstantFieldref:
			pi.constantPool[i] = ReadConstantFieldrefInfo(r)
		case ConstantMethodref:
			pi.constantPool[i] = ReadConstantMethodrefInfo(r)
		case ConstantInterfaceMethodref:
			pi.constantPool[i] = ReadConstantInterfaceMethodrefInfo(r)
		case ConstantString:
			pi.constantPool[i] = ReadConstantStringInfo(r)
		case ConstantInteger:
			pi.constantPool[i] = ReadConstantIntegerInfo(r)
		case ConstantFloat:
			pi.constantPool[i] = ReadConstantFloatInfo(r)
		case ConstantLong:
			pi.constantPool[i] = ReadConstantLongInfo(r)
		case ConstantDouble:
			pi.constantPool[i] = ReadConstantDoubleInfo(r)
		case ConstantNameAndType:
			pi.constantPool[i] = ReadConstantNameAndTypeInfo(r)
		case ConstantUtf8:
			pi.constantPool[i] = ReadConstantUtf8Info(r)
		case ConstantMethodHandle:
			pi.constantPool[i] = ReadConstantMethodHandleInfo(r)
		case ConstantMethodType:
			pi.constantPool[i] = ReadConstantMethodTypeInfo(r)
		case ConstantInvokeDynamic:
			pi.constantPool[i] = ReadConstantInvokeDynamicInfo(r)
		default:
			panic("unknown tag in constantPool : " + strconv.Itoa(int(tag)))
		}

		// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
		/*
			All 8-byte constants take up two entries in the constant_pool table of the class file.
			If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
			table at index n, then the next usable item in the pool is located at index n+2.
			The constant_pool index n+1 must be valid but is considered unusable.
		*/
		switch tag {
		case ConstantDouble, ConstantLong:
			i++
		}
	}

	return pi, nil
}

type poolImpl struct {
	constantPool []interface{}
}

func (cp *poolImpl) Lookup(index uint16) interface{} {
	return cp.constantPool[index]
}

func (cp *poolImpl) Visit(f func(interface{})) {
	for _, c := range cp.constantPool {
		f(c)
	}
}

func (cp *poolImpl) DebugOut() {
	for i, pe := range cp.constantPool {
		// skip first, long, and double entries that are nil
		if pe == nil {
			continue
		}
		fmt.Printf("pool %04X %-30s %s\n", i, TypeName(pe), pe)
	}
}

// TODO: find a better way to represent this
func TypeName(i interface{}) (name string) {
	name = reflect.TypeOf(i).String()
	np := strings.Split(name, ".")
	return np[len(np)-1]
}
