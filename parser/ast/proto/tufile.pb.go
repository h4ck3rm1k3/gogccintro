// Code generated by protoc-gen-go.
// source: tufile.proto
// DO NOT EDIT!

/*
Package astproto is a generated protocol buffer package.

It is generated from these files:
	tufile.proto

It has these top-level messages:
	FileRef
	IntAttr
	NodeRef
	OpNodeRef
	NumberedNodeRef
	Attr
	Node
	File
*/
package astproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TuNodeType int32

const (
	TuNodeType_TUFILE            TuNodeType = 0
	TuNodeType_OpAttr            TuNodeType = 1
	TuNodeType_NTSourceAttr      TuNodeType = 6
	TuNodeType_NTIntAttr         TuNodeType = 7
	TuNodeType_NTIntAttr3        TuNodeType = 8
	TuNodeType_NTIntAttr2        TuNodeType = 9
	TuNodeType_NTSignedIntAttr   TuNodeType = 10
	TuNodeType_NTAddr            TuNodeType = 11
	TuNodeType_NTHex             TuNodeType = 12
	TuNodeType_NTAddrAttr        TuNodeType = 13
	TuNodeType_NTTagAttr         TuNodeType = 14
	TuNodeType_NTBodyAttr        TuNodeType = 15
	TuNodeType_NTLinkAttr        TuNodeType = 16
	TuNodeType_NTNoteAttr        TuNodeType = 17
	TuNodeType_NTAccsAttr        TuNodeType = 18
	TuNodeType_NTQualAttr        TuNodeType = 19
	TuNodeType_NTSignAttr        TuNodeType = 20
	TuNodeType_NTNodeName        TuNodeType = 21
	TuNodeType_NTNodeAttr        TuNodeType = 22
	TuNodeType_NTSpecValue       TuNodeType = 23
	TuNodeType_NTLngtAttr        TuNodeType = 24
	TuNodeType_NTStringAttr      TuNodeType = 25
	TuNodeType_NTRandomSpec      TuNodeType = 27
	TuNodeType_NTOneAttr         TuNodeType = 28
	TuNodeType_NTAttrs           TuNodeType = 29
	TuNodeType_NTAttr            TuNodeType = 30
	TuNodeType_NTStatement       TuNodeType = 31
	TuNodeType_NTNode            TuNodeType = 32
	TuNodeType_NTInteger         TuNodeType = 33
	TuNodeType_NTNodeType        TuNodeType = 34
	TuNodeType_NTNumberedNodeRef TuNodeType = 35
)

var TuNodeType_name = map[int32]string{
	0:  "TUFILE",
	1:  "OpAttr",
	6:  "NTSourceAttr",
	7:  "NTIntAttr",
	8:  "NTIntAttr3",
	9:  "NTIntAttr2",
	10: "NTSignedIntAttr",
	11: "NTAddr",
	12: "NTHex",
	13: "NTAddrAttr",
	14: "NTTagAttr",
	15: "NTBodyAttr",
	16: "NTLinkAttr",
	17: "NTNoteAttr",
	18: "NTAccsAttr",
	19: "NTQualAttr",
	20: "NTSignAttr",
	21: "NTNodeName",
	22: "NTNodeAttr",
	23: "NTSpecValue",
	24: "NTLngtAttr",
	25: "NTStringAttr",
	27: "NTRandomSpec",
	28: "NTOneAttr",
	29: "NTAttrs",
	30: "NTAttr",
	31: "NTStatement",
	32: "NTNode",
	33: "NTInteger",
	34: "NTNodeType",
	35: "NTNumberedNodeRef",
}
var TuNodeType_value = map[string]int32{
	"TUFILE":            0,
	"OpAttr":            1,
	"NTSourceAttr":      6,
	"NTIntAttr":         7,
	"NTIntAttr3":        8,
	"NTIntAttr2":        9,
	"NTSignedIntAttr":   10,
	"NTAddr":            11,
	"NTHex":             12,
	"NTAddrAttr":        13,
	"NTTagAttr":         14,
	"NTBodyAttr":        15,
	"NTLinkAttr":        16,
	"NTNoteAttr":        17,
	"NTAccsAttr":        18,
	"NTQualAttr":        19,
	"NTSignAttr":        20,
	"NTNodeName":        21,
	"NTNodeAttr":        22,
	"NTSpecValue":       23,
	"NTLngtAttr":        24,
	"NTStringAttr":      25,
	"NTRandomSpec":      27,
	"NTOneAttr":         28,
	"NTAttrs":           29,
	"NTAttr":            30,
	"NTStatement":       31,
	"NTNode":            32,
	"NTInteger":         33,
	"NTNodeType":        34,
	"NTNumberedNodeRef": 35,
}

