package main

type TypeRef struct {
}

type Ref struct {
}

type FileObj struct {
}

type ScopeObj struct {
}

type BasicBlock struct {
	// these fields occur with srcp/type in this order descending
	Srcp FileObj;
	Type TypeRef;
	Scope ScopeObj;
	Chain Ref;
	Name Ref;
	Align Ref;
	Size Ref;
	Used Ref;
}


type ParamDecl struct {
	Used Ref;
	Srcp FileObj;
	Align Ref;
	Type TypeRef;
	// 1729:	NT parm_decl 12893456
	// 1806:	CO 'used:' -> 'parm_decl', 12893456
	// 1807:	CO 'parm_decl' -> 'algn:', 12893456
	// 1808:	CO 'srcp:' -> 'parm_decl', 12893456
	// 1809:	CO 'type:' -> 'parm_decl', 12893456

	// optional
	
	// 1810:	CO 'size:' -> 'parm_decl', 12847185
	Size Ref;
	
	// 1815:	CO 'parm_decl' -> 'argt:', 12626102
	ArgType Ref;
	
	// 1817:	CO 'scpe:' -> 'parm_decl', 12197982
	Scope ScopeObj;
	
	// 1826:	CO 'parm_decl'-> 'chain:', 8195605
	Chain Ref;
	
	// 1871:	CO 'parm_decl' -> 'name:', 3535815
	Name Ref;
	
	// 1979:	CO 'parm_decl' -> 'note:', 785261
	Note Ref;
	
	// 1994:	CO 'parm_decl' -> 'lang:', 678787
	Lang Ref;
	
	// 2178:	CO 'parm_decl' -> 'orig:', 115224
	Orig Ref;

}

type FunctionDecl struct{

	// 	   1648:	NT function_decl 5662695
	// 1840:	CO 'name:' -> 'function_decl', 5662695
	// 1841:	CO 'link:' -> 'function_decl', 5662695
	// 1842:	CO 'srcp:' -> 'function_decl', 5662695
	// 1843:	CO 'type:' -> 'function_decl', 5662695
	
	// 1847:	CO 'function_decl' -> 'body:', 5624471
	// 1849:	CO 'function_decl' -> 'chain:', 5384020
	// 1855:	CO 'scpe:' -> 'function_decl', 4973725
	// 1866:	CO 'function_decl' -> 'args:', 4679597
	// 1899:	CO 'note:' -> 'function_decl', 1897628
	// 1919:	CO 'lang:' -> 'function_decl', 1409012
	// 1962:	CO 'function_decl' -> 'accs:', 948389
	// 2095:	CO 'mngl:' -> 'function_decl', 273510
	// 2149:	CO 'orig:' -> 'function_decl', 174134
	// 2370:	CO 'spec:' -> 'function_decl', 20794
}
type ScopeSource struct {
	Scope ScopeObj;
	Srcp FileObj;
}
