package node_types

import (
	"testing"
)

func TestLoadType(*testing.T){
	f:=NodeFactory{}
	f.StartGraph(nil)
	f.EndGraph()
}

func TestGenerateCode(*testing.T){
	GenerateCode()
}
