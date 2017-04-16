package main

import (
	"fmt"
	//"github.com/golang/protobuf/proto"
	"strconv"
	//"filepath"
	"strings"
	filepath "path"
	"log"
	"os"
	"io"
	"io/ioutil"
	//"encoding/json"	
)




func Parse2(filename string, b string) (*GccNodeTest, error){
	calc := &GccNodeTest{Buffer: b}
	calc.Init()
	if err := calc.Parse(); err != nil {
		log.Fatal(err)
	}
	return calc,nil
}

func ParseReader2(filename string, r io.Reader) (*GccNodeTest, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse2(filename, string(b))
}

func ParseFile2(filename string) (*GccNodeTest, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader2(filename, f)
}


func testTUFiles(rootDir string) []string {
//	const rootDir = "/home/mdupont/experiments/gcc_py_introspector/tests"

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

func Atoi(s string) int32 {
	var r int32
	r = -1
	l,e:=strconv.Atoi(s)
	if e == nil {
		r =int32(l)
	} else {
		r = -1
		fmt.Printf("error err:%s input:%s got:%d\n",s,e, l)
		panic("error converting")
	}
	return r
}

func Atoi64(s string) int64 {
	var r int64
	r = -1
	l,e:=strconv.Atoi(s)
	if e == nil {
		r =int64(l)
	} else {
		r = -1
		fmt.Printf("error err:%s input:%s got:%d\n",s,e, l)
		panic("error converting")
	}
	return r
}


func main() {


	//func TestOne(t *testing.T) {
	if len(os.Args) < 2 {
		name := os.Args[0]
		fmt.Printf("Usage: %v \"Test Directory\"\n", name)
		os.Exit(1)
	}
	dir := os.Args[1]
	
	files := testTUFiles(dir)
	for _, cfile := range files {
		
		
		//pgot, err := ParseFile(file,Debug(true))
		os.Remove("file.tu")
		os.Symlink(cfile, "file.tu")
		tree, err := ParseFile2(cfile)
		// pgot

		
		if err != nil {
			fmt.Printf("In %s\n",cfile)
			fmt.Printf("err %s\n",err)
			//t.Errorf("%s: pigeon.ParseFile: %v", file, err)
			//_, err := ParseFile2(file)

			//fmt.Printf("err %s\n",err)
			continue
		} else {
			fmt.Printf("OK %s\n",cfile)

			//fmt.Printf("syntax tree\n")
			//tree.PrintSyntaxTree();
			file.Filename = &cfile
			//fmt.Printf("exec\n")
			tree.Execute();
			tree.Report();
			//fmt.Printf("tokens: %#v\n",tree.Tokens())
			//fmt.Printf("got %#v\n",pgot)
		}
		//fmt.Printf("File %s\n",file.String())
		//body, _ := json.Marshal(file)
		//err = ioutil.WriteFile(fmt.Sprintf("%s.json",cfile), body, 0644)
		//if err != nil { panic (err); }

		//data, err := proto.Marshal(&file)
		//if err != nil {
		//log.Fatal("marshaling error: ", err)
		//}
		//err = ioutil.WriteFile(fmt.Sprintf("%s.proto",cfile), data, 0644)
		//if err != nil { panic (err); }
		//fmt.Printf("File : %s\n",body)

		resetFile()
	}
}
