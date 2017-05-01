package main
var Fields  = [...]string {
	"accs",
	"addr",
	"algn",
	"args","argt","bases","binf","body","bpos","chain","chan","clas","clnp","cnst","cond","csts","decl","domn","elts","expr","flds","fn","fncs","high","idx","init","int","labl","link","lngt","low ","max","min","mngl","name","note","op","op 0","op 1","op 2","op 3","orig","prec","prms","ptd","purp","qual","refd","retn","saturating","scpe","sign","size","spec","srcp","strg","tag","type","-uid","unql","used","val",
	"valu",
	"vars",
	"rslt",
}

var Flags  = [...]string {
	"bitfield", // flag
	"spec", 
}

var Notes = [...] string { //, "note",
	"operator",
	"artificial",
	"ptrmem",
	"member",
	"constructor",
	"destructor",
	"conversion",
	"global init",
	"global fini",
	"pseudo tmpl",
	"thunk",
	"this adjusting",
	"result adjusting",
	"cleanup",
 }

var Access = [...] string { //, "accs",
	"prot",
	"priv",
	"pub",
}

var Operators = [...] string { // "note operator"
	"deref",
	"not",
	"lnot",
	"preinc",
	"predec",
	"plus",
	"minus",
	"addr",
	"and",
	"andassign",
	"assign",
	"call",
	"compound",
	"delete",
	"div",
	"divassign",
	"eq",
	"ge",
	"gt",
	"land",
	"le",
	"lor",
	"lshift",
	"function",
	"lshiftassign",
	"lt",
	"memref",
	"minusassign",
	"mod",
	"modassign",
	"multassign",
	"mult",
	"ne",
	"neg",
	"new",
	"or",
	"orassign",
	"plusassign",
	"pos",
	"postdec",
	"postinc",
	"ref",
	"rshift",
	"rshiftassign",
	"subs",
	"vecdelete",
	"vecnew",
	"xor",
	"xorassign",
}

var Spec = [...] string {   // "spec"
	"virt",
	"mutable",
	"pure",
}

// "link", "static");
var Link = [...] string {   // "spec"
	"static",
	"extern",
}

