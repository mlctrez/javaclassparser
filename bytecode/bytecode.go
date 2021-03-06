package bytecode

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/mlctrez/javaclassparser/cpool"
	"github.com/mlctrez/javaclassparser/ioutil"
)

type ByteCode struct {
	Offset    uint32
	Operand   string
	Arguments []byte
	IndexByte bool
}

var cache map[byte]Reader

func Read(r io.Reader) (codes []*ByteCode, err error) {

	if cache == nil {
		cache = buildOpCodeFunctionMap()
	}

	var codeLength uint32
	if err = ioutil.ReadUint32(r, &codeLength); err != nil {
		return nil, err
	}

	var p uint32
	var oc = []byte{0}
	var rc = &Context{r, &p}

	for p = 0; p < codeLength; p++ {

		var n int
		if n, err = io.ReadFull(r, oc); err != nil {
			return
		}
		if n != 1 {
			return nil, errors.New("error reading opcode")
		}

		reader := cache[oc[0]]

		var code *ByteCode
		if code, err = reader(rc); err != nil {
			return nil, err
		}
		codes = append(codes, code)
	}
	return
}

type Context struct {
	io.Reader
	p *uint32
}

type Reader func(rc *Context) (*ByteCode, error)

func Simple(op string, c *Context) (*ByteCode, error) {
	return &ByteCode{Operand: op, Offset: *c.p}, nil
}

func WithArgs(op string, c *Context, index bool, count int) (bc *ByteCode, err error) {
	bc = &ByteCode{Operand: op, Offset: *c.p, IndexByte: index, Arguments: make([]byte, count)}
	if _, err = io.ReadFull(c, bc.Arguments); err != nil {
		return nil, err
	}
	*c.p = *c.p + uint32(count)
	return
}

func (bc ByteCode) String() string {
	if bc.Arguments != nil {
		return fmt.Sprintf("%s %v", bc.Operand, bc.Arguments)
	}
	return bc.Operand
}

func (bc *ByteCode) StringWithIndex(pool cpool.ConstantPool) string {

	if len(bc.Arguments) > 0 {
		if bc.IndexByte {
			index := uint16(bc.Arguments[0])*255 + uint16(bc.Arguments[1])
			return fmt.Sprintf("%s %v", bc.Operand, pool.Lookup(index))
		}
		if bc.Operand == "ldc" {
			return fmt.Sprintf("%s %q", bc.Operand, pool.Lookup(uint16(bc.Arguments[0])))
		}

	}
	return bc.String()
}

/*
A tableswitch is a variable-length instruction. Immediately after the tableswitch opcode, between zero and
three bytes must act as padding, such that defaultbyte1 begins at an address that is a multiple of four bytes
from the start of the current method (the opcode of its first instruction). Immediately after the padding are
bytes constituting three signed 32-bit values: default, low, and high. Immediately following are bytes constituting
a series of high - low + 1 signed 32-bit offsets. The value low must be less than or equal to high.
The high - low + 1 signed 32-bit offsets are treated as a 0-based jump table. Each of these signed 32-bit
values is constructed as (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4.
*/
func TableSwitch(op string, p *uint32, r io.Reader) (bc *ByteCode, err error) {
	bc = &ByteCode{Operand: op, Offset: *p}

	padding := 3 - *p%4
	if padding > 0 {
		*p = *p + padding
		pb := make([]byte, padding)
		_, err = io.ReadFull(r, pb)
		if err != nil {
			return nil, err
		}
	}

	var defaultByte int32
	var lowByte int32
	var highByte int32

	if err = ioutil.ReadInt32(r, &defaultByte); err != nil {
		return nil, err
	}
	if err = ioutil.ReadInt32(r, &lowByte); err != nil {
		return nil, err
	}
	if err = ioutil.ReadInt32(r, &highByte); err != nil {
		return nil, err
	}

	jumpCodeLength := (highByte - lowByte) + 1
	if jumpCodeLength < 1 {
		err = fmt.Errorf("error calculating jumpCodeLength low=%d high=%d default=%d", lowByte, highByte, defaultByte)
		return
	}
	jumpCodes := make([]int32, jumpCodeLength)

	for i := 0; i < int(jumpCodeLength); i++ {
		var jumpCode int32

		if err = ioutil.ReadInt32(r, &jumpCode); err != nil {
			return nil, err
		}
		jumpCodes = append(jumpCodes, jumpCode)
	}

	*p = *p + 12 + uint32(jumpCodeLength*4)
	return bc, nil
}

type MatchOffset struct {
	Match  int32
	Offset int32
}

func LookupSwitch(op string, p *uint32, r io.Reader) (bc *ByteCode, err error) {
	bc = &ByteCode{Operand: op, Offset: *p}

	padding := 3 - *p%4
	if padding > 0 {
		*p = *p + padding
		pb := make([]byte, padding)
		if _, err = io.ReadFull(r, pb); err != nil {
			return nil, err
		}
	}

	var defaultByte int32
	var nPairs int32
	if err = ioutil.ReadInt32(r, &defaultByte); err != nil {
		return nil, err
	}
	if err = ioutil.ReadInt32(r, &nPairs); err != nil {
		return nil, err
	}
	if nPairs < 0 {
		panic(nPairs)
	}
	matchOffsets := make([]MatchOffset, nPairs)

	for i := 0; i < int(nPairs); i++ {
		var match int32
		if err = ioutil.ReadInt32(r, &match); err != nil {
			return nil, err
		}
		var offset int32
		if err = ioutil.ReadInt32(r, &offset); err != nil {
			return nil, err
		}
		matchOffsets = append(matchOffsets, MatchOffset{Match: match, Offset: offset})
	}

	*p = *p + 8 + uint32(nPairs*8)
	return
}

func Wide(op string, p *uint32, r io.Reader) (bc *ByteCode, err error) {
	bc = &ByteCode{Operand: op, Offset: *p}

	// TODO: binary.Read is overkill here. could be io.Read
	inst := []byte{0}
	if err = binary.Read(r, binary.BigEndian, inst); err != nil {
		return
	}

	index := []byte{0, 0}
	if err = binary.Read(r, binary.BigEndian, index); err != nil {
		return
	}
	if inst[0] == 0x84 {
		c := []byte{0, 0}

		if err = binary.Read(r, binary.BigEndian, c); err != nil {
			return
		}
		*p = *p + 5
		return
	}
	*p = *p + 3
	return
}
