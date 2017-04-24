package lex
import (
	"testing"
	"strings"
	"reflect"
	//"fmt"
)

var Fields  = [...]string {
	"accs",
	"addr",
	"algn",
	"args","argt","bases","binf","body","bpos","chain","chan","clas","clnp","cnst","cond","csts","decl","domn","elts","expr","flds","fn","fncs","high","idx","init","int","labl","link","lngt","low ","max","min","mngl","name","note","op","op 0","op 1","op 2","op 3","orig","prec","prms","ptd","purp","qual","refd","retn","saturating","scpe","sign","size","spec","srcp","strg","tag","type","-uid","unql","used","val",
	"valu",
	"vars",
	"rslt",
}

type Foo struct {
	Accs int
	Addr int
	Algn int
	Args int
	Argt int
	Bases int
	Binf int
	Body int
	Bpos int
	Chain int
	Chan int
	Clas int
	Clnp int
	Cnst int
	Cond int
	Csts int
	Decl int
	Domn int
	Elts int
	Expr int
	Flds int
	Fn int
	Fncs int
	High int
	Idx int
	Init int
	Int int
	Labl int
	Link int
	Lngt int
	Low int
	Max int
	Min int
	Mngl int
	Name int
	Note int
	Op int
	Op0 int
	Op1 int
	Op2 int
	Op3 int
	Orig int
	Prec int
	Prms int
	Ptd int
	Purp int
	Qual int
	Refd int
	Retn int
	Saturating int
	Scpe int
	Sign int
	Size int
	Spec int
	Srcp int
	Strg int
	Tag int
	Type int
	Uid int
	Unql int
	Used int
	Val int
	Valu int
	Vars int
	Rslt int
}

func TestDFA(b *testing.T){
	for _,l := range(Fields) {
		_,_ = matchFunc(l)
		_,_ = matchFunc(l)		
	}
}

func BenchmarkSwitch(b *testing.B){
	//c := 0
	h := map[string]string{
		"low " : "Low",
		"op 0" : "Op0",
		"op 1" : "Op1",
		"op 2" : "Op2",
		"op 3" : "Op3",
		"-uid" : "Uid",
	}	
	for _,l := range(Fields) {
		if _,ok := h[l]; ok  {} else{
			
			h[l]=strings.Title(l)
			
		}

		//fmt.Printf("case \"%s\": X.%s=1\n",l,h[l])
	}

	b.ResetTimer()
	X :=  &Foo {
		Accs : 1,
	}

	for i := 0; i < b.N; i++ {
		for _,l := range(Fields) {
			switch (l) {

			case "accs": X.Accs=1
case "addr": X.Addr=1
case "algn": X.Algn=1
case "args": X.Args=1
case "argt": X.Argt=1
case "bases": X.Bases=1
case "binf": X.Binf=1
case "body": X.Body=1
case "bpos": X.Bpos=1
case "chain": X.Chain=1
case "chan": X.Chan=1
case "clas": X.Clas=1
case "clnp": X.Clnp=1
case "cnst": X.Cnst=1
case "cond": X.Cond=1
case "csts": X.Csts=1
case "decl": X.Decl=1
case "domn": X.Domn=1
case "elts": X.Elts=1
case "expr": X.Expr=1
case "flds": X.Flds=1
case "fn": X.Fn=1
case "fncs": X.Fncs=1
case "high": X.High=1
case "idx": X.Idx=1
case "init": X.Init=1
case "int": X.Int=1
case "labl": X.Labl=1
case "link": X.Link=1
case "lngt": X.Lngt=1
case "low ": X.Low=1
case "max": X.Max=1
case "min": X.Min=1
case "mngl": X.Mngl=1
case "name": X.Name=1
case "note": X.Note=1
case "op": X.Op=1
case "op 0": X.Op0=1
case "op 1": X.Op1=1
case "op 2": X.Op2=1
case "op 3": X.Op3=1
case "orig": X.Orig=1
case "prec": X.Prec=1
case "prms": X.Prms=1
case "ptd": X.Ptd=1
case "purp": X.Purp=1
case "qual": X.Qual=1
case "refd": X.Refd=1
case "retn": X.Retn=1
case "saturating": X.Saturating=1
case "scpe": X.Scpe=1
case "sign": X.Sign=1
case "size": X.Size=1
case "spec": X.Spec=1
case "srcp": X.Srcp=1
case "strg": X.Strg=1
case "tag": X.Tag=1
case "type": X.Type=1
case "-uid": X.Uid=1
case "unql": X.Unql=1
case "used": X.Used=1
case "val": X.Val=1
case "valu": X.Valu=1
case "vars": X.Vars=1
			case "rslt": X.Rslt=1
			default:
				panic(l)
			}
		}
	}
}


func BenchmarkReflect(b *testing.B){
	c := 0
	h := map[string]string{
		"low " : "Low",
		"op 0" : "Op0",
		"op 1" : "Op1",
		"op 2" : "Op2",
		"op 3" : "Op3",
		"-uid" : "Uid",
	}	
	for _,l := range(Fields) {
		if _,ok := h[l]; ok  {} else{
			
			h[l]=strings.Title(l)
			//fmt.Printf("adding %s %s %s\n",l,val,h[l])
		}
	}

	b.ResetTimer()
	X :=  &Foo {
		Accs : 1,
	}
	v := reflect.ValueOf(X).Elem()
	//i := reflect.Indirect(v)
	//t := v.Type()
	for i := 0; i < b.N; i++ {
		for _,l := range(Fields) {
			if l2,ok := h[l]; ok  {
				c++				
				//for i := 0; i < t.NumField(); i++ {
				//setTo := CreateA(t.Field(l).Name)
				field := v.FieldByName(l2)
				//if (!field.CanAddr()) {
				//fmt.Printf("check %s -> %s: %s\n",l,l2, field)
				//panic("addr")
				//}
				//if (!field.CanSet()) {
				//	panic("cannot set")
				//}
				
				//fmt.Printf("field %s", field)
				//fieldVal := 
				field.SetInt(1)
			//}
			} else {
				panic(l)
			}
			
		}
	}
}

func BenchmarkDFA(b *testing.B){
	for i := 0; i < b.N; i++ {
		for _,l := range(Fields) {
			_,_ = matchFunc(l)
			
		}
	}
}

func BenchmarkHash(b *testing.B){
	c := 0
	h := make(map [string]int)
	for _,l := range(Fields) {
		h[l]=1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _,l := range(Fields) {
			if _,ok := h[l]; ok  {
				c++
			}
			
		}
	}
}
