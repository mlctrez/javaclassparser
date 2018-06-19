package cpool

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type ConstantPool interface {
	Lookup(index uint16) interface{}
	DebugOut()
}

func Read(r io.Reader) (ConstantPool, error) {
	bs := []byte{0, 0}
	if _, err := io.ReadFull(r, bs); err != nil {
		return nil, err
	}
	cpLen := binary.BigEndian.Uint16(bs)
	pi := &poolImpl{}
	pi.constantPool = make([]interface{}, cpLen)

	ts := []byte{0}
	for i := 1; i < int(cpLen); i++ {
		if _, err := io.ReadFull(r, ts); err != nil {
			return nil, err
		}
		tag := ts[0]
		switch tag {
		case CONSTANT_Class:
			pi.constantPool[i] = ReadCONSTANT_Class_info(r)
		case CONSTANT_Fieldref:
			pi.constantPool[i] = ReadCONSTANT_Fieldref_info(r)
		case CONSTANT_Methodref:
			pi.constantPool[i] = ReadCONSTANT_Methodref_info(r)
		case CONSTANT_InterfaceMethodref:
			pi.constantPool[i] = ReadCONSTANT_InterfaceMethodref_info(r)
		case CONSTANT_String:
			pi.constantPool[i] = ReadCONSTANT_String_info(r)
		case CONSTANT_Integer:
			pi.constantPool[i] = ReadCONSTANT_Integer_info(r)
		case CONSTANT_Float:
			pi.constantPool[i] = ReadCONSTANT_Float_info(r)
		case CONSTANT_Long:
			pi.constantPool[i] = ReadCONSTANT_Long_info(r)
		case CONSTANT_Double:
			pi.constantPool[i] = ReadCONSTANT_Double_info(r)
		case CONSTANT_NameAndType:
			pi.constantPool[i] = ReadCONSTANT_NameAndType_info(r)
		case CONSTANT_Utf8:
			pi.constantPool[i] = ReadCONSTANT_Utf8_info(r)
		case CONSTANT_MethodHandle:
			pi.constantPool[i] = ReadCONSTANT_MethodHandle_info(r)
		case CONSTANT_MethodType:
			pi.constantPool[i] = ReadCONSTANT_MethodType_info(r)
		case CONSTANT_InvokeDynamic:
			pi.constantPool[i] = ReadCONSTANT_InvokeDynamic_info(r)
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
		case CONSTANT_Double, CONSTANT_Long:
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

func (cp *poolImpl) DebugOut() {
	for i, pe := range cp.constantPool {
		// skip first, long, and double entries that are nil
		if pe == nil {
			continue
		}
		fmt.Printf("pool %04X %-25s %s\n", i, TypeName(pe), pe)
	}
}

// TODO: remove this in leiu of strings in the type
func TypeName(i interface{}) (name string) {
	name = reflect.TypeOf(i).String()
	np := strings.Split(name, ".")
	return np[len(np)-1]
}
