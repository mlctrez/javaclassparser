package aflag

import (
	"errors"
	"io"
	"strings"

	"github.com/mlctrez/javaclassparser/ioutil"
)

type flag struct {
	Bits uint16
	Name string
}

// FieldAccessFlags contains the public, private, protected, etc flags for a field
// See http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.5-200-A.1
type FieldAccessFlags uint16

var fi [9]flag

func init() {
	fi[0] = flag{0x0001, "public"}
	fi[1] = flag{0x0002, "private"}
	fi[2] = flag{0x0004, "protected"}
	fi[3] = flag{0x0008, "static"}
	fi[4] = flag{0x0010, "final"}
	fi[5] = flag{0x0040, "volatile"}
	fi[6] = flag{0x0080, "transient"}
	fi[7] = flag{0x1000, "synthetic"}
	fi[8] = flag{0x4000, "enum"}
}

// ReadFieldAccessFlags reads the field access flags from the provided reader.
// The result is invalid if the error returned is not nil.
func ReadFieldAccessFlags(r io.Reader) (FieldAccessFlags, error) {
	var af uint16
	err := ioutil.ReadUint16(r, &af)
	return FieldAccessFlags(af), err
}

func (f FieldAccessFlags) String() string {
	var mods []string
	for _, flag := range fi {
		if (uint16(f) & flag.Bits) != 0 {
			mods = append(mods, flag.Name)
		}
	}
	return strings.Join(mods, " ")
}

// MethodAccessFlags contains the public, private, protected, etc flags for a field
// See http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.6-200-A.1
type MethodAccessFlags uint16

var mi [12]flag

func init() {
	mi[0] = flag{0x0001, "public"}
	mi[1] = flag{0x0002, "private"}
	mi[2] = flag{0x0004, "protected"}
	mi[3] = flag{0x0008, "static"}
	mi[4] = flag{0x0010, "final"}
	mi[5] = flag{0x0020, "synchronized"}
	mi[6] = flag{0x0040, "bridge"}
	mi[7] = flag{0x0080, "varargs"}
	mi[8] = flag{0x0100, "native"}
	mi[9] = flag{0x0400, "abstract"}
	mi[10] = flag{0x0800, "strict"}
	mi[11] = flag{0x1000, "synthetic"}
}

// ReadMethodAccessFlags reads the method access flags from the provided reader.
// The result is invalid if the error returned is not nil.
func ReadMethodAccessFlags(r io.Reader) (MethodAccessFlags, error) {
	var af uint16
	err := ioutil.ReadUint16(r, &af)
	return MethodAccessFlags(af), err
}

func (f MethodAccessFlags) String() string {
	var mods []string
	for _, flag := range mi {
		if (uint16(f) & flag.Bits) != 0 {
			mods = append(mods, flag.Name)
		}
	}
	return strings.Join(mods, " ")
}

// ClassAccessFlags contains the public, abstract, etc flags for a field
// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.1-200-E.1
type ClassAccessFlags uint16

var ci [9]flag

func init() {
	ci[0] = flag{0x0001, "public"}
	ci[1] = flag{0x0010, "final"}
	ci[2] = flag{0x0020, "super"}
	ci[3] = flag{0x0200, "interface"}
	ci[4] = flag{0x0400, "abstract"}
	ci[5] = flag{0x1000, "synthetic"}
	ci[6] = flag{0x2000, "annotation"}
	ci[7] = flag{0x4000, "enum"}
}

// ReadClassAccessFlags reads the method access flags from the provided reader.
// The result is invalid if the error returned is not nil.
func ReadClassAccessFlags(r io.Reader) (ClassAccessFlags, error) {
	var af uint16
	err := ioutil.ReadUint16(r, &af)
	return ClassAccessFlags(af), err
}

func (f *ClassAccessFlags) Validate() error {
	// TODO: revisit if this is needed
	if (*f & 0x0200) != 0 {
		*f |= 0x0400
	}
	if ((*f & 0x0400) != 0) && ((*f & 0x0010) != 0) {
		return errors.New("class can't be both final and abstract")
	}
	return nil
}

func (f ClassAccessFlags) String() string {
	var mods []string

	for _, flag := range ci {
		if (uint16(f) & flag.Bits) != 0 {
			mods = append(mods, flag.Name)
		}
	}
	return strings.Join(mods, " ")
}

// ReadClassAccessFlags reads the method access flags from the provided reader.
// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.24
// TODO: this is not currently used in when parsing but is here for completeness
type MethodParameterAccessFlags uint16

var mp [3]flag

func init() {
	mp[0] = flag{0x0001, "final"}
	mp[1] = flag{0x1000, "synthetic"}
	mp[2] = flag{0x8000, "mandated"}
}

func (f MethodParameterAccessFlags) String() string {
	var mods []string
	for _, flag := range mp {
		if (uint16(f) & flag.Bits) != 0 {
			mods = append(mods, flag.Name)
		}
	}
	return strings.Join(mods, " ")
}