func (x TuNodeType) Enum() *TuNodeType {
	p := new(TuNodeType)
	*p = x
	return p
}
func (x TuNodeType) String() string {
	return proto.EnumName(TuNodeType_name, int32(x))
}
func (x *TuNodeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TuNodeType_value, data, "TuNodeType")
	if err != nil {
		return err
	}
	*x = TuNodeType(value)
	return nil
}
func (TuNodeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AccessType int32

const (
	AccessType_ACC_priv AccessType = 0
	AccessType_ACC_pub  AccessType = 2
	AccessType_ACC_prot AccessType = 3
)

var AccessType_name = map[int32]string{
	0: "ACC_priv",
	2: "ACC_pub",
	3: "ACC_prot",
}
var AccessType_value = map[string]int32{
	"ACC_priv": 0,
	"ACC_pub":  2,
	"ACC_prot": 3,
}

func (x AccessType) Enum() *AccessType {
	p := new(AccessType)
	*p = x
	return p
}
func (x AccessType) String() string {
	return proto.EnumName(AccessType_name, int32(x))
}
func (x *AccessType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AccessType_value, data, "AccessType")
	if err != nil {
		return err
	}
	*x = AccessType(value)
	return nil
}
func (AccessType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type LinkType int32

const (
	LinkType_extern LinkType = 0
	LinkType_static LinkType = 1
)

var LinkType_name = map[int32]string{
	0: "extern",
	1: "static",
}
var LinkType_value = map[string]int32{
	"extern": 0,
	"static": 1,
}

func (x LinkType) Enum() *LinkType {
	p := new(LinkType)
	*p = x
	return p
}
func (x LinkType) String() string {
	return proto.EnumName(LinkType_name, int32(x))
}
func (x *LinkType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LinkType_value, data, "LinkType")
	if err != nil {
		return err
	}
	*x = LinkType(value)
	return nil
}
func (LinkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SignType int32

const (
	SignType_signed   SignType = 0
	SignType_unsigned SignType = 1
)

var SignType_name = map[int32]string{
	0: "signed",
	1: "unsigned",
}
var SignType_value = map[string]int32{
	"signed":   0,
	"unsigned": 1,
}

func (x SignType) Enum() *SignType {
	p := new(SignType)
	*p = x
	return p
}
func (x SignType) String() string {
	return proto.EnumName(SignType_name, int32(x))
}
func (x *SignType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SignType_value, data, "SignType")
	if err != nil {
		return err
	}
	*x = SignType(value)
	return nil
}
func (SignType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type QualType int32

const (
	QualType_c QualType = 0
	QualType_v QualType = 1
	QualType_r QualType = 2
)

var QualType_name = map[int32]string{
	0: "c",
	1: "v",
	2: "r",
}
var QualType_value = map[string]int32{
	"c": 0,
	"v": 1,
	"r": 2,
}

func (x QualType) Enum() *QualType {
	p := new(QualType)
	*p = x
	return p
}
func (x QualType) String() string {
	return proto.EnumName(QualType_name, int32(x))
}
func (x *QualType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(QualType_value, data, "QualType")
	if err != nil {
		return err
	}
	*x = QualType(value)
	return nil
}
func (QualType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type TagType int32

const (
	TagType_struct TagType = 0
	TagType_union  TagType = 1
)

var TagType_name = map[int32]string{
	0: "struct",
	1: "union",
}
var TagType_value = map[string]int32{
	"struct": 0,
	"union":  1,
}

func (x TagType) Enum() *TagType {
	p := new(TagType)
	*p = x
	return p
}
func (x TagType) String() string {
	return proto.EnumName(TagType_name, int32(x))
}
func (x *TagType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TagType_value, data, "TagType")
	if err != nil {
		return err
	}
	*x = TagType(value)
	return nil
}
func (TagType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type NodeType int32

const (
	NodeType_addr_expr             NodeType = 1
	NodeType_array_type            NodeType = 2
	NodeType_asm_expr              NodeType = 3
	NodeType_binfo                 NodeType = 4
	NodeType_boolean_type          NodeType = 5
	NodeType_component_ref         NodeType = 6
	NodeType_constructor           NodeType = 7
	NodeType_enumeral_type         NodeType = 8
	NodeType_error_mark            NodeType = 9
	NodeType_field_decl            NodeType = 10
	NodeType_function_decl         NodeType = 11
	NodeType_function_type         NodeType = 12
	NodeType_identifier_node       NodeType = 13
	NodeType_indirect_ref          NodeType = 14
	NodeType_integer_cst           NodeType = 15
	NodeType_integer_type          NodeType = 16
	NodeType_lshift_expr           NodeType = 17
	NodeType_mem_ref               NodeType = 18
	NodeType_namespace_decl        NodeType = 19
	NodeType_offset_type           NodeType = 20
	NodeType_parm_decl             NodeType = 21
	NodeType_pointer_type          NodeType = 22
	NodeType_real_type             NodeType = 23
	NodeType_record_type           NodeType = 24
	NodeType_reference_type        NodeType = 25
	NodeType_result_decl           NodeType = 26
	NodeType_save_expr             NodeType = 27
	NodeType_string_cst            NodeType = 28
	NodeType_template_type_parm    NodeType = 29
	NodeType_translation_unit_decl NodeType = 30
	NodeType_tree_list             NodeType = 31
	NodeType_truth_andif_expr      NodeType = 32
	NodeType_type_decl             NodeType = 33
	NodeType_typename_type         NodeType = 34
	NodeType_var_decl              NodeType = 35
	NodeType_void_type             NodeType = 36
)

var NodeType_name = map[int32]string{
	1:  "addr_expr",
	2:  "array_type",
	3:  "asm_expr",
	4:  "binfo",
	5:  "boolean_type",
	6:  "component_ref",
	7:  "constructor",
	8:  "enumeral_type",
	9:  "error_mark",
	10: "field_decl",
	11: "function_decl",
	12: "function_type",
	13: "identifier_node",
	14: "indirect_ref",
	15: "integer_cst",
	16: "integer_type",
	17: "lshift_expr",
	18: "mem_ref",
	19: "namespace_decl",
	20: "offset_type",
	21: "parm_decl",
	22: "pointer_type",
	23: "real_type",
	24: "record_type",
	25: "reference_type",
	26: "result_decl",
	27: "save_expr",
	28: "string_cst",
	29: "template_type_parm",
	30: "translation_unit_decl",
	31: "tree_list",
	32: "truth_andif_expr",
	33: "type_decl",
	34: "typename_type",
	35: "var_decl",
	36: "void_type",
}
var NodeType_value = map[string]int32{
	"addr_expr":             1,
	"array_type":            2,
	"asm_expr":              3,
	"binfo":                 4,
	"boolean_type":          5,
	"component_ref":         6,
	"constructor":           7,
	"enumeral_type":         8,
	"error_mark":            9,
	"field_decl":            10,
	"function_decl":         11,
	"function_type":         12,
	"identifier_node":       13,
	"indirect_ref":          14,
	"integer_cst":           15,
	"integer_type":          16,
	"lshift_expr":           17,
	"mem_ref":               18,
	"namespace_decl":        19,
	"offset_type":           20,
	"parm_decl":             21,
	"pointer_type":          22,
	"real_type":             23,
	"record_type":           24,
	"reference_type":        25,
	"result_decl":           26,
	"save_expr":             27,
	"string_cst":            28,
	"template_type_parm":    29,
	"translation_unit_decl": 30,
	"tree_list":             31,
	"truth_andif_expr":      32,
	"type_decl":             33,
	"typename_type":         34,
	"var_decl":              35,
	"void_type":             36,
}

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}
func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (x *NodeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NodeType_value, data, "NodeType")
	if err != nil {
		return err
	}
	*x = NodeType(value)
	return nil
}
func (NodeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Field int32

const (
	Field_accs       Field = 14
	Field_addr       Field = 9
	Field_algn       Field = 8
	Field_args       Field = 20
	Field_argt       Field = 24
	Field_base       Field = 21
	Field_bases      Field = 2
	Field_binf       Field = 46
	Field_body       Field = 11
	Field_bpos       Field = 55
	Field_chain      Field = 54
	Field_chan       Field = 34
	Field_cls        Field = 39
	Field_cnst       Field = 30
	Field_csts       Field = 40
	Field_dcls       Field = 26
	Field_decl       Field = 25
	Field_domn       Field = 51
	Field_elts       Field = 52
	Field_expr       Field = 27
	Field_flds       Field = 37
	Field_fn         Field = 29
	Field_fncs       Field = 28
	Field_high       Field = 7
	Field_idx        Field = 57
	Field_init       Field = 53
	Field_labl       Field = 41
	Field_link       Field = 12
	Field_lngt       Field = 61
	Field_low        Field = 3
	Field_max        Field = 45
	Field_min        Field = 44
	Field_mngl       Field = 48
	Field_name_field Field = 47
	Field_note       Field = 13
	Field_orig       Field = 22
	Field_prec       Field = 5
	Field_prms       Field = 18
	Field_ptd        Field = 38
	Field_purp       Field = 35
	Field_qual       Field = 15
	Field_refd       Field = 36
	Field_retn       Field = 17
	Field_scpe       Field = 50
	Field_sign       Field = 16
	Field_size       Field = 56
	Field_spec       Field = 60
	Field_srcp       Field = 1
	Field_strg       Field = 62
	Field_tag        Field = 10
	Field_type       Field = 42
	Field_unql       Field = 43
	Field_used       Field = 4
	Field_val        Field = 58
	Field_valu       Field = 33
	Field_vars       Field = 31
	Field_vfld       Field = 32
	Field_OP0        Field = 23
	Field_cond       Field = 63
	// all indexed after this
	Field_OPN Field = 101
)

var Field_name = map[int32]string{
	14:  "accs",
	9:   "addr",
	8:   "algn",
	20:  "args",
	24:  "argt",
	21:  "base",
	2:   "bases",
	46:  "binf",
	11:  "body",
	55:  "bpos",
	54:  "chain",
	34:  "chan",
	39:  "cls",
	30:  "cnst",
	40:  "csts",
	26:  "dcls",
	25:  "decl",
	51:  "domn",
	52:  "elts",
	27:  "expr",
	37:  "flds",
	29:  "fn",
	28:  "fncs",
	7:   "high",
	57:  "idx",
	53:  "init",
	41:  "labl",
	12:  "link",
	61:  "lngt",
	3:   "low",
	45:  "max",
	44:  "min",
	48:  "mngl",
	47:  "name_field",
	13:  "note",
	22:  "orig",
	5:   "prec",
	18:  "prms",
	38:  "ptd",
	35:  "purp",
	15:  "qual",
	36:  "refd",
	17:  "retn",
	50:  "scpe",
	16:  "sign",
	56:  "size",
	60:  "spec",
	1:   "srcp",
	62:  "strg",
	10:  "tag",
	42:  "type",
	43:  "unql",
	4:   "used",
	58:  "val",
	33:  "valu",
	31:  "vars",
	32:  "vfld",
	23:  "OP0",
	63:  "cond",
	101: "OPN",
}
var Field_value = map[string]int32{
	"accs":       14,
	"addr":       9,
	"algn":       8,
	"args":       20,
	"argt":       24,
	"base":       21,
	"bases":      2,
	"binf":       46,
	"body":       11,
	"bpos":       55,
	"chain":      54,
	"chan":       34,
	"cls":        39,
	"cnst":       30,
	"csts":       40,
	"dcls":       26,
	"decl":       25,
	"domn":       51,
	"elts":       52,
	"expr":       27,
	"flds":       37,
	"fn":         29,
	"fncs":       28,
	"high":       7,
	"idx":        57,
	"init":       53,
	"labl":       41,
	"link":       12,
	"lngt":       61,
	"low":        3,
	"max":        45,
	"min":        44,
	"mngl":       48,
	"name_field": 47,
	"note":       13,
	"orig":       22,
	"prec":       5,
	"prms":       18,
	"ptd":        38,
	"purp":       35,
	"qual":       15,
	"refd":       36,
	"retn":       17,
	"scpe":       50,
	"sign":       16,
	"size":       56,
	"spec":       60,
	"srcp":       1,
	"strg":       62,
	"tag":        10,
	"type":       42,
	"unql":       43,
	"used":       4,
	"val":        58,
	"valu":       33,
	"vars":       31,
	"vfld":       32,
	"OP0":        23,
	"cond":       63,
	"OPN":        101,
}

func (x Field) Enum() *Field {
	p := new(Field)
	*p = x
	return p
}
func (x Field) String() string {
	return proto.EnumName(Field_name, int32(x))
}
func (x *Field) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Field_value, data, "Field")
	if err != nil {
		return err
	}
	*x = Field(value)
	return nil
}
func (Field) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Spec int32

