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

func TestDependencyManager_extractNamePackage(t *testing.T) {
	d := &DependencyManager{}

	d.packageOnly = true
	if d.extractName("com/domain/Foo") != "com/domain" {
		t.Fatal("extractName com/domain/Foo to package name fail")
	}
}

func TestDependencyManager_extractNameDepth(t *testing.T) {
	d := &DependencyManager{}

	d.depth = 3
	if d.extractName("a/b/c/d/e") != "a/b/c" {
		t.Fatal("depth shorten test failed", d.extractName("a/b/c/d/e"))
	}
	if d.extractName("a/b") != "a/b" {
		t.Fatal("depth min test failed", d.extractName("a/b"))
	}

}

func TestDependencyManager_extractConflict(t *testing.T) {
	d := &DependencyManager{}
	defer func() {
		if r := recover(); r != nil {
			if s, ok := r.(string); ok {
				if s != "cannot set packageOnly and depth" {
					t.Fatal("unexpected panic value " + s)
				}
			}
		}
	}()
	d.packageOnly = true
	d.depth = 3
	if d.extractName("a/b/c/d/e") != "a/b/c" {
		t.Fatal("depth shorten test failed", d.extractName("a/b/c/d/e"))
	}

}
