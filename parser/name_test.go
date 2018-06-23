package parser

import (
	"testing"
)

func TestExtractName(t *testing.T) {

	if ExtractName("") != "" {
		t.Fatal("extractName empty fail")
	}
	if ExtractName("com") != "com" {
		t.Fatal("extractName com fail")
	}
	if ExtractName("com/domain") != "com/domain" {
		t.Fatal("extractName com/domain fail")
	}
	if ExtractName("com/domain/Foo$2") != "com/domain/Foo" {
		t.Fatal("extractName com/domain/Foo$2 fail")
	}
	if ExtractName("[Lcom/domain/Foo;") != "com/domain/Foo" {
		t.Fatal("extractName com/domain/Foo with [L")
	}

}
