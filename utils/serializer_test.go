package utils

import "testing"

const FILE_DIR = "../cache.data"

func TestSaveToFile(t *testing.T) {
	m := make(map[interface{}]interface{})
	m[123] = 123
	m[122] = 1222
	t.Log(m)
	err := SaveToFile(FILE_DIR, m)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadFromFile(t *testing.T) {
	m, err := LoadFromFile(FILE_DIR)
	if err != nil {
		t.Error(err)
	}
	t.Log(m)
}
