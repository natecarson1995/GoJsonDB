package gojsondb

import (
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New("test.json")

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSet(t *testing.T) {
	db, err := New("test.json")
	if err != nil {
		t.Fatal(err)
	}

	db.SetString("item", "test test")
	data, err := db.GetString("item")
	if err != nil {
		t.Fatal(err)
	}
	if data != "test test" {
		t.Fatal("Testing of data get set failed")
	}
}

func TestListKeys(t *testing.T) {
	db, err := New("test.json")
	if err != nil {
		t.Fatal(err)
	}

	db.SetString("item", "test test")
	keys := db.ListKeys()
	if keys[0] != "item" {
		t.Fatal("Testing of listkeys failed")
	}
}