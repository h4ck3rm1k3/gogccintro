package main

import (
//	"encoding/json"
	"io/ioutil"
	"path/filepath"
//	"reflect"
	"testing"
	"fmt"
	"strings"
)

func TestOne(t *testing.T) {
	files := testTUFiles(t)
	for _, file := range files {
		fmt.Printf("Reading %s\n",file)
		
		//pgot, err := ParseFile(file,Debug(true))
		_, err := ParseFile(file)
		// pgot
		if err != nil {
			fmt.Printf("err %s\n",err)
			//t.Errorf("%s: pigeon.ParseFile: %v", file, err)
			_, err := ParseFile2(file, Debug2(true))
			fmt.Printf("err %s\n",err)
			continue
		} else {
			fmt.Printf("End %s\n",file)
			//fmt.Printf("got %#v\n",pgot)
		}
		
		
	}
}

func testTUFiles(t *testing.T) []string {
	const rootDir = "/home/mdupont/experiments/gcc_py_introspector/tests"

	fis, err := ioutil.ReadDir(rootDir)
	if err != nil {
		t.Fatal(err)
	}
	files := make([]string, 0, len(fis))
	for _, fi := range fis {
		if !strings.Contains(fi.Name(),"#")  {
			if filepath.Ext(fi.Name()) == ".tu" {
				files = append(files, filepath.Join(rootDir, fi.Name()))
			}
		}
	}
	return files
}
