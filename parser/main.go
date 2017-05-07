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
	//"io"
	"io/ioutil"
	//"encoding/json"
	"flag"
	//"github.com/jeffail/tunny"
	"github.com/goinggo/workpool"
)

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

func ProcessFile (filename string, args *ParserGlobal) {
	os.Remove("file.tu")
	os.Symlink(filename, "file.tu")
//	tree, err := ParseFile2(cfile,args)
	f, err := os.Open(filename)
	if err != nil {
		//return nil, err
		panic(err)
	}
	defer f.Close()
	//	return ParseReader2(filename, f, args)
	b, err := ioutil.ReadAll(f)
	if err != nil {
		//return nil, err
		panic(err)
	}
	//return Parse2(filename, string(b),args)
		//defer profile.Start(profile.TraceProfile).Stop()

	tree := &GccNodeTest{Buffer: string(b)}
	tree.Init(filename, args)
//	args.StartFile(filename)
	if err := tree.Parse(); err != nil {
		log.Fatal(err)
	}
	//args.EndFile(filename)

	//return calc,nil
	
	if err != nil {
		fmt.Printf("In %s\n",filename)
		fmt.Printf("err %s\n",err)
		
	} else {
		fmt.Printf("OK %s\n",filename)
		//file.Filename = &cfile
		tree.Execute();
		
	}
	//resetFile()
}

type FileArgs struct {
	File string 
	Globals  *ParserGlobal
	Args * ProgramArgs
}


func (t * FileArgs) DoWork (workRoutine int) {
	StateLookup = CreateStateLookup()
	var c TreeConsumer
	
	if *t.Args.consumer== "mem"{
		c = NewConsumer(*t.Args.name)
	} else if *t.Args.consumer== "mem2"{
		c = NewGraphTransactionConsumer(t.File) //*t.Args.name
		
	//	c = NewMemoryGraphConsumer(*t.Args.name)
	} else if *t.Args.consumer== "graph"{
		c = NewGraphTransactionConsumer(t.File) //*t.Args.name
		//c = NewGraphConsumer(t.File) //*t.Args.name
	} else if *t.Args.consumer== "grapht"{
		c = NewGraphTransactionConsumer(t.File) //*t.Args.name
	}
	args2 := NewParser(c)
	t.Globals =args2
	
	if *t.Args.debug_level != 0 {
		args2.DebugLevel = *t.Args.debug_level
	}
	fmt.Printf("process %d\n", workRoutine)

	fmt.Printf("the process %s\n", t.File)
	ProcessFile(
		t.File,
		t.Globals,
	)
	fmt.Printf("after process %s\n", t.File)
	c.Report()
}

type ProgramArgs struct {
	debug_level *int
	scan_dir *string
	consumer *string
	read_file *string
	operation *string
	name  *string
}
	
func main() {
	args := &ProgramArgs {
		debug_level : flag.Int("debuglevel", 0, "debug level"),
		scan_dir : flag.String("directory", "", "scan directory"),
		consumer : flag.String("consumer", "mem", "consumer name"),
		read_file : flag.String("file", "", "scan file"),
		operation : flag.String("operation", "parse", "parse or readcache"),
		name  : flag.String("name", "test", "name of cachekey"),
	}
	
	flag.Parse()
	if *args.operation=="parse" {
		parser_main(args)
	} else if *args.operation=="readcache" {
		read_cache_main(args)
	}
}
	
func parser_main(args *ProgramArgs) {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs *3) // numCPUs hot threads + one for async tasks.
	defer profile.Start(profile.CPUProfile).Stop()

	//pool, _ := tunny.CreatePool(numCPUs, process).Open()
	workPool := workpool.New(runtime.NumCPU(), 800)

	//func TestOne(t *testing.T) {


	// if len(os.Args) < 2 {
	// 	name := os.Args[0]
	// 	fmt.Printf("Usage: %v \"Test Directory\"\n", name)
	// 	os.Exit(1)
	// }
	// dir := os.Args[1]

	if (*args.scan_dir != "") {
		files := testTUFiles(*args.scan_dir)
		for _, cfile := range files {

			fmt.Printf("reading file %s\n", cfile)
			
			a := FileArgs {
				Args: args,
				File: cfile,
				
			}
			
			//workPool.PostWork("routine",&a)
			a.DoWork(1)
			//fmt.Printf("res %#v\n", result2)
			
		}
	}

	if (*args.read_file != "") {
		fmt.Printf("read just one file %s", *args.read_file)
		//ProcessFile(*read_file, args)
		a := FileArgs {
			Args: args,
			File: *args.read_file,
			//Globals :args2,
		}
		result2 := workPool.PostWork("routine",&a)
		fmt.Printf("res %#v\n", result2)

	}

	workPool.Shutdown("routine")
	
	// for i,_ := range(args.Attrnames) {
	//  	fmt.Printf("attrnames %s\n",i)
	// }
	fmt.Printf("report\n")
	
}
