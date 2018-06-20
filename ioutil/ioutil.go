package ioutil

import (
	"encoding/binary"
	"io"
)

func ReadUint8(r io.Reader, i *uint8) (err error) {
	b := []byte{0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = b[0]
	}
	return err
}

func ReadUint16(r io.Reader, i *uint16) (err error) {
	b := []byte{0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = binary.BigEndian.Uint16(b)
	}
	return err
}

func ReadUint32(r io.Reader, i *uint32) (err error) {
	b := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = binary.BigEndian.Uint32(b)
	}
	return err
}

func ReadInt32(r io.Reader, i *int32) (err error) {
	b := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = int32(binary.BigEndian.Uint32(b))
	}
	return err
}

func ReadInt64(r io.Reader, i *int64) (err error) {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = int64(binary.BigEndian.Uint64(b))
	}
	return err
}
