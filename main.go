package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	// "path/filepath"
	// "strings"
)

func dirTree(out io.Writer, path string, printFiles bool) error {

	err := traverseDir(out, "", path, printFiles)
	if err != nil {
		return err
	}

	return nil
}

func traverseDir(out io.Writer, tabs string, path string, printFiles bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if !printFiles {
		for i := 0; i < len(files); {
			if !files[i].IsDir() {
				files[i] = files[len(files)-1]
				files = files[:len(files)-1]
			} else {
				i++
			}
		}
	}
	sort.SliceStable(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	for i, file := range files {
		if file.IsDir() {
			if i != len(files)-1 {
				out.Write([]byte(fmt.Sprintln(tabs + "├───" + file.Name())))
				traverseDir(out, tabs+"│\t", path+string(os.PathSeparator)+file.Name(), printFiles)
			} else {
				out.Write([]byte(fmt.Sprintln(tabs + "└───" + file.Name())))
				traverseDir(out, tabs+"\t", path+string(os.PathSeparator)+file.Name(), printFiles)
			}
		} else {
			if file.Name() == ".DS_Store" {
				continue
			}
			var strSize string
			size := file.Size()
			if size == 0 {
				strSize = " (empty)"
			} else {
				strSize = " (" + fmt.Sprint(size) + "b)"
			}
			if i != len(files)-1 {
				out.Write([]byte(fmt.Sprintln(tabs + "├───" + file.Name() + strSize)))
			} else {
				out.Write([]byte(fmt.Sprintln(tabs + "└───" + file.Name() + strSize)))
			}
		}

	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
	// err := dirTree(os.Stdout, ".", false)
	// if err != nil {
	// 	panic(err.Error())
	// }

}
