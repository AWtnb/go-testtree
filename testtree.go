package testtree

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AWtnb/go-walk"
)

func MakeTestDir(path string) error {
	err := os.MkdirAll(path, 0700)
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
	d := walk.Dir{All: true, Root: root}
	d.SetWalkDepth(-1)
	d.SetWalkException("")
	found, err := d.GetChildItem()
	if err != nil {
		fmt.Println(err)
	}
	return found
}
