package ioutil

import (
	"encoding/binary"
	"io"
)

// TODO: this class could leverage pools for the byte buffers to avoid some allocation overhead

// ReadUint8 reads one byte into the provided uint8 pointer
func ReadUint8(r io.Reader, i *uint8) (err error) {
	b := []byte{0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = b[0]
	}
	return err
}

// ReadUint16 reads two bytes into the provided uint16 pointer in BigEndian order
func ReadUint16(r io.Reader, i *uint16) (err error) {
	b := []byte{0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = binary.BigEndian.Uint16(b)
	}
	return err
}

// ReadUint32 reads four bytes into the provided uint32 pointer in BigEndian order
func ReadUint32(r io.Reader, i *uint32) (err error) {
	b := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = binary.BigEndian.Uint32(b)
	}
	return err
}

// ReadInt32 reads four bytes into the provided int32 pointer in BigEndian order
func ReadInt32(r io.Reader, i *int32) (err error) {
	b := []byte{0, 0, 0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = int32(binary.BigEndian.Uint32(b))
	}
	return err
}

// ReadInt32 reads eight bytes into the provided int64 pointer in BigEndian order
func ReadInt64(r io.Reader, i *int64) (err error) {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if _, err := io.ReadFull(r, b); err == nil {
		*i = int64(binary.BigEndian.Uint64(b))
	}
	return err
}
