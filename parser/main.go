package main

import (
	"fmt"
		"strconv"
	//"filepath"
	"strings"
	filepath "path"
	"log"
	"os"
	"io"
	"io/ioutil"
)

func Parse2(filename string, b string) (interface{}, error){
	calc := &GccNode{Buffer: b}
	calc.Init()
	if err := calc.Parse(); err != nil {
		log.Fatal(err)
	}
	return nil,nil
}

func ParseReader2(filename string, r io.Reader) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse2(filename, string(b))
}

func ParseFile2(filename string) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader2(filename, f)
}

func main() {


	//func TestOne(t *testing.T) {
	
	files := testTUFiles()
	for _, file := range files {
		
		
		//pgot, err := ParseFile(file,Debug(true))
		os.Remove("file.tu")
		os.Symlink(file, "file.tu")
		_, err := ParseFile2(file)
		// pgot

		
		if err != nil {
			fmt.Printf("In %s\n",file)
			fmt.Printf("err %s\n",err)
			//t.Errorf("%s: pigeon.ParseFile: %v", file, err)
			//_, err := ParseFile2(file)

			//fmt.Printf("err %s\n",err)
			continue
		} else {
			//fmt.Printf("End %s\n",file)
			//fmt.Printf("got %#v\n",pgot)
		}
		
		
	}
}

func testTUFiles() []string {
	const rootDir = "/home/mdupont/experiments/gcc_py_introspector/tests"

	fis, err := ioutil.ReadDir(rootDir)
	if err != nil {
		fmt.Printf("Error:%s",err)
		panic(err)
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

func (e *parseError) Error() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "Error\nfile.tu:%d:  parse error near %v\n (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf(format,
			translations[begin].line,

			
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return error
}
