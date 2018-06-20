package aflag

import (
	"testing"
)

func TestAccessFlag_String(t *testing.T) {
	if "public" != MethodAccessFlags(mi[0].Bits).String() {
		t.Error("access flags public test fail")
	}
	if "public static" != MethodAccessFlags(mi[0].Bits | mi[3].Bits).String() {
		t.Error("access flags 'public static' test fail")
	}

	if "public static final" != MethodAccessFlags(mi[0].Bits | mi[3].Bits | mi[4].Bits).String() {
		t.Error("access flags 'public static final' test fail")
	}

	if "" != ClassAccessFlags(0).String() {
		t.Error("access flags '' test fail")
	}
}
