package main

//p1body_type
//"statement_list","P*"
//"bind_expr", "body",type, vars
// switch_expr "body/cond, type
//return_expr expr,type
//cleanup_point_expr/op_0,type
//cond_expr /"op_0/op_1/op_2.type



// levelone_types := [
//	"bind_expr",
//	"statement_list",
//	"cleanup_point_expr",
//	"return_expr",
//	"cond_expr",
//	"switch_expr"
// ]

// // what is under the first function body
// levelone_predicates := [
//	"P0",
//	"P1",
//	"P2",
//	"P3",
//	"body",
//	"cond"
//	"expr",
//	"op_0",
//	"op_1",
//	"op_2",
//	"vars",
// //	"type",
// ]


import (
	"fmt"
	"log"
	"strings"
	"sort"
	//"time"
	"github.com/h4ck3rm1k3/sparql"
)

// prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
// PREFIX gcc: <https://h4ck3rm1k3.github.io/gogccintro/gcc/ontology/2017/05/20/gcc_compiler.owl#>
// SELECT ?p3  WHERE {
//	?name gcc:strg ?namestr.
//		?f14 gcc:name ?name.
//		?f14 rdf:type gcc:function_decl.
//		?f14 gcc:body ?f13.
//		?f13 ?p1 ?f2.
//		?f2 ?p2 ?f3.
//		?f3 ?p3 ?f4.
//		?f4 ?p4 ?f5.
//		?s1 ?p1 ?f5.
//		?s2 ?p2 ?s1.
//		?s3 ?p3 ?s2.
//		?s4 ?p4 ?s3.
//		?s5 ?p5 ?s4.
//		?s6 ?p6 ?s5.
//		?s7 ?p7 ?s6.
//		?s8 ?p8 ?s7.
//		?s9 ?p9 ?s8.
//		?s10 ?p10 ?s9.
//		?s11 ?p11 ?s10.
//		?s12 ?p12 ?s11.
//		?s13 ?p13 ?s12.
//		?s14 ?p14 ?s13.
//		?s15 ?p15 ?s14.
//		?s16 ?p16 ?s15.
//		?s17 ?p17 ?s16.
//		?s18 ?p18 ?s17.
//		?s19 ?p19 ?s18.
//		?s20 ?p20 ?s19.
//		?s21 ?p21 ?s20.
//		?s22 ?p22 ?s21.
//		?s23 ?p23 ?s22.
//		?s24 ?p24 ?s23.
//		?s25 ?p25 ?s24.
//		?s26 ?p26 ?s25.
//		?s27 ?p27 ?s26.
//		?s28 ?p28 ?s27.
//		?s29 ?p29 ?s28.
//		?s30 ?p30 ?s29.
//		?s31 ?p31 ?s30.
//		?s32 ?p32 ?s31.
//		?s33 ?p33 ?s32.
//		?s34 ?p34 ?s33.
//		?s35 ?p35 ?s34.
//		?s36 ?p36 ?s35.
//		?s37 ?p37 ?s36.
//		?s38 ?p38 ?s37.
//		?s39 ?p39 ?s38.
//		?s40 ?p40 ?s39.
//		?s41 ?p41 ?s40.
//		?s42 ?p42 ?s41.
//		?s43 ?p43 ?s42.
//		?s44 ?p44 ?s43.
//		?s45 ?p45 ?s44.
//		?s46 ?p46 ?s45.
//		?s47 ?p47 ?s46.
//		?s48 ?p48 ?s47.
//		?s49 ?p49 ?s48.
//		?s50 ?p50 ?s49.
//		?s51 ?p51 ?s50.
//		?s52 ?p52 ?s51.
//		?s53 ?p53 ?s52.
//		?s54 ?p54 ?s53.
//		?s55 ?p55 ?s54.
//		?s56 ?p56 ?s55.
//		?s57 ?p57 ?s56.
//		?s58 ?p58 ?s57.
//		?s59 ?p59 ?s58.
//		?s60 ?p60 ?s59.
//		?s61 ?p61 ?s60.
//		?s62 ?p62 ?s61.
//		?s63 ?p63 ?s62.
//		?s64 ?p64 ?s63.
//		?s65 ?p65 ?s64.
//		?s66 ?p66 ?s65.
//		?s67 ?p67 ?s66.
//		?s68 ?p68 ?s67.
//		?s69 ?p69 ?s68.
//		?s70 ?p70 ?s69.
//		?s71 ?p71 ?s70.
//		?s72 ?p72 ?s71.
//		?s73 ?p73 ?s72.
//		?s74 ?p74 ?s73.
//		?s75 ?p75 ?s74.
//		?s76 ?p76 ?s75.
//		?s77 ?p77 ?s76.
//		?s78 ?p78 ?s77.
//		?s79 ?p79 ?s78.
//		?s80 ?p80 ?s79.
//		?s81 ?p81 ?s80.
//		?s82 ?p82 ?s81.
//		?s83 ?p83 ?s82.
//		?s84 ?p84 ?s83.
//		?s85 ?p85 ?s84.
//		?s86 ?p86 ?s85.
//		?s87 ?p87 ?s86.
//		?s88 ?p88 ?s87.
//		?s89 ?p89 ?s88.
//		?s90 ?p90 ?s89.
//		?s91 ?p91 ?s90.
//		?s92 ?p92 ?s91.
//		?s93 ?p93 ?s92.
//		?s94 ?p94 ?s93.
//		?s95 ?p95 ?s94.
//		?s96 ?p96 ?s95.
//		?s97 ?p97 ?s96.
//		?s98 ?p98 ?s97.
//		?s99 ?p99 ?s98.
//		?s100 ?p100 ?s99.