const (
	Spec_mutable  Spec = 1
	Spec_bitfield Spec = 2
	Spec_pure     Spec = 3
	Spec_virt     Spec = 4
	Spec_register Spec = 5
)

var Spec_name = map[int32]string{
	1: "mutable",
	2: "bitfield",
	3: "pure",
	4: "virt",
	5: "register",
}
var Spec_value = map[string]int32{
	"mutable":  1,
	"bitfield": 2,
	"pure":     3,
	"virt":     4,
	"register": 5,
}

func (x Spec) Enum() *Spec {
	p := new(Spec)
	*p = x
	return p
}
func (x Spec) String() string {
	return proto.EnumName(Spec_name, int32(x))
}
func (x *Spec) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Spec_value, data, "Spec")
	if err != nil {
		return err
	}
	*x = Spec(value)
	return nil
}
func (Spec) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type FileRef struct {
	FileName         *string `protobuf:"bytes,1,req,name=FileName,json=fileName" json:"FileName,omitempty"`
	LineNumber       *int32  `protobuf:"varint,2,req,name=LineNumber,json=lineNumber" json:"LineNumber,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FileRef) Reset()                    { *m = FileRef{} }
func (m *FileRef) String() string            { return proto.CompactTextString(m) }
func (*FileRef) ProtoMessage()               {}
func (*FileRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FileRef) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *FileRef) GetLineNumber() int32 {
	if m != nil && m.LineNumber != nil {
		return *m.LineNumber
	}
	return 0
}

type IntAttr struct {
	// required int64 LValue = 1;
	Value            *string `protobuf:"bytes,2,req,name=Value,json=value" json:"Value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IntAttr) Reset()                    { *m = IntAttr{} }
func (m *IntAttr) String() string            { return proto.CompactTextString(m) }
func (*IntAttr) ProtoMessage()               {}
func (*IntAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IntAttr) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

// message AccsAttr {}
type NodeRef struct {
	NodeId           *int32 `protobuf:"varint,1,req,name=NodeId,json=nodeId" json:"NodeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NodeRef) Reset()                    { *m = NodeRef{} }
func (m *NodeRef) String() string            { return proto.CompactTextString(m) }
func (*NodeRef) ProtoMessage()               {}
func (*NodeRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NodeRef) GetNodeId() int32 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

type OpNodeRef struct {
	OpNo             *int32 `protobuf:"varint,1,req,name=OpNo,json=opNo" json:"OpNo,omitempty"`
	NodeId           *int32 `protobuf:"varint,2,req,name=NodeId,json=nodeId" json:"NodeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *OpNodeRef) Reset()                    { *m = OpNodeRef{} }
func (m *OpNodeRef) String() string            { return proto.CompactTextString(m) }
func (*OpNodeRef) ProtoMessage()               {}
func (*OpNodeRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OpNodeRef) GetOpNo() int32 {
	if m != nil && m.OpNo != nil {
		return *m.OpNo
	}
	return 0
}

func (m *OpNodeRef) GetNodeId() int32 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

type NumberedNodeRef struct {
	Number           *int32 `protobuf:"varint,1,req,name=Number,json=number" json:"Number,omitempty"`
	NodeId           *int32 `protobuf:"varint,2,req,name=NodeId,json=nodeId" json:"NodeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NumberedNodeRef) Reset()                    { *m = NumberedNodeRef{} }
func (m *NumberedNodeRef) String() string            { return proto.CompactTextString(m) }
func (*NumberedNodeRef) ProtoMessage()               {}
func (*NumberedNodeRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NumberedNodeRef) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *NumberedNodeRef) GetNodeId() int32 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

type Attr struct {
	NodeType *TuNodeType `protobuf:"varint,1,req,name=NodeType,json=nodeType,enum=github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.TuNodeType" json:"NodeType,omitempty"`
	AttrName *Field      `protobuf:"varint,2,req,name=AttrName,json=attrName,enum=github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.Field" json:"AttrName,omitempty"`
	// and one of these values, dependant on the type of node
	SourceAttr      *FileRef         `protobuf:"bytes,6,opt,name=SourceAttr,json=sourceAttr" json:"SourceAttr,omitempty"`
	IntAttr         *IntAttr         `protobuf:"bytes,7,opt,name=IntAttr,json=intAttr" json:"IntAttr,omitempty"`
	Access          *AccessType      `protobuf:"varint,18,opt,name=Access,json=access,enum=github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.AccessType" json:"Access,omitempty"`
	AddrAttr        *string          `protobuf:"bytes,13,opt,name=AddrAttr,json=addrAttr" json:"AddrAttr,omitempty"`
	NodeAttr        *NodeRef         `protobuf:"bytes,22,opt,name=NodeAttr,json=nodeAttr" json:"NodeAttr,omitempty"`
	NumberedNodeRef *NumberedNodeRef `protobuf:"bytes,35,opt,name=NumberedNodeRef,json=numberedNodeRef" json:"NumberedNodeRef,omitempty"`
	NodeOpAttr      *OpNodeRef       `protobuf:"bytes,26,opt,name=NodeOpAttr,json=nodeOpAttr" json:"NodeOpAttr,omitempty"`
	// 	optional NodeRef BodyAttr=15; body is a standard node ref
	LinkAttr *LinkType `protobuf:"varint,16,opt,name=LinkAttr,json=linkAttr,enum=github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.LinkType" json:"LinkAttr,omitempty"`
	// optional NodeAttr NodeAttr=22;
	NoteAttr         *string   `protobuf:"bytes,17,opt,name=NoteAttr,json=noteAttr" json:"NoteAttr,omitempty"`
	QualAttr         *QualType `protobuf:"varint,19,opt,name=QualAttr,json=qualAttr,enum=github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.QualType" json:"QualAttr,omitempty"`
	RandomSpec       *bool     `protobuf:"varint,27,opt,name=RandomSpec,json=randomSpec" json:"RandomSpec,omitempty"`
	SignAttr         *SignType `protobuf:"varint,20,opt,name=SignAttr,json=signAttr,enum=github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.SignType" json:"SignAttr,omitempty"`
	SignedIntAttr    *int32    `protobuf:"varint,10,opt,name=SignedIntAttr,json=signedIntAttr" json:"SignedIntAttr,omitempty"`
	StringAttr       *string   `protobuf:"bytes,25,opt,name=StringAttr,json=stringAttr" json:"StringAttr,omitempty"`
	TagAttr          *TagType  `protobuf:"varint,14,opt,name=TagAttr,json=tagAttr,enum=github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.TagType" json:"TagAttr,omitempty"`
	SpecAttr         *Spec     `protobuf:"varint,28,opt,name=SpecAttr,json=specAttr,enum=github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.Spec" json:"SpecAttr,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Attr) Reset()                    { *m = Attr{} }
func (m *Attr) String() string            { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()               {}
func (*Attr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Attr) GetNodeType() TuNodeType {
	if m != nil && m.NodeType != nil {
		return *m.NodeType
	}
	return TuNodeType_TUFILE
}

func (m *Attr) GetAttrName() Field {
	if m != nil && m.AttrName != nil {
		return *m.AttrName
	}
	return Field_accs
}

func (m *Attr) GetSourceAttr() *FileRef {
	if m != nil {
		return m.SourceAttr
	}
	return nil
}

func (m *Attr) GetIntAttr() *IntAttr {
	if m != nil {
		return m.IntAttr
	}
	return nil
}

func (m *Attr) GetAccess() AccessType {
	if m != nil && m.Access != nil {
		return *m.Access
	}
	return AccessType_ACC_priv
}

func (m *Attr) GetAddrAttr() string {
	if m != nil && m.AddrAttr != nil {
		return *m.AddrAttr
	}
	return ""
}

func (m *Attr) GetNodeAttr() *NodeRef {
	if m != nil {
		return m.NodeAttr
	}
	return nil
}

func (m *Attr) GetNumberedNodeRef() *NumberedNodeRef {
	if m != nil {
		return m.NumberedNodeRef
	}
	return nil
}

func (m *Attr) GetNodeOpAttr() *OpNodeRef {
	if m != nil {
		return m.NodeOpAttr
	}
	return nil
}

func (m *Attr) GetLinkAttr() LinkType {
	if m != nil && m.LinkAttr != nil {
		return *m.LinkAttr
	}
	return LinkType_extern
}

func (m *Attr) GetNoteAttr() string {
	if m != nil && m.NoteAttr != nil {
		return *m.NoteAttr
	}
	return ""
}

func (m *Attr) GetQualAttr() QualType {
	if m != nil && m.QualAttr != nil {
		return *m.QualAttr
	}
	return QualType_c
}

func (m *Attr) GetRandomSpec() bool {
	if m != nil && m.RandomSpec != nil {
		return *m.RandomSpec
	}
	return false
}

func (m *Attr) GetSignAttr() SignType {
	if m != nil && m.SignAttr != nil {
		return *m.SignAttr
	}
	return SignType_signed
}

func (m *Attr) GetSignedIntAttr() int32 {
	if m != nil && m.SignedIntAttr != nil {
		return *m.SignedIntAttr
	}
	return 0
}

func (m *Attr) GetStringAttr() string {
	if m != nil && m.StringAttr != nil {
		return *m.StringAttr
	}
	return ""
}

func (m *Attr) GetTagAttr() TagType {
	if m != nil && m.TagAttr != nil {
		return *m.TagAttr
	}
	return TagType_struct
}

func (m *Attr) GetSpecAttr() Spec {
	if m != nil && m.SpecAttr != nil {
		return *m.SpecAttr
	}
	return Spec_mutable
}

type Node struct {
	NodeID           *int32    `protobuf:"varint,1,req,name=NodeID,json=nodeID" json:"NodeID,omitempty"`
	NodeType         *NodeType `protobuf:"varint,2,opt,name=NodeType,json=nodeType,enum=github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.NodeType" json:"NodeType,omitempty"`
	Attrs            []*Attr   `protobuf:"bytes,3,rep,name=attrs" json:"attrs,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Node) GetNodeID() int32 {
	if m != nil && m.NodeID != nil {
		return *m.NodeID
	}
	return 0
}

func (m *Node) GetNodeType() NodeType {
	if m != nil && m.NodeType != nil {
		return *m.NodeType
	}
	return NodeType_addr_expr
}

func (m *Node) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type File struct {
	Filename         *string `protobuf:"bytes,1,req,name=Filename,json=filename" json:"Filename,omitempty"`
	Nodes            []*Node `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *File) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *File) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*FileRef)(nil), "github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.FileRef")
	proto.RegisterType((*IntAttr)(nil), "github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.IntAttr")
	proto.RegisterType((*NodeRef)(nil), "github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.NodeRef")
	proto.RegisterType((*OpNodeRef)(nil), "github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.OpNodeRef")
	proto.RegisterType((*NumberedNodeRef)(nil), "github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.NumberedNodeRef")
	proto.RegisterType((*Attr)(nil), "github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.Attr")
	proto.RegisterType((*Node)(nil), "github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.Node")
	proto.RegisterType((*File)(nil), "github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.File")
	proto.RegisterEnum("github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.TuNodeType", TuNodeType_name, TuNodeType_value)
	proto.RegisterEnum("github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.AccessType", AccessType_name, AccessType_value)
	proto.RegisterEnum("github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.LinkType", LinkType_name, LinkType_value)
	proto.RegisterEnum("github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.SignType", SignType_name, SignType_value)
	proto.RegisterEnum("github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.QualType", QualType_name, QualType_value)
	proto.RegisterEnum("github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.TagType", TagType_name, TagType_value)
	proto.RegisterEnum("github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.Field", Field_name, Field_value)
	proto.RegisterEnum("github.com.h4ck3rm1k3.gogccintro.parser.ast.proto.Spec", Spec_name, Spec_value)
}

