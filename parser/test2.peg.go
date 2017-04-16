package main

import (
	"./ast/proto"
	"fmt"
	"math"
	"sort"
	"strconv"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleTUFILE
	rulews
	rulews2
	rulews3
	ruleOtherType
	ruleNodeType
	ruleOpAttr
	ruleOpAttr2
	ruleLineNumber
	ruleFileName
	ruleFileRef
	ruleSourceAttr
	ruleLangAttr
	ruleIntAttr
	ruleIntAttrLow
	ruleIntAttr3
	ruleAddr
	ruleAddrAttr
	ruleTagAttr
	ruleBodyAttr
	ruleLinkAttr
	ruleNoteAttr
	ruleAccsAttr
	ruleQualAttr
	ruleSignAttr
	ruleOtherField
	ruleNodeFieldName
	ruleNodeAttr
	ruleSpecValue
	ruleLngtAttr
	ruleRandomSpec
	ruleOneAttr
	ruleAttrs
	ruleAttr
	ruleStatement
	ruleStringAttr
	ruleEOF
	ruleDecimalDigit
	ruleNonZeroDecimalDigit
	ruleNode
	rulePNode
	rulePosInt
	ruleInteger
	ruleHex
	ruleHexValue
	ruleNegInt
	ruleSignedIntAttr
	ruleIntAttr2
	ruleAction0
	ruleAction1
	ruleAction2
	rulePegText
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6
	ruleAction7
	ruleAction8
	ruleAction9
	ruleAction10
	ruleAction11
	ruleAction12
	ruleAction13
	ruleAction14
	ruleAction15
	ruleAction16
	ruleAction17
	ruleAction18
	ruleAction19
	ruleAction20
	ruleAction21
	ruleAction22
	ruleAction23
	ruleAction24
	ruleAction25
	ruleAction26
	ruleAction27
	ruleAction28
	ruleAction29
	ruleAction30
	ruleAction31
	ruleAction32
	ruleAction33
	ruleAction34
	ruleAction35
	ruleAction36
	ruleAction37
	ruleAction38
	ruleAction39
)

var rul3s = [...]string{
	"Unknown",
	"TUFILE",
	"ws",
	"ws2",
	"ws3",
	"OtherType",
	"NodeType",
	"OpAttr",
	"OpAttr2",
	"LineNumber",
	"FileName",
	"FileRef",
	"SourceAttr",
	"LangAttr",
	"IntAttr",
	"IntAttrLow",
	"IntAttr3",
	"Addr",
	"AddrAttr",
	"TagAttr",
	"BodyAttr",
	"LinkAttr",
	"NoteAttr",
	"AccsAttr",
	"QualAttr",
	"SignAttr",
	"OtherField",
	"NodeFieldName",
	"NodeAttr",
	"SpecValue",
	"LngtAttr",
	"RandomSpec",
	"OneAttr",
	"Attrs",
	"Attr",
	"Statement",
	"StringAttr",
	"EOF",
	"DecimalDigit",
	"NonZeroDecimalDigit",
	"Node",
	"PNode",
	"PosInt",
	"Integer",
	"Hex",
	"HexValue",
	"NegInt",
	"SignedIntAttr",
	"IntAttr2",
	"Action0",
	"Action1",
	"Action2",
	"PegText",
	"Action3",
	"Action4",
	"Action5",
	"Action6",
	"Action7",
	"Action8",
	"Action9",
	"Action10",
	"Action11",
	"Action12",
	"Action13",
	"Action14",
	"Action15",
	"Action16",
	"Action17",
	"Action18",
	"Action19",
	"Action20",
	"Action21",
	"Action22",
	"Action23",
	"Action24",
	"Action25",
	"Action26",
	"Action27",
	"Action28",
	"Action29",
	"Action30",
	"Action31",
	"Action32",
	"Action33",
	"Action34",
	"Action35",
	"Action36",
	"Action37",
	"Action38",
	"Action39",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Printf(" ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Printf("%v %v\n", rule, quote)
			} else {
				fmt.Printf("\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(buffer string) {
	node.print(false, buffer)
}

func (node *node32) PrettyPrint(buffer string) {
	node.print(true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	if tree := t.tree; int(index) >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	t.tree[index] = token32{
		pegRule: rule,
		begin:   begin,
		end:     end,
	}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type GccNode struct {
	Buffer string
	buffer []rune
	rules  [90]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *GccNode) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *GccNode) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *GccNode
	max token32
}

func (e *parseError) Error2() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return error
}

func (p *GccNode) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *GccNode) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for _, token := range p.Tokens() {
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:

		case ruleAction1:

		case ruleAction2:

		case ruleAction3:

			fmt.Printf("Other Type: %s\n", buffer[begin:end])

		case ruleAction4:

			s := astproto.NodeType(astproto.NodeType_value[buffer[begin:end]])
			NodeType = s

		case ruleAction5:

			OpNumber = Atoi(buffer[begin:end])
			FieldName = "opn"
			//fmt.Printf("op check1 %s %d\n", string(buffer[begin:end]), OpNumber)

		case ruleAction6:

			OpNumber = Atoi(buffer[begin:end])
			FieldName = "opn"
			//fmt.Printf("op check1 %s %d\n", string(buffer[begin:end]), OpNumber)

		case ruleAction7:

			LineNum = Atoi(buffer[begin:end])

		case ruleAction8:

			FileName = string(buffer[begin:end])

		case ruleAction9:

			//fmt.Printf("FileName and Number: '%s:%d'\n",FileName, LineNum)
			getNode().AddFileRef(FileName, LineNum)

		case ruleAction10:

			getNode().AddStringField("lang", "C")

		case ruleAction11:

			// base
			getNode().AddIntField(FieldName, buffer[begin:end])

		case ruleAction12:

			if FieldType == TNodeRef {
				getNode().AddNodeRef("low", NodeNumber)
			} else if FieldType == TInteger {
				getNode().AddIntField("low", IntVal)
			} else if FieldType == THex {
				getNode().AddHexField("low", HexVal)
			} else {
				fmt.Printf("unkown field type : %s\n", buffer[begin:end])
				fmt.Printf("unkown field type : %s\n", buffer[begin-30:end+30])
				//panic("unkown field")
				getNode().AddHexField("low", buffer[begin:end])
			}

		case ruleAction13:

			getNode().AddIntField(FieldName, buffer[begin:end])

		case ruleAction14:

			// addr :
			getNode().AddHexField("addr", buffer[begin:end])

		case ruleAction15:

			getNode().AddTag(astproto.TagType(astproto.TagType_value[buffer[begin:end]]))

		case ruleAction16:

		case ruleAction17:

			getNode().AddLink(astproto.LinkType(astproto.LinkType_value[buffer[begin:end]]))

		case ruleAction18:

			getNode().AddNote(buffer[begin:end])

		case ruleAction19:

			getNode().AddAccess(astproto.AccessType(astproto.AccessType_value[buffer[begin:end]]))

		case ruleAction20:

			getNode().AddQual(astproto.QualType(astproto.QualType_value[buffer[begin:end]]))

		case ruleAction21:

			getNode().AddSign(astproto.SignType(astproto.SignType_value[buffer[begin:end]]))

		case ruleAction22:

			fmt.Printf("Other Field: %s\n", buffer[begin:end])

		case ruleAction23:

			FieldName = buffer[begin:end]
			//fmt.Printf("set fieldname %s\n",FieldName)

		case ruleAction24:

			if FieldName == "opn" {
				getNode().AddOpNodeRef(FieldName, OpNumber, NodeNumber)
			} else {
				getNode().AddNodeRef(FieldName, NodeNumber)
			}

		case ruleAction25:

			s := astproto.Spec(astproto.Spec_value[buffer[begin:end]])
			getNode().AddSpec(s)

		case ruleAction26:

		case ruleAction27:

			FieldType = TUnknown

		case ruleAction28:

			// clear it
			n := getNode()
			//fmt.Printf("stmt %#v\n",n)
			nt := NodeType // copy it or it will get changed
			n.NodeType = &nt
			file.AddNode(n)
			clearNode()

		case ruleAction29:

			// fmt.Printf("found string : bytes %d\n", end - begin)
			//	fmt.Printf("found string : %s\n", buffer[begin:end])
			getNode().AddStringField(FieldName, buffer[begin:end])
			FieldType = TString

		case ruleAction30:

		case ruleAction31:

		case ruleAction32:

			//s:=buffer[begin:end]
			//fmt.Printf("noderef %s\n",s)
			NodeNumber = Atoi(buffer[begin:end])
			FieldType = TNodeRef

		case ruleAction33:

			MainNodeNumber = NodeNumber

		case ruleAction34:

		case ruleAction35:

			IntVal = buffer[begin:end]
			FieldType = TInteger

		case ruleAction36:

			HexVal = string(buffer[begin:end])
			FieldType = THex

		case ruleAction37:

			// NegInt
			IntVal = buffer[begin:end]
			FieldType = TInteger

		case ruleAction38:

			// int
			//fmt.Printf("signed check %s\n", string(buffer[0:end]))
			//fmt.Printf("signed check %s\n", string(buffer[begin:end]))
			getNode().AddIntStringField(FieldName, buffer[begin:end])

		case ruleAction39:

			// algn
			//fmt.Printf("algn %s\n", string(buffer[begin:end]))
			getNode().AddIntField("algn", buffer[begin:end])

		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
}

