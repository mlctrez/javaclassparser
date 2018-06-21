package cpool

import (
	"github.com/mlctrez/javaclassparser/ioutil"
)

type ConstantMethodHandleInfo struct {
	ConstBase
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func ReadConstantMethodHandleInfo(r PoolReader) (mh *ConstantMethodHandleInfo, err error) {
	mh = &ConstantMethodHandleInfo{}
	mh.Pool = r.ConstantPool
	mh.Tag = ConstantMethodHandle
	mh.Type = "CONSTANT_MethodHandle_info"
	if err = ioutil.ReadUint8(r, &mh.ReferenceKind); err != nil {
		return
	}
	err = ioutil.ReadUint16(r, &mh.ReferenceIndex)
	return
}

type ConstantMethodTypeInfo struct {
	ConstBase
	DescriptorIndex uint16
}

func ReadConstantMethodTypeInfo(r PoolReader) (mt *ConstantMethodTypeInfo, err error) {
	mt = &ConstantMethodTypeInfo{}
	mt.Pool = r.ConstantPool
	mt.Tag = ConstantMethodType
	mt.Type = "CONSTANT_MethodType_info"
	err = ioutil.ReadUint16(r, &mt.DescriptorIndex)
	return
}

type ConstantInvokeDynamicInfo struct {
	ConstBase
	BoostrapMethodAttrIndex uint16
	NameAndTypeIndex        uint16
}

func ReadConstantInvokeDynamicInfo(r PoolReader) (cid *ConstantInvokeDynamicInfo, err error) {
	cid = &ConstantInvokeDynamicInfo{}
	cid.Pool = r.ConstantPool
	cid.Tag = ConstantInvokeDynamic
	cid.Type = "CONSTANT_InvokeDynamic_info"
	if err = ioutil.ReadUint16(r, &cid.BoostrapMethodAttrIndex); err != nil {
		return
	}
	err = ioutil.ReadUint16(r, &cid.NameAndTypeIndex)
	return
}