func init() { proto.RegisterFile("tufile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x57, 0xef, 0x72, 0x1b, 0x49,
	0x11, 0xb7, 0x64, 0xc9, 0x92, 0xdb, 0xff, 0xda, 0x13, 0xc7, 0x91, 0xf3, 0x57, 0x51, 0x02, 0x18,
	0x01, 0xe2, 0x2e, 0xb9, 0x23, 0xc7, 0x71, 0x07, 0xe5, 0xe4, 0x2e, 0x45, 0xaa, 0x82, 0x02, 0xb2,
	0xee, 0x28, 0x52, 0x54, 0xa9, 0x46, 0xbb, 0xb3, 0xeb, 0x29, 0xef, 0xce, 0x6e, 0x66, 0x66, 0x8d,
	0xc3, 0x93, 0xf1, 0x04, 0xbc, 0x04, 0x2f, 0xc1, 0x67, 0x3e, 0x51, 0xdd, 0xb3, 0x92, 0xe3, 0x00,
	0x1f, 0xd0, 0x27, 0xfd, 0xba, 0xa7, 0xe7, 0xd7, 0x3d, 0x3d, 0xdd, 0xbd, 0x23, 0xd8, 0xf6, 0x55,
	0xa2, 0x33, 0x35, 0x2a, 0x6d, 0xe1, 0x0b, 0xf1, 0x69, 0xaa, 0xfd, 0x59, 0x35, 0x1f, 0x45, 0x45,
	0x3e, 0x3a, 0xfb, 0x2c, 0x3a, 0x7f, 0x6a, 0xf3, 0x4f, 0xcf, 0x9f, 0x8e, 0xd2, 0x22, 0x8d, 0x22,
	0x6d, 0xbc, 0x2d, 0x46, 0xa5, 0xb4, 0x4e, 0xd9, 0x91, 0x74, 0x3e, 0x6c, 0x19, 0x7c, 0x0b, 0x9d,
	0x97, 0x3a, 0x53, 0x13, 0x95, 0x88, 0xdb, 0xd0, 0x25, 0x38, 0x96, 0xb9, 0xea, 0x35, 0xfa, 0xcd,
	0xe3, 0xcd, 0x49, 0x37, 0xa9, 0x65, 0x71, 0x1f, 0xe0, 0xb5, 0x36, 0x6a, 0x5c, 0xe5, 0x73, 0x65,
	0x7b, 0xcd, 0x7e, 0xf3, 0xb8, 0x3d, 0x81, 0x6c, 0xa9, 0x19, 0x3c, 0x80, 0xce, 0x2b, 0xe3, 0x4f,
	0xbc, 0xb7, 0xe2, 0x00, 0xda, 0xdf, 0xcb, 0xac, 0x52, 0x6c, 0xb5, 0x39, 0x69, 0x5f, 0x90, 0x30,
	0x78, 0x08, 0x9d, 0x71, 0x11, 0xb3, 0x9f, 0x43, 0xd8, 0x20, 0xf8, 0x2a, 0x66, 0x2f, 0xed, 0xc9,
	0x86, 0x61, 0x69, 0xf0, 0x0c, 0x36, 0xdf, 0x94, 0x0b, 0x23, 0x01, 0x2d, 0x12, 0x6a, 0x93, 0x56,
	0x51, 0x8e, 0x8b, 0x0f, 0x36, 0x36, 0xaf, 0x6d, 0x3c, 0x81, 0xbd, 0x10, 0x86, 0x8a, 0x3f, 0xf4,
	0x11, 0x62, 0x5d, 0xf8, 0x60, 0xe9, 0x7f, 0x52, 0xfc, 0x0d, 0xa0, 0xc5, 0xd1, 0xff, 0x09, 0xba,
	0x64, 0x30, 0x7d, 0x5f, 0x86, 0x24, 0xec, 0x3e, 0xf9, 0x7a, 0xf4, 0x7f, 0x67, 0x75, 0x34, 0xad,
	0x16, 0x24, 0x93, 0xae, 0xa9, 0x91, 0x98, 0x42, 0x97, 0x5c, 0x70, 0x7e, 0x9b, 0x4c, 0xfd, 0xc5,
	0x0a, 0xd4, 0x2f, 0xb5, 0xca, 0xe2, 0x49, 0x57, 0xd6, 0x4c, 0xe2, 0x2d, 0xc0, 0x69, 0x51, 0xd9,
	0x48, 0x11, 0x77, 0x6f, 0xa3, 0xdf, 0x38, 0xde, 0x7a, 0xf2, 0xe5, 0x4a, 0xbc, 0x5c, 0x05, 0x13,
	0x70, 0x4b, 0x36, 0x31, 0x5d, 0xde, 0x6a, 0xaf, 0xb3, 0x32, 0x71, 0xcd, 0x30, 0xe9, 0xe8, 0xba,
	0x40, 0xbe, 0x83, 0x8d, 0x93, 0x28, 0x52, 0xce, 0xf5, 0x44, 0xbf, 0xb1, 0x62, 0x82, 0x03, 0x01,
	0x27, 0x78, 0x43, 0x32, 0xa6, 0xf2, 0x3d, 0x89, 0x63, 0xcb, 0xd1, 0xee, 0xf4, 0x1b, 0x54, 0xbe,
	0xb2, 0x96, 0xc5, 0xf7, 0xe1, 0x56, 0x79, 0xed, 0x70, 0xe5, 0x93, 0xd4, 0xc5, 0x15, 0xae, 0x94,
	0x79, 0xb3, 0xff, 0xa8, 0xbc, 0xde, 0x23, 0xa6, 0x7f, 0xbe, 0x0a, 0xfd, 0x75, 0xa6, 0xc9, 0x9e,
	0xf9, 0xa8, 0xa8, 0xff, 0x0c, 0x40, 0xf0, 0x4d, 0xc9, 0xe7, 0xb8, 0xcd, 0x8e, 0xbe, 0x5a, 0xc1,
	0xd1, 0xb2, 0xcb, 0x26, 0x60, 0x96, 0x7c, 0xe2, 0x8f, 0xd0, 0x7d, 0xad, 0xcd, 0x39, 0x73, 0x23,
	0x5f, 0xcc, 0xaf, 0x56, 0xe0, 0x26, 0x8a, 0x50, 0xf7, 0x59, 0x4d, 0x46, 0x17, 0x33, 0x2e, 0x7c,
	0x48, 0xfe, 0x7e, 0xb8, 0x18, 0x53, 0xcb, 0xe4, 0xf4, 0x0f, 0x95, 0xcc, 0x78, 0xed, 0xc6, 0xca,
	0x4e, 0x89, 0x22, 0x38, 0x7d, 0x57, 0x93, 0xd1, 0xc0, 0x9a, 0x48, 0x13, 0x17, 0xf9, 0x69, 0xa9,
	0xa2, 0xde, 0x9d, 0x7e, 0xe3, 0xb8, 0x3b, 0x01, 0xbb, 0xd4, 0x90, 0xe3, 0x53, 0x9d, 0x1a, 0x76,
	0x7c, 0xb0, 0xb2, 0x63, 0xa2, 0x08, 0x8e, 0x5d, 0x4d, 0x26, 0x1e, 0xc3, 0x0e, 0x69, 0x55, 0xbc,
	0xe8, 0x1c, 0xe8, 0x37, 0x8e, 0xdb, 0x93, 0x1d, 0xf7, 0xa1, 0x92, 0xc2, 0x3b, 0xf5, 0x56, 0x9b,
	0x94, 0x4d, 0x8e, 0x38, 0x2b, 0xe0, 0x96, 0x1a, 0xea, 0xbc, 0xa9, 0x0c, 0x8b, 0xbb, 0x1c, 0xdd,
	0x2a, 0xf5, 0x3a, 0x95, 0x29, 0x07, 0xd7, 0xf1, 0x81, 0x4a, 0x9c, 0x42, 0x97, 0x0e, 0xcf, 0xb4,
	0x77, 0x99, 0xf6, 0xd9, 0x2a, 0x87, 0x2e, 0x55, 0x34, 0xe9, 0xba, 0x9a, 0x68, 0xf0, 0xf7, 0x06,
	0xb4, 0xa8, 0x9e, 0x96, 0xb3, 0xf5, 0x9b, 0x6b, 0x73, 0xfd, 0x1b, 0x4a, 0xf5, 0x72, 0xa4, 0x36,
	0x57, 0x4e, 0xf5, 0x7f, 0x19, 0xa8, 0xbf, 0x83, 0x36, 0x8d, 0x41, 0xd7, 0x5b, 0xef, 0xaf, 0x1f,
	0x6f, 0xad, 0x74, 0x16, 0x9e, 0x4c, 0x81, 0x65, 0xf0, 0x0e, 0x5a, 0x34, 0x04, 0x17, 0xdf, 0x41,
	0xf3, 0xd1, 0x77, 0x90, 0x64, 0x72, 0x49, 0xee, 0x5d, 0xaf, 0xb9, 0xb2, 0x4b, 0xee, 0xbd, 0xc0,
	0x32, 0xfc, 0xe7, 0x3a, 0xc0, 0xd5, 0xb7, 0x42, 0x00, 0x6c, 0x4c, 0xbf, 0x7b, 0xf9, 0xea, 0xf5,
	0xb7, 0xb8, 0x46, 0x38, 0x34, 0x26, 0x36, 0x04, 0xc2, 0xf6, 0x78, 0x7a, 0x35, 0xe5, 0x71, 0x43,
	0xec, 0xc0, 0xe6, 0x78, 0x5a, 0x17, 0x13, 0x76, 0xc4, 0x2e, 0xc0, 0x52, 0x7c, 0x8a, 0xdd, 0x6b,
	0xf2, 0x13, 0xdc, 0x14, 0x37, 0x60, 0x6f, 0x3c, 0xbd, 0x56, 0x96, 0x08, 0xe4, 0x61, 0x3c, 0xa5,
	0x91, 0x89, 0x5b, 0x62, 0x13, 0xda, 0xe3, 0xe9, 0x6f, 0xd5, 0x25, 0x6e, 0x87, 0xbd, 0x8b, 0x49,
	0x8a, 0x3b, 0xc1, 0x55, 0x5d, 0x8c, 0xb8, 0x1b, 0x96, 0x9f, 0x17, 0xf1, 0x7b, 0x96, 0xf7, 0x82,
	0xbc, 0x18, 0x1c, 0x88, 0x41, 0x5e, 0xf4, 0x3b, 0xee, 0xd7, 0x74, 0x51, 0xe4, 0x58, 0x16, 0x41,
	0x5e, 0xf4, 0x3c, 0xde, 0x08, 0xf2, 0xa2, 0x15, 0xf1, 0x60, 0xb1, 0x3f, 0xe6, 0x77, 0x07, 0xde,
	0xbc, 0x92, 0x79, 0xfd, 0x50, 0xec, 0xc1, 0xd6, 0x78, 0x4a, 0x25, 0xc8, 0x8f, 0x0c, 0xbc, 0x55,
	0x07, 0x60, 0xd2, 0x70, 0xac, 0x5e, 0x9d, 0xac, 0x65, 0x2b, 0xe1, 0x51, 0xd0, 0x5c, 0x4d, 0x03,
	0xbc, 0x13, 0xce, 0xf4, 0xc6, 0x04, 0xce, 0xbb, 0x62, 0x0b, 0x3a, 0xe3, 0x29, 0x61, 0x87, 0xf7,
	0xea, 0xb4, 0xd0, 0xc2, 0xfd, 0xda, 0x99, 0x97, 0x5e, 0xe5, 0xca, 0x78, 0x7c, 0x10, 0x16, 0x29,
	0x1a, 0xec, 0x2f, 0xef, 0x40, 0xa5, 0xca, 0xe2, 0xc3, 0xab, 0x40, 0xe9, 0x2a, 0x71, 0x20, 0x6e,
	0xc2, 0xfe, 0x78, 0xfa, 0xd1, 0x4c, 0xc7, 0x47, 0xc3, 0xcf, 0x01, 0xae, 0x3e, 0x5e, 0x62, 0x1b,
	0xba, 0x27, 0x2f, 0x5e, 0xcc, 0x4a, 0xab, 0x2f, 0x70, 0x8d, 0xe2, 0x60, 0xa9, 0x9a, 0x63, 0xf3,
	0x6a, 0xa9, 0xf0, 0xb8, 0x3e, 0x1c, 0x84, 0xe9, 0xbc, 0x28, 0x13, 0x75, 0xe9, 0x95, 0x35, 0xa1,
	0x4c, 0x9c, 0x97, 0x5e, 0x47, 0xd8, 0x18, 0x3e, 0x0e, 0x33, 0x6d, 0x61, 0x13, 0x26, 0x0e, 0xae,
	0x11, 0x53, 0x65, 0x6a, 0xa9, 0x31, 0x7c, 0x10, 0x46, 0x2e, 0x5b, 0xb5, 0xa1, 0x11, 0xe1, 0x1a,
	0xfd, 0x5c, 0x60, 0x83, 0x7e, 0x2c, 0x36, 0x87, 0x7d, 0x9e, 0x3d, 0x4b, 0x16, 0x6f, 0xab, 0xc8,
	0xe3, 0x1a, 0x95, 0x48, 0x65, 0x74, 0x61, 0xb0, 0x31, 0xfc, 0x57, 0xeb, 0xaa, 0xa5, 0x29, 0x0d,
	0xf4, 0x9d, 0x9d, 0xa9, 0xcb, 0x92, 0x6a, 0x75, 0x17, 0x40, 0x5a, 0x2b, 0xdf, 0xcf, 0x3c, 0xa5,
	0x81, 0x8f, 0x21, 0x5d, 0x1e, 0x56, 0xd7, 0x89, 0x64, 0xae, 0x4d, 0x52, 0x60, 0x8b, 0x6e, 0x65,
	0x5e, 0x14, 0x99, 0x92, 0x26, 0x98, 0xb6, 0xc5, 0x3e, 0xec, 0x44, 0x45, 0x5e, 0x16, 0x46, 0x19,
	0x3f, 0xb3, 0x2a, 0xc1, 0x0d, 0xba, 0x80, 0xa8, 0x30, 0x21, 0x86, 0x82, 0x2a, 0x7d, 0x1f, 0x76,
	0x94, 0xa9, 0x72, 0x65, 0x65, 0x16, 0xb6, 0x71, 0xb1, 0x2b, 0x6b, 0x0b, 0x3b, 0xcb, 0xa5, 0x3d,
	0xc7, 0x4d, 0x92, 0x13, 0x7a, 0x24, 0xcd, 0x62, 0x15, 0x65, 0x08, 0xb4, 0x25, 0xa9, 0x4c, 0xe4,
	0x75, 0x61, 0x82, 0x6a, 0xeb, 0x9a, 0x8a, 0x59, 0xb6, 0xa9, 0x45, 0x74, 0xac, 0x8c, 0xd7, 0x89,
	0x56, 0x76, 0x46, 0xed, 0x89, 0x3b, 0x14, 0xa3, 0x36, 0xb1, 0xb6, 0x2a, 0x0a, 0x01, 0xed, 0x52,
	0x40, 0x3a, 0x5c, 0xf9, 0x2c, 0x72, 0x1e, 0xf7, 0x82, 0x49, 0x50, 0x30, 0x13, 0x92, 0x49, 0xe6,
	0xce, 0x74, 0xe2, 0xc3, 0xa1, 0xf7, 0xe9, 0x5a, 0x73, 0x95, 0x33, 0x81, 0x10, 0x02, 0x76, 0x69,
	0x92, 0xb8, 0x52, 0x46, 0x2a, 0x84, 0x73, 0x83, 0x76, 0x14, 0x49, 0xe2, 0x94, 0x0f, 0x14, 0x07,
	0x94, 0xd3, 0x52, 0xda, 0x3c, 0xac, 0xdf, 0x24, 0x1f, 0x65, 0x41, 0x5e, 0x6a, 0x1f, 0x87, 0x64,
	0x60, 0xd5, 0x22, 0x05, 0xb7, 0x88, 0xc0, 0xaa, 0xa8, 0xb0, 0x71, 0x50, 0xf4, 0xc8, 0x8b, 0x55,
	0x89, 0xb2, 0xca, 0x44, 0x2a, 0xe8, 0x8e, 0x82, 0x91, 0xab, 0x32, 0x1f, 0x68, 0x6f, 0x13, 0x89,
	0x93, 0x17, 0x2a, 0x84, 0x79, 0x87, 0xf2, 0x16, 0xbe, 0x40, 0x7c, 0xb2, 0xbb, 0xe2, 0x10, 0x84,
	0x57, 0x79, 0x99, 0x49, 0x1f, 0x28, 0x66, 0x14, 0x12, 0xde, 0x13, 0x47, 0x70, 0xd3, 0x5b, 0x69,
	0x5c, 0x26, 0x39, 0x7f, 0x95, 0xd1, 0x35, 0xe3, 0x7d, 0x62, 0xf4, 0x56, 0xa9, 0x59, 0xa6, 0x1d,
	0x75, 0xcb, 0x01, 0xa0, 0xb7, 0x95, 0x3f, 0x9b, 0x49, 0x13, 0xeb, 0x24, 0xf8, 0xe1, 0xbe, 0x61,
	0x3a, 0xde, 0xf3, 0x90, 0xee, 0x82, 0x44, 0x4a, 0x4a, 0x88, 0x74, 0x40, 0x35, 0x73, 0x21, 0x6d,
	0x30, 0x78, 0x44, 0xf6, 0x17, 0x85, 0xae, 0x8f, 0xf6, 0x78, 0xf8, 0x8f, 0x16, 0xb4, 0xf9, 0x11,
	0x2c, 0xba, 0xd0, 0x92, 0x51, 0xe4, 0x70, 0x97, 0x11, 0x0d, 0xb2, 0x4d, 0x46, 0x59, 0x6a, 0xb0,
	0xcb, 0xc8, 0xa6, 0x0e, 0x0f, 0x6a, 0xe4, 0xb1, 0x47, 0x68, 0x2e, 0x1d, 0x8d, 0x15, 0x2a, 0x44,
	0xe9, 0x94, 0xc3, 0x26, 0x2b, 0xb5, 0x49, 0x70, 0xc4, 0xa8, 0x88, 0xdf, 0xe3, 0x16, 0xa3, 0xb2,
	0x70, 0xf8, 0x8c, 0x0c, 0xa3, 0x33, 0xa9, 0x0d, 0xfe, 0x82, 0x94, 0xd1, 0x99, 0x34, 0x38, 0x10,
	0x1d, 0x58, 0x8f, 0x32, 0x87, 0x3f, 0x62, 0x95, 0x71, 0x1e, 0xef, 0x33, 0x72, 0xde, 0xe1, 0x31,
	0xa1, 0x98, 0x56, 0x6f, 0x33, 0xa2, 0x33, 0x1c, 0x31, 0x2a, 0x72, 0x83, 0x4f, 0x09, 0xa9, 0xcc,
	0x3b, 0xfc, 0x8c, 0x51, 0xc8, 0x7c, 0x17, 0x5a, 0x49, 0x16, 0x3b, 0xfc, 0x81, 0xd8, 0x80, 0x66,
	0x62, 0xf0, 0x1e, 0x6b, 0x4c, 0xe4, 0xf0, 0x2e, 0xa1, 0x33, 0x9d, 0x9e, 0x61, 0x87, 0x9c, 0xea,
	0xf8, 0x12, 0x7f, 0x49, 0x2a, 0x6d, 0xb4, 0xc7, 0xcf, 0x09, 0x65, 0x72, 0x9e, 0xe1, 0x8f, 0x19,
	0x69, 0x73, 0x8e, 0xdb, 0x8c, 0x4c, 0xea, 0xf1, 0x6b, 0xda, 0x90, 0x15, 0x7f, 0xc1, 0x75, 0x02,
	0xb9, 0xbc, 0xc4, 0x9f, 0x31, 0xd0, 0x06, 0x7f, 0x4a, 0x46, 0xb9, 0x49, 0x33, 0xfc, 0x84, 0x6e,
	0x9d, 0x53, 0xcf, 0x2d, 0x83, 0x3f, 0xa7, 0x15, 0x7a, 0x9d, 0xe1, 0x0e, 0xa1, 0xc2, 0xea, 0x14,
	0x0f, 0x09, 0x95, 0x56, 0x45, 0xd8, 0x0e, 0x28, 0x77, 0x28, 0x88, 0xaa, 0xf4, 0x31, 0xfe, 0x90,
	0x55, 0x95, 0x2d, 0xf1, 0x11, 0x21, 0x7a, 0x7f, 0xe1, 0x1e, 0x21, 0xab, 0x92, 0x18, 0x1f, 0x07,
	0xe4, 0x0d, 0xee, 0x13, 0x72, 0x51, 0xa9, 0xf0, 0x09, 0x23, 0x9d, 0x1a, 0xc4, 0x80, 0xfe, 0xaa,
	0xf0, 0x0b, 0x46, 0x34, 0x9b, 0xbf, 0x62, 0x64, 0xa3, 0x12, 0x1b, 0x8c, 0xbc, 0x4d, 0xf1, 0xd7,
	0xe4, 0xcc, 0xcb, 0x14, 0x81, 0x54, 0x5c, 0x06, 0x43, 0x42, 0x95, 0x79, 0x97, 0xe1, 0x4f, 0x18,
	0x39, 0x15, 0x63, 0x8b, 0xcc, 0x2e, 0x64, 0x86, 0x5f, 0x92, 0x8a, 0xfe, 0x76, 0xe2, 0xc3, 0x80,
	0xac, 0xc3, 0x07, 0x8c, 0x92, 0x2c, 0xc6, 0x3e, 0x99, 0xbd, 0xf9, 0xfd, 0x27, 0x78, 0x8b, 0xef,
	0xac, 0x30, 0x31, 0xfe, 0x26, 0xa8, 0xc6, 0xa8, 0x86, 0x2f, 0xa0, 0xc5, 0xef, 0x43, 0xea, 0xd9,
	0xca, 0xcb, 0x79, 0xa6, 0xb0, 0x41, 0xf5, 0x38, 0xd7, 0x3e, 0x64, 0xa8, 0x59, 0x1f, 0x58, 0xe1,
	0x3a, 0x53, 0x6a, 0xeb, 0xb1, 0x45, 0x16, 0x56, 0xa5, 0xda, 0x79, 0x65, 0xb1, 0xfd, 0x1c, 0xde,
	0x76, 0xa5, 0xf3, 0xfc, 0xb9, 0xff, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0x08, 0x97, 0x94,
	0xa2, 0x0f, 0x00, 0x00,
}
