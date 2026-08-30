package main

import (
	"testing"

	"github.com/veandco/go-sdl2/sdl"
)

// These exercise the pure split-tree logic (leaves/replaceNode/nextFocus)
// without spawning a real PTY or SDL window -- the part of Phase 1's
// i3-primitive split that's actually worth regression-testing on its own,
// same as vterm/font already get their own package tests.

func fakeLeaf() *node {
	return &node{pane: &pane{}}
}

func TestLeavesSingle(t *testing.T) {
	n := fakeLeaf()
	got := leaves(n)
	if len(got) != 1 || got[0] != n {
		t.Fatalf("leaves(single) = %v, want [n]", got)
	}
}

func TestLeavesSplitOrder(t *testing.T) {
	a, b, c := fakeLeaf(), fakeLeaf(), fakeLeaf()
	// ((a b) c), vertical over horizontal, arbitrary nesting.
	inner := &node{axis: splitVertical, children: [2]*node{a, b}}
	root := &node{axis: splitHorizontal, children: [2]*node{inner, c}}
	got := leaves(root)
	if len(got) != 3 || got[0] != a || got[1] != b || got[2] != c {
		t.Fatalf("leaves(nested) = %v, want [a b c]", got)
	}
}

func TestReplaceNodeAtRoot(t *testing.T) {
	root := fakeLeaf()
	replacement := fakeLeaf()
	rootPtr := root
	replaceNode(&rootPtr, root, replacement)
	if rootPtr != replacement {
		t.Fatalf("replaceNode at root did not swap: got %v want %v", rootPtr, replacement)
	}
}

func TestReplaceNodeInContainer(t *testing.T) {
	a, b := fakeLeaf(), fakeLeaf()
	root := &node{axis: splitVertical, children: [2]*node{a, b}}
	replacement := &node{axis: splitHorizontal, children: [2]*node{fakeLeaf(), fakeLeaf()}}
	rootPtr := root
	replaceNode(&rootPtr, a, replacement)
	if root.children[0] != replacement {
		t.Fatalf("replaceNode did not swap child a: got %v want %v", root.children[0], replacement)
	}
	if root.children[1] != b {
		t.Fatalf("replaceNode disturbed sibling b: got %v want %v", root.children[1], b)
	}
}

func TestNextFocusCyclesForwardAndBack(t *testing.T) {
	a, b, c := fakeLeaf(), fakeLeaf(), fakeLeaf()
	root := &node{axis: splitVertical, children: [2]*node{
		{axis: splitHorizontal, children: [2]*node{a, b}},
		c,
	}}
	if got := nextFocus(root, a, sdl.K_RIGHT); got != b {
		t.Fatalf("nextFocus(a, right) = %v, want b", got)
	}
	if got := nextFocus(root, c, sdl.K_RIGHT); got != a {
		t.Fatalf("nextFocus(c, right) should wrap to a, got %v", got)
	}
	if got := nextFocus(root, a, sdl.K_LEFT); got != c {
		t.Fatalf("nextFocus(a, left) should wrap to c, got %v", got)
	}
}

func TestNextFocusSingleLeafIsNoOp(t *testing.T) {
	a := fakeLeaf()
	if got := nextFocus(a, a, sdl.K_RIGHT); got != a {
		t.Fatalf("nextFocus(single leaf) = %v, want a (no-op)", got)
	}
}
