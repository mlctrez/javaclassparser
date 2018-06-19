package cpool

import (
	"io"
)

type CONSTANT_MethodHandle_info struct {
	ConstBase
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func ReadCONSTANT_MethodHandle_info(r io.Reader) *CONSTANT_MethodHandle_info {
	mh := &CONSTANT_MethodHandle_info{}
	mh.Tag = CONSTANT_MethodHandle
	failErr(readUint8(r, &mh.ReferenceKind))
	failErr(readUint16(r, &mh.ReferenceIndex))
	return mh
}

type CONSTANT_MethodType_info struct {
	ConstBase
	DescriptorIndex uint16
}

func ReadCONSTANT_MethodType_info(r io.Reader) *CONSTANT_MethodType_info {
	mt := &CONSTANT_MethodType_info{}
	mt.Tag = CONSTANT_MethodType
	failErr(readUint16(r,&mt.DescriptorIndex))
	return mt
}

type CONSTANT_InvokeDynamic_info struct {
	ConstBase
	BoostrapMethodAttrIndex uint16
	NameAndTypeIndex        uint16
}

func ReadCONSTANT_InvokeDynamic_info(r io.Reader) *CONSTANT_InvokeDynamic_info {
	cid := &CONSTANT_InvokeDynamic_info{}
	cid.Tag = CONSTANT_InvokeDynamic
	failErr(readUint16(r, &cid.BoostrapMethodAttrIndex))
	failErr(readUint16(r, &cid.NameAndTypeIndex))
	return cid
}


