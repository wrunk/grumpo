package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestLoadMeta(t *testing.T) {

	m, err := loadMeta("testfile.md")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Meta: (%s)", spew.Sdump(m))
	}

}
