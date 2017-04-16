package main
import "fmt"
import "strings"

type GccNodeTest struct {
	Buffer string;
	Attrnames map[string]int;
}

func (t *GccNodeTest) Init(){
	fmt.Printf("init\n")
	t.Attrnames =make(map[string]int);
}

type Token int 
const (
	START Token = iota
	EOF
	AT
	NODEID
	EXPECT_NODE_TYPE
	NODETYPE
	EXPECT_ATTRNAME
	ATTRNAME
	EXPECT_ATTRVALUE
	ATTRVALUE
	AFTERATTRVALUE
	AFTERATTRVALUE2
	NEWLINE
	COLON
	DATA
)
func isIdent(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func (t*GccNodeTest) Parse() (* int ){
	//r := ni
	//fmt.Printf("parser %s\n",t.Buffer)

	var state Token
	state = START
	
	pos := 0
	var token string


	var nodeid string
	var nodetype string
	var attrname string
	//var attrvalue string
	var attrvalues []string

	var nodeids map[string]int =make(map[string]int);
	var nodetypes map[string]map[string]int =make(map[string]map[string]int);
	//var nodeids map[string]int =make(map[string]int);
	//var map[string]int attrnames;
	
	for _,c := range(t.Buffer) {
		//c := t.Buffer[pos]
		// first split out statements
		if state == START {
			if (c == '@') {
				state = AT
				token =""
			}
		} else if state == AT {
			
			if isDigit(c) {
				state = NODEID
				token = token + string(c)
			}
		} else if state == NODEID {
			if isDigit(c) {
				token = token + string(c)
			} else if c == ' ' {
				nodeid = token
				nodeids[nodeid]=1
				//fmt.Printf("found node id %s\n", nodeid)
				token = ""
				state = EXPECT_NODE_TYPE
			}
		} else if state == EXPECT_NODE_TYPE {
			if c == ' ' {
				token = ""
			} else {
				token = token + string(c)
				state = NODETYPE
			}
		} else if state == NODETYPE {
			if c == ' ' {
				nodetype = token
				//fmt.Printf("found node type %s %s\n", nodeid, nodetype)
				if val, ok :=nodetypes[nodetype]; ok {
					val[nodeid]=1
				} else {
					nodetypes[nodetype]=make(map[string]int)
					nodetypes[nodetype][nodeid]=1
				}
				token = ""
				state = EXPECT_ATTRNAME
			} else {
				token = token + string(c)				
			}
		} else if state == EXPECT_ATTRNAME {
			if c == ' ' {
				token = ""
			} else {
				token = token + string(c)
				state = ATTRNAME
			}
		} else if state == ATTRNAME {
			if c == ' ' {
				if (token == "op") {
					token = token + string(c)				
				} else {
					//size := utf8.RuneCountInString(token)
					if strings.HasSuffix(token,":") {
						attrname = token
						//fmt.Printf("found attrname %s %s %s\n", nodeid, nodetype, attrname)
						t.Attrnames[attrname]=t.Attrnames[attrname]+1
						token = ""
						attrvalues=make([]string,0)
						state = EXPECT_ATTRVALUE
					} else {
						token = token + string(c)
					}
				}
			} else {
				token = token + string(c)				
			}
			
		} else if state == EXPECT_ATTRVALUE {
			if c == ' ' {
				token = ""
			} else {
				state = ATTRVALUE
				token = token + string(c)				
			}
		} else if state == ATTRVALUE {
			if c == ' ' {

				if strings.HasPrefix(token,"@") {
					// prefix
					attrvalues=append(attrvalues,token)
					//fmt.Printf("found attrvalue @ %s %s %s %s\n", nodeid, nodetype,attrname,						attrvalues)
					token = ""
					state = AFTERATTRVALUE2
				} else if strings.HasSuffix(token,":") {
					// it is an attribute name
					attrname = token
					//fmt.Printf("found attrname %s %s %s\n", nodeid, nodetype, attrname)
					t.Attrnames[attrname]=t.Attrnames[attrname]+1
					attrvalues=make([]string,0)
					token = ""
					state = EXPECT_ATTRVALUE				
				} else {
					if attrname == "note:" {
						l := len(attrvalues)
						if (l > 1) {
							//fmt.Printf("check len %d\n", l)
							//fmt.Printf("check values %s\n", attrvalues)
							v :=attrvalues[l-1]
							//fmt.Printf("check note %d %s\n", l, v)
							if v == "operator" {
								attrvalues=append(attrvalues,token)
								token = ""
								// state on state
								state = AFTERATTRVALUE
							}
						} else {
							attrvalues=append(attrvalues,token)
							//fmt.Printf("found first note %s %s %s %s\n", nodeid, nodetype,attrname,								attrvalues)
							token = ""
							state = AFTERATTRVALUE
						}
					} else {
						//attrvalue = token
						attrvalues=append(attrvalues,token)
						//fmt.Printf("found attrvalue %s %s %s %s\n", nodeid, nodetype,attrname,							attrvalues)
						token = ""
						state = AFTERATTRVALUE
					}
				}
			} else {
				token = token + string(c)				
			}
		} else if state == AFTERATTRVALUE {
			if c == '\n' {
				state = NEWLINE
			} else if c == ' ' {
				state = AFTERATTRVALUE
			} else {
				state = ATTRVALUE
				token = token + string(c)
			}
			
		} else if state == AFTERATTRVALUE2 {
			if c == '\n' {
				state = NEWLINE
			} else if c == ' ' {
				state = AFTERATTRVALUE2
			} else {
				state = ATTRNAME
				token = token + string(c)
			}
			
		} else if state == NEWLINE {
			// new statement
			if c == '@' {				
				state = NODEID
			} else {
				state = EXPECT_ATTRNAME	
			}
		}
		
		
		pos = pos + 1
	} // while 

	return nil
}

func (t*GccNodeTest) Execute() (* int ){
	r := 0
	//fmt.Printf("exec %#v\n",t)
	return &r
}

func (t*GccNodeTest) Report() {
	for i,_ := range(t.Attrnames) {
	 	fmt.Printf("attrnames %s\n",i)
	}

}
