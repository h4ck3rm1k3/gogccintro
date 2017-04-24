package lex
import (
	"testing"
	"strings"
	"reflect"
//	"fmt"
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
