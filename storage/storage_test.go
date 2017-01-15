package storage

import (
	"errors"
	"os"
	"testing"
)

const storefile = "/tmp/peernotify.db"

func dropStorage(fname string) error {
	return os.Remove(fname)
}

func writeAndClose(fname string, start, end int) error {
	s, err := NewStorage(fname)
	if err != nil {
		return err
	}
	defer s.Close()
	for i := start; i < end; i++ {
		key := "key" + string(i)
		value := "value" + string(i)
		if err := s.Store([]byte(key), []byte(value)); err != nil {
			return err
		}
	}
	return nil
}

func readAndVerify(fname string, start, end int) error {
	s, err := NewStorage(fname)
	if err != nil {
		return err
	}
	defer s.Close()
	for i := start; i < end; i++ {
		key := "key" + string(i)
		testValue := "value" + string(i)
		value, err := s.Get([]byte(key))
		if err != nil {
			return err
		}
		if string(value) != testValue {
			return errors.New("Incorrect restored value")
		}
	}
	return nil
}

func TestEmptyStorage(t *testing.T) {
	err := writeAndClose(storefile, 0, 100)
	if err != nil {
		t.Error()
	}
	err = readAndVerify(storefile, 0, 100)
	if err != nil {
		t.Error()
	}
	err = dropStorage(storefile)
	if err != nil {
		t.Error(err)
	}
}

func TestNonEmptyStorage(t *testing.T) {
	err := writeAndClose(storefile, 0, 100)
	if err != nil {
		t.Error()
	}
	err = writeAndClose(storefile, 100, 200)
	if err != nil {
		t.Error()
	}
	err = readAndVerify(storefile, 0, 200)
	if err != nil {
		t.Error()
	}
	err = dropStorage(storefile)
	if err != nil {
		t.Error(err)
	}
}
