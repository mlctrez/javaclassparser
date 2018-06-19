package javaclassparser

import (
	"testing"
)

func TestAccessFlag_String(t *testing.T) {
	if "public" != AccessFlag(ACC_PUBLIC).String() {
		t.Error("access flags public test fail")
	}
	if "public static" != AccessFlag(ACC_PUBLIC | ACC_STATIC).String() {
		t.Error("access flags 'public static' test fail")
	}

	if "public static final" != AccessFlag(ACC_PUBLIC | ACC_STATIC | ACC_FINAL).String() {
		t.Error("access flags 'public static final' test fail")
	}

	if "" != AccessFlag(0).String() {
		t.Error("access flags '' test fail")
	}
}
