package tokens

import (
	"bytes"
	"testing"
)

func TestCopying(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0, 0, 1}
	ts := newTokenSet(rawTS)
	rawTS[3] = 1
	if bytes.Equal([]byte(ts), []byte{0, 1, 1, 1, 0, 1}) {
		t.Error()
	}
}

func TestAddStart(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0, 0, 1}
	expTS := []byte{1, 1, 1, 0, 0, 1}
	ts := newTokenSet(rawTS)
	newTS := ts.addAt(0)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestAddInner(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0, 0, 1}
	expTS := []byte{0, 1, 1, 1, 0, 1}
	ts := newTokenSet(rawTS)
	newTS := ts.addAt(3)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestAddAfter(t *testing.T) {
	rawTS := []byte{0, 1, 1}
	expTS := []byte{0, 1, 1, 0, 0, 1}
	ts := newTokenSet(rawTS)
	newTS := ts.addAt(5)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestAddRightAfter(t *testing.T) {
	rawTS := []byte{0, 1, 1}
	expTS := []byte{0, 1, 1, 1}
	ts := newTokenSet(rawTS)
	newTS := ts.addAt(3)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestGetAt(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0}
	ts := newTokenSet(rawTS)
	if !(!ts.getAt(0) && ts.getAt(1) && ts.getAt(2) && !ts.getAt(3)) {
		t.Error()
	}
}

func TestGetAfter(t *testing.T) {
	rawTS := []byte{0, 1, 1, 0}
	ts := newTokenSet(rawTS)
	if !(!ts.getAt(4) && !ts.getAt(8)) {
		t.Error()
	}
}

func TestDropStart(t *testing.T) {
	rawTS := []byte{1, 1, 1, 0, 1, 0, 1}
	expTS := []byte{1, 1, 1, 0, 1, 0, 1}
	ts := newTokenSet(rawTS)
	newTS := ts.dropUntil(0)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestDropUntilInner(t *testing.T) {
	rawTS := []byte{1, 1, 1, 0, 1, 0, 1}
	expTS := []byte{1, 0, 1, 0, 1}
	ts := newTokenSet(rawTS)
	newTS := ts.dropUntil(2)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestDropUntilAfter(t *testing.T) {
	rawTS := []byte{1, 1, 1, 0, 1, 0, 1}
	expTS := []byte{}
	ts := newTokenSet(rawTS)
	newTS := ts.dropUntil(10)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}

func TestDropUntilRightAfter(t *testing.T) {
	rawTS := []byte{1, 1, 1, 0, 1, 0, 1}
	expTS := []byte{1}
	ts := newTokenSet(rawTS)
	newTS := ts.dropUntil(6)
	if !bytes.Equal([]byte(newTS), expTS) {
		t.Error()
	}
}