//		FILTER ( ?p3 NOT IN (rdf:type))
//	FILTER ( ?p2 NOT IN (rdf:type))
//	FILTER ( ?p1 NOT IN (rdf:type))
//	FILTER ( ?p4 NOT IN (rdf:type))
// }
// group by ?p3

func depth(x int) string {

	skiptypes := []string{
		"owl:NamedIndividual",
		"gcc:boolean_type",
		"gcc:enumeral_type",
		"gcc:binfo",	
		"gcc:baselink",
			"gcc:function_type",			
			"gcc:integer_type",			
			"gcc:pointer_type",
			"gcc:record_type",
			"gcc:reference_type",			
			"gcc:union_type",
		"gcc:void_type"}

	skip := []string{
			"gcc:accs",
			"gcc:algn",
			"gcc:argt",
			"gcc:base",
			"gcc:base",
			"gcc:bfld",
			"gcc:binf",
			"gcc:chain",
			"gcc:chan",
			"gcc:crnt",
			"gcc:ctor",
			"gcc:domn",
			"gcc:elts",
			"gcc:externbody",
			"gcc:int",
			"gcc:lang",
			"gcc:line",
			"gcc:link",
			"gcc:lngt",
			"gcc:note",
			"gcc:prec",
			"gcc:ptd",
			"gcc:scpe",
			"gcc:sign",
			"gcc:size",
			"gcc:spec",
			"gcc:tag",
			"gcc:type",
			"gcc:unql",
			"gcc:used",
			"rdf:type",
	}

	var names string
	var where string
	names = ""
	where = ""

	for y := 1; y < x; y++ {
		pobj   := fmt.Sprintf(" ?N%d_obj",y-1)
		obj   := fmt.Sprintf(" ?N%d_obj",y)
		otype := fmt.Sprintf(" ?N%d_type",y)
		pred  := fmt.Sprintf(" ?N%d_pred",y)
		names = names + otype
//		names = names + obj
		names = names + pred

		where =  where + fmt.Sprintf("%s rdf:type %s. ", obj, otype)
		//where =  where + fmt.Sprintf("FILTER(%s NOT IN ())\n", otype)

				
		where =  where + fmt.Sprintf("FILTER(%s NOT IN (", otype)
		where = where + strings.Join(skiptypes,",")
		where = where + "))\n"


	// wanted types...
			// "abs_expr",
			// "addr_expr",
			// "aggr_init_expr",
			// "array_ref",
			// "arrow_expr",
			// "bind_expr",			
			// "bit_and_expr",
			// "bit_field_ref",
			// "bit_ior_expr",
			// "bit_not_expr",
			// "bit_xor_expr",
			// "break_stmt",
			// "call_expr",
			// "case_label_expr",
			// "cast_expr",
			// "cleanup_point_expr",
			// "component_ref",
			// "compound_expr",
			// "cond_expr",
			// "const_cast_expr",
			// "const_decl",
			// "constructor",
			// "continue_stmt",
			// "convert_expr",
			// "ctor_initializer",
			// "decl_expr",
			// "dl_expr",
			// "do_stmt",
			// "eq_expr",
			// "expr_stmt",
			// "field_decl",
			// "float_expr",
			// "for_stmt",
			// "function_decl",
			// "ge_expr",
			// "goto_expr",
			// "gt_expr",
			// "identifier_node",
			// "if_stmt",
			// "indirect_ref",
			// "init_expr",
			// "integer_cst",
			// "label_decl",
			// "label_expr",
			// "le_expr",
			// "loop_expr",
			// "lshift_expr",
			// "lt_expr",
			// "max_expr",
			// "mem_ref",
			// "minus_expr",
			// "modify_expr",
			// "modop_expr",
			// "mult_expr",
			// "ne_expr",
			// "negate_expr",
			// "non_lvalue_expr",
			// "nop_expr",
			// "nw_expr",
			// "offsetof_expr",
			// "overload",
			// "parm_decl",
			// "plus_expr",
			// "pointer_plus_expr",
			// "postdecrement_expr",
			// "postincrement_expr",
			// "predecrement_expr",
			// "preincrement_expr",
			// "rdiv_expr",
			// "real_cst",
			// "reinterpret_cast_expr",
			// "result_decl",
			// "return_expr",
			// "rshift_expr",
			// "scope_ref",
			// "sizeof_expr",
			// "statement_list",
			// "static_cast_expr",
			// "string_cst",
			// "switch_expr",
			// "target_expr",
			// "template_decl",
			// "template_id_expr",
			// "template_parm_index",
			// "template_type_parm",
			// "trait_expr",
			// "tree_list",
			// "tree_vec",
			// "trunc_div_expr",
			// "trunc_mod_expr",
			// "truth_andif_expr",
			// "truth_not_expr",
			// "truth_orif_expr",
			// "try_catch_expr",
			// "try_finally",
			// "type_decl",
			// "var_decl",
			// "void_cst",
			// "while_stmt",
			// ]

		where =  where + fmt.Sprintf("%s %s %s. ", pobj, pred, obj)
		where =  where + fmt.Sprintf("FILTER(%s NOT IN (",pred)					
		where = where + strings.Join(skip,",")
		where = where + "))\n"

	}

	var s = " prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>\n" +
		"PREFIX gcc: <https://h4ck3rm1k3.github.io/gogccintro/gcc/ontology/2017/05/20/gcc_compiler.owl#>\n" +
"\nSELECT " +
		  "?function "+
		  "?namestr " +
		  "?body_type " +
		  names +
"\nWHERE {" +
		"?name gcc:strg ?namestr.\n" +
		"?function gcc:name ?name.\n" +
		"?function rdf:type gcc:function_decl.\n" +
		"?function gcc:body ?body.\n" +
		"?body rdf:type ?body_type.\n" +
		" FILTER(?body_type NOT IN (owl:NamedIndividual))\n" +
		"?body ?p1 ?N0_obj .\n" +
		" FILTER ( ?p1 NOT IN (rdf:type,gcc:type))\n" +
		where +

	"\n }"

	fmt.Printf("%s\n", s)
	return s
}

