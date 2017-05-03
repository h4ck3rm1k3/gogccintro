package main

import "encoding/gob"
import "bytes"
import "fmt"

func Gob(t*TestNode){

	fmt.Printf("GOB:%#v\n",t)
	
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)

	// Encoding the map
	err := e.Encode(t)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TODOBUF:%s\n",b)
}
