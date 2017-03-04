package node_types_examples
import (
	"testing"
	"reflect"
	//"gopkg.in/oleiade/reflections.v1"
	"fmt"
	"encoding/json"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	_ "github.com/h4ck3rm1k3/gogccintro/models"
	"io/ioutil"
	//"os"
)

func TestLoad(*testing.T){
	str_fields := []string{
		"AttrsString",
		"AttrsTypeName",
		"AttrsNote",
		"AttrsTypeSize",
		"AttrsAddr",
		"AttrsType",
	}
	ref_fields := []string{
		"RefsArgs",
		"RefsScpe",
		"RefsArgt",
		"RefsBody",
		"RefsBpos",
		"RefsChan",
		"RefsCnst",
		"RefsCond",
		"RefsCsts",
		"RefsDecl",
		"RefsDomn",
		"RefsE",
		"RefsElts",
		"RefsExpr",
		"RefsFlds",
		"RefsFn",
		"RefsIdx",
		"RefsInit",
		"RefsLabl",
		"RefsLow",
		"RefsMax",
		"RefsMin",
		"RefsMngl",
		"RefsName",
		"RefsOp0",
		"RefsOp1",
		"RefsOp2",
		"RefsPrms",
		"RefsPtd",
		"RefsPurp",
		"RefsRefd",
		"RefsRetn",
		"RefsSize",
		"RefsType",
		"RefsUnql",
		"RefsVal",
		"RefsValu",
		"RefsVars",
	}

	const filename = "funct_decl_key_get_conv_rpc_auth.json"
	
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("err %v\n", err)
		panic(err)
	}
	//fmt.Printf("content %v\n", content)

	treemap := tree.NewTreeMap(0)
	//Do something
	err = json.Unmarshal(content, &treemap.Nodes)
	if err != nil {
		fmt.Printf("err %v\n", err)
		panic(err)
	}
	fmt.Printf("map %v\n", treemap.Nodes)
	for _,v := range treemap.Nodes {
		fmt.Printf("node id %d %s\n", v.NodeID,v.NodeType)
		//fmt.Printf("map %s %v\n", k, v)

		for k,fn := range str_fields {
			objValue := reflect.ValueOf(v).Elem()
			field := objValue.FieldByName(fn).String()
			if field != "" {
				fmt.Printf("\tField %d %s : %v\n", k,fn,field)
			}

		}	

		for k,fn := range ref_fields {
			objValue := reflect.ValueOf(v).Elem()
			field := objValue.FieldByName(fn)
			valid := field.FieldByName("Valid").Bool()
			rid := field.FieldByName("Int64").Int()			
			if valid {

				o:=treemap.Nodes[int(rid)]
				if o != nil {
					fmt.Printf("\treflect %d %s %v %d %v\n", k,fn,rid,o.NodeID, o.NodeType)
				} else{
					fmt.Printf("\treflect %d %s %v NULL\n", k,fn,rid)
				}
			}
		
		
		}	
	}
}