func (p *GccNode) Init() {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := tokens32{tree: make([]token32, math.MaxInt16)}
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 TUFILE <- <(ws Statement+ EOF)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				if !_rules[rulews]() {
					goto l0
				}
				{
					position4 := position
					{
						position5 := position
						{
							position6 := position
							if !_rules[ruleNode]() {
								goto l0
							}
							{
								add(ruleAction33, position)
							}
							add(rulePNode, position6)
						}
						add(rulePegText, position5)
					}
					if !_rules[rulews]() {
						goto l0
					}
					{
						position8 := position
						{
							position9 := position
							{
								position10 := position
								{
									position11 := position
									{
										position14, tokenIndex14 := position, tokenIndex
										if c := buffer[position]; c < rune('a') || c > rune('z') {
											goto l15
										}
										position++
										goto l14
									l15:
										position, tokenIndex = position14, tokenIndex14
										if buffer[position] != rune('_') {
											goto l0
										}
										position++
									}
								l14:
								l12:
									{
										position13, tokenIndex13 := position, tokenIndex
										{
											position16, tokenIndex16 := position, tokenIndex
											if c := buffer[position]; c < rune('a') || c > rune('z') {
												goto l17
											}
											position++
											goto l16
										l17:
											position, tokenIndex = position16, tokenIndex16
											if buffer[position] != rune('_') {
												goto l13
											}
											position++
										}
									l16:
										goto l12
									l13:
										position, tokenIndex = position13, tokenIndex13
									}
									add(rulePegText, position11)
								}
								{
									add(ruleAction3, position)
								}
								add(ruleOtherType, position10)
							}
							{
								add(ruleAction4, position)
							}
							add(ruleNodeType, position9)
						}
						add(rulePegText, position8)
					}
					if !_rules[rulews]() {
						goto l0
					}
					{
						position20 := position
						{
							position21 := position
							{
								position22, tokenIndex22 := position, tokenIndex
								if !_rules[ruleOneAttr]() {
									goto l22
								}
								if !_rules[rulews]() {
									goto l22
								}
								{
									position24 := position
								l25:
									{
										position26, tokenIndex26 := position, tokenIndex
										if !_rules[rulews]() {
											goto l26
										}
										if !_rules[ruleOneAttr]() {
											goto l26
										}
										goto l25
									l26:
										position, tokenIndex = position26, tokenIndex26
									}
									add(ruleAttrs, position24)
								}
								if !_rules[rulews]() {
									goto l22
								}
								goto l23
							l22:
								position, tokenIndex = position22, tokenIndex22
							}
						l23:
							add(ruleAttr, position21)
						}
						add(rulePegText, position20)
					}
					if !_rules[rulews]() {
						goto l0
					}
					{
						add(ruleAction28, position)
					}
					add(ruleStatement, position4)
				}
			l2:
				{
					position3, tokenIndex3 := position, tokenIndex
					{
						position28 := position
						{
							position29 := position
							{
								position30 := position
								if !_rules[ruleNode]() {
									goto l3
								}
								{
									add(ruleAction33, position)
								}
								add(rulePNode, position30)
							}
							add(rulePegText, position29)
						}
						if !_rules[rulews]() {
							goto l3
						}
						{
							position32 := position
							{
								position33 := position
								{
									position34 := position
									{
										position35 := position
										{
											position38, tokenIndex38 := position, tokenIndex
											if c := buffer[position]; c < rune('a') || c > rune('z') {
												goto l39
											}
											position++
											goto l38
										l39:
											position, tokenIndex = position38, tokenIndex38
											if buffer[position] != rune('_') {
												goto l3
											}
											position++
										}
									l38:
									l36:
										{
											position37, tokenIndex37 := position, tokenIndex
											{
												position40, tokenIndex40 := position, tokenIndex
												if c := buffer[position]; c < rune('a') || c > rune('z') {
													goto l41
												}
												position++
												goto l40
											l41:
												position, tokenIndex = position40, tokenIndex40
												if buffer[position] != rune('_') {
													goto l37
												}
												position++
											}
										l40:
											goto l36
										l37:
											position, tokenIndex = position37, tokenIndex37
										}
										add(rulePegText, position35)
									}
									{
										add(ruleAction3, position)
									}
									add(ruleOtherType, position34)
								}
								{
									add(ruleAction4, position)
								}
								add(ruleNodeType, position33)
							}
							add(rulePegText, position32)
						}
						if !_rules[rulews]() {
							goto l3
						}
						{
							position44 := position
							{
								position45 := position
								{
									position46, tokenIndex46 := position, tokenIndex
									if !_rules[ruleOneAttr]() {
										goto l46
									}
									if !_rules[rulews]() {
										goto l46
									}
									{
										position48 := position
									l49:
										{
											position50, tokenIndex50 := position, tokenIndex
											if !_rules[rulews]() {
												goto l50
											}
											if !_rules[ruleOneAttr]() {
												goto l50
											}
											goto l49
										l50:
											position, tokenIndex = position50, tokenIndex50
										}
										add(ruleAttrs, position48)
									}
									if !_rules[rulews]() {
										goto l46
									}
									goto l47
								l46:
									position, tokenIndex = position46, tokenIndex46
								}
							l47:
								add(ruleAttr, position45)
							}
							add(rulePegText, position44)
						}
						if !_rules[rulews]() {
							goto l3
						}
						{
							add(ruleAction28, position)
						}
						add(ruleStatement, position28)
					}
					goto l2
				l3:
					position, tokenIndex = position3, tokenIndex3
				}
				{
					position52 := position
					{
						position53, tokenIndex53 := position, tokenIndex
						if !matchDot() {
							goto l53
						}
						goto l0
					l53:
						position, tokenIndex = position53, tokenIndex53
					}
					add(ruleEOF, position52)
				}
				add(ruleTUFILE, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 ws <- <(((&('\n') '\n') | (&('\r') '\r') | (&('\t') '\t') | (&(' ') ' '))* Action0)> */
		func() bool {
			{
				position55 := position
			l56:
				{
					position57, tokenIndex57 := position, tokenIndex
					{
						switch buffer[position] {
						case '\n':
							if buffer[position] != rune('\n') {
								goto l57
							}
							position++
							break
						case '\r':
							if buffer[position] != rune('\r') {
								goto l57
							}
							position++
							break
						case '\t':
							if buffer[position] != rune('\t') {
								goto l57
							}
							position++
							break
						default:
							if buffer[position] != rune(' ') {
								goto l57
							}
							position++
							break
						}
					}

					goto l56
				l57:
					position, tokenIndex = position57, tokenIndex57
				}
				{
					add(ruleAction0, position)
				}
				add(rulews, position55)
			}
			return true
		},
		/* 2 ws2 <- <(' ' Action1)> */
		nil,
		/* 3 ws3 <- <(' '+ '\n' ' '+ Action2)> */
		nil,
		/* 4 OtherType <- <(<([a-z] / '_')+> Action3)> */
		nil,
		/* 5 NodeType <- <(OtherType Action4)> */
		nil,
		/* 6 OpAttr <- <(('o' / 'O') ('p' / 'P') ' ' <[0-9]> Action5)> */
		nil,
		/* 7 OpAttr2 <- <(<[0-9]> Action6)> */
		nil,
		/* 8 LineNumber <- <(<Integer> Action7)> */
		nil,
		/* 9 FileName <- <(<((((&('_') '_') | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') [0-9]) | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') [a-z]) | (&('-') '-') | (&('+') '+') | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]))+ ('.' ((&('_') '_') | (&('a' | 'b' | 'c' | 'd' | 'e' | 'f' | 'g' | 'h' | 'i' | 'j' | 'k' | 'l' | 'm' | 'n' | 'o' | 'p' | 'q' | 'r' | 's' | 't' | 'u' | 'v' | 'w' | 'x' | 'y' | 'z') [a-z]) | (&('A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H' | 'I' | 'J' | 'K' | 'L' | 'M' | 'N' | 'O' | 'P' | 'Q' | 'R' | 'S' | 'T' | 'U' | 'V' | 'W' | 'X' | 'Y' | 'Z') [A-Z]))+)*) / ('<' ('b' / 'B') ('u' / 'U') ('i' / 'I') ('l' / 'L') ('t' / 'T') '-' ('i' / 'I') ('n' / 'N') '>'))> Action8)> */
		nil,
		/* 10 FileRef <- <('_'? <FileName> ':' <LineNumber> ws Action9)> */
		nil,
		/* 11 SourceAttr <- <(<('s' 'r' 'c' 'p')> ':' ws <FileRef>)> */
		nil,
		/* 12 LangAttr <- <(<('l' 'a' 'n' 'g')> ':' ws ('c' / 'C') Action10)> */
		nil,
		/* 13 IntAttr <- <(<((&('P' | 'p') (('p' / 'P') ('r' / 'R') ('e' / 'E') ('c' / 'C'))) | (&('C' | 'c') (('c' / 'C') ('t' / 'T') ('o' / 'O') ('r' / 'R'))) | (&('U' | 'u') (('u' / 'U') ('s' / 'S') ('e' / 'E') ('d' / 'D'))) | (&('B' | 'b') (('b' / 'B') ('a' / 'A') ('s' / 'S') ('e' / 'E') ('s' / 'S'))) | (&('L' | 'l') (('l' / 'L') ('i' / 'I') ('n' / 'N') ('e' / 'E'))))> ws ':' ws <Integer> ws Action11)> */
		nil,
		/* 14 IntAttrLow <- <(<(('l' / 'L') ('o' / 'O') ('w' / 'W'))> ws ':' ws (HexValue / ((&('-') NegInt) | (&('@') Node) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') Integer))) ws Action12)> */
		nil,
		/* 15 IntAttr3 <- <(<(('h' / 'H') ('i' / 'I') ('g' / 'G') ('h' / 'H'))> ws ':' ws <('-'? Integer)> ws Action13)> */
		nil,
		/* 16 Addr <- <(DecimalDigit / Hex)+> */
		nil,
		/* 17 AddrAttr <- <(<(('a' / 'A') ('d' / 'D') ('d' / 'D') ('r' / 'R') ':')> ws <Addr> Action14)> */
		nil,
		/* 18 TagAttr <- <(<(('t' / 'T') ('a' / 'A') ('g' / 'G'))> ws ':' ws <((('s' / 'S') ('t' / 'T') ('r' / 'R') ('u' / 'U') ('c' / 'C') ('t' / 'T')) / (('u' / 'U') ('n' / 'N') ('i' / 'I') ('o' / 'O') ('n' / 'N')))> Action15)> */
		nil,
		/* 19 BodyAttr <- <(<(('b' / 'B') ('o' / 'O') ('d' / 'D') ('y' / 'Y'))> ws ':' ws <((('u' / 'U') ('n' / 'N') ('d' / 'D') ('e' / 'E') ('f' / 'F') ('i' / 'I') ('n' / 'N') ('e' / 'E') ('d' / 'D')) / Node)> Action16)> */
		nil,
		/* 20 LinkAttr <- <(<(('l' / 'L') ('i' / 'I') ('n' / 'N') ('k' / 'K'))> ws ':' ws <((('e' / 'E') ('x' / 'X') ('t' / 'T') ('e' / 'E') ('r' / 'R') ('n' / 'N')) / (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('i' / 'I') ('c' / 'C')))> Action17)> */
		nil,
		/* 21 NoteAttr <- <(<(('n' / 'N') ('o' / 'O') ('t' / 'T') ('e' / 'E'))> ws ':' ws <((('p' / 'P') ('t' / 'T') ('r' / 'R') ('m' / 'M') ('e' / 'E') ('m' / 'M')) / (('c' / 'C') ('o' / 'O') ('n' / 'N') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('u' / 'U') ('c' / 'C') ('t' / 'T') ('o' / 'O') ('r' / 'R')) / ((&('D' | 'd') (('d' / 'D') ('e' / 'E') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('u' / 'U') ('c' / 'C') ('t' / 'T') ('o' / 'O') ('r' / 'R'))) | (&('P' | 'p') (('p' / 'P') ('s' / 'S') ('e' / 'E') ('u' / 'U') ('d' / 'D') ('o' / 'O') ' ' ('t' / 'T') ('m' / 'M') ('p' / 'P') ('l' / 'L'))) | (&('C' | 'c') (('c' / 'C') ('o' / 'O') ('n' / 'N') ('v' / 'V') ('e' / 'E') ('r' / 'R') ('s' / 'S') ('i' / 'I') ('o' / 'O') ('n' / 'N'))) | (&('A' | 'a') (('a' / 'A') ('r' / 'R') ('t' / 'T') ('i' / 'I') ('f' / 'F') ('i' / 'I') ('c' / 'C') ('i' / 'I') ('a' / 'A') ('l' / 'L')))))> ws Action18)> */
		nil,
		/* 22 AccsAttr <- <(<(('a' / 'A') ('c' / 'C') ('c' / 'C') ('s' / 'S'))> ws ':' ws <((('p' / 'P') ('r' / 'R') ('i' / 'I') ('v' / 'V')) / (('p' / 'P') ('u' / 'U') ('b' / 'B')) / (('p' / 'P') ('r' / 'R') ('o' / 'O') ('t' / 'T')))> Action19)> */
		nil,
		/* 23 QualAttr <- <(<(('q' / 'Q') ('u' / 'U') ('a' / 'A') ('l' / 'L'))> ws ':' ws <((&('R' | 'r') ('r' / 'R')) | (&('C') 'C') | (&('c') 'c') | (&('V' | 'v') ('v' / 'V')))>+ ws Action20)> */
		nil,
		/* 24 SignAttr <- <(<(('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N'))> ws ':' ws <((('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ('e' / 'E') ('d' / 'D')) / (('u' / 'U') ('n' / 'N') ('s' / 'S') ('i' / 'I') ('g' / 'G') ('n' / 'N') ('e' / 'E') ('d' / 'D')))> Action21)> */
		nil,
		/* 25 OtherField <- <(<[a-z]+> Action22)> */
		nil,
		/* 26 NodeFieldName <- <(<((('r' / 'R') ('e' / 'E') ('t' / 'T') ('n' / 'N')) / (('p' / 'P') ('r' / 'R') ('m' / 'M') ('s' / 'S')) / (('a' / 'A') ('r' / 'R') ('g' / 'G') ('s' / 'S')) / (('b' / 'B') ('a' / 'A') ('s' / 'S') ('e' / 'E')) / (('o' / 'O') ('r' / 'R') ('i' / 'I') ('g' / 'G')) / (('d' / 'D') ('e' / 'E') ('c' / 'C') ('l' / 'L')) / (('d' / 'D') ('c' / 'C') ('l' / 'L') ('s' / 'S')) / (('e' / 'E') ('x' / 'X') ('p' / 'P') ('r' / 'R')) / (('f' / 'F') ('n' / 'N') ('c' / 'C') ('s' / 'S')) / (('f' / 'F') ('n' / 'N')) / (('c' / 'C') ('n' / 'N') ('s' / 'S') ('t' / 'T')) / (('c' / 'C') ('o' / 'O') ('n' / 'N') ('d' / 'D')) / (('v' / 'V') ('a' / 'A') ('r' / 'R') ('s' / 'S')) / (('v' / 'V') ('f' / 'F') ('l' / 'L') ('d' / 'D')) / (('v' / 'V') ('a' / 'A') ('l' / 'L') ('u' / 'U')) / (('c' / 'C') ('h' / 'H') ('a' / 'A') ('n' / 'N')) / (('p' / 'P') ('u' / 'U') ('r' / 'R') ('p' / 'P')) / (('r' / 'R') ('e' / 'E') ('f' / 'F') ('d' / 'D')) / (('c' / 'C') ('l' / 'L') ('a' / 'A') ('s' / 'S')) / (('c' / 'C') ('l' / 'L') ('s' / 'S')) / (('c' / 'C') ('s' / 'S') ('t' / 'T') ('s' / 'S')) / (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) / (('m' / 'M') ('i' / 'I') ('n' / 'N')) / (('m' / 'M') ('a' / 'A') ('x' / 'X')) / (('b' / 'B') ('i' / 'I') ('n' / 'N') ('f' / 'F')) / (('s' / 'S') ('c' / 'C') ('p' / 'P') ('e' / 'E')) / (('i' / 'I') ('n' / 'N') ('i' / 'I') ('t' / 'T')) / ((&('V' | 'v') (('v' / 'V') ('a' / 'A') ('l' / 'L'))) | (&('I' | 'i') (('i' / 'I') ('d' / 'D') ('x' / 'X'))) | (&('S' | 's') (('s' / 'S') ('i' / 'I') ('z' / 'Z') ('e' / 'E'))) | (&('B' | 'b') (('b' / 'B') ('p' / 'P') ('o' / 'O') ('s' / 'S'))) | (&('C' | 'c') (('c' / 'C') ('h' / 'H') ('a' / 'A') ('i' / 'I') ('n' / 'N'))) | (&('R' | 'r') (('r' / 'R') ('s' / 'S') ('l' / 'L') ('t' / 'T'))) | (&('E' | 'e') (('e' / 'E') ('l' / 'L') ('t' / 'T') ('s' / 'S'))) | (&('D' | 'd') (('d' / 'D') ('o' / 'O') ('m' / 'M') ('n' / 'N'))) | (&('T' | 't') (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E'))) | (&('M' | 'm') (('m' / 'M') ('n' / 'N') ('g' / 'G') ('l' / 'L'))) | (&('N' | 'n') (('n' / 'N') ('a' / 'A') ('m' / 'M') ('e' / 'E'))) | (&('U' | 'u') (('u' / 'U') ('n' / 'N') ('q' / 'Q') ('l' / 'L'))) | (&('L' | 'l') (('l' / 'L') ('a' / 'A') ('b' / 'B') ('l' / 'L'))) | (&('P' | 'p') (('p' / 'P') ('t' / 'T') ('d' / 'D'))) | (&('F' | 'f') (('f' / 'F') ('l' / 'L') ('d' / 'D') ('s' / 'S'))) | (&('A' | 'a') (('a' / 'A') ('r' / 'R') ('g' / 'G') ('t' / 'T'))) | (&('O' | 'o') (('o' / 'O') ('p' / 'P') '0')) | (&('0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9') Integer)))> Action23)> */
		nil,
		/* 27 NodeAttr <- <(ws (NodeFieldName / OpAttr / OpAttr2 / OtherField) ws ':' ws <Node> ws Action24)> */
		nil,
		/* 28 SpecValue <- <(ws (('s' / 'S') ('p' / 'P') ('e' / 'E') ('c' / 'C') ':')? ws (<((&('R' | 'r') (('r' / 'R') ('e' / 'E') ('g' / 'G') ('i' / 'I') ('s' / 'S') ('t' / 'T') ('e' / 'E') ('r' / 'R'))) | (&('V' | 'v') (('v' / 'V') ('i' / 'I') ('r' / 'R') ('t' / 'T'))) | (&('P' | 'p') (('p' / 'P') ('u' / 'U') ('r' / 'R') ('e' / 'E'))) | (&('B' | 'b') (('b' / 'B') ('i' / 'I') ('t' / 'T') ('f' / 'F') ('i' / 'I') ('e' / 'E') ('l' / 'L') ('d' / 'D'))) | (&('M' | 'm') (('m' / 'M') ('u' / 'U') ('t' / 'T') ('a' / 'A') ('b' / 'B') ('l' / 'L') ('e' / 'E'))))> ws)+ Action25)> */
		nil,
		/* 29 LngtAttr <- <(('l' / 'L') ('n' / 'N') ('g' / 'G') ('t' / 'T') ws ':' ws <Integer> ws)> */
		nil,
		/* 30 RandomSpec <- <(ws (('s' / 'S') ('p' / 'P') ('e' / 'E') ('c' / 'C')) ws Action26)> */
		nil,
		/* 31 OneAttr <- <((StringAttr / SpecValue / NodeAttr / SourceAttr / IntAttr / SignAttr / IntAttrLow / AccsAttr / LangAttr / LinkAttr / IntAttr2 / ((&('A' | 'a') AddrAttr) | (&('L' | 'l') LngtAttr) | (&('I' | 'i') SignedIntAttr) | (&('Q' | 'q') QualAttr) | (&('N' | 'n') NoteAttr) | (&('B' | 'b') BodyAttr) | (&('T' | 't') TagAttr) | (&('H' | 'h') IntAttr3) | (&('\t' | '\n' | '\r' | ' ' | 'S' | 's') RandomSpec))) Action27)> */
		func() bool {
			position89, tokenIndex89 := position, tokenIndex
			{
				position90 := position
				{
					position91, tokenIndex91 := position, tokenIndex
					{
						position93 := position
						{
							position94, tokenIndex94 := position, tokenIndex
							if buffer[position] != rune('s') {
								goto l95
							}
							position++
							goto l94
						l95:
							position, tokenIndex = position94, tokenIndex94
							if buffer[position] != rune('S') {
								goto l92
							}
							position++
						}
					l94:
						{
							position96, tokenIndex96 := position, tokenIndex
							if buffer[position] != rune('t') {
								goto l97
							}
							position++
							goto l96
						l97:
							position, tokenIndex = position96, tokenIndex96
							if buffer[position] != rune('T') {
								goto l92
							}
							position++
						}
					l96:
						{
							position98, tokenIndex98 := position, tokenIndex
							if buffer[position] != rune('r') {
								goto l99
							}
							position++
							goto l98
						l99:
							position, tokenIndex = position98, tokenIndex98
							if buffer[position] != rune('R') {
								goto l92
							}
							position++
						}
					l98:
						{
							position100, tokenIndex100 := position, tokenIndex
							if buffer[position] != rune('g') {
								goto l101
							}
							position++
							goto l100
						l101:
							position, tokenIndex = position100, tokenIndex100
							if buffer[position] != rune('G') {
								goto l92
							}
							position++
						}
					l100:
						if buffer[position] != rune(':') {
							goto l92
						}
						position++
						if buffer[position] != rune(' ') {
							goto l92
						}
						position++
						{
							position102 := position
							{
								position105, tokenIndex105 := position, tokenIndex
								{
									position107, tokenIndex107 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l107
									}
									position++
									goto l106
								l107:
									position, tokenIndex = position107, tokenIndex107
								}
								if !matchDot() {
									goto l106
								}
								goto l105
							l106:
								position, tokenIndex = position105, tokenIndex105
								{
									position109, tokenIndex109 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l110
									}
									position++
									goto l109
								l110:
									position, tokenIndex = position109, tokenIndex109
									if buffer[position] != rune('L') {
										goto l108
									}
									position++
								}
							l109:
								{
									position111, tokenIndex111 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l111
									}
									position++
									goto l108
								l111:
									position, tokenIndex = position111, tokenIndex111
								}
								if !matchDot() {
									goto l108
								}
								goto l105
							l108:
								position, tokenIndex = position105, tokenIndex105
								{
									position113, tokenIndex113 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l114
									}
									position++
									goto l113
								l114:
									position, tokenIndex = position113, tokenIndex113
									if buffer[position] != rune('L') {
										goto l112
									}
									position++
								}
							l113:
								{
									position115, tokenIndex115 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l116
									}
									position++
									goto l115
								l116:
									position, tokenIndex = position115, tokenIndex115
									if buffer[position] != rune('N') {
										goto l112
									}
									position++
								}
							l115:
								{
									position117, tokenIndex117 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l117
									}
									position++
									goto l112
								l117:
									position, tokenIndex = position117, tokenIndex117
								}
								if !matchDot() {
									goto l112
								}
								goto l105
							l112:
								position, tokenIndex = position105, tokenIndex105
								{
									position119, tokenIndex119 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l120
									}
									position++
									goto l119
								l120:
									position, tokenIndex = position119, tokenIndex119
									if buffer[position] != rune('L') {
										goto l118
									}
									position++
								}
							l119:
								{
									position121, tokenIndex121 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l122
									}
									position++
									goto l121
								l122:
									position, tokenIndex = position121, tokenIndex121
									if buffer[position] != rune('N') {
										goto l118
									}
									position++
								}
							l121:
								{
									position123, tokenIndex123 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l124
									}
									position++
									goto l123
								l124:
									position, tokenIndex = position123, tokenIndex123
									if buffer[position] != rune('G') {
										goto l118
									}
									position++
								}
							l123:
								{
									position125, tokenIndex125 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l125
									}
									position++
									goto l118
								l125:
									position, tokenIndex = position125, tokenIndex125
								}
								if !matchDot() {
									goto l118
								}
								goto l105
							l118:
								position, tokenIndex = position105, tokenIndex105
								{
									position126, tokenIndex126 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l127
									}
									position++
									goto l126
								l127:
									position, tokenIndex = position126, tokenIndex126
									if buffer[position] != rune('L') {
										goto l92
									}
									position++
								}
							l126:
								{
									position128, tokenIndex128 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l129
									}
									position++
									goto l128
								l129:
									position, tokenIndex = position128, tokenIndex128
									if buffer[position] != rune('N') {
										goto l92
									}
									position++
								}
							l128:
								{
									position130, tokenIndex130 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l131
									}
									position++
									goto l130
								l131:
									position, tokenIndex = position130, tokenIndex130
									if buffer[position] != rune('G') {
										goto l92
									}
									position++
								}
							l130:
								{
									position132, tokenIndex132 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l133
									}
									position++
									goto l132
								l133:
									position, tokenIndex = position132, tokenIndex132
									if buffer[position] != rune('T') {
										goto l92
									}
									position++
								}
							l132:
								{
									position134, tokenIndex134 := position, tokenIndex
									if buffer[position] != rune(':') {
										goto l134
									}
									position++
									goto l92
								l134:
									position, tokenIndex = position134, tokenIndex134
								}
								if !matchDot() {
									goto l92
								}
							}
						l105:
						l103:
							{
								position104, tokenIndex104 := position, tokenIndex
								{
									position135, tokenIndex135 := position, tokenIndex
									{
										position137, tokenIndex137 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l137
										}
										position++
										goto l136
									l137:
										position, tokenIndex = position137, tokenIndex137
									}
									if !matchDot() {
										goto l136
									}
									goto l135
								l136:
									position, tokenIndex = position135, tokenIndex135
									{
										position139, tokenIndex139 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l140
										}
										position++
										goto l139
									l140:
										position, tokenIndex = position139, tokenIndex139
										if buffer[position] != rune('L') {
											goto l138
										}
										position++
									}
								l139:
									{
										position141, tokenIndex141 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l141
										}
										position++
										goto l138
									l141:
										position, tokenIndex = position141, tokenIndex141
									}
									if !matchDot() {
										goto l138
									}
									goto l135
								l138:
									position, tokenIndex = position135, tokenIndex135
									{
										position143, tokenIndex143 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l144
										}
										position++
										goto l143
									l144:
										position, tokenIndex = position143, tokenIndex143
										if buffer[position] != rune('L') {
											goto l142
										}
										position++
									}
								l143:
									{
										position145, tokenIndex145 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l146
										}
										position++
										goto l145
									l146:
										position, tokenIndex = position145, tokenIndex145
										if buffer[position] != rune('N') {
											goto l142
										}
										position++
									}
								l145:
									{
										position147, tokenIndex147 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l147
										}
										position++
										goto l142
									l147:
										position, tokenIndex = position147, tokenIndex147
									}
									if !matchDot() {
										goto l142
									}
									goto l135
								l142:
									position, tokenIndex = position135, tokenIndex135
									{
										position149, tokenIndex149 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l150
										}
										position++
										goto l149
									l150:
										position, tokenIndex = position149, tokenIndex149
										if buffer[position] != rune('L') {
											goto l148
										}
										position++
									}
								l149:
									{
										position151, tokenIndex151 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l152
										}
										position++
										goto l151
									l152:
										position, tokenIndex = position151, tokenIndex151
										if buffer[position] != rune('N') {
											goto l148
										}
										position++
									}
								l151:
									{
										position153, tokenIndex153 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l154
										}
										position++
										goto l153
									l154:
										position, tokenIndex = position153, tokenIndex153
										if buffer[position] != rune('G') {
											goto l148
										}
										position++
									}
								l153:
									{
										position155, tokenIndex155 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l155
										}
										position++
										goto l148
									l155:
										position, tokenIndex = position155, tokenIndex155
									}
									if !matchDot() {
										goto l148
									}
									goto l135
								l148:
									position, tokenIndex = position135, tokenIndex135
									{
										position156, tokenIndex156 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l157
										}
										position++
										goto l156
									l157:
										position, tokenIndex = position156, tokenIndex156
										if buffer[position] != rune('L') {
											goto l104
										}
										position++
									}
								l156:
									{
										position158, tokenIndex158 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l159
										}
										position++
										goto l158
									l159:
										position, tokenIndex = position158, tokenIndex158
										if buffer[position] != rune('N') {
											goto l104
										}
										position++
									}
								l158:
									{
										position160, tokenIndex160 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l161
										}
										position++
										goto l160
									l161:
										position, tokenIndex = position160, tokenIndex160
										if buffer[position] != rune('G') {
											goto l104
										}
										position++
									}
								l160:
									{
										position162, tokenIndex162 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l163
										}
										position++
										goto l162
									l163:
										position, tokenIndex = position162, tokenIndex162
										if buffer[position] != rune('T') {
											goto l104
										}
										position++
									}
								l162:
									{
										position164, tokenIndex164 := position, tokenIndex
										if buffer[position] != rune(':') {
											goto l164
										}
										position++
										goto l104
									l164:
										position, tokenIndex = position164, tokenIndex164
									}
									if !matchDot() {
										goto l104
									}
								}
							l135:
								goto l103
							l104:
								position, tokenIndex = position104, tokenIndex104
							}
							add(rulePegText, position102)
						}
						{
							position165, tokenIndex165 := position, tokenIndex
							{
								position166, tokenIndex166 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l167
								}
								position++
								goto l166
							l167:
								position, tokenIndex = position166, tokenIndex166
								if buffer[position] != rune('L') {
									goto l92
								}
								position++
							}
						l166:
							{
								position168, tokenIndex168 := position, tokenIndex
								if buffer[position] != rune('n') {
									goto l169
								}
								position++
								goto l168
							l169:
								position, tokenIndex = position168, tokenIndex168
								if buffer[position] != rune('N') {
									goto l92
								}
								position++
							}
						l168:
							{
								position170, tokenIndex170 := position, tokenIndex
								if buffer[position] != rune('g') {
									goto l171
								}
								position++
								goto l170
							l171:
								position, tokenIndex = position170, tokenIndex170
								if buffer[position] != rune('G') {
									goto l92
								}
								position++
							}
						l170:
							{
								position172, tokenIndex172 := position, tokenIndex
								if buffer[position] != rune('t') {
									goto l173
								}
								position++
								goto l172
							l173:
								position, tokenIndex = position172, tokenIndex172
								if buffer[position] != rune('T') {
									goto l92
								}
								position++
							}
						l172:
							if buffer[position] != rune(':') {
								goto l92
							}
							position++
							position, tokenIndex = position165, tokenIndex165
						}
						{
							add(ruleAction29, position)
						}
						add(ruleStringAttr, position93)
					}
					goto l91
				l92:
					position, tokenIndex = position91, tokenIndex91
					{
						position176 := position
						if !_rules[rulews]() {
							goto l175
						}
						{
							position177, tokenIndex177 := position, tokenIndex
							{
								position179, tokenIndex179 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l180
								}
								position++
								goto l179
							l180:
								position, tokenIndex = position179, tokenIndex179
								if buffer[position] != rune('S') {
									goto l177
								}
								position++
							}
						l179:
							{
								position181, tokenIndex181 := position, tokenIndex
								if buffer[position] != rune('p') {
									goto l182
								}
								position++
								goto l181
							l182:
								position, tokenIndex = position181, tokenIndex181
								if buffer[position] != rune('P') {
									goto l177
								}
								position++
							}
						l181:
							{
								position183, tokenIndex183 := position, tokenIndex
								if buffer[position] != rune('e') {
									goto l184
								}
								position++
								goto l183
							l184:
								position, tokenIndex = position183, tokenIndex183
								if buffer[position] != rune('E') {
									goto l177
								}
								position++
							}
						l183:
							{
								position185, tokenIndex185 := position, tokenIndex
								if buffer[position] != rune('c') {
									goto l186
								}
								position++
								goto l185
							l186:
								position, tokenIndex = position185, tokenIndex185
								if buffer[position] != rune('C') {
									goto l177
								}
								position++
							}
						l185:
							if buffer[position] != rune(':') {
								goto l177
							}
							position++
							goto l178
						l177:
							position, tokenIndex = position177, tokenIndex177
						}
					l178:
						if !_rules[rulews]() {
							goto l175
						}
						{
							position189 := position
							{
								switch buffer[position] {
								case 'R', 'r':
									{
										position191, tokenIndex191 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l192
										}
										position++
										goto l191
									l192:
										position, tokenIndex = position191, tokenIndex191
										if buffer[position] != rune('R') {
											goto l175
										}
										position++
									}
								l191:
									{
										position193, tokenIndex193 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l194
										}
										position++
										goto l193
									l194:
										position, tokenIndex = position193, tokenIndex193
										if buffer[position] != rune('E') {
											goto l175
										}
										position++
									}
								l193:
									{
										position195, tokenIndex195 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l196
										}
										position++
										goto l195
									l196:
										position, tokenIndex = position195, tokenIndex195
										if buffer[position] != rune('G') {
											goto l175
										}
										position++
									}
								l195:
									{
										position197, tokenIndex197 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l198
										}
										position++
										goto l197
									l198:
										position, tokenIndex = position197, tokenIndex197
										if buffer[position] != rune('I') {
											goto l175
										}
										position++
									}
								l197:
									{
										position199, tokenIndex199 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l200
										}
										position++
										goto l199
									l200:
										position, tokenIndex = position199, tokenIndex199
										if buffer[position] != rune('S') {
											goto l175
										}
										position++
									}
								l199:
									{
										position201, tokenIndex201 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l202
										}
										position++
										goto l201
									l202:
										position, tokenIndex = position201, tokenIndex201
										if buffer[position] != rune('T') {
											goto l175
										}
										position++
									}
								l201:
									{
										position203, tokenIndex203 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l204
										}
										position++
										goto l203
									l204:
										position, tokenIndex = position203, tokenIndex203
										if buffer[position] != rune('E') {
											goto l175
										}
										position++
									}
								l203:
									{
										position205, tokenIndex205 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l206
										}
										position++
										goto l205
									l206:
										position, tokenIndex = position205, tokenIndex205
										if buffer[position] != rune('R') {
											goto l175
										}
										position++
									}
								l205:
									break
								case 'V', 'v':
									{
										position207, tokenIndex207 := position, tokenIndex
										if buffer[position] != rune('v') {
											goto l208
										}
										position++
										goto l207
									l208:
										position, tokenIndex = position207, tokenIndex207
										if buffer[position] != rune('V') {
											goto l175
										}
										position++
									}
								l207:
									{
										position209, tokenIndex209 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l210
										}
										position++
										goto l209
									l210:
										position, tokenIndex = position209, tokenIndex209
										if buffer[position] != rune('I') {
											goto l175
										}
										position++
									}
								l209:
									{
										position211, tokenIndex211 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l212
										}
										position++
										goto l211
									l212:
										position, tokenIndex = position211, tokenIndex211
										if buffer[position] != rune('R') {
											goto l175
										}
										position++
									}
								l211:
									{
										position213, tokenIndex213 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l214
										}
										position++
										goto l213
									l214:
										position, tokenIndex = position213, tokenIndex213
										if buffer[position] != rune('T') {
											goto l175
										}
										position++
									}
								l213:
									break
								case 'P', 'p':
									{
										position215, tokenIndex215 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l216
										}
										position++
										goto l215
									l216:
										position, tokenIndex = position215, tokenIndex215
										if buffer[position] != rune('P') {
											goto l175
										}
										position++
									}
								l215:
									{
										position217, tokenIndex217 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l218
										}
										position++
										goto l217
									l218:
										position, tokenIndex = position217, tokenIndex217
										if buffer[position] != rune('U') {
											goto l175
										}
										position++
									}
								l217:
									{
										position219, tokenIndex219 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l220
										}
										position++
										goto l219
									l220:
										position, tokenIndex = position219, tokenIndex219
										if buffer[position] != rune('R') {
											goto l175
										}
										position++
									}
								l219:
									{
										position221, tokenIndex221 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l222
										}
										position++
										goto l221
									l222:
										position, tokenIndex = position221, tokenIndex221
										if buffer[position] != rune('E') {
											goto l175
										}
										position++
									}
								l221:
									break
								case 'B', 'b':
									{
										position223, tokenIndex223 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l224
										}
										position++
										goto l223
									l224:
										position, tokenIndex = position223, tokenIndex223
										if buffer[position] != rune('B') {
											goto l175
										}
										position++
									}
								l223:
									{
										position225, tokenIndex225 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l226
										}
										position++
										goto l225
									l226:
										position, tokenIndex = position225, tokenIndex225
										if buffer[position] != rune('I') {
											goto l175
										}
										position++
									}
								l225:
									{
										position227, tokenIndex227 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l228
										}
										position++
										goto l227
									l228:
										position, tokenIndex = position227, tokenIndex227
										if buffer[position] != rune('T') {
											goto l175
										}
										position++
									}
								l227:
									{
										position229, tokenIndex229 := position, tokenIndex
										if buffer[position] != rune('f') {
											goto l230
										}
										position++
										goto l229
									l230:
										position, tokenIndex = position229, tokenIndex229
										if buffer[position] != rune('F') {
											goto l175
										}
										position++
									}
								l229:
									{
										position231, tokenIndex231 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l232
										}
										position++
										goto l231
									l232:
										position, tokenIndex = position231, tokenIndex231
										if buffer[position] != rune('I') {
											goto l175
										}
										position++
									}
								l231:
									{
										position233, tokenIndex233 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l234
										}
										position++
										goto l233
									l234:
										position, tokenIndex = position233, tokenIndex233
										if buffer[position] != rune('E') {
											goto l175
										}
										position++
									}
								l233:
									{
										position235, tokenIndex235 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l236
										}
										position++
										goto l235
									l236:
										position, tokenIndex = position235, tokenIndex235
										if buffer[position] != rune('L') {
											goto l175
										}
										position++
									}
								l235:
									{
										position237, tokenIndex237 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l238
										}
										position++
										goto l237
									l238:
										position, tokenIndex = position237, tokenIndex237
										if buffer[position] != rune('D') {
											goto l175
										}
										position++
									}
								l237:
									break
								default:
									{
										position239, tokenIndex239 := position, tokenIndex
										if buffer[position] != rune('m') {
											goto l240
										}
										position++
										goto l239
									l240:
										position, tokenIndex = position239, tokenIndex239
										if buffer[position] != rune('M') {
											goto l175
										}
										position++
									}
								l239:
									{
										position241, tokenIndex241 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l242
										}
										position++
										goto l241
									l242:
										position, tokenIndex = position241, tokenIndex241
										if buffer[position] != rune('U') {
											goto l175
										}
										position++
									}
								l241:
									{
										position243, tokenIndex243 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l244
										}
										position++
										goto l243
									l244:
										position, tokenIndex = position243, tokenIndex243
										if buffer[position] != rune('T') {
											goto l175
										}
										position++
									}
								l243:
									{
										position245, tokenIndex245 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l246
										}
										position++
										goto l245
									l246:
										position, tokenIndex = position245, tokenIndex245
										if buffer[position] != rune('A') {
											goto l175
										}
										position++
									}
								l245:
									{
										position247, tokenIndex247 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l248
										}
										position++
										goto l247
									l248:
										position, tokenIndex = position247, tokenIndex247
										if buffer[position] != rune('B') {
											goto l175
										}
										position++
									}
								l247:
									{
										position249, tokenIndex249 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l250
										}
										position++
										goto l249
									l250:
										position, tokenIndex = position249, tokenIndex249
										if buffer[position] != rune('L') {
											goto l175
										}
										position++
									}
								l249:
									{
										position251, tokenIndex251 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l252
										}
										position++
										goto l251
									l252:
										position, tokenIndex = position251, tokenIndex251
										if buffer[position] != rune('E') {
											goto l175
										}
										position++
									}
								l251:
									break
								}
							}

							add(rulePegText, position189)
						}
						if !_rules[rulews]() {
							goto l175
						}
					l187:
						{
							position188, tokenIndex188 := position, tokenIndex
							{
								position253 := position
								{
									switch buffer[position] {
									case 'R', 'r':
										{
											position255, tokenIndex255 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l256
											}
											position++
											goto l255
										l256:
											position, tokenIndex = position255, tokenIndex255
											if buffer[position] != rune('R') {
												goto l188
											}
											position++
										}
									l255:
										{
											position257, tokenIndex257 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l258
											}
											position++
											goto l257
										l258:
											position, tokenIndex = position257, tokenIndex257
											if buffer[position] != rune('E') {
												goto l188
											}
											position++
										}
									l257:
										{
											position259, tokenIndex259 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l260
											}
											position++
											goto l259
										l260:
											position, tokenIndex = position259, tokenIndex259
											if buffer[position] != rune('G') {
												goto l188
											}
											position++
										}
									l259:
										{
											position261, tokenIndex261 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l262
											}
											position++
											goto l261
										l262:
											position, tokenIndex = position261, tokenIndex261
											if buffer[position] != rune('I') {
												goto l188
											}
											position++
										}
									l261:
										{
											position263, tokenIndex263 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l264
											}
											position++
											goto l263
										l264:
											position, tokenIndex = position263, tokenIndex263
											if buffer[position] != rune('S') {
												goto l188
											}
											position++
										}
									l263:
										{
											position265, tokenIndex265 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l266
											}
											position++
											goto l265
										l266:
											position, tokenIndex = position265, tokenIndex265
											if buffer[position] != rune('T') {
												goto l188
											}
											position++
										}
									l265:
										{
											position267, tokenIndex267 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l268
											}
											position++
											goto l267
										l268:
											position, tokenIndex = position267, tokenIndex267
											if buffer[position] != rune('E') {
												goto l188
											}
											position++
										}
									l267:
										{
											position269, tokenIndex269 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l270
											}
											position++
											goto l269
										l270:
											position, tokenIndex = position269, tokenIndex269
											if buffer[position] != rune('R') {
												goto l188
											}
											position++
										}
									l269:
										break
									case 'V', 'v':
										{
											position271, tokenIndex271 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l272
											}
											position++
											goto l271
										l272:
											position, tokenIndex = position271, tokenIndex271
											if buffer[position] != rune('V') {
												goto l188
											}
											position++
										}
									l271:
										{
											position273, tokenIndex273 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l274
											}
											position++
											goto l273
										l274:
											position, tokenIndex = position273, tokenIndex273
											if buffer[position] != rune('I') {
												goto l188
											}
											position++
										}
									l273:
										{
											position275, tokenIndex275 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l276
											}
											position++
											goto l275
										l276:
											position, tokenIndex = position275, tokenIndex275
											if buffer[position] != rune('R') {
												goto l188
											}
											position++
										}
									l275:
										{
											position277, tokenIndex277 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l278
											}
											position++
											goto l277
										l278:
											position, tokenIndex = position277, tokenIndex277
											if buffer[position] != rune('T') {
												goto l188
											}
											position++
										}
									l277:
										break
									case 'P', 'p':
										{
											position279, tokenIndex279 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l280
											}
											position++
											goto l279
										l280:
											position, tokenIndex = position279, tokenIndex279
											if buffer[position] != rune('P') {
												goto l188
											}
											position++
										}
									l279:
										{
											position281, tokenIndex281 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l282
											}
											position++
											goto l281
										l282:
											position, tokenIndex = position281, tokenIndex281
											if buffer[position] != rune('U') {
												goto l188
											}
											position++
										}
									l281:
										{
											position283, tokenIndex283 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l284
											}
											position++
											goto l283
										l284:
											position, tokenIndex = position283, tokenIndex283
											if buffer[position] != rune('R') {
												goto l188
											}
											position++
										}
									l283:
										{
											position285, tokenIndex285 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l286
											}
											position++
											goto l285
										l286:
											position, tokenIndex = position285, tokenIndex285
											if buffer[position] != rune('E') {
												goto l188
											}
											position++
										}
									l285:
										break
									case 'B', 'b':
										{
											position287, tokenIndex287 := position, tokenIndex
											if buffer[position] != rune('b') {
												goto l288
											}
											position++
											goto l287
										l288:
											position, tokenIndex = position287, tokenIndex287
											if buffer[position] != rune('B') {
												goto l188
											}
											position++
										}
									l287:
										{
											position289, tokenIndex289 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l290
											}
											position++
											goto l289
										l290:
											position, tokenIndex = position289, tokenIndex289
											if buffer[position] != rune('I') {
												goto l188
											}
											position++
										}
									l289:
										{
											position291, tokenIndex291 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l292
											}
											position++
											goto l291
										l292:
											position, tokenIndex = position291, tokenIndex291
											if buffer[position] != rune('T') {
												goto l188
											}
											position++
										}
									l291:
										{
											position293, tokenIndex293 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l294
											}
											position++
											goto l293
										l294:
											position, tokenIndex = position293, tokenIndex293
											if buffer[position] != rune('F') {
												goto l188
											}
											position++
										}
									l293:
										{
											position295, tokenIndex295 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l296
											}
											position++
											goto l295
										l296:
											position, tokenIndex = position295, tokenIndex295
											if buffer[position] != rune('I') {
												goto l188
											}
											position++
										}
									l295:
										{
											position297, tokenIndex297 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l298
											}
											position++
											goto l297
										l298:
											position, tokenIndex = position297, tokenIndex297
											if buffer[position] != rune('E') {
												goto l188
											}
											position++
										}
									l297:
										{
											position299, tokenIndex299 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l300
											}
											position++
											goto l299
										l300:
											position, tokenIndex = position299, tokenIndex299
											if buffer[position] != rune('L') {
												goto l188
											}
											position++
										}
									l299:
										{
											position301, tokenIndex301 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l302
											}
											position++
											goto l301
										l302:
											position, tokenIndex = position301, tokenIndex301
											if buffer[position] != rune('D') {
												goto l188
											}
											position++
										}
									l301:
										break
									default:
										{
											position303, tokenIndex303 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l304
											}
											position++
											goto l303
										l304:
											position, tokenIndex = position303, tokenIndex303
											if buffer[position] != rune('M') {
												goto l188
											}
											position++
										}
									l303:
										{
											position305, tokenIndex305 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l306
											}
											position++
											goto l305
										l306:
											position, tokenIndex = position305, tokenIndex305
											if buffer[position] != rune('U') {
												goto l188
											}
											position++
										}
									l305:
										{
											position307, tokenIndex307 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l308
											}
											position++
											goto l307
										l308:
											position, tokenIndex = position307, tokenIndex307
											if buffer[position] != rune('T') {
												goto l188
											}
											position++
										}
									l307:
										{
											position309, tokenIndex309 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l310
											}
											position++
											goto l309
										l310:
											position, tokenIndex = position309, tokenIndex309
											if buffer[position] != rune('A') {
												goto l188
											}
											position++
										}
									l309:
										{
											position311, tokenIndex311 := position, tokenIndex
											if buffer[position] != rune('b') {
												goto l312
											}
											position++
											goto l311
										l312:
											position, tokenIndex = position311, tokenIndex311
											if buffer[position] != rune('B') {
												goto l188
											}
											position++
										}
									l311:
										{
											position313, tokenIndex313 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l314
											}
											position++
											goto l313
										l314:
											position, tokenIndex = position313, tokenIndex313
											if buffer[position] != rune('L') {
												goto l188
											}
											position++
										}
									l313:
										{
											position315, tokenIndex315 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l316
											}
											position++
											goto l315
										l316:
											position, tokenIndex = position315, tokenIndex315
											if buffer[position] != rune('E') {
												goto l188
											}
											position++
										}
									l315:
										break
									}
								}

								add(rulePegText, position253)
							}
							if !_rules[rulews]() {
								goto l188
							}
							goto l187
						l188:
							position, tokenIndex = position188, tokenIndex188
						}
						{
							add(ruleAction25, position)
						}
						add(ruleSpecValue, position176)
					}
					goto l91
				l175:
					position, tokenIndex = position91, tokenIndex91
					{
						position319 := position
						if !_rules[rulews]() {
							goto l318
						}
						{
							position320, tokenIndex320 := position, tokenIndex
							{
								position322 := position
								{
									position323 := position
									{
										position324, tokenIndex324 := position, tokenIndex
										{
											position326, tokenIndex326 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l327
											}
											position++
											goto l326
										l327:
											position, tokenIndex = position326, tokenIndex326
											if buffer[position] != rune('R') {
												goto l325
											}
											position++
										}
									l326:
										{
											position328, tokenIndex328 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l329
											}
											position++
											goto l328
										l329:
											position, tokenIndex = position328, tokenIndex328
											if buffer[position] != rune('E') {
												goto l325
											}
											position++
										}
									l328:
										{
											position330, tokenIndex330 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l331
											}
											position++
											goto l330
										l331:
											position, tokenIndex = position330, tokenIndex330
											if buffer[position] != rune('T') {
												goto l325
											}
											position++
										}
									l330:
										{
											position332, tokenIndex332 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l333
											}
											position++
											goto l332
										l333:
											position, tokenIndex = position332, tokenIndex332
											if buffer[position] != rune('N') {
												goto l325
											}
											position++
										}
									l332:
										goto l324
									l325:
										position, tokenIndex = position324, tokenIndex324
										{
											position335, tokenIndex335 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l336
											}
											position++
											goto l335
										l336:
											position, tokenIndex = position335, tokenIndex335
											if buffer[position] != rune('P') {
												goto l334
											}
											position++
										}
									l335:
										{
											position337, tokenIndex337 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l338
											}
											position++
											goto l337
										l338:
											position, tokenIndex = position337, tokenIndex337
											if buffer[position] != rune('R') {
												goto l334
											}
											position++
										}
									l337:
										{
											position339, tokenIndex339 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l340
											}
											position++
											goto l339
										l340:
											position, tokenIndex = position339, tokenIndex339
											if buffer[position] != rune('M') {
												goto l334
											}
											position++
										}
									l339:
										{
											position341, tokenIndex341 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l342
											}
											position++
											goto l341
										l342:
											position, tokenIndex = position341, tokenIndex341
											if buffer[position] != rune('S') {
												goto l334
											}
											position++
										}
									l341:
										goto l324
									l334:
										position, tokenIndex = position324, tokenIndex324
										{
											position344, tokenIndex344 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l345
											}
											position++
											goto l344
										l345:
											position, tokenIndex = position344, tokenIndex344
											if buffer[position] != rune('A') {
												goto l343
											}
											position++
										}
									l344:
										{
											position346, tokenIndex346 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l347
											}
											position++
											goto l346
										l347:
											position, tokenIndex = position346, tokenIndex346
											if buffer[position] != rune('R') {
												goto l343
											}
											position++
										}
									l346:
										{
											position348, tokenIndex348 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l349
											}
											position++
											goto l348
										l349:
											position, tokenIndex = position348, tokenIndex348
											if buffer[position] != rune('G') {
												goto l343
											}
											position++
										}
									l348:
										{
											position350, tokenIndex350 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l351
											}
											position++
											goto l350
										l351:
											position, tokenIndex = position350, tokenIndex350
											if buffer[position] != rune('S') {
												goto l343
											}
											position++
										}
									l350:
										goto l324
									l343:
										position, tokenIndex = position324, tokenIndex324
										{
											position353, tokenIndex353 := position, tokenIndex
											if buffer[position] != rune('b') {
												goto l354
											}
											position++
											goto l353
										l354:
											position, tokenIndex = position353, tokenIndex353
											if buffer[position] != rune('B') {
												goto l352
											}
											position++
										}
									l353:
										{
											position355, tokenIndex355 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l356
											}
											position++
											goto l355
										l356:
											position, tokenIndex = position355, tokenIndex355
											if buffer[position] != rune('A') {
												goto l352
											}
											position++
										}
									l355:
										{
											position357, tokenIndex357 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l358
											}
											position++
											goto l357
										l358:
											position, tokenIndex = position357, tokenIndex357
											if buffer[position] != rune('S') {
												goto l352
											}
											position++
										}
									l357:
										{
											position359, tokenIndex359 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l360
											}
											position++
											goto l359
										l360:
											position, tokenIndex = position359, tokenIndex359
											if buffer[position] != rune('E') {
												goto l352
											}
											position++
										}
									l359:
										goto l324
									l352:
										position, tokenIndex = position324, tokenIndex324
										{
											position362, tokenIndex362 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l363
											}
											position++
											goto l362
										l363:
											position, tokenIndex = position362, tokenIndex362
											if buffer[position] != rune('O') {
												goto l361
											}
											position++
										}
									l362:
										{
											position364, tokenIndex364 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l365
											}
											position++
											goto l364
										l365:
											position, tokenIndex = position364, tokenIndex364
											if buffer[position] != rune('R') {
												goto l361
											}
											position++
										}
									l364:
										{
											position366, tokenIndex366 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l367
											}
											position++
											goto l366
										l367:
											position, tokenIndex = position366, tokenIndex366
											if buffer[position] != rune('I') {
												goto l361
											}
											position++
										}
									l366:
										{
											position368, tokenIndex368 := position, tokenIndex
											if buffer[position] != rune('g') {
												goto l369
											}
											position++
											goto l368
										l369:
											position, tokenIndex = position368, tokenIndex368
											if buffer[position] != rune('G') {
												goto l361
											}
											position++
										}
									l368:
										goto l324
									l361:
										position, tokenIndex = position324, tokenIndex324
										{
											position371, tokenIndex371 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l372
											}
											position++
											goto l371
										l372:
											position, tokenIndex = position371, tokenIndex371
											if buffer[position] != rune('D') {
												goto l370
											}
											position++
										}
									l371:
										{
											position373, tokenIndex373 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l374
											}
											position++
											goto l373
										l374:
											position, tokenIndex = position373, tokenIndex373
											if buffer[position] != rune('E') {
												goto l370
											}
											position++
										}
									l373:
										{
											position375, tokenIndex375 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l376
											}
											position++
											goto l375
										l376:
											position, tokenIndex = position375, tokenIndex375
											if buffer[position] != rune('C') {
												goto l370
											}
											position++
										}
									l375:
										{
											position377, tokenIndex377 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l378
											}
											position++
											goto l377
										l378:
											position, tokenIndex = position377, tokenIndex377
											if buffer[position] != rune('L') {
												goto l370
											}
											position++
										}
									l377:
										goto l324
									l370:
										position, tokenIndex = position324, tokenIndex324
										{
											position380, tokenIndex380 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l381
											}
											position++
											goto l380
										l381:
											position, tokenIndex = position380, tokenIndex380
											if buffer[position] != rune('D') {
												goto l379
											}
											position++
										}
									l380:
										{
											position382, tokenIndex382 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l383
											}
											position++
											goto l382
										l383:
											position, tokenIndex = position382, tokenIndex382
											if buffer[position] != rune('C') {
												goto l379
											}
											position++
										}
									l382:
										{
											position384, tokenIndex384 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l385
											}
											position++
											goto l384
										l385:
											position, tokenIndex = position384, tokenIndex384
											if buffer[position] != rune('L') {
												goto l379
											}
											position++
										}
									l384:
										{
											position386, tokenIndex386 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l387
											}
											position++
											goto l386
										l387:
											position, tokenIndex = position386, tokenIndex386
											if buffer[position] != rune('S') {
												goto l379
											}
											position++
										}
									l386:
										goto l324
									l379:
										position, tokenIndex = position324, tokenIndex324
										{
											position389, tokenIndex389 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l390
											}
											position++
											goto l389
										l390:
											position, tokenIndex = position389, tokenIndex389
											if buffer[position] != rune('E') {
												goto l388
											}
											position++
										}
									l389:
										{
											position391, tokenIndex391 := position, tokenIndex
											if buffer[position] != rune('x') {
												goto l392
											}
											position++
											goto l391
										l392:
											position, tokenIndex = position391, tokenIndex391
											if buffer[position] != rune('X') {
												goto l388
											}
											position++
										}
									l391:
										{
											position393, tokenIndex393 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l394
											}
											position++
											goto l393
										l394:
											position, tokenIndex = position393, tokenIndex393
											if buffer[position] != rune('P') {
												goto l388
											}
											position++
										}
									l393:
										{
											position395, tokenIndex395 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l396
											}
											position++
											goto l395
										l396:
											position, tokenIndex = position395, tokenIndex395
											if buffer[position] != rune('R') {
												goto l388
											}
											position++
										}
									l395:
										goto l324
									l388:
										position, tokenIndex = position324, tokenIndex324
										{
											position398, tokenIndex398 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l399
											}
											position++
											goto l398
										l399:
											position, tokenIndex = position398, tokenIndex398
											if buffer[position] != rune('F') {
												goto l397
											}
											position++
										}
									l398:
										{
											position400, tokenIndex400 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l401
											}
											position++
											goto l400
										l401:
											position, tokenIndex = position400, tokenIndex400
											if buffer[position] != rune('N') {
												goto l397
											}
											position++
										}
									l400:
										{
											position402, tokenIndex402 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l403
											}
											position++
											goto l402
										l403:
											position, tokenIndex = position402, tokenIndex402
											if buffer[position] != rune('C') {
												goto l397
											}
											position++
										}
									l402:
										{
											position404, tokenIndex404 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l405
											}
											position++
											goto l404
										l405:
											position, tokenIndex = position404, tokenIndex404
											if buffer[position] != rune('S') {
												goto l397
											}
											position++
										}
									l404:
										goto l324
									l397:
										position, tokenIndex = position324, tokenIndex324
										{
											position407, tokenIndex407 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l408
											}
											position++
											goto l407
										l408:
											position, tokenIndex = position407, tokenIndex407
											if buffer[position] != rune('F') {
												goto l406
											}
											position++
										}
									l407:
										{
											position409, tokenIndex409 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l410
											}
											position++
											goto l409
										l410:
											position, tokenIndex = position409, tokenIndex409
											if buffer[position] != rune('N') {
												goto l406
											}
											position++
										}
									l409:
										goto l324
									l406:
										position, tokenIndex = position324, tokenIndex324
										{
											position412, tokenIndex412 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l413
											}
											position++
											goto l412
										l413:
											position, tokenIndex = position412, tokenIndex412
											if buffer[position] != rune('C') {
												goto l411
											}
											position++
										}
									l412:
										{
											position414, tokenIndex414 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l415
											}
											position++
											goto l414
										l415:
											position, tokenIndex = position414, tokenIndex414
											if buffer[position] != rune('N') {
												goto l411
											}
											position++
										}
									l414:
										{
											position416, tokenIndex416 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l417
											}
											position++
											goto l416
										l417:
											position, tokenIndex = position416, tokenIndex416
											if buffer[position] != rune('S') {
												goto l411
											}
											position++
										}
									l416:
										{
											position418, tokenIndex418 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l419
											}
											position++
											goto l418
										l419:
											position, tokenIndex = position418, tokenIndex418
											if buffer[position] != rune('T') {
												goto l411
											}
											position++
										}
									l418:
										goto l324
									l411:
										position, tokenIndex = position324, tokenIndex324
										{
											position421, tokenIndex421 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l422
											}
											position++
											goto l421
										l422:
											position, tokenIndex = position421, tokenIndex421
											if buffer[position] != rune('C') {
												goto l420
											}
											position++
										}
									l421:
										{
											position423, tokenIndex423 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l424
											}
											position++
											goto l423
										l424:
											position, tokenIndex = position423, tokenIndex423
											if buffer[position] != rune('O') {
												goto l420
											}
											position++
										}
									l423:
										{
											position425, tokenIndex425 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l426
											}
											position++
											goto l425
										l426:
											position, tokenIndex = position425, tokenIndex425
											if buffer[position] != rune('N') {
												goto l420
											}
											position++
										}
									l425:
										{
											position427, tokenIndex427 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l428
											}
											position++
											goto l427
										l428:
											position, tokenIndex = position427, tokenIndex427
											if buffer[position] != rune('D') {
												goto l420
											}
											position++
										}
									l427:
										goto l324
									l420:
										position, tokenIndex = position324, tokenIndex324
										{
											position430, tokenIndex430 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l431
											}
											position++
											goto l430
										l431:
											position, tokenIndex = position430, tokenIndex430
											if buffer[position] != rune('V') {
												goto l429
											}
											position++
										}
									l430:
										{
											position432, tokenIndex432 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l433
											}
											position++
											goto l432
										l433:
											position, tokenIndex = position432, tokenIndex432
											if buffer[position] != rune('A') {
												goto l429
											}
											position++
										}
									l432:
										{
											position434, tokenIndex434 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l435
											}
											position++
											goto l434
										l435:
											position, tokenIndex = position434, tokenIndex434
											if buffer[position] != rune('R') {
												goto l429
											}
											position++
										}
									l434:
										{
											position436, tokenIndex436 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l437
											}
											position++
											goto l436
										l437:
											position, tokenIndex = position436, tokenIndex436
											if buffer[position] != rune('S') {
												goto l429
											}
											position++
										}
									l436:
										goto l324
									l429:
										position, tokenIndex = position324, tokenIndex324
										{
											position439, tokenIndex439 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l440
											}
											position++
											goto l439
										l440:
											position, tokenIndex = position439, tokenIndex439
											if buffer[position] != rune('V') {
												goto l438
											}
											position++
										}
									l439:
										{
											position441, tokenIndex441 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l442
											}
											position++
											goto l441
										l442:
											position, tokenIndex = position441, tokenIndex441
											if buffer[position] != rune('F') {
												goto l438
											}
											position++
										}
									l441:
										{
											position443, tokenIndex443 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l444
											}
											position++
											goto l443
										l444:
											position, tokenIndex = position443, tokenIndex443
											if buffer[position] != rune('L') {
												goto l438
											}
											position++
										}
									l443:
										{
											position445, tokenIndex445 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l446
											}
											position++
											goto l445
										l446:
											position, tokenIndex = position445, tokenIndex445
											if buffer[position] != rune('D') {
												goto l438
											}
											position++
										}
									l445:
										goto l324
									l438:
										position, tokenIndex = position324, tokenIndex324
										{
											position448, tokenIndex448 := position, tokenIndex
											if buffer[position] != rune('v') {
												goto l449
											}
											position++
											goto l448
										l449:
											position, tokenIndex = position448, tokenIndex448
											if buffer[position] != rune('V') {
												goto l447
											}
											position++
										}
									l448:
										{
											position450, tokenIndex450 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l451
											}
											position++
											goto l450
										l451:
											position, tokenIndex = position450, tokenIndex450
											if buffer[position] != rune('A') {
												goto l447
											}
											position++
										}
									l450:
										{
											position452, tokenIndex452 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l453
											}
											position++
											goto l452
										l453:
											position, tokenIndex = position452, tokenIndex452
											if buffer[position] != rune('L') {
												goto l447
											}
											position++
										}
									l452:
										{
											position454, tokenIndex454 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l455
											}
											position++
											goto l454
										l455:
											position, tokenIndex = position454, tokenIndex454
											if buffer[position] != rune('U') {
												goto l447
											}
											position++
										}
									l454:
										goto l324
									l447:
										position, tokenIndex = position324, tokenIndex324
										{
											position457, tokenIndex457 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l458
											}
											position++
											goto l457
										l458:
											position, tokenIndex = position457, tokenIndex457
											if buffer[position] != rune('C') {
												goto l456
											}
											position++
										}
									l457:
										{
											position459, tokenIndex459 := position, tokenIndex
											if buffer[position] != rune('h') {
												goto l460
											}
											position++
											goto l459
										l460:
											position, tokenIndex = position459, tokenIndex459
											if buffer[position] != rune('H') {
												goto l456
											}
											position++
										}
									l459:
										{
											position461, tokenIndex461 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l462
											}
											position++
											goto l461
										l462:
											position, tokenIndex = position461, tokenIndex461
											if buffer[position] != rune('A') {
												goto l456
											}
											position++
										}
									l461:
										{
											position463, tokenIndex463 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l464
											}
											position++
											goto l463
										l464:
											position, tokenIndex = position463, tokenIndex463
											if buffer[position] != rune('N') {
												goto l456
											}
											position++
										}
									l463:
										goto l324
									l456:
										position, tokenIndex = position324, tokenIndex324
										{
											position466, tokenIndex466 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l467
											}
											position++
											goto l466
										l467:
											position, tokenIndex = position466, tokenIndex466
											if buffer[position] != rune('P') {
												goto l465
											}
											position++
										}
									l466:
										{
											position468, tokenIndex468 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l469
											}
											position++
											goto l468
										l469:
											position, tokenIndex = position468, tokenIndex468
											if buffer[position] != rune('U') {
												goto l465
											}
											position++
										}
									l468:
										{
											position470, tokenIndex470 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l471
											}
											position++
											goto l470
										l471:
											position, tokenIndex = position470, tokenIndex470
											if buffer[position] != rune('R') {
												goto l465
											}
											position++
										}
									l470:
										{
											position472, tokenIndex472 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l473
											}
											position++
											goto l472
										l473:
											position, tokenIndex = position472, tokenIndex472
											if buffer[position] != rune('P') {
												goto l465
											}
											position++
										}
									l472:
										goto l324
									l465:
										position, tokenIndex = position324, tokenIndex324
										{
											position475, tokenIndex475 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l476
											}
											position++
											goto l475
										l476:
											position, tokenIndex = position475, tokenIndex475
											if buffer[position] != rune('R') {
												goto l474
											}
											position++
										}
									l475:
										{
											position477, tokenIndex477 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l478
											}
											position++
											goto l477
										l478:
											position, tokenIndex = position477, tokenIndex477
											if buffer[position] != rune('E') {
												goto l474
											}
											position++
										}
									l477:
										{
											position479, tokenIndex479 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l480
											}
											position++
											goto l479
										l480:
											position, tokenIndex = position479, tokenIndex479
											if buffer[position] != rune('F') {
												goto l474
											}
											position++
										}
									l479:
										{
											position481, tokenIndex481 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l482
											}
											position++
											goto l481
										l482:
											position, tokenIndex = position481, tokenIndex481
											if buffer[position] != rune('D') {
												goto l474
											}
											position++
										}
									l481:
										goto l324
									l474:
										position, tokenIndex = position324, tokenIndex324
										{
											position484, tokenIndex484 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l485
											}
											position++
											goto l484
										l485:
											position, tokenIndex = position484, tokenIndex484
											if buffer[position] != rune('C') {
												goto l483
											}
											position++
										}
									l484:
										{
											position486, tokenIndex486 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l487
											}
											position++
											goto l486
										l487:
											position, tokenIndex = position486, tokenIndex486
											if buffer[position] != rune('L') {
												goto l483
											}
											position++
										}
									l486:
										{
											position488, tokenIndex488 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l489
											}
											position++
											goto l488
										l489:
											position, tokenIndex = position488, tokenIndex488
											if buffer[position] != rune('A') {
												goto l483
											}
											position++
										}
									l488:
										{
											position490, tokenIndex490 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l491
											}
											position++
											goto l490
										l491:
											position, tokenIndex = position490, tokenIndex490
											if buffer[position] != rune('S') {
												goto l483
											}
											position++
										}
									l490:
										goto l324
									l483:
										position, tokenIndex = position324, tokenIndex324
										{
											position493, tokenIndex493 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l494
											}
											position++
											goto l493
										l494:
											position, tokenIndex = position493, tokenIndex493
											if buffer[position] != rune('C') {
												goto l492
											}
											position++
										}
									l493:
										{
											position495, tokenIndex495 := position, tokenIndex
											if buffer[position] != rune('l') {
												goto l496
											}
											position++
											goto l495
										l496:
											position, tokenIndex = position495, tokenIndex495
											if buffer[position] != rune('L') {
												goto l492
											}
											position++
										}
									l495:
										{
											position497, tokenIndex497 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l498
											}
											position++
											goto l497
										l498:
											position, tokenIndex = position497, tokenIndex497
											if buffer[position] != rune('S') {
												goto l492
											}
											position++
										}
									l497:
										goto l324
									l492:
										position, tokenIndex = position324, tokenIndex324
										{
											position500, tokenIndex500 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l501
											}
											position++
											goto l500
										l501:
											position, tokenIndex = position500, tokenIndex500
											if buffer[position] != rune('C') {
												goto l499
											}
											position++
										}
									l500:
										{
											position502, tokenIndex502 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l503
											}
											position++
											goto l502
										l503:
											position, tokenIndex = position502, tokenIndex502
											if buffer[position] != rune('S') {
												goto l499
											}
											position++
										}
									l502:
										{
											position504, tokenIndex504 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l505
											}
											position++
											goto l504
										l505:
											position, tokenIndex = position504, tokenIndex504
											if buffer[position] != rune('T') {
												goto l499
											}
											position++
										}
									l504:
										{
											position506, tokenIndex506 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l507
											}
											position++
											goto l506
										l507:
											position, tokenIndex = position506, tokenIndex506
											if buffer[position] != rune('S') {
												goto l499
											}
											position++
										}
									l506:
										goto l324
									l499:
										position, tokenIndex = position324, tokenIndex324
										{
											position509, tokenIndex509 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l510
											}
											position++
											goto l509
										l510:
											position, tokenIndex = position509, tokenIndex509
											if buffer[position] != rune('T') {
												goto l508
											}
											position++
										}
									l509:
										{
											position511, tokenIndex511 := position, tokenIndex
											if buffer[position] != rune('y') {
												goto l512
											}
											position++
											goto l511
										l512:
											position, tokenIndex = position511, tokenIndex511
											if buffer[position] != rune('Y') {
												goto l508
											}
											position++
										}
									l511:
										{
											position513, tokenIndex513 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l514
											}
											position++
											goto l513
										l514:
											position, tokenIndex = position513, tokenIndex513
											if buffer[position] != rune('P') {
												goto l508
											}
											position++
										}
									l513:
										{
											position515, tokenIndex515 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l516
											}
											position++
											goto l515
										l516:
											position, tokenIndex = position515, tokenIndex515
											if buffer[position] != rune('E') {
												goto l508
											}
											position++
										}
									l515:
										goto l324
									l508:
										position, tokenIndex = position324, tokenIndex324
										{
											position518, tokenIndex518 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l519
											}
											position++
											goto l518
										l519:
											position, tokenIndex = position518, tokenIndex518
											if buffer[position] != rune('M') {
												goto l517
											}
											position++
										}
									l518:
										{
											position520, tokenIndex520 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l521
											}
											position++
											goto l520
										l521:
											position, tokenIndex = position520, tokenIndex520
											if buffer[position] != rune('I') {
												goto l517
											}
											position++
										}
									l520:
										{
											position522, tokenIndex522 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l523
											}
											position++
											goto l522
										l523:
											position, tokenIndex = position522, tokenIndex522
											if buffer[position] != rune('N') {
												goto l517
											}
											position++
										}
									l522:
										goto l324
									l517:
										position, tokenIndex = position324, tokenIndex324
										{
											position525, tokenIndex525 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l526
											}
											position++
											goto l525
										l526:
											position, tokenIndex = position525, tokenIndex525
											if buffer[position] != rune('M') {
												goto l524
											}
											position++
										}
									l525:
										{
											position527, tokenIndex527 := position, tokenIndex
											if buffer[position] != rune('a') {
												goto l528
											}
											position++
											goto l527
										l528:
											position, tokenIndex = position527, tokenIndex527
											if buffer[position] != rune('A') {
												goto l524
											}
											position++
										}
									l527:
										{
											position529, tokenIndex529 := position, tokenIndex
											if buffer[position] != rune('x') {
												goto l530
											}
											position++
											goto l529
										l530:
											position, tokenIndex = position529, tokenIndex529
											if buffer[position] != rune('X') {
												goto l524
											}
											position++
										}
									l529:
										goto l324
									l524:
										position, tokenIndex = position324, tokenIndex324
										{
											position532, tokenIndex532 := position, tokenIndex
											if buffer[position] != rune('b') {
												goto l533
											}
											position++
											goto l532
										l533:
											position, tokenIndex = position532, tokenIndex532
											if buffer[position] != rune('B') {
												goto l531
											}
											position++
										}
									l532:
										{
											position534, tokenIndex534 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l535
											}
											position++
											goto l534
										l535:
											position, tokenIndex = position534, tokenIndex534
											if buffer[position] != rune('I') {
												goto l531
											}
											position++
										}
									l534:
										{
											position536, tokenIndex536 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l537
											}
											position++
											goto l536
										l537:
											position, tokenIndex = position536, tokenIndex536
											if buffer[position] != rune('N') {
												goto l531
											}
											position++
										}
									l536:
										{
											position538, tokenIndex538 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l539
											}
											position++
											goto l538
										l539:
											position, tokenIndex = position538, tokenIndex538
											if buffer[position] != rune('F') {
												goto l531
											}
											position++
										}
									l538:
										goto l324
									l531:
										position, tokenIndex = position324, tokenIndex324
										{
											position541, tokenIndex541 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l542
											}
											position++
											goto l541
										l542:
											position, tokenIndex = position541, tokenIndex541
											if buffer[position] != rune('S') {
												goto l540
											}
											position++
										}
									l541:
										{
											position543, tokenIndex543 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l544
											}
											position++
											goto l543
										l544:
											position, tokenIndex = position543, tokenIndex543
											if buffer[position] != rune('C') {
												goto l540
											}
											position++
										}
									l543:
										{
											position545, tokenIndex545 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l546
											}
											position++
											goto l545
										l546:
											position, tokenIndex = position545, tokenIndex545
											if buffer[position] != rune('P') {
												goto l540
											}
											position++
										}
									l545:
										{
											position547, tokenIndex547 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l548
											}
											position++
											goto l547
										l548:
											position, tokenIndex = position547, tokenIndex547
											if buffer[position] != rune('E') {
												goto l540
											}
											position++
										}
									l547:
										goto l324
									l540:
										position, tokenIndex = position324, tokenIndex324
										{
											position550, tokenIndex550 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l551
											}
											position++
											goto l550
										l551:
											position, tokenIndex = position550, tokenIndex550
											if buffer[position] != rune('I') {
												goto l549
											}
											position++
										}
									l550:
										{
											position552, tokenIndex552 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l553
											}
											position++
											goto l552
										l553:
											position, tokenIndex = position552, tokenIndex552
											if buffer[position] != rune('N') {
												goto l549
											}
											position++
										}
									l552:
										{
											position554, tokenIndex554 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l555
											}
											position++
											goto l554
										l555:
											position, tokenIndex = position554, tokenIndex554
											if buffer[position] != rune('I') {
												goto l549
											}
											position++
										}
									l554:
										{
											position556, tokenIndex556 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l557
											}
											position++
											goto l556
										l557:
											position, tokenIndex = position556, tokenIndex556
											if buffer[position] != rune('T') {
												goto l549
											}
											position++
										}
									l556:
										goto l324
									l549:
										position, tokenIndex = position324, tokenIndex324
										{
											switch buffer[position] {
											case 'V', 'v':
												{
													position559, tokenIndex559 := position, tokenIndex
													if buffer[position] != rune('v') {
														goto l560
													}
													position++
													goto l559
												l560:
													position, tokenIndex = position559, tokenIndex559
													if buffer[position] != rune('V') {
														goto l321
													}
													position++
												}
											l559:
												{
													position561, tokenIndex561 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l562
													}
													position++
													goto l561
												l562:
													position, tokenIndex = position561, tokenIndex561
													if buffer[position] != rune('A') {
														goto l321
													}
													position++
												}
											l561:
												{
													position563, tokenIndex563 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l564
													}
													position++
													goto l563
												l564:
													position, tokenIndex = position563, tokenIndex563
													if buffer[position] != rune('L') {
														goto l321
													}
													position++
												}
											l563:
												break
											case 'I', 'i':
												{
													position565, tokenIndex565 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l566
													}
													position++
													goto l565
												l566:
													position, tokenIndex = position565, tokenIndex565
													if buffer[position] != rune('I') {
														goto l321
													}
													position++
												}
											l565:
												{
													position567, tokenIndex567 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l568
													}
													position++
													goto l567
												l568:
													position, tokenIndex = position567, tokenIndex567
													if buffer[position] != rune('D') {
														goto l321
													}
													position++
												}
											l567:
												{
													position569, tokenIndex569 := position, tokenIndex
													if buffer[position] != rune('x') {
														goto l570
													}
													position++
													goto l569
												l570:
													position, tokenIndex = position569, tokenIndex569
													if buffer[position] != rune('X') {
														goto l321
													}
													position++
												}
											l569:
												break
											case 'S', 's':
												{
													position571, tokenIndex571 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l572
													}
													position++
													goto l571
												l572:
													position, tokenIndex = position571, tokenIndex571
													if buffer[position] != rune('S') {
														goto l321
													}
													position++
												}
											l571:
												{
													position573, tokenIndex573 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l574
													}
													position++
													goto l573
												l574:
													position, tokenIndex = position573, tokenIndex573
													if buffer[position] != rune('I') {
														goto l321
													}
													position++
												}
											l573:
												{
													position575, tokenIndex575 := position, tokenIndex
													if buffer[position] != rune('z') {
														goto l576
													}
													position++
													goto l575
												l576:
													position, tokenIndex = position575, tokenIndex575
													if buffer[position] != rune('Z') {
														goto l321
													}
													position++
												}
											l575:
												{
													position577, tokenIndex577 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l578
													}
													position++
													goto l577
												l578:
													position, tokenIndex = position577, tokenIndex577
													if buffer[position] != rune('E') {
														goto l321
													}
													position++
												}
											l577:
												break
											case 'B', 'b':
												{
													position579, tokenIndex579 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l580
													}
													position++
													goto l579
												l580:
													position, tokenIndex = position579, tokenIndex579
													if buffer[position] != rune('B') {
														goto l321
													}
													position++
												}
											l579:
												{
													position581, tokenIndex581 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l582
													}
													position++
													goto l581
												l582:
													position, tokenIndex = position581, tokenIndex581
													if buffer[position] != rune('P') {
														goto l321
													}
													position++
												}
											l581:
												{
													position583, tokenIndex583 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l584
													}
													position++
													goto l583
												l584:
													position, tokenIndex = position583, tokenIndex583
													if buffer[position] != rune('O') {
														goto l321
													}
													position++
												}
											l583:
												{
													position585, tokenIndex585 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l586
													}
													position++
													goto l585
												l586:
													position, tokenIndex = position585, tokenIndex585
													if buffer[position] != rune('S') {
														goto l321
													}
													position++
												}
											l585:
												break
											case 'C', 'c':
												{
													position587, tokenIndex587 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l588
													}
													position++
													goto l587
												l588:
													position, tokenIndex = position587, tokenIndex587
													if buffer[position] != rune('C') {
														goto l321
													}
													position++
												}
											l587:
												{
													position589, tokenIndex589 := position, tokenIndex
													if buffer[position] != rune('h') {
														goto l590
													}
													position++
													goto l589
												l590:
													position, tokenIndex = position589, tokenIndex589
													if buffer[position] != rune('H') {
														goto l321
													}
													position++
												}
											l589:
												{
													position591, tokenIndex591 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l592
													}
													position++
													goto l591
												l592:
													position, tokenIndex = position591, tokenIndex591
													if buffer[position] != rune('A') {
														goto l321
													}
													position++
												}
											l591:
												{
													position593, tokenIndex593 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l594
													}
													position++
													goto l593
												l594:
													position, tokenIndex = position593, tokenIndex593
													if buffer[position] != rune('I') {
														goto l321
													}
													position++
												}
											l593:
												{
													position595, tokenIndex595 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l596
													}
													position++
													goto l595
												l596:
													position, tokenIndex = position595, tokenIndex595
													if buffer[position] != rune('N') {
														goto l321
													}
													position++
												}
											l595:
												break
											case 'R', 'r':
												{
													position597, tokenIndex597 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l598
													}
													position++
													goto l597
												l598:
													position, tokenIndex = position597, tokenIndex597
													if buffer[position] != rune('R') {
														goto l321
													}
													position++
												}
											l597:
												{
													position599, tokenIndex599 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l600
													}
													position++
													goto l599
												l600:
													position, tokenIndex = position599, tokenIndex599
													if buffer[position] != rune('S') {
														goto l321
													}
													position++
												}
											l599:
												{
													position601, tokenIndex601 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l602
													}
													position++
													goto l601
												l602:
													position, tokenIndex = position601, tokenIndex601
													if buffer[position] != rune('L') {
														goto l321
													}
													position++
												}
											l601:
												{
													position603, tokenIndex603 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l604
													}
													position++
													goto l603
												l604:
													position, tokenIndex = position603, tokenIndex603
													if buffer[position] != rune('T') {
														goto l321
													}
													position++
												}
											l603:
												break
											case 'E', 'e':
												{
													position605, tokenIndex605 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l606
													}
													position++
													goto l605
												l606:
													position, tokenIndex = position605, tokenIndex605
													if buffer[position] != rune('E') {
														goto l321
													}
													position++
												}
											l605:
												{
													position607, tokenIndex607 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l608
													}
													position++
													goto l607
												l608:
													position, tokenIndex = position607, tokenIndex607
													if buffer[position] != rune('L') {
														goto l321
													}
													position++
												}
											l607:
												{
													position609, tokenIndex609 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l610
													}
													position++
													goto l609
												l610:
													position, tokenIndex = position609, tokenIndex609
													if buffer[position] != rune('T') {
														goto l321
													}
													position++
												}
											l609:
												{
													position611, tokenIndex611 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l612
													}
													position++
													goto l611
												l612:
													position, tokenIndex = position611, tokenIndex611
													if buffer[position] != rune('S') {
														goto l321
													}
													position++
												}
											l611:
												break
											case 'D', 'd':
												{
													position613, tokenIndex613 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l614
													}
													position++
													goto l613
												l614:
													position, tokenIndex = position613, tokenIndex613
													if buffer[position] != rune('D') {
														goto l321
													}
													position++
												}
											l613:
												{
													position615, tokenIndex615 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l616
													}
													position++
													goto l615
												l616:
													position, tokenIndex = position615, tokenIndex615
													if buffer[position] != rune('O') {
														goto l321
													}
													position++
												}
											l615:
												{
													position617, tokenIndex617 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l618
													}
													position++
													goto l617
												l618:
													position, tokenIndex = position617, tokenIndex617
													if buffer[position] != rune('M') {
														goto l321
													}
													position++
												}
											l617:
												{
													position619, tokenIndex619 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l620
													}
													position++
													goto l619
												l620:
													position, tokenIndex = position619, tokenIndex619
													if buffer[position] != rune('N') {
														goto l321
													}
													position++
												}
											l619:
												break
											case 'T', 't':
												{
													position621, tokenIndex621 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l622
													}
													position++
													goto l621
												l622:
													position, tokenIndex = position621, tokenIndex621
													if buffer[position] != rune('T') {
														goto l321
													}
													position++
												}
											l621:
												{
													position623, tokenIndex623 := position, tokenIndex
													if buffer[position] != rune('y') {
														goto l624
													}
													position++
													goto l623
												l624:
													position, tokenIndex = position623, tokenIndex623
													if buffer[position] != rune('Y') {
														goto l321
													}
													position++
												}
											l623:
												{
													position625, tokenIndex625 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l626
													}
													position++
													goto l625
												l626:
													position, tokenIndex = position625, tokenIndex625
													if buffer[position] != rune('P') {
														goto l321
													}
													position++
												}
											l625:
												{
													position627, tokenIndex627 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l628
													}
													position++
													goto l627
												l628:
													position, tokenIndex = position627, tokenIndex627
													if buffer[position] != rune('E') {
														goto l321
													}
													position++
												}
											l627:
												break
											case 'M', 'm':
												{
													position629, tokenIndex629 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l630
													}
													position++
													goto l629
												l630:
													position, tokenIndex = position629, tokenIndex629
													if buffer[position] != rune('M') {
														goto l321
													}
													position++
												}
											l629:
												{
													position631, tokenIndex631 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l632
													}
													position++
													goto l631
												l632:
													position, tokenIndex = position631, tokenIndex631
													if buffer[position] != rune('N') {
														goto l321
													}
													position++
												}
											l631:
												{
													position633, tokenIndex633 := position, tokenIndex
													if buffer[position] != rune('g') {
														goto l634
													}
													position++
													goto l633
												l634:
													position, tokenIndex = position633, tokenIndex633
													if buffer[position] != rune('G') {
														goto l321
													}
													position++
												}
											l633:
												{
													position635, tokenIndex635 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l636
													}
													position++
													goto l635
												l636:
													position, tokenIndex = position635, tokenIndex635
													if buffer[position] != rune('L') {
														goto l321
													}
													position++
												}
											l635:
												break
											case 'N', 'n':
												{
													position637, tokenIndex637 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l638
													}
													position++
													goto l637
												l638:
													position, tokenIndex = position637, tokenIndex637
													if buffer[position] != rune('N') {
														goto l321
													}
													position++
												}
											l637:
												{
													position639, tokenIndex639 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l640
													}
													position++
													goto l639
												l640:
													position, tokenIndex = position639, tokenIndex639
													if buffer[position] != rune('A') {
														goto l321
													}
													position++
												}
											l639:
												{
													position641, tokenIndex641 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l642
													}
													position++
													goto l641
												l642:
													position, tokenIndex = position641, tokenIndex641
													if buffer[position] != rune('M') {
														goto l321
													}
													position++
												}
											l641:
												{
													position643, tokenIndex643 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l644
													}
													position++
													goto l643
												l644:
													position, tokenIndex = position643, tokenIndex643
													if buffer[position] != rune('E') {
														goto l321
													}
													position++
												}
											l643:
												break
											case 'U', 'u':
												{
													position645, tokenIndex645 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l646
													}
													position++
													goto l645
												l646:
													position, tokenIndex = position645, tokenIndex645
													if buffer[position] != rune('U') {
														goto l321
													}
													position++
												}
											l645:
												{
													position647, tokenIndex647 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l648
													}
													position++
													goto l647
												l648:
													position, tokenIndex = position647, tokenIndex647
													if buffer[position] != rune('N') {
														goto l321
													}
													position++
												}
											l647:
												{
													position649, tokenIndex649 := position, tokenIndex
													if buffer[position] != rune('q') {
														goto l650
													}
													position++
													goto l649
												l650:
													position, tokenIndex = position649, tokenIndex649
													if buffer[position] != rune('Q') {
														goto l321
													}
													position++
												}
											l649:
												{
													position651, tokenIndex651 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l652
													}
													position++
													goto l651
												l652:
													position, tokenIndex = position651, tokenIndex651
													if buffer[position] != rune('L') {
														goto l321
													}
													position++
												}
											l651:
												break
											case 'L', 'l':
												{
													position653, tokenIndex653 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l654
													}
													position++
													goto l653
												l654:
													position, tokenIndex = position653, tokenIndex653
													if buffer[position] != rune('L') {
														goto l321
													}
													position++
												}
											l653:
												{
													position655, tokenIndex655 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l656
													}
													position++
													goto l655
												l656:
													position, tokenIndex = position655, tokenIndex655
													if buffer[position] != rune('A') {
														goto l321
													}
													position++
												}
											l655:
												{
													position657, tokenIndex657 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l658
													}
													position++
													goto l657
												l658:
													position, tokenIndex = position657, tokenIndex657
													if buffer[position] != rune('B') {
														goto l321
													}
													position++
												}
											l657:
												{
													position659, tokenIndex659 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l660
													}
													position++
													goto l659
												l660:
													position, tokenIndex = position659, tokenIndex659
													if buffer[position] != rune('L') {
														goto l321
													}
													position++
												}
											l659:
												break
											case 'P', 'p':
												{
													position661, tokenIndex661 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l662
													}
													position++
													goto l661
												l662:
													position, tokenIndex = position661, tokenIndex661
													if buffer[position] != rune('P') {
														goto l321
													}
													position++
												}
											l661:
												{
													position663, tokenIndex663 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l664
													}
													position++
													goto l663
												l664:
													position, tokenIndex = position663, tokenIndex663
													if buffer[position] != rune('T') {
														goto l321
													}
													position++
												}
											l663:
												{
													position665, tokenIndex665 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l666
													}
													position++
													goto l665
												l666:
													position, tokenIndex = position665, tokenIndex665
													if buffer[position] != rune('D') {
														goto l321
													}
													position++
												}
											l665:
												break
											case 'F', 'f':
												{
													position667, tokenIndex667 := position, tokenIndex
													if buffer[position] != rune('f') {
														goto l668
													}
													position++
													goto l667
												l668:
													position, tokenIndex = position667, tokenIndex667
													if buffer[position] != rune('F') {
														goto l321
													}
													position++
												}
											l667:
												{
													position669, tokenIndex669 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l670
													}
													position++
													goto l669
												l670:
													position, tokenIndex = position669, tokenIndex669
													if buffer[position] != rune('L') {
														goto l321
													}
													position++
												}
											l669:
												{
													position671, tokenIndex671 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l672
													}
													position++
													goto l671
												l672:
													position, tokenIndex = position671, tokenIndex671
													if buffer[position] != rune('D') {
														goto l321
													}
													position++
												}
											l671:
												{
													position673, tokenIndex673 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l674
													}
													position++
													goto l673
												l674:
													position, tokenIndex = position673, tokenIndex673
													if buffer[position] != rune('S') {
														goto l321
													}
													position++
												}
											l673:
												break
											case 'A', 'a':
												{
													position675, tokenIndex675 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l676
													}
													position++
													goto l675
												l676:
													position, tokenIndex = position675, tokenIndex675
													if buffer[position] != rune('A') {
														goto l321
													}
													position++
												}
											l675:
												{
													position677, tokenIndex677 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l678
													}
													position++
													goto l677
												l678:
													position, tokenIndex = position677, tokenIndex677
													if buffer[position] != rune('R') {
														goto l321
													}
													position++
												}
											l677:
												{
													position679, tokenIndex679 := position, tokenIndex
													if buffer[position] != rune('g') {
														goto l680
													}
													position++
													goto l679
												l680:
													position, tokenIndex = position679, tokenIndex679
													if buffer[position] != rune('G') {
														goto l321
													}
													position++
												}
											l679:
												{
													position681, tokenIndex681 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l682
													}
													position++
													goto l681
												l682:
													position, tokenIndex = position681, tokenIndex681
													if buffer[position] != rune('T') {
														goto l321
													}
													position++
												}
											l681:
												break
											case 'O', 'o':
												{
													position683, tokenIndex683 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l684
													}
													position++
													goto l683
												l684:
													position, tokenIndex = position683, tokenIndex683
													if buffer[position] != rune('O') {
														goto l321
													}
													position++
												}
											l683:
												{
													position685, tokenIndex685 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l686
													}
													position++
													goto l685
												l686:
													position, tokenIndex = position685, tokenIndex685
													if buffer[position] != rune('P') {
														goto l321
													}
													position++
												}
											l685:
												if buffer[position] != rune('0') {
													goto l321
												}
												position++
												break
											default:
												if !_rules[ruleInteger]() {
													goto l321
												}
												break
											}
										}

									}
								l324:
									add(rulePegText, position323)
								}
								{
									add(ruleAction23, position)
								}
								add(ruleNodeFieldName, position322)
							}
							goto l320
						l321:
							position, tokenIndex = position320, tokenIndex320
							{
								position689 := position
								{
									position690, tokenIndex690 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l691
									}
									position++
									goto l690
								l691:
									position, tokenIndex = position690, tokenIndex690
									if buffer[position] != rune('O') {
										goto l688
									}
									position++
								}
							l690:
								{
									position692, tokenIndex692 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l693
									}
									position++
									goto l692
								l693:
									position, tokenIndex = position692, tokenIndex692
									if buffer[position] != rune('P') {
										goto l688
									}
									position++
								}
							l692:
								if buffer[position] != rune(' ') {
									goto l688
								}
								position++
								{
									position694 := position
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l688
									}
									position++
									add(rulePegText, position694)
								}
								{
									add(ruleAction5, position)
								}
								add(ruleOpAttr, position689)
							}
							goto l320
						l688:
							position, tokenIndex = position320, tokenIndex320
							{
								position697 := position
								{
									position698 := position
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l696
									}
									position++
									add(rulePegText, position698)
								}
								{
									add(ruleAction6, position)
								}
								add(ruleOpAttr2, position697)
							}
							goto l320
						l696:
							position, tokenIndex = position320, tokenIndex320
							{
								position700 := position
								{
									position701 := position
									if c := buffer[position]; c < rune('a') || c > rune('z') {
										goto l318
									}
									position++
								l702:
									{
										position703, tokenIndex703 := position, tokenIndex
										if c := buffer[position]; c < rune('a') || c > rune('z') {
											goto l703
										}
										position++
										goto l702
									l703:
										position, tokenIndex = position703, tokenIndex703
									}
									add(rulePegText, position701)
								}
								{
									add(ruleAction22, position)
								}
								add(ruleOtherField, position700)
							}
						}
					l320:
						if !_rules[rulews]() {
							goto l318
						}
						if buffer[position] != rune(':') {
							goto l318
						}
						position++
						if !_rules[rulews]() {
							goto l318
						}
						{
							position705 := position
							if !_rules[ruleNode]() {
								goto l318
							}
							add(rulePegText, position705)
						}
						if !_rules[rulews]() {
							goto l318
						}
						{
							add(ruleAction24, position)
						}
						add(ruleNodeAttr, position319)
					}
					goto l91
				l318:
					position, tokenIndex = position91, tokenIndex91
					{
						position708 := position
						{
							position709 := position
							if buffer[position] != rune('s') {
								goto l707
							}
							position++
							if buffer[position] != rune('r') {
								goto l707
							}
							position++
							if buffer[position] != rune('c') {
								goto l707
							}
							position++
							if buffer[position] != rune('p') {
								goto l707
							}
							position++
							add(rulePegText, position709)
						}
						if buffer[position] != rune(':') {
							goto l707
						}
						position++
						if !_rules[rulews]() {
							goto l707
						}
						{
							position710 := position
							{
								position711 := position
								{
									position712, tokenIndex712 := position, tokenIndex
									if buffer[position] != rune('_') {
										goto l712
									}
									position++
									goto l713
								l712:
									position, tokenIndex = position712, tokenIndex712
								}
							l713:
								{
									position714 := position
									{
										position715 := position
										{
											position716 := position
											{
												position717, tokenIndex717 := position, tokenIndex
												{
													switch buffer[position] {
													case '_':
														if buffer[position] != rune('_') {
															goto l718
														}
														position++
														break
													case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
														if c := buffer[position]; c < rune('0') || c > rune('9') {
															goto l718
														}
														position++
														break
													case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
														if c := buffer[position]; c < rune('a') || c > rune('z') {
															goto l718
														}
														position++
														break
													case '-':
														if buffer[position] != rune('-') {
															goto l718
														}
														position++
														break
													case '+':
														if buffer[position] != rune('+') {
															goto l718
														}
														position++
														break
													default:
														if c := buffer[position]; c < rune('A') || c > rune('Z') {
															goto l718
														}
														position++
														break
													}
												}

											l719:
												{
													position720, tokenIndex720 := position, tokenIndex
													{
														switch buffer[position] {
														case '_':
															if buffer[position] != rune('_') {
																goto l720
															}
															position++
															break
														case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
															if c := buffer[position]; c < rune('0') || c > rune('9') {
																goto l720
															}
															position++
															break
														case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
															if c := buffer[position]; c < rune('a') || c > rune('z') {
																goto l720
															}
															position++
															break
														case '-':
															if buffer[position] != rune('-') {
																goto l720
															}
															position++
															break
														case '+':
															if buffer[position] != rune('+') {
																goto l720
															}
															position++
															break
														default:
															if c := buffer[position]; c < rune('A') || c > rune('Z') {
																goto l720
															}
															position++
															break
														}
													}

													goto l719
												l720:
													position, tokenIndex = position720, tokenIndex720
												}
											l723:
												{
													position724, tokenIndex724 := position, tokenIndex
													if buffer[position] != rune('.') {
														goto l724
													}
													position++
													{
														switch buffer[position] {
														case '_':
															if buffer[position] != rune('_') {
																goto l724
															}
															position++
															break
														case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
															if c := buffer[position]; c < rune('a') || c > rune('z') {
																goto l724
															}
															position++
															break
														default:
															if c := buffer[position]; c < rune('A') || c > rune('Z') {
																goto l724
															}
															position++
															break
														}
													}

												l725:
													{
														position726, tokenIndex726 := position, tokenIndex
														{
															switch buffer[position] {
															case '_':
																if buffer[position] != rune('_') {
																	goto l726
																}
																position++
																break
															case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
																if c := buffer[position]; c < rune('a') || c > rune('z') {
																	goto l726
																}
																position++
																break
															default:
																if c := buffer[position]; c < rune('A') || c > rune('Z') {
																	goto l726
																}
																position++
																break
															}
														}

														goto l725
													l726:
														position, tokenIndex = position726, tokenIndex726
													}
													goto l723
												l724:
													position, tokenIndex = position724, tokenIndex724
												}
												goto l717
											l718:
												position, tokenIndex = position717, tokenIndex717
												if buffer[position] != rune('<') {
													goto l707
												}
												position++
												{
													position729, tokenIndex729 := position, tokenIndex
													if buffer[position] != rune('b') {
														goto l730
													}
													position++
													goto l729
												l730:
													position, tokenIndex = position729, tokenIndex729
													if buffer[position] != rune('B') {
														goto l707
													}
													position++
												}
											l729:
												{
													position731, tokenIndex731 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l732
													}
													position++
													goto l731
												l732:
													position, tokenIndex = position731, tokenIndex731
													if buffer[position] != rune('U') {
														goto l707
													}
													position++
												}
											l731:
												{
													position733, tokenIndex733 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l734
													}
													position++
													goto l733
												l734:
													position, tokenIndex = position733, tokenIndex733
													if buffer[position] != rune('I') {
														goto l707
													}
													position++
												}
											l733:
												{
													position735, tokenIndex735 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l736
													}
													position++
													goto l735
												l736:
													position, tokenIndex = position735, tokenIndex735
													if buffer[position] != rune('L') {
														goto l707
													}
													position++
												}
											l735:
												{
													position737, tokenIndex737 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l738
													}
													position++
													goto l737
												l738:
													position, tokenIndex = position737, tokenIndex737
													if buffer[position] != rune('T') {
														goto l707
													}
													position++
												}
											l737:
												if buffer[position] != rune('-') {
													goto l707
												}
												position++
												{
													position739, tokenIndex739 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l740
													}
													position++
													goto l739
												l740:
													position, tokenIndex = position739, tokenIndex739
													if buffer[position] != rune('I') {
														goto l707
													}
													position++
												}
											l739:
												{
													position741, tokenIndex741 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l742
													}
													position++
													goto l741
												l742:
													position, tokenIndex = position741, tokenIndex741
													if buffer[position] != rune('N') {
														goto l707
													}
													position++
												}
											l741:
												if buffer[position] != rune('>') {
													goto l707
												}
												position++
											}
										l717:
											add(rulePegText, position716)
										}
										{
											add(ruleAction8, position)
										}
										add(ruleFileName, position715)
									}
									add(rulePegText, position714)
								}
								if buffer[position] != rune(':') {
									goto l707
								}
								position++
								{
									position744 := position
									{
										position745 := position
										{
											position746 := position
											if !_rules[ruleInteger]() {
												goto l707
											}
											add(rulePegText, position746)
										}
										{
											add(ruleAction7, position)
										}
										add(ruleLineNumber, position745)
									}
									add(rulePegText, position744)
								}
								if !_rules[rulews]() {
									goto l707
								}
								{
									add(ruleAction9, position)
								}
								add(ruleFileRef, position711)
							}
							add(rulePegText, position710)
						}
						add(ruleSourceAttr, position708)
					}
					goto l91
				l707:
					position, tokenIndex = position91, tokenIndex91
					{
						position750 := position
						{
							position751 := position
							{
								switch buffer[position] {
								case 'P', 'p':
									{
										position753, tokenIndex753 := position, tokenIndex
										if buffer[position] != rune('p') {
											goto l754
										}
										position++
										goto l753
									l754:
										position, tokenIndex = position753, tokenIndex753
										if buffer[position] != rune('P') {
											goto l749
										}
										position++
									}
								l753:
									{
										position755, tokenIndex755 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l756
										}
										position++
										goto l755
									l756:
										position, tokenIndex = position755, tokenIndex755
										if buffer[position] != rune('R') {
											goto l749
										}
										position++
									}
								l755:
									{
										position757, tokenIndex757 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l758
										}
										position++
										goto l757
									l758:
										position, tokenIndex = position757, tokenIndex757
										if buffer[position] != rune('E') {
											goto l749
										}
										position++
									}
								l757:
									{
										position759, tokenIndex759 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l760
										}
										position++
										goto l759
									l760:
										position, tokenIndex = position759, tokenIndex759
										if buffer[position] != rune('C') {
											goto l749
										}
										position++
									}
								l759:
									break
								case 'C', 'c':
									{
										position761, tokenIndex761 := position, tokenIndex
										if buffer[position] != rune('c') {
											goto l762
										}
										position++
										goto l761
									l762:
										position, tokenIndex = position761, tokenIndex761
										if buffer[position] != rune('C') {
											goto l749
										}
										position++
									}
								l761:
									{
										position763, tokenIndex763 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l764
										}
										position++
										goto l763
									l764:
										position, tokenIndex = position763, tokenIndex763
										if buffer[position] != rune('T') {
											goto l749
										}
										position++
									}
								l763:
									{
										position765, tokenIndex765 := position, tokenIndex
										if buffer[position] != rune('o') {
											goto l766
										}
										position++
										goto l765
									l766:
										position, tokenIndex = position765, tokenIndex765
										if buffer[position] != rune('O') {
											goto l749
										}
										position++
									}
								l765:
									{
										position767, tokenIndex767 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l768
										}
										position++
										goto l767
									l768:
										position, tokenIndex = position767, tokenIndex767
										if buffer[position] != rune('R') {
											goto l749
										}
										position++
									}
								l767:
									break
								case 'U', 'u':
									{
										position769, tokenIndex769 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l770
										}
										position++
										goto l769
									l770:
										position, tokenIndex = position769, tokenIndex769
										if buffer[position] != rune('U') {
											goto l749
										}
										position++
									}
								l769:
									{
										position771, tokenIndex771 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l772
										}
										position++
										goto l771
									l772:
										position, tokenIndex = position771, tokenIndex771
										if buffer[position] != rune('S') {
											goto l749
										}
										position++
									}
								l771:
									{
										position773, tokenIndex773 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l774
										}
										position++
										goto l773
									l774:
										position, tokenIndex = position773, tokenIndex773
										if buffer[position] != rune('E') {
											goto l749
										}
										position++
									}
								l773:
									{
										position775, tokenIndex775 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l776
										}
										position++
										goto l775
									l776:
										position, tokenIndex = position775, tokenIndex775
										if buffer[position] != rune('D') {
											goto l749
										}
										position++
									}
								l775:
									break
								case 'B', 'b':
									{
										position777, tokenIndex777 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l778
										}
										position++
										goto l777
									l778:
										position, tokenIndex = position777, tokenIndex777
										if buffer[position] != rune('B') {
											goto l749
										}
										position++
									}
								l777:
									{
										position779, tokenIndex779 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l780
										}
										position++
										goto l779
									l780:
										position, tokenIndex = position779, tokenIndex779
										if buffer[position] != rune('A') {
											goto l749
										}
										position++
									}
								l779:
									{
										position781, tokenIndex781 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l782
										}
										position++
										goto l781
									l782:
										position, tokenIndex = position781, tokenIndex781
										if buffer[position] != rune('S') {
											goto l749
										}
										position++
									}
								l781:
									{
										position783, tokenIndex783 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l784
										}
										position++
										goto l783
									l784:
										position, tokenIndex = position783, tokenIndex783
										if buffer[position] != rune('E') {
											goto l749
										}
										position++
									}
								l783:
									{
										position785, tokenIndex785 := position, tokenIndex
										if buffer[position] != rune('s') {
											goto l786
										}
										position++
										goto l785
									l786:
										position, tokenIndex = position785, tokenIndex785
										if buffer[position] != rune('S') {
											goto l749
										}
										position++
									}
								l785:
									break
								default:
									{
										position787, tokenIndex787 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l788
										}
										position++
										goto l787
									l788:
										position, tokenIndex = position787, tokenIndex787
										if buffer[position] != rune('L') {
											goto l749
										}
										position++
									}
								l787:
									{
										position789, tokenIndex789 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l790
										}
										position++
										goto l789
									l790:
										position, tokenIndex = position789, tokenIndex789
										if buffer[position] != rune('I') {
											goto l749
										}
										position++
									}
								l789:
									{
										position791, tokenIndex791 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l792
										}
										position++
										goto l791
									l792:
										position, tokenIndex = position791, tokenIndex791
										if buffer[position] != rune('N') {
											goto l749
										}
										position++
									}
								l791:
									{
										position793, tokenIndex793 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l794
										}
										position++
										goto l793
									l794:
										position, tokenIndex = position793, tokenIndex793
										if buffer[position] != rune('E') {
											goto l749
										}
										position++
									}
								l793:
									break
								}
							}

							add(rulePegText, position751)
						}
						if !_rules[rulews]() {
							goto l749
						}
						if buffer[position] != rune(':') {
							goto l749
						}
						position++
						if !_rules[rulews]() {
							goto l749
						}
						{
							position795 := position
							if !_rules[ruleInteger]() {
								goto l749
							}
							add(rulePegText, position795)
						}
						if !_rules[rulews]() {
							goto l749
						}
						{
							add(ruleAction11, position)
						}
						add(ruleIntAttr, position750)
					}
					goto l91
				l749:
					position, tokenIndex = position91, tokenIndex91
					{
						position798 := position
						{
							position799 := position
							{
								position800, tokenIndex800 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l801
								}
								position++
								goto l800
							l801:
								position, tokenIndex = position800, tokenIndex800
								if buffer[position] != rune('S') {
									goto l797
								}
								position++
							}
						l800:
							{
								position802, tokenIndex802 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l803
								}
								position++
								goto l802
							l803:
								position, tokenIndex = position802, tokenIndex802
								if buffer[position] != rune('I') {
									goto l797
								}
								position++
							}
						l802:
							{
								position804, tokenIndex804 := position, tokenIndex
								if buffer[position] != rune('g') {
									goto l805
								}
								position++
								goto l804
							l805:
								position, tokenIndex = position804, tokenIndex804
								if buffer[position] != rune('G') {
									goto l797
								}
								position++
							}
						l804:
							{
								position806, tokenIndex806 := position, tokenIndex
								if buffer[position] != rune('n') {
									goto l807
								}
								position++
								goto l806
							l807:
								position, tokenIndex = position806, tokenIndex806
								if buffer[position] != rune('N') {
									goto l797
								}
								position++
							}
						l806:
							add(rulePegText, position799)
						}
						if !_rules[rulews]() {
							goto l797
						}
						if buffer[position] != rune(':') {
							goto l797
						}
						position++
						if !_rules[rulews]() {
							goto l797
						}
						{
							position808 := position
							{
								position809, tokenIndex809 := position, tokenIndex
								{
									position811, tokenIndex811 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l812
									}
									position++
									goto l811
								l812:
									position, tokenIndex = position811, tokenIndex811
									if buffer[position] != rune('S') {
										goto l810
									}
									position++
								}
							l811:
								{
									position813, tokenIndex813 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l814
									}
									position++
									goto l813
								l814:
									position, tokenIndex = position813, tokenIndex813
									if buffer[position] != rune('I') {
										goto l810
									}
									position++
								}
							l813:
								{
									position815, tokenIndex815 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l816
									}
									position++
									goto l815
								l816:
									position, tokenIndex = position815, tokenIndex815
									if buffer[position] != rune('G') {
										goto l810
									}
									position++
								}
							l815:
								{
									position817, tokenIndex817 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l818
									}
									position++
									goto l817
								l818:
									position, tokenIndex = position817, tokenIndex817
									if buffer[position] != rune('N') {
										goto l810
									}
									position++
								}
							l817:
								{
									position819, tokenIndex819 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l820
									}
									position++
									goto l819
								l820:
									position, tokenIndex = position819, tokenIndex819
									if buffer[position] != rune('E') {
										goto l810
									}
									position++
								}
							l819:
								{
									position821, tokenIndex821 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l822
									}
									position++
									goto l821
								l822:
									position, tokenIndex = position821, tokenIndex821
									if buffer[position] != rune('D') {
										goto l810
									}
									position++
								}
							l821:
								goto l809
							l810:
								position, tokenIndex = position809, tokenIndex809
								{
									position823, tokenIndex823 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l824
									}
									position++
									goto l823
								l824:
									position, tokenIndex = position823, tokenIndex823
									if buffer[position] != rune('U') {
										goto l797
									}
									position++
								}
							l823:
								{
									position825, tokenIndex825 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l826
									}
									position++
									goto l825
								l826:
									position, tokenIndex = position825, tokenIndex825
									if buffer[position] != rune('N') {
										goto l797
									}
									position++
								}
							l825:
								{
									position827, tokenIndex827 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l828
									}
									position++
									goto l827
								l828:
									position, tokenIndex = position827, tokenIndex827
									if buffer[position] != rune('S') {
										goto l797
									}
									position++
								}
							l827:
								{
									position829, tokenIndex829 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l830
									}
									position++
									goto l829
								l830:
									position, tokenIndex = position829, tokenIndex829
									if buffer[position] != rune('I') {
										goto l797
									}
									position++
								}
							l829:
								{
									position831, tokenIndex831 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l832
									}
									position++
									goto l831
								l832:
									position, tokenIndex = position831, tokenIndex831
									if buffer[position] != rune('G') {
										goto l797
									}
									position++
								}
							l831:
								{
									position833, tokenIndex833 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l834
									}
									position++
									goto l833
								l834:
									position, tokenIndex = position833, tokenIndex833
									if buffer[position] != rune('N') {
										goto l797
									}
									position++
								}
							l833:
								{
									position835, tokenIndex835 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l836
									}
									position++
									goto l835
								l836:
									position, tokenIndex = position835, tokenIndex835
									if buffer[position] != rune('E') {
										goto l797
									}
									position++
								}
							l835:
								{
									position837, tokenIndex837 := position, tokenIndex
									if buffer[position] != rune('d') {
										goto l838
									}
									position++
									goto l837
								l838:
									position, tokenIndex = position837, tokenIndex837
									if buffer[position] != rune('D') {
										goto l797
									}
									position++
								}
							l837:
							}
						l809:
							add(rulePegText, position808)
						}
						{
							add(ruleAction21, position)
						}
						add(ruleSignAttr, position798)
					}
					goto l91
				l797:
					position, tokenIndex = position91, tokenIndex91
					{
						position841 := position
						{
							position842 := position
							{
								position843, tokenIndex843 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l844
								}
								position++
								goto l843
							l844:
								position, tokenIndex = position843, tokenIndex843
								if buffer[position] != rune('L') {
									goto l840
								}
								position++
							}
						l843:
							{
								position845, tokenIndex845 := position, tokenIndex
								if buffer[position] != rune('o') {
									goto l846
								}
								position++
								goto l845
							l846:
								position, tokenIndex = position845, tokenIndex845
								if buffer[position] != rune('O') {
									goto l840
								}
								position++
							}
						l845:
							{
								position847, tokenIndex847 := position, tokenIndex
								if buffer[position] != rune('w') {
									goto l848
								}
								position++
								goto l847
							l848:
								position, tokenIndex = position847, tokenIndex847
								if buffer[position] != rune('W') {
									goto l840
								}
								position++
							}
						l847:
							add(rulePegText, position842)
						}
						if !_rules[rulews]() {
							goto l840
						}
						if buffer[position] != rune(':') {
							goto l840
						}
						position++
						if !_rules[rulews]() {
							goto l840
						}
						{
							position849, tokenIndex849 := position, tokenIndex
							if !_rules[ruleHexValue]() {
								goto l850
							}
							goto l849
						l850:
							position, tokenIndex = position849, tokenIndex849
							{
								switch buffer[position] {
								case '-':
									if !_rules[ruleNegInt]() {
										goto l840
									}
									break
								case '@':
									if !_rules[ruleNode]() {
										goto l840
									}
									break
								default:
									if !_rules[ruleInteger]() {
										goto l840
									}
									break
								}
							}

						}
					l849:
						if !_rules[rulews]() {
							goto l840
						}
						{
							add(ruleAction12, position)
						}
						add(ruleIntAttrLow, position841)
					}
					goto l91
				l840:
					position, tokenIndex = position91, tokenIndex91
					{
						position854 := position
						{
							position855 := position
							{
								position856, tokenIndex856 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l857
								}
								position++
								goto l856
							l857:
								position, tokenIndex = position856, tokenIndex856
								if buffer[position] != rune('A') {
									goto l853
								}
								position++
							}
						l856:
							{
								position858, tokenIndex858 := position, tokenIndex
								if buffer[position] != rune('c') {
									goto l859
								}
								position++
								goto l858
							l859:
								position, tokenIndex = position858, tokenIndex858
								if buffer[position] != rune('C') {
									goto l853
								}
								position++
							}
						l858:
							{
								position860, tokenIndex860 := position, tokenIndex
								if buffer[position] != rune('c') {
									goto l861
								}
								position++
								goto l860
							l861:
								position, tokenIndex = position860, tokenIndex860
								if buffer[position] != rune('C') {
									goto l853
								}
								position++
							}
						l860:
							{
								position862, tokenIndex862 := position, tokenIndex
								if buffer[position] != rune('s') {
									goto l863
								}
								position++
								goto l862
							l863:
								position, tokenIndex = position862, tokenIndex862
								if buffer[position] != rune('S') {
									goto l853
								}
								position++
							}
						l862:
							add(rulePegText, position855)
						}
						if !_rules[rulews]() {
							goto l853
						}
						if buffer[position] != rune(':') {
							goto l853
						}
						position++
						if !_rules[rulews]() {
							goto l853
						}
						{
							position864 := position
							{
								position865, tokenIndex865 := position, tokenIndex
								{
									position867, tokenIndex867 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l868
									}
									position++
									goto l867
								l868:
									position, tokenIndex = position867, tokenIndex867
									if buffer[position] != rune('P') {
										goto l866
									}
									position++
								}
							l867:
								{
									position869, tokenIndex869 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l870
									}
									position++
									goto l869
								l870:
									position, tokenIndex = position869, tokenIndex869
									if buffer[position] != rune('R') {
										goto l866
									}
									position++
								}
							l869:
								{
									position871, tokenIndex871 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l872
									}
									position++
									goto l871
								l872:
									position, tokenIndex = position871, tokenIndex871
									if buffer[position] != rune('I') {
										goto l866
									}
									position++
								}
							l871:
								{
									position873, tokenIndex873 := position, tokenIndex
									if buffer[position] != rune('v') {
										goto l874
									}
									position++
									goto l873
								l874:
									position, tokenIndex = position873, tokenIndex873
									if buffer[position] != rune('V') {
										goto l866
									}
									position++
								}
							l873:
								goto l865
							l866:
								position, tokenIndex = position865, tokenIndex865
								{
									position876, tokenIndex876 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l877
									}
									position++
									goto l876
								l877:
									position, tokenIndex = position876, tokenIndex876
									if buffer[position] != rune('P') {
										goto l875
									}
									position++
								}
							l876:
								{
									position878, tokenIndex878 := position, tokenIndex
									if buffer[position] != rune('u') {
										goto l879
									}
									position++
									goto l878
								l879:
									position, tokenIndex = position878, tokenIndex878
									if buffer[position] != rune('U') {
										goto l875
									}
									position++
								}
							l878:
								{
									position880, tokenIndex880 := position, tokenIndex
									if buffer[position] != rune('b') {
										goto l881
									}
									position++
									goto l880
								l881:
									position, tokenIndex = position880, tokenIndex880
									if buffer[position] != rune('B') {
										goto l875
									}
									position++
								}
							l880:
								goto l865
							l875:
								position, tokenIndex = position865, tokenIndex865
								{
									position882, tokenIndex882 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l883
									}
									position++
									goto l882
								l883:
									position, tokenIndex = position882, tokenIndex882
									if buffer[position] != rune('P') {
										goto l853
									}
									position++
								}
							l882:
								{
									position884, tokenIndex884 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l885
									}
									position++
									goto l884
								l885:
									position, tokenIndex = position884, tokenIndex884
									if buffer[position] != rune('R') {
										goto l853
									}
									position++
								}
							l884:
								{
									position886, tokenIndex886 := position, tokenIndex
									if buffer[position] != rune('o') {
										goto l887
									}
									position++
									goto l886
								l887:
									position, tokenIndex = position886, tokenIndex886
									if buffer[position] != rune('O') {
										goto l853
									}
									position++
								}
							l886:
								{
									position888, tokenIndex888 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l889
									}
									position++
									goto l888
								l889:
									position, tokenIndex = position888, tokenIndex888
									if buffer[position] != rune('T') {
										goto l853
									}
									position++
								}
							l888:
							}
						l865:
							add(rulePegText, position864)
						}
						{
							add(ruleAction19, position)
						}
						add(ruleAccsAttr, position854)
					}
					goto l91
				l853:
					position, tokenIndex = position91, tokenIndex91
					{
						position892 := position
						{
							position893 := position
							if buffer[position] != rune('l') {
								goto l891
							}
							position++
							if buffer[position] != rune('a') {
								goto l891
							}
							position++
							if buffer[position] != rune('n') {
								goto l891
							}
							position++
							if buffer[position] != rune('g') {
								goto l891
							}
							position++
							add(rulePegText, position893)
						}
						if buffer[position] != rune(':') {
							goto l891
						}
						position++
						if !_rules[rulews]() {
							goto l891
						}
						{
							position894, tokenIndex894 := position, tokenIndex
							if buffer[position] != rune('c') {
								goto l895
							}
							position++
							goto l894
						l895:
							position, tokenIndex = position894, tokenIndex894
							if buffer[position] != rune('C') {
								goto l891
							}
							position++
						}
					l894:
						{
							add(ruleAction10, position)
						}
						add(ruleLangAttr, position892)
					}
					goto l91
				l891:
					position, tokenIndex = position91, tokenIndex91
					{
						position898 := position
						{
							position899 := position
							{
								position900, tokenIndex900 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l901
								}
								position++
								goto l900
							l901:
								position, tokenIndex = position900, tokenIndex900
								if buffer[position] != rune('L') {
									goto l897
								}
								position++
							}
						l900:
							{
								position902, tokenIndex902 := position, tokenIndex
								if buffer[position] != rune('i') {
									goto l903
								}
								position++
								goto l902
							l903:
								position, tokenIndex = position902, tokenIndex902
								if buffer[position] != rune('I') {
									goto l897
								}
								position++
							}
						l902:
							{
								position904, tokenIndex904 := position, tokenIndex
								if buffer[position] != rune('n') {
									goto l905
								}
								position++
								goto l904
							l905:
								position, tokenIndex = position904, tokenIndex904
								if buffer[position] != rune('N') {
									goto l897
								}
								position++
							}
						l904:
							{
								position906, tokenIndex906 := position, tokenIndex
								if buffer[position] != rune('k') {
									goto l907
								}
								position++
								goto l906
							l907:
								position, tokenIndex = position906, tokenIndex906
								if buffer[position] != rune('K') {
									goto l897
								}
								position++
							}
						l906:
							add(rulePegText, position899)
						}
						if !_rules[rulews]() {
							goto l897
						}
						if buffer[position] != rune(':') {
							goto l897
						}
						position++
						if !_rules[rulews]() {
							goto l897
						}
						{
							position908 := position
							{
								position909, tokenIndex909 := position, tokenIndex
								{
									position911, tokenIndex911 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l912
									}
									position++
									goto l911
								l912:
									position, tokenIndex = position911, tokenIndex911
									if buffer[position] != rune('E') {
										goto l910
									}
									position++
								}
							l911:
								{
									position913, tokenIndex913 := position, tokenIndex
									if buffer[position] != rune('x') {
										goto l914
									}
									position++
									goto l913
								l914:
									position, tokenIndex = position913, tokenIndex913
									if buffer[position] != rune('X') {
										goto l910
									}
									position++
								}
							l913:
								{
									position915, tokenIndex915 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l916
									}
									position++
									goto l915
								l916:
									position, tokenIndex = position915, tokenIndex915
									if buffer[position] != rune('T') {
										goto l910
									}
									position++
								}
							l915:
								{
									position917, tokenIndex917 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l918
									}
									position++
									goto l917
								l918:
									position, tokenIndex = position917, tokenIndex917
									if buffer[position] != rune('E') {
										goto l910
									}
									position++
								}
							l917:
								{
									position919, tokenIndex919 := position, tokenIndex
									if buffer[position] != rune('r') {
										goto l920
									}
									position++
									goto l919
								l920:
									position, tokenIndex = position919, tokenIndex919
									if buffer[position] != rune('R') {
										goto l910
									}
									position++
								}
							l919:
								{
									position921, tokenIndex921 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l922
									}
									position++
									goto l921
								l922:
									position, tokenIndex = position921, tokenIndex921
									if buffer[position] != rune('N') {
										goto l910
									}
									position++
								}
							l921:
								goto l909
							l910:
								position, tokenIndex = position909, tokenIndex909
								{
									position923, tokenIndex923 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l924
									}
									position++
									goto l923
								l924:
									position, tokenIndex = position923, tokenIndex923
									if buffer[position] != rune('S') {
										goto l897
									}
									position++
								}
							l923:
								{
									position925, tokenIndex925 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l926
									}
									position++
									goto l925
								l926:
									position, tokenIndex = position925, tokenIndex925
									if buffer[position] != rune('T') {
										goto l897
									}
									position++
								}
							l925:
								{
									position927, tokenIndex927 := position, tokenIndex
									if buffer[position] != rune('a') {
										goto l928
									}
									position++
									goto l927
								l928:
									position, tokenIndex = position927, tokenIndex927
									if buffer[position] != rune('A') {
										goto l897
									}
									position++
								}
							l927:
								{
									position929, tokenIndex929 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l930
									}
									position++
									goto l929
								l930:
									position, tokenIndex = position929, tokenIndex929
									if buffer[position] != rune('T') {
										goto l897
									}
									position++
								}
							l929:
								{
									position931, tokenIndex931 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l932
									}
									position++
									goto l931
								l932:
									position, tokenIndex = position931, tokenIndex931
									if buffer[position] != rune('I') {
										goto l897
									}
									position++
								}
							l931:
								{
									position933, tokenIndex933 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l934
									}
									position++
									goto l933
								l934:
									position, tokenIndex = position933, tokenIndex933
									if buffer[position] != rune('C') {
										goto l897
									}
									position++
								}
							l933:
							}
						l909:
							add(rulePegText, position908)
						}
						{
							add(ruleAction17, position)
						}
						add(ruleLinkAttr, position898)
					}
					goto l91
				l897:
					position, tokenIndex = position91, tokenIndex91
					{
						position937 := position
						{
							position938 := position
							{
								position939, tokenIndex939 := position, tokenIndex
								if buffer[position] != rune('a') {
									goto l940
								}
								position++
								goto l939
							l940:
								position, tokenIndex = position939, tokenIndex939
								if buffer[position] != rune('A') {
									goto l936
								}
								position++
							}
						l939:
							{
								position941, tokenIndex941 := position, tokenIndex
								if buffer[position] != rune('l') {
									goto l942
								}
								position++
								goto l941
							l942:
								position, tokenIndex = position941, tokenIndex941
								if buffer[position] != rune('L') {
									goto l936
								}
								position++
							}
						l941:
							{
								position943, tokenIndex943 := position, tokenIndex
								if buffer[position] != rune('g') {
									goto l944
								}
								position++
								goto l943
							l944:
								position, tokenIndex = position943, tokenIndex943
								if buffer[position] != rune('G') {
									goto l936
								}
								position++
							}
						l943:
							{
								position945, tokenIndex945 := position, tokenIndex
								if buffer[position] != rune('n') {
									goto l946
								}
								position++
								goto l945
							l946:
								position, tokenIndex = position945, tokenIndex945
								if buffer[position] != rune('N') {
									goto l936
								}
								position++
							}
						l945:
							add(rulePegText, position938)
						}
						if buffer[position] != rune(':') {
							goto l936
						}
						position++
						if !_rules[rulews]() {
							goto l936
						}
						{
							position947 := position
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l936
							}
							position++
						l948:
							{
								position949, tokenIndex949 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l949
								}
								position++
								goto l948
							l949:
								position, tokenIndex = position949, tokenIndex949
							}
							add(rulePegText, position947)
						}
						if !_rules[rulews]() {
							goto l936
						}
						{
							add(ruleAction39, position)
						}
						add(ruleIntAttr2, position937)
					}
					goto l91
				l936:
					position, tokenIndex = position91, tokenIndex91
					{
						switch buffer[position] {
						case 'A', 'a':
							{
								position952 := position
								{
									position953 := position
									{
										position954, tokenIndex954 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l955
										}
										position++
										goto l954
									l955:
										position, tokenIndex = position954, tokenIndex954
										if buffer[position] != rune('A') {
											goto l89
										}
										position++
									}
								l954:
									{
										position956, tokenIndex956 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l957
										}
										position++
										goto l956
									l957:
										position, tokenIndex = position956, tokenIndex956
										if buffer[position] != rune('D') {
											goto l89
										}
										position++
									}
								l956:
									{
										position958, tokenIndex958 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l959
										}
										position++
										goto l958
									l959:
										position, tokenIndex = position958, tokenIndex958
										if buffer[position] != rune('D') {
											goto l89
										}
										position++
									}
								l958:
									{
										position960, tokenIndex960 := position, tokenIndex
										if buffer[position] != rune('r') {
											goto l961
										}
										position++
										goto l960
									l961:
										position, tokenIndex = position960, tokenIndex960
										if buffer[position] != rune('R') {
											goto l89
										}
										position++
									}
								l960:
									if buffer[position] != rune(':') {
										goto l89
									}
									position++
									add(rulePegText, position953)
								}
								if !_rules[rulews]() {
									goto l89
								}
								{
									position962 := position
									{
										position963 := position
										{
											position966, tokenIndex966 := position, tokenIndex
											if !_rules[ruleDecimalDigit]() {
												goto l967
											}
											goto l966
										l967:
											position, tokenIndex = position966, tokenIndex966
											if !_rules[ruleHex]() {
												goto l89
											}
										}
									l966:
									l964:
										{
											position965, tokenIndex965 := position, tokenIndex
											{
												position968, tokenIndex968 := position, tokenIndex
												if !_rules[ruleDecimalDigit]() {
													goto l969
												}
												goto l968
											l969:
												position, tokenIndex = position968, tokenIndex968
												if !_rules[ruleHex]() {
													goto l965
												}
											}
										l968:
											goto l964
										l965:
											position, tokenIndex = position965, tokenIndex965
										}
										add(ruleAddr, position963)
									}
									add(rulePegText, position962)
								}
								{
									add(ruleAction14, position)
								}
								add(ruleAddrAttr, position952)
							}
							break
						case 'L', 'l':
							{
								position971 := position
								{
									position972, tokenIndex972 := position, tokenIndex
									if buffer[position] != rune('l') {
										goto l973
									}
									position++
									goto l972
								l973:
									position, tokenIndex = position972, tokenIndex972
									if buffer[position] != rune('L') {
										goto l89
									}
									position++
								}
							l972:
								{
									position974, tokenIndex974 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l975
									}
									position++
									goto l974
								l975:
									position, tokenIndex = position974, tokenIndex974
									if buffer[position] != rune('N') {
										goto l89
									}
									position++
								}
							l974:
								{
									position976, tokenIndex976 := position, tokenIndex
									if buffer[position] != rune('g') {
										goto l977
									}
									position++
									goto l976
								l977:
									position, tokenIndex = position976, tokenIndex976
									if buffer[position] != rune('G') {
										goto l89
									}
									position++
								}
							l976:
								{
									position978, tokenIndex978 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l979
									}
									position++
									goto l978
								l979:
									position, tokenIndex = position978, tokenIndex978
									if buffer[position] != rune('T') {
										goto l89
									}
									position++
								}
							l978:
								if !_rules[rulews]() {
									goto l89
								}
								if buffer[position] != rune(':') {
									goto l89
								}
								position++
								if !_rules[rulews]() {
									goto l89
								}
								{
									position980 := position
									if !_rules[ruleInteger]() {
										goto l89
									}
									add(rulePegText, position980)
								}
								if !_rules[rulews]() {
									goto l89
								}
								add(ruleLngtAttr, position971)
							}
							break
						case 'I', 'i':
							{
								position981 := position
								{
									position982, tokenIndex982 := position, tokenIndex
									if buffer[position] != rune('i') {
										goto l983
									}
									position++
									goto l982
								l983:
									position, tokenIndex = position982, tokenIndex982
									if buffer[position] != rune('I') {
										goto l89
									}
									position++
								}
							l982:
								{
									position984, tokenIndex984 := position, tokenIndex
									if buffer[position] != rune('n') {
										goto l985
									}
									position++
									goto l984
								l985:
									position, tokenIndex = position984, tokenIndex984
									if buffer[position] != rune('N') {
										goto l89
									}
									position++
								}
							l984:
								{
									position986, tokenIndex986 := position, tokenIndex
									if buffer[position] != rune('t') {
										goto l987
									}
									position++
									goto l986
								l987:
									position, tokenIndex = position986, tokenIndex986
									if buffer[position] != rune('T') {
										goto l89
									}
									position++
								}
							l986:
								if buffer[position] != rune(':') {
									goto l89
								}
								position++
								if !_rules[rulews]() {
									goto l89
								}
								{
									position988, tokenIndex988 := position, tokenIndex
									if !_rules[ruleHexValue]() {
										goto l989
									}
									goto l988
								l989:
									position, tokenIndex = position988, tokenIndex988
									if !_rules[ruleNegInt]() {
										goto l990
									}
									goto l988
								l990:
									position, tokenIndex = position988, tokenIndex988
									if !_rules[ruleInteger]() {
										goto l89
									}
								}
							l988:
								{
									add(ruleAction38, position)
								}
								add(ruleSignedIntAttr, position981)
							}
							break
						case 'Q', 'q':
							{
								position992 := position
								{
									position993 := position
									{
										position994, tokenIndex994 := position, tokenIndex
										if buffer[position] != rune('q') {
											goto l995
										}
										position++
										goto l994
									l995:
										position, tokenIndex = position994, tokenIndex994
										if buffer[position] != rune('Q') {
											goto l89
										}
										position++
									}
								l994:
									{
										position996, tokenIndex996 := position, tokenIndex
										if buffer[position] != rune('u') {
											goto l997
										}
										position++
										goto l996
									l997:
										position, tokenIndex = position996, tokenIndex996
										if buffer[position] != rune('U') {
											goto l89
										}
										position++
									}
								l996:
									{
										position998, tokenIndex998 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l999
										}
										position++
										goto l998
									l999:
										position, tokenIndex = position998, tokenIndex998
										if buffer[position] != rune('A') {
											goto l89
										}
										position++
									}
								l998:
									{
										position1000, tokenIndex1000 := position, tokenIndex
										if buffer[position] != rune('l') {
											goto l1001
										}
										position++
										goto l1000
									l1001:
										position, tokenIndex = position1000, tokenIndex1000
										if buffer[position] != rune('L') {
											goto l89
										}
										position++
									}
								l1000:
									add(rulePegText, position993)
								}
								if !_rules[rulews]() {
									goto l89
								}
								if buffer[position] != rune(':') {
									goto l89
								}
								position++
								if !_rules[rulews]() {
									goto l89
								}
								{
									position1004 := position
									{
										switch buffer[position] {
										case 'R', 'r':
											{
												position1006, tokenIndex1006 := position, tokenIndex
												if buffer[position] != rune('r') {
													goto l1007
												}
												position++
												goto l1006
											l1007:
												position, tokenIndex = position1006, tokenIndex1006
												if buffer[position] != rune('R') {
													goto l89
												}
												position++
											}
										l1006:
											break
										case 'C':
											if buffer[position] != rune('C') {
												goto l89
											}
											position++
											break
										case 'c':
											if buffer[position] != rune('c') {
												goto l89
											}
											position++
											break
										default:
											{
												position1008, tokenIndex1008 := position, tokenIndex
												if buffer[position] != rune('v') {
													goto l1009
												}
												position++
												goto l1008
											l1009:
												position, tokenIndex = position1008, tokenIndex1008
												if buffer[position] != rune('V') {
													goto l89
												}
												position++
											}
										l1008:
											break
										}
									}

									add(rulePegText, position1004)
								}
							l1002:
								{
									position1003, tokenIndex1003 := position, tokenIndex
									{
										position1010 := position
										{
											switch buffer[position] {
											case 'R', 'r':
												{
													position1012, tokenIndex1012 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l1013
													}
													position++
													goto l1012
												l1013:
													position, tokenIndex = position1012, tokenIndex1012
													if buffer[position] != rune('R') {
														goto l1003
													}
													position++
												}
											l1012:
												break
											case 'C':
												if buffer[position] != rune('C') {
													goto l1003
												}
												position++
												break
											case 'c':
												if buffer[position] != rune('c') {
													goto l1003
												}
												position++
												break
											default:
												{
													position1014, tokenIndex1014 := position, tokenIndex
													if buffer[position] != rune('v') {
														goto l1015
													}
													position++
													goto l1014
												l1015:
													position, tokenIndex = position1014, tokenIndex1014
													if buffer[position] != rune('V') {
														goto l1003
													}
													position++
												}
											l1014:
												break
											}
										}

										add(rulePegText, position1010)
									}
									goto l1002
								l1003:
									position, tokenIndex = position1003, tokenIndex1003
								}
								if !_rules[rulews]() {
									goto l89
								}
								{
									add(ruleAction20, position)
								}
								add(ruleQualAttr, position992)
							}
							break
						case 'N', 'n':
							{
								position1017 := position
								{
									position1018 := position
									{
										position1019, tokenIndex1019 := position, tokenIndex
										if buffer[position] != rune('n') {
											goto l1020
										}
										position++
										goto l1019
									l1020:
										position, tokenIndex = position1019, tokenIndex1019
										if buffer[position] != rune('N') {
											goto l89
										}
										position++
									}
								l1019:
									{
										position1021, tokenIndex1021 := position, tokenIndex
										if buffer[position] != rune('o') {
											goto l1022
										}
										position++
										goto l1021
									l1022:
										position, tokenIndex = position1021, tokenIndex1021
										if buffer[position] != rune('O') {
											goto l89
										}
										position++
									}
								l1021:
									{
										position1023, tokenIndex1023 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l1024
										}
										position++
										goto l1023
									l1024:
										position, tokenIndex = position1023, tokenIndex1023
										if buffer[position] != rune('T') {
											goto l89
										}
										position++
									}
								l1023:
									{
										position1025, tokenIndex1025 := position, tokenIndex
										if buffer[position] != rune('e') {
											goto l1026
										}
										position++
										goto l1025
									l1026:
										position, tokenIndex = position1025, tokenIndex1025
										if buffer[position] != rune('E') {
											goto l89
										}
										position++
									}
								l1025:
									add(rulePegText, position1018)
								}
								if !_rules[rulews]() {
									goto l89
								}
								if buffer[position] != rune(':') {
									goto l89
								}
								position++
								if !_rules[rulews]() {
									goto l89
								}
								{
									position1027 := position
									{
										position1028, tokenIndex1028 := position, tokenIndex
										{
											position1030, tokenIndex1030 := position, tokenIndex
											if buffer[position] != rune('p') {
												goto l1031
											}
											position++
											goto l1030
										l1031:
											position, tokenIndex = position1030, tokenIndex1030
											if buffer[position] != rune('P') {
												goto l1029
											}
											position++
										}
									l1030:
										{
											position1032, tokenIndex1032 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l1033
											}
											position++
											goto l1032
										l1033:
											position, tokenIndex = position1032, tokenIndex1032
											if buffer[position] != rune('T') {
												goto l1029
											}
											position++
										}
									l1032:
										{
											position1034, tokenIndex1034 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l1035
											}
											position++
											goto l1034
										l1035:
											position, tokenIndex = position1034, tokenIndex1034
											if buffer[position] != rune('R') {
												goto l1029
											}
											position++
										}
									l1034:
										{
											position1036, tokenIndex1036 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l1037
											}
											position++
											goto l1036
										l1037:
											position, tokenIndex = position1036, tokenIndex1036
											if buffer[position] != rune('M') {
												goto l1029
											}
											position++
										}
									l1036:
										{
											position1038, tokenIndex1038 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l1039
											}
											position++
											goto l1038
										l1039:
											position, tokenIndex = position1038, tokenIndex1038
											if buffer[position] != rune('E') {
												goto l1029
											}
											position++
										}
									l1038:
										{
											position1040, tokenIndex1040 := position, tokenIndex
											if buffer[position] != rune('m') {
												goto l1041
											}
											position++
											goto l1040
										l1041:
											position, tokenIndex = position1040, tokenIndex1040
											if buffer[position] != rune('M') {
												goto l1029
											}
											position++
										}
									l1040:
										goto l1028
									l1029:
										position, tokenIndex = position1028, tokenIndex1028
										{
											position1043, tokenIndex1043 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l1044
											}
											position++
											goto l1043
										l1044:
											position, tokenIndex = position1043, tokenIndex1043
											if buffer[position] != rune('C') {
												goto l1042
											}
											position++
										}
									l1043:
										{
											position1045, tokenIndex1045 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l1046
											}
											position++
											goto l1045
										l1046:
											position, tokenIndex = position1045, tokenIndex1045
											if buffer[position] != rune('O') {
												goto l1042
											}
											position++
										}
									l1045:
										{
											position1047, tokenIndex1047 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l1048
											}
											position++
											goto l1047
										l1048:
											position, tokenIndex = position1047, tokenIndex1047
											if buffer[position] != rune('N') {
												goto l1042
											}
											position++
										}
									l1047:
										{
											position1049, tokenIndex1049 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l1050
											}
											position++
											goto l1049
										l1050:
											position, tokenIndex = position1049, tokenIndex1049
											if buffer[position] != rune('S') {
												goto l1042
											}
											position++
										}
									l1049:
										{
											position1051, tokenIndex1051 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l1052
											}
											position++
											goto l1051
										l1052:
											position, tokenIndex = position1051, tokenIndex1051
											if buffer[position] != rune('T') {
												goto l1042
											}
											position++
										}
									l1051:
										{
											position1053, tokenIndex1053 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l1054
											}
											position++
											goto l1053
										l1054:
											position, tokenIndex = position1053, tokenIndex1053
											if buffer[position] != rune('R') {
												goto l1042
											}
											position++
										}
									l1053:
										{
											position1055, tokenIndex1055 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l1056
											}
											position++
											goto l1055
										l1056:
											position, tokenIndex = position1055, tokenIndex1055
											if buffer[position] != rune('U') {
												goto l1042
											}
											position++
										}
									l1055:
										{
											position1057, tokenIndex1057 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l1058
											}
											position++
											goto l1057
										l1058:
											position, tokenIndex = position1057, tokenIndex1057
											if buffer[position] != rune('C') {
												goto l1042
											}
											position++
										}
									l1057:
										{
											position1059, tokenIndex1059 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l1060
											}
											position++
											goto l1059
										l1060:
											position, tokenIndex = position1059, tokenIndex1059
											if buffer[position] != rune('T') {
												goto l1042
											}
											position++
										}
									l1059:
										{
											position1061, tokenIndex1061 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l1062
											}
											position++
											goto l1061
										l1062:
											position, tokenIndex = position1061, tokenIndex1061
											if buffer[position] != rune('O') {
												goto l1042
											}
											position++
										}
									l1061:
										{
											position1063, tokenIndex1063 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l1064
											}
											position++
											goto l1063
										l1064:
											position, tokenIndex = position1063, tokenIndex1063
											if buffer[position] != rune('R') {
												goto l1042
											}
											position++
										}
									l1063:
										goto l1028
									l1042:
										position, tokenIndex = position1028, tokenIndex1028
										{
											switch buffer[position] {
											case 'D', 'd':
												{
													position1066, tokenIndex1066 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l1067
													}
													position++
													goto l1066
												l1067:
													position, tokenIndex = position1066, tokenIndex1066
													if buffer[position] != rune('D') {
														goto l89
													}
													position++
												}
											l1066:
												{
													position1068, tokenIndex1068 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l1069
													}
													position++
													goto l1068
												l1069:
													position, tokenIndex = position1068, tokenIndex1068
													if buffer[position] != rune('E') {
														goto l89
													}
													position++
												}
											l1068:
												{
													position1070, tokenIndex1070 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l1071
													}
													position++
													goto l1070
												l1071:
													position, tokenIndex = position1070, tokenIndex1070
													if buffer[position] != rune('S') {
														goto l89
													}
													position++
												}
											l1070:
												{
													position1072, tokenIndex1072 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1073
													}
													position++
													goto l1072
												l1073:
													position, tokenIndex = position1072, tokenIndex1072
													if buffer[position] != rune('T') {
														goto l89
													}
													position++
												}
											l1072:
												{
													position1074, tokenIndex1074 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l1075
													}
													position++
													goto l1074
												l1075:
													position, tokenIndex = position1074, tokenIndex1074
													if buffer[position] != rune('R') {
														goto l89
													}
													position++
												}
											l1074:
												{
													position1076, tokenIndex1076 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1077
													}
													position++
													goto l1076
												l1077:
													position, tokenIndex = position1076, tokenIndex1076
													if buffer[position] != rune('U') {
														goto l89
													}
													position++
												}
											l1076:
												{
													position1078, tokenIndex1078 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l1079
													}
													position++
													goto l1078
												l1079:
													position, tokenIndex = position1078, tokenIndex1078
													if buffer[position] != rune('C') {
														goto l89
													}
													position++
												}
											l1078:
												{
													position1080, tokenIndex1080 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1081
													}
													position++
													goto l1080
												l1081:
													position, tokenIndex = position1080, tokenIndex1080
													if buffer[position] != rune('T') {
														goto l89
													}
													position++
												}
											l1080:
												{
													position1082, tokenIndex1082 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1083
													}
													position++
													goto l1082
												l1083:
													position, tokenIndex = position1082, tokenIndex1082
													if buffer[position] != rune('O') {
														goto l89
													}
													position++
												}
											l1082:
												{
													position1084, tokenIndex1084 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l1085
													}
													position++
													goto l1084
												l1085:
													position, tokenIndex = position1084, tokenIndex1084
													if buffer[position] != rune('R') {
														goto l89
													}
													position++
												}
											l1084:
												break
											case 'P', 'p':
												{
													position1086, tokenIndex1086 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l1087
													}
													position++
													goto l1086
												l1087:
													position, tokenIndex = position1086, tokenIndex1086
													if buffer[position] != rune('P') {
														goto l89
													}
													position++
												}
											l1086:
												{
													position1088, tokenIndex1088 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l1089
													}
													position++
													goto l1088
												l1089:
													position, tokenIndex = position1088, tokenIndex1088
													if buffer[position] != rune('S') {
														goto l89
													}
													position++
												}
											l1088:
												{
													position1090, tokenIndex1090 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l1091
													}
													position++
													goto l1090
												l1091:
													position, tokenIndex = position1090, tokenIndex1090
													if buffer[position] != rune('E') {
														goto l89
													}
													position++
												}
											l1090:
												{
													position1092, tokenIndex1092 := position, tokenIndex
													if buffer[position] != rune('u') {
														goto l1093
													}
													position++
													goto l1092
												l1093:
													position, tokenIndex = position1092, tokenIndex1092
													if buffer[position] != rune('U') {
														goto l89
													}
													position++
												}
											l1092:
												{
													position1094, tokenIndex1094 := position, tokenIndex
													if buffer[position] != rune('d') {
														goto l1095
													}
													position++
													goto l1094
												l1095:
													position, tokenIndex = position1094, tokenIndex1094
													if buffer[position] != rune('D') {
														goto l89
													}
													position++
												}
											l1094:
												{
													position1096, tokenIndex1096 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1097
													}
													position++
													goto l1096
												l1097:
													position, tokenIndex = position1096, tokenIndex1096
													if buffer[position] != rune('O') {
														goto l89
													}
													position++
												}
											l1096:
												if buffer[position] != rune(' ') {
													goto l89
												}
												position++
												{
													position1098, tokenIndex1098 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1099
													}
													position++
													goto l1098
												l1099:
													position, tokenIndex = position1098, tokenIndex1098
													if buffer[position] != rune('T') {
														goto l89
													}
													position++
												}
											l1098:
												{
													position1100, tokenIndex1100 := position, tokenIndex
													if buffer[position] != rune('m') {
														goto l1101
													}
													position++
													goto l1100
												l1101:
													position, tokenIndex = position1100, tokenIndex1100
													if buffer[position] != rune('M') {
														goto l89
													}
													position++
												}
											l1100:
												{
													position1102, tokenIndex1102 := position, tokenIndex
													if buffer[position] != rune('p') {
														goto l1103
													}
													position++
													goto l1102
												l1103:
													position, tokenIndex = position1102, tokenIndex1102
													if buffer[position] != rune('P') {
														goto l89
													}
													position++
												}
											l1102:
												{
													position1104, tokenIndex1104 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l1105
													}
													position++
													goto l1104
												l1105:
													position, tokenIndex = position1104, tokenIndex1104
													if buffer[position] != rune('L') {
														goto l89
													}
													position++
												}
											l1104:
												break
											case 'C', 'c':
												{
													position1106, tokenIndex1106 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l1107
													}
													position++
													goto l1106
												l1107:
													position, tokenIndex = position1106, tokenIndex1106
													if buffer[position] != rune('C') {
														goto l89
													}
													position++
												}
											l1106:
												{
													position1108, tokenIndex1108 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1109
													}
													position++
													goto l1108
												l1109:
													position, tokenIndex = position1108, tokenIndex1108
													if buffer[position] != rune('O') {
														goto l89
													}
													position++
												}
											l1108:
												{
													position1110, tokenIndex1110 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l1111
													}
													position++
													goto l1110
												l1111:
													position, tokenIndex = position1110, tokenIndex1110
													if buffer[position] != rune('N') {
														goto l89
													}
													position++
												}
											l1110:
												{
													position1112, tokenIndex1112 := position, tokenIndex
													if buffer[position] != rune('v') {
														goto l1113
													}
													position++
													goto l1112
												l1113:
													position, tokenIndex = position1112, tokenIndex1112
													if buffer[position] != rune('V') {
														goto l89
													}
													position++
												}
											l1112:
												{
													position1114, tokenIndex1114 := position, tokenIndex
													if buffer[position] != rune('e') {
														goto l1115
													}
													position++
													goto l1114
												l1115:
													position, tokenIndex = position1114, tokenIndex1114
													if buffer[position] != rune('E') {
														goto l89
													}
													position++
												}
											l1114:
												{
													position1116, tokenIndex1116 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l1117
													}
													position++
													goto l1116
												l1117:
													position, tokenIndex = position1116, tokenIndex1116
													if buffer[position] != rune('R') {
														goto l89
													}
													position++
												}
											l1116:
												{
													position1118, tokenIndex1118 := position, tokenIndex
													if buffer[position] != rune('s') {
														goto l1119
													}
													position++
													goto l1118
												l1119:
													position, tokenIndex = position1118, tokenIndex1118
													if buffer[position] != rune('S') {
														goto l89
													}
													position++
												}
											l1118:
												{
													position1120, tokenIndex1120 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1121
													}
													position++
													goto l1120
												l1121:
													position, tokenIndex = position1120, tokenIndex1120
													if buffer[position] != rune('I') {
														goto l89
													}
													position++
												}
											l1120:
												{
													position1122, tokenIndex1122 := position, tokenIndex
													if buffer[position] != rune('o') {
														goto l1123
													}
													position++
													goto l1122
												l1123:
													position, tokenIndex = position1122, tokenIndex1122
													if buffer[position] != rune('O') {
														goto l89
													}
													position++
												}
											l1122:
												{
													position1124, tokenIndex1124 := position, tokenIndex
													if buffer[position] != rune('n') {
														goto l1125
													}
													position++
													goto l1124
												l1125:
													position, tokenIndex = position1124, tokenIndex1124
													if buffer[position] != rune('N') {
														goto l89
													}
													position++
												}
											l1124:
												break
											default:
												{
													position1126, tokenIndex1126 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l1127
													}
													position++
													goto l1126
												l1127:
													position, tokenIndex = position1126, tokenIndex1126
													if buffer[position] != rune('A') {
														goto l89
													}
													position++
												}
											l1126:
												{
													position1128, tokenIndex1128 := position, tokenIndex
													if buffer[position] != rune('r') {
														goto l1129
													}
													position++
													goto l1128
												l1129:
													position, tokenIndex = position1128, tokenIndex1128
													if buffer[position] != rune('R') {
														goto l89
													}
													position++
												}
											l1128:
												{
													position1130, tokenIndex1130 := position, tokenIndex
													if buffer[position] != rune('t') {
														goto l1131
													}
													position++
													goto l1130
												l1131:
													position, tokenIndex = position1130, tokenIndex1130
													if buffer[position] != rune('T') {
														goto l89
													}
													position++
												}
											l1130:
												{
													position1132, tokenIndex1132 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1133
													}
													position++
													goto l1132
												l1133:
													position, tokenIndex = position1132, tokenIndex1132
													if buffer[position] != rune('I') {
														goto l89
													}
													position++
												}
											l1132:
												{
													position1134, tokenIndex1134 := position, tokenIndex
													if buffer[position] != rune('f') {
														goto l1135
													}
													position++
													goto l1134
												l1135:
													position, tokenIndex = position1134, tokenIndex1134
													if buffer[position] != rune('F') {
														goto l89
													}
													position++
												}
											l1134:
												{
													position1136, tokenIndex1136 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1137
													}
													position++
													goto l1136
												l1137:
													position, tokenIndex = position1136, tokenIndex1136
													if buffer[position] != rune('I') {
														goto l89
													}
													position++
												}
											l1136:
												{
													position1138, tokenIndex1138 := position, tokenIndex
													if buffer[position] != rune('c') {
														goto l1139
													}
													position++
													goto l1138
												l1139:
													position, tokenIndex = position1138, tokenIndex1138
													if buffer[position] != rune('C') {
														goto l89
													}
													position++
												}
											l1138:
												{
													position1140, tokenIndex1140 := position, tokenIndex
													if buffer[position] != rune('i') {
														goto l1141
													}
													position++
													goto l1140
												l1141:
													position, tokenIndex = position1140, tokenIndex1140
													if buffer[position] != rune('I') {
														goto l89
													}
													position++
												}
											l1140:
												{
													position1142, tokenIndex1142 := position, tokenIndex
													if buffer[position] != rune('a') {
														goto l1143
													}
													position++
													goto l1142
												l1143:
													position, tokenIndex = position1142, tokenIndex1142
													if buffer[position] != rune('A') {
														goto l89
													}
													position++
												}
											l1142:
												{
													position1144, tokenIndex1144 := position, tokenIndex
													if buffer[position] != rune('l') {
														goto l1145
													}
													position++
													goto l1144
												l1145:
													position, tokenIndex = position1144, tokenIndex1144
													if buffer[position] != rune('L') {
														goto l89
													}
													position++
												}
											l1144:
												break
											}
										}

									}
								l1028:
									add(rulePegText, position1027)
								}
								if !_rules[rulews]() {
									goto l89
								}
								{
									add(ruleAction18, position)
								}
								add(ruleNoteAttr, position1017)
							}
							break
						case 'B', 'b':
							{
								position1147 := position
								{
									position1148 := position
									{
										position1149, tokenIndex1149 := position, tokenIndex
										if buffer[position] != rune('b') {
											goto l1150
										}
										position++
										goto l1149
									l1150:
										position, tokenIndex = position1149, tokenIndex1149
										if buffer[position] != rune('B') {
											goto l89
										}
										position++
									}
								l1149:
									{
										position1151, tokenIndex1151 := position, tokenIndex
										if buffer[position] != rune('o') {
											goto l1152
										}
										position++
										goto l1151
									l1152:
										position, tokenIndex = position1151, tokenIndex1151
										if buffer[position] != rune('O') {
											goto l89
										}
										position++
									}
								l1151:
									{
										position1153, tokenIndex1153 := position, tokenIndex
										if buffer[position] != rune('d') {
											goto l1154
										}
										position++
										goto l1153
									l1154:
										position, tokenIndex = position1153, tokenIndex1153
										if buffer[position] != rune('D') {
											goto l89
										}
										position++
									}
								l1153:
									{
										position1155, tokenIndex1155 := position, tokenIndex
										if buffer[position] != rune('y') {
											goto l1156
										}
										position++
										goto l1155
									l1156:
										position, tokenIndex = position1155, tokenIndex1155
										if buffer[position] != rune('Y') {
											goto l89
										}
										position++
									}
								l1155:
									add(rulePegText, position1148)
								}
								if !_rules[rulews]() {
									goto l89
								}
								if buffer[position] != rune(':') {
									goto l89
								}
								position++
								if !_rules[rulews]() {
									goto l89
								}
								{
									position1157 := position
									{
										position1158, tokenIndex1158 := position, tokenIndex
										{
											position1160, tokenIndex1160 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l1161
											}
											position++
											goto l1160
										l1161:
											position, tokenIndex = position1160, tokenIndex1160
											if buffer[position] != rune('U') {
												goto l1159
											}
											position++
										}
									l1160:
										{
											position1162, tokenIndex1162 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l1163
											}
											position++
											goto l1162
										l1163:
											position, tokenIndex = position1162, tokenIndex1162
											if buffer[position] != rune('N') {
												goto l1159
											}
											position++
										}
									l1162:
										{
											position1164, tokenIndex1164 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l1165
											}
											position++
											goto l1164
										l1165:
											position, tokenIndex = position1164, tokenIndex1164
											if buffer[position] != rune('D') {
												goto l1159
											}
											position++
										}
									l1164:
										{
											position1166, tokenIndex1166 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l1167
											}
											position++
											goto l1166
										l1167:
											position, tokenIndex = position1166, tokenIndex1166
											if buffer[position] != rune('E') {
												goto l1159
											}
											position++
										}
									l1166:
										{
											position1168, tokenIndex1168 := position, tokenIndex
											if buffer[position] != rune('f') {
												goto l1169
											}
											position++
											goto l1168
										l1169:
											position, tokenIndex = position1168, tokenIndex1168
											if buffer[position] != rune('F') {
												goto l1159
											}
											position++
										}
									l1168:
										{
											position1170, tokenIndex1170 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l1171
											}
											position++
											goto l1170
										l1171:
											position, tokenIndex = position1170, tokenIndex1170
											if buffer[position] != rune('I') {
												goto l1159
											}
											position++
										}
									l1170:
										{
											position1172, tokenIndex1172 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l1173
											}
											position++
											goto l1172
										l1173:
											position, tokenIndex = position1172, tokenIndex1172
											if buffer[position] != rune('N') {
												goto l1159
											}
											position++
										}
									l1172:
										{
											position1174, tokenIndex1174 := position, tokenIndex
											if buffer[position] != rune('e') {
												goto l1175
											}
											position++
											goto l1174
										l1175:
											position, tokenIndex = position1174, tokenIndex1174
											if buffer[position] != rune('E') {
												goto l1159
											}
											position++
										}
									l1174:
										{
											position1176, tokenIndex1176 := position, tokenIndex
											if buffer[position] != rune('d') {
												goto l1177
											}
											position++
											goto l1176
										l1177:
											position, tokenIndex = position1176, tokenIndex1176
											if buffer[position] != rune('D') {
												goto l1159
											}
											position++
										}
									l1176:
										goto l1158
									l1159:
										position, tokenIndex = position1158, tokenIndex1158
										if !_rules[ruleNode]() {
											goto l89
										}
									}
								l1158:
									add(rulePegText, position1157)
								}
								{
									add(ruleAction16, position)
								}
								add(ruleBodyAttr, position1147)
							}
							break
						case 'T', 't':
							{
								position1179 := position
								{
									position1180 := position
									{
										position1181, tokenIndex1181 := position, tokenIndex
										if buffer[position] != rune('t') {
											goto l1182
										}
										position++
										goto l1181
									l1182:
										position, tokenIndex = position1181, tokenIndex1181
										if buffer[position] != rune('T') {
											goto l89
										}
										position++
									}
								l1181:
									{
										position1183, tokenIndex1183 := position, tokenIndex
										if buffer[position] != rune('a') {
											goto l1184
										}
										position++
										goto l1183
									l1184:
										position, tokenIndex = position1183, tokenIndex1183
										if buffer[position] != rune('A') {
											goto l89
										}
										position++
									}
								l1183:
									{
										position1185, tokenIndex1185 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l1186
										}
										position++
										goto l1185
									l1186:
										position, tokenIndex = position1185, tokenIndex1185
										if buffer[position] != rune('G') {
											goto l89
										}
										position++
									}
								l1185:
									add(rulePegText, position1180)
								}
								if !_rules[rulews]() {
									goto l89
								}
								if buffer[position] != rune(':') {
									goto l89
								}
								position++
								if !_rules[rulews]() {
									goto l89
								}
								{
									position1187 := position
									{
										position1188, tokenIndex1188 := position, tokenIndex
										{
											position1190, tokenIndex1190 := position, tokenIndex
											if buffer[position] != rune('s') {
												goto l1191
											}
											position++
											goto l1190
										l1191:
											position, tokenIndex = position1190, tokenIndex1190
											if buffer[position] != rune('S') {
												goto l1189
											}
											position++
										}
									l1190:
										{
											position1192, tokenIndex1192 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l1193
											}
											position++
											goto l1192
										l1193:
											position, tokenIndex = position1192, tokenIndex1192
											if buffer[position] != rune('T') {
												goto l1189
											}
											position++
										}
									l1192:
										{
											position1194, tokenIndex1194 := position, tokenIndex
											if buffer[position] != rune('r') {
												goto l1195
											}
											position++
											goto l1194
										l1195:
											position, tokenIndex = position1194, tokenIndex1194
											if buffer[position] != rune('R') {
												goto l1189
											}
											position++
										}
									l1194:
										{
											position1196, tokenIndex1196 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l1197
											}
											position++
											goto l1196
										l1197:
											position, tokenIndex = position1196, tokenIndex1196
											if buffer[position] != rune('U') {
												goto l1189
											}
											position++
										}
									l1196:
										{
											position1198, tokenIndex1198 := position, tokenIndex
											if buffer[position] != rune('c') {
												goto l1199
											}
											position++
											goto l1198
										l1199:
											position, tokenIndex = position1198, tokenIndex1198
											if buffer[position] != rune('C') {
												goto l1189
											}
											position++
										}
									l1198:
										{
											position1200, tokenIndex1200 := position, tokenIndex
											if buffer[position] != rune('t') {
												goto l1201
											}
											position++
											goto l1200
										l1201:
											position, tokenIndex = position1200, tokenIndex1200
											if buffer[position] != rune('T') {
												goto l1189
											}
											position++
										}
									l1200:
										goto l1188
									l1189:
										position, tokenIndex = position1188, tokenIndex1188
										{
											position1202, tokenIndex1202 := position, tokenIndex
											if buffer[position] != rune('u') {
												goto l1203
											}
											position++
											goto l1202
										l1203:
											position, tokenIndex = position1202, tokenIndex1202
											if buffer[position] != rune('U') {
												goto l89
											}
											position++
										}
									l1202:
										{
											position1204, tokenIndex1204 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l1205
											}
											position++
											goto l1204
										l1205:
											position, tokenIndex = position1204, tokenIndex1204
											if buffer[position] != rune('N') {
												goto l89
											}
											position++
										}
									l1204:
										{
											position1206, tokenIndex1206 := position, tokenIndex
											if buffer[position] != rune('i') {
												goto l1207
											}
											position++
											goto l1206
										l1207:
											position, tokenIndex = position1206, tokenIndex1206
											if buffer[position] != rune('I') {
												goto l89
											}
											position++
										}
									l1206:
										{
											position1208, tokenIndex1208 := position, tokenIndex
											if buffer[position] != rune('o') {
												goto l1209
											}
											position++
											goto l1208
										l1209:
											position, tokenIndex = position1208, tokenIndex1208
											if buffer[position] != rune('O') {
												goto l89
											}
											position++
										}
									l1208:
										{
											position1210, tokenIndex1210 := position, tokenIndex
											if buffer[position] != rune('n') {
												goto l1211
											}
											position++
											goto l1210
										l1211:
											position, tokenIndex = position1210, tokenIndex1210
											if buffer[position] != rune('N') {
												goto l89
											}
											position++
										}
									l1210:
									}
								l1188:
									add(rulePegText, position1187)
								}
								{
									add(ruleAction15, position)
								}
								add(ruleTagAttr, position1179)
							}
							break
						case 'H', 'h':
							{
								position1213 := position
								{
									position1214 := position
									{
										position1215, tokenIndex1215 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1216
										}
										position++
										goto l1215
									l1216:
										position, tokenIndex = position1215, tokenIndex1215
										if buffer[position] != rune('H') {
											goto l89
										}
										position++
									}
								l1215:
									{
										position1217, tokenIndex1217 := position, tokenIndex
										if buffer[position] != rune('i') {
											goto l1218
										}
										position++
										goto l1217
									l1218:
										position, tokenIndex = position1217, tokenIndex1217
										if buffer[position] != rune('I') {
											goto l89
										}
										position++
									}
								l1217:
									{
										position1219, tokenIndex1219 := position, tokenIndex
										if buffer[position] != rune('g') {
											goto l1220
										}
										position++
										goto l1219
									l1220:
										position, tokenIndex = position1219, tokenIndex1219
										if buffer[position] != rune('G') {
											goto l89
										}
										position++
									}
								l1219:
									{
										position1221, tokenIndex1221 := position, tokenIndex
										if buffer[position] != rune('h') {
											goto l1222
										}
										position++
										goto l1221
									l1222:
										position, tokenIndex = position1221, tokenIndex1221
										if buffer[position] != rune('H') {
											goto l89
										}
										position++
									}
								l1221:
									add(rulePegText, position1214)
								}
								if !_rules[rulews]() {
									goto l89
								}
								if buffer[position] != rune(':') {
									goto l89
								}
								position++
								if !_rules[rulews]() {
									goto l89
								}
								{
									position1223 := position
									{
										position1224, tokenIndex1224 := position, tokenIndex
										if buffer[position] != rune('-') {
											goto l1224
										}
										position++
										goto l1225
									l1224:
										position, tokenIndex = position1224, tokenIndex1224
									}
								l1225:
									if !_rules[ruleInteger]() {
										goto l89
									}
									add(rulePegText, position1223)
								}
								if !_rules[rulews]() {
									goto l89
								}
								{
									add(ruleAction13, position)
								}
								add(ruleIntAttr3, position1213)
							}
							break
						default:
							{
								position1227 := position
								if !_rules[rulews]() {
									goto l89
								}
								{
									position1228, tokenIndex1228 := position, tokenIndex
									if buffer[position] != rune('s') {
										goto l1229
									}
									position++
									goto l1228
								l1229:
									position, tokenIndex = position1228, tokenIndex1228
									if buffer[position] != rune('S') {
										goto l89
									}
									position++
								}
							l1228:
								{
									position1230, tokenIndex1230 := position, tokenIndex
									if buffer[position] != rune('p') {
										goto l1231
									}
									position++
									goto l1230
								l1231:
									position, tokenIndex = position1230, tokenIndex1230
									if buffer[position] != rune('P') {
										goto l89
									}
									position++
								}
							l1230:
								{
									position1232, tokenIndex1232 := position, tokenIndex
									if buffer[position] != rune('e') {
										goto l1233
									}
									position++
									goto l1232
								l1233:
									position, tokenIndex = position1232, tokenIndex1232
									if buffer[position] != rune('E') {
										goto l89
									}
									position++
								}
							l1232:
								{
									position1234, tokenIndex1234 := position, tokenIndex
									if buffer[position] != rune('c') {
										goto l1235
									}
									position++
									goto l1234
								l1235:
									position, tokenIndex = position1234, tokenIndex1234
									if buffer[position] != rune('C') {
										goto l89
									}
									position++
								}
							l1234:
								if !_rules[rulews]() {
									goto l89
								}
								{
									add(ruleAction26, position)
								}
								add(ruleRandomSpec, position1227)
							}
							break
						}
					}

				}
			l91:
				{
					add(ruleAction27, position)
				}
				add(ruleOneAttr, position90)
			}
			return true
		l89:
			position, tokenIndex = position89, tokenIndex89
			return false
		},
		/* 32 Attrs <- <(ws OneAttr)*> */
		nil,
		/* 33 Attr <- <(OneAttr ws Attrs ws)?> */
		nil,
		/* 34 Statement <- <(<PNode> ws <NodeType> ws <Attr> ws Action28)> */
		nil,
		/* 35 StringAttr <- <(('s' / 'S') ('t' / 'T') ('r' / 'R') ('g' / 'G') ':' ' ' <((!'l' .) / (('l' / 'L') (!'n' .)) / (('l' / 'L') ('n' / 'N') (!'g' .)) / (('l' / 'L') ('n' / 'N') ('g' / 'G') (!'t' .)) / (('l' / 'L') ('n' / 'N') ('g' / 'G') ('t' / 'T') (!':' .)))+> &(('l' / 'L') ('n' / 'N') ('g' / 'G') ('t' / 'T') ':') Action29)> */
		nil,
		/* 36 EOF <- <!.> */
		nil,
		/* 37 DecimalDigit <- <([0-9] Action30)> */
		func() bool {
			position1243, tokenIndex1243 := position, tokenIndex
			{
				position1244 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l1243
				}
				position++
				{
					add(ruleAction30, position)
				}
				add(ruleDecimalDigit, position1244)
			}
			return true
		l1243:
			position, tokenIndex = position1243, tokenIndex1243
			return false
		},
		/* 38 NonZeroDecimalDigit <- <([1-9] Action31)> */
		func() bool {
			position1246, tokenIndex1246 := position, tokenIndex
			{
				position1247 := position
				if c := buffer[position]; c < rune('1') || c > rune('9') {
					goto l1246
				}
				position++
				{
					add(ruleAction31, position)
				}
				add(ruleNonZeroDecimalDigit, position1247)
			}
			return true
		l1246:
			position, tokenIndex = position1246, tokenIndex1246
			return false
		},
		/* 39 Node <- <('@' <(NonZeroDecimalDigit DecimalDigit*)> Action32)> */
		func() bool {
			position1249, tokenIndex1249 := position, tokenIndex
			{
				position1250 := position
				if buffer[position] != rune('@') {
					goto l1249
				}
				position++
				{
					position1251 := position
					if !_rules[ruleNonZeroDecimalDigit]() {
						goto l1249
					}
				l1252:
					{
						position1253, tokenIndex1253 := position, tokenIndex
						if !_rules[ruleDecimalDigit]() {
							goto l1253
						}
						goto l1252
					l1253:
						position, tokenIndex = position1253, tokenIndex1253
					}
					add(rulePegText, position1251)
				}
				{
					add(ruleAction32, position)
				}
				add(ruleNode, position1250)
			}
			return true
		l1249:
			position, tokenIndex = position1249, tokenIndex1249
			return false
		},
		/* 40 PNode <- <(Node Action33)> */
		nil,
		/* 41 PosInt <- <(NonZeroDecimalDigit DecimalDigit* Action34)> */
		func() bool {
			position1256, tokenIndex1256 := position, tokenIndex
			{
				position1257 := position
				if !_rules[ruleNonZeroDecimalDigit]() {
					goto l1256
				}
			l1258:
				{
					position1259, tokenIndex1259 := position, tokenIndex
					if !_rules[ruleDecimalDigit]() {
						goto l1259
					}
					goto l1258
				l1259:
					position, tokenIndex = position1259, tokenIndex1259
				}
				{
					add(ruleAction34, position)
				}
				add(rulePosInt, position1257)
			}
			return true
		l1256:
			position, tokenIndex = position1256, tokenIndex1256
			return false
		},
		/* 42 Integer <- <(<('0' / PosInt)> Action35)> */
		func() bool {
			position1261, tokenIndex1261 := position, tokenIndex
			{
				position1262 := position
				{
					position1263 := position
					{
						position1264, tokenIndex1264 := position, tokenIndex
						if buffer[position] != rune('0') {
							goto l1265
						}
						position++
						goto l1264
					l1265:
						position, tokenIndex = position1264, tokenIndex1264
						if !_rules[rulePosInt]() {
							goto l1261
						}
					}
				l1264:
					add(rulePegText, position1263)
				}
				{
					add(ruleAction35, position)
				}
				add(ruleInteger, position1262)
			}
			return true
		l1261:
			position, tokenIndex = position1261, tokenIndex1261
			return false
		},
		/* 43 Hex <- <[a-h]> */
		func() bool {
			position1267, tokenIndex1267 := position, tokenIndex
			{
				position1268 := position
				if c := buffer[position]; c < rune('a') || c > rune('h') {
					goto l1267
				}
				position++
				add(ruleHex, position1268)
			}
			return true
		l1267:
			position, tokenIndex = position1267, tokenIndex1267
			return false
		},
		/* 44 HexValue <- <('0' 'x' <(DecimalDigit / Hex)+> Action36)> */
		func() bool {
			position1269, tokenIndex1269 := position, tokenIndex
			{
				position1270 := position
				if buffer[position] != rune('0') {
					goto l1269
				}
				position++
				if buffer[position] != rune('x') {
					goto l1269
				}
				position++
				{
					position1271 := position
					{
						position1274, tokenIndex1274 := position, tokenIndex
						if !_rules[ruleDecimalDigit]() {
							goto l1275
						}
						goto l1274
					l1275:
						position, tokenIndex = position1274, tokenIndex1274
						if !_rules[ruleHex]() {
							goto l1269
						}
					}
				l1274:
				l1272:
					{
						position1273, tokenIndex1273 := position, tokenIndex
						{
							position1276, tokenIndex1276 := position, tokenIndex
							if !_rules[ruleDecimalDigit]() {
								goto l1277
							}
							goto l1276
						l1277:
							position, tokenIndex = position1276, tokenIndex1276
							if !_rules[ruleHex]() {
								goto l1273
							}
						}
					l1276:
						goto l1272
					l1273:
						position, tokenIndex = position1273, tokenIndex1273
					}
					add(rulePegText, position1271)
				}
				{
					add(ruleAction36, position)
				}
				add(ruleHexValue, position1270)
			}
			return true
		l1269:
			position, tokenIndex = position1269, tokenIndex1269
			return false
		},
		/* 45 NegInt <- <(<('-' PosInt)> Action37)> */
		func() bool {
			position1279, tokenIndex1279 := position, tokenIndex
			{
				position1280 := position
				{
					position1281 := position
					if buffer[position] != rune('-') {
						goto l1279
					}
					position++
					if !_rules[rulePosInt]() {
						goto l1279
					}
					add(rulePegText, position1281)
				}
				{
					add(ruleAction37, position)
				}
				add(ruleNegInt, position1280)
			}
			return true
		l1279:
			position, tokenIndex = position1279, tokenIndex1279
			return false
		},
		/* 46 SignedIntAttr <- <(('i' / 'I') ('n' / 'N') ('t' / 'T') ':' ws (HexValue / NegInt / Integer) Action38)> */
		nil,
		/* 47 IntAttr2 <- <(<(('a' / 'A') ('l' / 'L') ('g' / 'G') ('n' / 'N'))> ':' ws <[0-9]+> ws Action39)> */
		nil,
		/* 49 Action0 <- <{}> */
		nil,
		/* 50 Action1 <- <{}> */
		nil,
		/* 51 Action2 <- <{}> */
		nil,
		nil,
		/* 53 Action3 <- <{
			fmt.Printf("Other Type: %s\n",buffer[begin:end])
		}> */
		nil,
		/* 54 Action4 <- <{
			s := astproto.NodeType(astproto.NodeType_value[buffer[begin:end]])
			NodeType=s
		}> */
		nil,
		/* 55 Action5 <- <{
		    OpNumber = Atoi(buffer[begin:end])
		    FieldName="opn"
		    //fmt.Printf("op check1 %s %d\n", string(buffer[begin:end]), OpNumber)
		}> */
		nil,
		/* 56 Action6 <- <{
		    OpNumber = Atoi(buffer[begin:end])
		    FieldName="opn"
		    //fmt.Printf("op check1 %s %d\n", string(buffer[begin:end]), OpNumber)
		}> */
		nil,
		/* 57 Action7 <- <{
		   LineNum = Atoi(buffer[begin:end])
		}> */
		nil,
		/* 58 Action8 <- <{
			FileName=string(buffer[begin:end])
		}> */
		nil,
		/* 59 Action9 <- <{
			//fmt.Printf("FileName and Number: '%s:%d'\n",FileName, LineNum)
			getNode().AddFileRef(FileName,LineNum)
		}> */
		nil,
		/* 60 Action10 <- <{
			getNode().AddStringField("lang", "C")
		}> */
		nil,
		/* 61 Action11 <- <{
			// base
			getNode().AddIntField(FieldName,buffer[begin:end])
		}> */
		nil,
		/* 62 Action12 <- <{
			if FieldType == TNodeRef {
				getNode().AddNodeRef("low",NodeNumber)
			} else if FieldType == TInteger {
				getNode().AddIntField("low",IntVal)
			} else if FieldType == THex {
				getNode().AddHexField("low",HexVal)
			} else {
				fmt.Printf("unkown field type : %s\n",buffer[begin:end])
				fmt.Printf("unkown field type : %s\n",buffer[begin-30:end+30])
				//panic("unkown field")
				getNode().AddHexField("low",buffer[begin:end])
			}

		}> */
		nil,
		/* 63 Action13 <- <{
		    getNode().AddIntField(FieldName,buffer[begin:end])
		}> */
		nil,
		/* 64 Action14 <- <{
			// addr :
			getNode().AddHexField("addr",buffer[begin:end])
		}> */
		nil,
		/* 65 Action15 <- <{
			getNode().AddTag(astproto.TagType(astproto.TagType_value[buffer[begin:end]]))
		}> */
		nil,
		/* 66 Action16 <- <{

		}> */
		nil,
		/* 67 Action17 <- <{
			getNode().AddLink(astproto.LinkType(astproto.LinkType_value[buffer[begin:end]]))
		}> */
		nil,
		/* 68 Action18 <- <{
			getNode().AddNote(buffer[begin:end])
		}> */
		nil,
		/* 69 Action19 <- <{
			getNode().AddAccess(astproto.AccessType(astproto.AccessType_value[buffer[begin:end]]))
		}> */
		nil,
		/* 70 Action20 <- <{
			getNode().AddQual(astproto.QualType(astproto.QualType_value[buffer[begin:end]]))
		}> */
		nil,
		/* 71 Action21 <- <{
			getNode().AddSign(astproto.SignType(astproto.SignType_value[buffer[begin:end]]))
		}> */
		nil,
		/* 72 Action22 <- <{
			fmt.Printf("Other Field: %s\n",buffer[begin:end])
		}> */
		nil,
		/* 73 Action23 <- <{
			       FieldName=buffer[begin:end]
			//fmt.Printf("set fieldname %s\n",FieldName)
		   }> */
		nil,
		/* 74 Action24 <- <{
			if FieldName == "opn" {
				getNode().AddOpNodeRef(FieldName,OpNumber,NodeNumber)
			}else {
				getNode().AddNodeRef(FieldName,NodeNumber)
			}
		}> */
		nil,
		/* 75 Action25 <- <{
			s := astproto.Spec( astproto.Spec_value[buffer[begin:end]] )
			getNode().AddSpec(s)
		}> */
		nil,
		/* 76 Action26 <- <{}> */
		nil,
		/* 77 Action27 <- <{
		FieldType=TUnknown
		}> */
		nil,
		/* 78 Action28 <- <{
			// clear it
			n := getNode()
			//fmt.Printf("stmt %#v\n",n)
			nt := NodeType // copy it or it will get changed
			n.NodeType=&nt
			file.AddNode(n)
			clearNode()
		}> */
		nil,
		/* 79 Action29 <- <{
		// fmt.Printf("found string : bytes %d\n", end - begin)
		//	fmt.Printf("found string : %s\n", buffer[begin:end])
		    getNode().AddStringField(FieldName,buffer[begin:end])
		    FieldType=TString
		}> */
		nil,
		/* 80 Action30 <- <{}> */
		nil,
		/* 81 Action31 <- <{}> */
		nil,
		/* 82 Action32 <- <{
			//s:=buffer[begin:end]
			//fmt.Printf("noderef %s\n",s)
		    NodeNumber=Atoi(buffer[begin:end])
		    FieldType=TNodeRef
		}> */
		nil,
		/* 83 Action33 <- <{
		    MainNodeNumber=NodeNumber
		}> */
		nil,
		/* 84 Action34 <- <{}> */
		nil,
		/* 85 Action35 <- <{
		    IntVal=buffer[begin:end]
		    FieldType = TInteger
		}> */
		nil,
		/* 86 Action36 <- <{
		    HexVal = string(buffer[begin:end])
		    FieldType = THex
		}> */
		nil,
		/* 87 Action37 <- <{
		// NegInt
		    IntVal=buffer[begin:end]
		    FieldType = TInteger
		}> */
		nil,
		/* 88 Action38 <- <{
			// int
			//fmt.Printf("signed check %s\n", string(buffer[0:end]))
		//fmt.Printf("signed check %s\n", string(buffer[begin:end]))
			getNode().AddIntStringField(FieldName,buffer[begin:end])
		}> */
		nil,
		/* 89 Action39 <- <{
		    // algn
		    //fmt.Printf("algn %s\n", string(buffer[begin:end]))
		    getNode().AddIntField("algn",buffer[begin:end])
		}> */
		nil,
	}
	p.rules = _rules
}
