package node_types_examples
import (
	"testing"
	"strconv"
	"reflect"
	"fmt"
	"encoding/json"
	"github.com/h4ck3rm1k3/gogccintro/tree"
	//"github.com/h4ck3rm1k3/gogccintro/models"
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
	}
	str_ref_fields := []string{
		"AttrsType",
	}

	ref_fields := []string{
		"RefsArgs",
		//"RefsScpe",
		"RefsArgt",
		"RefsBody",
		"RefsBpos",
		"RefsChan",
		"RefsChain",
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
	//fmt.Printf("map %v\n", treemap.Nodes)
	for _,v := range treemap.Nodes {
		s:=treemap.FindName(v)
		fmt.Printf("node id %d %s name:%s\n", v.NodeID,v.NodeType,s)
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
					s2:=treemap.FindName(o)
					fmt.Printf("\treflect %d %s %v %d %v %s\n", k,fn,rid,o.NodeID, o.NodeType, s2)
				} else{
					fmt.Printf("\treflect %d %s %v NULL\n", k,fn,rid)
				}
			}			
		}
		
		for k,fn := range str_ref_fields {
			objValue := reflect.ValueOf(v).Elem()
			field := objValue.FieldByName(fn)
			rid := field.String()
			if rid == "" {

			} else {
			
				i, err := strconv.Atoi(rid)			
				if err != nil {
					fmt.Printf("\treflect2 %d %s %v %sERR\n", k,fn,rid, err)
				} else {
					o:=treemap.Nodes[i]
					if o != nil {
						s2:=treemap.FindName(o)
						fmt.Printf("\treflect2 %d %s %v %d %v %s\n", k,fn,rid,o.NodeID, o.NodeType, s2)
					} else{
					fmt.Printf("\treflect2 %d %s %v NULL\n", k,fn,rid)
					}
				}
			}
		}	
	}
}