func foo( repo * sparql.Repo, x int) {
	res, err := repo.Query(depth(x))

		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("all %s\n",res)
		//fmt.Printf("res %s\n",res.Results)
		// // for c, r := range res.Results {
		// //
		//fmt.Printf("%#v\n",c,r)
		// }
		//fmt.Printf("bindings %s\n",res.Bindings())

		//map[int][string][string]

		for c, r := range res.Bindings() {
			//fmt.Printf("bval %s\n",c)

			cnt := make(map[string]int)// simple count
			n := map[int][]string{}
			for _, v := range r {
				v2 := strings.Replace(v.String(),"https://h4ck3rm1k3.github.io/gogccintro/gcc/ontology/2017/05/20/gcc_compiler.owl#","gcc:",1)
				//fmt.Printf("bval2 %s %d %s\n", c, k, v2)

				if val, ok := cnt[v2]; ok {
					cnt[v2] = val +1
				} else {
					cnt[v2] = 1
				}

			}

			///////////////////////
			for k, v := range cnt {
				n[v] = append(n[v], k)
			}

//			fmt.Printf("%s %v",c, cnt)

			var a []int // what items occur at this frequency
			for k2 := range n {
				a = append(a, k2)
			}
			sort.Sort(sort.Reverse(sort.IntSlice(a)))
			for _, k := range a {
				for _, s := range n[k] {
					fmt.Printf("%s : %s, %d\n",c, s, k)
				}

			}
		}

}

func main() {
	repo, err := sparql.NewRepo("http://localhost:9999/blazegraph/sparql") //	sparql.DigestAuth("dba", "dba"),
	//sparql.Timeout(time.Millisecond*15000),

	if err != nil {
		log.Fatal(err)
	}
	for x := 160; x > 1; x-- {
		//depth(x)
		foo( repo, x)
	}
}
