package cpool

import (
	"encoding/binary"
	"io"
)

func failErr(err error) {
	if err != nil {
		panic(err)
	}
}

func read(r io.Reader, data interface{}) {
	failErr(binary.Read(r, binary.BigEndian, data))
}

type ConstBase struct {
	Pool ConstantPool
	Tag  uint8
}

func (cb ConstBase) SetPool(cp ConstantPool) {
	cb.Pool = cp
}

func readUint8(r io.Reader, i *uint8) (err error) {
	b := []byte{0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = b[0]
	}
	return err
}

func readUint16(r io.Reader, i *uint16) (err error) {
	b := []byte{0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = binary.BigEndian.Uint16(b)
	}
	return err
}

func readInt32(r io.Reader, i *int32) (err error) {
	b := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = int32(binary.BigEndian.Uint32(b))
	}
	return err
}

func readUint32(r io.Reader, i *uint32) (err error) {
	b := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = binary.BigEndian.Uint32(b)
	}
	return err
}

func readInt64(r io.Reader, i *int64) (err error) {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = int64(binary.BigEndian.Uint64(b))
	}
	return err
}
