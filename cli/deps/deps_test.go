package main

import (
	"testing"
)

func TestDependencyManager_included(t *testing.T) {
	d := &DependencyManager{prefixes: []string{"com/"}}
	if d.included("java/") {
		t.Fatal("java should not be included")
	}
	if !d.included("com/") {
		t.Fatal("com should be included")
	}
}

func TestDependencyManager_extractName(t *testing.T) {
	d := &DependencyManager{prefixes: []string{"com/"}}

	if d.extractName("") != "" {
		t.Fatal("extractName empty fail")
	}
	if d.extractName("com") != "com" {
		t.Fatal("extractName com fail")
	}
	if d.extractName("com/domain") != "com/domain" {
		t.Fatal("extractName com/domain fail")
	}
	if d.extractName("com/domain/Foo$2") != "com/domain/Foo" {
		t.Fatal("extractName com/domain/Foo$2 fail")
	}
	if d.extractName("com/domain/Foo method args") != "com/domain/Foo" {
		t.Fatal("extractName com/domain/Foo with methods fail")
	}

	d.packageOnly = true
	if d.extractName("com/domain/Foo") != "com/domain" {
		t.Fatal("extractName com/domain/Foo to package name fail")
	}

	d.depth = 3
	if d.extractName("a/b/c/d/e") != "a/b/c" {
		t.Fatal("depth shorten test failed", d.extractName("a/b/c/d/e"))
	}
	if d.extractName("a/b") != "a/b" {
		t.Fatal("depth min test failed", d.extractName("a/b"))
	}

}
