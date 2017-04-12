package astproto

import "fmt"
import "strconv"

func (t *Node) AddFileRef(file string, line int32) {
	s := Field_srcp
	nt := TuNodeType_NTSourceAttr
	t.Attrs = append(t.Attrs,
		&Attr{
			AttrName: &s,
			NodeType: &nt,
			SourceAttr: &FileRef{
				FileName:   &file,
				LineNumber: &line,
			},
		})
}

func (t *Node) AddNodeRef(field string, node int32) {
	var s Field
	s = -1
	if val, ok := Field_value[field]; ok {
		s = Field(val)
	} else {
		if field == "name" {
			s = Field_name_field
		} else {

			// look for a numeric field
			l, e := strconv.Atoi(field)
			if e == nil {
				//fmt.Printf("numeric field :%d\n",l)
				// TODO add field with number
				var l2 int32
				l2 = int32(l)
				nt := TuNodeType_NTNumberedNodeRef
				t.Attrs = append(t.Attrs,
					&Attr{
						AttrName: &s,
						NodeType: &nt,
						NumberedNodeRef: &NumberedNodeRef{
							Number: &l2,
							NodeId: &node,
						},
					})

			} else {
				fmt.Printf("error err:%s input:%s got:%d\n", s, e, l)

				fmt.Printf("Missing field %s\n", field)
				panic(field)

				panic("Unkown field")
			}
		}
	}

	nt := TuNodeType_NTNodeAttr
	t.Attrs = append(t.Attrs,
		&Attr{
			AttrName: &s,
			NodeType: &nt, //nt:=TuNodeType_NTNodeAttr
			NodeAttr: &NodeRef{
				NodeId: &node,
			},
		})
}

func (t *Node) AddIntField(field string, val string) {
	// size

	if s, ok := Field_value[field]; ok {
		s := Field(s)
		nt := TuNodeType_NTIntAttr
		t.Attrs = append(t.Attrs,
			&Attr{
				AttrName: &s,
				NodeType: &nt,
				IntAttr:  &IntAttr{Value: &val},
			})
	} else {
		fmt.Printf("AddIntField(%s,%d)\n", field, val)
		panic(field)
	}

}

func (t *Node) AddLowIntField(field string, val string) {
	// this could be : @node or -int or 0x444 hex or 222 int
	t.AddIntField(field, val)
}

func (t *Node) AddHexField(field string, val string) {
	t.AddIntField(field, val)
}

func (t *Node) AddIntStringField(field string, val string) {
	t.AddIntField(field, val)
}

func (t *Node) AddStringField(field string, val string) {
	var s Field
	s = -1
	if val, ok := Field_value[field]; ok {
		s = Field(val)
	}
	nt := TuNodeType_NTStringAttr
	t.Attrs = append(t.Attrs, &Attr{AttrName: &s,
		NodeType:   &nt,
		StringAttr: &val})
}

func (t *Node) AddTag(val TagType) {
	nt := TuNodeType_NTTagAttr
	t.AddGeneric("tag", &Attr{NodeType: &nt, TagAttr: &val})
}

func (t *Node) AddAccess(val AccessType) {
	nt := TuNodeType_NTAccsAttr
	t.AddGeneric("accs", &Attr{NodeType: &nt, Access: &val})
}

func (t *Node) AddQual(val QualType) {
	nt := TuNodeType_NTQualAttr
	t.AddGeneric("qual", &Attr{NodeType: &nt, QualAttr: &val})
}

func (t *Node) AddSign(val SignType) {
	nt := TuNodeType_NTSignAttr
	t.AddGeneric("sign", &Attr{NodeType: &nt, SignAttr: &val})
}

func (t *Node) AddNote(val string) {
	nt := TuNodeType_NTNoteAttr
	t.AddGeneric("note", &Attr{NodeType: &nt, NoteAttr: &val})
}

func (t *Node) AddLink(val LinkType) {
	nt := TuNodeType_NTLinkAttr
	t.AddGeneric("link", &Attr{NodeType: &nt, LinkAttr: &val})
}

func (t *Node) AddGeneric(field string, val *Attr) {
	var s Field
	if val, ok := Field_value[field]; ok {
		s = Field(val)
	}
	val.AttrName = &s
	t.Attrs = append(t.Attrs, val)
}

func (t *Node) AddOpNodeRef(field string, opn int32, node int32) {
	var s Field
	s = -1
	if val, ok := Field_value[field]; ok {
		s = Field(val)
	} else {
		if field == "name" {
			s = Field_name_field
		} 
	}
	nt := TuNodeType_NTNumberedNodeRef
	t.Attrs = append(t.Attrs,
		&Attr{
			AttrName: &s,
			NodeType: &nt,
			NumberedNodeRef: &NumberedNodeRef{
				NodeId: &node,
				Number: &opn,
			},
		})
}

func (t *Node) AddSpec(spec Spec) {
	s := Field_spec
	nt := TuNodeType_NTSpecValue
	t.Attrs = append(t.Attrs,
		&Attr{
			AttrName: &s,
			NodeType: &nt,
			SpecAttr: &spec,
		})
}

func (t *File) AddNode(n *Node) {
	t.Nodes = append(t.Nodes, n)
}
