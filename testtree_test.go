package testtree_test

import (
	"testing"

	"github.com/AWtnb/go-testtree"
)

func TestGetPerm(t *testing.T) {
	t.Log(testtree.GetPerm(`C:\Personal\gotemp\hoge`))
}

func TestMakeTestDir(t *testing.T) {
	err := testtree.MakeTestDir(`C:\Personal\gotemp\huga`)
	if err != nil {
		t.Error(err)
	}
}
