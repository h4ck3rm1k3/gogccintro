package main
import "github.com/pkg/profile"

//func main() {

	
import (
	"fmt"
	"runtime"
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
	"flag"
	//"github.com/jeffail/tunny"
	"github.com/goinggo/workpool"
)

func Parse2(filename string, b string,args*ParserGlobal) (*GccNodeTest, error){
	//defer profile.Start(profile.TraceProfile).Stop()

	calc := &GccNodeTest{Buffer: b}
	calc.Init(filename, args)
	if err := calc.Parse(); err != nil {
		log.Fatal(err)
	}

	return calc,nil
}

func ParseReader2(filename string, r io.Reader,args*ParserGlobal) (*GccNodeTest, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse2(filename, string(b),args)
}

func ParseFile2(filename string,args*ParserGlobal) (*GccNodeTest, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader2(filename, f, args)
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

func ProcessFile (cfile string, args *ParserGlobal) {
	os.Remove("file.tu")
	os.Symlink(cfile, "file.tu")
	tree, err := ParseFile2(cfile,args)
	
	if err != nil {
		fmt.Printf("In %s\n",cfile)
		fmt.Printf("err %s\n",err)
		
	} else {
		fmt.Printf("OK %s\n",cfile)
		file.Filename = &cfile
		tree.Execute();
		
	}
	resetFile()
}

type FileArgs struct {
	File string 
	Globals  *ParserGlobal
}


func (t * FileArgs) DoWork (workRoutine int) {
	fmt.Printf("process %d\n", workRoutine)

	fmt.Printf("the process %s\n", t.File)
	ProcessFile(
		t.File,
		t.Globals,
	)
	fmt.Printf("after process %s\n", t.File)
}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs *3) // numCPUs hot threads + one for async tasks.
	defer profile.Start(profile.CPUProfile).Stop()

	//pool, _ := tunny.CreatePool(numCPUs, process).Open()
	workPool := workpool.New(runtime.NumCPU(), 800)
	
	StateLookup = CreateStateLookup()

	c := NewConsumer()
	args := NewParser(c)	
	//func TestOne(t *testing.T) {

	var debug_level = flag.Int("debuglevel", 0, "debug level")

	var scan_dir = flag.String("directory", "", "scan directory")
	var read_file = flag.String("file", "", "scan file")
	
	flag.Parse()

	if *debug_level != 0 {
		args.DebugLevel = *debug_level
	}
	// if len(os.Args) < 2 {
	// 	name := os.Args[0]
	// 	fmt.Printf("Usage: %v \"Test Directory\"\n", name)
	// 	os.Exit(1)
	// }
	// dir := os.Args[1]

	if (*scan_dir != "") {
		files := testTUFiles(*scan_dir)
		for _, cfile := range files {

			a := FileArgs {
				File: cfile,
				Globals :args,
			}

			result2 := workPool.PostWork("routine",&a)
			fmt.Printf("res %#v\n", result2)
			
		}
	}

	if (*read_file != "") {
		fmt.Printf("read just one file %s", *read_file)
		//ProcessFile(*read_file, args)
		a := FileArgs {
			File: *read_file,
			Globals :args,
		}
		result2 := workPool.PostWork("routine",&a)
		fmt.Printf("res %#v\n", result2)

	}

	workPool.Shutdown("routine")
	
	// for i,_ := range(args.Attrnames) {
	//  	fmt.Printf("attrnames %s\n",i)
	// }
	fmt.Printf("report\n")
	c.Report()
}
