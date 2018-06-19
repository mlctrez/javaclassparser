package javaclassparser

import (
	"strings"
)

// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4

// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4-140
const CONSTANT_Class = 7
const CONSTANT_Fieldref = 9
const CONSTANT_Methodref = 10
const CONSTANT_InterfaceMethodref = 11
const CONSTANT_String = 8
const CONSTANT_Integer = 3
const CONSTANT_Float = 4
const CONSTANT_Long = 5
const CONSTANT_Double = 6
const CONSTANT_NameAndType = 12
const CONSTANT_Utf8 = 1
const CONSTANT_MethodHandle = 15
const CONSTANT_MethodType = 16
const CONSTANT_InvokeDynamic = 18

// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.5-200-A.1
// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.6-200-A.1
// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.1-200-E.1
// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.24
const ACC_PUBLIC AccessFlag = 0x0001
const ACC_PRIVATE AccessFlag = 0x0002
const ACC_PROTECTED AccessFlag = 0x0004
const ACC_STATIC AccessFlag = 0x0008
const ACC_FINAL AccessFlag = 0x0010
const ACC_SYNCHRONIZED AccessFlag = 0x0020
const ACC_VOLATILE AccessFlag = 0x0040
const ACC_BRIDGE AccessFlag = 0x0040
const ACC_TRANSIENT AccessFlag = 0x0080
const ACC_VARARGS AccessFlag = 0x0080
const ACC_NATIVE AccessFlag = 0x0100
const ACC_INTERFACE AccessFlag = 0x0200
const ACC_ABSTRACT AccessFlag = 0x0400
const ACC_STRICT AccessFlag = 0x0800
const ACC_SYNTHETIC AccessFlag = 0x1000
const ACC_ANNOTATION AccessFlag = 0x2000
const ACC_ENUM AccessFlag = 0x4000
const ACC_MANDATED AccessFlag = 0x8000
const ACC_SUPER AccessFlag = 0x0020

type AccessFlag uint16

func (f AccessFlag) String() string {
	var mods []string
	if (ACC_PUBLIC & f) != 0 {
		mods = append(mods, "public")
	}
	if (ACC_PRIVATE & f) != 0 {
		mods = append(mods, "private")
	}
	if (ACC_PROTECTED & f) != 0 {
		mods = append(mods, "protected")
	}
	if (ACC_STATIC & f) != 0 {
		mods = append(mods, "static")
	}
	if (ACC_FINAL & f) != 0 {
		mods = append(mods, "final")
	}
	if (ACC_SYNCHRONIZED & f) != 0 {
		mods = append(mods, "synchronized")
	}
	if (ACC_VOLATILE & f) != 0 {
		mods = append(mods, "volatile")
	}
	if (ACC_BRIDGE & f) != 0 {
		mods = append(mods, "bridge")
	}
	if (ACC_TRANSIENT & f) != 0 {
		mods = append(mods, "transient")
	}
	if (ACC_VARARGS & f) != 0 {
		mods = append(mods, "varargs")
	}
	if (ACC_NATIVE & f) != 0 {
		mods = append(mods, "native")
	}
	if (ACC_INTERFACE & f) != 0 {
		mods = append(mods, "interface")
	}
	if (ACC_ABSTRACT & f) != 0 {
		mods = append(mods, "abstract")
	}
	if (ACC_STRICT & f) != 0 {
		mods = append(mods, "strict")
	}
	if (ACC_SYNTHETIC & f) != 0 {
		mods = append(mods, "synthetic")
	}
	if (ACC_ANNOTATION & f) != 0 {
		mods = append(mods, "annotation")
	}
	if (ACC_ENUM & f) != 0 {
		mods = append(mods, "enum")
	}
	if (ACC_MANDATED & f) != 0 {
		mods = append(mods, "mandated")
	}
	if (ACC_SUPER & f) != 0 {
		mods = append(mods, "super")
	}
	return strings.Join(mods, " ")
}
