package main;

import (
	"github.com/anknown/ahocorasick"
	"bytes"
	"fmt"
)

type Machine struct {
	Machine *goahocorasick.Machine
	MaxLen int
}

//type [][]rune DictType

func GenericTrieFields(
	dict [][]rune,
	maxlen int,
) (*Machine){
	m := new(goahocorasick.Machine)
	if err := m.Build(dict); err != nil {
		fmt.Println(err)
		return nil
	}
	return &Machine{
		Machine: m,
		MaxLen: maxlen,
	}

}

func NewTrieFields(in []string) (*Machine){
	dict := [][]rune{}
	maxlen := 0
	// process the fields
	for _,l := range(in) {
		ln := len(l)
		if ln > maxlen {
			maxlen = ln
		}		
		if len(l) == 3 {
			dict = append(dict, bytes.Runes( []byte(l + " :")))
		} else if len(l) == 4 {
			dict = append(dict, bytes.Runes( []byte(l + ":")))
		}
	}
	return GenericTrieFields(dict, maxlen)
}

func NewTrie(in []string) (*Machine){
	maxlen := 0
	dict := [][]rune{}
	for _,l := range(in) {
		ln := len(l)
		if ln > maxlen {
			maxlen = ln
		}
		dict = append(dict, bytes.Runes( []byte(l)))
	}
	return GenericTrieFields(dict,maxlen)
}
