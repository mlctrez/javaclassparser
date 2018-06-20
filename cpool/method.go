package cpool

import (
	"github.com/mlctrez/javaclassparser/ioutil"
)

type ConstantMethodHandleInfo struct {
	ConstBase
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func ReadConstantMethodHandleInfo(r PoolReader) *ConstantMethodHandleInfo {
	mh := &ConstantMethodHandleInfo{}
	mh.Pool = r.ConstantPool
	mh.Tag = ConstantMethodHandle
	mh.Type = "CONSTANT_MethodHandle_info"
	failErr(ioutil.ReadUint8(r, &mh.ReferenceKind))
	failErr(ioutil.ReadUint16(r, &mh.ReferenceIndex))
	return mh
}

type ConstantMethodTypeInfo struct {
	ConstBase
	DescriptorIndex uint16
}

func ReadConstantMethodTypeInfo(r PoolReader) *ConstantMethodTypeInfo {
	mt := &ConstantMethodTypeInfo{}
	mt.Pool = r.ConstantPool
	mt.Tag = ConstantMethodType
	mt.Type = "CONSTANT_MethodType_info"
	failErr(ioutil.ReadUint16(r, &mt.DescriptorIndex))
	return mt
}

type ConstantInvokeDynamicInfo struct {
	ConstBase
	BoostrapMethodAttrIndex uint16
	NameAndTypeIndex        uint16
}

func ReadConstantInvokeDynamicInfo(r PoolReader) *ConstantInvokeDynamicInfo {
	cid := &ConstantInvokeDynamicInfo{}
	cid.Pool = r.ConstantPool
	cid.Tag = ConstantInvokeDynamic
	cid.Type = "CONSTANT_InvokeDynamic_info"
	failErr(ioutil.ReadUint16(r, &cid.BoostrapMethodAttrIndex))
	failErr(ioutil.ReadUint16(r, &cid.NameAndTypeIndex))
	return cid
}
