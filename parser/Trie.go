package main;

import (
	//"github.com/anknown/ahocorasick"
	//"github.com/cloudflare/ahocorasick"
	"github.com/derekparker/trie"
	//"bytes"
	"fmt"
)

type Machine struct {
	Machine *trie.Trie
	Dict[] string
	MaxLen int
}

//type [][]rune DictType

func GenericTrieFields(
	dict []string,
	maxlen int,
) (*Machine){
	//fmt.Printf("Build %s", dict)
	m :=  trie.New()
	for _,x := range(dict){
		m.Add(x,1)
	}
	return &Machine{
		Machine: m,
		Dict: dict,
		MaxLen: maxlen,
	}
}

func NewTrieFields(in []string) (*Machine){
	dict := []string{}
	maxlen := 0
	// process the fields
	for _,l := range(in) {
		ln := len(l)
		if ln > maxlen {
			maxlen = ln
		}

		if len(l) == 3 {
			//fmt.Printf("adding2 %s\n", l + " :")
			dict = append(dict, l + " :")
		} else if len(l) == 2 {
			//fmt.Printf("adding3 %s\n", l + "  :")
			dict = append(dict, l + "  :")
		} else if len(l) >= 4 {
			//fmt.Printf("adding2 %s\n", l + ":")
			dict = append(dict, l + ":")
		} else {
			panic(l)
		}
	}
	return GenericTrieFields(dict, maxlen)
}

func NewTrie(in []string) (*Machine){
	maxlen := 0
	dict := []string{}

	for _,l := range(in) {
		ln := len(l)
		if ln > maxlen {
			maxlen = ln
		}
		//fmt.Printf("adding %s\n", l)
		dict = append(dict, l)
	}
	return GenericTrieFields(dict,maxlen)
}

func (t *Machine) MultiPatternSearch(s []byte) (bool){
	//fmt.Printf("looking for %s in %s", s, t.Dict)
	r := t.Machine.HasKeysWithPrefix(string(s))
	return r 
}
func (t *Machine) ExactSearch(s []byte) (bool){
	_,r := t.Machine.Find(string(s))
	return r
	
	// if (len(r) == 1 ){
	// 	fmt.Printf("Exact found: %s = %s\n", r, s)
	// 	return true
	// }
	// fmt.Printf("Not Exact %s\n", r)
	// return false
}
func (t*Machine) PrintOutput(){
	fmt.Printf("Machine %s", t)
}
