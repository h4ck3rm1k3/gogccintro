+++
date = "2017-05-15T05:26:49-04:00"
draft = false
title = "TreeDump"

+++

Example generating a call graph from the a standard gcc .tu tree dump.

1. Compile with gcc and get the tu files

    CFLAGS=-O0 -fdump-translation-unit -save-temps
    CPPFLAGS=-O0 -fdump-translation-unit -save-temps
    CXXFLAGS=-O0 -fdump-translation-unit -save-temps

2. convert the tu file to a nquads

    ./parser -file ~/experiments/gcc-1/build/gcc/tree-dump.c.001t.tu -consumer grapht

3. convert to a pquads for load speed

    GOPATH=/home/mdupont/gocode ~/gocode/src/github.com/cayleygraph/cayley/cmd/cayley/cayley conv -i ~/experiments/gcc-1/build/gcc/tree-dump.c.001t.tu.nq -o ~/experiments/gcc-1/build/gcc/tree-dump.c.001t.tu.pq

# query

    nts = g.Morphism().In("Contained")

    nts2 = g.Morphism().In([ "0",  "1",  "10",  "100",  "101", "102",  "103",  "104",      "105",  "106",  "107",  "108",   "109",  "11",   "110",  "111",   "112",  "113",  "114",  "115",  "116",  "117",   "118",  "119",  "12",  "120",  "121",  "122",  "123",  "124",  "125",  "126",  "127",  "128",  "129",  "13",  "130",  "131",  "132",  "133",  "134",  "135",  "136",  "137",  "138",  "139",  "14",  "140",  "141",  "142",  "143",  "144",  "145",  "146",  "147",  "148",  "149",  "15",  "150",  "151",  "152",  "153", "154",  "155",  "156",  "157",  "158",  "159",  "16",  "160",  "161",  "162",  "163",  "164",  "165",  "166",  "167",  "168",  "169",  "17",  "170",  "171",  "172",  "173",  "174",  "175",  "176",  "177",  "178",  "179",  "18",  "180",  "181",  "182",  "183",  "184",  "185",  "186",  "187",  "188",  "189",  "19",  "190", "191",  "192",  "193",  "194",  "195",  "196",  "197",  "198",  "199",  "2",  "20",  "200",  "201",  "202",  "203",  "204",  "205",  "206",  "207",  "208",  "209",  "21",  "210",  "211",  "212",  "213",  "214",  "215",  "22",  "23",  "24",  "25",  "26",  "27",  "28",  "29",  "3",  "30",  "31",  "32",  "33",  "34",  "35",  "36",  "37",  "38",  "39",  "4",  "40",  "41",  "42",  "43",  "44",  "45",  "46",  "47",  "48",  "49",  "5",  "50",  "51",  "52",  "53",  "54",  "55",  "56",  "57",  "58",  "59",  "6",  "60",  "61",  "62",  "63",  "64",  "65",  "66",  "67",  "68",  "69",  "7",  "70",  "71",  "72",  "73",  "74",  "75",  "76",  "77",  "78",  "79",  "8",  "80",  "81",  "82",  "83",  "84",  "85",  "86",  "87",  "88",  "89",  "9",  "90",  "91",  "92",  "93",  "94",  "95",  "96",  "97",  "98", "99",       "args","body","bpos","cnst","cond","csts","dcls","decl","else","expr","flds","fn","fncs","idx","init","inst","labl","low","max","min","mngl","name","op_0","op_1","op_2","prms","purp","retn","rslt","spcs","srcp","strg","then","val","valu","vars","vfld"])
	start =  g.Morphism().Has("nodetype","function_decl").Out("name").Out("strg").Tag("source")
	finish = g.Morphism().In("body").Has("nodetype").Out("name").Out("strg").Tag("target")
	g.V()
	.Follow(start)
	.FollowRecursive(nts)
	.Follow(finish)
	    .All();

# json

![json](test.json)

# convert json to dot

  unique |.[] | "\"" + .source + "\"" + "->" + "\"" + .target + "\"" + ";\n"

# dot
![json](test.dot)

# png
![CallGraph](test.png)

# svg
![CallGraph](test.svg)


