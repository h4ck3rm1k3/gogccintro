// register variables to be used during parsing
package main
import "./ast/proto"

var LineNum int32
var FileName string

var FieldName string
var OpNumber int32

var NodeType astproto.NodeType
var NodeNumber int32 // last node number seen
var MainNodeNumber int32 // node number of the statement

var IntVal  string // large int
var HexVal string  //  hex values

const (
	TUnknown = iota 
	TNodeRef = iota 
	TInteger = iota
	THex = iota 
	TString  = iota 
)
var FieldType int32


//var Spec speco

type LNodeType int32


//var NodeRefs map [string]
var last_attr *astproto.Attr
var last_node *astproto.Node
var file astproto.File

func getNode() (*astproto.Node) {
	id :=MainNodeNumber
	if last_node == nil {
		last_node = &astproto.Node{
			NodeID: &id,
		}
	}
	return last_node
}

func clearNode() {
	// when last node is finished clear it
	last_node = nil
}

func resetFile() {
	// when last node is finished clear it
	file=astproto.File{}
	last_node = nil
}
