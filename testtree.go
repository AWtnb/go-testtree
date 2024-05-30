package testtree

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/AWtnb/go-walk"
)

func GetPerm(path string) fs.FileMode {
	s := string(os.PathSeparator)
	elems := strings.Split(path, s)
	for i := 0; i < len(elems); i++ {
		ln := len(elems) - i
		p := strings.Join(elems[0:ln], s)
		if fs, err := os.Stat(p); err == nil {
			return fs.Mode() & os.ModePerm
		}

	}
	return 0700
}

func MakeTestDir(path string) error {
	p := GetPerm(path)
	err := os.MkdirAll(path, p)
	return err
}

func MakeTestFile(path string) error {
	_, err := os.Create(path)
	return err
}

func MakeTestTree(root string, dirs []string, files []string) error {
	if err := MakeTestDir(root); err != nil {
		return err
	}
	for _, d := range dirs {
		p := filepath.Join(root, d)
		if err := MakeTestDir(p); err != nil {
			return err
		}
	}
	for _, f := range files {
		p := filepath.Join(root, f)
		if err := MakeTestFile(p); err != nil {
			return err
		}
	}
	return nil
}

func GetChildItems(root string) []string {
	var d walk.Dir
	d.Init(root, true, -1, "")
	found, err := d.GetChildItem()
	if err != nil {
		fmt.Println(err)
	}
	return found
}
