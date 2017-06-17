package tokens

import (
	"bytes"
	"testing"
)

func TestCopying(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0, 0, 1}
	ts := NewTokenSet(rawTS)
	rawTS[3] = 1
	if bytes.Equal([]byte(ts), []byte{0, 1, 1, 1, 0, 1}) {
		t.Error()
	}
}

func TestAddStart(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0, 0, 1}
	expTS := []byte{1, 1, 1, 0, 0, 1}
	ts := NewTokenSet(rawTS)
	newTS := ts.AddAt(0)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestAddInner(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0, 0, 1}
	expTS := []byte{0, 1, 1, 1, 0, 1}
	ts := NewTokenSet(rawTS)
	newTS := ts.AddAt(3)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestAddAfter(t *testing.T) {
	rawTS := []byte{0, 1, 1}
	expTS := []byte{0, 1, 1, 0, 0, 1}
	ts := NewTokenSet(rawTS)
	newTS := ts.AddAt(5)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestAddRightAfter(t *testing.T) {
	rawTS := []byte{0, 1, 1}
	expTS := []byte{0, 1, 1, 1}
	ts := NewTokenSet(rawTS)
	newTS := ts.AddAt(3)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestGetAt(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0}
	ts := NewTokenSet(rawTS)
	if !(!ts.GetAt(0) && ts.GetAt(1) && ts.GetAt(2) && !ts.GetAt(3)) {
		t.Error()
	}
}

func TestGetAfter(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0}
	ts := NewTokenSet(rawTS)
	if !(!ts.GetAt(4) && !ts.GetAt(8)) {
		t.Error()
	}
}

func TestDropStart(t *testing.T) {
	rawTS := []byte{1, 1, 1, 0, 1, 0, 1}
	expTS := []byte{1, 1, 1, 0, 1, 0, 1}
	ts := NewTokenSet(rawTS)
	newTS := ts.DropUntil(0)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestDropUntilInner(t *testing.T) {
	rawTS := []byte{1, 1, 1, 0, 1, 0, 1}
	expTS := []byte{1, 0, 1, 0, 1}
	ts := NewTokenSet(rawTS)
	newTS := ts.DropUntil(2)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestDropUntilAfter(t *testing.T) {
	rawTS := []byte{1, 1, 1, 0, 1, 0, 1}
	expTS := []byte{}
	ts := NewTokenSet(rawTS)
	newTS := ts.DropUntil(10)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestDropUntilRightAfter(t *testing.T) {
	rawTS := []byte{1, 1, 1, 0, 1, 0, 1}
	expTS := []byte{1}
	ts := NewTokenSet(rawTS)
	newTS := ts.DropUntil(6)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}
