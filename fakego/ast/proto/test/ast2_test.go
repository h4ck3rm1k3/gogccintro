package tests

import (
	"testing"
	//"reflect"
	//"fmt"

	///
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast"
	"github.com/h4ck3rm1k3/gogccintro/fakego/ast/proto"
	//"go/parser"
	"github.com/h4ck3rm1k3/gogccintro/fakego/token"
)

func TestVistAst(b* testing.T)  {
	////
	p:= astproto.CreateTable()

	p.StrmapFile("0xc4200a2380",&ast.File/*struct*/{
  Package: 249,
  Name: p.StrmapIdent("0xc420010280",&ast.Ident/*struct*/{
    NamePos: 257,
    Name: "ast",
  }/*struct*/),
  Decls: []ast.Decl /*Slice*/{
    0: p.StrmapGenDecl("0xc420012900",&ast.GenDecl/*struct*/{
      TokPos: 262,
      Tok: token.Token(75)/*import*/,
      Lparen: 269,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapImportSpec("0xc42000a5d0",&ast.ImportSpec/*struct*/{
          Path: p.StrmapBasicLit("0xc4200102a0",&ast.BasicLit/*struct*/{
            ValuePos: 272,
            Kind: token.Token(9)/*STRING*/,
            Value: "\"go/token\"",
          }/*struct*/),
          EndPos: 0,
        }/*struct*/),/* slice_item: 0*/1: p.StrmapImportSpec("0xc42000a600",&ast.ImportSpec/*struct*/{
          Path: p.StrmapBasicLit("0xc4200102c0",&ast.BasicLit/*struct*/{
            ValuePos: 284,
            Kind: token.Token(9)/*STRING*/,
            Value: "\"strings\"",
          }/*struct*/),
          EndPos: 0,
        }/*struct*/),/* slice_item: 1*/2: p.StrmapImportSpec("0xc42000a630",&ast.ImportSpec/*struct*/{
          Path: p.StrmapBasicLit("0xc420010300",&ast.BasicLit/*struct*/{
            ValuePos: 295,
            Kind: token.Token(9)/*STRING*/,
            Value: "\"unicode\"",
          }/*struct*/),
          EndPos: 0,
        }/*struct*/),/* slice_item: 2*/3: p.StrmapImportSpec("0xc42000a660",&ast.ImportSpec/*struct*/{
          Path: p.StrmapBasicLit("0xc420010340",&ast.BasicLit/*struct*/{
            ValuePos: 306,
            Kind: token.Token(9)/*STRING*/,
            Value: "\"unicode/utf8\"",
          }/*struct*/),
          EndPos: 0,
        }/*struct*/),/* slice_item: 3*/}/*slice*/,
      Rparen: 321,
    }/*struct*/),/* slice_item: 0*/1: p.StrmapGenDecl("0xc420012b00",&ast.GenDecl/*struct*/{
      TokPos: 1209,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc42000a6f0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc420010380",&ast.Ident/*struct*/{
            NamePos: 1214,
            Name: "Node",
            Obj: p.StrmapObject("0xc420050550",&ast.Object/*struct*/{
              Kind: "type",
              Name: "Node",
              Decl: p.PtrmapTypeSpec("0xc42000a6f0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapInterfaceType("0xc4200104e0",&ast.InterfaceType/*struct*/{
            Interface: 1219,
            Methods: p.StrmapFieldList("0xc42000a7e0",&ast.FieldList/*struct*/{
              Opening: 1229,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200129c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200103a0",&ast.Ident/*struct*/{
                      NamePos: 1232,
                      Name: "Pos",
                      Obj: p.StrmapObject("0xc4200505a0",&ast.Object/*struct*/{
                        Kind: "func",
                        Name: "Pos",
                        Decl: p.PtrmapField("0xc4200129c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapFuncType("0xc420010420",&ast.FuncType/*struct*/{
                    Func: 0,
                    Params: p.StrmapFieldList("0xc42000a720",&ast.FieldList/*struct*/{
                      Opening: 1235,
                      Closing: 1236,
                    }/*struct*/),
                    Results: p.StrmapFieldList("0xc42000a750",&ast.FieldList/*struct*/{
                      Opening: 0,
                      List: []*ast.Field /*Slice*/{
                        0: p.StrmapField("0xc420012940",&ast.Field/*struct*/{
                          Type: p.StrmapSelectorExpr("0xc420010400",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200103c0",&ast.Ident/*struct*/{
                              NamePos: 1238,
                              Name: "token",
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200103e0",&ast.Ident/*struct*/{
                              NamePos: 1244,
                              Name: "Pos",
                            }/*struct*/),
                          }/*struct*/),
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Closing: 0,
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc420012a80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc420010440",&ast.Ident/*struct*/{
                      NamePos: 1302,
                      Name: "End",
                      Obj: p.StrmapObject("0xc4200505f0",&ast.Object/*struct*/{
                        Kind: "func",
                        Name: "End",
                        Decl: p.PtrmapField("0xc420012a80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapFuncType("0xc4200104c0",&ast.FuncType/*struct*/{
                    Func: 0,
                    Params: p.StrmapFieldList("0xc42000a780",&ast.FieldList/*struct*/{
                      Opening: 1305,
                      Closing: 1306,
                    }/*struct*/),
                    Results: p.StrmapFieldList("0xc42000a7b0",&ast.FieldList/*struct*/{
                      Opening: 0,
                      List: []*ast.Field /*Slice*/{
                        0: p.StrmapField("0xc420012a00",&ast.Field/*struct*/{
                          Type: p.StrmapSelectorExpr("0xc4200104a0",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc420010460",&ast.Ident/*struct*/{
                              NamePos: 1308,
                              Name: "token",
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc420010480",&ast.Ident/*struct*/{
                              NamePos: 1314,
                              Name: "Pos",
                            }/*struct*/),
                          }/*struct*/),
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Closing: 0,
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 1376,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 1*/2: p.StrmapGenDecl("0xc420012c00",&ast.GenDecl/*struct*/{
      TokPos: 1433,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc42000a810",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc420010520",&ast.Ident/*struct*/{
            NamePos: 1438,
            Name: "Expr",
            Obj: p.StrmapObject("0xc420050640",&ast.Object/*struct*/{
              Kind: "type",
              Name: "Expr",
              Decl: p.PtrmapTypeSpec("0xc42000a810"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapInterfaceType("0xc4200105a0",&ast.InterfaceType/*struct*/{
            Interface: 1443,
            Methods: p.StrmapFieldList("0xc42000a870",&ast.FieldList/*struct*/{
              Opening: 1453,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc420012b40",&ast.Field/*struct*/{
                  Type: p.StrmapIdent("0xc420010540",&ast.Ident/*struct*/{
                    NamePos: 1456,
                    Name: "Node",
                    Obj: p.PtrmapObject("0xc420050550"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc420012b80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc420010560",&ast.Ident/*struct*/{
                      NamePos: 1462,
                      Name: "exprNode",
                      Obj: p.StrmapObject("0xc420050690",&ast.Object/*struct*/{
                        Kind: "func",
                        Name: "exprNode",
                        Decl: p.PtrmapField("0xc420012b80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapFuncType("0xc420010580",&ast.FuncType/*struct*/{
                    Func: 0,
                    Params: p.StrmapFieldList("0xc42000a840",&ast.FieldList/*struct*/{
                      Opening: 1470,
                      Closing: 1471,
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 1473,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 2*/3: p.StrmapGenDecl("0xc420012d40",&ast.GenDecl/*struct*/{
      TokPos: 1529,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc42000a8a0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200105c0",&ast.Ident/*struct*/{
            NamePos: 1534,
            Name: "Stmt",
            Obj: p.StrmapObject("0xc4200506e0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "Stmt",
              Decl: p.PtrmapTypeSpec("0xc42000a8a0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapInterfaceType("0xc420010640",&ast.InterfaceType/*struct*/{
            Interface: 1539,
            Methods: p.StrmapFieldList("0xc42000a900",&ast.FieldList/*struct*/{
              Opening: 1549,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc420012c80",&ast.Field/*struct*/{
                  Type: p.StrmapIdent("0xc4200105e0",&ast.Ident/*struct*/{
                    NamePos: 1552,
                    Name: "Node",
                    Obj: p.PtrmapObject("0xc420050550"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc420012cc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc420010600",&ast.Ident/*struct*/{
                      NamePos: 1558,
                      Name: "stmtNode",
                      Obj: p.StrmapObject("0xc420050730",&ast.Object/*struct*/{
                        Kind: "func",
                        Name: "stmtNode",
                        Decl: p.PtrmapField("0xc420012cc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapFuncType("0xc420010620",&ast.FuncType/*struct*/{
                    Func: 0,
                    Params: p.StrmapFieldList("0xc42000a8d0",&ast.FieldList/*struct*/{
                      Opening: 1566,
                      Closing: 1567,
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 1569,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 3*/4: p.StrmapGenDecl("0xc420012e00",&ast.GenDecl/*struct*/{
      TokPos: 1627,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc42000a930",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc420010660",&ast.Ident/*struct*/{
            NamePos: 1632,
            Name: "Decl",
            Obj: p.StrmapObject("0xc420050780",&ast.Object/*struct*/{
              Kind: "type",
              Name: "Decl",
              Decl: p.PtrmapTypeSpec("0xc42000a930"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapInterfaceType("0xc4200106e0",&ast.InterfaceType/*struct*/{
            Interface: 1637,
            Methods: p.StrmapFieldList("0xc42000a990",&ast.FieldList/*struct*/{
              Opening: 1647,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc420012d80",&ast.Field/*struct*/{
                  Type: p.StrmapIdent("0xc420010680",&ast.Ident/*struct*/{
                    NamePos: 1650,
                    Name: "Node",
                    Obj: p.PtrmapObject("0xc420050550"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc420012dc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200106a0",&ast.Ident/*struct*/{
                      NamePos: 1656,
                      Name: "declNode",
                      Obj: p.StrmapObject("0xc4200507d0",&ast.Object/*struct*/{
                        Kind: "func",
                        Name: "declNode",
                        Decl: p.PtrmapField("0xc420012dc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapFuncType("0xc4200106c0",&ast.FuncType/*struct*/{
                    Func: 0,
                    Params: p.StrmapFieldList("0xc42000a960",&ast.FieldList/*struct*/{
                      Opening: 1664,
                      Closing: 1665,
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 1667,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 4*/5: p.StrmapGenDecl("0xc420012f00",&ast.GenDecl/*struct*/{
      TokPos: 1831,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc42000a9c0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc420010700",&ast.Ident/*struct*/{
            NamePos: 1836,
            Name: "Comment",
            Obj: p.StrmapObject("0xc4200508c0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "Comment",
              Decl: p.PtrmapTypeSpec("0xc42000a9c0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc420010800",&ast.StructType/*struct*/{
            Struct: 1844,
            Fields: p.StrmapFieldList("0xc42000aa20",&ast.FieldList/*struct*/{
              Opening: 1851,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc420012e40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc420010720",&ast.Ident/*struct*/{
                      NamePos: 1854,
                      Name: "Slash",
                      Obj: p.StrmapObject("0xc420050910",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Slash",
                        Decl: p.PtrmapField("0xc420012e40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200107a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc420010740",&ast.Ident/*struct*/{
                      NamePos: 1860,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc420010780",&ast.Ident/*struct*/{
                      NamePos: 1866,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc420012ec0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200107c0",&ast.Ident/*struct*/{
                      NamePos: 1911,
                      Name: "Text",
                      Obj: p.StrmapObject("0xc420050960",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Text",
                        Decl: p.PtrmapField("0xc420012ec0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200107e0",&ast.Ident/*struct*/{
                    NamePos: 1917,
                    Name: "string",
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 1982,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 5*/6: p.StrmapFuncDecl("0xc42000ab70",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc42000aa80",&ast.FieldList/*struct*/{
        Opening: 1990,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc420012f40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc420010820",&ast.Ident/*struct*/{
                NamePos: 1991,
                Name: "c",
                Obj: p.StrmapObject("0xc4200509b0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "c",
                  Decl: p.PtrmapField("0xc420012f40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc420010860",&ast.StarExpr/*struct*/{
              Star: 1993,
              X: p.StrmapIdent("0xc420010840",&ast.Ident/*struct*/{
                NamePos: 1994,
                Name: "Comment",
                Obj: p.PtrmapObject("0xc4200508c0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 2001,
      }/*struct*/),
      Name: p.StrmapIdent("0xc420010880",&ast.Ident/*struct*/{
        NamePos: 2003,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200109a0",&ast.FuncType/*struct*/{
        Func: 1985,
        Params: p.StrmapFieldList("0xc42000aab0",&ast.FieldList/*struct*/{
          Opening: 2006,
          Closing: 2007,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc42000aae0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc420012fc0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200108e0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200108a0",&ast.Ident/*struct*/{
                  NamePos: 2009,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200108c0",&ast.Ident/*struct*/{
                  NamePos: 2015,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc42000ab40",&ast.BlockStmt/*struct*/{
        Lbrace: 2019,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc420010980",&ast.ReturnStmt/*struct*/{
            Return: 2021,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc420010960",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc420010920",&ast.Ident/*struct*/{
                  NamePos: 2028,
                  Name: "c",
                  Obj: p.PtrmapObject("0xc4200509b0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc420010940",&ast.Ident/*struct*/{
                  NamePos: 2030,
                  Name: "Slash",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 2036,
      }/*struct*/),
    }/*struct*/),/* slice_item: 6*/7: p.StrmapFuncDecl("0xc42000acf0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc42000abd0",&ast.FieldList/*struct*/{
        Opening: 2043,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc420013000",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200109c0",&ast.Ident/*struct*/{
                NamePos: 2044,
                Name: "c",
                Obj: p.StrmapObject("0xc420050a00",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "c",
                  Decl: p.PtrmapField("0xc420013000"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc420010a00",&ast.StarExpr/*struct*/{
              Star: 2046,
              X: p.StrmapIdent("0xc4200109e0",&ast.Ident/*struct*/{
                NamePos: 2047,
                Name: "Comment",
                Obj: p.PtrmapObject("0xc4200508c0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 2054,
      }/*struct*/),
      Name: p.StrmapIdent("0xc420010a20",&ast.Ident/*struct*/{
        NamePos: 2056,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc420010c20",&ast.FuncType/*struct*/{
        Func: 2038,
        Params: p.StrmapFieldList("0xc42000ac00",&ast.FieldList/*struct*/{
          Opening: 2059,
          Closing: 2060,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc42000ac30",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc420013040",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc420010a80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc420010a40",&ast.Ident/*struct*/{
                  NamePos: 2062,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc420010a60",&ast.Ident/*struct*/{
                  NamePos: 2068,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc42000acc0",&ast.BlockStmt/*struct*/{
        Lbrace: 2072,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc420010c00",&ast.ReturnStmt/*struct*/{
            Return: 2074,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc420013100",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc420010ae0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc420010aa0",&ast.Ident/*struct*/{
                    NamePos: 2081,
                    Name: "token",
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc420010ac0",&ast.Ident/*struct*/{
                    NamePos: 2087,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 2090,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapBinaryExpr("0xc42000ac90",&ast.BinaryExpr/*struct*/{
                    X: p.StrmapCallExpr("0xc420013080",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapIdent("0xc420010b00",&ast.Ident/*struct*/{
                        NamePos: 2091,
                        Name: "int",
                      }/*struct*/),
                      Lparen: 2094,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapSelectorExpr("0xc420010b60",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc420010b20",&ast.Ident/*struct*/{
                            NamePos: 2095,
                            Name: "c",
                            Obj: p.PtrmapObject("0xc420050a00"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc420010b40",&ast.Ident/*struct*/{
                            NamePos: 2097,
                            Name: "Slash",
                          }/*struct*/),
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 2102,
                    }/*struct*/),
                    OpPos: 2104,
                    Op: token.Token(12)/*+*/,
                    Y: p.StrmapCallExpr("0xc4200130c0",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapIdent("0xc420010b80",&ast.Ident/*struct*/{
                        NamePos: 2106,
                        Name: "len",
                      }/*struct*/),
                      Lparen: 2109,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapSelectorExpr("0xc420010be0",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc420010ba0",&ast.Ident/*struct*/{
                            NamePos: 2110,
                            Name: "c",
                            Obj: p.PtrmapObject("0xc420050a00"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc420010bc0",&ast.Ident/*struct*/{
                            NamePos: 2112,
                            Name: "Text",
                          }/*struct*/),
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 2116,
                    }/*struct*/),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 2117,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 2119,
      }/*struct*/),
    }/*struct*/),/* slice_item: 7*/8: p.StrmapGenDecl("0xc420013200",&ast.GenDecl/*struct*/{
      TokPos: 2229,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc42000ad20",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc420010c40",&ast.Ident/*struct*/{
            NamePos: 2234,
            Name: "CommentGroup",
            Obj: p.StrmapObject("0xc420050a50",&ast.Object/*struct*/{
              Kind: "type",
              Name: "CommentGroup",
              Decl: p.PtrmapTypeSpec("0xc42000ad20"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc420010cc0",&ast.StructType/*struct*/{
            Struct: 2247,
            Fields: p.StrmapFieldList("0xc42000ad80",&ast.FieldList/*struct*/{
              Opening: 2254,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200131c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc420010c60",&ast.Ident/*struct*/{
                      NamePos: 2257,
                      Name: "List",
                      Obj: p.StrmapObject("0xc420050aa0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "List",
                        Decl: p.PtrmapField("0xc4200131c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc42000ad50",&ast.ArrayType/*struct*/{
                    Lbrack: 2262,
                    Elt: p.StrmapStarExpr("0xc420010ca0",&ast.StarExpr/*struct*/{
                      Star: 2264,
                      X: p.StrmapIdent("0xc420010c80",&ast.Ident/*struct*/{
                        NamePos: 2265,
                        Name: "Comment",
                        Obj: p.PtrmapObject("0xc4200508c0"),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Closing: 2290,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 8*/9: p.StrmapFuncDecl("0xc42000af00",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc42000ade0",&ast.FieldList/*struct*/{
        Opening: 2298,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc420013240",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc420010ce0",&ast.Ident/*struct*/{
                NamePos: 2299,
                Name: "g",
                Obj: p.StrmapObject("0xc420050af0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "g",
                  Decl: p.PtrmapField("0xc420013240"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc420010d20",&ast.StarExpr/*struct*/{
              Star: 2301,
              X: p.StrmapIdent("0xc420010d00",&ast.Ident/*struct*/{
                NamePos: 2302,
                Name: "CommentGroup",
                Obj: p.PtrmapObject("0xc420050a50"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 2314,
      }/*struct*/),
      Name: p.StrmapIdent("0xc420010d40",&ast.Ident/*struct*/{
        NamePos: 2316,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc420010ea0",&ast.FuncType/*struct*/{
        Func: 2293,
        Params: p.StrmapFieldList("0xc42000ae10",&ast.FieldList/*struct*/{
          Opening: 2319,
          Closing: 2320,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc42000ae40",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc420013280",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc420010da0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc420010d60",&ast.Ident/*struct*/{
                  NamePos: 2322,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc420010d80",&ast.Ident/*struct*/{
                  NamePos: 2328,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc42000aed0",&ast.BlockStmt/*struct*/{
        Lbrace: 2332,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc420010e80",&ast.ReturnStmt/*struct*/{
            Return: 2334,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200132c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc420010e60",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIndexExpr("0xc42000aea0",&ast.IndexExpr/*struct*/{
                    X: p.StrmapSelectorExpr("0xc420010e00",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc420010dc0",&ast.Ident/*struct*/{
                        NamePos: 2341,
                        Name: "g",
                        Obj: p.PtrmapObject("0xc420050af0"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc420010de0",&ast.Ident/*struct*/{
                        NamePos: 2343,
                        Name: "List",
                      }/*struct*/),
                    }/*struct*/),
                    Lbrack: 2347,
                    Index: p.StrmapBasicLit("0xc420010e20",&ast.BasicLit/*struct*/{
                      ValuePos: 2348,
                      Kind: token.Token(5)/*INT*/,
                      Value: "0",
                    }/*struct*/),
                    Rbrack: 2349,
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc420010e40",&ast.Ident/*struct*/{
                    NamePos: 2351,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 2354,
                Ellipsis: 0,
                Rparen: 2355,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 2357,
      }/*struct*/),
    }/*struct*/),/* slice_item: 9*/10: p.StrmapFuncDecl("0xc42000b0b0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc42000af60",&ast.FieldList/*struct*/{
        Opening: 2364,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc420013300",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc420010ec0",&ast.Ident/*struct*/{
                NamePos: 2365,
                Name: "g",
                Obj: p.StrmapObject("0xc420050b40",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "g",
                  Decl: p.PtrmapField("0xc420013300"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc420010f00",&ast.StarExpr/*struct*/{
              Star: 2367,
              X: p.StrmapIdent("0xc420010ee0",&ast.Ident/*struct*/{
                NamePos: 2368,
                Name: "CommentGroup",
                Obj: p.PtrmapObject("0xc420050a50"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 2380,
      }/*struct*/),
      Name: p.StrmapIdent("0xc420010f20",&ast.Ident/*struct*/{
        NamePos: 2382,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc420011100",&ast.FuncType/*struct*/{
        Func: 2359,
        Params: p.StrmapFieldList("0xc42000af90",&ast.FieldList/*struct*/{
          Opening: 2385,
          Closing: 2386,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc42000afc0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc420013340",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc420010f80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc420010f40",&ast.Ident/*struct*/{
                  NamePos: 2388,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc420010f60",&ast.Ident/*struct*/{
                  NamePos: 2394,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc42000b080",&ast.BlockStmt/*struct*/{
        Lbrace: 2398,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200110e0",&ast.ReturnStmt/*struct*/{
            Return: 2400,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200133c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200110c0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIndexExpr("0xc42000b050",&ast.IndexExpr/*struct*/{
                    X: p.StrmapSelectorExpr("0xc420010fe0",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc420010fa0",&ast.Ident/*struct*/{
                        NamePos: 2407,
                        Name: "g",
                        Obj: p.PtrmapObject("0xc420050b40"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc420010fc0",&ast.Ident/*struct*/{
                        NamePos: 2409,
                        Name: "List",
                      }/*struct*/),
                    }/*struct*/),
                    Lbrack: 2413,
                    Index: p.StrmapBinaryExpr("0xc42000b020",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapCallExpr("0xc420013380",&ast.CallExpr/*struct*/{
                        Fun: p.StrmapIdent("0xc420011000",&ast.Ident/*struct*/{
                          NamePos: 2414,
                          Name: "len",
                        }/*struct*/),
                        Lparen: 2417,
                        Args: []ast.Expr /*Slice*/{
                          0: p.StrmapSelectorExpr("0xc420011060",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc420011020",&ast.Ident/*struct*/{
                              NamePos: 2418,
                              Name: "g",
                              Obj: p.PtrmapObject("0xc420050b40"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc420011040",&ast.Ident/*struct*/{
                              NamePos: 2420,
                              Name: "List",
                            }/*struct*/),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                        Ellipsis: 0,
                        Rparen: 2424,
                      }/*struct*/),
                      OpPos: 2425,
                      Op: token.Token(13)/*-*/,
                      Y: p.StrmapBasicLit("0xc420011080",&ast.BasicLit/*struct*/{
                        ValuePos: 2426,
                        Kind: token.Token(5)/*INT*/,
                        Value: "1",
                      }/*struct*/),
                    }/*struct*/),
                    Rbrack: 2427,
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200110a0",&ast.Ident/*struct*/{
                    NamePos: 2429,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 2432,
                Ellipsis: 0,
                Rparen: 2433,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 2435,
      }/*struct*/),
    }/*struct*/),/* slice_item: 10*/11: p.StrmapFuncDecl("0xc42000b320",&ast.FuncDecl/*struct*/{
      Name: p.StrmapIdent("0xc420011120",&ast.Ident/*struct*/{
        NamePos: 2443,
        Name: "isWhitespace",
        Obj: p.StrmapObject("0xc420050be0",&ast.Object/*struct*/{
          Kind: "func",
          Name: "isWhitespace",
          Decl: p.PtrmapFuncDecl("0xc42000b320"),
        }/*struct*/),
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200112c0",&ast.FuncType/*struct*/{
        Func: 2438,
        Params: p.StrmapFieldList("0xc42000b110",&ast.FieldList/*struct*/{
          Opening: 2455,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc420013400",&ast.Field/*struct*/{
              Names: []*ast.Ident /*Slice*/{
                0: p.StrmapIdent("0xc420011140",&ast.Ident/*struct*/{
                  NamePos: 2456,
                  Name: "ch",
                  Obj: p.StrmapObject("0xc420050b90",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "ch",
                    Decl: p.PtrmapField("0xc420013400"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Type: p.StrmapIdent("0xc420011160",&ast.Ident/*struct*/{
                NamePos: 2459,
                Name: "byte",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 2463,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc42000b140",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc420013440",&ast.Field/*struct*/{
              Type: p.StrmapIdent("0xc420011180",&ast.Ident/*struct*/{
                NamePos: 2465,
                Name: "bool",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc42000b2f0",&ast.BlockStmt/*struct*/{
        Lbrace: 2470,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200112a0",&ast.ReturnStmt/*struct*/{
            Return: 2472,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc42000b2c0",&ast.BinaryExpr/*struct*/{
                X: p.StrmapBinaryExpr("0xc42000b260",&ast.BinaryExpr/*struct*/{
                  X: p.StrmapBinaryExpr("0xc42000b200",&ast.BinaryExpr/*struct*/{
                    X: p.StrmapBinaryExpr("0xc42000b1a0",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200111a0",&ast.Ident/*struct*/{
                        NamePos: 2479,
                        Name: "ch",
                        Obj: p.PtrmapObject("0xc420050b90"),
                      }/*struct*/),
                      OpPos: 2482,
                      Op: token.Token(39)/*==*/,
                      Y: p.StrmapBasicLit("0xc4200111c0",&ast.BasicLit/*struct*/{
                        ValuePos: 2485,
                        Kind: token.Token(8)/*CHAR*/,
                        Value: "' '",
                      }/*struct*/),
                    }/*struct*/),
                    OpPos: 2489,
                    Op: token.Token(35)/*||*/,
                    Y: p.StrmapBinaryExpr("0xc42000b1d0",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200111e0",&ast.Ident/*struct*/{
                        NamePos: 2492,
                        Name: "ch",
                        Obj: p.PtrmapObject("0xc420050b90"),
                      }/*struct*/),
                      OpPos: 2495,
                      Op: token.Token(39)/*==*/,
                      Y: p.StrmapBasicLit("0xc420011200",&ast.BasicLit/*struct*/{
                        ValuePos: 2498,
                        Kind: token.Token(8)/*CHAR*/,
                        Value: "'\\t'",
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                  OpPos: 2503,
                  Op: token.Token(35)/*||*/,
                  Y: p.StrmapBinaryExpr("0xc42000b230",&ast.BinaryExpr/*struct*/{
                    X: p.StrmapIdent("0xc420011220",&ast.Ident/*struct*/{
                      NamePos: 2506,
                      Name: "ch",
                      Obj: p.PtrmapObject("0xc420050b90"),
                    }/*struct*/),
                    OpPos: 2509,
                    Op: token.Token(39)/*==*/,
                    Y: p.StrmapBasicLit("0xc420011240",&ast.BasicLit/*struct*/{
                      ValuePos: 2512,
                      Kind: token.Token(8)/*CHAR*/,
                      Value: "'\\n'",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),
                OpPos: 2517,
                Op: token.Token(35)/*||*/,
                Y: p.StrmapBinaryExpr("0xc42000b290",&ast.BinaryExpr/*struct*/{
                  X: p.StrmapIdent("0xc420011260",&ast.Ident/*struct*/{
                    NamePos: 2520,
                    Name: "ch",
                    Obj: p.PtrmapObject("0xc420050b90"),
                  }/*struct*/),
                  OpPos: 2523,
                  Op: token.Token(39)/*==*/,
                  Y: p.StrmapBasicLit("0xc420011280",&ast.BasicLit/*struct*/{
                    ValuePos: 2526,
                    Kind: token.Token(8)/*CHAR*/,
                    Value: "'\\r'",
                  }/*struct*/),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 2531,
      }/*struct*/),
    }/*struct*/),/* slice_item: 11*/12: p.StrmapFuncDecl("0xc42000b5f0",&ast.FuncDecl/*struct*/{
      Name: p.StrmapIdent("0xc420011300",&ast.Ident/*struct*/{
        NamePos: 2539,
        Name: "stripTrailingWhitespace",
        Obj: p.StrmapObject("0xc420050dc0",&ast.Object/*struct*/{
          Kind: "func",
          Name: "stripTrailingWhitespace",
          Decl: p.PtrmapFuncDecl("0xc42000b5f0"),
        }/*struct*/),
      }/*struct*/),
      Type: p.StrmapFuncType("0xc420011580",&ast.FuncType/*struct*/{
        Func: 2534,
        Params: p.StrmapFieldList("0xc42000b380",&ast.FieldList/*struct*/{
          Opening: 2562,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc420013480",&ast.Field/*struct*/{
              Names: []*ast.Ident /*Slice*/{
                0: p.StrmapIdent("0xc420011320",&ast.Ident/*struct*/{
                  NamePos: 2563,
                  Name: "s",
                  Obj: p.StrmapObject("0xc420050c30",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "s",
                    Decl: p.PtrmapField("0xc420013480"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Type: p.StrmapIdent("0xc420011340",&ast.Ident/*struct*/{
                NamePos: 2565,
                Name: "string",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 2571,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc42000b3b0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200134c0",&ast.Field/*struct*/{
              Type: p.StrmapIdent("0xc420011360",&ast.Ident/*struct*/{
                NamePos: 2573,
                Name: "string",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc42000b560",&ast.BlockStmt/*struct*/{
        Lbrace: 2580,
	      List:
	      []ast.Stmt /*Slice*/{
          0: p.StrmapAssignStmt("0xc420013540",&ast.AssignStmt/*struct*/{
            Lhs: []ast.Expr /*Slice*/{
              0: p.StrmapIdent("0xc420011380",&ast.Ident/*struct*/{
                NamePos: 2583,
                Name: "i",
                Obj: p.StrmapObject("0xc420050c80",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "i",
                  Decl: p.PtrmapAssignStmt("0xc420013540"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            TokPos: 2585,
            Tok: token.Token(47)/*:=*/,
            Rhs: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc420013500",&ast.CallExpr/*struct*/{
                Fun: p.StrmapIdent("0xc4200113a0",&ast.Ident/*struct*/{
                  NamePos: 2588,
                  Name: "len",
                }/*struct*/),
                Lparen: 2591,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapIdent("0xc4200113c0",&ast.Ident/*struct*/{
                    NamePos: 2592,
                    Name: "s",
                    Obj: p.PtrmapObject("0xc420050c30"),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 2593,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/1: p.StrmapForStmt("0xc4200135c0",&ast.ForStmt/*struct*/{
            For: 2596,
            Cond: p.StrmapBinaryExpr("0xc42000b4d0",&ast.BinaryExpr/*struct*/{
              X: p.StrmapBinaryExpr("0xc42000b440",&ast.BinaryExpr/*struct*/{
                X: p.StrmapIdent("0xc4200113e0",&ast.Ident/*struct*/{
                  NamePos: 2600,
                  Name: "i",
                  Obj: p.PtrmapObject("0xc420050c80"),
                }/*struct*/),
                OpPos: 2602,
                Op: token.Token(41)/*>*/,
                Y: p.StrmapBasicLit("0xc420011400",&ast.BasicLit/*struct*/{
                  ValuePos: 2604,
                  Kind: token.Token(5)/*INT*/,
                  Value: "0",
                }/*struct*/),
              }/*struct*/),
              OpPos: 2606,
              Op: token.Token(34)/*&&*/,
              Y: p.StrmapCallExpr("0xc420013580",&ast.CallExpr/*struct*/{
                Fun: p.StrmapIdent("0xc420011420",&ast.Ident/*struct*/{
                  NamePos: 2609,
                  Name: "isWhitespace",
                  Obj: p.PtrmapObject("0xc420050be0"),
                }/*struct*/),
                Lparen: 2621,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapIndexExpr("0xc42000b4a0",&ast.IndexExpr/*struct*/{
                    X: p.StrmapIdent("0xc420011440",&ast.Ident/*struct*/{
                      NamePos: 2622,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc420050c30"),
                    }/*struct*/),
                    Lbrack: 2623,
                    Index: p.StrmapBinaryExpr("0xc42000b470",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapIdent("0xc420011460",&ast.Ident/*struct*/{
                        NamePos: 2624,
                        Name: "i",
                        Obj: p.PtrmapObject("0xc420050c80"),
                      }/*struct*/),
                      OpPos: 2625,
                      Op: token.Token(13)/*-*/,
                      Y: p.StrmapBasicLit("0xc420011480",&ast.BasicLit/*struct*/{
                        ValuePos: 2626,
                        Kind: token.Token(5)/*INT*/,
                        Value: "1",
                      }/*struct*/),
                    }/*struct*/),
                    Rbrack: 2627,
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 2628,
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc42000b530",&ast.BlockStmt/*struct*/{
              Lbrace: 2630,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapIncDecStmt("0xc4200114c0",&ast.IncDecStmt/*struct*/{
                  X: p.StrmapIdent("0xc4200114a0",&ast.Ident/*struct*/{
                    NamePos: 2634,
                    Name: "i",
                    Obj: p.PtrmapObject("0xc420050c80"),
                  }/*struct*/),
                  TokPos: 2635,
                  Tok: token.Token(38)/*--*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 2639,
            }/*struct*/),
          }/*struct*/),/* slice_item: 1*/2: p.StrmapReturnStmt("0xc420011560",&ast.ReturnStmt/*struct*/{
            Return: 2642,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSliceExpr("0xc42005a420",&ast.SliceExpr/*struct*/{
                X: p.StrmapIdent("0xc420011500",&ast.Ident/*struct*/{
                  NamePos: 2649,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc420050c30"),
                }/*struct*/),
                Lbrack: 2650,
                Low: p.StrmapBasicLit("0xc420011520",&ast.BasicLit/*struct*/{
                  ValuePos: 2651,
                  Kind: token.Token(5)/*INT*/,
                  Value: "0",
                }/*struct*/),
                High: p.StrmapIdent("0xc420011540",&ast.Ident/*struct*/{
                  NamePos: 2653,
                  Name: "i",
                  Obj: p.PtrmapObject("0xc420050c80"),
                }/*struct*/),
                Slice3: false,
                Rbrack: 2654,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 2*/}/*slice*/,
        Rbrace: 2656,
      }/*struct*/),
    }/*struct*/),/* slice_item: 12*/13: p.StrmapFuncDecl("0xc4200b41e0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc42000b650",&ast.FieldList/*struct*/{
        Opening: 2974,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc420013640",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200115a0",&ast.Ident/*struct*/{
                NamePos: 2975,
                Name: "g",
                Obj: p.StrmapObject("0xc420050e10",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "g",
                  Decl: p.PtrmapField("0xc420013640"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200115e0",&ast.StarExpr/*struct*/{
              Star: 2977,
              X: p.StrmapIdent("0xc4200115c0",&ast.Ident/*struct*/{
                NamePos: 2978,
                Name: "CommentGroup",
                Obj: p.PtrmapObject("0xc420050a50"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 2990,
      }/*struct*/),
      Name: p.StrmapIdent("0xc420011600",&ast.Ident/*struct*/{
        NamePos: 2992,
        Name: "Text",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200b0640",&ast.FuncType/*struct*/{
        Func: 2969,
        Params: p.StrmapFieldList("0xc42000b680",&ast.FieldList/*struct*/{
          Opening: 2996,
          Closing: 2997,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc42000b6b0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc420013680",&ast.Field/*struct*/{
              Type: p.StrmapIdent("0xc420011620",&ast.Ident/*struct*/{
                NamePos: 2999,
                Name: "string",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200b41b0",&ast.BlockStmt/*struct*/{
        Lbrace: 3006,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200136c0",&ast.IfStmt/*struct*/{
            If: 3009,
            Cond: p.StrmapBinaryExpr("0xc42000b740",&ast.BinaryExpr/*struct*/{
              X: p.StrmapIdent("0xc420011640",&ast.Ident/*struct*/{
                NamePos: 3012,
                Name: "g",
                Obj: p.PtrmapObject("0xc420050e10"),
              }/*struct*/),
              OpPos: 3014,
              Op: token.Token(39)/*==*/,
              Y: p.StrmapIdent("0xc420011660",&ast.Ident/*struct*/{
                NamePos: 3017,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc42000b7a0",&ast.BlockStmt/*struct*/{
              Lbrace: 3021,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200116a0",&ast.ReturnStmt/*struct*/{
                  Return: 3025,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapBasicLit("0xc420011680",&ast.BasicLit/*struct*/{
                      ValuePos: 3032,
                      Kind: token.Token(9)/*STRING*/,
                      Value: "\"\"",
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 3036,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapAssignStmt("0xc420013780",&ast.AssignStmt/*struct*/{
            Lhs: []ast.Expr /*Slice*/{
              0: p.StrmapIdent("0xc4200116c0",&ast.Ident/*struct*/{
                NamePos: 3039,
                Name: "comments",
                Obj: p.StrmapObject("0xc420050e60",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "comments",
                  Decl: p.PtrmapAssignStmt("0xc420013780"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            TokPos: 3048,
            Tok: token.Token(47)/*:=*/,
            Rhs: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc420013740",&ast.CallExpr/*struct*/{
                Fun: p.StrmapIdent("0xc4200116e0",&ast.Ident/*struct*/{
                  NamePos: 3051,
                  Name: "make",
                }/*struct*/),
                Lparen: 3055,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapArrayType("0xc42000b7d0",&ast.ArrayType/*struct*/{
                    Lbrack: 3056,
                    Elt: p.StrmapIdent("0xc420011700",&ast.Ident/*struct*/{
                      NamePos: 3058,
                      Name: "string",
                    }/*struct*/),
                  }/*struct*/),/* slice_item: 0*/1: p.StrmapCallExpr("0xc420013700",&ast.CallExpr/*struct*/{
                    Fun: p.StrmapIdent("0xc420011720",&ast.Ident/*struct*/{
                      NamePos: 3066,
                      Name: "len",
                    }/*struct*/),
                    Lparen: 3069,
                    Args: []ast.Expr /*Slice*/{
                      0: p.StrmapSelectorExpr("0xc420011780",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIdent("0xc420011740",&ast.Ident/*struct*/{
                          NamePos: 3070,
                          Name: "g",
                          Obj: p.PtrmapObject("0xc420050e10"),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc420011760",&ast.Ident/*struct*/{
                          NamePos: 3072,
                          Name: "List",
                        }/*struct*/),
                      }/*struct*/),/* slice_item: 0*/}/*slice*/,
                    Ellipsis: 0,
                    Rparen: 3076,
                  }/*struct*/),/* slice_item: 1*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 3077,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/2: p.StrmapRangeStmt("0xc420050f50",&ast.RangeStmt/*struct*/{
            For: 3080,
            Key: p.StrmapIdent("0xc4200117e0",&ast.Ident/*struct*/{
              NamePos: 3084,
              Name: "i",
              Obj: p.StrmapObject("0xc420050eb0",&ast.Object/*struct*/{
                Kind: "var",
                Name: "i",
                Decl: p.StrmapAssignStmt("0xc4200137c0",&ast.AssignStmt/*struct*/{
                  Lhs: []ast.Expr /*Slice*/{
                    0: p.PtrmapIdent("0xc4200117e0"),/* slice_item: 0*/1: p.StrmapIdent("0xc420011800",&ast.Ident/*struct*/{
                      NamePos: 3087,
                      Name: "c",
                      Obj: p.StrmapObject("0xc420050f00",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "c",
                        Decl: p.PtrmapAssignStmt("0xc4200137c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 1*/}/*slice*/,
                  TokPos: 3089,
                  Tok: token.Token(47)/*:=*/,
                  Rhs: []ast.Expr /*Slice*/{
                    0: p.StrmapUnaryExpr("0xc4200118a0",&ast.UnaryExpr/*struct*/{
                      OpPos: 3092,
                      Op: token.Token(79)/*range*/,
                      X: p.StrmapSelectorExpr("0xc420011880",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIdent("0xc420011840",&ast.Ident/*struct*/{
                          NamePos: 3098,
                          Name: "g",
                          Obj: p.PtrmapObject("0xc420050e10"),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc420011860",&ast.Ident/*struct*/{
                          NamePos: 3100,
                          Name: "List",
                        }/*struct*/),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),
            Value: p.PtrmapIdent("0xc420011800"),
            TokPos: 3089,
            Tok: token.Token(47)/*:=*/,
            X: p.PtrmapSelectorExpr("0xc420011880"),
            Body: p.StrmapBlockStmt("0xc42000b890",&ast.BlockStmt/*struct*/{
              Lbrace: 3105,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapAssignStmt("0xc420013800",&ast.AssignStmt/*struct*/{
                  Lhs: []ast.Expr /*Slice*/{
                    0: p.StrmapIndexExpr("0xc42000b860",&ast.IndexExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200118c0",&ast.Ident/*struct*/{
                        NamePos: 3109,
                        Name: "comments",
                        Obj: p.PtrmapObject("0xc420050e60"),
                      }/*struct*/),
                      Lbrack: 3117,
                      Index: p.StrmapIdent("0xc4200118e0",&ast.Ident/*struct*/{
                        NamePos: 3118,
                        Name: "i",
                        Obj: p.PtrmapObject("0xc420050eb0"),
                      }/*struct*/),
                      Rbrack: 3119,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  TokPos: 3121,
                  Tok: token.Token(42)/*=*/,
                  Rhs: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc420011940",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc420011900",&ast.Ident/*struct*/{
                        NamePos: 3123,
                        Name: "c",
                        Obj: p.PtrmapObject("0xc420050f00"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc420011920",&ast.Ident/*struct*/{
                        NamePos: 3125,
                        Name: "Text",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 3131,
            }/*struct*/),
          }/*struct*/),/* slice_item: 2*/3: p.StrmapAssignStmt("0xc420013900",&ast.AssignStmt/*struct*/{
            Lhs: []ast.Expr /*Slice*/{
              0: p.StrmapIdent("0xc420011960",&ast.Ident/*struct*/{
                NamePos: 3135,
                Name: "lines",
                Obj: p.StrmapObject("0xc420050fa0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "lines",
                  Decl: p.PtrmapAssignStmt("0xc420013900"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            TokPos: 3141,
            Tok: token.Token(47)/*:=*/,
            Rhs: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200138c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapIdent("0xc420011980",&ast.Ident/*struct*/{
                  NamePos: 3144,
                  Name: "make",
                }/*struct*/),
                Lparen: 3148,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapArrayType("0xc42000b8c0",&ast.ArrayType/*struct*/{
                    Lbrack: 3149,
                    Elt: p.StrmapIdent("0xc4200119a0",&ast.Ident/*struct*/{
                      NamePos: 3151,
                      Name: "string",
                    }/*struct*/),
                  }/*struct*/),/* slice_item: 0*/1: p.StrmapBasicLit("0xc4200119c0",&ast.BasicLit/*struct*/{
                    ValuePos: 3159,
                    Kind: token.Token(5)/*INT*/,
                    Value: "0",
                  }/*struct*/),/* slice_item: 1*/2: p.StrmapBasicLit("0xc420011a00",&ast.BasicLit/*struct*/{
                    ValuePos: 3162,
                    Kind: token.Token(5)/*INT*/,
                    Value: "10",
                  }/*struct*/),/* slice_item: 2*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 3164,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 3*/4: p.StrmapRangeStmt("0xc420051220",&ast.RangeStmt/*struct*/{
            For: 3207,
            Key: p.StrmapIdent("0xc420011a20",&ast.Ident/*struct*/{
              NamePos: 3211,
              Name: "_",
              Obj: p.StrmapObject("0xc420050ff0",&ast.Object/*struct*/{
                Kind: "var",
                Name: "_",
                Decl: p.StrmapAssignStmt("0xc420013940",&ast.AssignStmt/*struct*/{
                  Lhs: []ast.Expr /*Slice*/{
                    0: p.PtrmapIdent("0xc420011a20"),/* slice_item: 0*/1: p.StrmapIdent("0xc420011a40",&ast.Ident/*struct*/{
                      NamePos: 3214,
                      Name: "c",
                      Obj: p.StrmapObject("0xc420051040",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "c",
                        Decl: p.PtrmapAssignStmt("0xc420013940"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 1*/}/*slice*/,
                  TokPos: 3216,
                  Tok: token.Token(47)/*:=*/,
                  Rhs: []ast.Expr /*Slice*/{
                    0: p.StrmapUnaryExpr("0xc420011aa0",&ast.UnaryExpr/*struct*/{
                      OpPos: 3219,
                      Op: token.Token(79)/*range*/,
                      X: p.StrmapIdent("0xc420011a80",&ast.Ident/*struct*/{
                        NamePos: 3225,
                        Name: "comments",
                        Obj: p.PtrmapObject("0xc420050e60"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),
            Value: p.PtrmapIdent("0xc420011a40"),
            TokPos: 3216,
            Tok: token.Token(47)/*:=*/,
            X: p.PtrmapIdent("0xc420011a80"),
            Body: p.StrmapBlockStmt("0xc42000bce0",&ast.BlockStmt/*struct*/{
              Lbrace: 3234,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapSwitchStmt("0xc42000bc20",&ast.SwitchStmt/*struct*/{
                  Switch: 3322,
                  Tag: p.StrmapIndexExpr("0xc42000b9b0",&ast.IndexExpr/*struct*/{
                    X: p.StrmapIdent("0xc420011ae0",&ast.Ident/*struct*/{
                      NamePos: 3329,
                      Name: "c",
                      Obj: p.PtrmapObject("0xc420051040"),
                    }/*struct*/),
                    Lbrack: 3330,
                    Index: p.StrmapBasicLit("0xc420011b00",&ast.BasicLit/*struct*/{
                      ValuePos: 3331,
                      Kind: token.Token(5)/*INT*/,
                      Value: "1",
                    }/*struct*/),
                    Rbrack: 3332,
                  }/*struct*/),
                  Body: p.StrmapBlockStmt("0xc42000bbf0",&ast.BlockStmt/*struct*/{
                    Lbrace: 3334,
                    List: []ast.Stmt /*Slice*/{
                      0: p.StrmapCaseClause("0xc420013b80",&ast.CaseClause/*struct*/{
                        Case: 3338,
                        List: []ast.Expr /*Slice*/{
                          0: p.StrmapBasicLit("0xc420011b20",&ast.BasicLit/*struct*/{
                            ValuePos: 3343,
                            Kind: token.Token(8)/*CHAR*/,
                            Value: "'/'",
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                        Colon: 3346,
                        Body: []ast.Stmt /*Slice*/{
                          0: p.StrmapAssignStmt("0xc420013a00",&ast.AssignStmt/*struct*/{
                            Lhs: []ast.Expr /*Slice*/{
                              0: p.StrmapIdent("0xc420011b40",&ast.Ident/*struct*/{
                                NamePos: 3395,
                                Name: "c",
                                Obj: p.PtrmapObject("0xc420051040"),
                              }/*struct*/),/* slice_item: 0*/}/*slice*/,
                            TokPos: 3397,
                            Tok: token.Token(42)/*=*/,
                            Rhs: []ast.Expr /*Slice*/{
                              0: p.StrmapSliceExpr("0xc42005a480",&ast.SliceExpr/*struct*/{
                                X: p.StrmapIdent("0xc420011b60",&ast.Ident/*struct*/{
                                  NamePos: 3399,
                                  Name: "c",
                                  Obj: p.PtrmapObject("0xc420051040"),
                                }/*struct*/),
                                Lbrack: 3400,
                                Low: p.StrmapBasicLit("0xc420011b80",&ast.BasicLit/*struct*/{
                                  ValuePos: 3401,
                                  Kind: token.Token(5)/*INT*/,
                                  Value: "2",
                                }/*struct*/),
                                Slice3: false,
                                Rbrack: 3403,
                              }/*struct*/),/* slice_item: 0*/}/*slice*/,
                          }/*struct*/),/* slice_item: 0*/1: p.StrmapIfStmt("0xc420013b40",&ast.IfStmt/*struct*/{
                            If: 3461,
                            Cond: p.StrmapBinaryExpr("0xc42000bb00",&ast.BinaryExpr/*struct*/{
                              X: p.StrmapBinaryExpr("0xc42000ba70",&ast.BinaryExpr/*struct*/{
                                X: p.StrmapCallExpr("0xc420013ac0",&ast.CallExpr/*struct*/{
                                  Fun: p.StrmapIdent("0xc420011ba0",&ast.Ident/*struct*/{
                                    NamePos: 3464,
                                    Name: "len",
                                  }/*struct*/),
                                  Lparen: 3467,
                                  Args: []ast.Expr /*Slice*/{
                                    0: p.StrmapIdent("0xc420011bc0",&ast.Ident/*struct*/{
                                      NamePos: 3468,
                                      Name: "c",
                                      Obj: p.PtrmapObject("0xc420051040"),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Ellipsis: 0,
                                  Rparen: 3469,
                                }/*struct*/),
                                OpPos: 3471,
                                Op: token.Token(41)/*>*/,
                                Y: p.StrmapBasicLit("0xc420011be0",&ast.BasicLit/*struct*/{
                                  ValuePos: 3473,
                                  Kind: token.Token(5)/*INT*/,
                                  Value: "0",
                                }/*struct*/),
                              }/*struct*/),
                              OpPos: 3475,
                              Op: token.Token(34)/*&&*/,
                              Y: p.StrmapBinaryExpr("0xc42000bad0",&ast.BinaryExpr/*struct*/{
                                X: p.StrmapIndexExpr("0xc42000baa0",&ast.IndexExpr/*struct*/{
                                  X: p.StrmapIdent("0xc420011c00",&ast.Ident/*struct*/{
                                    NamePos: 3478,
                                    Name: "c",
                                    Obj: p.PtrmapObject("0xc420051040"),
                                  }/*struct*/),
                                  Lbrack: 3479,
                                  Index: p.StrmapBasicLit("0xc420011c20",&ast.BasicLit/*struct*/{
                                    ValuePos: 3480,
                                    Kind: token.Token(5)/*INT*/,
                                    Value: "0",
                                  }/*struct*/),
                                  Rbrack: 3481,
                                }/*struct*/),
                                OpPos: 3483,
                                Op: token.Token(39)/*==*/,
                                Y: p.StrmapBasicLit("0xc420011c40",&ast.BasicLit/*struct*/{
                                  ValuePos: 3486,
                                  Kind: token.Token(8)/*CHAR*/,
                                  Value: "' '",
                                }/*struct*/),
                              }/*struct*/),
                            }/*struct*/),
                            Body: p.StrmapBlockStmt("0xc42000bb60",&ast.BlockStmt/*struct*/{
                              Lbrace: 3490,
                              List: []ast.Stmt /*Slice*/{
                                0: p.StrmapAssignStmt("0xc420013b00",&ast.AssignStmt/*struct*/{
                                  Lhs: []ast.Expr /*Slice*/{
                                    0: p.StrmapIdent("0xc420011c60",&ast.Ident/*struct*/{
                                      NamePos: 3496,
                                      Name: "c",
                                      Obj: p.PtrmapObject("0xc420051040"),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  TokPos: 3498,
                                  Tok: token.Token(42)/*=*/,
                                  Rhs: []ast.Expr /*Slice*/{
                                    0: p.StrmapSliceExpr("0xc42005a4e0",&ast.SliceExpr/*struct*/{
                                      X: p.StrmapIdent("0xc420011c80",&ast.Ident/*struct*/{
                                        NamePos: 3500,
                                        Name: "c",
                                        Obj: p.PtrmapObject("0xc420051040"),
                                      }/*struct*/),
                                      Lbrack: 3501,
                                      Low: p.StrmapBasicLit("0xc420011ca0",&ast.BasicLit/*struct*/{
                                        ValuePos: 3502,
                                        Kind: token.Token(5)/*INT*/,
                                        Value: "1",
                                      }/*struct*/),
                                      Slice3: false,
                                      Rbrack: 3504,
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                }/*struct*/),/* slice_item: 0*/}/*slice*/,
                              Rbrace: 3509,
                            }/*struct*/),
                          }/*struct*/),/* slice_item: 1*/}/*slice*/,
                      }/*struct*/),/* slice_item: 0*/1: p.StrmapCaseClause("0xc420013c40",&ast.CaseClause/*struct*/{
                        Case: 3513,
                        List: []ast.Expr /*Slice*/{
                          0: p.StrmapBasicLit("0xc420011ce0",&ast.BasicLit/*struct*/{
                            ValuePos: 3518,
                            Kind: token.Token(8)/*CHAR*/,
                            Value: "'*'",
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                        Colon: 3521,
                        Body: []ast.Stmt /*Slice*/{
                          0: p.StrmapAssignStmt("0xc420013c00",&ast.AssignStmt/*struct*/{
                            Lhs: []ast.Expr /*Slice*/{
                              0: p.StrmapIdent("0xc420011d20",&ast.Ident/*struct*/{
                                NamePos: 3549,
                                Name: "c",
                                Obj: p.PtrmapObject("0xc420051040"),
                              }/*struct*/),/* slice_item: 0*/}/*slice*/,
                            TokPos: 3551,
                            Tok: token.Token(42)/*=*/,
                            Rhs: []ast.Expr /*Slice*/{
                              0: p.StrmapSliceExpr("0xc42005a540",&ast.SliceExpr/*struct*/{
                                X: p.StrmapIdent("0xc420011d40",&ast.Ident/*struct*/{
                                  NamePos: 3553,
                                  Name: "c",
                                  Obj: p.PtrmapObject("0xc420051040"),
                                }/*struct*/),
                                Lbrack: 3554,
                                Low: p.StrmapBasicLit("0xc420011d60",&ast.BasicLit/*struct*/{
                                  ValuePos: 3555,
                                  Kind: token.Token(5)/*INT*/,
                                  Value: "2",
                                }/*struct*/),
                                High: p.StrmapBinaryExpr("0xc42000bbc0",&ast.BinaryExpr/*struct*/{
                                  X: p.StrmapCallExpr("0xc420013bc0",&ast.CallExpr/*struct*/{
                                    Fun: p.StrmapIdent("0xc420011d80",&ast.Ident/*struct*/{
                                      NamePos: 3559,
                                      Name: "len",
                                    }/*struct*/),
                                    Lparen: 3562,
                                    Args: []ast.Expr /*Slice*/{
                                      0: p.StrmapIdent("0xc420011da0",&ast.Ident/*struct*/{
                                        NamePos: 3563,
                                        Name: "c",
                                        Obj: p.PtrmapObject("0xc420051040"),
                                      }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                    Ellipsis: 0,
                                    Rparen: 3564,
                                  }/*struct*/),
                                  OpPos: 3565,
                                  Op: token.Token(13)/*-*/,
                                  Y: p.StrmapBasicLit("0xc420011dc0",&ast.BasicLit/*struct*/{
                                    ValuePos: 3566,
                                    Kind: token.Token(5)/*INT*/,
                                    Value: "2",
                                  }/*struct*/),
                                }/*struct*/),
                                Slice3: false,
                                Rbrack: 3567,
                              }/*struct*/),/* slice_item: 0*/}/*slice*/,
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      }/*struct*/),/* slice_item: 1*/}/*slice*/,
                    Rbrace: 3571,
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapAssignStmt("0xc420013cc0",&ast.AssignStmt/*struct*/{
                  Lhs: []ast.Expr /*Slice*/{
                    0: p.StrmapIdent("0xc420011e20",&ast.Ident/*struct*/{
                      NamePos: 3600,
                      Name: "cl",
                      Obj: p.StrmapObject("0xc420051090",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "cl",
                        Decl: p.PtrmapAssignStmt("0xc420013cc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  TokPos: 3603,
                  Tok: token.Token(47)/*:=*/,
                  Rhs: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc420013c80",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc420011e80",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIdent("0xc420011e40",&ast.Ident/*struct*/{
                          NamePos: 3606,
                          Name: "strings",
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc420011e60",&ast.Ident/*struct*/{
                          NamePos: 3614,
                          Name: "Split",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 3619,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapIdent("0xc420011ea0",&ast.Ident/*struct*/{
                          NamePos: 3620,
                          Name: "c",
                          Obj: p.PtrmapObject("0xc420051040"),
                        }/*struct*/),/* slice_item: 0*/1: p.StrmapBasicLit("0xc420011ec0",&ast.BasicLit/*struct*/{
                          ValuePos: 3623,
                          Kind: token.Token(9)/*STRING*/,
                          Value: "\"\\n\"",
                        }/*struct*/),/* slice_item: 1*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 3627,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 1*/2: p.StrmapRangeStmt("0xc4200511d0",&ast.RangeStmt/*struct*/{
                  For: 3700,
                  Key: p.StrmapIdent("0xc420011f20",&ast.Ident/*struct*/{
                    NamePos: 3704,
                    Name: "_",
                    Obj: p.StrmapObject("0xc420051130",&ast.Object/*struct*/{
                      Kind: "var",
                      Name: "_",
                      Decl: p.StrmapAssignStmt("0xc420013d00",&ast.AssignStmt/*struct*/{
                        Lhs: []ast.Expr /*Slice*/{
                          0: p.PtrmapIdent("0xc420011f20"),/* slice_item: 0*/1: p.StrmapIdent("0xc420011f40",&ast.Ident/*struct*/{
                            NamePos: 3707,
                            Name: "l",
                            Obj: p.StrmapObject("0xc420051180",&ast.Object/*struct*/{
                              Kind: "var",
                              Name: "l",
                              Decl: p.PtrmapAssignStmt("0xc420013d00"),
                            }/*struct*/),
                          }/*struct*/),/* slice_item: 1*/}/*slice*/,
                        TokPos: 3709,
                        Tok: token.Token(47)/*:=*/,
                        Rhs: []ast.Expr /*Slice*/{
                          0: p.StrmapUnaryExpr("0xc420011fa0",&ast.UnaryExpr/*struct*/{
                            OpPos: 3712,
                            Op: token.Token(79)/*range*/,
                            X: p.StrmapIdent("0xc420011f80",&ast.Ident/*struct*/{
                              NamePos: 3718,
                              Name: "cl",
                              Obj: p.PtrmapObject("0xc420051090"),
                            }/*struct*/),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                  Value: p.PtrmapIdent("0xc420011f40"),
                  TokPos: 3709,
                  Tok: token.Token(47)/*:=*/,
                  X: p.PtrmapIdent("0xc420011f80"),
                  Body: p.StrmapBlockStmt("0xc42000bcb0",&ast.BlockStmt/*struct*/{
                    Lbrace: 3721,
                    List: []ast.Stmt /*Slice*/{
                      0: p.StrmapAssignStmt("0xc420013dc0",&ast.AssignStmt/*struct*/{
                        Lhs: []ast.Expr /*Slice*/{
                          0: p.StrmapIdent("0xc420011fc0",&ast.Ident/*struct*/{
                            NamePos: 3726,
                            Name: "lines",
                            Obj: p.PtrmapObject("0xc420050fa0"),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                        TokPos: 3732,
                        Tok: token.Token(42)/*=*/,
                        Rhs: []ast.Expr /*Slice*/{
                          0: p.StrmapCallExpr("0xc420013d80",&ast.CallExpr/*struct*/{
                            Fun: p.StrmapIdent("0xc420011fe0",&ast.Ident/*struct*/{
                              NamePos: 3734,
                              Name: "append",
                            }/*struct*/),
                            Lparen: 3740,
                            Args: []ast.Expr /*Slice*/{
                              0: p.StrmapIdent("0xc4200b0000",&ast.Ident/*struct*/{
                                NamePos: 3741,
                                Name: "lines",
                                Obj: p.PtrmapObject("0xc420050fa0"),
                              }/*struct*/),/* slice_item: 0*/1: p.StrmapCallExpr("0xc420013d40",&ast.CallExpr/*struct*/{
                                Fun: p.StrmapIdent("0xc4200b0040",&ast.Ident/*struct*/{
                                  NamePos: 3748,
                                  Name: "stripTrailingWhitespace",
                                  Obj: p.PtrmapObject("0xc420050dc0"),
                                }/*struct*/),
                                Lparen: 3771,
                                Args: []ast.Expr /*Slice*/{
                                  0: p.StrmapIdent("0xc4200b0060",&ast.Ident/*struct*/{
                                    NamePos: 3772,
                                    Name: "l",
                                    Obj: p.PtrmapObject("0xc420051180"),
                                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                Ellipsis: 0,
                                Rparen: 3773,
                              }/*struct*/),/* slice_item: 1*/}/*slice*/,
                            Ellipsis: 0,
                            Rparen: 3774,
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      }/*struct*/),/* slice_item: 0*/}/*slice*/,
                    Rbrace: 3778,
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Rbrace: 3781,
            }/*struct*/),
          }/*struct*/),/* slice_item: 4*/5: p.StrmapAssignStmt("0xc420013e40",&ast.AssignStmt/*struct*/{
            Lhs: []ast.Expr /*Slice*/{
              0: p.StrmapIdent("0xc4200b00a0",&ast.Ident/*struct*/{
                NamePos: 3882,
                Name: "n",
                Obj: p.StrmapObject("0xc420051270",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "n",
                  Decl: p.PtrmapAssignStmt("0xc420013e40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            TokPos: 3884,
            Tok: token.Token(47)/*:=*/,
            Rhs: []ast.Expr /*Slice*/{
              0: p.StrmapBasicLit("0xc4200b00c0",&ast.BasicLit/*struct*/{
                ValuePos: 3887,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 5*/6: p.StrmapRangeStmt("0xc420051360",&ast.RangeStmt/*struct*/{
            For: 3890,
            Key: p.StrmapIdent("0xc4200b00e0",&ast.Ident/*struct*/{
              NamePos: 3894,
              Name: "_",
              Obj: p.StrmapObject("0xc4200512c0",&ast.Object/*struct*/{
                Kind: "var",
                Name: "_",
                Decl: p.StrmapAssignStmt("0xc420013e80",&ast.AssignStmt/*struct*/{
                  Lhs: []ast.Expr /*Slice*/{
                    0: p.PtrmapIdent("0xc4200b00e0"),/* slice_item: 0*/1: p.StrmapIdent("0xc4200b0100",&ast.Ident/*struct*/{
                      NamePos: 3897,
                      Name: "line",
                      Obj: p.StrmapObject("0xc420051310",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "line",
                        Decl: p.PtrmapAssignStmt("0xc420013e80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 1*/}/*slice*/,
                  TokPos: 3902,
                  Tok: token.Token(47)/*:=*/,
                  Rhs: []ast.Expr /*Slice*/{
                    0: p.StrmapUnaryExpr("0xc4200b0160",&ast.UnaryExpr/*struct*/{
                      OpPos: 3905,
                      Op: token.Token(79)/*range*/,
                      X: p.StrmapIdent("0xc4200b0140",&ast.Ident/*struct*/{
                        NamePos: 3911,
                        Name: "lines",
                        Obj: p.PtrmapObject("0xc420050fa0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),
            Value: p.PtrmapIdent("0xc4200b0100"),
            TokPos: 3902,
            Tok: token.Token(47)/*:=*/,
            X: p.PtrmapIdent("0xc4200b0140"),
            Body: p.StrmapBlockStmt("0xc4200b4000",&ast.BlockStmt/*struct*/{
              Lbrace: 3917,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapIfStmt("0xc420013f00",&ast.IfStmt/*struct*/{
                  If: 3921,
                  Cond: p.StrmapBinaryExpr("0xc42000bf20",&ast.BinaryExpr/*struct*/{
                    X: p.StrmapBinaryExpr("0xc42000be00",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200b0180",&ast.Ident/*struct*/{
                        NamePos: 3924,
                        Name: "line",
                        Obj: p.PtrmapObject("0xc420051310"),
                      }/*struct*/),
                      OpPos: 3929,
                      Op: token.Token(44)/*!=*/,
                      Y: p.StrmapBasicLit("0xc4200b01a0",&ast.BasicLit/*struct*/{
                        ValuePos: 3932,
                        Kind: token.Token(9)/*STRING*/,
                        Value: "\"\"",
                      }/*struct*/),
                    }/*struct*/),
                    OpPos: 3935,
                    Op: token.Token(35)/*||*/,
                    Y: p.StrmapBinaryExpr("0xc42000bef0",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapBinaryExpr("0xc42000be30",&ast.BinaryExpr/*struct*/{
                        X: p.StrmapIdent("0xc4200b01c0",&ast.Ident/*struct*/{
                          NamePos: 3938,
                          Name: "n",
                          Obj: p.PtrmapObject("0xc420051270"),
                        }/*struct*/),
                        OpPos: 3940,
                        Op: token.Token(41)/*>*/,
                        Y: p.StrmapBasicLit("0xc4200b01e0",&ast.BasicLit/*struct*/{
                          ValuePos: 3942,
                          Kind: token.Token(5)/*INT*/,
                          Value: "0",
                        }/*struct*/),
                      }/*struct*/),
                      OpPos: 3944,
                      Op: token.Token(34)/*&&*/,
                      Y: p.StrmapBinaryExpr("0xc42000bec0",&ast.BinaryExpr/*struct*/{
                        X: p.StrmapIndexExpr("0xc42000be90",&ast.IndexExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200b0200",&ast.Ident/*struct*/{
                            NamePos: 3947,
                            Name: "lines",
                            Obj: p.PtrmapObject("0xc420050fa0"),
                          }/*struct*/),
                          Lbrack: 3952,
                          Index: p.StrmapBinaryExpr("0xc42000be60",&ast.BinaryExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200b0220",&ast.Ident/*struct*/{
                              NamePos: 3953,
                              Name: "n",
                              Obj: p.PtrmapObject("0xc420051270"),
                            }/*struct*/),
                            OpPos: 3954,
                            Op: token.Token(13)/*-*/,
                            Y: p.StrmapBasicLit("0xc4200b0240",&ast.BasicLit/*struct*/{
                              ValuePos: 3955,
                              Kind: token.Token(5)/*INT*/,
                              Value: "1",
                            }/*struct*/),
                          }/*struct*/),
                          Rbrack: 3956,
                        }/*struct*/),
                        OpPos: 3958,
                        Op: token.Token(44)/*!=*/,
                        Y: p.StrmapBasicLit("0xc4200b0260",&ast.BasicLit/*struct*/{
                          ValuePos: 3961,
                          Kind: token.Token(9)/*STRING*/,
                          Value: "\"\"",
                        }/*struct*/),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                  Body: p.StrmapBlockStmt("0xc42000bfb0",&ast.BlockStmt/*struct*/{
                    Lbrace: 3964,
                    List: []ast.Stmt /*Slice*/{
                      0: p.StrmapAssignStmt("0xc420013ec0",&ast.AssignStmt/*struct*/{
                        Lhs: []ast.Expr /*Slice*/{
                          0: p.StrmapIndexExpr("0xc42000bf80",&ast.IndexExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200b0280",&ast.Ident/*struct*/{
                              NamePos: 3969,
                              Name: "lines",
                              Obj: p.PtrmapObject("0xc420050fa0"),
                            }/*struct*/),
                            Lbrack: 3974,
                            Index: p.StrmapIdent("0xc4200b02a0",&ast.Ident/*struct*/{
                              NamePos: 3975,
                              Name: "n",
                              Obj: p.PtrmapObject("0xc420051270"),
                            }/*struct*/),
                            Rbrack: 3976,
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                        TokPos: 3978,
                        Tok: token.Token(42)/*=*/,
                        Rhs: []ast.Expr /*Slice*/{
                          0: p.StrmapIdent("0xc4200b02c0",&ast.Ident/*struct*/{
                            NamePos: 3980,
                            Name: "line",
                            Obj: p.PtrmapObject("0xc420051310"),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      }/*struct*/),/* slice_item: 0*/1: p.StrmapIncDecStmt("0xc4200b0300",&ast.IncDecStmt/*struct*/{
                        X: p.StrmapIdent("0xc4200b02e0",&ast.Ident/*struct*/{
                          NamePos: 3988,
                          Name: "n",
                          Obj: p.PtrmapObject("0xc420051270"),
                        }/*struct*/),
                        TokPos: 3989,
                        Tok: token.Token(37)/*++*/,
                      }/*struct*/),/* slice_item: 1*/}/*slice*/,
                    Rbrace: 3994,
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 3997,
            }/*struct*/),
          }/*struct*/),/* slice_item: 6*/7: p.StrmapAssignStmt("0xc420013f40",&ast.AssignStmt/*struct*/{
            Lhs: []ast.Expr /*Slice*/{
              0: p.StrmapIdent("0xc4200b0340",&ast.Ident/*struct*/{
                NamePos: 4000,
                Name: "lines",
                Obj: p.PtrmapObject("0xc420050fa0"),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            TokPos: 4006,
            Tok: token.Token(42)/*=*/,
            Rhs: []ast.Expr /*Slice*/{
              0: p.StrmapSliceExpr("0xc42005a5a0",&ast.SliceExpr/*struct*/{
                X: p.StrmapIdent("0xc4200b0360",&ast.Ident/*struct*/{
                  NamePos: 4008,
                  Name: "lines",
                  Obj: p.PtrmapObject("0xc420050fa0"),
                }/*struct*/),
                Lbrack: 4013,
                Low: p.StrmapBasicLit("0xc4200b0380",&ast.BasicLit/*struct*/{
                  ValuePos: 4014,
                  Kind: token.Token(5)/*INT*/,
                  Value: "0",
                }/*struct*/),
                High: p.StrmapIdent("0xc4200b03a0",&ast.Ident/*struct*/{
                  NamePos: 4016,
                  Name: "n",
                  Obj: p.PtrmapObject("0xc420051270"),
                }/*struct*/),
                Slice3: false,
                Rbrack: 4017,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 7*/8: p.StrmapIfStmt("0xc4200b6040",&ast.IfStmt/*struct*/{
            If: 4079,
            Cond: p.StrmapBinaryExpr("0xc4200b4120",&ast.BinaryExpr/*struct*/{
              X: p.StrmapBinaryExpr("0xc4200b4060",&ast.BinaryExpr/*struct*/{
                X: p.StrmapIdent("0xc4200b03c0",&ast.Ident/*struct*/{
                  NamePos: 4082,
                  Name: "n",
                  Obj: p.PtrmapObject("0xc420051270"),
                }/*struct*/),
                OpPos: 4084,
                Op: token.Token(41)/*>*/,
                Y: p.StrmapBasicLit("0xc4200b03e0",&ast.BasicLit/*struct*/{
                  ValuePos: 4086,
                  Kind: token.Token(5)/*INT*/,
                  Value: "0",
                }/*struct*/),
              }/*struct*/),
              OpPos: 4088,
              Op: token.Token(34)/*&&*/,
              Y: p.StrmapBinaryExpr("0xc4200b40f0",&ast.BinaryExpr/*struct*/{
                X: p.StrmapIndexExpr("0xc4200b40c0",&ast.IndexExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200b0400",&ast.Ident/*struct*/{
                    NamePos: 4091,
                    Name: "lines",
                    Obj: p.PtrmapObject("0xc420050fa0"),
                  }/*struct*/),
                  Lbrack: 4096,
                  Index: p.StrmapBinaryExpr("0xc4200b4090",&ast.BinaryExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200b0420",&ast.Ident/*struct*/{
                      NamePos: 4097,
                      Name: "n",
                      Obj: p.PtrmapObject("0xc420051270"),
                    }/*struct*/),
                    OpPos: 4098,
                    Op: token.Token(13)/*-*/,
                    Y: p.StrmapBasicLit("0xc4200b0440",&ast.BasicLit/*struct*/{
                      ValuePos: 4099,
                      Kind: token.Token(5)/*INT*/,
                      Value: "1",
                    }/*struct*/),
                  }/*struct*/),
                  Rbrack: 4100,
                }/*struct*/),
                OpPos: 4102,
                Op: token.Token(44)/*!=*/,
                Y: p.StrmapBasicLit("0xc4200b0460",&ast.BasicLit/*struct*/{
                  ValuePos: 4105,
                  Kind: token.Token(9)/*STRING*/,
                  Value: "\"\"",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200b4180",&ast.BlockStmt/*struct*/{
              Lbrace: 4108,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapAssignStmt("0xc4200b6000",&ast.AssignStmt/*struct*/{
                  Lhs: []ast.Expr /*Slice*/{
                    0: p.StrmapIdent("0xc4200b0480",&ast.Ident/*struct*/{
                      NamePos: 4112,
                      Name: "lines",
                      Obj: p.PtrmapObject("0xc420050fa0"),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  TokPos: 4118,
                  Tok: token.Token(42)/*=*/,
                  Rhs: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc420013fc0",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapIdent("0xc4200b04a0",&ast.Ident/*struct*/{
                        NamePos: 4120,
                        Name: "append",
                      }/*struct*/),
                      Lparen: 4126,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapIdent("0xc4200b04c0",&ast.Ident/*struct*/{
                          NamePos: 4127,
                          Name: "lines",
                          Obj: p.PtrmapObject("0xc420050fa0"),
                        }/*struct*/),/* slice_item: 0*/1: p.StrmapBasicLit("0xc4200b04e0",&ast.BasicLit/*struct*/{
                          ValuePos: 4134,
                          Kind: token.Token(9)/*STRING*/,
                          Value: "\"\"",
                        }/*struct*/),/* slice_item: 1*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 4136,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 4139,
            }/*struct*/),
          }/*struct*/),/* slice_item: 8*/9: p.StrmapReturnStmt("0xc4200b05e0",&ast.ReturnStmt/*struct*/{
            Return: 4143,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200b6080",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200b0560",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200b0520",&ast.Ident/*struct*/{
                    NamePos: 4150,
                    Name: "strings",
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200b0540",&ast.Ident/*struct*/{
                    NamePos: 4158,
                    Name: "Join",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 4162,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapIdent("0xc4200b0580",&ast.Ident/*struct*/{
                    NamePos: 4163,
                    Name: "lines",
                    Obj: p.PtrmapObject("0xc420050fa0"),
                  }/*struct*/),/* slice_item: 0*/1: p.StrmapBasicLit("0xc4200b05a0",&ast.BasicLit/*struct*/{
                    ValuePos: 4170,
                    Kind: token.Token(9)/*STRING*/,
                    Value: "\"\\n\"",
                  }/*struct*/),/* slice_item: 1*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 4174,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 9*/}/*slice*/,
        Rbrace: 4176,
      }/*struct*/),
    }/*struct*/),/* slice_item: 13*/14: p.StrmapGenDecl("0xc4200b62c0",&ast.GenDecl/*struct*/{
      TokPos: 4445,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200b4210",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200b0660",&ast.Ident/*struct*/{
            NamePos: 4450,
            Name: "Field",
            Obj: p.StrmapObject("0xc420051450",&ast.Object/*struct*/{
              Kind: "type",
              Name: "Field",
              Decl: p.PtrmapTypeSpec("0xc4200b4210"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200b08c0",&ast.StructType/*struct*/{
            Struct: 4456,
            Fields: p.StrmapFieldList("0xc4200b42a0",&ast.FieldList/*struct*/{
              Opening: 4463,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b6100",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200b0680",&ast.Ident/*struct*/{
                      NamePos: 4466,
                      Name: "Doc",
                      Obj: p.StrmapObject("0xc4200514a0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Doc",
                        Decl: p.PtrmapField("0xc4200b6100"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200b06c0",&ast.StarExpr/*struct*/{
                    Star: 4474,
                    X: p.StrmapIdent("0xc4200b06a0",&ast.Ident/*struct*/{
                      NamePos: 4475,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b6180",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200b06e0",&ast.Ident/*struct*/{
                      NamePos: 4525,
                      Name: "Names",
                      Obj: p.StrmapObject("0xc4200514f0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Names",
                        Decl: p.PtrmapField("0xc4200b6180"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200b4270",&ast.ArrayType/*struct*/{
                    Lbrack: 4533,
                    Elt: p.StrmapStarExpr("0xc4200b0720",&ast.StarExpr/*struct*/{
                      Star: 4535,
                      X: p.StrmapIdent("0xc4200b0700",&ast.Ident/*struct*/{
                        NamePos: 4536,
                        Name: "Ident",
                        Obj: p.StrmapObject("0xc420051c70",&ast.Object/*struct*/{
                          Kind: "type",
                          Name: "Ident",
                          Decl: p.StrmapTypeSpec("0xc4200b52f0",&ast.TypeSpec/*struct*/{
                            Name: p.StrmapIdent("0xc4200b1f60",&ast.Ident/*struct*/{
                              NamePos: 6431,
                              Name: "Ident",
                              Obj: p.PtrmapObject("0xc420051c70"),
                            }/*struct*/),
                            Type: p.StrmapStructType("0xc4200ba120",&ast.StructType/*struct*/{
                              Struct: 6437,
                              Fields: p.StrmapFieldList("0xc4200b5320",&ast.FieldList/*struct*/{
                                Opening: 6444,
                                List: []*ast.Field /*Slice*/{
                                  0: p.StrmapField("0xc4200b6f80",&ast.Field/*struct*/{
                                    Names: []*ast.Ident /*Slice*/{
                                      0: p.StrmapIdent("0xc4200b1f80",&ast.Ident/*struct*/{
                                        NamePos: 6448,
                                        Name: "NamePos",
                                        Obj: p.StrmapObject("0xc420051cc0",&ast.Object/*struct*/{
                                          Kind: "var",
                                          Name: "NamePos",
                                          Decl: p.PtrmapField("0xc4200b6f80"),
                                        }/*struct*/),
                                      }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                    Type: p.StrmapSelectorExpr("0xc4200b1fe0",&ast.SelectorExpr/*struct*/{
                                      X: p.StrmapIdent("0xc4200b1fa0",&ast.Ident/*struct*/{
                                        NamePos: 6456,
                                        Name: "token",
                                      }/*struct*/),
                                      Sel: p.StrmapIdent("0xc4200b1fc0",&ast.Ident/*struct*/{
                                        NamePos: 6462,
                                        Name: "Pos",
                                      }/*struct*/),
                                    }/*struct*/),
                                  }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b6fc0",&ast.Field/*struct*/{
                                    Names: []*ast.Ident /*Slice*/{
                                      0: p.StrmapIdent("0xc4200ba020",&ast.Ident/*struct*/{
                                        NamePos: 6491,
                                        Name: "Name",
                                        Obj: p.StrmapObject("0xc420051d10",&ast.Object/*struct*/{
                                          Kind: "var",
                                          Name: "Name",
                                          Decl: p.PtrmapField("0xc4200b6fc0"),
                                        }/*struct*/),
                                      }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                    Type: p.StrmapIdent("0xc4200ba040",&ast.Ident/*struct*/{
                                      NamePos: 6499,
                                      Name: "string",
                                    }/*struct*/),
                                  }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7000",&ast.Field/*struct*/{
                                    Names: []*ast.Ident /*Slice*/{
                                      0: p.StrmapIdent("0xc4200ba080",&ast.Ident/*struct*/{
                                        NamePos: 6530,
                                        Name: "Obj",
                                        Obj: p.StrmapObject("0xc420051d60",&ast.Object/*struct*/{
                                          Kind: "var",
                                          Name: "Obj",
                                          Decl: p.PtrmapField("0xc4200b7000"),
                                        }/*struct*/),
                                      }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                    Type: p.StrmapStarExpr("0xc4200ba0c0",&ast.StarExpr/*struct*/{
                                      Star: 6538,
                                      X: p.StrmapIdent("0xc4200ba0a0",&ast.Ident/*struct*/{
                                        NamePos: 6539,
                                        Name: "Object",
                                      }/*struct*/),
                                    }/*struct*/),
                                  }/*struct*/),/* slice_item: 2*/}/*slice*/,
                                Closing: 6575,
                              }/*struct*/),
                              Incomplete: false,
                            }/*struct*/),
                          }/*struct*/),
                        }/*struct*/),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b61c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200b0740",&ast.Ident/*struct*/{
                      NamePos: 4607,
                      Name: "Type",
                      Obj: p.StrmapObject("0xc420051540",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Type",
                        Decl: p.PtrmapField("0xc4200b61c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200b0760",&ast.Ident/*struct*/{
                    NamePos: 4615,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200b6200",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200b07c0",&ast.Ident/*struct*/{
                      NamePos: 4661,
                      Name: "Tag",
                      Obj: p.StrmapObject("0xc420051590",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Tag",
                        Decl: p.PtrmapField("0xc4200b6200"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200b0800",&ast.StarExpr/*struct*/{
                    Star: 4669,
                    X: p.StrmapIdent("0xc4200b07e0",&ast.Ident/*struct*/{
                      NamePos: 4670,
                      Name: "BasicLit",
                      Obj: p.StrmapObject("0xc420051ea0",&ast.Object/*struct*/{
                        Kind: "type",
                        Name: "BasicLit",
                        Decl: p.StrmapTypeSpec("0xc4200b53b0",&ast.TypeSpec/*struct*/{
                          Name: p.StrmapIdent("0xc4200ba280",&ast.Ident/*struct*/{
                            NamePos: 6890,
                            Name: "BasicLit",
                            Obj: p.PtrmapObject("0xc420051ea0"),
                          }/*struct*/),
                          Type: p.StrmapStructType("0xc4200ba420",&ast.StructType/*struct*/{
                            Struct: 6899,
                            Fields: p.StrmapFieldList("0xc4200b53e0",&ast.FieldList/*struct*/{
                              Opening: 6906,
                              List: []*ast.Field /*Slice*/{
                                0: p.StrmapField("0xc4200b7200",&ast.Field/*struct*/{
                                  Names: []*ast.Ident /*Slice*/{
                                    0: p.StrmapIdent("0xc4200ba2a0",&ast.Ident/*struct*/{
                                      NamePos: 6910,
                                      Name: "ValuePos",
                                      Obj: p.StrmapObject("0xc420051ef0",&ast.Object/*struct*/{
                                        Kind: "var",
                                        Name: "ValuePos",
                                        Decl: p.PtrmapField("0xc4200b7200"),
                                      }/*struct*/),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Type: p.StrmapSelectorExpr("0xc4200ba300",&ast.SelectorExpr/*struct*/{
                                    X: p.StrmapIdent("0xc4200ba2c0",&ast.Ident/*struct*/{
                                      NamePos: 6919,
                                      Name: "token",
                                    }/*struct*/),
                                    Sel: p.StrmapIdent("0xc4200ba2e0",&ast.Ident/*struct*/{
                                      NamePos: 6925,
                                      Name: "Pos",
                                    }/*struct*/),
                                  }/*struct*/),
                                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7240",&ast.Field/*struct*/{
                                  Names: []*ast.Ident /*Slice*/{
                                    0: p.StrmapIdent("0xc4200ba340",&ast.Ident/*struct*/{
                                      NamePos: 6953,
                                      Name: "Kind",
                                      Obj: p.StrmapObject("0xc420051f90",&ast.Object/*struct*/{
                                        Kind: "var",
                                        Name: "Kind",
                                        Decl: p.PtrmapField("0xc4200b7240"),
                                      }/*struct*/),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Type: p.StrmapSelectorExpr("0xc4200ba3a0",&ast.SelectorExpr/*struct*/{
                                    X: p.StrmapIdent("0xc4200ba360",&ast.Ident/*struct*/{
                                      NamePos: 6962,
                                      Name: "token",
                                    }/*struct*/),
                                    Sel: p.StrmapIdent("0xc4200ba380",&ast.Ident/*struct*/{
                                      NamePos: 6968,
                                      Name: "Token",
                                    }/*struct*/),
                                  }/*struct*/),
                                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7280",&ast.Field/*struct*/{
                                  Names: []*ast.Ident /*Slice*/{
                                    0: p.StrmapIdent("0xc4200ba3c0",&ast.Ident/*struct*/{
                                      NamePos: 7043,
                                      Name: "Value",
                                      Obj: p.StrmapObject("0xc4200c0000",&ast.Object/*struct*/{
                                        Kind: "var",
                                        Name: "Value",
                                        Decl: p.PtrmapField("0xc4200b7280"),
                                      }/*struct*/),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Type: p.StrmapIdent("0xc4200ba3e0",&ast.Ident/*struct*/{
                                    NamePos: 7052,
                                    Name: "string",
                                  }/*struct*/),
                                }/*struct*/),/* slice_item: 2*/}/*slice*/,
                              Closing: 7148,
                            }/*struct*/),
                            Incomplete: false,
                          }/*struct*/),
                        }/*struct*/),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200b6240",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200b0840",&ast.Ident/*struct*/{
                      NamePos: 4705,
                      Name: "Comment",
                      Obj: p.StrmapObject("0xc4200515e0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Comment",
                        Decl: p.PtrmapField("0xc4200b6240"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200b0880",&ast.StarExpr/*struct*/{
                    Star: 4713,
                    X: p.StrmapIdent("0xc4200b0860",&ast.Ident/*struct*/{
                      NamePos: 4714,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/}/*slice*/,
              Closing: 4752,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 14*/15: p.StrmapFuncDecl("0xc4200b44e0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200b4300",&ast.FieldList/*struct*/{
        Opening: 4760,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200b6300",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200b08e0",&ast.Ident/*struct*/{
                NamePos: 4761,
                Name: "f",
                Obj: p.StrmapObject("0xc420051630",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "f",
                  Decl: p.PtrmapField("0xc4200b6300"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200b0920",&ast.StarExpr/*struct*/{
              Star: 4763,
              X: p.StrmapIdent("0xc4200b0900",&ast.Ident/*struct*/{
                NamePos: 4764,
                Name: "Field",
                Obj: p.PtrmapObject("0xc420051450"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 4769,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200b0940",&ast.Ident/*struct*/{
        NamePos: 4771,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200b0c20",&ast.FuncType/*struct*/{
        Func: 4755,
        Params: p.StrmapFieldList("0xc4200b4330",&ast.FieldList/*struct*/{
          Opening: 4774,
          Closing: 4775,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200b4360",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200b6340",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200b09a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200b0960",&ast.Ident/*struct*/{
                  NamePos: 4777,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200b0980",&ast.Ident/*struct*/{
                  NamePos: 4783,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200b44b0",&ast.BlockStmt/*struct*/{
        Lbrace: 4787,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200b6400",&ast.IfStmt/*struct*/{
            If: 4790,
            Cond: p.StrmapBinaryExpr("0xc4200b43f0",&ast.BinaryExpr/*struct*/{
              X: p.StrmapCallExpr("0xc4200b6380",&ast.CallExpr/*struct*/{
                Fun: p.StrmapIdent("0xc4200b09c0",&ast.Ident/*struct*/{
                  NamePos: 4793,
                  Name: "len",
                }/*struct*/),
                Lparen: 4796,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapSelectorExpr("0xc4200b0a20",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200b09e0",&ast.Ident/*struct*/{
                      NamePos: 4797,
                      Name: "f",
                      Obj: p.PtrmapObject("0xc420051630"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200b0a00",&ast.Ident/*struct*/{
                      NamePos: 4799,
                      Name: "Names",
                    }/*struct*/),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 4804,
              }/*struct*/),
              OpPos: 4806,
              Op: token.Token(41)/*>*/,
              Y: p.StrmapBasicLit("0xc4200b0a40",&ast.BasicLit/*struct*/{
                ValuePos: 4808,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200b4480",&ast.BlockStmt/*struct*/{
              Lbrace: 4810,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200b0b20",&ast.ReturnStmt/*struct*/{
                  Return: 4814,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200b63c0",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200b0b00",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIndexExpr("0xc4200b4450",&ast.IndexExpr/*struct*/{
                          X: p.StrmapSelectorExpr("0xc4200b0aa0",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200b0a60",&ast.Ident/*struct*/{
                              NamePos: 4821,
                              Name: "f",
                              Obj: p.PtrmapObject("0xc420051630"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200b0a80",&ast.Ident/*struct*/{
                              NamePos: 4823,
                              Name: "Names",
                            }/*struct*/),
                          }/*struct*/),
                          Lbrack: 4828,
                          Index: p.StrmapBasicLit("0xc4200b0ac0",&ast.BasicLit/*struct*/{
                            ValuePos: 4829,
                            Kind: token.Token(5)/*INT*/,
                            Value: "0",
                          }/*struct*/),
                          Rbrack: 4830,
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200b0ae0",&ast.Ident/*struct*/{
                          NamePos: 4832,
                          Name: "Pos",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 4835,
                      Ellipsis: 0,
                      Rparen: 4836,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 4839,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200b0be0",&ast.ReturnStmt/*struct*/{
            Return: 4842,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200b6440",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200b0bc0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200b0b80",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200b0b40",&ast.Ident/*struct*/{
                      NamePos: 4849,
                      Name: "f",
                      Obj: p.PtrmapObject("0xc420051630"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200b0b60",&ast.Ident/*struct*/{
                      NamePos: 4851,
                      Name: "Type",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200b0ba0",&ast.Ident/*struct*/{
                    NamePos: 4856,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 4859,
                Ellipsis: 0,
                Rparen: 4860,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 4862,
      }/*struct*/),
    }/*struct*/),/* slice_item: 15*/16: p.StrmapFuncDecl("0xc4200b46f0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200b4540",&ast.FieldList/*struct*/{
        Opening: 4870,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200b6480",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200b0c40",&ast.Ident/*struct*/{
                NamePos: 4871,
                Name: "f",
                Obj: p.StrmapObject("0xc420051680",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "f",
                  Decl: p.PtrmapField("0xc4200b6480"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200b0c80",&ast.StarExpr/*struct*/{
              Star: 4873,
              X: p.StrmapIdent("0xc4200b0c60",&ast.Ident/*struct*/{
                NamePos: 4874,
                Name: "Field",
                Obj: p.PtrmapObject("0xc420051450"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 4879,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200b0ca0",&ast.Ident/*struct*/{
        NamePos: 4881,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200b0f40",&ast.FuncType/*struct*/{
        Func: 4865,
        Params: p.StrmapFieldList("0xc4200b4570",&ast.FieldList/*struct*/{
          Opening: 4884,
          Closing: 4885,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200b45a0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200b64c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200b0d00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200b0cc0",&ast.Ident/*struct*/{
                  NamePos: 4887,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200b0ce0",&ast.Ident/*struct*/{
                  NamePos: 4893,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200b46c0",&ast.BlockStmt/*struct*/{
        Lbrace: 4897,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200b6540",&ast.IfStmt/*struct*/{
            If: 4900,
            Cond: p.StrmapBinaryExpr("0xc4200b4630",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200b0d60",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200b0d20",&ast.Ident/*struct*/{
                  NamePos: 4903,
                  Name: "f",
                  Obj: p.PtrmapObject("0xc420051680"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200b0d40",&ast.Ident/*struct*/{
                  NamePos: 4905,
                  Name: "Tag",
                }/*struct*/),
              }/*struct*/),
              OpPos: 4909,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200b0d80",&ast.Ident/*struct*/{
                NamePos: 4912,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200b4690",&ast.BlockStmt/*struct*/{
              Lbrace: 4916,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200b0e40",&ast.ReturnStmt/*struct*/{
                  Return: 4920,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200b6500",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200b0e20",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapSelectorExpr("0xc4200b0de0",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200b0da0",&ast.Ident/*struct*/{
                            NamePos: 4927,
                            Name: "f",
                            Obj: p.PtrmapObject("0xc420051680"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200b0dc0",&ast.Ident/*struct*/{
                            NamePos: 4929,
                            Name: "Tag",
                          }/*struct*/),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200b0e00",&ast.Ident/*struct*/{
                          NamePos: 4933,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 4936,
                      Ellipsis: 0,
                      Rparen: 4937,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 4940,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200b0f00",&ast.ReturnStmt/*struct*/{
            Return: 4943,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200b6580",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200b0ee0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200b0ea0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200b0e60",&ast.Ident/*struct*/{
                      NamePos: 4950,
                      Name: "f",
                      Obj: p.PtrmapObject("0xc420051680"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200b0e80",&ast.Ident/*struct*/{
                      NamePos: 4952,
                      Name: "Type",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200b0ec0",&ast.Ident/*struct*/{
                    NamePos: 4957,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 4960,
                Ellipsis: 0,
                Rparen: 4961,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 4963,
      }/*struct*/),
    }/*struct*/),/* slice_item: 16*/17: p.StrmapGenDecl("0xc4200b6680",&ast.GenDecl/*struct*/{
      TokPos: 5045,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200b4720",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200b0f60",&ast.Ident/*struct*/{
            NamePos: 5050,
            Name: "FieldList",
            Obj: p.StrmapObject("0xc420051720",&ast.Object/*struct*/{
              Kind: "type",
              Name: "FieldList",
              Decl: p.PtrmapTypeSpec("0xc4200b4720"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200b1120",&ast.StructType/*struct*/{
            Struct: 5060,
            Fields: p.StrmapFieldList("0xc4200b47e0",&ast.FieldList/*struct*/{
              Opening: 5067,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b65c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200b0f80",&ast.Ident/*struct*/{
                      NamePos: 5070,
                      Name: "Opening",
                      Obj: p.StrmapObject("0xc420051770",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Opening",
                        Decl: p.PtrmapField("0xc4200b65c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200b0fe0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200b0fa0",&ast.Ident/*struct*/{
                      NamePos: 5078,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200b0fc0",&ast.Ident/*struct*/{
                      NamePos: 5084,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b6600",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200b1000",&ast.Ident/*struct*/{
                      NamePos: 5138,
                      Name: "List",
                      Obj: p.StrmapObject("0xc4200517c0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "List",
                        Decl: p.PtrmapField("0xc4200b6600"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200b4780",&ast.ArrayType/*struct*/{
                    Lbrack: 5146,
                    Elt: p.StrmapStarExpr("0xc4200b1040",&ast.StarExpr/*struct*/{
                      Star: 5148,
                      X: p.StrmapIdent("0xc4200b1020",&ast.Ident/*struct*/{
                        NamePos: 5149,
                        Name: "Field",
                        Obj: p.PtrmapObject("0xc420051450"),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b6640",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200b1080",&ast.Ident/*struct*/{
                      NamePos: 5179,
                      Name: "Closing",
                      Obj: p.StrmapObject("0xc420051810",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Closing",
                        Decl: p.PtrmapField("0xc4200b6640"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200b10e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200b10a0",&ast.Ident/*struct*/{
                      NamePos: 5187,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200b10c0",&ast.Ident/*struct*/{
                      NamePos: 5193,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 5246,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 17*/18: p.StrmapFuncDecl("0xc4200b4b10",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200b4840",&ast.FieldList/*struct*/{
        Opening: 5254,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200b66c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200b1140",&ast.Ident/*struct*/{
                NamePos: 5255,
                Name: "f",
                Obj: p.StrmapObject("0xc420051860",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "f",
                  Decl: p.PtrmapField("0xc4200b66c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200b1180",&ast.StarExpr/*struct*/{
              Star: 5257,
              X: p.StrmapIdent("0xc4200b1160",&ast.Ident/*struct*/{
                NamePos: 5258,
                Name: "FieldList",
                Obj: p.PtrmapObject("0xc420051720"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 5267,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200b11a0",&ast.Ident/*struct*/{
        NamePos: 5269,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200b1560",&ast.FuncType/*struct*/{
        Func: 5249,
        Params: p.StrmapFieldList("0xc4200b4870",&ast.FieldList/*struct*/{
          Opening: 5272,
          Closing: 5273,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200b48a0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200b6700",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200b1200",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200b11c0",&ast.Ident/*struct*/{
                  NamePos: 5275,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200b11e0",&ast.Ident/*struct*/{
                  NamePos: 5281,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200b4ae0",&ast.BlockStmt/*struct*/{
        Lbrace: 5285,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200b6780",&ast.IfStmt/*struct*/{
            If: 5288,
            Cond: p.StrmapCallExpr("0xc4200b6740",&ast.CallExpr/*struct*/{
              Fun: p.StrmapSelectorExpr("0xc4200b12a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200b1260",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200b1220",&ast.Ident/*struct*/{
                    NamePos: 5291,
                    Name: "f",
                    Obj: p.PtrmapObject("0xc420051860"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200b1240",&ast.Ident/*struct*/{
                    NamePos: 5293,
                    Name: "Opening",
                  }/*struct*/),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200b1280",&ast.Ident/*struct*/{
                  NamePos: 5301,
                  Name: "IsValid",
                }/*struct*/),
              }/*struct*/),
              Lparen: 5308,
              Ellipsis: 0,
              Rparen: 5309,
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200b4960",&ast.BlockStmt/*struct*/{
              Lbrace: 5311,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200b1320",&ast.ReturnStmt/*struct*/{
                  Return: 5315,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200b1300",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200b12c0",&ast.Ident/*struct*/{
                        NamePos: 5322,
                        Name: "f",
                        Obj: p.PtrmapObject("0xc420051860"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200b12e0",&ast.Ident/*struct*/{
                        NamePos: 5324,
                        Name: "Opening",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 5333,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapIfStmt("0xc4200b6840",&ast.IfStmt/*struct*/{
            If: 5430,
            Cond: p.StrmapBinaryExpr("0xc4200b4a20",&ast.BinaryExpr/*struct*/{
              X: p.StrmapCallExpr("0xc4200b67c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapIdent("0xc4200b1340",&ast.Ident/*struct*/{
                  NamePos: 5433,
                  Name: "len",
                }/*struct*/),
                Lparen: 5436,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapSelectorExpr("0xc4200b13a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200b1360",&ast.Ident/*struct*/{
                      NamePos: 5437,
                      Name: "f",
                      Obj: p.PtrmapObject("0xc420051860"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200b1380",&ast.Ident/*struct*/{
                      NamePos: 5439,
                      Name: "List",
                    }/*struct*/),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 5443,
              }/*struct*/),
              OpPos: 5445,
              Op: token.Token(41)/*>*/,
              Y: p.StrmapBasicLit("0xc4200b13c0",&ast.BasicLit/*struct*/{
                ValuePos: 5447,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200b4ab0",&ast.BlockStmt/*struct*/{
              Lbrace: 5449,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200b14a0",&ast.ReturnStmt/*struct*/{
                  Return: 5453,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200b6800",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200b1480",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIndexExpr("0xc4200b4a80",&ast.IndexExpr/*struct*/{
                          X: p.StrmapSelectorExpr("0xc4200b1420",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200b13e0",&ast.Ident/*struct*/{
                              NamePos: 5460,
                              Name: "f",
                              Obj: p.PtrmapObject("0xc420051860"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200b1400",&ast.Ident/*struct*/{
                              NamePos: 5462,
                              Name: "List",
                            }/*struct*/),
                          }/*struct*/),
                          Lbrack: 5466,
                          Index: p.StrmapBasicLit("0xc4200b1440",&ast.BasicLit/*struct*/{
                            ValuePos: 5467,
                            Kind: token.Token(5)/*INT*/,
                            Value: "0",
                          }/*struct*/),
                          Rbrack: 5468,
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200b1460",&ast.Ident/*struct*/{
                          NamePos: 5470,
                          Name: "Pos",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 5473,
                      Ellipsis: 0,
                      Rparen: 5474,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 5477,
            }/*struct*/),
          }/*struct*/),/* slice_item: 1*/2: p.StrmapReturnStmt("0xc4200b1540",&ast.ReturnStmt/*struct*/{
            Return: 5480,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200b1520",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200b14e0",&ast.Ident/*struct*/{
                  NamePos: 5487,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200b1500",&ast.Ident/*struct*/{
                  NamePos: 5493,
                  Name: "NoPos",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 2*/}/*slice*/,
        Rbrace: 5499,
      }/*struct*/),
    }/*struct*/),/* slice_item: 18*/19: p.StrmapFuncDecl("0xc4200b4ea0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200b4b70",&ast.FieldList/*struct*/{
        Opening: 5507,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200b68c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200b1580",&ast.Ident/*struct*/{
                NamePos: 5508,
                Name: "f",
                Obj: p.StrmapObject("0xc4200518b0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "f",
                  Decl: p.PtrmapField("0xc4200b68c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200b15c0",&ast.StarExpr/*struct*/{
              Star: 5510,
              X: p.StrmapIdent("0xc4200b15a0",&ast.Ident/*struct*/{
                NamePos: 5511,
                Name: "FieldList",
                Obj: p.PtrmapObject("0xc420051720"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 5520,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200b15e0",&ast.Ident/*struct*/{
        NamePos: 5522,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200b1a20",&ast.FuncType/*struct*/{
        Func: 5502,
        Params: p.StrmapFieldList("0xc4200b4ba0",&ast.FieldList/*struct*/{
          Opening: 5525,
          Closing: 5526,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200b4bd0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200b6900",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200b1640",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200b1600",&ast.Ident/*struct*/{
                  NamePos: 5528,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200b1620",&ast.Ident/*struct*/{
                  NamePos: 5534,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200b4e70",&ast.BlockStmt/*struct*/{
        Lbrace: 5538,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200b6980",&ast.IfStmt/*struct*/{
            If: 5541,
            Cond: p.StrmapCallExpr("0xc4200b6940",&ast.CallExpr/*struct*/{
              Fun: p.StrmapSelectorExpr("0xc4200b16e0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200b16a0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200b1660",&ast.Ident/*struct*/{
                    NamePos: 5544,
                    Name: "f",
                    Obj: p.PtrmapObject("0xc4200518b0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200b1680",&ast.Ident/*struct*/{
                    NamePos: 5546,
                    Name: "Closing",
                  }/*struct*/),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200b16c0",&ast.Ident/*struct*/{
                  NamePos: 5554,
                  Name: "IsValid",
                }/*struct*/),
              }/*struct*/),
              Lparen: 5561,
              Ellipsis: 0,
              Rparen: 5562,
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200b4cc0",&ast.BlockStmt/*struct*/{
              Lbrace: 5564,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200b1780",&ast.ReturnStmt/*struct*/{
                  Return: 5568,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapBinaryExpr("0xc4200b4c90",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapSelectorExpr("0xc4200b1740",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIdent("0xc4200b1700",&ast.Ident/*struct*/{
                          NamePos: 5575,
                          Name: "f",
                          Obj: p.PtrmapObject("0xc4200518b0"),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200b1720",&ast.Ident/*struct*/{
                          NamePos: 5577,
                          Name: "Closing",
                        }/*struct*/),
                      }/*struct*/),
                      OpPos: 5585,
                      Op: token.Token(12)/*+*/,
                      Y: p.StrmapBasicLit("0xc4200b1760",&ast.BasicLit/*struct*/{
                        ValuePos: 5587,
                        Kind: token.Token(5)/*INT*/,
                        Value: "1",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 5590,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapIfStmt("0xc4200b6a80",&ast.IfStmt/*struct*/{
            If: 5687,
            Init: p.StrmapAssignStmt("0xc4200b6a00",&ast.AssignStmt/*struct*/{
              Lhs: []ast.Expr /*Slice*/{
                0: p.StrmapIdent("0xc4200b17a0",&ast.Ident/*struct*/{
                  NamePos: 5690,
                  Name: "n",
                  Obj: p.StrmapObject("0xc420051900",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "n",
                    Decl: p.PtrmapAssignStmt("0xc4200b6a00"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              TokPos: 5692,
              Tok: token.Token(47)/*:=*/,
              Rhs: []ast.Expr /*Slice*/{
                0: p.StrmapCallExpr("0xc4200b69c0",&ast.CallExpr/*struct*/{
                  Fun: p.StrmapIdent("0xc4200b17c0",&ast.Ident/*struct*/{
                    NamePos: 5695,
                    Name: "len",
                  }/*struct*/),
                  Lparen: 5698,
                  Args: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200b1820",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200b17e0",&ast.Ident/*struct*/{
                        NamePos: 5699,
                        Name: "f",
                        Obj: p.PtrmapObject("0xc4200518b0"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200b1800",&ast.Ident/*struct*/{
                        NamePos: 5701,
                        Name: "List",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Ellipsis: 0,
                  Rparen: 5705,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
            }/*struct*/),
            Cond: p.StrmapBinaryExpr("0xc4200b4d80",&ast.BinaryExpr/*struct*/{
              X: p.StrmapIdent("0xc4200b1840",&ast.Ident/*struct*/{
                NamePos: 5708,
                Name: "n",
                Obj: p.PtrmapObject("0xc420051900"),
              }/*struct*/),
              OpPos: 5710,
              Op: token.Token(41)/*>*/,
              Y: p.StrmapBasicLit("0xc4200b1860",&ast.BasicLit/*struct*/{
                ValuePos: 5712,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200b4e40",&ast.BlockStmt/*struct*/{
              Lbrace: 5714,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200b1960",&ast.ReturnStmt/*struct*/{
                  Return: 5718,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200b6a40",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200b1940",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIndexExpr("0xc4200b4e10",&ast.IndexExpr/*struct*/{
                          X: p.StrmapSelectorExpr("0xc4200b18c0",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200b1880",&ast.Ident/*struct*/{
                              NamePos: 5725,
                              Name: "f",
                              Obj: p.PtrmapObject("0xc4200518b0"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200b18a0",&ast.Ident/*struct*/{
                              NamePos: 5727,
                              Name: "List",
                            }/*struct*/),
                          }/*struct*/),
                          Lbrack: 5731,
                          Index: p.StrmapBinaryExpr("0xc4200b4de0",&ast.BinaryExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200b18e0",&ast.Ident/*struct*/{
                              NamePos: 5732,
                              Name: "n",
                              Obj: p.PtrmapObject("0xc420051900"),
                            }/*struct*/),
                            OpPos: 5733,
                            Op: token.Token(13)/*-*/,
                            Y: p.StrmapBasicLit("0xc4200b1900",&ast.BasicLit/*struct*/{
                              ValuePos: 5734,
                              Kind: token.Token(5)/*INT*/,
                              Value: "1",
                            }/*struct*/),
                          }/*struct*/),
                          Rbrack: 5735,
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200b1920",&ast.Ident/*struct*/{
                          NamePos: 5737,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 5740,
                      Ellipsis: 0,
                      Rparen: 5741,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 5744,
            }/*struct*/),
          }/*struct*/),/* slice_item: 1*/2: p.StrmapReturnStmt("0xc4200b1a00",&ast.ReturnStmt/*struct*/{
            Return: 5747,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200b19e0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200b19a0",&ast.Ident/*struct*/{
                  NamePos: 5754,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200b19c0",&ast.Ident/*struct*/{
                  NamePos: 5760,
                  Name: "NoPos",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 2*/}/*slice*/,
        Rbrace: 5766,
      }/*struct*/),
    }/*struct*/),/* slice_item: 19*/20: p.StrmapFuncDecl("0xc4200b5200",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200b4f00",&ast.FieldList/*struct*/{
        Opening: 5854,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200b6b00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200b1a40",&ast.Ident/*struct*/{
                NamePos: 5855,
                Name: "f",
                Obj: p.StrmapObject("0xc4200519a0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "f",
                  Decl: p.PtrmapField("0xc4200b6b00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200b1a80",&ast.StarExpr/*struct*/{
              Star: 5857,
              X: p.StrmapIdent("0xc4200b1a60",&ast.Ident/*struct*/{
                NamePos: 5858,
                Name: "FieldList",
                Obj: p.PtrmapObject("0xc420051720"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 5867,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200b1aa0",&ast.Ident/*struct*/{
        NamePos: 5869,
        Name: "NumFields",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200b1e40",&ast.FuncType/*struct*/{
        Func: 5849,
        Params: p.StrmapFieldList("0xc4200b4f30",&ast.FieldList/*struct*/{
          Opening: 5878,
          Closing: 5879,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200b4f60",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200b6b40",&ast.Field/*struct*/{
              Type: p.StrmapIdent("0xc4200b1ac0",&ast.Ident/*struct*/{
                NamePos: 5881,
                Name: "int",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200b51d0",&ast.BlockStmt/*struct*/{
        Lbrace: 5885,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapAssignStmt("0xc4200b6b80",&ast.AssignStmt/*struct*/{
            Lhs: []ast.Expr /*Slice*/{
              0: p.StrmapIdent("0xc4200b1ae0",&ast.Ident/*struct*/{
                NamePos: 5888,
                Name: "n",
                Obj: p.StrmapObject("0xc4200519f0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "n",
                  Decl: p.PtrmapAssignStmt("0xc4200b6b80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            TokPos: 5890,
            Tok: token.Token(47)/*:=*/,
            Rhs: []ast.Expr /*Slice*/{
              0: p.StrmapBasicLit("0xc4200b1b00",&ast.BasicLit/*struct*/{
                ValuePos: 5893,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/1: p.StrmapIfStmt("0xc4200b6dc0",&ast.IfStmt/*struct*/{
            If: 5896,
            Cond: p.StrmapBinaryExpr("0xc4200b4ff0",&ast.BinaryExpr/*struct*/{
              X: p.StrmapIdent("0xc4200b1b20",&ast.Ident/*struct*/{
                NamePos: 5899,
                Name: "f",
                Obj: p.PtrmapObject("0xc4200519a0"),
              }/*struct*/),
              OpPos: 5901,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200b1b40",&ast.Ident/*struct*/{
                NamePos: 5904,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200b51a0",&ast.BlockStmt/*struct*/{
              Lbrace: 5908,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapRangeStmt("0xc420051b30",&ast.RangeStmt/*struct*/{
                  For: 5912,
                  Key: p.StrmapIdent("0xc4200b1b60",&ast.Ident/*struct*/{
                    NamePos: 5916,
                    Name: "_",
                    Obj: p.StrmapObject("0xc420051a40",&ast.Object/*struct*/{
                      Kind: "var",
                      Name: "_",
                      Decl: p.StrmapAssignStmt("0xc4200b6bc0",&ast.AssignStmt/*struct*/{
                        Lhs: []ast.Expr /*Slice*/{
                          0: p.PtrmapIdent("0xc4200b1b60"),/* slice_item: 0*/1: p.StrmapIdent("0xc4200b1b80",&ast.Ident/*struct*/{
                            NamePos: 5919,
                            Name: "g",
                            Obj: p.StrmapObject("0xc420051a90",&ast.Object/*struct*/{
                              Kind: "var",
                              Name: "g",
                              Decl: p.PtrmapAssignStmt("0xc4200b6bc0"),
                            }/*struct*/),
                          }/*struct*/),/* slice_item: 1*/}/*slice*/,
                        TokPos: 5921,
                        Tok: token.Token(47)/*:=*/,
                        Rhs: []ast.Expr /*Slice*/{
                          0: p.StrmapUnaryExpr("0xc4200b1c20",&ast.UnaryExpr/*struct*/{
                            OpPos: 5924,
                            Op: token.Token(79)/*range*/,
                            X: p.StrmapSelectorExpr("0xc4200b1c00",&ast.SelectorExpr/*struct*/{
                              X: p.StrmapIdent("0xc4200b1bc0",&ast.Ident/*struct*/{
                                NamePos: 5930,
                                Name: "f",
                                Obj: p.PtrmapObject("0xc4200519a0"),
                              }/*struct*/),
                              Sel: p.StrmapIdent("0xc4200b1be0",&ast.Ident/*struct*/{
                                NamePos: 5932,
                                Name: "List",
                              }/*struct*/),
                            }/*struct*/),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                  Value: p.PtrmapIdent("0xc4200b1b80"),
                  TokPos: 5921,
                  Tok: token.Token(47)/*:=*/,
                  X: p.PtrmapSelectorExpr("0xc4200b1c00"),
                  Body: p.StrmapBlockStmt("0xc4200b5170",&ast.BlockStmt/*struct*/{
                    Lbrace: 5937,
                    List: []ast.Stmt /*Slice*/{
                      0: p.StrmapAssignStmt("0xc4200b6c40",&ast.AssignStmt/*struct*/{
                        Lhs: []ast.Expr /*Slice*/{
                          0: p.StrmapIdent("0xc4200b1c40",&ast.Ident/*struct*/{
                            NamePos: 5942,
                            Name: "m",
                            Obj: p.StrmapObject("0xc420051ae0",&ast.Object/*struct*/{
                              Kind: "var",
                              Name: "m",
                              Decl: p.PtrmapAssignStmt("0xc4200b6c40"),
                            }/*struct*/),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                        TokPos: 5944,
                        Tok: token.Token(47)/*:=*/,
                        Rhs: []ast.Expr /*Slice*/{
                          0: p.StrmapCallExpr("0xc4200b6c00",&ast.CallExpr/*struct*/{
                            Fun: p.StrmapIdent("0xc4200b1c60",&ast.Ident/*struct*/{
                              NamePos: 5947,
                              Name: "len",
                            }/*struct*/),
                            Lparen: 5950,
                            Args: []ast.Expr /*Slice*/{
                              0: p.StrmapSelectorExpr("0xc4200b1cc0",&ast.SelectorExpr/*struct*/{
                                X: p.StrmapIdent("0xc4200b1c80",&ast.Ident/*struct*/{
                                  NamePos: 5951,
                                  Name: "g",
                                  Obj: p.PtrmapObject("0xc420051a90"),
                                }/*struct*/),
                                Sel: p.StrmapIdent("0xc4200b1ca0",&ast.Ident/*struct*/{
                                  NamePos: 5953,
                                  Name: "Names",
                                }/*struct*/),
                              }/*struct*/),/* slice_item: 0*/}/*slice*/,
                            Ellipsis: 0,
                            Rparen: 5958,
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      }/*struct*/),/* slice_item: 0*/1: p.StrmapIfStmt("0xc4200b6d00",&ast.IfStmt/*struct*/{
                        If: 5963,
                        Cond: p.StrmapBinaryExpr("0xc4200b50e0",&ast.BinaryExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200b1ce0",&ast.Ident/*struct*/{
                            NamePos: 5966,
                            Name: "m",
                            Obj: p.PtrmapObject("0xc420051ae0"),
                          }/*struct*/),
                          OpPos: 5968,
                          Op: token.Token(39)/*==*/,
                          Y: p.StrmapBasicLit("0xc4200b1d00",&ast.BasicLit/*struct*/{
                            ValuePos: 5971,
                            Kind: token.Token(5)/*INT*/,
                            Value: "0",
                          }/*struct*/),
                        }/*struct*/),
                        Body: p.StrmapBlockStmt("0xc4200b5140",&ast.BlockStmt/*struct*/{
                          Lbrace: 5973,
                          List: []ast.Stmt /*Slice*/{
                            0: p.StrmapAssignStmt("0xc4200b6cc0",&ast.AssignStmt/*struct*/{
                              Lhs: []ast.Expr /*Slice*/{
                                0: p.StrmapIdent("0xc4200b1d20",&ast.Ident/*struct*/{
                                  NamePos: 5979,
                                  Name: "m",
                                  Obj: p.PtrmapObject("0xc420051ae0"),
                                }/*struct*/),/* slice_item: 0*/}/*slice*/,
                              TokPos: 5981,
                              Tok: token.Token(42)/*=*/,
                              Rhs: []ast.Expr /*Slice*/{
                                0: p.StrmapBasicLit("0xc4200b1d40",&ast.BasicLit/*struct*/{
                                  ValuePos: 5983,
                                  Kind: token.Token(5)/*INT*/,
                                  Value: "1",
                                }/*struct*/),/* slice_item: 0*/}/*slice*/,
                            }/*struct*/),/* slice_item: 0*/}/*slice*/,
                          Rbrace: 6007,
                        }/*struct*/),
                      }/*struct*/),/* slice_item: 1*/2: p.StrmapAssignStmt("0xc4200b6d40",&ast.AssignStmt/*struct*/{
                        Lhs: []ast.Expr /*Slice*/{
                          0: p.StrmapIdent("0xc4200b1da0",&ast.Ident/*struct*/{
                            NamePos: 6012,
                            Name: "n",
                            Obj: p.PtrmapObject("0xc4200519f0"),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                        TokPos: 6014,
                        Tok: token.Token(23)/*+=*/,
                        Rhs: []ast.Expr /*Slice*/{
                          0: p.StrmapIdent("0xc4200b1dc0",&ast.Ident/*struct*/{
                            NamePos: 6017,
                            Name: "m",
                            Obj: p.PtrmapObject("0xc420051ae0"),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      }/*struct*/),/* slice_item: 2*/}/*slice*/,
                    Rbrace: 6021,
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 6024,
            }/*struct*/),
          }/*struct*/),/* slice_item: 1*/2: p.StrmapReturnStmt("0xc4200b1e20",&ast.ReturnStmt/*struct*/{
            Return: 6027,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapIdent("0xc4200b1e00",&ast.Ident/*struct*/{
                NamePos: 6034,
                Name: "n",
                Obj: p.PtrmapObject("0xc4200519f0"),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 2*/}/*slice*/,
        Rbrace: 6036,
      }/*struct*/),
    }/*struct*/),/* slice_item: 20*/21: p.StrmapGenDecl("0xc4200c2000",&ast.GenDecl/*struct*/{
      TokPos: 6157,
      Tok: token.Token(84)/*type*/,
      Lparen: 6162,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200b5230",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200b1e60",&ast.Ident/*struct*/{
            NamePos: 6308,
            Name: "BadExpr",
            Obj: p.StrmapObject("0xc420051b80",&ast.Object/*struct*/{
              Kind: "type",
              Name: "BadExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b5230"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200b1f40",&ast.StructType/*struct*/{
            Struct: 6316,
            Fields: p.StrmapFieldList("0xc4200b5290",&ast.FieldList/*struct*/{
              Opening: 6323,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b6f40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200b1e80",&ast.Ident/*struct*/{
                      NamePos: 6327,
                      Name: "From",
                      Obj: p.StrmapObject("0xc420051bd0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "From",
                        Decl: p.PtrmapField("0xc4200b6f40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/1: p.StrmapIdent("0xc4200b1ea0",&ast.Ident/*struct*/{
                      NamePos: 6333,
                      Name: "To",
                      Obj: p.StrmapObject("0xc420051c20",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "To",
                        Decl: p.PtrmapField("0xc4200b6f40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 1*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200b1f20",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200b1ee0",&ast.Ident/*struct*/{
                      NamePos: 6336,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200b1f00",&ast.Ident/*struct*/{
                      NamePos: 6342,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Closing: 6383,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/1: p.PtrmapTypeSpec("0xc4200b52f0"),/* slice_item: 1*/2: p.StrmapTypeSpec("0xc4200b5350",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200ba160",&ast.Ident/*struct*/{
            NamePos: 6692,
            Name: "Ellipsis",
            Obj: p.StrmapObject("0xc420051db0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "Ellipsis",
              Decl: p.PtrmapTypeSpec("0xc4200b5350"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200ba260",&ast.StructType/*struct*/{
            Struct: 6701,
            Fields: p.StrmapFieldList("0xc4200b5380",&ast.FieldList/*struct*/{
              Opening: 6708,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b70c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba180",&ast.Ident/*struct*/{
                      NamePos: 6712,
                      Name: "Ellipsis",
                      Obj: p.StrmapObject("0xc420051e00",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Ellipsis",
                        Decl: p.PtrmapField("0xc4200b70c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200ba1e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ba1a0",&ast.Ident/*struct*/{
                      NamePos: 6721,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ba1c0",&ast.Ident/*struct*/{
                      NamePos: 6727,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7140",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba220",&ast.Ident/*struct*/{
                      NamePos: 6754,
                      Name: "Elt",
                      Obj: p.StrmapObject("0xc420051e50",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Elt",
                        Decl: p.PtrmapField("0xc4200b7140"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200ba240",&ast.Ident/*struct*/{
                    NamePos: 6763,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 6830,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 2*/3: p.PtrmapTypeSpec("0xc4200b53b0"),/* slice_item: 3*/4: p.StrmapTypeSpec("0xc4200b5440",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200ba440",&ast.Ident/*struct*/{
            NamePos: 7202,
            Name: "FuncLit",
            Obj: p.StrmapObject("0xc4200c0050",&ast.Object/*struct*/{
              Kind: "type",
              Name: "FuncLit",
              Decl: p.PtrmapTypeSpec("0xc4200b5440"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200ba520",&ast.StructType/*struct*/{
            Struct: 7210,
            Fields: p.StrmapFieldList("0xc4200b5470",&ast.FieldList/*struct*/{
              Opening: 7217,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b72c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba460",&ast.Ident/*struct*/{
                      NamePos: 7221,
                      Name: "Type",
                      Obj: p.StrmapObject("0xc4200c00a0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Type",
                        Decl: p.PtrmapField("0xc4200b72c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200ba4a0",&ast.StarExpr/*struct*/{
                    Star: 7226,
                    X: p.StrmapIdent("0xc4200ba480",&ast.Ident/*struct*/{
                      NamePos: 7227,
                      Name: "FuncType",
                      Obj: p.StrmapObject("0xc4200c1770",&ast.Object/*struct*/{
                        Kind: "type",
                        Name: "FuncType",
                        Decl: p.StrmapTypeSpec("0xc4200b5b60",&ast.TypeSpec/*struct*/{
                          Name: p.StrmapIdent("0xc4200c4020",&ast.Ident/*struct*/{
                            NamePos: 10960,
                            Name: "FuncType",
                            Obj: p.PtrmapObject("0xc4200c1770"),
                          }/*struct*/),
                          Type: p.StrmapStructType("0xc4200c41c0",&ast.StructType/*struct*/{
                            Struct: 10969,
                            Fields: p.StrmapFieldList("0xc4200b5bc0",&ast.FieldList/*struct*/{
                              Opening: 10976,
                              List: []*ast.Field /*Slice*/{
                                0: p.StrmapField("0xc4200c23c0",&ast.Field/*struct*/{
                                  Names: []*ast.Ident /*Slice*/{
                                    0: p.StrmapIdent("0xc4200c4040",&ast.Ident/*struct*/{
                                      NamePos: 10980,
                                      Name: "Func",
                                      Obj: p.StrmapObject("0xc4200c1810",&ast.Object/*struct*/{
                                        Kind: "var",
                                        Name: "Func",
                                        Decl: p.PtrmapField("0xc4200c23c0"),
                                      }/*struct*/),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Type: p.StrmapSelectorExpr("0xc4200c40a0",&ast.SelectorExpr/*struct*/{
                                    X: p.StrmapIdent("0xc4200c4060",&ast.Ident/*struct*/{
                                      NamePos: 10988,
                                      Name: "token",
                                    }/*struct*/),
                                    Sel: p.StrmapIdent("0xc4200c4080",&ast.Ident/*struct*/{
                                      NamePos: 10994,
                                      Name: "Pos",
                                    }/*struct*/),
                                  }/*struct*/),
                                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200c2400",&ast.Field/*struct*/{
                                  Names: []*ast.Ident /*Slice*/{
                                    0: p.StrmapIdent("0xc4200c40c0",&ast.Ident/*struct*/{
                                      NamePos: 11067,
                                      Name: "Params",
                                      Obj: p.StrmapObject("0xc4200c1860",&ast.Object/*struct*/{
                                        Kind: "var",
                                        Name: "Params",
                                        Decl: p.PtrmapField("0xc4200c2400"),
                                      }/*struct*/),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Type: p.StrmapStarExpr("0xc4200c4100",&ast.StarExpr/*struct*/{
                                    Star: 11075,
                                    X: p.StrmapIdent("0xc4200c40e0",&ast.Ident/*struct*/{
                                      NamePos: 11076,
                                      Name: "FieldList",
                                      Obj: p.PtrmapObject("0xc420051720"),
                                    }/*struct*/),
                                  }/*struct*/),
                                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200c2440",&ast.Field/*struct*/{
                                  Names: []*ast.Ident /*Slice*/{
                                    0: p.StrmapIdent("0xc4200c4120",&ast.Ident/*struct*/{
                                      NamePos: 11122,
                                      Name: "Results",
                                      Obj: p.StrmapObject("0xc4200c18b0",&ast.Object/*struct*/{
                                        Kind: "var",
                                        Name: "Results",
                                        Decl: p.PtrmapField("0xc4200c2440"),
                                      }/*struct*/),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Type: p.StrmapStarExpr("0xc4200c4160",&ast.StarExpr/*struct*/{
                                    Star: 11130,
                                    X: p.StrmapIdent("0xc4200c4140",&ast.Ident/*struct*/{
                                      NamePos: 11131,
                                      Name: "FieldList",
                                      Obj: p.PtrmapObject("0xc420051720"),
                                    }/*struct*/),
                                  }/*struct*/),
                                }/*struct*/),/* slice_item: 2*/}/*slice*/,
                              Closing: 11172,
                            }/*struct*/),
                            Incomplete: false,
                          }/*struct*/),
                        }/*struct*/),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7300",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba4c0",&ast.Ident/*struct*/{
                      NamePos: 7256,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200c00f0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200b7300"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200ba500",&ast.StarExpr/*struct*/{
                    Star: 7261,
                    X: p.StrmapIdent("0xc4200ba4e0",&ast.Ident/*struct*/{
                      NamePos: 7262,
                      Name: "BlockStmt",
                      Obj: p.StrmapObject("0xc4200c9b30",&ast.Object/*struct*/{
                        Kind: "type",
                        Name: "BlockStmt",
                        Decl: p.StrmapTypeSpec("0xc4200dc150",&ast.TypeSpec/*struct*/{
                          Name: p.StrmapIdent("0xc4200d9a20",&ast.Ident/*struct*/{
                            NamePos: 19233,
                            Name: "BlockStmt",
                            Obj: p.PtrmapObject("0xc4200c9b30"),
                          }/*struct*/),
                          Type: p.StrmapStructType("0xc4200d9be0",&ast.StructType/*struct*/{
                            Struct: 19243,
                            Fields: p.StrmapFieldList("0xc4200dc1b0",&ast.FieldList/*struct*/{
                              Opening: 19250,
                              List: []*ast.Field /*Slice*/{
                                0: p.StrmapField("0xc4200d1d80",&ast.Field/*struct*/{
                                  Names: []*ast.Ident /*Slice*/{
                                    0: p.StrmapIdent("0xc4200d9a40",&ast.Ident/*struct*/{
                                      NamePos: 19254,
                                      Name: "Lbrace",
                                      Obj: p.StrmapObject("0xc4200c9b80",&ast.Object/*struct*/{
                                        Kind: "var",
                                        Name: "Lbrace",
                                        Decl: p.PtrmapField("0xc4200d1d80"),
                                      }/*struct*/),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Type: p.StrmapSelectorExpr("0xc4200d9aa0",&ast.SelectorExpr/*struct*/{
                                    X: p.StrmapIdent("0xc4200d9a60",&ast.Ident/*struct*/{
                                      NamePos: 19261,
                                      Name: "token",
                                    }/*struct*/),
                                    Sel: p.StrmapIdent("0xc4200d9a80",&ast.Ident/*struct*/{
                                      NamePos: 19267,
                                      Name: "Pos",
                                    }/*struct*/),
                                  }/*struct*/),
                                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1dc0",&ast.Field/*struct*/{
                                  Names: []*ast.Ident /*Slice*/{
                                    0: p.StrmapIdent("0xc4200d9ae0",&ast.Ident/*struct*/{
                                      NamePos: 19292,
                                      Name: "List",
                                      Obj: p.StrmapObject("0xc4200c9bd0",&ast.Object/*struct*/{
                                        Kind: "var",
                                        Name: "List",
                                        Decl: p.PtrmapField("0xc4200d1dc0"),
                                      }/*struct*/),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Type: p.StrmapArrayType("0xc4200dc180",&ast.ArrayType/*struct*/{
                                    Lbrack: 19299,
                                    Elt: p.StrmapIdent("0xc4200d9b00",&ast.Ident/*struct*/{
                                      NamePos: 19301,
                                      Name: "Stmt",
                                      Obj: p.PtrmapObject("0xc4200506e0"),
                                    }/*struct*/),
                                  }/*struct*/),
                                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200d1e00",&ast.Field/*struct*/{
                                  Names: []*ast.Ident /*Slice*/{
                                    0: p.StrmapIdent("0xc4200d9b20",&ast.Ident/*struct*/{
                                      NamePos: 19308,
                                      Name: "Rbrace",
                                      Obj: p.StrmapObject("0xc4200c9c20",&ast.Object/*struct*/{
                                        Kind: "var",
                                        Name: "Rbrace",
                                        Decl: p.PtrmapField("0xc4200d1e00"),
                                      }/*struct*/),
                                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                                  Type: p.StrmapSelectorExpr("0xc4200d9b80",&ast.SelectorExpr/*struct*/{
                                    X: p.StrmapIdent("0xc4200d9b40",&ast.Ident/*struct*/{
                                      NamePos: 19315,
                                      Name: "token",
                                    }/*struct*/),
                                    Sel: p.StrmapIdent("0xc4200d9b60",&ast.Ident/*struct*/{
                                      NamePos: 19321,
                                      Name: "Pos",
                                    }/*struct*/),
                                  }/*struct*/),
                                }/*struct*/),/* slice_item: 2*/}/*slice*/,
                              Closing: 19345,
                            }/*struct*/),
                            Incomplete: false,
                          }/*struct*/),
                        }/*struct*/),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 7290,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 4*/5: p.StrmapTypeSpec("0xc4200b54a0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200ba540",&ast.Ident/*struct*/{
            NamePos: 7350,
            Name: "CompositeLit",
            Obj: p.StrmapObject("0xc4200c0140",&ast.Object/*struct*/{
              Kind: "type",
              Name: "CompositeLit",
              Decl: p.PtrmapTypeSpec("0xc4200b54a0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200ba760",&ast.StructType/*struct*/{
            Struct: 7363,
            Fields: p.StrmapFieldList("0xc4200b5530",&ast.FieldList/*struct*/{
              Opening: 7370,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7380",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba560",&ast.Ident/*struct*/{
                      NamePos: 7374,
                      Name: "Type",
                      Obj: p.StrmapObject("0xc4200c0190",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Type",
                        Decl: p.PtrmapField("0xc4200b7380"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200ba580",&ast.Ident/*struct*/{
                    NamePos: 7381,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b73c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba5c0",&ast.Ident/*struct*/{
                      NamePos: 7417,
                      Name: "Lbrace",
                      Obj: p.StrmapObject("0xc4200c01e0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Lbrace",
                        Decl: p.PtrmapField("0xc4200b73c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200ba620",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ba5e0",&ast.Ident/*struct*/{
                      NamePos: 7424,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ba600",&ast.Ident/*struct*/{
                      NamePos: 7430,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7400",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba660",&ast.Ident/*struct*/{
                      NamePos: 7455,
                      Name: "Elts",
                      Obj: p.StrmapObject("0xc4200c0230",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Elts",
                        Decl: p.PtrmapField("0xc4200b7400"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200b54d0",&ast.ArrayType/*struct*/{
                    Lbrack: 7462,
                    Elt: p.StrmapIdent("0xc4200ba680",&ast.Ident/*struct*/{
                      NamePos: 7464,
                      Name: "Expr",
                      Obj: p.PtrmapObject("0xc420050640"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200b7440",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba6c0",&ast.Ident/*struct*/{
                      NamePos: 7512,
                      Name: "Rbrace",
                      Obj: p.StrmapObject("0xc4200c0280",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Rbrace",
                        Decl: p.PtrmapField("0xc4200b7440"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200ba720",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ba6e0",&ast.Ident/*struct*/{
                      NamePos: 7519,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ba700",&ast.Ident/*struct*/{
                      NamePos: 7525,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 7549,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 5*/6: p.StrmapTypeSpec("0xc4200b5560",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200ba780",&ast.Ident/*struct*/{
            NamePos: 7613,
            Name: "ParenExpr",
            Obj: p.StrmapObject("0xc4200c02d0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "ParenExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b5560"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200ba960",&ast.StructType/*struct*/{
            Struct: 7623,
            Fields: p.StrmapFieldList("0xc4200b5590",&ast.FieldList/*struct*/{
              Opening: 7630,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b74c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba7a0",&ast.Ident/*struct*/{
                      NamePos: 7634,
                      Name: "Lparen",
                      Obj: p.StrmapObject("0xc4200c0320",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Lparen",
                        Decl: p.PtrmapField("0xc4200b74c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200ba800",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ba7c0",&ast.Ident/*struct*/{
                      NamePos: 7641,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ba7e0",&ast.Ident/*struct*/{
                      NamePos: 7647,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7500",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba840",&ast.Ident/*struct*/{
                      NamePos: 7672,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c0370",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200b7500"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200ba860",&ast.Ident/*struct*/{
                    NamePos: 7679,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7540",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba8a0",&ast.Ident/*struct*/{
                      NamePos: 7719,
                      Name: "Rparen",
                      Obj: p.StrmapObject("0xc4200c03c0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Rparen",
                        Decl: p.PtrmapField("0xc4200b7540"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200ba900",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ba8c0",&ast.Ident/*struct*/{
                      NamePos: 7726,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ba8e0",&ast.Ident/*struct*/{
                      NamePos: 7732,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 7756,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 6*/7: p.StrmapTypeSpec("0xc4200b55c0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200ba980",&ast.Ident/*struct*/{
            NamePos: 7833,
            Name: "SelectorExpr",
            Obj: p.StrmapObject("0xc4200c0460",&ast.Object/*struct*/{
              Kind: "type",
              Name: "SelectorExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b55c0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200baa60",&ast.StructType/*struct*/{
            Struct: 7846,
            Fields: p.StrmapFieldList("0xc4200b55f0",&ast.FieldList/*struct*/{
              Opening: 7853,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7580",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba9a0",&ast.Ident/*struct*/{
                      NamePos: 7857,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c04b0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200b7580"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200ba9c0",&ast.Ident/*struct*/{
                    NamePos: 7861,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b75c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200ba9e0",&ast.Ident/*struct*/{
                      NamePos: 7884,
                      Name: "Sel",
                      Obj: p.StrmapObject("0xc4200c0500",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Sel",
                        Decl: p.PtrmapField("0xc4200b75c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200baa20",&ast.StarExpr/*struct*/{
                    Star: 7888,
                    X: p.StrmapIdent("0xc4200baa00",&ast.Ident/*struct*/{
                      NamePos: 7889,
                      Name: "Ident",
                      Obj: p.PtrmapObject("0xc420051c70"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 7914,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 7*/8: p.StrmapTypeSpec("0xc4200b5620",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200baa80",&ast.Ident/*struct*/{
            NamePos: 7987,
            Name: "IndexExpr",
            Obj: p.StrmapObject("0xc4200c05a0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "IndexExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b5620"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200baca0",&ast.StructType/*struct*/{
            Struct: 7997,
            Fields: p.StrmapFieldList("0xc4200b5650",&ast.FieldList/*struct*/{
              Opening: 8004,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7600",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200baaa0",&ast.Ident/*struct*/{
                      NamePos: 8008,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c05f0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200b7600"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200baac0",&ast.Ident/*struct*/{
                    NamePos: 8015,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7640",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200baae0",&ast.Ident/*struct*/{
                      NamePos: 8041,
                      Name: "Lbrack",
                      Obj: p.StrmapObject("0xc4200c0640",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Lbrack",
                        Decl: p.PtrmapField("0xc4200b7640"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bab40",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bab00",&ast.Ident/*struct*/{
                      NamePos: 8048,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bab20",&ast.Ident/*struct*/{
                      NamePos: 8054,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7680",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bab80",&ast.Ident/*struct*/{
                      NamePos: 8079,
                      Name: "Index",
                      Obj: p.StrmapObject("0xc4200c0690",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Index",
                        Decl: p.PtrmapField("0xc4200b7680"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200baba0",&ast.Ident/*struct*/{
                    NamePos: 8086,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200b76c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bac00",&ast.Ident/*struct*/{
                      NamePos: 8118,
                      Name: "Rbrack",
                      Obj: p.StrmapObject("0xc4200c06e0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Rbrack",
                        Decl: p.PtrmapField("0xc4200b76c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bac60",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bac20",&ast.Ident/*struct*/{
                      NamePos: 8125,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bac40",&ast.Ident/*struct*/{
                      NamePos: 8131,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 8155,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 8*/9: p.StrmapTypeSpec("0xc4200b5680",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bacc0",&ast.Ident/*struct*/{
            NamePos: 8233,
            Name: "SliceExpr",
            Obj: p.StrmapObject("0xc4200c0780",&ast.Object/*struct*/{
              Kind: "type",
              Name: "SliceExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b5680"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200bafc0",&ast.StructType/*struct*/{
            Struct: 8243,
            Fields: p.StrmapFieldList("0xc4200b5710",&ast.FieldList/*struct*/{
              Opening: 8250,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7700",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bace0",&ast.Ident/*struct*/{
                      NamePos: 8254,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c07d0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200b7700"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bad00",&ast.Ident/*struct*/{
                    NamePos: 8261,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7740",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bad20",&ast.Ident/*struct*/{
                      NamePos: 8287,
                      Name: "Lbrack",
                      Obj: p.StrmapObject("0xc4200c0820",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Lbrack",
                        Decl: p.PtrmapField("0xc4200b7740"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bad80",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bad40",&ast.Ident/*struct*/{
                      NamePos: 8294,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bad60",&ast.Ident/*struct*/{
                      NamePos: 8300,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7780",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200badc0",&ast.Ident/*struct*/{
                      NamePos: 8325,
                      Name: "Low",
                      Obj: p.StrmapObject("0xc4200c0870",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Low",
                        Decl: p.PtrmapField("0xc4200b7780"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bade0",&ast.Ident/*struct*/{
                    NamePos: 8332,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200b77c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bae40",&ast.Ident/*struct*/{
                      NamePos: 8376,
                      Name: "High",
                      Obj: p.StrmapObject("0xc4200c08c0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "High",
                        Decl: p.PtrmapField("0xc4200b77c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bae60",&ast.Ident/*struct*/{
                    NamePos: 8383,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200b7800",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200baea0",&ast.Ident/*struct*/{
                      NamePos: 8425,
                      Name: "Max",
                      Obj: p.StrmapObject("0xc4200c0910",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Max",
                        Decl: p.PtrmapField("0xc4200b7800"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200baec0",&ast.Ident/*struct*/{
                    NamePos: 8432,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/5: p.StrmapField("0xc4200b7880",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200baee0",&ast.Ident/*struct*/{
                      NamePos: 8481,
                      Name: "Slice3",
                      Obj: p.StrmapObject("0xc4200c0960",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Slice3",
                        Decl: p.PtrmapField("0xc4200b7880"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200baf00",&ast.Ident/*struct*/{
                    NamePos: 8488,
                    Name: "bool",
                  }/*struct*/),
                }/*struct*/),/* slice_item: 5*/6: p.StrmapField("0xc4200b78c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200baf20",&ast.Ident/*struct*/{
                      NamePos: 8544,
                      Name: "Rbrack",
                      Obj: p.StrmapObject("0xc4200c09b0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Rbrack",
                        Decl: p.PtrmapField("0xc4200b78c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200baf80",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200baf40",&ast.Ident/*struct*/{
                      NamePos: 8551,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200baf60",&ast.Ident/*struct*/{
                      NamePos: 8557,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 6*/}/*slice*/,
              Closing: 8581,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 9*/10: p.StrmapTypeSpec("0xc4200b5740",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bb000",&ast.Ident/*struct*/{
            NamePos: 8674,
            Name: "TypeAssertExpr",
            Obj: p.StrmapObject("0xc4200c0a00",&ast.Object/*struct*/{
              Kind: "type",
              Name: "TypeAssertExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b5740"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200bb200",&ast.StructType/*struct*/{
            Struct: 8689,
            Fields: p.StrmapFieldList("0xc4200b57a0",&ast.FieldList/*struct*/{
              Opening: 8696,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7940",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb020",&ast.Ident/*struct*/{
                      NamePos: 8700,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c0a50",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200b7940"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bb040",&ast.Ident/*struct*/{
                    NamePos: 8707,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7980",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb060",&ast.Ident/*struct*/{
                      NamePos: 8733,
                      Name: "Lparen",
                      Obj: p.StrmapObject("0xc4200c0aa0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Lparen",
                        Decl: p.PtrmapField("0xc4200b7980"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb0c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb080",&ast.Ident/*struct*/{
                      NamePos: 8740,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb0a0",&ast.Ident/*struct*/{
                      NamePos: 8746,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b79c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb100",&ast.Ident/*struct*/{
                      NamePos: 8771,
                      Name: "Type",
                      Obj: p.StrmapObject("0xc4200c0af0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Type",
                        Decl: p.PtrmapField("0xc4200b79c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bb120",&ast.Ident/*struct*/{
                    NamePos: 8778,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200b7a00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb160",&ast.Ident/*struct*/{
                      NamePos: 8839,
                      Name: "Rparen",
                      Obj: p.StrmapObject("0xc4200c0b40",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Rparen",
                        Decl: p.PtrmapField("0xc4200b7a00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb1c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb180",&ast.Ident/*struct*/{
                      NamePos: 8846,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb1a0",&ast.Ident/*struct*/{
                      NamePos: 8852,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 8876,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 10*/11: p.StrmapTypeSpec("0xc4200b57d0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bb220",&ast.Ident/*struct*/{
            NamePos: 8955,
            Name: "CallExpr",
            Obj: p.StrmapObject("0xc4200c0be0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "CallExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b57d0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200bb500",&ast.StructType/*struct*/{
            Struct: 8964,
            Fields: p.StrmapFieldList("0xc4200b5830",&ast.FieldList/*struct*/{
              Opening: 8971,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7a40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb240",&ast.Ident/*struct*/{
                      NamePos: 8975,
                      Name: "Fun",
                      Obj: p.StrmapObject("0xc4200c0c30",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Fun",
                        Decl: p.PtrmapField("0xc4200b7a40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bb260",&ast.Ident/*struct*/{
                    NamePos: 8984,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7a80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb2a0",&ast.Ident/*struct*/{
                      NamePos: 9019,
                      Name: "Lparen",
                      Obj: p.StrmapObject("0xc4200c0c80",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Lparen",
                        Decl: p.PtrmapField("0xc4200b7a80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb300",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb2c0",&ast.Ident/*struct*/{
                      NamePos: 9028,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb2e0",&ast.Ident/*struct*/{
                      NamePos: 9034,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7ac0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb340",&ast.Ident/*struct*/{
                      NamePos: 9059,
                      Name: "Args",
                      Obj: p.StrmapObject("0xc4200c0cd0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Args",
                        Decl: p.PtrmapField("0xc4200b7ac0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200b5800",&ast.ArrayType/*struct*/{
                    Lbrack: 9068,
                    Elt: p.StrmapIdent("0xc4200bb360",&ast.Ident/*struct*/{
                      NamePos: 9070,
                      Name: "Expr",
                      Obj: p.PtrmapObject("0xc420050640"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200b7b00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb3c0",&ast.Ident/*struct*/{
                      NamePos: 9110,
                      Name: "Ellipsis",
                      Obj: p.StrmapObject("0xc4200c0d20",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Ellipsis",
                        Decl: p.PtrmapField("0xc4200b7b00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb420",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb3e0",&ast.Ident/*struct*/{
                      NamePos: 9119,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb400",&ast.Ident/*struct*/{
                      NamePos: 9125,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200b7b40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb460",&ast.Ident/*struct*/{
                      NamePos: 9160,
                      Name: "Rparen",
                      Obj: p.StrmapObject("0xc4200c0d70",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Rparen",
                        Decl: p.PtrmapField("0xc4200b7b40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb4c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb480",&ast.Ident/*struct*/{
                      NamePos: 9169,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb4a0",&ast.Ident/*struct*/{
                      NamePos: 9175,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/}/*slice*/,
              Closing: 9199,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 11*/12: p.StrmapTypeSpec("0xc4200b5860",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bb520",&ast.Ident/*struct*/{
            NamePos: 9352,
            Name: "StarExpr",
            Obj: p.StrmapObject("0xc4200c0e60",&ast.Object/*struct*/{
              Kind: "type",
              Name: "StarExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b5860"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200bb620",&ast.StructType/*struct*/{
            Struct: 9361,
            Fields: p.StrmapFieldList("0xc4200b5890",&ast.FieldList/*struct*/{
              Opening: 9368,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7bc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb540",&ast.Ident/*struct*/{
                      NamePos: 9372,
                      Name: "Star",
                      Obj: p.StrmapObject("0xc4200c0eb0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Star",
                        Decl: p.PtrmapField("0xc4200b7bc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb5a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb560",&ast.Ident/*struct*/{
                      NamePos: 9377,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb580",&ast.Ident/*struct*/{
                      NamePos: 9383,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7c00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb5e0",&ast.Ident/*struct*/{
                      NamePos: 9408,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c0f00",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200b7c00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bb600",&ast.Ident/*struct*/{
                    NamePos: 9413,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 9435,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 12*/13: p.StrmapTypeSpec("0xc4200b58c0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bb640",&ast.Ident/*struct*/{
            NamePos: 9557,
            Name: "UnaryExpr",
            Obj: p.StrmapObject("0xc4200c0f50",&ast.Object/*struct*/{
              Kind: "type",
              Name: "UnaryExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b58c0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200bb7e0",&ast.StructType/*struct*/{
            Struct: 9567,
            Fields: p.StrmapFieldList("0xc4200b58f0",&ast.FieldList/*struct*/{
              Opening: 9574,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7cc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb660",&ast.Ident/*struct*/{
                      NamePos: 9578,
                      Name: "OpPos",
                      Obj: p.StrmapObject("0xc4200c0fa0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "OpPos",
                        Decl: p.PtrmapField("0xc4200b7cc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb6c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb680",&ast.Ident/*struct*/{
                      NamePos: 9584,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb6a0",&ast.Ident/*struct*/{
                      NamePos: 9590,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7d00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb700",&ast.Ident/*struct*/{
                      NamePos: 9616,
                      Name: "Op",
                      Obj: p.StrmapObject("0xc4200c0ff0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Op",
                        Decl: p.PtrmapField("0xc4200b7d00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb760",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb720",&ast.Ident/*struct*/{
                      NamePos: 9622,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb740",&ast.Ident/*struct*/{
                      NamePos: 9628,
                      Name: "Token",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7d40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb780",&ast.Ident/*struct*/{
                      NamePos: 9648,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c1040",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200b7d40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bb7a0",&ast.Ident/*struct*/{
                    NamePos: 9654,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 9678,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 13*/14: p.StrmapTypeSpec("0xc4200b5920",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bb800",&ast.Ident/*struct*/{
            NamePos: 9736,
            Name: "BinaryExpr",
            Obj: p.StrmapObject("0xc4200c1090",&ast.Object/*struct*/{
              Kind: "type",
              Name: "BinaryExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b5920"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200bb9e0",&ast.StructType/*struct*/{
            Struct: 9747,
            Fields: p.StrmapFieldList("0xc4200b5950",&ast.FieldList/*struct*/{
              Opening: 9754,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7dc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb820",&ast.Ident/*struct*/{
                      NamePos: 9758,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c10e0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200b7dc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bb840",&ast.Ident/*struct*/{
                    NamePos: 9764,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7e00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb860",&ast.Ident/*struct*/{
                      NamePos: 9794,
                      Name: "OpPos",
                      Obj: p.StrmapObject("0xc4200c1130",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "OpPos",
                        Decl: p.PtrmapField("0xc4200b7e00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb8c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb880",&ast.Ident/*struct*/{
                      NamePos: 9800,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb8a0",&ast.Ident/*struct*/{
                      NamePos: 9806,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7e40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb900",&ast.Ident/*struct*/{
                      NamePos: 9832,
                      Name: "Op",
                      Obj: p.StrmapObject("0xc4200c1180",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Op",
                        Decl: p.PtrmapField("0xc4200b7e40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bb960",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bb920",&ast.Ident/*struct*/{
                      NamePos: 9838,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bb940",&ast.Ident/*struct*/{
                      NamePos: 9844,
                      Name: "Token",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200b7e80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bb9a0",&ast.Ident/*struct*/{
                      NamePos: 9864,
                      Name: "Y",
                      Obj: p.StrmapObject("0xc4200c11d0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Y",
                        Decl: p.PtrmapField("0xc4200b7e80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bb9c0",&ast.Ident/*struct*/{
                    NamePos: 9870,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 9900,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 14*/15: p.StrmapTypeSpec("0xc4200b5980",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bba20",&ast.Ident/*struct*/{
            NamePos: 9990,
            Name: "KeyValueExpr",
            Obj: p.StrmapObject("0xc4200c1220",&ast.Object/*struct*/{
              Kind: "type",
              Name: "KeyValueExpr",
              Decl: p.PtrmapTypeSpec("0xc4200b5980"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200bbb80",&ast.StructType/*struct*/{
            Struct: 10003,
            Fields: p.StrmapFieldList("0xc4200b59b0",&ast.FieldList/*struct*/{
              Opening: 10010,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200b7f00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bba40",&ast.Ident/*struct*/{
                      NamePos: 10014,
                      Name: "Key",
                      Obj: p.StrmapObject("0xc4200c1270",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Key",
                        Decl: p.PtrmapField("0xc4200b7f00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bba60",&ast.Ident/*struct*/{
                    NamePos: 10020,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200b7f40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bba80",&ast.Ident/*struct*/{
                      NamePos: 10027,
                      Name: "Colon",
                      Obj: p.StrmapObject("0xc4200c12c0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Colon",
                        Decl: p.PtrmapField("0xc4200b7f40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bbae0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bbaa0",&ast.Ident/*struct*/{
                      NamePos: 10033,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bbac0",&ast.Ident/*struct*/{
                      NamePos: 10039,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200b7f80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bbb20",&ast.Ident/*struct*/{
                      NamePos: 10064,
                      Name: "Value",
                      Obj: p.StrmapObject("0xc4200c1310",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Value",
                        Decl: p.PtrmapField("0xc4200b7f80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bbb40",&ast.Ident/*struct*/{
                    NamePos: 10070,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 10076,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 15*/}/*slice*/,
      Rparen: 10078,
    }/*struct*/),/* slice_item: 21*/22: p.StrmapGenDecl("0xc4200c2040",&ast.GenDecl/*struct*/{
      TokPos: 10170,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200b59e0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bbbc0",&ast.Ident/*struct*/{
            NamePos: 10175,
            Name: "ChanDir",
            Obj: p.StrmapObject("0xc4200c1360",&ast.Object/*struct*/{
              Kind: "type",
              Name: "ChanDir",
              Decl: p.PtrmapTypeSpec("0xc4200b59e0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapIdent("0xc4200bbbe0",&ast.Ident/*struct*/{
            NamePos: 10183,
            Name: "int",
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 22*/23: p.StrmapGenDecl("0xc4200c2100",&ast.GenDecl/*struct*/{
      TokPos: 10188,
      Tok: token.Token(64)/*const*/,
      Lparen: 10194,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapValueSpec("0xc4200c13b0",&ast.ValueSpec/*struct*/{
          Names: []*ast.Ident /*Slice*/{
            0: p.StrmapIdent("0xc4200bbc00",&ast.Ident/*struct*/{
              NamePos: 10197,
              Name: "SEND",
              Obj: p.StrmapObject("0xc4200c1400",&ast.Object/*struct*/{
                Kind: "const",
                Name: "SEND",
                Decl: p.PtrmapValueSpec("0xc4200c13b0"),
                Data: 0,
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Type: p.StrmapIdent("0xc4200bbc20",&ast.Ident/*struct*/{
            NamePos: 10202,
            Name: "ChanDir",
            Obj: p.PtrmapObject("0xc4200c1360"),
          }/*struct*/),
          Values: []ast.Expr /*Slice*/{
            0: p.StrmapBinaryExpr("0xc4200b5a10",&ast.BinaryExpr/*struct*/{
              X: p.StrmapBasicLit("0xc4200bbc40",&ast.BasicLit/*struct*/{
                ValuePos: 10212,
                Kind: token.Token(5)/*INT*/,
                Value: "1",
              }/*struct*/),
              OpPos: 10214,
              Op: token.Token(20)/*<<*/,
              Y: p.StrmapIdent("0xc4200bbc60",&ast.Ident/*struct*/{
                NamePos: 10217,
                Name: "iota",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
        }/*struct*/),/* slice_item: 0*/1: p.StrmapValueSpec("0xc4200c1450",&ast.ValueSpec/*struct*/{
          Names: []*ast.Ident /*Slice*/{
            0: p.StrmapIdent("0xc4200bbc80",&ast.Ident/*struct*/{
              NamePos: 10223,
              Name: "RECV",
              Obj: p.StrmapObject("0xc4200c14a0",&ast.Object/*struct*/{
                Kind: "const",
                Name: "RECV",
                Decl: p.PtrmapValueSpec("0xc4200c1450"),
                Data: 1,
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
        }/*struct*/),/* slice_item: 1*/}/*slice*/,
      Rparen: 10228,
    }/*struct*/),/* slice_item: 23*/24: p.StrmapGenDecl("0xc4200c2880",&ast.GenDecl/*struct*/{
      TokPos: 10350,
      Tok: token.Token(84)/*type*/,
      Lparen: 10355,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200b5a40",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bbcc0",&ast.Ident/*struct*/{
            NamePos: 10415,
            Name: "ArrayType",
            Obj: p.StrmapObject("0xc4200c14f0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "ArrayType",
              Decl: p.PtrmapTypeSpec("0xc4200b5a40"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200bbe20",&ast.StructType/*struct*/{
            Struct: 10425,
            Fields: p.StrmapFieldList("0xc4200b5a70",&ast.FieldList/*struct*/{
              Opening: 10432,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200c2180",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bbce0",&ast.Ident/*struct*/{
                      NamePos: 10436,
                      Name: "Lbrack",
                      Obj: p.StrmapObject("0xc4200c1540",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Lbrack",
                        Decl: p.PtrmapField("0xc4200c2180"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bbd40",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bbd00",&ast.Ident/*struct*/{
                      NamePos: 10443,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bbd20",&ast.Ident/*struct*/{
                      NamePos: 10449,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200c2200",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bbd80",&ast.Ident/*struct*/{
                      NamePos: 10474,
                      Name: "Len",
                      Obj: p.StrmapObject("0xc4200c1590",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Len",
                        Decl: p.PtrmapField("0xc4200c2200"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bbda0",&ast.Ident/*struct*/{
                    NamePos: 10481,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200c2240",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bbdc0",&ast.Ident/*struct*/{
                      NamePos: 10554,
                      Name: "Elt",
                      Obj: p.StrmapObject("0xc4200c15e0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Elt",
                        Decl: p.PtrmapField("0xc4200c2240"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bbde0",&ast.Ident/*struct*/{
                    NamePos: 10561,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 10588,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/1: p.StrmapTypeSpec("0xc4200b5ad0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200bbe40",&ast.Ident/*struct*/{
            NamePos: 10640,
            Name: "StructType",
            Obj: p.StrmapObject("0xc4200c1630",&ast.Object/*struct*/{
              Kind: "type",
              Name: "StructType",
              Decl: p.PtrmapTypeSpec("0xc4200b5ad0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200bbfe0",&ast.StructType/*struct*/{
            Struct: 10651,
            Fields: p.StrmapFieldList("0xc4200b5b00",&ast.FieldList/*struct*/{
              Opening: 10658,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200c2280",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bbe60",&ast.Ident/*struct*/{
                      NamePos: 10662,
                      Name: "Struct",
                      Obj: p.StrmapObject("0xc4200c1680",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Struct",
                        Decl: p.PtrmapField("0xc4200c2280"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200bbec0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200bbe80",&ast.Ident/*struct*/{
                      NamePos: 10673,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200bbea0",&ast.Ident/*struct*/{
                      NamePos: 10679,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200c22c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bbf00",&ast.Ident/*struct*/{
                      NamePos: 10718,
                      Name: "Fields",
                      Obj: p.StrmapObject("0xc4200c16d0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Fields",
                        Decl: p.PtrmapField("0xc4200c22c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200bbf40",&ast.StarExpr/*struct*/{
                    Star: 10729,
                    X: p.StrmapIdent("0xc4200bbf20",&ast.Ident/*struct*/{
                      NamePos: 10730,
                      Name: "FieldList",
                      Obj: p.PtrmapObject("0xc420051720"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200c2340",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200bbf80",&ast.Ident/*struct*/{
                      NamePos: 10772,
                      Name: "Incomplete",
                      Obj: p.StrmapObject("0xc4200c1720",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Incomplete",
                        Decl: p.PtrmapField("0xc4200c2340"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200bbfa0",&ast.Ident/*struct*/{
                    NamePos: 10783,
                    Name: "bool",
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 10853,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 1*/2: p.PtrmapTypeSpec("0xc4200b5b60"),/* slice_item: 2*/3: p.StrmapTypeSpec("0xc4200b5bf0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200c41e0",&ast.Ident/*struct*/{
            NamePos: 11232,
            Name: "InterfaceType",
            Obj: p.StrmapObject("0xc4200c1900",&ast.Object/*struct*/{
              Kind: "type",
              Name: "InterfaceType",
              Decl: p.PtrmapTypeSpec("0xc4200b5bf0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200c4360",&ast.StructType/*struct*/{
            Struct: 11246,
            Fields: p.StrmapFieldList("0xc4200b5c50",&ast.FieldList/*struct*/{
              Opening: 11253,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200c2500",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c4200",&ast.Ident/*struct*/{
                      NamePos: 11257,
                      Name: "Interface",
                      Obj: p.StrmapObject("0xc4200c1950",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Interface",
                        Decl: p.PtrmapField("0xc4200c2500"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200c4260",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c4220",&ast.Ident/*struct*/{
                      NamePos: 11268,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c4240",&ast.Ident/*struct*/{
                      NamePos: 11274,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200c2540",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c4280",&ast.Ident/*struct*/{
                      NamePos: 11316,
                      Name: "Methods",
                      Obj: p.StrmapObject("0xc4200c19a0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Methods",
                        Decl: p.PtrmapField("0xc4200c2540"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200c42c0",&ast.StarExpr/*struct*/{
                    Star: 11327,
                    X: p.StrmapIdent("0xc4200c42a0",&ast.Ident/*struct*/{
                      NamePos: 11328,
                      Name: "FieldList",
                      Obj: p.PtrmapObject("0xc420051720"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200c25c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c4300",&ast.Ident/*struct*/{
                      NamePos: 11359,
                      Name: "Incomplete",
                      Obj: p.StrmapObject("0xc4200c19f0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Incomplete",
                        Decl: p.PtrmapField("0xc4200c25c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200c4320",&ast.Ident/*struct*/{
                    NamePos: 11370,
                    Name: "bool",
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 11442,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 3*/4: p.StrmapTypeSpec("0xc4200b5cb0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200c4380",&ast.Ident/*struct*/{
            NamePos: 11488,
            Name: "MapType",
            Obj: p.StrmapObject("0xc4200c1a40",&ast.Object/*struct*/{
              Kind: "type",
              Name: "MapType",
              Decl: p.PtrmapTypeSpec("0xc4200b5cb0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200c44e0",&ast.StructType/*struct*/{
            Struct: 11496,
            Fields: p.StrmapFieldList("0xc4200b5ce0",&ast.FieldList/*struct*/{
              Opening: 11503,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200c2600",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c43a0",&ast.Ident/*struct*/{
                      NamePos: 11507,
                      Name: "Map",
                      Obj: p.StrmapObject("0xc4200c1a90",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Map",
                        Decl: p.PtrmapField("0xc4200c2600"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200c4400",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c43c0",&ast.Ident/*struct*/{
                      NamePos: 11513,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c43e0",&ast.Ident/*struct*/{
                      NamePos: 11519,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200c2640",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c4440",&ast.Ident/*struct*/{
                      NamePos: 11554,
                      Name: "Key",
                      Obj: p.StrmapObject("0xc4200c1ae0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Key",
                        Decl: p.PtrmapField("0xc4200c2640"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200c4460",&ast.Ident/*struct*/{
                    NamePos: 11560,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200c2680",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c4480",&ast.Ident/*struct*/{
                      NamePos: 11567,
                      Name: "Value",
                      Obj: p.StrmapObject("0xc4200c1b30",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Value",
                        Decl: p.PtrmapField("0xc4200c2680"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200c44a0",&ast.Ident/*struct*/{
                    NamePos: 11573,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 11579,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 4*/5: p.StrmapTypeSpec("0xc4200b5d40",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200c4500",&ast.Ident/*struct*/{
            NamePos: 11630,
            Name: "ChanType",
            Obj: p.StrmapObject("0xc4200c1b80",&ast.Object/*struct*/{
              Kind: "type",
              Name: "ChanType",
              Decl: p.PtrmapTypeSpec("0xc4200b5d40"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200c46e0",&ast.StructType/*struct*/{
            Struct: 11639,
            Fields: p.StrmapFieldList("0xc4200b5d70",&ast.FieldList/*struct*/{
              Opening: 11646,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200c2700",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c4520",&ast.Ident/*struct*/{
                      NamePos: 11650,
                      Name: "Begin",
                      Obj: p.StrmapObject("0xc4200c1bd0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Begin",
                        Decl: p.PtrmapField("0xc4200c2700"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200c4580",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c4540",&ast.Ident/*struct*/{
                      NamePos: 11656,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c4560",&ast.Ident/*struct*/{
                      NamePos: 11662,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200c2780",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c45a0",&ast.Ident/*struct*/{
                      NamePos: 11730,
                      Name: "Arrow",
                      Obj: p.StrmapObject("0xc4200c1c20",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Arrow",
                        Decl: p.PtrmapField("0xc4200c2780"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200c4600",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c45c0",&ast.Ident/*struct*/{
                      NamePos: 11736,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c45e0",&ast.Ident/*struct*/{
                      NamePos: 11742,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200c27c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c4620",&ast.Ident/*struct*/{
                      NamePos: 11802,
                      Name: "Dir",
                      Obj: p.StrmapObject("0xc4200c1c70",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Dir",
                        Decl: p.PtrmapField("0xc4200c27c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200c4640",&ast.Ident/*struct*/{
                    NamePos: 11808,
                    Name: "ChanDir",
                    Obj: p.PtrmapObject("0xc4200c1360"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200c2800",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200c46a0",&ast.Ident/*struct*/{
                      NamePos: 11841,
                      Name: "Value",
                      Obj: p.StrmapObject("0xc4200c1cc0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Value",
                        Decl: p.PtrmapField("0xc4200c2800"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200c46c0",&ast.Ident/*struct*/{
                    NamePos: 11847,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 11872,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 5*/}/*slice*/,
      Rparen: 11874,
    }/*struct*/),/* slice_item: 24*/25: p.StrmapFuncDecl("0xc4200b5ec0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200b5dd0",&ast.FieldList/*struct*/{
        Opening: 11941,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c28c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c4700",&ast.Ident/*struct*/{
                NamePos: 11942,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c1d10",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c28c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c4740",&ast.StarExpr/*struct*/{
              Star: 11944,
              X: p.StrmapIdent("0xc4200c4720",&ast.Ident/*struct*/{
                NamePos: 11945,
                Name: "BadExpr",
                Obj: p.PtrmapObject("0xc420051b80"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 11952,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c4760",&ast.Ident/*struct*/{
        NamePos: 11954,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c4860",&ast.FuncType/*struct*/{
        Func: 11936,
        Params: p.StrmapFieldList("0xc4200b5e00",&ast.FieldList/*struct*/{
          Opening: 11957,
          Closing: 11958,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200b5e30",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2900",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c47c0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4780",&ast.Ident/*struct*/{
                  NamePos: 11960,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c47a0",&ast.Ident/*struct*/{
                  NamePos: 11966,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200b5e90",&ast.BlockStmt/*struct*/{
        Lbrace: 11971,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c4840",&ast.ReturnStmt/*struct*/{
            Return: 11973,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200c4820",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c47e0",&ast.Ident/*struct*/{
                  NamePos: 11980,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c1d10"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4800",&ast.Ident/*struct*/{
                  NamePos: 11982,
                  Name: "From",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 11987,
      }/*struct*/),
    }/*struct*/),/* slice_item: 25*/26: p.StrmapFuncDecl("0xc4200c6030",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200b5f20",&ast.FieldList/*struct*/{
        Opening: 11994,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c2940",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c4880",&ast.Ident/*struct*/{
                NamePos: 11995,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c1d60",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c2940"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c48c0",&ast.StarExpr/*struct*/{
              Star: 11997,
              X: p.StrmapIdent("0xc4200c48a0",&ast.Ident/*struct*/{
                NamePos: 11998,
                Name: "Ident",
                Obj: p.PtrmapObject("0xc420051c70"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12003,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c48e0",&ast.Ident/*struct*/{
        NamePos: 12005,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c49e0",&ast.FuncType/*struct*/{
        Func: 11989,
        Params: p.StrmapFieldList("0xc4200b5f50",&ast.FieldList/*struct*/{
          Opening: 12008,
          Closing: 12009,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200b5f80",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2980",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c4940",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4900",&ast.Ident/*struct*/{
                  NamePos: 12011,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4920",&ast.Ident/*struct*/{
                  NamePos: 12017,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c6000",&ast.BlockStmt/*struct*/{
        Lbrace: 12024,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c49c0",&ast.ReturnStmt/*struct*/{
            Return: 12026,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200c49a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4960",&ast.Ident/*struct*/{
                  NamePos: 12033,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c1d60"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4980",&ast.Ident/*struct*/{
                  NamePos: 12035,
                  Name: "NamePos",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12043,
      }/*struct*/),
    }/*struct*/),/* slice_item: 26*/27: p.StrmapFuncDecl("0xc4200c6180",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6090",&ast.FieldList/*struct*/{
        Opening: 12050,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c29c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c4a00",&ast.Ident/*struct*/{
                NamePos: 12051,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c1db0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c29c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c4a40",&ast.StarExpr/*struct*/{
              Star: 12053,
              X: p.StrmapIdent("0xc4200c4a20",&ast.Ident/*struct*/{
                NamePos: 12054,
                Name: "Ellipsis",
                Obj: p.PtrmapObject("0xc420051db0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12062,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c4a60",&ast.Ident/*struct*/{
        NamePos: 12064,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c4b60",&ast.FuncType/*struct*/{
        Func: 12045,
        Params: p.StrmapFieldList("0xc4200c60c0",&ast.FieldList/*struct*/{
          Opening: 12067,
          Closing: 12068,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c60f0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2a00",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c4ac0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4a80",&ast.Ident/*struct*/{
                  NamePos: 12070,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4aa0",&ast.Ident/*struct*/{
                  NamePos: 12076,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c6150",&ast.BlockStmt/*struct*/{
        Lbrace: 12080,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c4b40",&ast.ReturnStmt/*struct*/{
            Return: 12082,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200c4b20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4ae0",&ast.Ident/*struct*/{
                  NamePos: 12089,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c1db0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4b00",&ast.Ident/*struct*/{
                  NamePos: 12091,
                  Name: "Ellipsis",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12100,
      }/*struct*/),
    }/*struct*/),/* slice_item: 27*/28: p.StrmapFuncDecl("0xc4200c62d0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c61e0",&ast.FieldList/*struct*/{
        Opening: 12107,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c2a40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c4b80",&ast.Ident/*struct*/{
                NamePos: 12108,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c1e00",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c2a40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c4bc0",&ast.StarExpr/*struct*/{
              Star: 12110,
              X: p.StrmapIdent("0xc4200c4ba0",&ast.Ident/*struct*/{
                NamePos: 12111,
                Name: "BasicLit",
                Obj: p.PtrmapObject("0xc420051ea0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12119,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c4be0",&ast.Ident/*struct*/{
        NamePos: 12121,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c4ce0",&ast.FuncType/*struct*/{
        Func: 12102,
        Params: p.StrmapFieldList("0xc4200c6210",&ast.FieldList/*struct*/{
          Opening: 12124,
          Closing: 12125,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c6240",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2a80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c4c40",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4c00",&ast.Ident/*struct*/{
                  NamePos: 12127,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4c20",&ast.Ident/*struct*/{
                  NamePos: 12133,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c62a0",&ast.BlockStmt/*struct*/{
        Lbrace: 12137,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c4cc0",&ast.ReturnStmt/*struct*/{
            Return: 12139,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200c4ca0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4c60",&ast.Ident/*struct*/{
                  NamePos: 12146,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c1e00"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4c80",&ast.Ident/*struct*/{
                  NamePos: 12148,
                  Name: "ValuePos",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12157,
      }/*struct*/),
    }/*struct*/),/* slice_item: 28*/29: p.StrmapFuncDecl("0xc4200c6420",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6330",&ast.FieldList/*struct*/{
        Opening: 12164,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c2ac0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c4d00",&ast.Ident/*struct*/{
                NamePos: 12165,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c1e50",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c2ac0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c4d40",&ast.StarExpr/*struct*/{
              Star: 12167,
              X: p.StrmapIdent("0xc4200c4d20",&ast.Ident/*struct*/{
                NamePos: 12168,
                Name: "FuncLit",
                Obj: p.PtrmapObject("0xc4200c0050"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12175,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c4d60",&ast.Ident/*struct*/{
        NamePos: 12177,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c4ea0",&ast.FuncType/*struct*/{
        Func: 12159,
        Params: p.StrmapFieldList("0xc4200c6360",&ast.FieldList/*struct*/{
          Opening: 12180,
          Closing: 12181,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c6390",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2b00",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c4dc0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4d80",&ast.Ident/*struct*/{
                  NamePos: 12183,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4da0",&ast.Ident/*struct*/{
                  NamePos: 12189,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c63f0",&ast.BlockStmt/*struct*/{
        Lbrace: 12194,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c4e80",&ast.ReturnStmt/*struct*/{
            Return: 12196,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c2b40",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200c4e60",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200c4e20",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c4de0",&ast.Ident/*struct*/{
                      NamePos: 12203,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c1e50"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c4e00",&ast.Ident/*struct*/{
                      NamePos: 12205,
                      Name: "Type",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200c4e40",&ast.Ident/*struct*/{
                    NamePos: 12210,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 12213,
                Ellipsis: 0,
                Rparen: 12214,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12216,
      }/*struct*/),
    }/*struct*/),/* slice_item: 29*/30: p.StrmapFuncDecl("0xc4200c6630",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6480",&ast.FieldList/*struct*/{
        Opening: 12223,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c2b80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c4ec0",&ast.Ident/*struct*/{
                NamePos: 12224,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c1ea0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c2b80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c4f00",&ast.StarExpr/*struct*/{
              Star: 12226,
              X: p.StrmapIdent("0xc4200c4ee0",&ast.Ident/*struct*/{
                NamePos: 12227,
                Name: "CompositeLit",
                Obj: p.PtrmapObject("0xc4200c0140"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12239,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c4f20",&ast.Ident/*struct*/{
        NamePos: 12241,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c5180",&ast.FuncType/*struct*/{
        Func: 12218,
        Params: p.StrmapFieldList("0xc4200c64b0",&ast.FieldList/*struct*/{
          Opening: 12244,
          Closing: 12245,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c64e0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2bc0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c4f80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4f40",&ast.Ident/*struct*/{
                  NamePos: 12247,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4f60",&ast.Ident/*struct*/{
                  NamePos: 12253,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c6600",&ast.BlockStmt/*struct*/{
        Lbrace: 12257,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200c2c40",&ast.IfStmt/*struct*/{
            If: 12260,
            Cond: p.StrmapBinaryExpr("0xc4200c6570",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200c4fe0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c4fa0",&ast.Ident/*struct*/{
                  NamePos: 12263,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c1ea0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c4fc0",&ast.Ident/*struct*/{
                  NamePos: 12265,
                  Name: "Type",
                }/*struct*/),
              }/*struct*/),
              OpPos: 12270,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200c5000",&ast.Ident/*struct*/{
                NamePos: 12273,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200c65d0",&ast.BlockStmt/*struct*/{
              Lbrace: 12277,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200c50c0",&ast.ReturnStmt/*struct*/{
                  Return: 12281,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200c2c00",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200c50a0",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapSelectorExpr("0xc4200c5060",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200c5020",&ast.Ident/*struct*/{
                            NamePos: 12288,
                            Name: "x",
                            Obj: p.PtrmapObject("0xc4200c1ea0"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200c5040",&ast.Ident/*struct*/{
                            NamePos: 12290,
                            Name: "Type",
                          }/*struct*/),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200c5080",&ast.Ident/*struct*/{
                          NamePos: 12295,
                          Name: "Pos",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 12298,
                      Ellipsis: 0,
                      Rparen: 12299,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 12302,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200c5140",&ast.ReturnStmt/*struct*/{
            Return: 12305,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200c5120",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c50e0",&ast.Ident/*struct*/{
                  NamePos: 12312,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c1ea0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5100",&ast.Ident/*struct*/{
                  NamePos: 12314,
                  Name: "Lbrace",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 12321,
      }/*struct*/),
    }/*struct*/),/* slice_item: 30*/31: p.StrmapFuncDecl("0xc4200c6780",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6690",&ast.FieldList/*struct*/{
        Opening: 12328,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c2c80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c51a0",&ast.Ident/*struct*/{
                NamePos: 12329,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c1ef0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c2c80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c51e0",&ast.StarExpr/*struct*/{
              Star: 12331,
              X: p.StrmapIdent("0xc4200c51c0",&ast.Ident/*struct*/{
                NamePos: 12332,
                Name: "ParenExpr",
                Obj: p.PtrmapObject("0xc4200c02d0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12341,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c5200",&ast.Ident/*struct*/{
        NamePos: 12343,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c5300",&ast.FuncType/*struct*/{
        Func: 12323,
        Params: p.StrmapFieldList("0xc4200c66c0",&ast.FieldList/*struct*/{
          Opening: 12346,
          Closing: 12347,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c66f0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2cc0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c5260",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5220",&ast.Ident/*struct*/{
                  NamePos: 12349,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5240",&ast.Ident/*struct*/{
                  NamePos: 12355,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c6750",&ast.BlockStmt/*struct*/{
        Lbrace: 12364,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c52e0",&ast.ReturnStmt/*struct*/{
            Return: 12366,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200c52c0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5280",&ast.Ident/*struct*/{
                  NamePos: 12373,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c1ef0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c52a0",&ast.Ident/*struct*/{
                  NamePos: 12375,
                  Name: "Lparen",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12382,
      }/*struct*/),
    }/*struct*/),/* slice_item: 31*/32: p.StrmapFuncDecl("0xc4200c68d0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c67e0",&ast.FieldList/*struct*/{
        Opening: 12389,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c2d00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c5320",&ast.Ident/*struct*/{
                NamePos: 12390,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c1f40",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c2d00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c5360",&ast.StarExpr/*struct*/{
              Star: 12392,
              X: p.StrmapIdent("0xc4200c5340",&ast.Ident/*struct*/{
                NamePos: 12393,
                Name: "SelectorExpr",
                Obj: p.PtrmapObject("0xc4200c0460"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12405,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c5380",&ast.Ident/*struct*/{
        NamePos: 12407,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c54c0",&ast.FuncType/*struct*/{
        Func: 12384,
        Params: p.StrmapFieldList("0xc4200c6810",&ast.FieldList/*struct*/{
          Opening: 12410,
          Closing: 12411,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c6840",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2d40",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c53e0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c53a0",&ast.Ident/*struct*/{
                  NamePos: 12413,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c53c0",&ast.Ident/*struct*/{
                  NamePos: 12419,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c68a0",&ast.BlockStmt/*struct*/{
        Lbrace: 12425,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c54a0",&ast.ReturnStmt/*struct*/{
            Return: 12427,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c2d80",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200c5480",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200c5440",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c5400",&ast.Ident/*struct*/{
                      NamePos: 12434,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c1f40"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c5420",&ast.Ident/*struct*/{
                      NamePos: 12436,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200c5460",&ast.Ident/*struct*/{
                    NamePos: 12438,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 12441,
                Ellipsis: 0,
                Rparen: 12442,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12444,
      }/*struct*/),
    }/*struct*/),/* slice_item: 32*/33: p.StrmapFuncDecl("0xc4200c6a20",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6930",&ast.FieldList/*struct*/{
        Opening: 12451,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c2dc0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c54e0",&ast.Ident/*struct*/{
                NamePos: 12452,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c1f90",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c2dc0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c5520",&ast.StarExpr/*struct*/{
              Star: 12454,
              X: p.StrmapIdent("0xc4200c5500",&ast.Ident/*struct*/{
                NamePos: 12455,
                Name: "IndexExpr",
                Obj: p.PtrmapObject("0xc4200c05a0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12464,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c5540",&ast.Ident/*struct*/{
        NamePos: 12466,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c5680",&ast.FuncType/*struct*/{
        Func: 12446,
        Params: p.StrmapFieldList("0xc4200c6960",&ast.FieldList/*struct*/{
          Opening: 12469,
          Closing: 12470,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c6990",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2e00",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c55a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5560",&ast.Ident/*struct*/{
                  NamePos: 12472,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5580",&ast.Ident/*struct*/{
                  NamePos: 12478,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c69f0",&ast.BlockStmt/*struct*/{
        Lbrace: 12487,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c5660",&ast.ReturnStmt/*struct*/{
            Return: 12489,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c2e40",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200c5640",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200c5600",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c55c0",&ast.Ident/*struct*/{
                      NamePos: 12496,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c1f90"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c55e0",&ast.Ident/*struct*/{
                      NamePos: 12498,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200c5620",&ast.Ident/*struct*/{
                    NamePos: 12500,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 12503,
                Ellipsis: 0,
                Rparen: 12504,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12506,
      }/*struct*/),
    }/*struct*/),/* slice_item: 33*/34: p.StrmapFuncDecl("0xc4200c6b70",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6a80",&ast.FieldList/*struct*/{
        Opening: 12513,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c2e80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c56a0",&ast.Ident/*struct*/{
                NamePos: 12514,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8000",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c2e80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c56e0",&ast.StarExpr/*struct*/{
              Star: 12516,
              X: p.StrmapIdent("0xc4200c56c0",&ast.Ident/*struct*/{
                NamePos: 12517,
                Name: "SliceExpr",
                Obj: p.PtrmapObject("0xc4200c0780"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12526,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c5700",&ast.Ident/*struct*/{
        NamePos: 12528,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c5840",&ast.FuncType/*struct*/{
        Func: 12508,
        Params: p.StrmapFieldList("0xc4200c6ab0",&ast.FieldList/*struct*/{
          Opening: 12531,
          Closing: 12532,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c6ae0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2ec0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c5760",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5720",&ast.Ident/*struct*/{
                  NamePos: 12534,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5740",&ast.Ident/*struct*/{
                  NamePos: 12540,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c6b40",&ast.BlockStmt/*struct*/{
        Lbrace: 12549,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c5820",&ast.ReturnStmt/*struct*/{
            Return: 12551,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c2f00",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200c5800",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200c57c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c5780",&ast.Ident/*struct*/{
                      NamePos: 12558,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8000"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c57a0",&ast.Ident/*struct*/{
                      NamePos: 12560,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200c57e0",&ast.Ident/*struct*/{
                    NamePos: 12562,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 12565,
                Ellipsis: 0,
                Rparen: 12566,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12568,
      }/*struct*/),
    }/*struct*/),/* slice_item: 34*/35: p.StrmapFuncDecl("0xc4200c6cc0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6bd0",&ast.FieldList/*struct*/{
        Opening: 12575,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c2f40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c5860",&ast.Ident/*struct*/{
                NamePos: 12576,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8050",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c2f40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c58a0",&ast.StarExpr/*struct*/{
              Star: 12578,
              X: p.StrmapIdent("0xc4200c5880",&ast.Ident/*struct*/{
                NamePos: 12579,
                Name: "TypeAssertExpr",
                Obj: p.PtrmapObject("0xc4200c0a00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12593,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c58c0",&ast.Ident/*struct*/{
        NamePos: 12595,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c5a00",&ast.FuncType/*struct*/{
        Func: 12570,
        Params: p.StrmapFieldList("0xc4200c6c00",&ast.FieldList/*struct*/{
          Opening: 12598,
          Closing: 12599,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c6c30",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c2f80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c5920",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c58e0",&ast.Ident/*struct*/{
                  NamePos: 12601,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5900",&ast.Ident/*struct*/{
                  NamePos: 12607,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c6c90",&ast.BlockStmt/*struct*/{
        Lbrace: 12611,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c59e0",&ast.ReturnStmt/*struct*/{
            Return: 12613,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c2fc0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200c59c0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200c5980",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c5940",&ast.Ident/*struct*/{
                      NamePos: 12620,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8050"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c5960",&ast.Ident/*struct*/{
                      NamePos: 12622,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200c59a0",&ast.Ident/*struct*/{
                    NamePos: 12624,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 12627,
                Ellipsis: 0,
                Rparen: 12628,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12630,
      }/*struct*/),
    }/*struct*/),/* slice_item: 35*/36: p.StrmapFuncDecl("0xc4200c6e10",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6d20",&ast.FieldList/*struct*/{
        Opening: 12637,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3000",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c5a20",&ast.Ident/*struct*/{
                NamePos: 12638,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c80a0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3000"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c5a60",&ast.StarExpr/*struct*/{
              Star: 12640,
              X: p.StrmapIdent("0xc4200c5a40",&ast.Ident/*struct*/{
                NamePos: 12641,
                Name: "CallExpr",
                Obj: p.PtrmapObject("0xc4200c0be0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12649,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c5a80",&ast.Ident/*struct*/{
        NamePos: 12651,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c5bc0",&ast.FuncType/*struct*/{
        Func: 12632,
        Params: p.StrmapFieldList("0xc4200c6d50",&ast.FieldList/*struct*/{
          Opening: 12654,
          Closing: 12655,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c6d80",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3040",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c5ae0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5aa0",&ast.Ident/*struct*/{
                  NamePos: 12657,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5ac0",&ast.Ident/*struct*/{
                  NamePos: 12663,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c6de0",&ast.BlockStmt/*struct*/{
        Lbrace: 12673,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c5ba0",&ast.ReturnStmt/*struct*/{
            Return: 12675,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c3080",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200c5b80",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200c5b40",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c5b00",&ast.Ident/*struct*/{
                      NamePos: 12682,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c80a0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c5b20",&ast.Ident/*struct*/{
                      NamePos: 12684,
                      Name: "Fun",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200c5b60",&ast.Ident/*struct*/{
                    NamePos: 12688,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 12691,
                Ellipsis: 0,
                Rparen: 12692,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12694,
      }/*struct*/),
    }/*struct*/),/* slice_item: 36*/37: p.StrmapFuncDecl("0xc4200c6f60",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6e70",&ast.FieldList/*struct*/{
        Opening: 12701,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c30c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c5be0",&ast.Ident/*struct*/{
                NamePos: 12702,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c80f0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c30c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c5c20",&ast.StarExpr/*struct*/{
              Star: 12704,
              X: p.StrmapIdent("0xc4200c5c00",&ast.Ident/*struct*/{
                NamePos: 12705,
                Name: "StarExpr",
                Obj: p.PtrmapObject("0xc4200c0e60"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12713,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c5c40",&ast.Ident/*struct*/{
        NamePos: 12715,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c5d40",&ast.FuncType/*struct*/{
        Func: 12696,
        Params: p.StrmapFieldList("0xc4200c6ea0",&ast.FieldList/*struct*/{
          Opening: 12718,
          Closing: 12719,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c6ed0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3100",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c5ca0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5c60",&ast.Ident/*struct*/{
                  NamePos: 12721,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5c80",&ast.Ident/*struct*/{
                  NamePos: 12727,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c6f30",&ast.BlockStmt/*struct*/{
        Lbrace: 12737,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c5d20",&ast.ReturnStmt/*struct*/{
            Return: 12739,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200c5d00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5cc0",&ast.Ident/*struct*/{
                  NamePos: 12746,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c80f0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5ce0",&ast.Ident/*struct*/{
                  NamePos: 12748,
                  Name: "Star",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12753,
      }/*struct*/),
    }/*struct*/),/* slice_item: 37*/38: p.StrmapFuncDecl("0xc4200c70b0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c6fc0",&ast.FieldList/*struct*/{
        Opening: 12760,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3140",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c5d60",&ast.Ident/*struct*/{
                NamePos: 12761,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8140",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3140"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c5da0",&ast.StarExpr/*struct*/{
              Star: 12763,
              X: p.StrmapIdent("0xc4200c5d80",&ast.Ident/*struct*/{
                NamePos: 12764,
                Name: "UnaryExpr",
                Obj: p.PtrmapObject("0xc4200c0f50"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12773,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c5dc0",&ast.Ident/*struct*/{
        NamePos: 12775,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200c5ec0",&ast.FuncType/*struct*/{
        Func: 12755,
        Params: p.StrmapFieldList("0xc4200c6ff0",&ast.FieldList/*struct*/{
          Opening: 12778,
          Closing: 12779,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c7020",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3180",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c5e20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5de0",&ast.Ident/*struct*/{
                  NamePos: 12781,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5e00",&ast.Ident/*struct*/{
                  NamePos: 12787,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c7080",&ast.BlockStmt/*struct*/{
        Lbrace: 12796,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200c5ea0",&ast.ReturnStmt/*struct*/{
            Return: 12798,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200c5e80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5e40",&ast.Ident/*struct*/{
                  NamePos: 12805,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c8140"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5e60",&ast.Ident/*struct*/{
                  NamePos: 12807,
                  Name: "OpPos",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12813,
      }/*struct*/),
    }/*struct*/),/* slice_item: 38*/39: p.StrmapFuncDecl("0xc4200c7200",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c7110",&ast.FieldList/*struct*/{
        Opening: 12820,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c31c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200c5ee0",&ast.Ident/*struct*/{
                NamePos: 12821,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8190",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c31c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200c5f20",&ast.StarExpr/*struct*/{
              Star: 12823,
              X: p.StrmapIdent("0xc4200c5f00",&ast.Ident/*struct*/{
                NamePos: 12824,
                Name: "BinaryExpr",
                Obj: p.PtrmapObject("0xc4200c1090"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12834,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200c5f40",&ast.Ident/*struct*/{
        NamePos: 12836,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ca080",&ast.FuncType/*struct*/{
        Func: 12815,
        Params: p.StrmapFieldList("0xc4200c7140",&ast.FieldList/*struct*/{
          Opening: 12839,
          Closing: 12840,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c7170",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3200",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200c5fa0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200c5f60",&ast.Ident/*struct*/{
                  NamePos: 12842,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200c5f80",&ast.Ident/*struct*/{
                  NamePos: 12848,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c71d0",&ast.BlockStmt/*struct*/{
        Lbrace: 12856,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ca060",&ast.ReturnStmt/*struct*/{
            Return: 12858,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c3240",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ca040",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ca000",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200c5fc0",&ast.Ident/*struct*/{
                      NamePos: 12865,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8190"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200c5fe0",&ast.Ident/*struct*/{
                      NamePos: 12867,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ca020",&ast.Ident/*struct*/{
                    NamePos: 12869,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 12872,
                Ellipsis: 0,
                Rparen: 12873,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12875,
      }/*struct*/),
    }/*struct*/),/* slice_item: 39*/40: p.StrmapFuncDecl("0xc4200c7350",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c7260",&ast.FieldList/*struct*/{
        Opening: 12882,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3280",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ca0a0",&ast.Ident/*struct*/{
                NamePos: 12883,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c81e0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3280"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ca0e0",&ast.StarExpr/*struct*/{
              Star: 12885,
              X: p.StrmapIdent("0xc4200ca0c0",&ast.Ident/*struct*/{
                NamePos: 12886,
                Name: "KeyValueExpr",
                Obj: p.PtrmapObject("0xc4200c1220"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12898,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ca100",&ast.Ident/*struct*/{
        NamePos: 12900,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ca240",&ast.FuncType/*struct*/{
        Func: 12877,
        Params: p.StrmapFieldList("0xc4200c7290",&ast.FieldList/*struct*/{
          Opening: 12903,
          Closing: 12904,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c72c0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c32c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ca160",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ca120",&ast.Ident/*struct*/{
                  NamePos: 12906,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ca140",&ast.Ident/*struct*/{
                  NamePos: 12912,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c7320",&ast.BlockStmt/*struct*/{
        Lbrace: 12918,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ca220",&ast.ReturnStmt/*struct*/{
            Return: 12920,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c3300",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ca200",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ca1c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ca180",&ast.Ident/*struct*/{
                      NamePos: 12927,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c81e0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ca1a0",&ast.Ident/*struct*/{
                      NamePos: 12929,
                      Name: "Key",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ca1e0",&ast.Ident/*struct*/{
                    NamePos: 12933,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 12936,
                Ellipsis: 0,
                Rparen: 12937,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 12939,
      }/*struct*/),
    }/*struct*/),/* slice_item: 40*/41: p.StrmapFuncDecl("0xc4200c74a0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c73b0",&ast.FieldList/*struct*/{
        Opening: 12946,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3340",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ca260",&ast.Ident/*struct*/{
                NamePos: 12947,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8230",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3340"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ca2a0",&ast.StarExpr/*struct*/{
              Star: 12949,
              X: p.StrmapIdent("0xc4200ca280",&ast.Ident/*struct*/{
                NamePos: 12950,
                Name: "ArrayType",
                Obj: p.PtrmapObject("0xc4200c14f0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 12959,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ca2c0",&ast.Ident/*struct*/{
        NamePos: 12961,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ca3c0",&ast.FuncType/*struct*/{
        Func: 12941,
        Params: p.StrmapFieldList("0xc4200c73e0",&ast.FieldList/*struct*/{
          Opening: 12964,
          Closing: 12965,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c7410",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3380",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ca320",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ca2e0",&ast.Ident/*struct*/{
                  NamePos: 12967,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ca300",&ast.Ident/*struct*/{
                  NamePos: 12973,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c7470",&ast.BlockStmt/*struct*/{
        Lbrace: 12982,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ca3a0",&ast.ReturnStmt/*struct*/{
            Return: 12984,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200ca380",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ca340",&ast.Ident/*struct*/{
                  NamePos: 12991,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c8230"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ca360",&ast.Ident/*struct*/{
                  NamePos: 12993,
                  Name: "Lbrack",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13000,
      }/*struct*/),
    }/*struct*/),/* slice_item: 41*/42: p.StrmapFuncDecl("0xc4200c75f0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c7500",&ast.FieldList/*struct*/{
        Opening: 13007,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c33c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ca3e0",&ast.Ident/*struct*/{
                NamePos: 13008,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8280",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c33c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ca420",&ast.StarExpr/*struct*/{
              Star: 13010,
              X: p.StrmapIdent("0xc4200ca400",&ast.Ident/*struct*/{
                NamePos: 13011,
                Name: "StructType",
                Obj: p.PtrmapObject("0xc4200c1630"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13021,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ca440",&ast.Ident/*struct*/{
        NamePos: 13023,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ca540",&ast.FuncType/*struct*/{
        Func: 13002,
        Params: p.StrmapFieldList("0xc4200c7530",&ast.FieldList/*struct*/{
          Opening: 13026,
          Closing: 13027,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c7560",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3400",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ca4a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ca460",&ast.Ident/*struct*/{
                  NamePos: 13029,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ca480",&ast.Ident/*struct*/{
                  NamePos: 13035,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c75c0",&ast.BlockStmt/*struct*/{
        Lbrace: 13043,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ca520",&ast.ReturnStmt/*struct*/{
            Return: 13045,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200ca500",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ca4c0",&ast.Ident/*struct*/{
                  NamePos: 13052,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c8280"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ca4e0",&ast.Ident/*struct*/{
                  NamePos: 13054,
                  Name: "Struct",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13061,
      }/*struct*/),
    }/*struct*/),/* slice_item: 42*/43: p.StrmapFuncDecl("0xc4200c7830",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c7650",&ast.FieldList/*struct*/{
        Opening: 13068,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3440",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ca560",&ast.Ident/*struct*/{
                NamePos: 13069,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c82d0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3440"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ca5a0",&ast.StarExpr/*struct*/{
              Star: 13071,
              X: p.StrmapIdent("0xc4200ca580",&ast.Ident/*struct*/{
                NamePos: 13072,
                Name: "FuncType",
                Obj: p.PtrmapObject("0xc4200c1770"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13080,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ca5c0",&ast.Ident/*struct*/{
        NamePos: 13082,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ca8e0",&ast.FuncType/*struct*/{
        Func: 13063,
        Params: p.StrmapFieldList("0xc4200c7680",&ast.FieldList/*struct*/{
          Opening: 13085,
          Closing: 13086,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c76b0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3480",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ca620",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ca5e0",&ast.Ident/*struct*/{
                  NamePos: 13088,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ca600",&ast.Ident/*struct*/{
                  NamePos: 13094,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c7800",&ast.BlockStmt/*struct*/{
        Lbrace: 13098,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200c3500",&ast.IfStmt/*struct*/{
            If: 13101,
            Cond: p.StrmapBinaryExpr("0xc4200c7770",&ast.BinaryExpr/*struct*/{
              X: p.StrmapCallExpr("0xc4200c34c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ca6c0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ca680",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ca640",&ast.Ident/*struct*/{
                      NamePos: 13104,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c82d0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ca660",&ast.Ident/*struct*/{
                      NamePos: 13106,
                      Name: "Func",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ca6a0",&ast.Ident/*struct*/{
                    NamePos: 13111,
                    Name: "IsValid",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 13118,
                Ellipsis: 0,
                Rparen: 13119,
              }/*struct*/),
              OpPos: 13121,
              Op: token.Token(35)/*||*/,
              Y: p.StrmapBinaryExpr("0xc4200c7740",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200ca720",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200ca6e0",&ast.Ident/*struct*/{
                    NamePos: 13124,
                    Name: "x",
                    Obj: p.PtrmapObject("0xc4200c82d0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ca700",&ast.Ident/*struct*/{
                    NamePos: 13126,
                    Name: "Params",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 13133,
                Op: token.Token(39)/*==*/,
                Y: p.StrmapIdent("0xc4200ca740",&ast.Ident/*struct*/{
                  NamePos: 13136,
                  Name: "nil",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200c77d0",&ast.BlockStmt/*struct*/{
              Lbrace: 13140,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200ca7e0",&ast.ReturnStmt/*struct*/{
                  Return: 13162,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200ca7c0",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200ca780",&ast.Ident/*struct*/{
                        NamePos: 13169,
                        Name: "x",
                        Obj: p.PtrmapObject("0xc4200c82d0"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200ca7a0",&ast.Ident/*struct*/{
                        NamePos: 13171,
                        Name: "Func",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 13177,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200ca8a0",&ast.ReturnStmt/*struct*/{
            Return: 13180,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c3540",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ca880",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ca840",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ca800",&ast.Ident/*struct*/{
                      NamePos: 13187,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c82d0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ca820",&ast.Ident/*struct*/{
                      NamePos: 13189,
                      Name: "Params",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ca860",&ast.Ident/*struct*/{
                    NamePos: 13196,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 13199,
                Ellipsis: 0,
                Rparen: 13200,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 13258,
      }/*struct*/),
    }/*struct*/),/* slice_item: 43*/44: p.StrmapFuncDecl("0xc4200c7980",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c7890",&ast.FieldList/*struct*/{
        Opening: 13265,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c35c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ca900",&ast.Ident/*struct*/{
                NamePos: 13266,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8320",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c35c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ca940",&ast.StarExpr/*struct*/{
              Star: 13268,
              X: p.StrmapIdent("0xc4200ca920",&ast.Ident/*struct*/{
                NamePos: 13269,
                Name: "InterfaceType",
                Obj: p.PtrmapObject("0xc4200c1900"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13282,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ca960",&ast.Ident/*struct*/{
        NamePos: 13284,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200caa60",&ast.FuncType/*struct*/{
        Func: 13260,
        Params: p.StrmapFieldList("0xc4200c78c0",&ast.FieldList/*struct*/{
          Opening: 13287,
          Closing: 13288,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c78f0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3600",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ca9c0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ca980",&ast.Ident/*struct*/{
                  NamePos: 13290,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ca9a0",&ast.Ident/*struct*/{
                  NamePos: 13296,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c7950",&ast.BlockStmt/*struct*/{
        Lbrace: 13300,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200caa40",&ast.ReturnStmt/*struct*/{
            Return: 13302,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200caa20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ca9e0",&ast.Ident/*struct*/{
                  NamePos: 13309,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c8320"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200caa00",&ast.Ident/*struct*/{
                  NamePos: 13311,
                  Name: "Interface",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13321,
      }/*struct*/),
    }/*struct*/),/* slice_item: 44*/45: p.StrmapFuncDecl("0xc4200c7ad0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c79e0",&ast.FieldList/*struct*/{
        Opening: 13328,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3640",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200caa80",&ast.Ident/*struct*/{
                NamePos: 13329,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8370",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3640"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200caac0",&ast.StarExpr/*struct*/{
              Star: 13331,
              X: p.StrmapIdent("0xc4200caaa0",&ast.Ident/*struct*/{
                NamePos: 13332,
                Name: "MapType",
                Obj: p.PtrmapObject("0xc4200c1a40"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13339,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200caae0",&ast.Ident/*struct*/{
        NamePos: 13341,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cabe0",&ast.FuncType/*struct*/{
        Func: 13323,
        Params: p.StrmapFieldList("0xc4200c7a10",&ast.FieldList/*struct*/{
          Opening: 13344,
          Closing: 13345,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c7a40",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3680",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cab40",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cab00",&ast.Ident/*struct*/{
                  NamePos: 13347,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cab20",&ast.Ident/*struct*/{
                  NamePos: 13353,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c7aa0",&ast.BlockStmt/*struct*/{
        Lbrace: 13363,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cabc0",&ast.ReturnStmt/*struct*/{
            Return: 13365,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200caba0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cab60",&ast.Ident/*struct*/{
                  NamePos: 13372,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c8370"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cab80",&ast.Ident/*struct*/{
                  NamePos: 13374,
                  Name: "Map",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13378,
      }/*struct*/),
    }/*struct*/),/* slice_item: 45*/46: p.StrmapFuncDecl("0xc4200c7c20",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c7b30",&ast.FieldList/*struct*/{
        Opening: 13385,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c36c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cac00",&ast.Ident/*struct*/{
                NamePos: 13386,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c83c0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c36c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cac40",&ast.StarExpr/*struct*/{
              Star: 13388,
              X: p.StrmapIdent("0xc4200cac20",&ast.Ident/*struct*/{
                NamePos: 13389,
                Name: "ChanType",
                Obj: p.PtrmapObject("0xc4200c1b80"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13397,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cac60",&ast.Ident/*struct*/{
        NamePos: 13399,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cad60",&ast.FuncType/*struct*/{
        Func: 13380,
        Params: p.StrmapFieldList("0xc4200c7b60",&ast.FieldList/*struct*/{
          Opening: 13402,
          Closing: 13403,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c7b90",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3700",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cacc0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cac80",&ast.Ident/*struct*/{
                  NamePos: 13405,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200caca0",&ast.Ident/*struct*/{
                  NamePos: 13411,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c7bf0",&ast.BlockStmt/*struct*/{
        Lbrace: 13420,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cad40",&ast.ReturnStmt/*struct*/{
            Return: 13422,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200cad20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cace0",&ast.Ident/*struct*/{
                  NamePos: 13429,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c83c0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cad00",&ast.Ident/*struct*/{
                  NamePos: 13431,
                  Name: "Begin",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13437,
      }/*struct*/),
    }/*struct*/),/* slice_item: 46*/47: p.StrmapFuncDecl("0xc4200c7d70",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c7c80",&ast.FieldList/*struct*/{
        Opening: 13445,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3740",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cad80",&ast.Ident/*struct*/{
                NamePos: 13446,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8410",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3740"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cadc0",&ast.StarExpr/*struct*/{
              Star: 13448,
              X: p.StrmapIdent("0xc4200cada0",&ast.Ident/*struct*/{
                NamePos: 13449,
                Name: "BadExpr",
                Obj: p.PtrmapObject("0xc420051b80"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13456,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cade0",&ast.Ident/*struct*/{
        NamePos: 13458,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200caee0",&ast.FuncType/*struct*/{
        Func: 13440,
        Params: p.StrmapFieldList("0xc4200c7cb0",&ast.FieldList/*struct*/{
          Opening: 13461,
          Closing: 13462,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c7ce0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3780",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cae40",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cae00",&ast.Ident/*struct*/{
                  NamePos: 13464,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cae20",&ast.Ident/*struct*/{
                  NamePos: 13470,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c7d40",&ast.BlockStmt/*struct*/{
        Lbrace: 13474,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200caec0",&ast.ReturnStmt/*struct*/{
            Return: 13476,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200caea0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cae60",&ast.Ident/*struct*/{
                  NamePos: 13483,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c8410"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cae80",&ast.Ident/*struct*/{
                  NamePos: 13485,
                  Name: "To",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13488,
      }/*struct*/),
    }/*struct*/),/* slice_item: 47*/48: p.StrmapFuncDecl("0xc4200c7ef0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c7dd0",&ast.FieldList/*struct*/{
        Opening: 13495,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c37c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200caf00",&ast.Ident/*struct*/{
                NamePos: 13496,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8460",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c37c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200caf40",&ast.StarExpr/*struct*/{
              Star: 13498,
              X: p.StrmapIdent("0xc4200caf20",&ast.Ident/*struct*/{
                NamePos: 13499,
                Name: "Ident",
                Obj: p.PtrmapObject("0xc420051c70"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13504,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200caf60",&ast.Ident/*struct*/{
        NamePos: 13506,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cb160",&ast.FuncType/*struct*/{
        Func: 13490,
        Params: p.StrmapFieldList("0xc4200c7e00",&ast.FieldList/*struct*/{
          Opening: 13509,
          Closing: 13510,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c7e30",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3800",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cafc0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200caf80",&ast.Ident/*struct*/{
                  NamePos: 13512,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cafa0",&ast.Ident/*struct*/{
                  NamePos: 13518,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200c7ec0",&ast.BlockStmt/*struct*/{
        Lbrace: 13524,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cb140",&ast.ReturnStmt/*struct*/{
            Return: 13526,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c38c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200cb020",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200cafe0",&ast.Ident/*struct*/{
                    NamePos: 13533,
                    Name: "token",
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cb000",&ast.Ident/*struct*/{
                    NamePos: 13539,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 13542,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapBinaryExpr("0xc4200c7e90",&ast.BinaryExpr/*struct*/{
                    X: p.StrmapCallExpr("0xc4200c3840",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapIdent("0xc4200cb040",&ast.Ident/*struct*/{
                        NamePos: 13543,
                        Name: "int",
                      }/*struct*/),
                      Lparen: 13546,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapSelectorExpr("0xc4200cb0a0",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200cb060",&ast.Ident/*struct*/{
                            NamePos: 13547,
                            Name: "x",
                            Obj: p.PtrmapObject("0xc4200c8460"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200cb080",&ast.Ident/*struct*/{
                            NamePos: 13549,
                            Name: "NamePos",
                          }/*struct*/),
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 13556,
                    }/*struct*/),
                    OpPos: 13558,
                    Op: token.Token(12)/*+*/,
                    Y: p.StrmapCallExpr("0xc4200c3880",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapIdent("0xc4200cb0c0",&ast.Ident/*struct*/{
                        NamePos: 13560,
                        Name: "len",
                      }/*struct*/),
                      Lparen: 13563,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapSelectorExpr("0xc4200cb120",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200cb0e0",&ast.Ident/*struct*/{
                            NamePos: 13564,
                            Name: "x",
                            Obj: p.PtrmapObject("0xc4200c8460"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200cb100",&ast.Ident/*struct*/{
                            NamePos: 13566,
                            Name: "Name",
                          }/*struct*/),
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 13570,
                    }/*struct*/),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 13571,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13573,
      }/*struct*/),
    }/*struct*/),/* slice_item: 48*/49: p.StrmapFuncDecl("0xc4200cc150",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200c7f50",&ast.FieldList/*struct*/{
        Opening: 13580,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3900",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cb180",&ast.Ident/*struct*/{
                NamePos: 13581,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c84b0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3900"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cb1c0",&ast.StarExpr/*struct*/{
              Star: 13583,
              X: p.StrmapIdent("0xc4200cb1a0",&ast.Ident/*struct*/{
                NamePos: 13584,
                Name: "Ellipsis",
                Obj: p.PtrmapObject("0xc420051db0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13592,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cb1e0",&ast.Ident/*struct*/{
        NamePos: 13594,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cb460",&ast.FuncType/*struct*/{
        Func: 13575,
        Params: p.StrmapFieldList("0xc4200c7f80",&ast.FieldList/*struct*/{
          Opening: 13597,
          Closing: 13598,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200c7fb0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3940",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cb240",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cb200",&ast.Ident/*struct*/{
                  NamePos: 13600,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cb220",&ast.Ident/*struct*/{
                  NamePos: 13606,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cc120",&ast.BlockStmt/*struct*/{
        Lbrace: 13610,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200c39c0",&ast.IfStmt/*struct*/{
            If: 13613,
            Cond: p.StrmapBinaryExpr("0xc4200cc060",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200cb2a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cb260",&ast.Ident/*struct*/{
                  NamePos: 13616,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c84b0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cb280",&ast.Ident/*struct*/{
                  NamePos: 13618,
                  Name: "Elt",
                }/*struct*/),
              }/*struct*/),
              OpPos: 13622,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200cb2c0",&ast.Ident/*struct*/{
                NamePos: 13625,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200cc0c0",&ast.BlockStmt/*struct*/{
              Lbrace: 13629,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200cb380",&ast.ReturnStmt/*struct*/{
                  Return: 13633,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200c3980",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200cb360",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapSelectorExpr("0xc4200cb320",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200cb2e0",&ast.Ident/*struct*/{
                            NamePos: 13640,
                            Name: "x",
                            Obj: p.PtrmapObject("0xc4200c84b0"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200cb300",&ast.Ident/*struct*/{
                            NamePos: 13642,
                            Name: "Elt",
                          }/*struct*/),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200cb340",&ast.Ident/*struct*/{
                          NamePos: 13646,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 13649,
                      Ellipsis: 0,
                      Rparen: 13650,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 13653,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200cb420",&ast.ReturnStmt/*struct*/{
            Return: 13656,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200cc0f0",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200cb3e0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200cb3a0",&ast.Ident/*struct*/{
                    NamePos: 13663,
                    Name: "x",
                    Obj: p.PtrmapObject("0xc4200c84b0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cb3c0",&ast.Ident/*struct*/{
                    NamePos: 13665,
                    Name: "Ellipsis",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 13674,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200cb400",&ast.BasicLit/*struct*/{
                  ValuePos: 13676,
                  Kind: token.Token(5)/*INT*/,
                  Value: "3",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 13692,
      }/*struct*/),
    }/*struct*/),/* slice_item: 49*/50: p.StrmapFuncDecl("0xc4200cc2d0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cc1b0",&ast.FieldList/*struct*/{
        Opening: 13699,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3a00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cb480",&ast.Ident/*struct*/{
                NamePos: 13700,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8500",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3a00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cb4c0",&ast.StarExpr/*struct*/{
              Star: 13702,
              X: p.StrmapIdent("0xc4200cb4a0",&ast.Ident/*struct*/{
                NamePos: 13703,
                Name: "BasicLit",
                Obj: p.PtrmapObject("0xc420051ea0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13711,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cb4e0",&ast.Ident/*struct*/{
        NamePos: 13713,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cb6e0",&ast.FuncType/*struct*/{
        Func: 13694,
        Params: p.StrmapFieldList("0xc4200cc1e0",&ast.FieldList/*struct*/{
          Opening: 13716,
          Closing: 13717,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cc210",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3a40",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cb540",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cb500",&ast.Ident/*struct*/{
                  NamePos: 13719,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cb520",&ast.Ident/*struct*/{
                  NamePos: 13725,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cc2a0",&ast.BlockStmt/*struct*/{
        Lbrace: 13735,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cb6c0",&ast.ReturnStmt/*struct*/{
            Return: 13737,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c3b00",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200cb5a0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200cb560",&ast.Ident/*struct*/{
                    NamePos: 13744,
                    Name: "token",
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cb580",&ast.Ident/*struct*/{
                    NamePos: 13750,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 13753,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapBinaryExpr("0xc4200cc270",&ast.BinaryExpr/*struct*/{
                    X: p.StrmapCallExpr("0xc4200c3a80",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapIdent("0xc4200cb5c0",&ast.Ident/*struct*/{
                        NamePos: 13754,
                        Name: "int",
                      }/*struct*/),
                      Lparen: 13757,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapSelectorExpr("0xc4200cb620",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200cb5e0",&ast.Ident/*struct*/{
                            NamePos: 13758,
                            Name: "x",
                            Obj: p.PtrmapObject("0xc4200c8500"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200cb600",&ast.Ident/*struct*/{
                            NamePos: 13760,
                            Name: "ValuePos",
                          }/*struct*/),
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 13768,
                    }/*struct*/),
                    OpPos: 13770,
                    Op: token.Token(12)/*+*/,
                    Y: p.StrmapCallExpr("0xc4200c3ac0",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapIdent("0xc4200cb640",&ast.Ident/*struct*/{
                        NamePos: 13772,
                        Name: "len",
                      }/*struct*/),
                      Lparen: 13775,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapSelectorExpr("0xc4200cb6a0",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200cb660",&ast.Ident/*struct*/{
                            NamePos: 13776,
                            Name: "x",
                            Obj: p.PtrmapObject("0xc4200c8500"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200cb680",&ast.Ident/*struct*/{
                            NamePos: 13778,
                            Name: "Value",
                          }/*struct*/),
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 13783,
                    }/*struct*/),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 13784,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13786,
      }/*struct*/),
    }/*struct*/),/* slice_item: 50*/51: p.StrmapFuncDecl("0xc4200cc420",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cc330",&ast.FieldList/*struct*/{
        Opening: 13793,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3b40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cb700",&ast.Ident/*struct*/{
                NamePos: 13794,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8550",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3b40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cb740",&ast.StarExpr/*struct*/{
              Star: 13796,
              X: p.StrmapIdent("0xc4200cb720",&ast.Ident/*struct*/{
                NamePos: 13797,
                Name: "FuncLit",
                Obj: p.PtrmapObject("0xc4200c0050"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13804,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cb760",&ast.Ident/*struct*/{
        NamePos: 13806,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cb8a0",&ast.FuncType/*struct*/{
        Func: 13788,
        Params: p.StrmapFieldList("0xc4200cc360",&ast.FieldList/*struct*/{
          Opening: 13809,
          Closing: 13810,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cc390",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3b80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cb7c0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cb780",&ast.Ident/*struct*/{
                  NamePos: 13812,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cb7a0",&ast.Ident/*struct*/{
                  NamePos: 13818,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cc3f0",&ast.BlockStmt/*struct*/{
        Lbrace: 13829,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cb880",&ast.ReturnStmt/*struct*/{
            Return: 13831,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c3bc0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200cb860",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200cb820",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200cb7e0",&ast.Ident/*struct*/{
                      NamePos: 13838,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8550"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200cb800",&ast.Ident/*struct*/{
                      NamePos: 13840,
                      Name: "Body",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cb840",&ast.Ident/*struct*/{
                    NamePos: 13845,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 13848,
                Ellipsis: 0,
                Rparen: 13849,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13851,
      }/*struct*/),
    }/*struct*/),/* slice_item: 51*/52: p.StrmapFuncDecl("0xc4200cc5a0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cc480",&ast.FieldList/*struct*/{
        Opening: 13858,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3c00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cb8c0",&ast.Ident/*struct*/{
                NamePos: 13859,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c85a0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3c00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cb900",&ast.StarExpr/*struct*/{
              Star: 13861,
              X: p.StrmapIdent("0xc4200cb8e0",&ast.Ident/*struct*/{
                NamePos: 13862,
                Name: "CompositeLit",
                Obj: p.PtrmapObject("0xc4200c0140"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13874,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cb920",&ast.Ident/*struct*/{
        NamePos: 13876,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cba40",&ast.FuncType/*struct*/{
        Func: 13853,
        Params: p.StrmapFieldList("0xc4200cc4b0",&ast.FieldList/*struct*/{
          Opening: 13879,
          Closing: 13880,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cc4e0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3c40",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cb980",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cb940",&ast.Ident/*struct*/{
                  NamePos: 13882,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cb960",&ast.Ident/*struct*/{
                  NamePos: 13888,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cc570",&ast.BlockStmt/*struct*/{
        Lbrace: 13894,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cba20",&ast.ReturnStmt/*struct*/{
            Return: 13896,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200cc540",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200cb9e0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200cb9a0",&ast.Ident/*struct*/{
                    NamePos: 13903,
                    Name: "x",
                    Obj: p.PtrmapObject("0xc4200c85a0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cb9c0",&ast.Ident/*struct*/{
                    NamePos: 13905,
                    Name: "Rbrace",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 13912,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200cba00",&ast.BasicLit/*struct*/{
                  ValuePos: 13914,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13916,
      }/*struct*/),
    }/*struct*/),/* slice_item: 52*/53: p.StrmapFuncDecl("0xc4200cc720",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cc600",&ast.FieldList/*struct*/{
        Opening: 13923,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3c80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cba60",&ast.Ident/*struct*/{
                NamePos: 13924,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c85f0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3c80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cbaa0",&ast.StarExpr/*struct*/{
              Star: 13926,
              X: p.StrmapIdent("0xc4200cba80",&ast.Ident/*struct*/{
                NamePos: 13927,
                Name: "ParenExpr",
                Obj: p.PtrmapObject("0xc4200c02d0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 13936,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cbac0",&ast.Ident/*struct*/{
        NamePos: 13938,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cbbe0",&ast.FuncType/*struct*/{
        Func: 13918,
        Params: p.StrmapFieldList("0xc4200cc630",&ast.FieldList/*struct*/{
          Opening: 13941,
          Closing: 13942,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cc660",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3cc0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cbb20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cbae0",&ast.Ident/*struct*/{
                  NamePos: 13944,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cbb00",&ast.Ident/*struct*/{
                  NamePos: 13950,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cc6f0",&ast.BlockStmt/*struct*/{
        Lbrace: 13959,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cbbc0",&ast.ReturnStmt/*struct*/{
            Return: 13961,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200cc6c0",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200cbb80",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200cbb40",&ast.Ident/*struct*/{
                    NamePos: 13968,
                    Name: "x",
                    Obj: p.PtrmapObject("0xc4200c85f0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cbb60",&ast.Ident/*struct*/{
                    NamePos: 13970,
                    Name: "Rparen",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 13977,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200cbba0",&ast.BasicLit/*struct*/{
                  ValuePos: 13979,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 13981,
      }/*struct*/),
    }/*struct*/),/* slice_item: 53*/54: p.StrmapFuncDecl("0xc4200cc870",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cc780",&ast.FieldList/*struct*/{
        Opening: 13988,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3d00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cbc00",&ast.Ident/*struct*/{
                NamePos: 13989,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8640",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3d00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cbc40",&ast.StarExpr/*struct*/{
              Star: 13991,
              X: p.StrmapIdent("0xc4200cbc20",&ast.Ident/*struct*/{
                NamePos: 13992,
                Name: "SelectorExpr",
                Obj: p.PtrmapObject("0xc4200c0460"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14004,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cbc60",&ast.Ident/*struct*/{
        NamePos: 14006,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cbda0",&ast.FuncType/*struct*/{
        Func: 13983,
        Params: p.StrmapFieldList("0xc4200cc7b0",&ast.FieldList/*struct*/{
          Opening: 14009,
          Closing: 14010,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cc7e0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3d40",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cbcc0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cbc80",&ast.Ident/*struct*/{
                  NamePos: 14012,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cbca0",&ast.Ident/*struct*/{
                  NamePos: 14018,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cc840",&ast.BlockStmt/*struct*/{
        Lbrace: 14024,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cbd80",&ast.ReturnStmt/*struct*/{
            Return: 14026,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200c3d80",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200cbd60",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200cbd20",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200cbce0",&ast.Ident/*struct*/{
                      NamePos: 14033,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8640"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200cbd00",&ast.Ident/*struct*/{
                      NamePos: 14035,
                      Name: "Sel",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cbd40",&ast.Ident/*struct*/{
                    NamePos: 14039,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14042,
                Ellipsis: 0,
                Rparen: 14043,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14045,
      }/*struct*/),
    }/*struct*/),/* slice_item: 54*/55: p.StrmapFuncDecl("0xc4200cc9f0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cc8d0",&ast.FieldList/*struct*/{
        Opening: 14052,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3dc0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cbdc0",&ast.Ident/*struct*/{
                NamePos: 14053,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8690",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3dc0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cbe00",&ast.StarExpr/*struct*/{
              Star: 14055,
              X: p.StrmapIdent("0xc4200cbde0",&ast.Ident/*struct*/{
                NamePos: 14056,
                Name: "IndexExpr",
                Obj: p.PtrmapObject("0xc4200c05a0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14065,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cbe20",&ast.Ident/*struct*/{
        NamePos: 14067,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cbf40",&ast.FuncType/*struct*/{
        Func: 14047,
        Params: p.StrmapFieldList("0xc4200cc900",&ast.FieldList/*struct*/{
          Opening: 14070,
          Closing: 14071,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cc930",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3e00",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cbe80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cbe40",&ast.Ident/*struct*/{
                  NamePos: 14073,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cbe60",&ast.Ident/*struct*/{
                  NamePos: 14079,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cc9c0",&ast.BlockStmt/*struct*/{
        Lbrace: 14088,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cbf20",&ast.ReturnStmt/*struct*/{
            Return: 14090,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200cc990",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200cbee0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200cbea0",&ast.Ident/*struct*/{
                    NamePos: 14097,
                    Name: "x",
                    Obj: p.PtrmapObject("0xc4200c8690"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cbec0",&ast.Ident/*struct*/{
                    NamePos: 14099,
                    Name: "Rbrack",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 14106,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200cbf00",&ast.BasicLit/*struct*/{
                  ValuePos: 14108,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14110,
      }/*struct*/),
    }/*struct*/),/* slice_item: 55*/56: p.StrmapFuncDecl("0xc4200ccb70",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cca50",&ast.FieldList/*struct*/{
        Opening: 14117,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3e40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cbf60",&ast.Ident/*struct*/{
                NamePos: 14118,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c86e0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3e40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cbfa0",&ast.StarExpr/*struct*/{
              Star: 14120,
              X: p.StrmapIdent("0xc4200cbf80",&ast.Ident/*struct*/{
                NamePos: 14121,
                Name: "SliceExpr",
                Obj: p.PtrmapObject("0xc4200c0780"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14130,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cbfc0",&ast.Ident/*struct*/{
        NamePos: 14132,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ce0e0",&ast.FuncType/*struct*/{
        Func: 14112,
        Params: p.StrmapFieldList("0xc4200cca80",&ast.FieldList/*struct*/{
          Opening: 14135,
          Closing: 14136,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200ccab0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3e80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ce020",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cbfe0",&ast.Ident/*struct*/{
                  NamePos: 14138,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ce000",&ast.Ident/*struct*/{
                  NamePos: 14144,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200ccb40",&ast.BlockStmt/*struct*/{
        Lbrace: 14153,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ce0c0",&ast.ReturnStmt/*struct*/{
            Return: 14155,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200ccb10",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200ce080",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200ce040",&ast.Ident/*struct*/{
                    NamePos: 14162,
                    Name: "x",
                    Obj: p.PtrmapObject("0xc4200c86e0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ce060",&ast.Ident/*struct*/{
                    NamePos: 14164,
                    Name: "Rbrack",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 14171,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200ce0a0",&ast.BasicLit/*struct*/{
                  ValuePos: 14173,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14175,
      }/*struct*/),
    }/*struct*/),/* slice_item: 56*/57: p.StrmapFuncDecl("0xc4200cccf0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200ccbd0",&ast.FieldList/*struct*/{
        Opening: 14182,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3ec0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ce100",&ast.Ident/*struct*/{
                NamePos: 14183,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8730",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3ec0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ce140",&ast.StarExpr/*struct*/{
              Star: 14185,
              X: p.StrmapIdent("0xc4200ce120",&ast.Ident/*struct*/{
                NamePos: 14186,
                Name: "TypeAssertExpr",
                Obj: p.PtrmapObject("0xc4200c0a00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14200,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ce160",&ast.Ident/*struct*/{
        NamePos: 14202,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ce280",&ast.FuncType/*struct*/{
        Func: 14177,
        Params: p.StrmapFieldList("0xc4200ccc00",&ast.FieldList/*struct*/{
          Opening: 14205,
          Closing: 14206,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200ccc30",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3f00",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ce1c0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ce180",&ast.Ident/*struct*/{
                  NamePos: 14208,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ce1a0",&ast.Ident/*struct*/{
                  NamePos: 14214,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cccc0",&ast.BlockStmt/*struct*/{
        Lbrace: 14218,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ce260",&ast.ReturnStmt/*struct*/{
            Return: 14220,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200ccc90",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200ce220",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200ce1e0",&ast.Ident/*struct*/{
                    NamePos: 14227,
                    Name: "x",
                    Obj: p.PtrmapObject("0xc4200c8730"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ce200",&ast.Ident/*struct*/{
                    NamePos: 14229,
                    Name: "Rparen",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 14236,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200ce240",&ast.BasicLit/*struct*/{
                  ValuePos: 14238,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14240,
      }/*struct*/),
    }/*struct*/),/* slice_item: 57*/58: p.StrmapFuncDecl("0xc4200cce70",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200ccd50",&ast.FieldList/*struct*/{
        Opening: 14247,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3f40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ce2a0",&ast.Ident/*struct*/{
                NamePos: 14248,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8780",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3f40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ce2e0",&ast.StarExpr/*struct*/{
              Star: 14250,
              X: p.StrmapIdent("0xc4200ce2c0",&ast.Ident/*struct*/{
                NamePos: 14251,
                Name: "CallExpr",
                Obj: p.PtrmapObject("0xc4200c0be0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14259,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ce300",&ast.Ident/*struct*/{
        NamePos: 14261,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ce420",&ast.FuncType/*struct*/{
        Func: 14242,
        Params: p.StrmapFieldList("0xc4200ccd80",&ast.FieldList/*struct*/{
          Opening: 14264,
          Closing: 14265,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200ccdb0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200c3f80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ce360",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ce320",&ast.Ident/*struct*/{
                  NamePos: 14267,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ce340",&ast.Ident/*struct*/{
                  NamePos: 14273,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cce40",&ast.BlockStmt/*struct*/{
        Lbrace: 14283,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ce400",&ast.ReturnStmt/*struct*/{
            Return: 14285,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200cce10",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200ce3c0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200ce380",&ast.Ident/*struct*/{
                    NamePos: 14292,
                    Name: "x",
                    Obj: p.PtrmapObject("0xc4200c8780"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ce3a0",&ast.Ident/*struct*/{
                    NamePos: 14294,
                    Name: "Rparen",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 14301,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200ce3e0",&ast.BasicLit/*struct*/{
                  ValuePos: 14303,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14305,
      }/*struct*/),
    }/*struct*/),/* slice_item: 58*/59: p.StrmapFuncDecl("0xc4200ccfc0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cced0",&ast.FieldList/*struct*/{
        Opening: 14312,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200c3fc0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ce440",&ast.Ident/*struct*/{
                NamePos: 14313,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c87d0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200c3fc0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ce480",&ast.StarExpr/*struct*/{
              Star: 14315,
              X: p.StrmapIdent("0xc4200ce460",&ast.Ident/*struct*/{
                NamePos: 14316,
                Name: "StarExpr",
                Obj: p.PtrmapObject("0xc4200c0e60"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14324,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ce4a0",&ast.Ident/*struct*/{
        NamePos: 14326,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ce5e0",&ast.FuncType/*struct*/{
        Func: 14307,
        Params: p.StrmapFieldList("0xc4200ccf00",&ast.FieldList/*struct*/{
          Opening: 14329,
          Closing: 14330,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200ccf30",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0000",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ce500",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ce4c0",&ast.Ident/*struct*/{
                  NamePos: 14332,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ce4e0",&ast.Ident/*struct*/{
                  NamePos: 14338,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200ccf90",&ast.BlockStmt/*struct*/{
        Lbrace: 14348,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ce5c0",&ast.ReturnStmt/*struct*/{
            Return: 14350,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d0040",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ce5a0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ce560",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ce520",&ast.Ident/*struct*/{
                      NamePos: 14357,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c87d0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ce540",&ast.Ident/*struct*/{
                      NamePos: 14359,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ce580",&ast.Ident/*struct*/{
                    NamePos: 14361,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14364,
                Ellipsis: 0,
                Rparen: 14365,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14367,
      }/*struct*/),
    }/*struct*/),/* slice_item: 59*/60: p.StrmapFuncDecl("0xc4200cd110",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cd020",&ast.FieldList/*struct*/{
        Opening: 14374,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0080",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ce600",&ast.Ident/*struct*/{
                NamePos: 14375,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8820",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200d0080"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ce640",&ast.StarExpr/*struct*/{
              Star: 14377,
              X: p.StrmapIdent("0xc4200ce620",&ast.Ident/*struct*/{
                NamePos: 14378,
                Name: "UnaryExpr",
                Obj: p.PtrmapObject("0xc4200c0f50"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14387,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ce660",&ast.Ident/*struct*/{
        NamePos: 14389,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ce7a0",&ast.FuncType/*struct*/{
        Func: 14369,
        Params: p.StrmapFieldList("0xc4200cd050",&ast.FieldList/*struct*/{
          Opening: 14392,
          Closing: 14393,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cd080",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d00c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ce6c0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ce680",&ast.Ident/*struct*/{
                  NamePos: 14395,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ce6a0",&ast.Ident/*struct*/{
                  NamePos: 14401,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cd0e0",&ast.BlockStmt/*struct*/{
        Lbrace: 14410,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ce780",&ast.ReturnStmt/*struct*/{
            Return: 14412,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d0100",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ce760",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ce720",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ce6e0",&ast.Ident/*struct*/{
                      NamePos: 14419,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8820"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ce700",&ast.Ident/*struct*/{
                      NamePos: 14421,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ce740",&ast.Ident/*struct*/{
                    NamePos: 14423,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14426,
                Ellipsis: 0,
                Rparen: 14427,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14429,
      }/*struct*/),
    }/*struct*/),/* slice_item: 60*/61: p.StrmapFuncDecl("0xc4200cd260",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cd170",&ast.FieldList/*struct*/{
        Opening: 14436,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0140",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ce7c0",&ast.Ident/*struct*/{
                NamePos: 14437,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8870",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200d0140"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ce800",&ast.StarExpr/*struct*/{
              Star: 14439,
              X: p.StrmapIdent("0xc4200ce7e0",&ast.Ident/*struct*/{
                NamePos: 14440,
                Name: "BinaryExpr",
                Obj: p.PtrmapObject("0xc4200c1090"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14450,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ce820",&ast.Ident/*struct*/{
        NamePos: 14452,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ce960",&ast.FuncType/*struct*/{
        Func: 14431,
        Params: p.StrmapFieldList("0xc4200cd1a0",&ast.FieldList/*struct*/{
          Opening: 14455,
          Closing: 14456,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cd1d0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0180",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ce880",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ce840",&ast.Ident/*struct*/{
                  NamePos: 14458,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ce860",&ast.Ident/*struct*/{
                  NamePos: 14464,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cd230",&ast.BlockStmt/*struct*/{
        Lbrace: 14472,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ce940",&ast.ReturnStmt/*struct*/{
            Return: 14474,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d01c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ce920",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ce8e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ce8a0",&ast.Ident/*struct*/{
                      NamePos: 14481,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8870"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ce8c0",&ast.Ident/*struct*/{
                      NamePos: 14483,
                      Name: "Y",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ce900",&ast.Ident/*struct*/{
                    NamePos: 14485,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14488,
                Ellipsis: 0,
                Rparen: 14489,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14491,
      }/*struct*/),
    }/*struct*/),/* slice_item: 61*/62: p.StrmapFuncDecl("0xc4200cd3b0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cd2c0",&ast.FieldList/*struct*/{
        Opening: 14498,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0200",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ce980",&ast.Ident/*struct*/{
                NamePos: 14499,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c88c0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200d0200"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ce9c0",&ast.StarExpr/*struct*/{
              Star: 14501,
              X: p.StrmapIdent("0xc4200ce9a0",&ast.Ident/*struct*/{
                NamePos: 14502,
                Name: "KeyValueExpr",
                Obj: p.PtrmapObject("0xc4200c1220"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14514,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ce9e0",&ast.Ident/*struct*/{
        NamePos: 14516,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ceb20",&ast.FuncType/*struct*/{
        Func: 14493,
        Params: p.StrmapFieldList("0xc4200cd2f0",&ast.FieldList/*struct*/{
          Opening: 14519,
          Closing: 14520,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cd320",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0240",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cea40",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cea00",&ast.Ident/*struct*/{
                  NamePos: 14522,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cea20",&ast.Ident/*struct*/{
                  NamePos: 14528,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cd380",&ast.BlockStmt/*struct*/{
        Lbrace: 14534,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ceb00",&ast.ReturnStmt/*struct*/{
            Return: 14536,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d0280",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ceae0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ceaa0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200cea60",&ast.Ident/*struct*/{
                      NamePos: 14543,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c88c0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200cea80",&ast.Ident/*struct*/{
                      NamePos: 14545,
                      Name: "Value",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ceac0",&ast.Ident/*struct*/{
                    NamePos: 14551,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14554,
                Ellipsis: 0,
                Rparen: 14555,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14557,
      }/*struct*/),
    }/*struct*/),/* slice_item: 62*/63: p.StrmapFuncDecl("0xc4200cd500",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cd410",&ast.FieldList/*struct*/{
        Opening: 14564,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d02c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ceb40",&ast.Ident/*struct*/{
                NamePos: 14565,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8910",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200d02c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ceb80",&ast.StarExpr/*struct*/{
              Star: 14567,
              X: p.StrmapIdent("0xc4200ceb60",&ast.Ident/*struct*/{
                NamePos: 14568,
                Name: "ArrayType",
                Obj: p.PtrmapObject("0xc4200c14f0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14577,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ceba0",&ast.Ident/*struct*/{
        NamePos: 14579,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cece0",&ast.FuncType/*struct*/{
        Func: 14559,
        Params: p.StrmapFieldList("0xc4200cd440",&ast.FieldList/*struct*/{
          Opening: 14582,
          Closing: 14583,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cd470",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0300",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cec00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cebc0",&ast.Ident/*struct*/{
                  NamePos: 14585,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cebe0",&ast.Ident/*struct*/{
                  NamePos: 14591,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cd4d0",&ast.BlockStmt/*struct*/{
        Lbrace: 14600,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cecc0",&ast.ReturnStmt/*struct*/{
            Return: 14602,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d0340",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ceca0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200cec60",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200cec20",&ast.Ident/*struct*/{
                      NamePos: 14609,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8910"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200cec40",&ast.Ident/*struct*/{
                      NamePos: 14611,
                      Name: "Elt",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cec80",&ast.Ident/*struct*/{
                    NamePos: 14615,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14618,
                Ellipsis: 0,
                Rparen: 14619,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14621,
      }/*struct*/),
    }/*struct*/),/* slice_item: 63*/64: p.StrmapFuncDecl("0xc4200cd650",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cd560",&ast.FieldList/*struct*/{
        Opening: 14628,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0380",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ced00",&ast.Ident/*struct*/{
                NamePos: 14629,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8960",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200d0380"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ced40",&ast.StarExpr/*struct*/{
              Star: 14631,
              X: p.StrmapIdent("0xc4200ced20",&ast.Ident/*struct*/{
                NamePos: 14632,
                Name: "StructType",
                Obj: p.PtrmapObject("0xc4200c1630"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14642,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ced60",&ast.Ident/*struct*/{
        NamePos: 14644,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ceea0",&ast.FuncType/*struct*/{
        Func: 14623,
        Params: p.StrmapFieldList("0xc4200cd590",&ast.FieldList/*struct*/{
          Opening: 14647,
          Closing: 14648,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cd5c0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d03c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cedc0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ced80",&ast.Ident/*struct*/{
                  NamePos: 14650,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ceda0",&ast.Ident/*struct*/{
                  NamePos: 14656,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cd620",&ast.BlockStmt/*struct*/{
        Lbrace: 14664,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cee80",&ast.ReturnStmt/*struct*/{
            Return: 14666,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d0400",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200cee60",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200cee20",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200cede0",&ast.Ident/*struct*/{
                      NamePos: 14673,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8960"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200cee00",&ast.Ident/*struct*/{
                      NamePos: 14675,
                      Name: "Fields",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cee40",&ast.Ident/*struct*/{
                    NamePos: 14682,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14685,
                Ellipsis: 0,
                Rparen: 14686,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14688,
      }/*struct*/),
    }/*struct*/),/* slice_item: 64*/65: p.StrmapFuncDecl("0xc4200cd860",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cd6b0",&ast.FieldList/*struct*/{
        Opening: 14695,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0440",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ceec0",&ast.Ident/*struct*/{
                NamePos: 14696,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c89b0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200d0440"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cef00",&ast.StarExpr/*struct*/{
              Star: 14698,
              X: p.StrmapIdent("0xc4200ceee0",&ast.Ident/*struct*/{
                NamePos: 14699,
                Name: "FuncType",
                Obj: p.PtrmapObject("0xc4200c1770"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14707,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cef20",&ast.Ident/*struct*/{
        NamePos: 14709,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cf1c0",&ast.FuncType/*struct*/{
        Func: 14690,
        Params: p.StrmapFieldList("0xc4200cd6e0",&ast.FieldList/*struct*/{
          Opening: 14712,
          Closing: 14713,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cd710",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0480",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cef80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cef40",&ast.Ident/*struct*/{
                  NamePos: 14715,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cef60",&ast.Ident/*struct*/{
                  NamePos: 14721,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cd830",&ast.BlockStmt/*struct*/{
        Lbrace: 14725,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200d0500",&ast.IfStmt/*struct*/{
            If: 14728,
            Cond: p.StrmapBinaryExpr("0xc4200cd7a0",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200cefe0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cefa0",&ast.Ident/*struct*/{
                  NamePos: 14731,
                  Name: "x",
                  Obj: p.PtrmapObject("0xc4200c89b0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cefc0",&ast.Ident/*struct*/{
                  NamePos: 14733,
                  Name: "Results",
                }/*struct*/),
              }/*struct*/),
              OpPos: 14741,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200cf000",&ast.Ident/*struct*/{
                NamePos: 14744,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200cd800",&ast.BlockStmt/*struct*/{
              Lbrace: 14748,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200cf0c0",&ast.ReturnStmt/*struct*/{
                  Return: 14752,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200d04c0",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200cf0a0",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapSelectorExpr("0xc4200cf060",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200cf020",&ast.Ident/*struct*/{
                            NamePos: 14759,
                            Name: "x",
                            Obj: p.PtrmapObject("0xc4200c89b0"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200cf040",&ast.Ident/*struct*/{
                            NamePos: 14761,
                            Name: "Results",
                          }/*struct*/),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200cf080",&ast.Ident/*struct*/{
                          NamePos: 14769,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 14772,
                      Ellipsis: 0,
                      Rparen: 14773,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 14776,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200cf180",&ast.ReturnStmt/*struct*/{
            Return: 14779,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d0540",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200cf160",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200cf120",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200cf0e0",&ast.Ident/*struct*/{
                      NamePos: 14786,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c89b0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200cf100",&ast.Ident/*struct*/{
                      NamePos: 14788,
                      Name: "Params",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cf140",&ast.Ident/*struct*/{
                    NamePos: 14795,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14798,
                Ellipsis: 0,
                Rparen: 14799,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 14801,
      }/*struct*/),
    }/*struct*/),/* slice_item: 65*/66: p.StrmapFuncDecl("0xc4200cd9b0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cd8c0",&ast.FieldList/*struct*/{
        Opening: 14808,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0580",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cf1e0",&ast.Ident/*struct*/{
                NamePos: 14809,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8a00",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200d0580"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cf220",&ast.StarExpr/*struct*/{
              Star: 14811,
              X: p.StrmapIdent("0xc4200cf200",&ast.Ident/*struct*/{
                NamePos: 14812,
                Name: "InterfaceType",
                Obj: p.PtrmapObject("0xc4200c1900"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14825,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cf240",&ast.Ident/*struct*/{
        NamePos: 14827,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cf380",&ast.FuncType/*struct*/{
        Func: 14803,
        Params: p.StrmapFieldList("0xc4200cd8f0",&ast.FieldList/*struct*/{
          Opening: 14830,
          Closing: 14831,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cd920",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d05c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cf2a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cf260",&ast.Ident/*struct*/{
                  NamePos: 14833,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cf280",&ast.Ident/*struct*/{
                  NamePos: 14839,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cd980",&ast.BlockStmt/*struct*/{
        Lbrace: 14843,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cf360",&ast.ReturnStmt/*struct*/{
            Return: 14845,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d0600",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200cf340",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200cf300",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200cf2c0",&ast.Ident/*struct*/{
                      NamePos: 14852,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8a00"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200cf2e0",&ast.Ident/*struct*/{
                      NamePos: 14854,
                      Name: "Methods",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cf320",&ast.Ident/*struct*/{
                    NamePos: 14862,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14865,
                Ellipsis: 0,
                Rparen: 14866,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14868,
      }/*struct*/),
    }/*struct*/),/* slice_item: 66*/67: p.StrmapFuncDecl("0xc4200cdb00",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cda10",&ast.FieldList/*struct*/{
        Opening: 14875,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0640",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cf3a0",&ast.Ident/*struct*/{
                NamePos: 14876,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8a50",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200d0640"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cf3e0",&ast.StarExpr/*struct*/{
              Star: 14878,
              X: p.StrmapIdent("0xc4200cf3c0",&ast.Ident/*struct*/{
                NamePos: 14879,
                Name: "MapType",
                Obj: p.PtrmapObject("0xc4200c1a40"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14886,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cf400",&ast.Ident/*struct*/{
        NamePos: 14888,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cf540",&ast.FuncType/*struct*/{
        Func: 14870,
        Params: p.StrmapFieldList("0xc4200cda40",&ast.FieldList/*struct*/{
          Opening: 14891,
          Closing: 14892,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cda70",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0680",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cf460",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cf420",&ast.Ident/*struct*/{
                  NamePos: 14894,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cf440",&ast.Ident/*struct*/{
                  NamePos: 14900,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cdad0",&ast.BlockStmt/*struct*/{
        Lbrace: 14910,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cf520",&ast.ReturnStmt/*struct*/{
            Return: 14912,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d06c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200cf500",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200cf4c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200cf480",&ast.Ident/*struct*/{
                      NamePos: 14919,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8a50"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200cf4a0",&ast.Ident/*struct*/{
                      NamePos: 14921,
                      Name: "Value",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cf4e0",&ast.Ident/*struct*/{
                    NamePos: 14927,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14930,
                Ellipsis: 0,
                Rparen: 14931,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14933,
      }/*struct*/),
    }/*struct*/),/* slice_item: 67*/68: p.StrmapFuncDecl("0xc4200cdc50",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cdb60",&ast.FieldList/*struct*/{
        Opening: 14940,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0700",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200cf560",&ast.Ident/*struct*/{
                NamePos: 14941,
                Name: "x",
                Obj: p.StrmapObject("0xc4200c8aa0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "x",
                  Decl: p.PtrmapField("0xc4200d0700"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200cf5a0",&ast.StarExpr/*struct*/{
              Star: 14943,
              X: p.StrmapIdent("0xc4200cf580",&ast.Ident/*struct*/{
                NamePos: 14944,
                Name: "ChanType",
                Obj: p.PtrmapObject("0xc4200c1b80"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 14952,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cf5c0",&ast.Ident/*struct*/{
        NamePos: 14954,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cf720",&ast.FuncType/*struct*/{
        Func: 14935,
        Params: p.StrmapFieldList("0xc4200cdb90",&ast.FieldList/*struct*/{
          Opening: 14957,
          Closing: 14958,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200cdbc0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0740",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200cf620",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200cf5e0",&ast.Ident/*struct*/{
                  NamePos: 14960,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200cf600",&ast.Ident/*struct*/{
                  NamePos: 14966,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cdc20",&ast.BlockStmt/*struct*/{
        Lbrace: 14975,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200cf6e0",&ast.ReturnStmt/*struct*/{
            Return: 14977,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d0780",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200cf6c0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200cf680",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200cf640",&ast.Ident/*struct*/{
                      NamePos: 14984,
                      Name: "x",
                      Obj: p.PtrmapObject("0xc4200c8aa0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200cf660",&ast.Ident/*struct*/{
                      NamePos: 14986,
                      Name: "Value",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200cf6a0",&ast.Ident/*struct*/{
                    NamePos: 14992,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 14995,
                Ellipsis: 0,
                Rparen: 14996,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 14998,
      }/*struct*/),
    }/*struct*/),/* slice_item: 68*/69: p.StrmapFuncDecl("0xc4200cdd70",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cdcb0",&ast.FieldList/*struct*/{
        Opening: 15094,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0800",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cf760",&ast.StarExpr/*struct*/{
              Star: 15095,
              X: p.StrmapIdent("0xc4200cf740",&ast.Ident/*struct*/{
                NamePos: 15096,
                Name: "BadExpr",
                Obj: p.PtrmapObject("0xc420051b80"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15103,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cf780",&ast.Ident/*struct*/{
        NamePos: 15105,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cf7a0",&ast.FuncType/*struct*/{
        Func: 15089,
        Params: p.StrmapFieldList("0xc4200cdce0",&ast.FieldList/*struct*/{
          Opening: 15113,
          Closing: 15114,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cdd40",&ast.BlockStmt/*struct*/{
        Lbrace: 15123,
        Rbrace: 15124,
      }/*struct*/),
    }/*struct*/),/* slice_item: 69*/70: p.StrmapFuncDecl("0xc4200cde90",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cddd0",&ast.FieldList/*struct*/{
        Opening: 15131,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0840",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cf7e0",&ast.StarExpr/*struct*/{
              Star: 15132,
              X: p.StrmapIdent("0xc4200cf7c0",&ast.Ident/*struct*/{
                NamePos: 15133,
                Name: "Ident",
                Obj: p.PtrmapObject("0xc420051c70"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15138,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cf800",&ast.Ident/*struct*/{
        NamePos: 15140,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cf820",&ast.FuncType/*struct*/{
        Func: 15126,
        Params: p.StrmapFieldList("0xc4200cde00",&ast.FieldList/*struct*/{
          Opening: 15148,
          Closing: 15149,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cde60",&ast.BlockStmt/*struct*/{
        Lbrace: 15160,
        Rbrace: 15161,
      }/*struct*/),
    }/*struct*/),/* slice_item: 70*/71: p.StrmapFuncDecl("0xc4200cdfb0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200cdef0",&ast.FieldList/*struct*/{
        Opening: 15168,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0880",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cf860",&ast.StarExpr/*struct*/{
              Star: 15169,
              X: p.StrmapIdent("0xc4200cf840",&ast.Ident/*struct*/{
                NamePos: 15170,
                Name: "Ellipsis",
                Obj: p.PtrmapObject("0xc420051db0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15178,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cf880",&ast.Ident/*struct*/{
        NamePos: 15180,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cf8a0",&ast.FuncType/*struct*/{
        Func: 15163,
        Params: p.StrmapFieldList("0xc4200cdf20",&ast.FieldList/*struct*/{
          Opening: 15188,
          Closing: 15189,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200cdf80",&ast.BlockStmt/*struct*/{
        Lbrace: 15197,
        Rbrace: 15198,
      }/*struct*/),
    }/*struct*/),/* slice_item: 71*/72: p.StrmapFuncDecl("0xc4200d60f0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6030",&ast.FieldList/*struct*/{
        Opening: 15205,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d08c0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cf8e0",&ast.StarExpr/*struct*/{
              Star: 15206,
              X: p.StrmapIdent("0xc4200cf8c0",&ast.Ident/*struct*/{
                NamePos: 15207,
                Name: "BasicLit",
                Obj: p.PtrmapObject("0xc420051ea0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15215,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cf900",&ast.Ident/*struct*/{
        NamePos: 15217,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cf920",&ast.FuncType/*struct*/{
        Func: 15200,
        Params: p.StrmapFieldList("0xc4200d6060",&ast.FieldList/*struct*/{
          Opening: 15225,
          Closing: 15226,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d60c0",&ast.BlockStmt/*struct*/{
        Lbrace: 15234,
        Rbrace: 15235,
      }/*struct*/),
    }/*struct*/),/* slice_item: 72*/73: p.StrmapFuncDecl("0xc4200d6210",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6150",&ast.FieldList/*struct*/{
        Opening: 15242,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0900",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cf960",&ast.StarExpr/*struct*/{
              Star: 15243,
              X: p.StrmapIdent("0xc4200cf940",&ast.Ident/*struct*/{
                NamePos: 15244,
                Name: "FuncLit",
                Obj: p.PtrmapObject("0xc4200c0050"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15251,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cf980",&ast.Ident/*struct*/{
        NamePos: 15253,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cf9a0",&ast.FuncType/*struct*/{
        Func: 15237,
        Params: p.StrmapFieldList("0xc4200d6180",&ast.FieldList/*struct*/{
          Opening: 15261,
          Closing: 15262,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d61e0",&ast.BlockStmt/*struct*/{
        Lbrace: 15271,
        Rbrace: 15272,
      }/*struct*/),
    }/*struct*/),/* slice_item: 73*/74: p.StrmapFuncDecl("0xc4200d6330",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6270",&ast.FieldList/*struct*/{
        Opening: 15279,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0940",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cf9e0",&ast.StarExpr/*struct*/{
              Star: 15280,
              X: p.StrmapIdent("0xc4200cf9c0",&ast.Ident/*struct*/{
                NamePos: 15281,
                Name: "CompositeLit",
                Obj: p.PtrmapObject("0xc4200c0140"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15293,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfa00",&ast.Ident/*struct*/{
        NamePos: 15295,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfa20",&ast.FuncType/*struct*/{
        Func: 15274,
        Params: p.StrmapFieldList("0xc4200d62a0",&ast.FieldList/*struct*/{
          Opening: 15303,
          Closing: 15304,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6300",&ast.BlockStmt/*struct*/{
        Lbrace: 15308,
        Rbrace: 15309,
      }/*struct*/),
    }/*struct*/),/* slice_item: 74*/75: p.StrmapFuncDecl("0xc4200d6450",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6390",&ast.FieldList/*struct*/{
        Opening: 15316,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0980",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfa60",&ast.StarExpr/*struct*/{
              Star: 15317,
              X: p.StrmapIdent("0xc4200cfa40",&ast.Ident/*struct*/{
                NamePos: 15318,
                Name: "ParenExpr",
                Obj: p.PtrmapObject("0xc4200c02d0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15327,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfa80",&ast.Ident/*struct*/{
        NamePos: 15329,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfaa0",&ast.FuncType/*struct*/{
        Func: 15311,
        Params: p.StrmapFieldList("0xc4200d63c0",&ast.FieldList/*struct*/{
          Opening: 15337,
          Closing: 15338,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6420",&ast.BlockStmt/*struct*/{
        Lbrace: 15345,
        Rbrace: 15346,
      }/*struct*/),
    }/*struct*/),/* slice_item: 75*/76: p.StrmapFuncDecl("0xc4200d6570",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d64b0",&ast.FieldList/*struct*/{
        Opening: 15353,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d09c0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfae0",&ast.StarExpr/*struct*/{
              Star: 15354,
              X: p.StrmapIdent("0xc4200cfac0",&ast.Ident/*struct*/{
                NamePos: 15355,
                Name: "SelectorExpr",
                Obj: p.PtrmapObject("0xc4200c0460"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15367,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfb00",&ast.Ident/*struct*/{
        NamePos: 15369,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfb20",&ast.FuncType/*struct*/{
        Func: 15348,
        Params: p.StrmapFieldList("0xc4200d64e0",&ast.FieldList/*struct*/{
          Opening: 15377,
          Closing: 15378,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6540",&ast.BlockStmt/*struct*/{
        Lbrace: 15382,
        Rbrace: 15383,
      }/*struct*/),
    }/*struct*/),/* slice_item: 76*/77: p.StrmapFuncDecl("0xc4200d6690",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d65d0",&ast.FieldList/*struct*/{
        Opening: 15390,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0a00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfb60",&ast.StarExpr/*struct*/{
              Star: 15391,
              X: p.StrmapIdent("0xc4200cfb40",&ast.Ident/*struct*/{
                NamePos: 15392,
                Name: "IndexExpr",
                Obj: p.PtrmapObject("0xc4200c05a0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15401,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfb80",&ast.Ident/*struct*/{
        NamePos: 15403,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfba0",&ast.FuncType/*struct*/{
        Func: 15385,
        Params: p.StrmapFieldList("0xc4200d6600",&ast.FieldList/*struct*/{
          Opening: 15411,
          Closing: 15412,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6660",&ast.BlockStmt/*struct*/{
        Lbrace: 15419,
        Rbrace: 15420,
      }/*struct*/),
    }/*struct*/),/* slice_item: 77*/78: p.StrmapFuncDecl("0xc4200d67b0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d66f0",&ast.FieldList/*struct*/{
        Opening: 15427,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0a40",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfbe0",&ast.StarExpr/*struct*/{
              Star: 15428,
              X: p.StrmapIdent("0xc4200cfbc0",&ast.Ident/*struct*/{
                NamePos: 15429,
                Name: "SliceExpr",
                Obj: p.PtrmapObject("0xc4200c0780"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15438,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfc00",&ast.Ident/*struct*/{
        NamePos: 15440,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfc20",&ast.FuncType/*struct*/{
        Func: 15422,
        Params: p.StrmapFieldList("0xc4200d6720",&ast.FieldList/*struct*/{
          Opening: 15448,
          Closing: 15449,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6780",&ast.BlockStmt/*struct*/{
        Lbrace: 15456,
        Rbrace: 15457,
      }/*struct*/),
    }/*struct*/),/* slice_item: 78*/79: p.StrmapFuncDecl("0xc4200d68d0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6810",&ast.FieldList/*struct*/{
        Opening: 15464,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0a80",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfc60",&ast.StarExpr/*struct*/{
              Star: 15465,
              X: p.StrmapIdent("0xc4200cfc40",&ast.Ident/*struct*/{
                NamePos: 15466,
                Name: "TypeAssertExpr",
                Obj: p.PtrmapObject("0xc4200c0a00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15480,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfc80",&ast.Ident/*struct*/{
        NamePos: 15482,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfca0",&ast.FuncType/*struct*/{
        Func: 15459,
        Params: p.StrmapFieldList("0xc4200d6840",&ast.FieldList/*struct*/{
          Opening: 15490,
          Closing: 15491,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d68a0",&ast.BlockStmt/*struct*/{
        Lbrace: 15493,
        Rbrace: 15494,
      }/*struct*/),
    }/*struct*/),/* slice_item: 79*/80: p.StrmapFuncDecl("0xc4200d69f0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6930",&ast.FieldList/*struct*/{
        Opening: 15501,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0ac0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfce0",&ast.StarExpr/*struct*/{
              Star: 15502,
              X: p.StrmapIdent("0xc4200cfcc0",&ast.Ident/*struct*/{
                NamePos: 15503,
                Name: "CallExpr",
                Obj: p.PtrmapObject("0xc4200c0be0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15511,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfd00",&ast.Ident/*struct*/{
        NamePos: 15513,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfd20",&ast.FuncType/*struct*/{
        Func: 15496,
        Params: p.StrmapFieldList("0xc4200d6960",&ast.FieldList/*struct*/{
          Opening: 15521,
          Closing: 15522,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d69c0",&ast.BlockStmt/*struct*/{
        Lbrace: 15530,
        Rbrace: 15531,
      }/*struct*/),
    }/*struct*/),/* slice_item: 80*/81: p.StrmapFuncDecl("0xc4200d6b10",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6a50",&ast.FieldList/*struct*/{
        Opening: 15538,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0b00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfd60",&ast.StarExpr/*struct*/{
              Star: 15539,
              X: p.StrmapIdent("0xc4200cfd40",&ast.Ident/*struct*/{
                NamePos: 15540,
                Name: "StarExpr",
                Obj: p.PtrmapObject("0xc4200c0e60"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15548,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfd80",&ast.Ident/*struct*/{
        NamePos: 15550,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfda0",&ast.FuncType/*struct*/{
        Func: 15533,
        Params: p.StrmapFieldList("0xc4200d6a80",&ast.FieldList/*struct*/{
          Opening: 15558,
          Closing: 15559,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6ae0",&ast.BlockStmt/*struct*/{
        Lbrace: 15567,
        Rbrace: 15568,
      }/*struct*/),
    }/*struct*/),/* slice_item: 81*/82: p.StrmapFuncDecl("0xc4200d6c30",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6b70",&ast.FieldList/*struct*/{
        Opening: 15575,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0b40",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfde0",&ast.StarExpr/*struct*/{
              Star: 15576,
              X: p.StrmapIdent("0xc4200cfdc0",&ast.Ident/*struct*/{
                NamePos: 15577,
                Name: "UnaryExpr",
                Obj: p.PtrmapObject("0xc4200c0f50"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15586,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfe00",&ast.Ident/*struct*/{
        NamePos: 15588,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfe20",&ast.FuncType/*struct*/{
        Func: 15570,
        Params: p.StrmapFieldList("0xc4200d6ba0",&ast.FieldList/*struct*/{
          Opening: 15596,
          Closing: 15597,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6c00",&ast.BlockStmt/*struct*/{
        Lbrace: 15604,
        Rbrace: 15605,
      }/*struct*/),
    }/*struct*/),/* slice_item: 82*/83: p.StrmapFuncDecl("0xc4200d6d50",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6c90",&ast.FieldList/*struct*/{
        Opening: 15612,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0b80",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfe60",&ast.StarExpr/*struct*/{
              Star: 15613,
              X: p.StrmapIdent("0xc4200cfe40",&ast.Ident/*struct*/{
                NamePos: 15614,
                Name: "BinaryExpr",
                Obj: p.PtrmapObject("0xc4200c1090"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15624,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cfe80",&ast.Ident/*struct*/{
        NamePos: 15626,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cfea0",&ast.FuncType/*struct*/{
        Func: 15607,
        Params: p.StrmapFieldList("0xc4200d6cc0",&ast.FieldList/*struct*/{
          Opening: 15634,
          Closing: 15635,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6d20",&ast.BlockStmt/*struct*/{
        Lbrace: 15641,
        Rbrace: 15642,
      }/*struct*/),
    }/*struct*/),/* slice_item: 83*/84: p.StrmapFuncDecl("0xc4200d6e70",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6db0",&ast.FieldList/*struct*/{
        Opening: 15649,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0bc0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cfee0",&ast.StarExpr/*struct*/{
              Star: 15650,
              X: p.StrmapIdent("0xc4200cfec0",&ast.Ident/*struct*/{
                NamePos: 15651,
                Name: "KeyValueExpr",
                Obj: p.PtrmapObject("0xc4200c1220"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15663,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cff00",&ast.Ident/*struct*/{
        NamePos: 15665,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cff20",&ast.FuncType/*struct*/{
        Func: 15644,
        Params: p.StrmapFieldList("0xc4200d6de0",&ast.FieldList/*struct*/{
          Opening: 15673,
          Closing: 15674,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6e40",&ast.BlockStmt/*struct*/{
        Lbrace: 15678,
        Rbrace: 15679,
      }/*struct*/),
    }/*struct*/),/* slice_item: 84*/85: p.StrmapFuncDecl("0xc4200d6f90",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6ed0",&ast.FieldList/*struct*/{
        Opening: 15687,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0c00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cff60",&ast.StarExpr/*struct*/{
              Star: 15688,
              X: p.StrmapIdent("0xc4200cff40",&ast.Ident/*struct*/{
                NamePos: 15689,
                Name: "ArrayType",
                Obj: p.PtrmapObject("0xc4200c14f0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15698,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200cff80",&ast.Ident/*struct*/{
        NamePos: 15700,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200cffa0",&ast.FuncType/*struct*/{
        Func: 15682,
        Params: p.StrmapFieldList("0xc4200d6f00",&ast.FieldList/*struct*/{
          Opening: 15708,
          Closing: 15709,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d6f60",&ast.BlockStmt/*struct*/{
        Lbrace: 15715,
        Rbrace: 15716,
      }/*struct*/),
    }/*struct*/),/* slice_item: 85*/86: p.StrmapFuncDecl("0xc4200d70b0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d6ff0",&ast.FieldList/*struct*/{
        Opening: 15723,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0c40",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200cffe0",&ast.StarExpr/*struct*/{
              Star: 15724,
              X: p.StrmapIdent("0xc4200cffc0",&ast.Ident/*struct*/{
                NamePos: 15725,
                Name: "StructType",
                Obj: p.PtrmapObject("0xc4200c1630"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15735,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200d8000",&ast.Ident/*struct*/{
        NamePos: 15737,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200d8020",&ast.FuncType/*struct*/{
        Func: 15718,
        Params: p.StrmapFieldList("0xc4200d7020",&ast.FieldList/*struct*/{
          Opening: 15745,
          Closing: 15746,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d7080",&ast.BlockStmt/*struct*/{
        Lbrace: 15751,
        Rbrace: 15752,
      }/*struct*/),
    }/*struct*/),/* slice_item: 86*/87: p.StrmapFuncDecl("0xc4200d71d0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d7110",&ast.FieldList/*struct*/{
        Opening: 15759,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0c80",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200d8060",&ast.StarExpr/*struct*/{
              Star: 15760,
              X: p.StrmapIdent("0xc4200d8040",&ast.Ident/*struct*/{
                NamePos: 15761,
                Name: "FuncType",
                Obj: p.PtrmapObject("0xc4200c1770"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15769,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200d8080",&ast.Ident/*struct*/{
        NamePos: 15771,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200d80a0",&ast.FuncType/*struct*/{
        Func: 15754,
        Params: p.StrmapFieldList("0xc4200d7140",&ast.FieldList/*struct*/{
          Opening: 15779,
          Closing: 15780,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d71a0",&ast.BlockStmt/*struct*/{
        Lbrace: 15787,
        Rbrace: 15788,
      }/*struct*/),
    }/*struct*/),/* slice_item: 87*/88: p.StrmapFuncDecl("0xc4200d72f0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d7230",&ast.FieldList/*struct*/{
        Opening: 15795,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0cc0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200d80e0",&ast.StarExpr/*struct*/{
              Star: 15796,
              X: p.StrmapIdent("0xc4200d80c0",&ast.Ident/*struct*/{
                NamePos: 15797,
                Name: "InterfaceType",
                Obj: p.PtrmapObject("0xc4200c1900"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15810,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200d8100",&ast.Ident/*struct*/{
        NamePos: 15812,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200d8120",&ast.FuncType/*struct*/{
        Func: 15790,
        Params: p.StrmapFieldList("0xc4200d7260",&ast.FieldList/*struct*/{
          Opening: 15820,
          Closing: 15821,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d72c0",&ast.BlockStmt/*struct*/{
        Lbrace: 15823,
        Rbrace: 15824,
      }/*struct*/),
    }/*struct*/),/* slice_item: 88*/89: p.StrmapFuncDecl("0xc4200d7410",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d7350",&ast.FieldList/*struct*/{
        Opening: 15831,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0d00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200d8160",&ast.StarExpr/*struct*/{
              Star: 15832,
              X: p.StrmapIdent("0xc4200d8140",&ast.Ident/*struct*/{
                NamePos: 15833,
                Name: "MapType",
                Obj: p.PtrmapObject("0xc4200c1a40"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15840,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200d8180",&ast.Ident/*struct*/{
        NamePos: 15842,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200d81a0",&ast.FuncType/*struct*/{
        Func: 15826,
        Params: p.StrmapFieldList("0xc4200d7380",&ast.FieldList/*struct*/{
          Opening: 15850,
          Closing: 15851,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d73e0",&ast.BlockStmt/*struct*/{
        Lbrace: 15859,
        Rbrace: 15860,
      }/*struct*/),
    }/*struct*/),/* slice_item: 89*/90: p.StrmapFuncDecl("0xc4200d7560",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d7470",&ast.FieldList/*struct*/{
        Opening: 15867,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d0d40",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200d81e0",&ast.StarExpr/*struct*/{
              Star: 15868,
              X: p.StrmapIdent("0xc4200d81c0",&ast.Ident/*struct*/{
                NamePos: 15869,
                Name: "ChanType",
                Obj: p.PtrmapObject("0xc4200c1b80"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 15877,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200d8200",&ast.Ident/*struct*/{
        NamePos: 15879,
        Name: "exprNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200d8220",&ast.FuncType/*struct*/{
        Func: 15862,
        Params: p.StrmapFieldList("0xc4200d74a0",&ast.FieldList/*struct*/{
          Opening: 15887,
          Closing: 15888,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d7500",&ast.BlockStmt/*struct*/{
        Lbrace: 15895,
        Rbrace: 15896,
      }/*struct*/),
    }/*struct*/),/* slice_item: 90*/91: p.StrmapFuncDecl("0xc4200d7680",&ast.FuncDecl/*struct*/{
      Name: p.StrmapIdent("0xc4200d8240",&ast.Ident/*struct*/{
        NamePos: 16137,
        Name: "NewIdent",
        Obj: p.StrmapObject("0xc4200c8b90",&ast.Object/*struct*/{
          Kind: "func",
          Name: "NewIdent",
          Decl: p.PtrmapFuncDecl("0xc4200d7680"),
        }/*struct*/),
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200d8400",&ast.FuncType/*struct*/{
        Func: 16132,
        Params: p.StrmapFieldList("0xc4200d75c0",&ast.FieldList/*struct*/{
          Opening: 16145,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0e00",&ast.Field/*struct*/{
              Names: []*ast.Ident /*Slice*/{
                0: p.StrmapIdent("0xc4200d8260",&ast.Ident/*struct*/{
                  NamePos: 16146,
                  Name: "name",
                  Obj: p.StrmapObject("0xc4200c8b40",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "name",
                    Decl: p.PtrmapField("0xc4200d0e00"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Type: p.StrmapIdent("0xc4200d8280",&ast.Ident/*struct*/{
                NamePos: 16151,
                Name: "string",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 16157,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200d75f0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0e40",&ast.Field/*struct*/{
              Type: p.StrmapStarExpr("0xc4200d82c0",&ast.StarExpr/*struct*/{
                Star: 16159,
                X: p.StrmapIdent("0xc4200d82a0",&ast.Ident/*struct*/{
                  NamePos: 16160,
                  Name: "Ident",
                  Obj: p.PtrmapObject("0xc420051c70"),
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d7650",&ast.BlockStmt/*struct*/{
        Lbrace: 16166,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200d83e0",&ast.ReturnStmt/*struct*/{
            Return: 16168,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapUnaryExpr("0xc4200d83c0",&ast.UnaryExpr/*struct*/{
                OpPos: 16175,
                Op: token.Token(17)/*&*/,
                X: p.StrmapCompositeLit("0xc4200d0ec0",&ast.CompositeLit/*struct*/{
                  Type: p.StrmapIdent("0xc4200d82e0",&ast.Ident/*struct*/{
                    NamePos: 16176,
                    Name: "Ident",
                    Obj: p.PtrmapObject("0xc420051c70"),
                  }/*struct*/),
                  Lbrace: 16181,
                  Elts: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200d8340",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200d8300",&ast.Ident/*struct*/{
                        NamePos: 16182,
                        Name: "token",
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200d8320",&ast.Ident/*struct*/{
                        NamePos: 16188,
                        Name: "NoPos",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/1: p.StrmapIdent("0xc4200d8360",&ast.Ident/*struct*/{
                      NamePos: 16195,
                      Name: "name",
                      Obj: p.PtrmapObject("0xc4200c8b40"),
                    }/*struct*/),/* slice_item: 1*/2: p.StrmapIdent("0xc4200d83a0",&ast.Ident/*struct*/{
                      NamePos: 16201,
                      Name: "nil",
                    }/*struct*/),/* slice_item: 2*/}/*slice*/,
                  Rbrace: 16204,
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 16206,
      }/*struct*/),
    }/*struct*/),/* slice_item: 91*/92: p.StrmapFuncDecl("0xc4200d77a0",&ast.FuncDecl/*struct*/{
      Name: p.StrmapIdent("0xc4200d8420",&ast.Ident/*struct*/{
        NamePos: 16336,
        Name: "IsExported",
        Obj: p.StrmapObject("0xc4200c8cd0",&ast.Object/*struct*/{
          Kind: "func",
          Name: "IsExported",
          Decl: p.PtrmapFuncDecl("0xc4200d77a0"),
        }/*struct*/),
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200d8660",&ast.FuncType/*struct*/{
        Func: 16331,
        Params: p.StrmapFieldList("0xc4200d76e0",&ast.FieldList/*struct*/{
          Opening: 16346,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0f80",&ast.Field/*struct*/{
              Names: []*ast.Ident /*Slice*/{
                0: p.StrmapIdent("0xc4200d8440",&ast.Ident/*struct*/{
                  NamePos: 16347,
                  Name: "name",
                  Obj: p.StrmapObject("0xc4200c8be0",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "name",
                    Decl: p.PtrmapField("0xc4200d0f80"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Type: p.StrmapIdent("0xc4200d8460",&ast.Ident/*struct*/{
                NamePos: 16352,
                Name: "string",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 16358,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200d7710",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d0fc0",&ast.Field/*struct*/{
              Type: p.StrmapIdent("0xc4200d8480",&ast.Ident/*struct*/{
                NamePos: 16360,
                Name: "bool",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d7770",&ast.BlockStmt/*struct*/{
        Lbrace: 16365,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapAssignStmt("0xc4200d1040",&ast.AssignStmt/*struct*/{
            Lhs: []ast.Expr /*Slice*/{
              0: p.StrmapIdent("0xc4200d84a0",&ast.Ident/*struct*/{
                NamePos: 16368,
                Name: "ch",
                Obj: p.StrmapObject("0xc4200c8c30",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "ch",
                  Decl: p.PtrmapAssignStmt("0xc4200d1040"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/1: p.StrmapIdent("0xc4200d84c0",&ast.Ident/*struct*/{
                NamePos: 16372,
                Name: "_",
                Obj: p.StrmapObject("0xc4200c8c80",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "_",
                  Decl: p.PtrmapAssignStmt("0xc4200d1040"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 1*/}/*slice*/,
            TokPos: 16374,
            Tok: token.Token(47)/*:=*/,
            Rhs: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d1000",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200d8560",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200d8500",&ast.Ident/*struct*/{
                    NamePos: 16377,
                    Name: "utf8",
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200d8540",&ast.Ident/*struct*/{
                    NamePos: 16382,
                    Name: "DecodeRuneInString",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 16400,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapIdent("0xc4200d8580",&ast.Ident/*struct*/{
                    NamePos: 16401,
                    Name: "name",
                    Obj: p.PtrmapObject("0xc4200c8be0"),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 16405,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200d8620",&ast.ReturnStmt/*struct*/{
            Return: 16408,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d1080",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200d85e0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200d85a0",&ast.Ident/*struct*/{
                    NamePos: 16415,
                    Name: "unicode",
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200d85c0",&ast.Ident/*struct*/{
                    NamePos: 16423,
                    Name: "IsUpper",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 16430,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapIdent("0xc4200d8600",&ast.Ident/*struct*/{
                    NamePos: 16431,
                    Name: "ch",
                    Obj: p.PtrmapObject("0xc4200c8c30"),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 16433,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 16435,
      }/*struct*/),
    }/*struct*/),/* slice_item: 92*/93: p.StrmapFuncDecl("0xc4200d78f0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d7800",&ast.FieldList/*struct*/{
        Opening: 16562,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d1140",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200d8680",&ast.Ident/*struct*/{
                NamePos: 16563,
                Name: "id",
                Obj: p.StrmapObject("0xc4200c8d20",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "id",
                  Decl: p.PtrmapField("0xc4200d1140"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200d86c0",&ast.StarExpr/*struct*/{
              Star: 16566,
              X: p.StrmapIdent("0xc4200d86a0",&ast.Ident/*struct*/{
                NamePos: 16567,
                Name: "Ident",
                Obj: p.PtrmapObject("0xc420051c70"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 16572,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200d86e0",&ast.Ident/*struct*/{
        NamePos: 16574,
        Name: "IsExported",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200d87c0",&ast.FuncType/*struct*/{
        Func: 16557,
        Params: p.StrmapFieldList("0xc4200d7830",&ast.FieldList/*struct*/{
          Opening: 16584,
          Closing: 16585,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200d7860",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d1180",&ast.Field/*struct*/{
              Type: p.StrmapIdent("0xc4200d8700",&ast.Ident/*struct*/{
                NamePos: 16587,
                Name: "bool",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d78c0",&ast.BlockStmt/*struct*/{
        Lbrace: 16592,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200d87a0",&ast.ReturnStmt/*struct*/{
            Return: 16594,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200d11c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapIdent("0xc4200d8720",&ast.Ident/*struct*/{
                  NamePos: 16601,
                  Name: "IsExported",
                  Obj: p.PtrmapObject("0xc4200c8cd0"),
                }/*struct*/),
                Lparen: 16611,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapSelectorExpr("0xc4200d8780",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d8740",&ast.Ident/*struct*/{
                      NamePos: 16612,
                      Name: "id",
                      Obj: p.PtrmapObject("0xc4200c8d20"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d8760",&ast.Ident/*struct*/{
                      NamePos: 16615,
                      Name: "Name",
                    }/*struct*/),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 16619,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 16621,
      }/*struct*/),
    }/*struct*/),/* slice_item: 93*/94: p.StrmapFuncDecl("0xc4200d7b00",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200d7950",&ast.FieldList/*struct*/{
        Opening: 16629,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200d1200",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200d87e0",&ast.Ident/*struct*/{
                NamePos: 16630,
                Name: "id",
                Obj: p.StrmapObject("0xc4200c8d70",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "id",
                  Decl: p.PtrmapField("0xc4200d1200"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200d8820",&ast.StarExpr/*struct*/{
              Star: 16633,
              X: p.StrmapIdent("0xc4200d8800",&ast.Ident/*struct*/{
                NamePos: 16634,
                Name: "Ident",
                Obj: p.PtrmapObject("0xc420051c70"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 16639,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200d8840",&ast.Ident/*struct*/{
        NamePos: 16641,
        Name: "String",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200d89a0",&ast.FuncType/*struct*/{
        Func: 16624,
        Params: p.StrmapFieldList("0xc4200d7980",&ast.FieldList/*struct*/{
          Opening: 16647,
          Closing: 16648,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200d79b0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200d1240",&ast.Field/*struct*/{
              Type: p.StrmapIdent("0xc4200d8860",&ast.Ident/*struct*/{
                NamePos: 16650,
                Name: "string",
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200d7ad0",&ast.BlockStmt/*struct*/{
        Lbrace: 16657,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200d1280",&ast.IfStmt/*struct*/{
            If: 16660,
            Cond: p.StrmapBinaryExpr("0xc4200d7a40",&ast.BinaryExpr/*struct*/{
              X: p.StrmapIdent("0xc4200d8880",&ast.Ident/*struct*/{
                NamePos: 16663,
                Name: "id",
                Obj: p.PtrmapObject("0xc4200c8d70"),
              }/*struct*/),
              OpPos: 16666,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200d88a0",&ast.Ident/*struct*/{
                NamePos: 16669,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200d7aa0",&ast.BlockStmt/*struct*/{
              Lbrace: 16673,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200d8920",&ast.ReturnStmt/*struct*/{
                  Return: 16677,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200d8900",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200d88c0",&ast.Ident/*struct*/{
                        NamePos: 16684,
                        Name: "id",
                        Obj: p.PtrmapObject("0xc4200c8d70"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200d88e0",&ast.Ident/*struct*/{
                        NamePos: 16687,
                        Name: "Name",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 16693,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200d8960",&ast.ReturnStmt/*struct*/{
            Return: 16696,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBasicLit("0xc4200d8940",&ast.BasicLit/*struct*/{
                ValuePos: 16703,
                Kind: token.Token(9)/*STRING*/,
                Value: "\"<nil>\"",
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 16711,
      }/*struct*/),
    }/*struct*/),/* slice_item: 94*/95: p.StrmapGenDecl("0xc4200de980",&ast.GenDecl/*struct*/{
      TokPos: 16924,
      Tok: token.Token(84)/*type*/,
      Lparen: 16929,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200d7b30",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d89c0",&ast.Ident/*struct*/{
            NamePos: 17073,
            Name: "BadStmt",
            Obj: p.StrmapObject("0xc4200c8e10",&ast.Object/*struct*/{
              Kind: "type",
              Name: "BadStmt",
              Decl: p.PtrmapTypeSpec("0xc4200d7b30"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d8aa0",&ast.StructType/*struct*/{
            Struct: 17081,
            Fields: p.StrmapFieldList("0xc4200d7b90",&ast.FieldList/*struct*/{
              Opening: 17088,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d13c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d89e0",&ast.Ident/*struct*/{
                      NamePos: 17092,
                      Name: "From",
                      Obj: p.StrmapObject("0xc4200c8e60",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "From",
                        Decl: p.PtrmapField("0xc4200d13c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/1: p.StrmapIdent("0xc4200d8a00",&ast.Ident/*struct*/{
                      NamePos: 17098,
                      Name: "To",
                      Obj: p.StrmapObject("0xc4200c8eb0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "To",
                        Decl: p.PtrmapField("0xc4200d13c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 1*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d8a80",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d8a40",&ast.Ident/*struct*/{
                      NamePos: 17101,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d8a60",&ast.Ident/*struct*/{
                      NamePos: 17107,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Closing: 17147,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/1: p.StrmapTypeSpec("0xc4200d7bc0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d8ac0",&ast.Ident/*struct*/{
            NamePos: 17217,
            Name: "DeclStmt",
            Obj: p.StrmapObject("0xc4200c8f00",&ast.Object/*struct*/{
              Kind: "type",
              Name: "DeclStmt",
              Decl: p.PtrmapTypeSpec("0xc4200d7bc0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d8b20",&ast.StructType/*struct*/{
            Struct: 17226,
            Fields: p.StrmapFieldList("0xc4200d7c20",&ast.FieldList/*struct*/{
              Opening: 17233,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1440",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8ae0",&ast.Ident/*struct*/{
                      NamePos: 17237,
                      Name: "Decl",
                      Obj: p.StrmapObject("0xc4200c8f50",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Decl",
                        Decl: p.PtrmapField("0xc4200d1440"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d8b00",&ast.Ident/*struct*/{
                    NamePos: 17242,
                    Name: "Decl",
                    Obj: p.PtrmapObject("0xc420050780"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Closing: 17291,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 1*/2: p.StrmapTypeSpec("0xc4200d7c50",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d8b60",&ast.Ident/*struct*/{
            NamePos: 17477,
            Name: "EmptyStmt",
            Obj: p.StrmapObject("0xc4200c8ff0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "EmptyStmt",
              Decl: p.PtrmapTypeSpec("0xc4200d7c50"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d8c60",&ast.StructType/*struct*/{
            Struct: 17487,
            Fields: p.StrmapFieldList("0xc4200d7cb0",&ast.FieldList/*struct*/{
              Opening: 17494,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1500",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8b80",&ast.Ident/*struct*/{
                      NamePos: 17498,
                      Name: "Semicolon",
                      Obj: p.StrmapObject("0xc4200c9040",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Semicolon",
                        Decl: p.PtrmapField("0xc4200d1500"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d8be0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d8ba0",&ast.Ident/*struct*/{
                      NamePos: 17508,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d8bc0",&ast.Ident/*struct*/{
                      NamePos: 17514,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1540",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8c20",&ast.Ident/*struct*/{
                      NamePos: 17549,
                      Name: "Implicit",
                      Obj: p.StrmapObject("0xc4200c9090",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Implicit",
                        Decl: p.PtrmapField("0xc4200d1540"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d8c40",&ast.Ident/*struct*/{
                    NamePos: 17559,
                    Name: "bool",
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 17611,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 2*/3: p.StrmapTypeSpec("0xc4200d7ce0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d8c80",&ast.Ident/*struct*/{
            NamePos: 17670,
            Name: "LabeledStmt",
            Obj: p.StrmapObject("0xc4200c90e0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "LabeledStmt",
              Decl: p.PtrmapTypeSpec("0xc4200d7ce0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d8e00",&ast.StructType/*struct*/{
            Struct: 17682,
            Fields: p.StrmapFieldList("0xc4200d7d10",&ast.FieldList/*struct*/{
              Opening: 17689,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1600",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8ca0",&ast.Ident/*struct*/{
                      NamePos: 17693,
                      Name: "Label",
                      Obj: p.StrmapObject("0xc4200c9130",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Label",
                        Decl: p.PtrmapField("0xc4200d1600"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200d8ce0",&ast.StarExpr/*struct*/{
                    Star: 17699,
                    X: p.StrmapIdent("0xc4200d8cc0",&ast.Ident/*struct*/{
                      NamePos: 17700,
                      Name: "Ident",
                      Obj: p.PtrmapObject("0xc420051c70"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1640",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8d00",&ast.Ident/*struct*/{
                      NamePos: 17708,
                      Name: "Colon",
                      Obj: p.StrmapObject("0xc4200c9180",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Colon",
                        Decl: p.PtrmapField("0xc4200d1640"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d8d60",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d8d20",&ast.Ident/*struct*/{
                      NamePos: 17714,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d8d40",&ast.Ident/*struct*/{
                      NamePos: 17720,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200d1680",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8da0",&ast.Ident/*struct*/{
                      NamePos: 17745,
                      Name: "Stmt",
                      Obj: p.StrmapObject("0xc4200c91d0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Stmt",
                        Decl: p.PtrmapField("0xc4200d1680"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d8dc0",&ast.Ident/*struct*/{
                    NamePos: 17751,
                    Name: "Stmt",
                    Obj: p.PtrmapObject("0xc4200506e0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 17757,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 3*/4: p.StrmapTypeSpec("0xc4200d7d40",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d8e40",&ast.Ident/*struct*/{
            NamePos: 17849,
            Name: "ExprStmt",
            Obj: p.StrmapObject("0xc4200c9220",&ast.Object/*struct*/{
              Kind: "type",
              Name: "ExprStmt",
              Decl: p.PtrmapTypeSpec("0xc4200d7d40"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d8ea0",&ast.StructType/*struct*/{
            Struct: 17858,
            Fields: p.StrmapFieldList("0xc4200d7d70",&ast.FieldList/*struct*/{
              Opening: 17865,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1700",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8e60",&ast.Ident/*struct*/{
                      NamePos: 17869,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c9270",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200d1700"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d8e80",&ast.Ident/*struct*/{
                    NamePos: 17871,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Closing: 17891,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 4*/5: p.StrmapTypeSpec("0xc4200d7dd0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d8ec0",&ast.Ident/*struct*/{
            NamePos: 17944,
            Name: "SendStmt",
            Obj: p.StrmapObject("0xc4200c92c0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "SendStmt",
              Decl: p.PtrmapTypeSpec("0xc4200d7dd0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d9020",&ast.StructType/*struct*/{
            Struct: 17953,
            Fields: p.StrmapFieldList("0xc4200d7e00",&ast.FieldList/*struct*/{
              Opening: 17960,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1740",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8ee0",&ast.Ident/*struct*/{
                      NamePos: 17964,
                      Name: "Chan",
                      Obj: p.StrmapObject("0xc4200c9310",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Chan",
                        Decl: p.PtrmapField("0xc4200d1740"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d8f00",&ast.Ident/*struct*/{
                    NamePos: 17970,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1780",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8f20",&ast.Ident/*struct*/{
                      NamePos: 17977,
                      Name: "Arrow",
                      Obj: p.StrmapObject("0xc4200c9360",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Arrow",
                        Decl: p.PtrmapField("0xc4200d1780"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d8f80",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d8f40",&ast.Ident/*struct*/{
                      NamePos: 17983,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d8f60",&ast.Ident/*struct*/{
                      NamePos: 17989,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200d17c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d8fc0",&ast.Ident/*struct*/{
                      NamePos: 18015,
                      Name: "Value",
                      Obj: p.StrmapObject("0xc4200c93b0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Value",
                        Decl: p.PtrmapField("0xc4200d17c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d8fe0",&ast.Ident/*struct*/{
                    NamePos: 18021,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 18027,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 5*/6: p.StrmapTypeSpec("0xc4200d7e30",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d9040",&ast.Ident/*struct*/{
            NamePos: 18102,
            Name: "IncDecStmt",
            Obj: p.StrmapObject("0xc4200c9450",&ast.Object/*struct*/{
              Kind: "type",
              Name: "IncDecStmt",
              Decl: p.PtrmapTypeSpec("0xc4200d7e30"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d91e0",&ast.StructType/*struct*/{
            Struct: 18113,
            Fields: p.StrmapFieldList("0xc4200d7e60",&ast.FieldList/*struct*/{
              Opening: 18120,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1800",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9060",&ast.Ident/*struct*/{
                      NamePos: 18124,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200c94a0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200d1800"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d9080",&ast.Ident/*struct*/{
                    NamePos: 18131,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1840",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d90a0",&ast.Ident/*struct*/{
                      NamePos: 18138,
                      Name: "TokPos",
                      Obj: p.StrmapObject("0xc4200c94f0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "TokPos",
                        Decl: p.PtrmapField("0xc4200d1840"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d9100",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d90c0",&ast.Ident/*struct*/{
                      NamePos: 18145,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d90e0",&ast.Ident/*struct*/{
                      NamePos: 18151,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200d1880",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9140",&ast.Ident/*struct*/{
                      NamePos: 18178,
                      Name: "Tok",
                      Obj: p.StrmapObject("0xc4200c9540",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Tok",
                        Decl: p.PtrmapField("0xc4200d1880"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d91a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d9160",&ast.Ident/*struct*/{
                      NamePos: 18185,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d9180",&ast.Ident/*struct*/{
                      NamePos: 18191,
                      Name: "Token",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 18212,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 6*/7: p.StrmapTypeSpec("0xc4200d7e90",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d9220",&ast.Ident/*struct*/{
            NamePos: 18305,
            Name: "AssignStmt",
            Obj: p.StrmapObject("0xc4200c9590",&ast.Object/*struct*/{
              Kind: "type",
              Name: "AssignStmt",
              Decl: p.PtrmapTypeSpec("0xc4200d7e90"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d9420",&ast.StructType/*struct*/{
            Struct: 18316,
            Fields: p.StrmapFieldList("0xc4200d7f20",&ast.FieldList/*struct*/{
              Opening: 18323,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1900",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9240",&ast.Ident/*struct*/{
                      NamePos: 18327,
                      Name: "Lhs",
                      Obj: p.StrmapObject("0xc4200c95e0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Lhs",
                        Decl: p.PtrmapField("0xc4200d1900"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200d7ec0",&ast.ArrayType/*struct*/{
                    Lbrack: 18334,
                    Elt: p.StrmapIdent("0xc4200d9260",&ast.Ident/*struct*/{
                      NamePos: 18336,
                      Name: "Expr",
                      Obj: p.PtrmapObject("0xc420050640"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1940",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9280",&ast.Ident/*struct*/{
                      NamePos: 18343,
                      Name: "TokPos",
                      Obj: p.StrmapObject("0xc4200c9630",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "TokPos",
                        Decl: p.PtrmapField("0xc4200d1940"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d92e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d92a0",&ast.Ident/*struct*/{
                      NamePos: 18350,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d92c0",&ast.Ident/*struct*/{
                      NamePos: 18356,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200d1980",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9320",&ast.Ident/*struct*/{
                      NamePos: 18383,
                      Name: "Tok",
                      Obj: p.StrmapObject("0xc4200c9680",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Tok",
                        Decl: p.PtrmapField("0xc4200d1980"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d9380",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d9340",&ast.Ident/*struct*/{
                      NamePos: 18390,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d9360",&ast.Ident/*struct*/{
                      NamePos: 18396,
                      Name: "Token",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200d19c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d93e0",&ast.Ident/*struct*/{
                      NamePos: 18432,
                      Name: "Rhs",
                      Obj: p.StrmapObject("0xc4200c96d0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Rhs",
                        Decl: p.PtrmapField("0xc4200d19c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200d7ef0",&ast.ArrayType/*struct*/{
                    Lbrack: 18439,
                    Elt: p.StrmapIdent("0xc4200d9400",&ast.Ident/*struct*/{
                      NamePos: 18441,
                      Name: "Expr",
                      Obj: p.PtrmapObject("0xc420050640"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 18447,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 7*/8: p.StrmapTypeSpec("0xc4200d7f80",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d9440",&ast.Ident/*struct*/{
            NamePos: 18496,
            Name: "GoStmt",
            Obj: p.StrmapObject("0xc4200c9720",&ast.Object/*struct*/{
              Kind: "type",
              Name: "GoStmt",
              Decl: p.PtrmapTypeSpec("0xc4200d7f80"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d9560",&ast.StructType/*struct*/{
            Struct: 18503,
            Fields: p.StrmapFieldList("0xc4200d7fb0",&ast.FieldList/*struct*/{
              Opening: 18510,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1a00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9460",&ast.Ident/*struct*/{
                      NamePos: 18514,
                      Name: "Go",
                      Obj: p.StrmapObject("0xc4200c9770",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Go",
                        Decl: p.PtrmapField("0xc4200d1a00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d94c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d9480",&ast.Ident/*struct*/{
                      NamePos: 18519,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d94a0",&ast.Ident/*struct*/{
                      NamePos: 18525,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1a40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9500",&ast.Ident/*struct*/{
                      NamePos: 18559,
                      Name: "Call",
                      Obj: p.StrmapObject("0xc4200c97c0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Call",
                        Decl: p.PtrmapField("0xc4200d1a40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200d9540",&ast.StarExpr/*struct*/{
                    Star: 18564,
                    X: p.StrmapIdent("0xc4200d9520",&ast.Ident/*struct*/{
                      NamePos: 18565,
                      Name: "CallExpr",
                      Obj: p.PtrmapObject("0xc4200c0be0"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 18575,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 8*/9: p.StrmapTypeSpec("0xc4200dc000",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d9580",&ast.Ident/*struct*/{
            NamePos: 18630,
            Name: "DeferStmt",
            Obj: p.StrmapObject("0xc4200c9810",&ast.Object/*struct*/{
              Kind: "type",
              Name: "DeferStmt",
              Decl: p.PtrmapTypeSpec("0xc4200dc000"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d96a0",&ast.StructType/*struct*/{
            Struct: 18640,
            Fields: p.StrmapFieldList("0xc4200dc030",&ast.FieldList/*struct*/{
              Opening: 18647,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1ac0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d95a0",&ast.Ident/*struct*/{
                      NamePos: 18651,
                      Name: "Defer",
                      Obj: p.StrmapObject("0xc4200c9860",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Defer",
                        Decl: p.PtrmapField("0xc4200d1ac0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d9600",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d95c0",&ast.Ident/*struct*/{
                      NamePos: 18657,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d95e0",&ast.Ident/*struct*/{
                      NamePos: 18663,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1b00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9640",&ast.Ident/*struct*/{
                      NamePos: 18700,
                      Name: "Call",
                      Obj: p.StrmapObject("0xc4200c98b0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Call",
                        Decl: p.PtrmapField("0xc4200d1b00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200d9680",&ast.StarExpr/*struct*/{
                    Star: 18706,
                    X: p.StrmapIdent("0xc4200d9660",&ast.Ident/*struct*/{
                      NamePos: 18707,
                      Name: "CallExpr",
                      Obj: p.PtrmapObject("0xc4200c0be0"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 18717,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 9*/10: p.StrmapTypeSpec("0xc4200dc060",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d96c0",&ast.Ident/*struct*/{
            NamePos: 18774,
            Name: "ReturnStmt",
            Obj: p.StrmapObject("0xc4200c9900",&ast.Object/*struct*/{
              Kind: "type",
              Name: "ReturnStmt",
              Decl: p.PtrmapTypeSpec("0xc4200dc060"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d97e0",&ast.StructType/*struct*/{
            Struct: 18785,
            Fields: p.StrmapFieldList("0xc4200dc0c0",&ast.FieldList/*struct*/{
              Opening: 18792,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1b80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d96e0",&ast.Ident/*struct*/{
                      NamePos: 18796,
                      Name: "Return",
                      Obj: p.StrmapObject("0xc4200c9950",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Return",
                        Decl: p.PtrmapField("0xc4200d1b80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d9740",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d9700",&ast.Ident/*struct*/{
                      NamePos: 18804,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d9720",&ast.Ident/*struct*/{
                      NamePos: 18810,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1bc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9780",&ast.Ident/*struct*/{
                      NamePos: 18848,
                      Name: "Results",
                      Obj: p.StrmapObject("0xc4200c99a0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Results",
                        Decl: p.PtrmapField("0xc4200d1bc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200dc090",&ast.ArrayType/*struct*/{
                    Lbrack: 18856,
                    Elt: p.StrmapIdent("0xc4200d97a0",&ast.Ident/*struct*/{
                      NamePos: 18858,
                      Name: "Expr",
                      Obj: p.PtrmapObject("0xc420050640"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 18897,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 10*/11: p.StrmapTypeSpec("0xc4200dc0f0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d9820",&ast.Ident/*struct*/{
            NamePos: 18993,
            Name: "BranchStmt",
            Obj: p.StrmapObject("0xc4200c99f0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "BranchStmt",
              Decl: p.PtrmapTypeSpec("0xc4200dc0f0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d9a00",&ast.StructType/*struct*/{
            Struct: 19004,
            Fields: p.StrmapFieldList("0xc4200dc120",&ast.FieldList/*struct*/{
              Opening: 19011,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1c40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9840",&ast.Ident/*struct*/{
                      NamePos: 19015,
                      Name: "TokPos",
                      Obj: p.StrmapObject("0xc4200c9a40",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "TokPos",
                        Decl: p.PtrmapField("0xc4200d1c40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d98a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d9860",&ast.Ident/*struct*/{
                      NamePos: 19022,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d9880",&ast.Ident/*struct*/{
                      NamePos: 19028,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1cc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d98e0",&ast.Ident/*struct*/{
                      NamePos: 19055,
                      Name: "Tok",
                      Obj: p.StrmapObject("0xc4200c9a90",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Tok",
                        Decl: p.PtrmapField("0xc4200d1cc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d9940",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d9900",&ast.Ident/*struct*/{
                      NamePos: 19062,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d9920",&ast.Ident/*struct*/{
                      NamePos: 19068,
                      Name: "Token",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200d1d00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9960",&ast.Ident/*struct*/{
                      NamePos: 19130,
                      Name: "Label",
                      Obj: p.StrmapObject("0xc4200c9ae0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Label",
                        Decl: p.PtrmapField("0xc4200d1d00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200d99a0",&ast.StarExpr/*struct*/{
                    Star: 19137,
                    X: p.StrmapIdent("0xc4200d9980",&ast.Ident/*struct*/{
                      NamePos: 19138,
                      Name: "Ident",
                      Obj: p.PtrmapObject("0xc420051c70"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/}/*slice*/,
              Closing: 19172,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 11*/12: p.PtrmapTypeSpec("0xc4200dc150"),/* slice_item: 12*/13: p.StrmapTypeSpec("0xc4200dc210",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d9c00",&ast.Ident/*struct*/{
            NamePos: 19396,
            Name: "IfStmt",
            Obj: p.StrmapObject("0xc4200c9c70",&ast.Object/*struct*/{
              Kind: "type",
              Name: "IfStmt",
              Decl: p.PtrmapTypeSpec("0xc4200dc210"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200d9e20",&ast.StructType/*struct*/{
            Struct: 19403,
            Fields: p.StrmapFieldList("0xc4200dc270",&ast.FieldList/*struct*/{
              Opening: 19410,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1e40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9c20",&ast.Ident/*struct*/{
                      NamePos: 19414,
                      Name: "If",
                      Obj: p.StrmapObject("0xc4200c9cc0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "If",
                        Decl: p.PtrmapField("0xc4200d1e40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d9c80",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d9c40",&ast.Ident/*struct*/{
                      NamePos: 19419,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d9c60",&ast.Ident/*struct*/{
                      NamePos: 19425,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200d1e80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9cc0",&ast.Ident/*struct*/{
                      NamePos: 19459,
                      Name: "Init",
                      Obj: p.StrmapObject("0xc4200c9d10",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Init",
                        Decl: p.PtrmapField("0xc4200d1e80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d9ce0",&ast.Ident/*struct*/{
                    NamePos: 19464,
                    Name: "Stmt",
                    Obj: p.PtrmapObject("0xc4200506e0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200d1ec0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9d00",&ast.Ident/*struct*/{
                      NamePos: 19512,
                      Name: "Cond",
                      Obj: p.StrmapObject("0xc4200c9d60",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Cond",
                        Decl: p.PtrmapField("0xc4200d1ec0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d9d20",&ast.Ident/*struct*/{
                    NamePos: 19517,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200d1f00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9d60",&ast.Ident/*struct*/{
                      NamePos: 19542,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200c9db0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200d1f00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200d9da0",&ast.StarExpr/*struct*/{
                    Star: 19547,
                    X: p.StrmapIdent("0xc4200d9d80",&ast.Ident/*struct*/{
                      NamePos: 19548,
                      Name: "BlockStmt",
                      Obj: p.PtrmapObject("0xc4200c9b30"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200d1f40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9dc0",&ast.Ident/*struct*/{
                      NamePos: 19560,
                      Name: "Else",
                      Obj: p.StrmapObject("0xc4200c9e00",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Else",
                        Decl: p.PtrmapField("0xc4200d1f40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200d9de0",&ast.Ident/*struct*/{
                    NamePos: 19565,
                    Name: "Stmt",
                    Obj: p.PtrmapObject("0xc4200506e0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/}/*slice*/,
              Closing: 19594,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 13*/14: p.StrmapTypeSpec("0xc4200dc2a0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200d9e40",&ast.Ident/*struct*/{
            NamePos: 19676,
            Name: "CaseClause",
            Obj: p.StrmapObject("0xc4200c9ea0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "CaseClause",
              Decl: p.PtrmapTypeSpec("0xc4200dc2a0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200e0040",&ast.StructType/*struct*/{
            Struct: 19687,
            Fields: p.StrmapFieldList("0xc4200dc360",&ast.FieldList/*struct*/{
              Opening: 19694,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200d1fc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9e60",&ast.Ident/*struct*/{
                      NamePos: 19698,
                      Name: "Case",
                      Obj: p.StrmapObject("0xc4200c9ef0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Case",
                        Decl: p.PtrmapField("0xc4200d1fc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d9ec0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d9e80",&ast.Ident/*struct*/{
                      NamePos: 19704,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d9ea0",&ast.Ident/*struct*/{
                      NamePos: 19710,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200de040",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9ee0",&ast.Ident/*struct*/{
                      NamePos: 19759,
                      Name: "List",
                      Obj: p.StrmapObject("0xc4200c9f40",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "List",
                        Decl: p.PtrmapField("0xc4200de040"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200dc300",&ast.ArrayType/*struct*/{
                    Lbrack: 19765,
                    Elt: p.StrmapIdent("0xc4200d9f00",&ast.Ident/*struct*/{
                      NamePos: 19767,
                      Name: "Expr",
                      Obj: p.PtrmapObject("0xc420050640"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200de080",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9f20",&ast.Ident/*struct*/{
                      NamePos: 19833,
                      Name: "Colon",
                      Obj: p.StrmapObject("0xc4200c9f90",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Colon",
                        Decl: p.PtrmapField("0xc4200de080"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200d9f80",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200d9f40",&ast.Ident/*struct*/{
                      NamePos: 19839,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200d9f60",&ast.Ident/*struct*/{
                      NamePos: 19845,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200de0c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200d9fe0",&ast.Ident/*struct*/{
                      NamePos: 19870,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200e2000",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200de0c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200dc330",&ast.ArrayType/*struct*/{
                    Lbrack: 19876,
                    Elt: p.StrmapIdent("0xc4200e0000",&ast.Ident/*struct*/{
                      NamePos: 19878,
                      Name: "Stmt",
                      Obj: p.PtrmapObject("0xc4200506e0"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 19913,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 14*/15: p.StrmapTypeSpec("0xc4200dc390",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200e0060",&ast.Ident/*struct*/{
            NamePos: 19982,
            Name: "SwitchStmt",
            Obj: p.StrmapObject("0xc4200e2050",&ast.Object/*struct*/{
              Kind: "type",
              Name: "SwitchStmt",
              Decl: p.PtrmapTypeSpec("0xc4200dc390"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200e0260",&ast.StructType/*struct*/{
            Struct: 19993,
            Fields: p.StrmapFieldList("0xc4200dc3f0",&ast.FieldList/*struct*/{
              Opening: 20000,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200de140",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0080",&ast.Ident/*struct*/{
                      NamePos: 20004,
                      Name: "Switch",
                      Obj: p.StrmapObject("0xc4200e20a0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Switch",
                        Decl: p.PtrmapField("0xc4200de140"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200e00e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e00a0",&ast.Ident/*struct*/{
                      NamePos: 20011,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e00c0",&ast.Ident/*struct*/{
                      NamePos: 20017,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200de180",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0120",&ast.Ident/*struct*/{
                      NamePos: 20056,
                      Name: "Init",
                      Obj: p.StrmapObject("0xc4200e20f0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Init",
                        Decl: p.PtrmapField("0xc4200de180"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e0140",&ast.Ident/*struct*/{
                    NamePos: 20063,
                    Name: "Stmt",
                    Obj: p.PtrmapObject("0xc4200506e0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200de1c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0160",&ast.Ident/*struct*/{
                      NamePos: 20112,
                      Name: "Tag",
                      Obj: p.StrmapObject("0xc4200e2140",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Tag",
                        Decl: p.PtrmapField("0xc4200de1c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e0180",&ast.Ident/*struct*/{
                    NamePos: 20119,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200de200",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e01e0",&ast.Ident/*struct*/{
                      NamePos: 20158,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200e2190",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200de200"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200e0220",&ast.StarExpr/*struct*/{
                    Star: 20165,
                    X: p.StrmapIdent("0xc4200e0200",&ast.Ident/*struct*/{
                      NamePos: 20166,
                      Name: "BlockStmt",
                      Obj: p.PtrmapObject("0xc4200c9b30"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 20197,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 15*/16: p.StrmapTypeSpec("0xc4200dc420",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200e0280",&ast.Ident/*struct*/{
            NamePos: 20264,
            Name: "TypeSwitchStmt",
            Obj: p.StrmapObject("0xc4200e21e0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "TypeSwitchStmt",
              Decl: p.PtrmapTypeSpec("0xc4200dc420"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200e0480",&ast.StructType/*struct*/{
            Struct: 20279,
            Fields: p.StrmapFieldList("0xc4200dc480",&ast.FieldList/*struct*/{
              Opening: 20286,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200de280",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e02a0",&ast.Ident/*struct*/{
                      NamePos: 20290,
                      Name: "Switch",
                      Obj: p.StrmapObject("0xc4200e2230",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Switch",
                        Decl: p.PtrmapField("0xc4200de280"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200e0300",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e02c0",&ast.Ident/*struct*/{
                      NamePos: 20297,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e02e0",&ast.Ident/*struct*/{
                      NamePos: 20303,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200de2c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0340",&ast.Ident/*struct*/{
                      NamePos: 20342,
                      Name: "Init",
                      Obj: p.StrmapObject("0xc4200e2280",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Init",
                        Decl: p.PtrmapField("0xc4200de2c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e0360",&ast.Ident/*struct*/{
                    NamePos: 20349,
                    Name: "Stmt",
                    Obj: p.PtrmapObject("0xc4200506e0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200de300",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0380",&ast.Ident/*struct*/{
                      NamePos: 20398,
                      Name: "Assign",
                      Obj: p.StrmapObject("0xc4200e22d0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Assign",
                        Decl: p.PtrmapField("0xc4200de300"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e03a0",&ast.Ident/*struct*/{
                    NamePos: 20405,
                    Name: "Stmt",
                    Obj: p.PtrmapObject("0xc4200506e0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200de340",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0400",&ast.Ident/*struct*/{
                      NamePos: 20447,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200e2320",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200de340"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200e0440",&ast.StarExpr/*struct*/{
                    Star: 20454,
                    X: p.StrmapIdent("0xc4200e0420",&ast.Ident/*struct*/{
                      NamePos: 20455,
                      Name: "BlockStmt",
                      Obj: p.PtrmapObject("0xc4200c9b30"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 20486,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 16*/17: p.StrmapTypeSpec("0xc4200dc4b0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200e04a0",&ast.Ident/*struct*/{
            NamePos: 20553,
            Name: "CommClause",
            Obj: p.StrmapObject("0xc4200e2370",&ast.Object/*struct*/{
              Kind: "type",
              Name: "CommClause",
              Decl: p.PtrmapTypeSpec("0xc4200dc4b0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200e06a0",&ast.StructType/*struct*/{
            Struct: 20564,
            Fields: p.StrmapFieldList("0xc4200dc540",&ast.FieldList/*struct*/{
              Opening: 20571,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200de3c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e04c0",&ast.Ident/*struct*/{
                      NamePos: 20575,
                      Name: "Case",
                      Obj: p.StrmapObject("0xc4200e23c0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Case",
                        Decl: p.PtrmapField("0xc4200de3c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200e0520",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e04e0",&ast.Ident/*struct*/{
                      NamePos: 20581,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e0500",&ast.Ident/*struct*/{
                      NamePos: 20587,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200de440",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0540",&ast.Ident/*struct*/{
                      NamePos: 20636,
                      Name: "Comm",
                      Obj: p.StrmapObject("0xc4200e2410",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Comm",
                        Decl: p.PtrmapField("0xc4200de440"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e0560",&ast.Ident/*struct*/{
                    NamePos: 20642,
                    Name: "Stmt",
                    Obj: p.PtrmapObject("0xc4200506e0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200de480",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0580",&ast.Ident/*struct*/{
                      NamePos: 20707,
                      Name: "Colon",
                      Obj: p.StrmapObject("0xc4200e2460",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Colon",
                        Decl: p.PtrmapField("0xc4200de480"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200e05e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e05a0",&ast.Ident/*struct*/{
                      NamePos: 20713,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e05c0",&ast.Ident/*struct*/{
                      NamePos: 20719,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200de4c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0640",&ast.Ident/*struct*/{
                      NamePos: 20744,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200e24b0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200de4c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200dc510",&ast.ArrayType/*struct*/{
                    Lbrack: 20750,
                    Elt: p.StrmapIdent("0xc4200e0660",&ast.Ident/*struct*/{
                      NamePos: 20752,
                      Name: "Stmt",
                      Obj: p.PtrmapObject("0xc4200506e0"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 20787,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 17*/18: p.StrmapTypeSpec("0xc4200dc570",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200e06c0",&ast.Ident/*struct*/{
            NamePos: 20845,
            Name: "SelectStmt",
            Obj: p.StrmapObject("0xc4200e2500",&ast.Object/*struct*/{
              Kind: "type",
              Name: "SelectStmt",
              Decl: p.PtrmapTypeSpec("0xc4200dc570"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200e0800",&ast.StructType/*struct*/{
            Struct: 20856,
            Fields: p.StrmapFieldList("0xc4200dc5a0",&ast.FieldList/*struct*/{
              Opening: 20863,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200de540",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e06e0",&ast.Ident/*struct*/{
                      NamePos: 20867,
                      Name: "Select",
                      Obj: p.StrmapObject("0xc4200e2550",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Select",
                        Decl: p.PtrmapField("0xc4200de540"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200e0740",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e0700",&ast.Ident/*struct*/{
                      NamePos: 20874,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e0720",&ast.Ident/*struct*/{
                      NamePos: 20880,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200de580",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0780",&ast.Ident/*struct*/{
                      NamePos: 20919,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200e25a0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200de580"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200e07c0",&ast.StarExpr/*struct*/{
                    Star: 20926,
                    X: p.StrmapIdent("0xc4200e07a0",&ast.Ident/*struct*/{
                      NamePos: 20927,
                      Name: "BlockStmt",
                      Obj: p.PtrmapObject("0xc4200c9b30"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 20958,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 18*/19: p.StrmapTypeSpec("0xc4200dc600",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200e0820",&ast.Ident/*struct*/{
            NamePos: 21004,
            Name: "ForStmt",
            Obj: p.StrmapObject("0xc4200e25f0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "ForStmt",
              Decl: p.PtrmapTypeSpec("0xc4200dc600"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200e0a40",&ast.StructType/*struct*/{
            Struct: 21012,
            Fields: p.StrmapFieldList("0xc4200dc690",&ast.FieldList/*struct*/{
              Opening: 21019,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200de5c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0840",&ast.Ident/*struct*/{
                      NamePos: 21023,
                      Name: "For",
                      Obj: p.StrmapObject("0xc4200e2640",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "For",
                        Decl: p.PtrmapField("0xc4200de5c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200e08a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e0860",&ast.Ident/*struct*/{
                      NamePos: 21028,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e0880",&ast.Ident/*struct*/{
                      NamePos: 21034,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200de600",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e08e0",&ast.Ident/*struct*/{
                      NamePos: 21069,
                      Name: "Init",
                      Obj: p.StrmapObject("0xc4200e2690",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Init",
                        Decl: p.PtrmapField("0xc4200de600"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e0900",&ast.Ident/*struct*/{
                    NamePos: 21074,
                    Name: "Stmt",
                    Obj: p.PtrmapObject("0xc4200506e0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200de640",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0920",&ast.Ident/*struct*/{
                      NamePos: 21122,
                      Name: "Cond",
                      Obj: p.StrmapObject("0xc4200e26e0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Cond",
                        Decl: p.PtrmapField("0xc4200de640"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e0940",&ast.Ident/*struct*/{
                    NamePos: 21127,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200de680",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e09a0",&ast.Ident/*struct*/{
                      NamePos: 21160,
                      Name: "Post",
                      Obj: p.StrmapObject("0xc4200e2730",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Post",
                        Decl: p.PtrmapField("0xc4200de680"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e09c0",&ast.Ident/*struct*/{
                    NamePos: 21165,
                    Name: "Stmt",
                    Obj: p.PtrmapObject("0xc4200506e0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200de6c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e09e0",&ast.Ident/*struct*/{
                      NamePos: 21213,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200e2780",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200de6c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200e0a20",&ast.StarExpr/*struct*/{
                    Star: 21218,
                    X: p.StrmapIdent("0xc4200e0a00",&ast.Ident/*struct*/{
                      NamePos: 21219,
                      Name: "BlockStmt",
                      Obj: p.PtrmapObject("0xc4200c9b30"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/}/*slice*/,
              Closing: 21230,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 19*/20: p.StrmapTypeSpec("0xc4200dc6c0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200e0a60",&ast.Ident/*struct*/{
            NamePos: 21298,
            Name: "RangeStmt",
            Obj: p.StrmapObject("0xc4200e27d0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "RangeStmt",
              Decl: p.PtrmapTypeSpec("0xc4200dc6c0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200e0da0",&ast.StructType/*struct*/{
            Struct: 21308,
            Fields: p.StrmapFieldList("0xc4200dc750",&ast.FieldList/*struct*/{
              Opening: 21315,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200de780",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0a80",&ast.Ident/*struct*/{
                      NamePos: 21319,
                      Name: "For",
                      Obj: p.StrmapObject("0xc4200e2820",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "For",
                        Decl: p.PtrmapField("0xc4200de780"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200e0ae0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e0aa0",&ast.Ident/*struct*/{
                      NamePos: 21330,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e0ac0",&ast.Ident/*struct*/{
                      NamePos: 21336,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200de7c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0b20",&ast.Ident/*struct*/{
                      NamePos: 21373,
                      Name: "Key",
                      Obj: p.StrmapObject("0xc4200e2870",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Key",
                        Decl: p.PtrmapField("0xc4200de7c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/1: p.StrmapIdent("0xc4200e0b40",&ast.Ident/*struct*/{
                      NamePos: 21378,
                      Name: "Value",
                      Obj: p.StrmapObject("0xc4200e28c0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Value",
                        Decl: p.PtrmapField("0xc4200de7c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 1*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e0b80",&ast.Ident/*struct*/{
                    NamePos: 21384,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200de800",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0bc0",&ast.Ident/*struct*/{
                      NamePos: 21423,
                      Name: "TokPos",
                      Obj: p.StrmapObject("0xc4200e2910",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "TokPos",
                        Decl: p.PtrmapField("0xc4200de800"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200e0c20",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e0be0",&ast.Ident/*struct*/{
                      NamePos: 21434,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e0c00",&ast.Ident/*struct*/{
                      NamePos: 21440,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200de840",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0c60",&ast.Ident/*struct*/{
                      NamePos: 21490,
                      Name: "Tok",
                      Obj: p.StrmapObject("0xc4200e2960",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Tok",
                        Decl: p.PtrmapField("0xc4200de840"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200e0cc0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e0c80",&ast.Ident/*struct*/{
                      NamePos: 21501,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e0ca0",&ast.Ident/*struct*/{
                      NamePos: 21507,
                      Name: "Token",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200de880",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0ce0",&ast.Ident/*struct*/{
                      NamePos: 21556,
                      Name: "X",
                      Obj: p.StrmapObject("0xc4200e29b0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "X",
                        Decl: p.PtrmapField("0xc4200de880"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200e0d00",&ast.Ident/*struct*/{
                    NamePos: 21567,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/5: p.StrmapField("0xc4200de900",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200e0d40",&ast.Ident/*struct*/{
                      NamePos: 21604,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200e2a00",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200de900"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200e0d80",&ast.StarExpr/*struct*/{
                    Star: 21615,
                    X: p.StrmapIdent("0xc4200e0d60",&ast.Ident/*struct*/{
                      NamePos: 21616,
                      Name: "BlockStmt",
                      Obj: p.PtrmapObject("0xc4200c9b30"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 5*/}/*slice*/,
              Closing: 21627,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 20*/}/*slice*/,
      Rparen: 21629,
    }/*struct*/),/* slice_item: 95*/96: p.StrmapFuncDecl("0xc4200dc8a0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dc7b0",&ast.FieldList/*struct*/{
        Opening: 21690,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200de9c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e0dc0",&ast.Ident/*struct*/{
                NamePos: 21691,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2a50",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200de9c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e0e00",&ast.StarExpr/*struct*/{
              Star: 21693,
              X: p.StrmapIdent("0xc4200e0de0",&ast.Ident/*struct*/{
                NamePos: 21694,
                Name: "BadStmt",
                Obj: p.PtrmapObject("0xc4200c8e10"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 21701,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e0e20",&ast.Ident/*struct*/{
        NamePos: 21703,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e0f20",&ast.FuncType/*struct*/{
        Func: 21685,
        Params: p.StrmapFieldList("0xc4200dc7e0",&ast.FieldList/*struct*/{
          Opening: 21706,
          Closing: 21707,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dc810",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dea00",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e0e80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e0e40",&ast.Ident/*struct*/{
                  NamePos: 21709,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e0e60",&ast.Ident/*struct*/{
                  NamePos: 21715,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dc870",&ast.BlockStmt/*struct*/{
        Lbrace: 21726,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e0f00",&ast.ReturnStmt/*struct*/{
            Return: 21728,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e0ee0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e0ea0",&ast.Ident/*struct*/{
                  NamePos: 21735,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2a50"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e0ec0",&ast.Ident/*struct*/{
                  NamePos: 21737,
                  Name: "From",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 21742,
      }/*struct*/),
    }/*struct*/),/* slice_item: 96*/97: p.StrmapFuncDecl("0xc4200dc9f0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dc900",&ast.FieldList/*struct*/{
        Opening: 21749,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dea40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e0f40",&ast.Ident/*struct*/{
                NamePos: 21750,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2aa0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dea40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e0f80",&ast.StarExpr/*struct*/{
              Star: 21752,
              X: p.StrmapIdent("0xc4200e0f60",&ast.Ident/*struct*/{
                NamePos: 21753,
                Name: "DeclStmt",
                Obj: p.PtrmapObject("0xc4200c8f00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 21761,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e0fa0",&ast.Ident/*struct*/{
        NamePos: 21763,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e10e0",&ast.FuncType/*struct*/{
        Func: 21744,
        Params: p.StrmapFieldList("0xc4200dc930",&ast.FieldList/*struct*/{
          Opening: 21766,
          Closing: 21767,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dc960",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dea80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e1000",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e0fc0",&ast.Ident/*struct*/{
                  NamePos: 21769,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e0fe0",&ast.Ident/*struct*/{
                  NamePos: 21775,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dc9c0",&ast.BlockStmt/*struct*/{
        Lbrace: 21785,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e10c0",&ast.ReturnStmt/*struct*/{
            Return: 21787,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200deac0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e10a0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e1060",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e1020",&ast.Ident/*struct*/{
                      NamePos: 21794,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e2aa0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e1040",&ast.Ident/*struct*/{
                      NamePos: 21796,
                      Name: "Decl",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e1080",&ast.Ident/*struct*/{
                    NamePos: 21801,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 21804,
                Ellipsis: 0,
                Rparen: 21805,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 21807,
      }/*struct*/),
    }/*struct*/),/* slice_item: 97*/98: p.StrmapFuncDecl("0xc4200dcb40",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dca50",&ast.FieldList/*struct*/{
        Opening: 21814,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200deb00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e1100",&ast.Ident/*struct*/{
                NamePos: 21815,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2af0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200deb00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e1140",&ast.StarExpr/*struct*/{
              Star: 21817,
              X: p.StrmapIdent("0xc4200e1120",&ast.Ident/*struct*/{
                NamePos: 21818,
                Name: "EmptyStmt",
                Obj: p.PtrmapObject("0xc4200c8ff0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 21827,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e1160",&ast.Ident/*struct*/{
        NamePos: 21829,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e1260",&ast.FuncType/*struct*/{
        Func: 21809,
        Params: p.StrmapFieldList("0xc4200dca80",&ast.FieldList/*struct*/{
          Opening: 21832,
          Closing: 21833,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dcab0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200deb40",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e11c0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1180",&ast.Ident/*struct*/{
                  NamePos: 21835,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e11a0",&ast.Ident/*struct*/{
                  NamePos: 21841,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dcb10",&ast.BlockStmt/*struct*/{
        Lbrace: 21850,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e1240",&ast.ReturnStmt/*struct*/{
            Return: 21852,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e1220",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e11e0",&ast.Ident/*struct*/{
                  NamePos: 21859,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2af0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1200",&ast.Ident/*struct*/{
                  NamePos: 21861,
                  Name: "Semicolon",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 21871,
      }/*struct*/),
    }/*struct*/),/* slice_item: 98*/99: p.StrmapFuncDecl("0xc4200dcc90",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dcba0",&ast.FieldList/*struct*/{
        Opening: 21878,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200deb80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e1280",&ast.Ident/*struct*/{
                NamePos: 21879,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2b40",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200deb80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e12c0",&ast.StarExpr/*struct*/{
              Star: 21881,
              X: p.StrmapIdent("0xc4200e12a0",&ast.Ident/*struct*/{
                NamePos: 21882,
                Name: "LabeledStmt",
                Obj: p.PtrmapObject("0xc4200c90e0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 21893,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e12e0",&ast.Ident/*struct*/{
        NamePos: 21895,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e1420",&ast.FuncType/*struct*/{
        Func: 21873,
        Params: p.StrmapFieldList("0xc4200dcbd0",&ast.FieldList/*struct*/{
          Opening: 21898,
          Closing: 21899,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dcc00",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200debc0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e1340",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1300",&ast.Ident/*struct*/{
                  NamePos: 21901,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1320",&ast.Ident/*struct*/{
                  NamePos: 21907,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dcc60",&ast.BlockStmt/*struct*/{
        Lbrace: 21914,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e1400",&ast.ReturnStmt/*struct*/{
            Return: 21916,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200dec00",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e13e0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e13a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e1360",&ast.Ident/*struct*/{
                      NamePos: 21923,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e2b40"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e1380",&ast.Ident/*struct*/{
                      NamePos: 21925,
                      Name: "Label",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e13c0",&ast.Ident/*struct*/{
                    NamePos: 21931,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 21934,
                Ellipsis: 0,
                Rparen: 21935,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 21937,
      }/*struct*/),
    }/*struct*/),/* slice_item: 99*/100: p.StrmapFuncDecl("0xc4200dcde0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dccf0",&ast.FieldList/*struct*/{
        Opening: 21944,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dec40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e1440",&ast.Ident/*struct*/{
                NamePos: 21945,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2b90",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dec40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e1480",&ast.StarExpr/*struct*/{
              Star: 21947,
              X: p.StrmapIdent("0xc4200e1460",&ast.Ident/*struct*/{
                NamePos: 21948,
                Name: "ExprStmt",
                Obj: p.PtrmapObject("0xc4200c9220"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 21956,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e14a0",&ast.Ident/*struct*/{
        NamePos: 21958,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e15e0",&ast.FuncType/*struct*/{
        Func: 21939,
        Params: p.StrmapFieldList("0xc4200dcd20",&ast.FieldList/*struct*/{
          Opening: 21961,
          Closing: 21962,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dcd50",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dec80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e1500",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e14c0",&ast.Ident/*struct*/{
                  NamePos: 21964,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e14e0",&ast.Ident/*struct*/{
                  NamePos: 21970,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dcdb0",&ast.BlockStmt/*struct*/{
        Lbrace: 21980,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e15c0",&ast.ReturnStmt/*struct*/{
            Return: 21982,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200decc0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e15a0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e1560",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e1520",&ast.Ident/*struct*/{
                      NamePos: 21989,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e2b90"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e1540",&ast.Ident/*struct*/{
                      NamePos: 21991,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e1580",&ast.Ident/*struct*/{
                    NamePos: 21993,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 21996,
                Ellipsis: 0,
                Rparen: 21997,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 21999,
      }/*struct*/),
    }/*struct*/),/* slice_item: 100*/101: p.StrmapFuncDecl("0xc4200dcf30",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dce40",&ast.FieldList/*struct*/{
        Opening: 22006,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ded00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e1600",&ast.Ident/*struct*/{
                NamePos: 22007,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2be0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ded00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e1640",&ast.StarExpr/*struct*/{
              Star: 22009,
              X: p.StrmapIdent("0xc4200e1620",&ast.Ident/*struct*/{
                NamePos: 22010,
                Name: "SendStmt",
                Obj: p.PtrmapObject("0xc4200c92c0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22018,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e1660",&ast.Ident/*struct*/{
        NamePos: 22020,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e17a0",&ast.FuncType/*struct*/{
        Func: 22001,
        Params: p.StrmapFieldList("0xc4200dce70",&ast.FieldList/*struct*/{
          Opening: 22023,
          Closing: 22024,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dcea0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ded40",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e16c0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1680",&ast.Ident/*struct*/{
                  NamePos: 22026,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e16a0",&ast.Ident/*struct*/{
                  NamePos: 22032,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dcf00",&ast.BlockStmt/*struct*/{
        Lbrace: 22042,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e1780",&ast.ReturnStmt/*struct*/{
            Return: 22044,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ded80",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e1760",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e1720",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e16e0",&ast.Ident/*struct*/{
                      NamePos: 22051,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e2be0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e1700",&ast.Ident/*struct*/{
                      NamePos: 22053,
                      Name: "Chan",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e1740",&ast.Ident/*struct*/{
                    NamePos: 22058,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 22061,
                Ellipsis: 0,
                Rparen: 22062,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22064,
      }/*struct*/),
    }/*struct*/),/* slice_item: 101*/102: p.StrmapFuncDecl("0xc4200dd080",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dcf90",&ast.FieldList/*struct*/{
        Opening: 22071,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dedc0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e17c0",&ast.Ident/*struct*/{
                NamePos: 22072,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2c30",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dedc0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e1800",&ast.StarExpr/*struct*/{
              Star: 22074,
              X: p.StrmapIdent("0xc4200e17e0",&ast.Ident/*struct*/{
                NamePos: 22075,
                Name: "IncDecStmt",
                Obj: p.PtrmapObject("0xc4200c9450"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22085,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e1820",&ast.Ident/*struct*/{
        NamePos: 22087,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e1960",&ast.FuncType/*struct*/{
        Func: 22066,
        Params: p.StrmapFieldList("0xc4200dcfc0",&ast.FieldList/*struct*/{
          Opening: 22090,
          Closing: 22091,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dcff0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dee00",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e1880",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1840",&ast.Ident/*struct*/{
                  NamePos: 22093,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1860",&ast.Ident/*struct*/{
                  NamePos: 22099,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dd050",&ast.BlockStmt/*struct*/{
        Lbrace: 22107,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e1940",&ast.ReturnStmt/*struct*/{
            Return: 22109,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200dee40",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e1920",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e18e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e18a0",&ast.Ident/*struct*/{
                      NamePos: 22116,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e2c30"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e18c0",&ast.Ident/*struct*/{
                      NamePos: 22118,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e1900",&ast.Ident/*struct*/{
                    NamePos: 22120,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 22123,
                Ellipsis: 0,
                Rparen: 22124,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22126,
      }/*struct*/),
    }/*struct*/),/* slice_item: 102*/103: p.StrmapFuncDecl("0xc4200dd200",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dd0e0",&ast.FieldList/*struct*/{
        Opening: 22133,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dee80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e1980",&ast.Ident/*struct*/{
                NamePos: 22134,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2c80",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dee80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e19c0",&ast.StarExpr/*struct*/{
              Star: 22136,
              X: p.StrmapIdent("0xc4200e19a0",&ast.Ident/*struct*/{
                NamePos: 22137,
                Name: "AssignStmt",
                Obj: p.PtrmapObject("0xc4200c9590"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22147,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e19e0",&ast.Ident/*struct*/{
        NamePos: 22149,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e1b40",&ast.FuncType/*struct*/{
        Func: 22128,
        Params: p.StrmapFieldList("0xc4200dd110",&ast.FieldList/*struct*/{
          Opening: 22152,
          Closing: 22153,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dd140",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200deec0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e1a40",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1a00",&ast.Ident/*struct*/{
                  NamePos: 22155,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1a20",&ast.Ident/*struct*/{
                  NamePos: 22161,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dd1d0",&ast.BlockStmt/*struct*/{
        Lbrace: 22169,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e1b20",&ast.ReturnStmt/*struct*/{
            Return: 22171,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200def00",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e1b00",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIndexExpr("0xc4200dd1a0",&ast.IndexExpr/*struct*/{
                    X: p.StrmapSelectorExpr("0xc4200e1aa0",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200e1a60",&ast.Ident/*struct*/{
                        NamePos: 22178,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200e2c80"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200e1a80",&ast.Ident/*struct*/{
                        NamePos: 22180,
                        Name: "Lhs",
                      }/*struct*/),
                    }/*struct*/),
                    Lbrack: 22183,
                    Index: p.StrmapBasicLit("0xc4200e1ac0",&ast.BasicLit/*struct*/{
                      ValuePos: 22184,
                      Kind: token.Token(5)/*INT*/,
                      Value: "0",
                    }/*struct*/),
                    Rbrack: 22185,
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e1ae0",&ast.Ident/*struct*/{
                    NamePos: 22187,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 22190,
                Ellipsis: 0,
                Rparen: 22191,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22193,
      }/*struct*/),
    }/*struct*/),/* slice_item: 103*/104: p.StrmapFuncDecl("0xc4200dd350",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dd260",&ast.FieldList/*struct*/{
        Opening: 22200,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200def40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e1b60",&ast.Ident/*struct*/{
                NamePos: 22201,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2cd0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200def40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e1ba0",&ast.StarExpr/*struct*/{
              Star: 22203,
              X: p.StrmapIdent("0xc4200e1b80",&ast.Ident/*struct*/{
                NamePos: 22204,
                Name: "GoStmt",
                Obj: p.PtrmapObject("0xc4200c9720"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22210,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e1bc0",&ast.Ident/*struct*/{
        NamePos: 22212,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e1cc0",&ast.FuncType/*struct*/{
        Func: 22195,
        Params: p.StrmapFieldList("0xc4200dd290",&ast.FieldList/*struct*/{
          Opening: 22215,
          Closing: 22216,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dd2c0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200def80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e1c20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1be0",&ast.Ident/*struct*/{
                  NamePos: 22218,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1c00",&ast.Ident/*struct*/{
                  NamePos: 22224,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dd320",&ast.BlockStmt/*struct*/{
        Lbrace: 22236,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e1ca0",&ast.ReturnStmt/*struct*/{
            Return: 22238,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e1c80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1c40",&ast.Ident/*struct*/{
                  NamePos: 22245,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2cd0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1c60",&ast.Ident/*struct*/{
                  NamePos: 22247,
                  Name: "Go",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22250,
      }/*struct*/),
    }/*struct*/),/* slice_item: 104*/105: p.StrmapFuncDecl("0xc4200dd4a0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dd3b0",&ast.FieldList/*struct*/{
        Opening: 22257,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200defc0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e1ce0",&ast.Ident/*struct*/{
                NamePos: 22258,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2d20",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200defc0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e1d20",&ast.StarExpr/*struct*/{
              Star: 22260,
              X: p.StrmapIdent("0xc4200e1d00",&ast.Ident/*struct*/{
                NamePos: 22261,
                Name: "DeferStmt",
                Obj: p.PtrmapObject("0xc4200c9810"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22270,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e1d40",&ast.Ident/*struct*/{
        NamePos: 22272,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e1e40",&ast.FuncType/*struct*/{
        Func: 22252,
        Params: p.StrmapFieldList("0xc4200dd3e0",&ast.FieldList/*struct*/{
          Opening: 22275,
          Closing: 22276,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dd410",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df000",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e1da0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1d60",&ast.Ident/*struct*/{
                  NamePos: 22278,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1d80",&ast.Ident/*struct*/{
                  NamePos: 22284,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dd470",&ast.BlockStmt/*struct*/{
        Lbrace: 22293,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e1e20",&ast.ReturnStmt/*struct*/{
            Return: 22295,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e1e00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1dc0",&ast.Ident/*struct*/{
                  NamePos: 22302,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2d20"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1de0",&ast.Ident/*struct*/{
                  NamePos: 22304,
                  Name: "Defer",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22310,
      }/*struct*/),
    }/*struct*/),/* slice_item: 105*/106: p.StrmapFuncDecl("0xc4200dd5f0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dd500",&ast.FieldList/*struct*/{
        Opening: 22317,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df040",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e1e60",&ast.Ident/*struct*/{
                NamePos: 22318,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2d70",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df040"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e1ea0",&ast.StarExpr/*struct*/{
              Star: 22320,
              X: p.StrmapIdent("0xc4200e1e80",&ast.Ident/*struct*/{
                NamePos: 22321,
                Name: "ReturnStmt",
                Obj: p.PtrmapObject("0xc4200c9900"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22331,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e1ec0",&ast.Ident/*struct*/{
        NamePos: 22333,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e1fc0",&ast.FuncType/*struct*/{
        Func: 22312,
        Params: p.StrmapFieldList("0xc4200dd530",&ast.FieldList/*struct*/{
          Opening: 22336,
          Closing: 22337,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dd560",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df080",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e1f20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1ee0",&ast.Ident/*struct*/{
                  NamePos: 22339,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1f00",&ast.Ident/*struct*/{
                  NamePos: 22345,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dd5c0",&ast.BlockStmt/*struct*/{
        Lbrace: 22353,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e1fa0",&ast.ReturnStmt/*struct*/{
            Return: 22355,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e1f80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e1f40",&ast.Ident/*struct*/{
                  NamePos: 22362,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2d70"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e1f60",&ast.Ident/*struct*/{
                  NamePos: 22364,
                  Name: "Return",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22371,
      }/*struct*/),
    }/*struct*/),/* slice_item: 106*/107: p.StrmapFuncDecl("0xc4200dd740",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dd650",&ast.FieldList/*struct*/{
        Opening: 22378,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df0c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e1fe0",&ast.Ident/*struct*/{
                NamePos: 22379,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2dc0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df0c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e4020",&ast.StarExpr/*struct*/{
              Star: 22381,
              X: p.StrmapIdent("0xc4200e4000",&ast.Ident/*struct*/{
                NamePos: 22382,
                Name: "BranchStmt",
                Obj: p.PtrmapObject("0xc4200c99f0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22392,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e4040",&ast.Ident/*struct*/{
        NamePos: 22394,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e4140",&ast.FuncType/*struct*/{
        Func: 22373,
        Params: p.StrmapFieldList("0xc4200dd680",&ast.FieldList/*struct*/{
          Opening: 22397,
          Closing: 22398,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dd6b0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df100",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e40a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4060",&ast.Ident/*struct*/{
                  NamePos: 22400,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4080",&ast.Ident/*struct*/{
                  NamePos: 22406,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dd710",&ast.BlockStmt/*struct*/{
        Lbrace: 22414,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e4120",&ast.ReturnStmt/*struct*/{
            Return: 22416,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4100",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e40c0",&ast.Ident/*struct*/{
                  NamePos: 22423,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2dc0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e40e0",&ast.Ident/*struct*/{
                  NamePos: 22425,
                  Name: "TokPos",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22432,
      }/*struct*/),
    }/*struct*/),/* slice_item: 107*/108: p.StrmapFuncDecl("0xc4200dd890",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dd7a0",&ast.FieldList/*struct*/{
        Opening: 22439,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df140",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e4160",&ast.Ident/*struct*/{
                NamePos: 22440,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2e10",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df140"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e41a0",&ast.StarExpr/*struct*/{
              Star: 22442,
              X: p.StrmapIdent("0xc4200e4180",&ast.Ident/*struct*/{
                NamePos: 22443,
                Name: "BlockStmt",
                Obj: p.PtrmapObject("0xc4200c9b30"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22452,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e41c0",&ast.Ident/*struct*/{
        NamePos: 22454,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e42c0",&ast.FuncType/*struct*/{
        Func: 22434,
        Params: p.StrmapFieldList("0xc4200dd7d0",&ast.FieldList/*struct*/{
          Opening: 22457,
          Closing: 22458,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dd800",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df180",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e4220",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e41e0",&ast.Ident/*struct*/{
                  NamePos: 22460,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4200",&ast.Ident/*struct*/{
                  NamePos: 22466,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dd860",&ast.BlockStmt/*struct*/{
        Lbrace: 22475,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e42a0",&ast.ReturnStmt/*struct*/{
            Return: 22477,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4280",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4240",&ast.Ident/*struct*/{
                  NamePos: 22484,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2e10"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4260",&ast.Ident/*struct*/{
                  NamePos: 22486,
                  Name: "Lbrace",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22493,
      }/*struct*/),
    }/*struct*/),/* slice_item: 108*/109: p.StrmapFuncDecl("0xc4200dd9e0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dd8f0",&ast.FieldList/*struct*/{
        Opening: 22500,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df1c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e42e0",&ast.Ident/*struct*/{
                NamePos: 22501,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2e60",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df1c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e4320",&ast.StarExpr/*struct*/{
              Star: 22503,
              X: p.StrmapIdent("0xc4200e4300",&ast.Ident/*struct*/{
                NamePos: 22504,
                Name: "IfStmt",
                Obj: p.PtrmapObject("0xc4200c9c70"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22510,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e4340",&ast.Ident/*struct*/{
        NamePos: 22512,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e4440",&ast.FuncType/*struct*/{
        Func: 22495,
        Params: p.StrmapFieldList("0xc4200dd920",&ast.FieldList/*struct*/{
          Opening: 22515,
          Closing: 22516,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dd950",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df200",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e43a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4360",&ast.Ident/*struct*/{
                  NamePos: 22518,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4380",&ast.Ident/*struct*/{
                  NamePos: 22524,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200dd9b0",&ast.BlockStmt/*struct*/{
        Lbrace: 22536,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e4420",&ast.ReturnStmt/*struct*/{
            Return: 22538,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4400",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e43c0",&ast.Ident/*struct*/{
                  NamePos: 22545,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2e60"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e43e0",&ast.Ident/*struct*/{
                  NamePos: 22547,
                  Name: "If",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22550,
      }/*struct*/),
    }/*struct*/),/* slice_item: 109*/110: p.StrmapFuncDecl("0xc4200ddb30",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dda40",&ast.FieldList/*struct*/{
        Opening: 22557,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df240",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e4460",&ast.Ident/*struct*/{
                NamePos: 22558,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2eb0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df240"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e44a0",&ast.StarExpr/*struct*/{
              Star: 22560,
              X: p.StrmapIdent("0xc4200e4480",&ast.Ident/*struct*/{
                NamePos: 22561,
                Name: "CaseClause",
                Obj: p.PtrmapObject("0xc4200c9ea0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22571,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e44c0",&ast.Ident/*struct*/{
        NamePos: 22573,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e45c0",&ast.FuncType/*struct*/{
        Func: 22552,
        Params: p.StrmapFieldList("0xc4200dda70",&ast.FieldList/*struct*/{
          Opening: 22576,
          Closing: 22577,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200ddaa0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df280",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e4520",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e44e0",&ast.Ident/*struct*/{
                  NamePos: 22579,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4500",&ast.Ident/*struct*/{
                  NamePos: 22585,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200ddb00",&ast.BlockStmt/*struct*/{
        Lbrace: 22593,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e45a0",&ast.ReturnStmt/*struct*/{
            Return: 22595,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4580",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4540",&ast.Ident/*struct*/{
                  NamePos: 22602,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2eb0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4560",&ast.Ident/*struct*/{
                  NamePos: 22604,
                  Name: "Case",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22609,
      }/*struct*/),
    }/*struct*/),/* slice_item: 110*/111: p.StrmapFuncDecl("0xc4200ddc80",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200ddb90",&ast.FieldList/*struct*/{
        Opening: 22616,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df2c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e45e0",&ast.Ident/*struct*/{
                NamePos: 22617,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2f00",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df2c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e4620",&ast.StarExpr/*struct*/{
              Star: 22619,
              X: p.StrmapIdent("0xc4200e4600",&ast.Ident/*struct*/{
                NamePos: 22620,
                Name: "SwitchStmt",
                Obj: p.PtrmapObject("0xc4200e2050"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22630,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e4640",&ast.Ident/*struct*/{
        NamePos: 22632,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e4740",&ast.FuncType/*struct*/{
        Func: 22611,
        Params: p.StrmapFieldList("0xc4200ddbc0",&ast.FieldList/*struct*/{
          Opening: 22635,
          Closing: 22636,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200ddbf0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df300",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e46a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4660",&ast.Ident/*struct*/{
                  NamePos: 22638,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4680",&ast.Ident/*struct*/{
                  NamePos: 22644,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200ddc50",&ast.BlockStmt/*struct*/{
        Lbrace: 22652,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e4720",&ast.ReturnStmt/*struct*/{
            Return: 22654,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4700",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e46c0",&ast.Ident/*struct*/{
                  NamePos: 22661,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2f00"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e46e0",&ast.Ident/*struct*/{
                  NamePos: 22663,
                  Name: "Switch",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22670,
      }/*struct*/),
    }/*struct*/),/* slice_item: 111*/112: p.StrmapFuncDecl("0xc4200dddd0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200ddce0",&ast.FieldList/*struct*/{
        Opening: 22677,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df340",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e4760",&ast.Ident/*struct*/{
                NamePos: 22678,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2f50",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df340"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e47a0",&ast.StarExpr/*struct*/{
              Star: 22680,
              X: p.StrmapIdent("0xc4200e4780",&ast.Ident/*struct*/{
                NamePos: 22681,
                Name: "TypeSwitchStmt",
                Obj: p.PtrmapObject("0xc4200e21e0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22695,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e47c0",&ast.Ident/*struct*/{
        NamePos: 22697,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e48c0",&ast.FuncType/*struct*/{
        Func: 22672,
        Params: p.StrmapFieldList("0xc4200ddd10",&ast.FieldList/*struct*/{
          Opening: 22700,
          Closing: 22701,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200ddd40",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df380",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e4820",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e47e0",&ast.Ident/*struct*/{
                  NamePos: 22703,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4800",&ast.Ident/*struct*/{
                  NamePos: 22709,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200ddda0",&ast.BlockStmt/*struct*/{
        Lbrace: 22713,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e48a0",&ast.ReturnStmt/*struct*/{
            Return: 22715,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4880",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4840",&ast.Ident/*struct*/{
                  NamePos: 22722,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2f50"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4860",&ast.Ident/*struct*/{
                  NamePos: 22724,
                  Name: "Switch",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22731,
      }/*struct*/),
    }/*struct*/),/* slice_item: 112*/113: p.StrmapFuncDecl("0xc4200ddf20",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200dde30",&ast.FieldList/*struct*/{
        Opening: 22738,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df3c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e48e0",&ast.Ident/*struct*/{
                NamePos: 22739,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2fa0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df3c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e4920",&ast.StarExpr/*struct*/{
              Star: 22741,
              X: p.StrmapIdent("0xc4200e4900",&ast.Ident/*struct*/{
                NamePos: 22742,
                Name: "CommClause",
                Obj: p.PtrmapObject("0xc4200e2370"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22752,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e4940",&ast.Ident/*struct*/{
        NamePos: 22754,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e4a40",&ast.FuncType/*struct*/{
        Func: 22733,
        Params: p.StrmapFieldList("0xc4200dde60",&ast.FieldList/*struct*/{
          Opening: 22757,
          Closing: 22758,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200dde90",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df400",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e49a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4960",&ast.Ident/*struct*/{
                  NamePos: 22760,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4980",&ast.Ident/*struct*/{
                  NamePos: 22766,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200ddef0",&ast.BlockStmt/*struct*/{
        Lbrace: 22774,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e4a20",&ast.ReturnStmt/*struct*/{
            Return: 22776,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4a00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e49c0",&ast.Ident/*struct*/{
                  NamePos: 22783,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2fa0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e49e0",&ast.Ident/*struct*/{
                  NamePos: 22785,
                  Name: "Case",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22790,
      }/*struct*/),
    }/*struct*/),/* slice_item: 113*/114: p.StrmapFuncDecl("0xc4200e8090",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200ddf80",&ast.FieldList/*struct*/{
        Opening: 22797,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df440",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e4a60",&ast.Ident/*struct*/{
                NamePos: 22798,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e2ff0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df440"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e4aa0",&ast.StarExpr/*struct*/{
              Star: 22800,
              X: p.StrmapIdent("0xc4200e4a80",&ast.Ident/*struct*/{
                NamePos: 22801,
                Name: "SelectStmt",
                Obj: p.PtrmapObject("0xc4200e2500"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22811,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e4ac0",&ast.Ident/*struct*/{
        NamePos: 22813,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e4bc0",&ast.FuncType/*struct*/{
        Func: 22792,
        Params: p.StrmapFieldList("0xc4200ddfb0",&ast.FieldList/*struct*/{
          Opening: 22816,
          Closing: 22817,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e8000",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df480",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e4b20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4ae0",&ast.Ident/*struct*/{
                  NamePos: 22819,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4b00",&ast.Ident/*struct*/{
                  NamePos: 22825,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e8060",&ast.BlockStmt/*struct*/{
        Lbrace: 22833,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e4ba0",&ast.ReturnStmt/*struct*/{
            Return: 22835,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4b80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4b40",&ast.Ident/*struct*/{
                  NamePos: 22842,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e2ff0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4b60",&ast.Ident/*struct*/{
                  NamePos: 22844,
                  Name: "Select",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22851,
      }/*struct*/),
    }/*struct*/),/* slice_item: 114*/115: p.StrmapFuncDecl("0xc4200e81e0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e80f0",&ast.FieldList/*struct*/{
        Opening: 22858,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df4c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e4be0",&ast.Ident/*struct*/{
                NamePos: 22859,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3040",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df4c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e4c20",&ast.StarExpr/*struct*/{
              Star: 22861,
              X: p.StrmapIdent("0xc4200e4c00",&ast.Ident/*struct*/{
                NamePos: 22862,
                Name: "ForStmt",
                Obj: p.PtrmapObject("0xc4200e25f0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22869,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e4c40",&ast.Ident/*struct*/{
        NamePos: 22871,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e4d40",&ast.FuncType/*struct*/{
        Func: 22853,
        Params: p.StrmapFieldList("0xc4200e8120",&ast.FieldList/*struct*/{
          Opening: 22874,
          Closing: 22875,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e8150",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df500",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e4ca0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4c60",&ast.Ident/*struct*/{
                  NamePos: 22877,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4c80",&ast.Ident/*struct*/{
                  NamePos: 22883,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e81b0",&ast.BlockStmt/*struct*/{
        Lbrace: 22894,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e4d20",&ast.ReturnStmt/*struct*/{
            Return: 22896,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4d00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4cc0",&ast.Ident/*struct*/{
                  NamePos: 22903,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e3040"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4ce0",&ast.Ident/*struct*/{
                  NamePos: 22905,
                  Name: "For",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22909,
      }/*struct*/),
    }/*struct*/),/* slice_item: 115*/116: p.StrmapFuncDecl("0xc4200e8330",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e8240",&ast.FieldList/*struct*/{
        Opening: 22916,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df540",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e4d60",&ast.Ident/*struct*/{
                NamePos: 22917,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3090",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df540"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e4da0",&ast.StarExpr/*struct*/{
              Star: 22919,
              X: p.StrmapIdent("0xc4200e4d80",&ast.Ident/*struct*/{
                NamePos: 22920,
                Name: "RangeStmt",
                Obj: p.PtrmapObject("0xc4200e27d0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22929,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e4dc0",&ast.Ident/*struct*/{
        NamePos: 22931,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e4ec0",&ast.FuncType/*struct*/{
        Func: 22911,
        Params: p.StrmapFieldList("0xc4200e8270",&ast.FieldList/*struct*/{
          Opening: 22934,
          Closing: 22935,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e82a0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df580",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e4e20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4de0",&ast.Ident/*struct*/{
                  NamePos: 22937,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4e00",&ast.Ident/*struct*/{
                  NamePos: 22943,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e8300",&ast.BlockStmt/*struct*/{
        Lbrace: 22952,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e4ea0",&ast.ReturnStmt/*struct*/{
            Return: 22954,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e4e80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4e40",&ast.Ident/*struct*/{
                  NamePos: 22961,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e3090"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4e60",&ast.Ident/*struct*/{
                  NamePos: 22963,
                  Name: "For",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 22967,
      }/*struct*/),
    }/*struct*/),/* slice_item: 116*/117: p.StrmapFuncDecl("0xc4200e8480",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e8390",&ast.FieldList/*struct*/{
        Opening: 22975,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df5c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e4ee0",&ast.Ident/*struct*/{
                NamePos: 22976,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e30e0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df5c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e4f20",&ast.StarExpr/*struct*/{
              Star: 22978,
              X: p.StrmapIdent("0xc4200e4f00",&ast.Ident/*struct*/{
                NamePos: 22979,
                Name: "BadStmt",
                Obj: p.PtrmapObject("0xc4200c8e10"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 22986,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e4f40",&ast.Ident/*struct*/{
        NamePos: 22988,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e5040",&ast.FuncType/*struct*/{
        Func: 22970,
        Params: p.StrmapFieldList("0xc4200e83c0",&ast.FieldList/*struct*/{
          Opening: 22991,
          Closing: 22992,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e83f0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df600",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e4fa0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4f60",&ast.Ident/*struct*/{
                  NamePos: 22994,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4f80",&ast.Ident/*struct*/{
                  NamePos: 23000,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e8450",&ast.BlockStmt/*struct*/{
        Lbrace: 23005,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e5020",&ast.ReturnStmt/*struct*/{
            Return: 23007,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200e5000",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e4fc0",&ast.Ident/*struct*/{
                  NamePos: 23014,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e30e0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e4fe0",&ast.Ident/*struct*/{
                  NamePos: 23016,
                  Name: "To",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 23019,
      }/*struct*/),
    }/*struct*/),/* slice_item: 117*/118: p.StrmapFuncDecl("0xc4200e85d0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e84e0",&ast.FieldList/*struct*/{
        Opening: 23026,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df640",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e5060",&ast.Ident/*struct*/{
                NamePos: 23027,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3130",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df640"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e50a0",&ast.StarExpr/*struct*/{
              Star: 23029,
              X: p.StrmapIdent("0xc4200e5080",&ast.Ident/*struct*/{
                NamePos: 23030,
                Name: "DeclStmt",
                Obj: p.PtrmapObject("0xc4200c8f00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23038,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e50c0",&ast.Ident/*struct*/{
        NamePos: 23040,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e5200",&ast.FuncType/*struct*/{
        Func: 23021,
        Params: p.StrmapFieldList("0xc4200e8510",&ast.FieldList/*struct*/{
          Opening: 23043,
          Closing: 23044,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e8540",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df680",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e5120",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e50e0",&ast.Ident/*struct*/{
                  NamePos: 23046,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e5100",&ast.Ident/*struct*/{
                  NamePos: 23052,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e85a0",&ast.BlockStmt/*struct*/{
        Lbrace: 23056,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e51e0",&ast.ReturnStmt/*struct*/{
            Return: 23058,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200df6c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e51c0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e5180",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e5140",&ast.Ident/*struct*/{
                      NamePos: 23065,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e3130"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e5160",&ast.Ident/*struct*/{
                      NamePos: 23067,
                      Name: "Decl",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e51a0",&ast.Ident/*struct*/{
                    NamePos: 23072,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 23075,
                Ellipsis: 0,
                Rparen: 23076,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 23078,
      }/*struct*/),
    }/*struct*/),/* slice_item: 118*/119: p.StrmapFuncDecl("0xc4200e87e0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e8630",&ast.FieldList/*struct*/{
        Opening: 23085,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df700",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e5220",&ast.Ident/*struct*/{
                NamePos: 23086,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3180",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df700"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e5260",&ast.StarExpr/*struct*/{
              Star: 23088,
              X: p.StrmapIdent("0xc4200e5240",&ast.Ident/*struct*/{
                NamePos: 23089,
                Name: "EmptyStmt",
                Obj: p.PtrmapObject("0xc4200c8ff0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23098,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e5280",&ast.Ident/*struct*/{
        NamePos: 23100,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e54a0",&ast.FuncType/*struct*/{
        Func: 23080,
        Params: p.StrmapFieldList("0xc4200e8660",&ast.FieldList/*struct*/{
          Opening: 23103,
          Closing: 23104,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e8690",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df740",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e52e0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e52a0",&ast.Ident/*struct*/{
                  NamePos: 23106,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e52c0",&ast.Ident/*struct*/{
                  NamePos: 23112,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e87b0",&ast.BlockStmt/*struct*/{
        Lbrace: 23116,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200df780",&ast.IfStmt/*struct*/{
            If: 23119,
            Cond: p.StrmapSelectorExpr("0xc4200e5340",&ast.SelectorExpr/*struct*/{
              X: p.StrmapIdent("0xc4200e5300",&ast.Ident/*struct*/{
                NamePos: 23122,
                Name: "s",
                Obj: p.PtrmapObject("0xc4200e3180"),
              }/*struct*/),
              Sel: p.StrmapIdent("0xc4200e5320",&ast.Ident/*struct*/{
                NamePos: 23124,
                Name: "Implicit",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200e8750",&ast.BlockStmt/*struct*/{
              Lbrace: 23133,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200e53c0",&ast.ReturnStmt/*struct*/{
                  Return: 23137,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200e53a0",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200e5360",&ast.Ident/*struct*/{
                        NamePos: 23144,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200e3180"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200e5380",&ast.Ident/*struct*/{
                        NamePos: 23146,
                        Name: "Semicolon",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 23157,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200e5460",&ast.ReturnStmt/*struct*/{
            Return: 23160,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200e8780",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200e5420",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200e53e0",&ast.Ident/*struct*/{
                    NamePos: 23167,
                    Name: "s",
                    Obj: p.PtrmapObject("0xc4200e3180"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e5400",&ast.Ident/*struct*/{
                    NamePos: 23169,
                    Name: "Semicolon",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 23179,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200e5440",&ast.BasicLit/*struct*/{
                  ValuePos: 23181,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 23198,
      }/*struct*/),
    }/*struct*/),/* slice_item: 119*/120: p.StrmapFuncDecl("0xc4200e8930",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e8840",&ast.FieldList/*struct*/{
        Opening: 23205,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df7c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e54c0",&ast.Ident/*struct*/{
                NamePos: 23206,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e31d0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df7c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e5500",&ast.StarExpr/*struct*/{
              Star: 23208,
              X: p.StrmapIdent("0xc4200e54e0",&ast.Ident/*struct*/{
                NamePos: 23209,
                Name: "LabeledStmt",
                Obj: p.PtrmapObject("0xc4200c90e0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23220,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e5520",&ast.Ident/*struct*/{
        NamePos: 23222,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e5660",&ast.FuncType/*struct*/{
        Func: 23200,
        Params: p.StrmapFieldList("0xc4200e8870",&ast.FieldList/*struct*/{
          Opening: 23225,
          Closing: 23226,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e88a0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df800",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e5580",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e5540",&ast.Ident/*struct*/{
                  NamePos: 23228,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e5560",&ast.Ident/*struct*/{
                  NamePos: 23234,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e8900",&ast.BlockStmt/*struct*/{
        Lbrace: 23238,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e5640",&ast.ReturnStmt/*struct*/{
            Return: 23240,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200df840",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e5620",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e55e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e55a0",&ast.Ident/*struct*/{
                      NamePos: 23247,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e31d0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e55c0",&ast.Ident/*struct*/{
                      NamePos: 23249,
                      Name: "Stmt",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e5600",&ast.Ident/*struct*/{
                    NamePos: 23254,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 23257,
                Ellipsis: 0,
                Rparen: 23258,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 23260,
      }/*struct*/),
    }/*struct*/),/* slice_item: 120*/121: p.StrmapFuncDecl("0xc4200e8a80",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e8990",&ast.FieldList/*struct*/{
        Opening: 23267,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df880",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e5680",&ast.Ident/*struct*/{
                NamePos: 23268,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3220",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df880"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e56c0",&ast.StarExpr/*struct*/{
              Star: 23270,
              X: p.StrmapIdent("0xc4200e56a0",&ast.Ident/*struct*/{
                NamePos: 23271,
                Name: "ExprStmt",
                Obj: p.PtrmapObject("0xc4200c9220"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23279,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e56e0",&ast.Ident/*struct*/{
        NamePos: 23281,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e5820",&ast.FuncType/*struct*/{
        Func: 23262,
        Params: p.StrmapFieldList("0xc4200e89c0",&ast.FieldList/*struct*/{
          Opening: 23284,
          Closing: 23285,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e89f0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df8c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e5740",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e5700",&ast.Ident/*struct*/{
                  NamePos: 23287,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e5720",&ast.Ident/*struct*/{
                  NamePos: 23293,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e8a50",&ast.BlockStmt/*struct*/{
        Lbrace: 23300,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e5800",&ast.ReturnStmt/*struct*/{
            Return: 23302,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200df900",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e57e0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e57a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e5760",&ast.Ident/*struct*/{
                      NamePos: 23309,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e3220"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e5780",&ast.Ident/*struct*/{
                      NamePos: 23311,
                      Name: "X",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e57c0",&ast.Ident/*struct*/{
                    NamePos: 23313,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 23316,
                Ellipsis: 0,
                Rparen: 23317,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 23319,
      }/*struct*/),
    }/*struct*/),/* slice_item: 121*/122: p.StrmapFuncDecl("0xc4200e8bd0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e8ae0",&ast.FieldList/*struct*/{
        Opening: 23326,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200df940",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e5840",&ast.Ident/*struct*/{
                NamePos: 23327,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3270",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200df940"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e5880",&ast.StarExpr/*struct*/{
              Star: 23329,
              X: p.StrmapIdent("0xc4200e5860",&ast.Ident/*struct*/{
                NamePos: 23330,
                Name: "SendStmt",
                Obj: p.PtrmapObject("0xc4200c92c0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23338,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e58a0",&ast.Ident/*struct*/{
        NamePos: 23340,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e59e0",&ast.FuncType/*struct*/{
        Func: 23321,
        Params: p.StrmapFieldList("0xc4200e8b10",&ast.FieldList/*struct*/{
          Opening: 23343,
          Closing: 23344,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e8b40",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200df980",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e5900",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e58c0",&ast.Ident/*struct*/{
                  NamePos: 23346,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e58e0",&ast.Ident/*struct*/{
                  NamePos: 23352,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e8ba0",&ast.BlockStmt/*struct*/{
        Lbrace: 23359,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e59c0",&ast.ReturnStmt/*struct*/{
            Return: 23361,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200df9c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e59a0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e5960",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e5920",&ast.Ident/*struct*/{
                      NamePos: 23368,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e3270"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e5940",&ast.Ident/*struct*/{
                      NamePos: 23370,
                      Name: "Value",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e5980",&ast.Ident/*struct*/{
                    NamePos: 23376,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 23379,
                Ellipsis: 0,
                Rparen: 23380,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 23382,
      }/*struct*/),
    }/*struct*/),/* slice_item: 122*/123: p.StrmapFuncDecl("0xc4200e8d50",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e8c30",&ast.FieldList/*struct*/{
        Opening: 23389,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dfa00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e5a00",&ast.Ident/*struct*/{
                NamePos: 23390,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e32c0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dfa00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e5a40",&ast.StarExpr/*struct*/{
              Star: 23392,
              X: p.StrmapIdent("0xc4200e5a20",&ast.Ident/*struct*/{
                NamePos: 23393,
                Name: "IncDecStmt",
                Obj: p.PtrmapObject("0xc4200c9450"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23403,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e5a60",&ast.Ident/*struct*/{
        NamePos: 23405,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e5b80",&ast.FuncType/*struct*/{
        Func: 23384,
        Params: p.StrmapFieldList("0xc4200e8c60",&ast.FieldList/*struct*/{
          Opening: 23408,
          Closing: 23409,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e8c90",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dfa40",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e5ac0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e5a80",&ast.Ident/*struct*/{
                  NamePos: 23411,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e5aa0",&ast.Ident/*struct*/{
                  NamePos: 23417,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e8d20",&ast.BlockStmt/*struct*/{
        Lbrace: 23421,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e5b60",&ast.ReturnStmt/*struct*/{
            Return: 23424,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200e8cf0",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200e5b20",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200e5ae0",&ast.Ident/*struct*/{
                    NamePos: 23431,
                    Name: "s",
                    Obj: p.PtrmapObject("0xc4200e32c0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e5b00",&ast.Ident/*struct*/{
                    NamePos: 23433,
                    Name: "TokPos",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 23440,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200e5b40",&ast.BasicLit/*struct*/{
                  ValuePos: 23442,
                  Kind: token.Token(5)/*INT*/,
                  Value: "2",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 23460,
      }/*struct*/),
    }/*struct*/),/* slice_item: 123*/124: p.StrmapFuncDecl("0xc4200e8f00",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e8db0",&ast.FieldList/*struct*/{
        Opening: 23467,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dfa80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e5ba0",&ast.Ident/*struct*/{
                NamePos: 23468,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3310",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dfa80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e5be0",&ast.StarExpr/*struct*/{
              Star: 23470,
              X: p.StrmapIdent("0xc4200e5bc0",&ast.Ident/*struct*/{
                NamePos: 23471,
                Name: "AssignStmt",
                Obj: p.PtrmapObject("0xc4200c9590"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23481,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e5c00",&ast.Ident/*struct*/{
        NamePos: 23483,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e5de0",&ast.FuncType/*struct*/{
        Func: 23462,
        Params: p.StrmapFieldList("0xc4200e8de0",&ast.FieldList/*struct*/{
          Opening: 23486,
          Closing: 23487,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e8e10",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dfac0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e5c60",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e5c20",&ast.Ident/*struct*/{
                  NamePos: 23489,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e5c40",&ast.Ident/*struct*/{
                  NamePos: 23495,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e8ed0",&ast.BlockStmt/*struct*/{
        Lbrace: 23499,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e5dc0",&ast.ReturnStmt/*struct*/{
            Return: 23501,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200dfb40",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e5da0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIndexExpr("0xc4200e8ea0",&ast.IndexExpr/*struct*/{
                    X: p.StrmapSelectorExpr("0xc4200e5cc0",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200e5c80",&ast.Ident/*struct*/{
                        NamePos: 23508,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200e3310"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200e5ca0",&ast.Ident/*struct*/{
                        NamePos: 23510,
                        Name: "Rhs",
                      }/*struct*/),
                    }/*struct*/),
                    Lbrack: 23513,
                    Index: p.StrmapBinaryExpr("0xc4200e8e70",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapCallExpr("0xc4200dfb00",&ast.CallExpr/*struct*/{
                        Fun: p.StrmapIdent("0xc4200e5ce0",&ast.Ident/*struct*/{
                          NamePos: 23514,
                          Name: "len",
                        }/*struct*/),
                        Lparen: 23517,
                        Args: []ast.Expr /*Slice*/{
                          0: p.StrmapSelectorExpr("0xc4200e5d40",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200e5d00",&ast.Ident/*struct*/{
                              NamePos: 23518,
                              Name: "s",
                              Obj: p.PtrmapObject("0xc4200e3310"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200e5d20",&ast.Ident/*struct*/{
                              NamePos: 23520,
                              Name: "Rhs",
                            }/*struct*/),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                        Ellipsis: 0,
                        Rparen: 23523,
                      }/*struct*/),
                      OpPos: 23524,
                      Op: token.Token(13)/*-*/,
                      Y: p.StrmapBasicLit("0xc4200e5d60",&ast.BasicLit/*struct*/{
                        ValuePos: 23525,
                        Kind: token.Token(5)/*INT*/,
                        Value: "1",
                      }/*struct*/),
                    }/*struct*/),
                    Rbrack: 23526,
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e5d80",&ast.Ident/*struct*/{
                    NamePos: 23528,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 23531,
                Ellipsis: 0,
                Rparen: 23532,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 23534,
      }/*struct*/),
    }/*struct*/),/* slice_item: 124*/125: p.StrmapFuncDecl("0xc4200e9050",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e8f60",&ast.FieldList/*struct*/{
        Opening: 23541,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dfb80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e5e00",&ast.Ident/*struct*/{
                NamePos: 23542,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3360",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dfb80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200e5e40",&ast.StarExpr/*struct*/{
              Star: 23544,
              X: p.StrmapIdent("0xc4200e5e20",&ast.Ident/*struct*/{
                NamePos: 23545,
                Name: "GoStmt",
                Obj: p.PtrmapObject("0xc4200c9720"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23551,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200e5e60",&ast.Ident/*struct*/{
        NamePos: 23553,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200e5fa0",&ast.FuncType/*struct*/{
        Func: 23536,
        Params: p.StrmapFieldList("0xc4200e8f90",&ast.FieldList/*struct*/{
          Opening: 23556,
          Closing: 23557,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e8fc0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dfbc0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200e5ec0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200e5e80",&ast.Ident/*struct*/{
                  NamePos: 23559,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200e5ea0",&ast.Ident/*struct*/{
                  NamePos: 23565,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e9020",&ast.BlockStmt/*struct*/{
        Lbrace: 23573,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200e5f80",&ast.ReturnStmt/*struct*/{
            Return: 23575,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200dfc00",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200e5f60",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200e5f20",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200e5ee0",&ast.Ident/*struct*/{
                      NamePos: 23582,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e3360"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200e5f00",&ast.Ident/*struct*/{
                      NamePos: 23584,
                      Name: "Call",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200e5f40",&ast.Ident/*struct*/{
                    NamePos: 23589,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 23592,
                Ellipsis: 0,
                Rparen: 23593,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 23595,
      }/*struct*/),
    }/*struct*/),/* slice_item: 125*/126: p.StrmapFuncDecl("0xc4200e91a0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e90b0",&ast.FieldList/*struct*/{
        Opening: 23602,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dfc40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200e5fc0",&ast.Ident/*struct*/{
                NamePos: 23603,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e33b0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dfc40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ec000",&ast.StarExpr/*struct*/{
              Star: 23605,
              X: p.StrmapIdent("0xc4200e5fe0",&ast.Ident/*struct*/{
                NamePos: 23606,
                Name: "DeferStmt",
                Obj: p.PtrmapObject("0xc4200c9810"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23615,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ec020",&ast.Ident/*struct*/{
        NamePos: 23617,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ec160",&ast.FuncType/*struct*/{
        Func: 23597,
        Params: p.StrmapFieldList("0xc4200e90e0",&ast.FieldList/*struct*/{
          Opening: 23620,
          Closing: 23621,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e9110",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dfc80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ec080",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ec040",&ast.Ident/*struct*/{
                  NamePos: 23623,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ec060",&ast.Ident/*struct*/{
                  NamePos: 23629,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e9170",&ast.BlockStmt/*struct*/{
        Lbrace: 23634,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ec140",&ast.ReturnStmt/*struct*/{
            Return: 23636,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200dfcc0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ec120",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ec0e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ec0a0",&ast.Ident/*struct*/{
                      NamePos: 23643,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e33b0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ec0c0",&ast.Ident/*struct*/{
                      NamePos: 23645,
                      Name: "Call",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ec100",&ast.Ident/*struct*/{
                    NamePos: 23650,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 23653,
                Ellipsis: 0,
                Rparen: 23654,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 23656,
      }/*struct*/),
    }/*struct*/),/* slice_item: 126*/127: p.StrmapFuncDecl("0xc4200e9440",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e9200",&ast.FieldList/*struct*/{
        Opening: 23663,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dfd00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ec180",&ast.Ident/*struct*/{
                NamePos: 23664,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3400",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dfd00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ec1c0",&ast.StarExpr/*struct*/{
              Star: 23666,
              X: p.StrmapIdent("0xc4200ec1a0",&ast.Ident/*struct*/{
                NamePos: 23667,
                Name: "ReturnStmt",
                Obj: p.PtrmapObject("0xc4200c9900"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23677,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ec1e0",&ast.Ident/*struct*/{
        NamePos: 23679,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ec500",&ast.FuncType/*struct*/{
        Func: 23658,
        Params: p.StrmapFieldList("0xc4200e9230",&ast.FieldList/*struct*/{
          Opening: 23682,
          Closing: 23683,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e9260",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dfd40",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ec240",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ec200",&ast.Ident/*struct*/{
                  NamePos: 23685,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ec220",&ast.Ident/*struct*/{
                  NamePos: 23691,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e9410",&ast.BlockStmt/*struct*/{
        Lbrace: 23695,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200dfe40",&ast.IfStmt/*struct*/{
            If: 23698,
            Init: p.StrmapAssignStmt("0xc4200dfdc0",&ast.AssignStmt/*struct*/{
              Lhs: []ast.Expr /*Slice*/{
                0: p.StrmapIdent("0xc4200ec260",&ast.Ident/*struct*/{
                  NamePos: 23701,
                  Name: "n",
                  Obj: p.StrmapObject("0xc4200e3450",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "n",
                    Decl: p.PtrmapAssignStmt("0xc4200dfdc0"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              TokPos: 23703,
              Tok: token.Token(47)/*:=*/,
              Rhs: []ast.Expr /*Slice*/{
                0: p.StrmapCallExpr("0xc4200dfd80",&ast.CallExpr/*struct*/{
                  Fun: p.StrmapIdent("0xc4200ec280",&ast.Ident/*struct*/{
                    NamePos: 23706,
                    Name: "len",
                  }/*struct*/),
                  Lparen: 23709,
                  Args: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200ec2e0",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200ec2a0",&ast.Ident/*struct*/{
                        NamePos: 23710,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200e3400"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200ec2c0",&ast.Ident/*struct*/{
                        NamePos: 23712,
                        Name: "Results",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Ellipsis: 0,
                  Rparen: 23719,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
            }/*struct*/),
            Cond: p.StrmapBinaryExpr("0xc4200e92f0",&ast.BinaryExpr/*struct*/{
              X: p.StrmapIdent("0xc4200ec300",&ast.Ident/*struct*/{
                NamePos: 23722,
                Name: "n",
                Obj: p.PtrmapObject("0xc4200e3450"),
              }/*struct*/),
              OpPos: 23724,
              Op: token.Token(41)/*>*/,
              Y: p.StrmapBasicLit("0xc4200ec320",&ast.BasicLit/*struct*/{
                ValuePos: 23726,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200e93b0",&ast.BlockStmt/*struct*/{
              Lbrace: 23728,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200ec420",&ast.ReturnStmt/*struct*/{
                  Return: 23732,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200dfe00",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200ec400",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIndexExpr("0xc4200e9380",&ast.IndexExpr/*struct*/{
                          X: p.StrmapSelectorExpr("0xc4200ec380",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200ec340",&ast.Ident/*struct*/{
                              NamePos: 23739,
                              Name: "s",
                              Obj: p.PtrmapObject("0xc4200e3400"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200ec360",&ast.Ident/*struct*/{
                              NamePos: 23741,
                              Name: "Results",
                            }/*struct*/),
                          }/*struct*/),
                          Lbrack: 23748,
                          Index: p.StrmapBinaryExpr("0xc4200e9350",&ast.BinaryExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200ec3a0",&ast.Ident/*struct*/{
                              NamePos: 23749,
                              Name: "n",
                              Obj: p.PtrmapObject("0xc4200e3450"),
                            }/*struct*/),
                            OpPos: 23750,
                            Op: token.Token(13)/*-*/,
                            Y: p.StrmapBasicLit("0xc4200ec3c0",&ast.BasicLit/*struct*/{
                              ValuePos: 23751,
                              Kind: token.Token(5)/*INT*/,
                              Value: "1",
                            }/*struct*/),
                          }/*struct*/),
                          Rbrack: 23752,
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200ec3e0",&ast.Ident/*struct*/{
                          NamePos: 23754,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 23757,
                      Ellipsis: 0,
                      Rparen: 23758,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 23761,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200ec4c0",&ast.ReturnStmt/*struct*/{
            Return: 23764,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200e93e0",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200ec480",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200ec440",&ast.Ident/*struct*/{
                    NamePos: 23771,
                    Name: "s",
                    Obj: p.PtrmapObject("0xc4200e3400"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ec460",&ast.Ident/*struct*/{
                    NamePos: 23773,
                    Name: "Return",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 23780,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200ec4a0",&ast.BasicLit/*struct*/{
                  ValuePos: 23782,
                  Kind: token.Token(5)/*INT*/,
                  Value: "6",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 23801,
      }/*struct*/),
    }/*struct*/),/* slice_item: 127*/128: p.StrmapFuncDecl("0xc4200e9680",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e94a0",&ast.FieldList/*struct*/{
        Opening: 23808,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200dfe80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ec520",&ast.Ident/*struct*/{
                NamePos: 23809,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e34a0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200dfe80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ec560",&ast.StarExpr/*struct*/{
              Star: 23811,
              X: p.StrmapIdent("0xc4200ec540",&ast.Ident/*struct*/{
                NamePos: 23812,
                Name: "BranchStmt",
                Obj: p.PtrmapObject("0xc4200c99f0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23822,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ec580",&ast.Ident/*struct*/{
        NamePos: 23824,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ec920",&ast.FuncType/*struct*/{
        Func: 23803,
        Params: p.StrmapFieldList("0xc4200e94d0",&ast.FieldList/*struct*/{
          Opening: 23827,
          Closing: 23828,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e9500",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200dfec0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ec5e0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ec5a0",&ast.Ident/*struct*/{
                  NamePos: 23830,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ec5c0",&ast.Ident/*struct*/{
                  NamePos: 23836,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e9650",&ast.BlockStmt/*struct*/{
        Lbrace: 23840,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200dff40",&ast.IfStmt/*struct*/{
            If: 23843,
            Cond: p.StrmapBinaryExpr("0xc4200e9590",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200ec640",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ec600",&ast.Ident/*struct*/{
                  NamePos: 23846,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e34a0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ec620",&ast.Ident/*struct*/{
                  NamePos: 23848,
                  Name: "Label",
                }/*struct*/),
              }/*struct*/),
              OpPos: 23854,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200ec660",&ast.Ident/*struct*/{
                NamePos: 23857,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200e95f0",&ast.BlockStmt/*struct*/{
              Lbrace: 23861,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200ec720",&ast.ReturnStmt/*struct*/{
                  Return: 23865,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200dff00",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200ec700",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapSelectorExpr("0xc4200ec6c0",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200ec680",&ast.Ident/*struct*/{
                            NamePos: 23872,
                            Name: "s",
                            Obj: p.PtrmapObject("0xc4200e34a0"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200ec6a0",&ast.Ident/*struct*/{
                            NamePos: 23874,
                            Name: "Label",
                          }/*struct*/),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200ec6e0",&ast.Ident/*struct*/{
                          NamePos: 23880,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 23883,
                      Ellipsis: 0,
                      Rparen: 23884,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 23887,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200ec8e0",&ast.ReturnStmt/*struct*/{
            Return: 23890,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ee040",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ec780",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200ec740",&ast.Ident/*struct*/{
                    NamePos: 23897,
                    Name: "token",
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ec760",&ast.Ident/*struct*/{
                    NamePos: 23903,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 23906,
                Args: []ast.Expr /*Slice*/{
                  0: p.StrmapBinaryExpr("0xc4200e9620",&ast.BinaryExpr/*struct*/{
                    X: p.StrmapCallExpr("0xc4200dff80",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapIdent("0xc4200ec7a0",&ast.Ident/*struct*/{
                        NamePos: 23907,
                        Name: "int",
                      }/*struct*/),
                      Lparen: 23910,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapSelectorExpr("0xc4200ec800",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200ec7c0",&ast.Ident/*struct*/{
                            NamePos: 23911,
                            Name: "s",
                            Obj: p.PtrmapObject("0xc4200e34a0"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200ec7e0",&ast.Ident/*struct*/{
                            NamePos: 23913,
                            Name: "TokPos",
                          }/*struct*/),
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 23919,
                    }/*struct*/),
                    OpPos: 23921,
                    Op: token.Token(12)/*+*/,
                    Y: p.StrmapCallExpr("0xc4200ee000",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapIdent("0xc4200ec820",&ast.Ident/*struct*/{
                        NamePos: 23923,
                        Name: "len",
                      }/*struct*/),
                      Lparen: 23926,
                      Args: []ast.Expr /*Slice*/{
                        0: p.StrmapCallExpr("0xc4200dffc0",&ast.CallExpr/*struct*/{
                          Fun: p.StrmapSelectorExpr("0xc4200ec8c0",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapSelectorExpr("0xc4200ec880",&ast.SelectorExpr/*struct*/{
                              X: p.StrmapIdent("0xc4200ec840",&ast.Ident/*struct*/{
                                NamePos: 23927,
                                Name: "s",
                                Obj: p.PtrmapObject("0xc4200e34a0"),
                              }/*struct*/),
                              Sel: p.StrmapIdent("0xc4200ec860",&ast.Ident/*struct*/{
                                NamePos: 23929,
                                Name: "Tok",
                              }/*struct*/),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200ec8a0",&ast.Ident/*struct*/{
                              NamePos: 23933,
                              Name: "String",
                            }/*struct*/),
                          }/*struct*/),
                          Lparen: 23939,
                          Ellipsis: 0,
                          Rparen: 23940,
                        }/*struct*/),/* slice_item: 0*/}/*slice*/,
                      Ellipsis: 0,
                      Rparen: 23941,
                    }/*struct*/),
                  }/*struct*/),/* slice_item: 0*/}/*slice*/,
                Ellipsis: 0,
                Rparen: 23942,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 23944,
      }/*struct*/),
    }/*struct*/),/* slice_item: 128*/129: p.StrmapFuncDecl("0xc4200e9800",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e96e0",&ast.FieldList/*struct*/{
        Opening: 23951,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee080",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ec940",&ast.Ident/*struct*/{
                NamePos: 23952,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e34f0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ee080"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ec980",&ast.StarExpr/*struct*/{
              Star: 23954,
              X: p.StrmapIdent("0xc4200ec960",&ast.Ident/*struct*/{
                NamePos: 23955,
                Name: "BlockStmt",
                Obj: p.PtrmapObject("0xc4200c9b30"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 23964,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ec9a0",&ast.Ident/*struct*/{
        NamePos: 23966,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ecac0",&ast.FuncType/*struct*/{
        Func: 23946,
        Params: p.StrmapFieldList("0xc4200e9710",&ast.FieldList/*struct*/{
          Opening: 23969,
          Closing: 23970,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e9740",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ee0c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200eca00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ec9c0",&ast.Ident/*struct*/{
                  NamePos: 23972,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ec9e0",&ast.Ident/*struct*/{
                  NamePos: 23978,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e97d0",&ast.BlockStmt/*struct*/{
        Lbrace: 23982,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ecaa0",&ast.ReturnStmt/*struct*/{
            Return: 23984,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200e97a0",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200eca60",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200eca20",&ast.Ident/*struct*/{
                    NamePos: 23991,
                    Name: "s",
                    Obj: p.PtrmapObject("0xc4200e34f0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200eca40",&ast.Ident/*struct*/{
                    NamePos: 23993,
                    Name: "Rbrace",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 24000,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200eca80",&ast.BasicLit/*struct*/{
                  ValuePos: 24002,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 24004,
      }/*struct*/),
    }/*struct*/),/* slice_item: 129*/130: p.StrmapFuncDecl("0xc4200e9a10",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e9860",&ast.FieldList/*struct*/{
        Opening: 24011,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee100",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ecae0",&ast.Ident/*struct*/{
                NamePos: 24012,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3540",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ee100"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ecb20",&ast.StarExpr/*struct*/{
              Star: 24014,
              X: p.StrmapIdent("0xc4200ecb00",&ast.Ident/*struct*/{
                NamePos: 24015,
                Name: "IfStmt",
                Obj: p.PtrmapObject("0xc4200c9c70"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24021,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ecb40",&ast.Ident/*struct*/{
        NamePos: 24023,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ecde0",&ast.FuncType/*struct*/{
        Func: 24006,
        Params: p.StrmapFieldList("0xc4200e9890",&ast.FieldList/*struct*/{
          Opening: 24026,
          Closing: 24027,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e98c0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ee140",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ecba0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ecb60",&ast.Ident/*struct*/{
                  NamePos: 24029,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ecb80",&ast.Ident/*struct*/{
                  NamePos: 24035,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e99e0",&ast.BlockStmt/*struct*/{
        Lbrace: 24039,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200ee1c0",&ast.IfStmt/*struct*/{
            If: 24042,
            Cond: p.StrmapBinaryExpr("0xc4200e9950",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200ecc00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ecbc0",&ast.Ident/*struct*/{
                  NamePos: 24045,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200e3540"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ecbe0",&ast.Ident/*struct*/{
                  NamePos: 24047,
                  Name: "Else",
                }/*struct*/),
              }/*struct*/),
              OpPos: 24052,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200ecc20",&ast.Ident/*struct*/{
                NamePos: 24055,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200e99b0",&ast.BlockStmt/*struct*/{
              Lbrace: 24059,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200ecce0",&ast.ReturnStmt/*struct*/{
                  Return: 24063,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200ee180",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200eccc0",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapSelectorExpr("0xc4200ecc80",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200ecc40",&ast.Ident/*struct*/{
                            NamePos: 24070,
                            Name: "s",
                            Obj: p.PtrmapObject("0xc4200e3540"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200ecc60",&ast.Ident/*struct*/{
                            NamePos: 24072,
                            Name: "Else",
                          }/*struct*/),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200ecca0",&ast.Ident/*struct*/{
                          NamePos: 24077,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 24080,
                      Ellipsis: 0,
                      Rparen: 24081,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 24084,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200ecda0",&ast.ReturnStmt/*struct*/{
            Return: 24087,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ee200",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ecd80",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ecd40",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ecd00",&ast.Ident/*struct*/{
                      NamePos: 24094,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e3540"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ecd20",&ast.Ident/*struct*/{
                      NamePos: 24096,
                      Name: "Body",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ecd60",&ast.Ident/*struct*/{
                    NamePos: 24101,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 24104,
                Ellipsis: 0,
                Rparen: 24105,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 24107,
      }/*struct*/),
    }/*struct*/),/* slice_item: 130*/131: p.StrmapFuncDecl("0xc4200e9cb0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e9a70",&ast.FieldList/*struct*/{
        Opening: 24114,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee240",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ece00",&ast.Ident/*struct*/{
                NamePos: 24115,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3590",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ee240"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ece40",&ast.StarExpr/*struct*/{
              Star: 24117,
              X: p.StrmapIdent("0xc4200ece20",&ast.Ident/*struct*/{
                NamePos: 24118,
                Name: "CaseClause",
                Obj: p.PtrmapObject("0xc4200c9ea0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24128,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ece60",&ast.Ident/*struct*/{
        NamePos: 24130,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ed180",&ast.FuncType/*struct*/{
        Func: 24109,
        Params: p.StrmapFieldList("0xc4200e9aa0",&ast.FieldList/*struct*/{
          Opening: 24133,
          Closing: 24134,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e9ad0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ee280",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ecec0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ece80",&ast.Ident/*struct*/{
                  NamePos: 24136,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ecea0",&ast.Ident/*struct*/{
                  NamePos: 24142,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e9c80",&ast.BlockStmt/*struct*/{
        Lbrace: 24146,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200ee380",&ast.IfStmt/*struct*/{
            If: 24149,
            Init: p.StrmapAssignStmt("0xc4200ee300",&ast.AssignStmt/*struct*/{
              Lhs: []ast.Expr /*Slice*/{
                0: p.StrmapIdent("0xc4200ecee0",&ast.Ident/*struct*/{
                  NamePos: 24152,
                  Name: "n",
                  Obj: p.StrmapObject("0xc4200e35e0",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "n",
                    Decl: p.PtrmapAssignStmt("0xc4200ee300"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              TokPos: 24154,
              Tok: token.Token(47)/*:=*/,
              Rhs: []ast.Expr /*Slice*/{
                0: p.StrmapCallExpr("0xc4200ee2c0",&ast.CallExpr/*struct*/{
                  Fun: p.StrmapIdent("0xc4200ecf00",&ast.Ident/*struct*/{
                    NamePos: 24157,
                    Name: "len",
                  }/*struct*/),
                  Lparen: 24160,
                  Args: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200ecf60",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200ecf20",&ast.Ident/*struct*/{
                        NamePos: 24161,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200e3590"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200ecf40",&ast.Ident/*struct*/{
                        NamePos: 24163,
                        Name: "Body",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Ellipsis: 0,
                  Rparen: 24167,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
            }/*struct*/),
            Cond: p.StrmapBinaryExpr("0xc4200e9b60",&ast.BinaryExpr/*struct*/{
              X: p.StrmapIdent("0xc4200ecf80",&ast.Ident/*struct*/{
                NamePos: 24170,
                Name: "n",
                Obj: p.PtrmapObject("0xc4200e35e0"),
              }/*struct*/),
              OpPos: 24172,
              Op: token.Token(41)/*>*/,
              Y: p.StrmapBasicLit("0xc4200ecfa0",&ast.BasicLit/*struct*/{
                ValuePos: 24174,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200e9c20",&ast.BlockStmt/*struct*/{
              Lbrace: 24176,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200ed0a0",&ast.ReturnStmt/*struct*/{
                  Return: 24180,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200ee340",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200ed080",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIndexExpr("0xc4200e9bf0",&ast.IndexExpr/*struct*/{
                          X: p.StrmapSelectorExpr("0xc4200ed000",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200ecfc0",&ast.Ident/*struct*/{
                              NamePos: 24187,
                              Name: "s",
                              Obj: p.PtrmapObject("0xc4200e3590"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200ecfe0",&ast.Ident/*struct*/{
                              NamePos: 24189,
                              Name: "Body",
                            }/*struct*/),
                          }/*struct*/),
                          Lbrack: 24193,
                          Index: p.StrmapBinaryExpr("0xc4200e9bc0",&ast.BinaryExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200ed020",&ast.Ident/*struct*/{
                              NamePos: 24194,
                              Name: "n",
                              Obj: p.PtrmapObject("0xc4200e35e0"),
                            }/*struct*/),
                            OpPos: 24195,
                            Op: token.Token(13)/*-*/,
                            Y: p.StrmapBasicLit("0xc4200ed040",&ast.BasicLit/*struct*/{
                              ValuePos: 24196,
                              Kind: token.Token(5)/*INT*/,
                              Value: "1",
                            }/*struct*/),
                          }/*struct*/),
                          Rbrack: 24197,
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200ed060",&ast.Ident/*struct*/{
                          NamePos: 24199,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 24202,
                      Ellipsis: 0,
                      Rparen: 24203,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 24206,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200ed140",&ast.ReturnStmt/*struct*/{
            Return: 24209,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200e9c50",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200ed100",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200ed0c0",&ast.Ident/*struct*/{
                    NamePos: 24216,
                    Name: "s",
                    Obj: p.PtrmapObject("0xc4200e3590"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ed0e0",&ast.Ident/*struct*/{
                    NamePos: 24218,
                    Name: "Colon",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 24224,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200ed120",&ast.BasicLit/*struct*/{
                  ValuePos: 24226,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 24228,
      }/*struct*/),
    }/*struct*/),/* slice_item: 131*/132: p.StrmapFuncDecl("0xc4200e9e00",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e9d10",&ast.FieldList/*struct*/{
        Opening: 24235,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee3c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ed1a0",&ast.Ident/*struct*/{
                NamePos: 24236,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3630",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ee3c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ed1e0",&ast.StarExpr/*struct*/{
              Star: 24238,
              X: p.StrmapIdent("0xc4200ed1c0",&ast.Ident/*struct*/{
                NamePos: 24239,
                Name: "SwitchStmt",
                Obj: p.PtrmapObject("0xc4200e2050"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24249,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ed200",&ast.Ident/*struct*/{
        NamePos: 24251,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ed340",&ast.FuncType/*struct*/{
        Func: 24230,
        Params: p.StrmapFieldList("0xc4200e9d40",&ast.FieldList/*struct*/{
          Opening: 24254,
          Closing: 24255,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e9d70",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ee400",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ed260",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ed220",&ast.Ident/*struct*/{
                  NamePos: 24257,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ed240",&ast.Ident/*struct*/{
                  NamePos: 24263,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e9dd0",&ast.BlockStmt/*struct*/{
        Lbrace: 24271,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ed320",&ast.ReturnStmt/*struct*/{
            Return: 24273,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ee440",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ed300",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ed2c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ed280",&ast.Ident/*struct*/{
                      NamePos: 24280,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e3630"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ed2a0",&ast.Ident/*struct*/{
                      NamePos: 24282,
                      Name: "Body",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ed2e0",&ast.Ident/*struct*/{
                    NamePos: 24287,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 24290,
                Ellipsis: 0,
                Rparen: 24291,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 24293,
      }/*struct*/),
    }/*struct*/),/* slice_item: 132*/133: p.StrmapFuncDecl("0xc4200e9f50",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e9e60",&ast.FieldList/*struct*/{
        Opening: 24300,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee480",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ed360",&ast.Ident/*struct*/{
                NamePos: 24301,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3680",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ee480"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ed3a0",&ast.StarExpr/*struct*/{
              Star: 24303,
              X: p.StrmapIdent("0xc4200ed380",&ast.Ident/*struct*/{
                NamePos: 24304,
                Name: "TypeSwitchStmt",
                Obj: p.PtrmapObject("0xc4200e21e0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24318,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ed3c0",&ast.Ident/*struct*/{
        NamePos: 24320,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ed500",&ast.FuncType/*struct*/{
        Func: 24295,
        Params: p.StrmapFieldList("0xc4200e9e90",&ast.FieldList/*struct*/{
          Opening: 24323,
          Closing: 24324,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200e9ec0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ee4c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ed420",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ed3e0",&ast.Ident/*struct*/{
                  NamePos: 24326,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ed400",&ast.Ident/*struct*/{
                  NamePos: 24332,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200e9f20",&ast.BlockStmt/*struct*/{
        Lbrace: 24336,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200ed4e0",&ast.ReturnStmt/*struct*/{
            Return: 24338,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ee500",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200ed4c0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ed480",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ed440",&ast.Ident/*struct*/{
                      NamePos: 24345,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e3680"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ed460",&ast.Ident/*struct*/{
                      NamePos: 24347,
                      Name: "Body",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ed4a0",&ast.Ident/*struct*/{
                    NamePos: 24352,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 24355,
                Ellipsis: 0,
                Rparen: 24356,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 24358,
      }/*struct*/),
    }/*struct*/),/* slice_item: 133*/134: p.StrmapFuncDecl("0xc4200f0210",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200e9fb0",&ast.FieldList/*struct*/{
        Opening: 24365,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee540",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ed520",&ast.Ident/*struct*/{
                NamePos: 24366,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e36d0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ee540"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ed560",&ast.StarExpr/*struct*/{
              Star: 24368,
              X: p.StrmapIdent("0xc4200ed540",&ast.Ident/*struct*/{
                NamePos: 24369,
                Name: "CommClause",
                Obj: p.PtrmapObject("0xc4200e2370"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24379,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ed580",&ast.Ident/*struct*/{
        NamePos: 24381,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ed8a0",&ast.FuncType/*struct*/{
        Func: 24360,
        Params: p.StrmapFieldList("0xc4200f0000",&ast.FieldList/*struct*/{
          Opening: 24384,
          Closing: 24385,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f0030",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ee580",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ed5e0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ed5a0",&ast.Ident/*struct*/{
                  NamePos: 24387,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ed5c0",&ast.Ident/*struct*/{
                  NamePos: 24393,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f01e0",&ast.BlockStmt/*struct*/{
        Lbrace: 24397,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200ee680",&ast.IfStmt/*struct*/{
            If: 24400,
            Init: p.StrmapAssignStmt("0xc4200ee600",&ast.AssignStmt/*struct*/{
              Lhs: []ast.Expr /*Slice*/{
                0: p.StrmapIdent("0xc4200ed600",&ast.Ident/*struct*/{
                  NamePos: 24403,
                  Name: "n",
                  Obj: p.StrmapObject("0xc4200e3720",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "n",
                    Decl: p.PtrmapAssignStmt("0xc4200ee600"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              TokPos: 24405,
              Tok: token.Token(47)/*:=*/,
              Rhs: []ast.Expr /*Slice*/{
                0: p.StrmapCallExpr("0xc4200ee5c0",&ast.CallExpr/*struct*/{
                  Fun: p.StrmapIdent("0xc4200ed620",&ast.Ident/*struct*/{
                    NamePos: 24408,
                    Name: "len",
                  }/*struct*/),
                  Lparen: 24411,
                  Args: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200ed680",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200ed640",&ast.Ident/*struct*/{
                        NamePos: 24412,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200e36d0"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200ed660",&ast.Ident/*struct*/{
                        NamePos: 24414,
                        Name: "Body",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Ellipsis: 0,
                  Rparen: 24418,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
            }/*struct*/),
            Cond: p.StrmapBinaryExpr("0xc4200f00c0",&ast.BinaryExpr/*struct*/{
              X: p.StrmapIdent("0xc4200ed6a0",&ast.Ident/*struct*/{
                NamePos: 24421,
                Name: "n",
                Obj: p.PtrmapObject("0xc4200e3720"),
              }/*struct*/),
              OpPos: 24423,
              Op: token.Token(41)/*>*/,
              Y: p.StrmapBasicLit("0xc4200ed6c0",&ast.BasicLit/*struct*/{
                ValuePos: 24425,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200f0180",&ast.BlockStmt/*struct*/{
              Lbrace: 24427,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200ed7c0",&ast.ReturnStmt/*struct*/{
                  Return: 24431,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200ee640",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200ed7a0",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIndexExpr("0xc4200f0150",&ast.IndexExpr/*struct*/{
                          X: p.StrmapSelectorExpr("0xc4200ed720",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200ed6e0",&ast.Ident/*struct*/{
                              NamePos: 24438,
                              Name: "s",
                              Obj: p.PtrmapObject("0xc4200e36d0"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200ed700",&ast.Ident/*struct*/{
                              NamePos: 24440,
                              Name: "Body",
                            }/*struct*/),
                          }/*struct*/),
                          Lbrack: 24444,
                          Index: p.StrmapBinaryExpr("0xc4200f0120",&ast.BinaryExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200ed740",&ast.Ident/*struct*/{
                              NamePos: 24445,
                              Name: "n",
                              Obj: p.PtrmapObject("0xc4200e3720"),
                            }/*struct*/),
                            OpPos: 24446,
                            Op: token.Token(13)/*-*/,
                            Y: p.StrmapBasicLit("0xc4200ed760",&ast.BasicLit/*struct*/{
                              ValuePos: 24447,
                              Kind: token.Token(5)/*INT*/,
                              Value: "1",
                            }/*struct*/),
                          }/*struct*/),
                          Rbrack: 24448,
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200ed780",&ast.Ident/*struct*/{
                          NamePos: 24450,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 24453,
                      Ellipsis: 0,
                      Rparen: 24454,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 24457,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200ed860",&ast.ReturnStmt/*struct*/{
            Return: 24460,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapBinaryExpr("0xc4200f01b0",&ast.BinaryExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200ed820",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200ed7e0",&ast.Ident/*struct*/{
                    NamePos: 24467,
                    Name: "s",
                    Obj: p.PtrmapObject("0xc4200e36d0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200ed800",&ast.Ident/*struct*/{
                    NamePos: 24469,
                    Name: "Colon",
                  }/*struct*/),
                }/*struct*/),
                OpPos: 24475,
                Op: token.Token(12)/*+*/,
                Y: p.StrmapBasicLit("0xc4200ed840",&ast.BasicLit/*struct*/{
                  ValuePos: 24477,
                  Kind: token.Token(5)/*INT*/,
                  Value: "1",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 24479,
      }/*struct*/),
    }/*struct*/),/* slice_item: 134*/135: p.StrmapFuncDecl("0xc4200f0360",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f0270",&ast.FieldList/*struct*/{
        Opening: 24486,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee6c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200ed8c0",&ast.Ident/*struct*/{
                NamePos: 24487,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3770",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ee6c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200ed900",&ast.StarExpr/*struct*/{
              Star: 24489,
              X: p.StrmapIdent("0xc4200ed8e0",&ast.Ident/*struct*/{
                NamePos: 24490,
                Name: "SelectStmt",
                Obj: p.PtrmapObject("0xc4200e2500"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24500,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ed920",&ast.Ident/*struct*/{
        NamePos: 24502,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200eda60",&ast.FuncType/*struct*/{
        Func: 24481,
        Params: p.StrmapFieldList("0xc4200f02a0",&ast.FieldList/*struct*/{
          Opening: 24505,
          Closing: 24506,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f02d0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ee700",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200ed980",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200ed940",&ast.Ident/*struct*/{
                  NamePos: 24508,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200ed960",&ast.Ident/*struct*/{
                  NamePos: 24514,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0330",&ast.BlockStmt/*struct*/{
        Lbrace: 24518,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200eda40",&ast.ReturnStmt/*struct*/{
            Return: 24520,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ee740",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200eda20",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200ed9e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200ed9a0",&ast.Ident/*struct*/{
                      NamePos: 24527,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e3770"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200ed9c0",&ast.Ident/*struct*/{
                      NamePos: 24529,
                      Name: "Body",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200eda00",&ast.Ident/*struct*/{
                    NamePos: 24534,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 24537,
                Ellipsis: 0,
                Rparen: 24538,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 24540,
      }/*struct*/),
    }/*struct*/),/* slice_item: 135*/136: p.StrmapFuncDecl("0xc4200f04b0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f03c0",&ast.FieldList/*struct*/{
        Opening: 24547,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee780",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200eda80",&ast.Ident/*struct*/{
                NamePos: 24548,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e37c0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ee780"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200edac0",&ast.StarExpr/*struct*/{
              Star: 24550,
              X: p.StrmapIdent("0xc4200edaa0",&ast.Ident/*struct*/{
                NamePos: 24551,
                Name: "ForStmt",
                Obj: p.PtrmapObject("0xc4200e25f0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24558,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200edae0",&ast.Ident/*struct*/{
        NamePos: 24560,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200edc20",&ast.FuncType/*struct*/{
        Func: 24542,
        Params: p.StrmapFieldList("0xc4200f03f0",&ast.FieldList/*struct*/{
          Opening: 24563,
          Closing: 24564,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f0420",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ee7c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200edb40",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200edb00",&ast.Ident/*struct*/{
                  NamePos: 24566,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200edb20",&ast.Ident/*struct*/{
                  NamePos: 24572,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0480",&ast.BlockStmt/*struct*/{
        Lbrace: 24579,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200edc00",&ast.ReturnStmt/*struct*/{
            Return: 24581,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ee800",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200edbe0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200edba0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200edb60",&ast.Ident/*struct*/{
                      NamePos: 24588,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e37c0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200edb80",&ast.Ident/*struct*/{
                      NamePos: 24590,
                      Name: "Body",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200edbc0",&ast.Ident/*struct*/{
                    NamePos: 24595,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 24598,
                Ellipsis: 0,
                Rparen: 24599,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 24601,
      }/*struct*/),
    }/*struct*/),/* slice_item: 136*/137: p.StrmapFuncDecl("0xc4200f0600",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f0510",&ast.FieldList/*struct*/{
        Opening: 24608,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee840",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200edc40",&ast.Ident/*struct*/{
                NamePos: 24609,
                Name: "s",
                Obj: p.StrmapObject("0xc4200e3810",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ee840"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200edc80",&ast.StarExpr/*struct*/{
              Star: 24611,
              X: p.StrmapIdent("0xc4200edc60",&ast.Ident/*struct*/{
                NamePos: 24612,
                Name: "RangeStmt",
                Obj: p.PtrmapObject("0xc4200e27d0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24621,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200edca0",&ast.Ident/*struct*/{
        NamePos: 24623,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ede00",&ast.FuncType/*struct*/{
        Func: 24603,
        Params: p.StrmapFieldList("0xc4200f0540",&ast.FieldList/*struct*/{
          Opening: 24626,
          Closing: 24627,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f0570",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ee880",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200edd00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200edcc0",&ast.Ident/*struct*/{
                  NamePos: 24629,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200edce0",&ast.Ident/*struct*/{
                  NamePos: 24635,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f05d0",&ast.BlockStmt/*struct*/{
        Lbrace: 24640,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200eddc0",&ast.ReturnStmt/*struct*/{
            Return: 24642,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ee8c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200edda0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200edd60",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200edd20",&ast.Ident/*struct*/{
                      NamePos: 24649,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200e3810"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200edd40",&ast.Ident/*struct*/{
                      NamePos: 24651,
                      Name: "Body",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200edd80",&ast.Ident/*struct*/{
                    NamePos: 24656,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 24659,
                Ellipsis: 0,
                Rparen: 24660,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 24662,
      }/*struct*/),
    }/*struct*/),/* slice_item: 137*/138: p.StrmapFuncDecl("0xc4200f0720",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f0660",&ast.FieldList/*struct*/{
        Opening: 24751,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee940",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200ede40",&ast.StarExpr/*struct*/{
              Star: 24752,
              X: p.StrmapIdent("0xc4200ede20",&ast.Ident/*struct*/{
                NamePos: 24753,
                Name: "BadStmt",
                Obj: p.PtrmapObject("0xc4200c8e10"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24760,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200ede60",&ast.Ident/*struct*/{
        NamePos: 24762,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200ede80",&ast.FuncType/*struct*/{
        Func: 24746,
        Params: p.StrmapFieldList("0xc4200f0690",&ast.FieldList/*struct*/{
          Opening: 24770,
          Closing: 24771,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f06f0",&ast.BlockStmt/*struct*/{
        Lbrace: 24780,
        Rbrace: 24781,
      }/*struct*/),
    }/*struct*/),/* slice_item: 138*/139: p.StrmapFuncDecl("0xc4200f0840",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f0780",&ast.FieldList/*struct*/{
        Opening: 24788,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee980",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200edec0",&ast.StarExpr/*struct*/{
              Star: 24789,
              X: p.StrmapIdent("0xc4200edea0",&ast.Ident/*struct*/{
                NamePos: 24790,
                Name: "DeclStmt",
                Obj: p.PtrmapObject("0xc4200c8f00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24798,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200edee0",&ast.Ident/*struct*/{
        NamePos: 24800,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200edf00",&ast.FuncType/*struct*/{
        Func: 24783,
        Params: p.StrmapFieldList("0xc4200f07b0",&ast.FieldList/*struct*/{
          Opening: 24808,
          Closing: 24809,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0810",&ast.BlockStmt/*struct*/{
        Lbrace: 24817,
        Rbrace: 24818,
      }/*struct*/),
    }/*struct*/),/* slice_item: 139*/140: p.StrmapFuncDecl("0xc4200f0960",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f08a0",&ast.FieldList/*struct*/{
        Opening: 24825,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ee9c0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200edf40",&ast.StarExpr/*struct*/{
              Star: 24826,
              X: p.StrmapIdent("0xc4200edf20",&ast.Ident/*struct*/{
                NamePos: 24827,
                Name: "EmptyStmt",
                Obj: p.PtrmapObject("0xc4200c8ff0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24836,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200edf60",&ast.Ident/*struct*/{
        NamePos: 24838,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200edf80",&ast.FuncType/*struct*/{
        Func: 24820,
        Params: p.StrmapFieldList("0xc4200f08d0",&ast.FieldList/*struct*/{
          Opening: 24846,
          Closing: 24847,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0930",&ast.BlockStmt/*struct*/{
        Lbrace: 24854,
        Rbrace: 24855,
      }/*struct*/),
    }/*struct*/),/* slice_item: 140*/141: p.StrmapFuncDecl("0xc4200f0a80",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f09c0",&ast.FieldList/*struct*/{
        Opening: 24862,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eea00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200edfc0",&ast.StarExpr/*struct*/{
              Star: 24863,
              X: p.StrmapIdent("0xc4200edfa0",&ast.Ident/*struct*/{
                NamePos: 24864,
                Name: "LabeledStmt",
                Obj: p.PtrmapObject("0xc4200c90e0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24875,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200edfe0",&ast.Ident/*struct*/{
        NamePos: 24877,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2000",&ast.FuncType/*struct*/{
        Func: 24857,
        Params: p.StrmapFieldList("0xc4200f09f0",&ast.FieldList/*struct*/{
          Opening: 24885,
          Closing: 24886,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0a50",&ast.BlockStmt/*struct*/{
        Lbrace: 24891,
        Rbrace: 24892,
      }/*struct*/),
    }/*struct*/),/* slice_item: 141*/142: p.StrmapFuncDecl("0xc4200f0ba0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f0ae0",&ast.FieldList/*struct*/{
        Opening: 24899,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eea40",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f2040",&ast.StarExpr/*struct*/{
              Star: 24900,
              X: p.StrmapIdent("0xc4200f2020",&ast.Ident/*struct*/{
                NamePos: 24901,
                Name: "ExprStmt",
                Obj: p.PtrmapObject("0xc4200c9220"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24909,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f2060",&ast.Ident/*struct*/{
        NamePos: 24911,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2080",&ast.FuncType/*struct*/{
        Func: 24894,
        Params: p.StrmapFieldList("0xc4200f0b10",&ast.FieldList/*struct*/{
          Opening: 24919,
          Closing: 24920,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0b70",&ast.BlockStmt/*struct*/{
        Lbrace: 24928,
        Rbrace: 24929,
      }/*struct*/),
    }/*struct*/),/* slice_item: 142*/143: p.StrmapFuncDecl("0xc4200f0cc0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f0c00",&ast.FieldList/*struct*/{
        Opening: 24936,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eea80",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f20c0",&ast.StarExpr/*struct*/{
              Star: 24937,
              X: p.StrmapIdent("0xc4200f20a0",&ast.Ident/*struct*/{
                NamePos: 24938,
                Name: "SendStmt",
                Obj: p.PtrmapObject("0xc4200c92c0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24946,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f20e0",&ast.Ident/*struct*/{
        NamePos: 24948,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2100",&ast.FuncType/*struct*/{
        Func: 24931,
        Params: p.StrmapFieldList("0xc4200f0c30",&ast.FieldList/*struct*/{
          Opening: 24956,
          Closing: 24957,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0c90",&ast.BlockStmt/*struct*/{
        Lbrace: 24965,
        Rbrace: 24966,
      }/*struct*/),
    }/*struct*/),/* slice_item: 143*/144: p.StrmapFuncDecl("0xc4200f0de0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f0d20",&ast.FieldList/*struct*/{
        Opening: 24973,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eeac0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f2140",&ast.StarExpr/*struct*/{
              Star: 24974,
              X: p.StrmapIdent("0xc4200f2120",&ast.Ident/*struct*/{
                NamePos: 24975,
                Name: "IncDecStmt",
                Obj: p.PtrmapObject("0xc4200c9450"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 24985,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f2160",&ast.Ident/*struct*/{
        NamePos: 24987,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2180",&ast.FuncType/*struct*/{
        Func: 24968,
        Params: p.StrmapFieldList("0xc4200f0d50",&ast.FieldList/*struct*/{
          Opening: 24995,
          Closing: 24996,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0db0",&ast.BlockStmt/*struct*/{
        Lbrace: 25002,
        Rbrace: 25003,
      }/*struct*/),
    }/*struct*/),/* slice_item: 144*/145: p.StrmapFuncDecl("0xc4200f0f00",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f0e40",&ast.FieldList/*struct*/{
        Opening: 25010,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eeb00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f21c0",&ast.StarExpr/*struct*/{
              Star: 25011,
              X: p.StrmapIdent("0xc4200f21a0",&ast.Ident/*struct*/{
                NamePos: 25012,
                Name: "AssignStmt",
                Obj: p.PtrmapObject("0xc4200c9590"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25022,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f21e0",&ast.Ident/*struct*/{
        NamePos: 25024,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2200",&ast.FuncType/*struct*/{
        Func: 25005,
        Params: p.StrmapFieldList("0xc4200f0e70",&ast.FieldList/*struct*/{
          Opening: 25032,
          Closing: 25033,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0ed0",&ast.BlockStmt/*struct*/{
        Lbrace: 25039,
        Rbrace: 25040,
      }/*struct*/),
    }/*struct*/),/* slice_item: 145*/146: p.StrmapFuncDecl("0xc4200f1020",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f0f60",&ast.FieldList/*struct*/{
        Opening: 25047,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eeb40",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f2240",&ast.StarExpr/*struct*/{
              Star: 25048,
              X: p.StrmapIdent("0xc4200f2220",&ast.Ident/*struct*/{
                NamePos: 25049,
                Name: "GoStmt",
                Obj: p.PtrmapObject("0xc4200c9720"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25055,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f2260",&ast.Ident/*struct*/{
        NamePos: 25057,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2280",&ast.FuncType/*struct*/{
        Func: 25042,
        Params: p.StrmapFieldList("0xc4200f0f90",&ast.FieldList/*struct*/{
          Opening: 25065,
          Closing: 25066,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f0ff0",&ast.BlockStmt/*struct*/{
        Lbrace: 25076,
        Rbrace: 25077,
      }/*struct*/),
    }/*struct*/),/* slice_item: 146*/147: p.StrmapFuncDecl("0xc4200f1140",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f1080",&ast.FieldList/*struct*/{
        Opening: 25084,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eeb80",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f22c0",&ast.StarExpr/*struct*/{
              Star: 25085,
              X: p.StrmapIdent("0xc4200f22a0",&ast.Ident/*struct*/{
                NamePos: 25086,
                Name: "DeferStmt",
                Obj: p.PtrmapObject("0xc4200c9810"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25095,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f22e0",&ast.Ident/*struct*/{
        NamePos: 25097,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2300",&ast.FuncType/*struct*/{
        Func: 25079,
        Params: p.StrmapFieldList("0xc4200f10b0",&ast.FieldList/*struct*/{
          Opening: 25105,
          Closing: 25106,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f1110",&ast.BlockStmt/*struct*/{
        Lbrace: 25113,
        Rbrace: 25114,
      }/*struct*/),
    }/*struct*/),/* slice_item: 147*/148: p.StrmapFuncDecl("0xc4200f1260",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f11a0",&ast.FieldList/*struct*/{
        Opening: 25121,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eebc0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f2340",&ast.StarExpr/*struct*/{
              Star: 25122,
              X: p.StrmapIdent("0xc4200f2320",&ast.Ident/*struct*/{
                NamePos: 25123,
                Name: "ReturnStmt",
                Obj: p.PtrmapObject("0xc4200c9900"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25133,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f2360",&ast.Ident/*struct*/{
        NamePos: 25135,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2380",&ast.FuncType/*struct*/{
        Func: 25116,
        Params: p.StrmapFieldList("0xc4200f11d0",&ast.FieldList/*struct*/{
          Opening: 25143,
          Closing: 25144,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f1230",&ast.BlockStmt/*struct*/{
        Lbrace: 25150,
        Rbrace: 25151,
      }/*struct*/),
    }/*struct*/),/* slice_item: 148*/149: p.StrmapFuncDecl("0xc4200f1380",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f12c0",&ast.FieldList/*struct*/{
        Opening: 25158,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eec00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f23c0",&ast.StarExpr/*struct*/{
              Star: 25159,
              X: p.StrmapIdent("0xc4200f23a0",&ast.Ident/*struct*/{
                NamePos: 25160,
                Name: "BranchStmt",
                Obj: p.PtrmapObject("0xc4200c99f0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25170,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f23e0",&ast.Ident/*struct*/{
        NamePos: 25172,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2400",&ast.FuncType/*struct*/{
        Func: 25153,
        Params: p.StrmapFieldList("0xc4200f12f0",&ast.FieldList/*struct*/{
          Opening: 25180,
          Closing: 25181,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f1350",&ast.BlockStmt/*struct*/{
        Lbrace: 25187,
        Rbrace: 25188,
      }/*struct*/),
    }/*struct*/),/* slice_item: 149*/150: p.StrmapFuncDecl("0xc4200f14a0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f13e0",&ast.FieldList/*struct*/{
        Opening: 25195,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eec40",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f2440",&ast.StarExpr/*struct*/{
              Star: 25196,
              X: p.StrmapIdent("0xc4200f2420",&ast.Ident/*struct*/{
                NamePos: 25197,
                Name: "BlockStmt",
                Obj: p.PtrmapObject("0xc4200c9b30"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25206,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f2460",&ast.Ident/*struct*/{
        NamePos: 25208,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2480",&ast.FuncType/*struct*/{
        Func: 25190,
        Params: p.StrmapFieldList("0xc4200f1410",&ast.FieldList/*struct*/{
          Opening: 25216,
          Closing: 25217,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f1470",&ast.BlockStmt/*struct*/{
        Lbrace: 25224,
        Rbrace: 25225,
      }/*struct*/),
    }/*struct*/),/* slice_item: 150*/151: p.StrmapFuncDecl("0xc4200f15c0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f1500",&ast.FieldList/*struct*/{
        Opening: 25232,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eec80",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f24c0",&ast.StarExpr/*struct*/{
              Star: 25233,
              X: p.StrmapIdent("0xc4200f24a0",&ast.Ident/*struct*/{
                NamePos: 25234,
                Name: "IfStmt",
                Obj: p.PtrmapObject("0xc4200c9c70"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25240,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f24e0",&ast.Ident/*struct*/{
        NamePos: 25242,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2500",&ast.FuncType/*struct*/{
        Func: 25227,
        Params: p.StrmapFieldList("0xc4200f1530",&ast.FieldList/*struct*/{
          Opening: 25250,
          Closing: 25251,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f1590",&ast.BlockStmt/*struct*/{
        Lbrace: 25261,
        Rbrace: 25262,
      }/*struct*/),
    }/*struct*/),/* slice_item: 151*/152: p.StrmapFuncDecl("0xc4200f16e0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f1620",&ast.FieldList/*struct*/{
        Opening: 25269,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eecc0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f2540",&ast.StarExpr/*struct*/{
              Star: 25270,
              X: p.StrmapIdent("0xc4200f2520",&ast.Ident/*struct*/{
                NamePos: 25271,
                Name: "CaseClause",
                Obj: p.PtrmapObject("0xc4200c9ea0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25281,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f2560",&ast.Ident/*struct*/{
        NamePos: 25283,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2580",&ast.FuncType/*struct*/{
        Func: 25264,
        Params: p.StrmapFieldList("0xc4200f1650",&ast.FieldList/*struct*/{
          Opening: 25291,
          Closing: 25292,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f16b0",&ast.BlockStmt/*struct*/{
        Lbrace: 25298,
        Rbrace: 25299,
      }/*struct*/),
    }/*struct*/),/* slice_item: 152*/153: p.StrmapFuncDecl("0xc4200f1800",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f1740",&ast.FieldList/*struct*/{
        Opening: 25306,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eed00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f25c0",&ast.StarExpr/*struct*/{
              Star: 25307,
              X: p.StrmapIdent("0xc4200f25a0",&ast.Ident/*struct*/{
                NamePos: 25308,
                Name: "SwitchStmt",
                Obj: p.PtrmapObject("0xc4200e2050"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25318,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f25e0",&ast.Ident/*struct*/{
        NamePos: 25320,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2600",&ast.FuncType/*struct*/{
        Func: 25301,
        Params: p.StrmapFieldList("0xc4200f1770",&ast.FieldList/*struct*/{
          Opening: 25328,
          Closing: 25329,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f17d0",&ast.BlockStmt/*struct*/{
        Lbrace: 25335,
        Rbrace: 25336,
      }/*struct*/),
    }/*struct*/),/* slice_item: 153*/154: p.StrmapFuncDecl("0xc4200f1920",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f1860",&ast.FieldList/*struct*/{
        Opening: 25343,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eed40",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f2640",&ast.StarExpr/*struct*/{
              Star: 25344,
              X: p.StrmapIdent("0xc4200f2620",&ast.Ident/*struct*/{
                NamePos: 25345,
                Name: "TypeSwitchStmt",
                Obj: p.PtrmapObject("0xc4200e21e0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25359,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f2660",&ast.Ident/*struct*/{
        NamePos: 25361,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2680",&ast.FuncType/*struct*/{
        Func: 25338,
        Params: p.StrmapFieldList("0xc4200f1890",&ast.FieldList/*struct*/{
          Opening: 25369,
          Closing: 25370,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f18f0",&ast.BlockStmt/*struct*/{
        Lbrace: 25372,
        Rbrace: 25373,
      }/*struct*/),
    }/*struct*/),/* slice_item: 154*/155: p.StrmapFuncDecl("0xc4200f1a40",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f1980",&ast.FieldList/*struct*/{
        Opening: 25380,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eed80",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f26c0",&ast.StarExpr/*struct*/{
              Star: 25381,
              X: p.StrmapIdent("0xc4200f26a0",&ast.Ident/*struct*/{
                NamePos: 25382,
                Name: "CommClause",
                Obj: p.PtrmapObject("0xc4200e2370"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25392,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f26e0",&ast.Ident/*struct*/{
        NamePos: 25394,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2700",&ast.FuncType/*struct*/{
        Func: 25375,
        Params: p.StrmapFieldList("0xc4200f19b0",&ast.FieldList/*struct*/{
          Opening: 25402,
          Closing: 25403,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f1a10",&ast.BlockStmt/*struct*/{
        Lbrace: 25409,
        Rbrace: 25410,
      }/*struct*/),
    }/*struct*/),/* slice_item: 155*/156: p.StrmapFuncDecl("0xc4200f1b60",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f1aa0",&ast.FieldList/*struct*/{
        Opening: 25417,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eedc0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f2740",&ast.StarExpr/*struct*/{
              Star: 25418,
              X: p.StrmapIdent("0xc4200f2720",&ast.Ident/*struct*/{
                NamePos: 25419,
                Name: "SelectStmt",
                Obj: p.PtrmapObject("0xc4200e2500"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25429,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f2760",&ast.Ident/*struct*/{
        NamePos: 25431,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2780",&ast.FuncType/*struct*/{
        Func: 25412,
        Params: p.StrmapFieldList("0xc4200f1ad0",&ast.FieldList/*struct*/{
          Opening: 25439,
          Closing: 25440,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f1b30",&ast.BlockStmt/*struct*/{
        Lbrace: 25446,
        Rbrace: 25447,
      }/*struct*/),
    }/*struct*/),/* slice_item: 156*/157: p.StrmapFuncDecl("0xc4200f1c80",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f1bc0",&ast.FieldList/*struct*/{
        Opening: 25454,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eee00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f27c0",&ast.StarExpr/*struct*/{
              Star: 25455,
              X: p.StrmapIdent("0xc4200f27a0",&ast.Ident/*struct*/{
                NamePos: 25456,
                Name: "ForStmt",
                Obj: p.PtrmapObject("0xc4200e25f0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25463,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f27e0",&ast.Ident/*struct*/{
        NamePos: 25465,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2800",&ast.FuncType/*struct*/{
        Func: 25449,
        Params: p.StrmapFieldList("0xc4200f1bf0",&ast.FieldList/*struct*/{
          Opening: 25473,
          Closing: 25474,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f1c50",&ast.BlockStmt/*struct*/{
        Lbrace: 25483,
        Rbrace: 25484,
      }/*struct*/),
    }/*struct*/),/* slice_item: 157*/158: p.StrmapFuncDecl("0xc4200f1dd0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f1ce0",&ast.FieldList/*struct*/{
        Opening: 25491,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200eee40",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f2840",&ast.StarExpr/*struct*/{
              Star: 25492,
              X: p.StrmapIdent("0xc4200f2820",&ast.Ident/*struct*/{
                NamePos: 25493,
                Name: "RangeStmt",
                Obj: p.PtrmapObject("0xc4200e27d0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 25502,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f2860",&ast.Ident/*struct*/{
        NamePos: 25504,
        Name: "stmtNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f2880",&ast.FuncType/*struct*/{
        Func: 25486,
        Params: p.StrmapFieldList("0xc4200f1d10",&ast.FieldList/*struct*/{
          Opening: 25512,
          Closing: 25513,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f1d70",&ast.BlockStmt/*struct*/{
        Lbrace: 25520,
        Rbrace: 25521,
      }/*struct*/),
    }/*struct*/),/* slice_item: 158*/159: p.StrmapGenDecl("0xc4200ef3c0",&ast.GenDecl/*struct*/{
      TokPos: 25731,
      Tok: token.Token(84)/*type*/,
      Lparen: 25736,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200f1e00",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200f28a0",&ast.Ident/*struct*/{
            NamePos: 25815,
            Name: "Spec",
            Obj: p.StrmapObject("0xc4200e3900",&ast.Object/*struct*/{
              Kind: "type",
              Name: "Spec",
              Decl: p.PtrmapTypeSpec("0xc4200f1e00"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapInterfaceType("0xc4200f2920",&ast.InterfaceType/*struct*/{
            Interface: 25820,
            Methods: p.StrmapFieldList("0xc4200f1e60",&ast.FieldList/*struct*/{
              Opening: 25830,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200eeec0",&ast.Field/*struct*/{
                  Type: p.StrmapIdent("0xc4200f28c0",&ast.Ident/*struct*/{
                    NamePos: 25834,
                    Name: "Node",
                    Obj: p.PtrmapObject("0xc420050550"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200eef00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f28e0",&ast.Ident/*struct*/{
                      NamePos: 25841,
                      Name: "specNode",
                      Obj: p.StrmapObject("0xc4200e3950",&ast.Object/*struct*/{
                        Kind: "func",
                        Name: "specNode",
                        Decl: p.PtrmapField("0xc4200eef00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapFuncType("0xc4200f2900",&ast.FuncType/*struct*/{
                    Func: 0,
                    Params: p.StrmapFieldList("0xc4200f1e30",&ast.FieldList/*struct*/{
                      Opening: 25849,
                      Closing: 25850,
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/}/*slice*/,
              Closing: 25853,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/1: p.StrmapTypeSpec("0xc4200f1e90",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200f2940",&ast.Ident/*struct*/{
            NamePos: 25916,
            Name: "ImportSpec",
            Obj: p.StrmapObject("0xc4200e39a0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "ImportSpec",
              Decl: p.PtrmapTypeSpec("0xc4200f1e90"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200f2ba0",&ast.StructType/*struct*/{
            Struct: 25927,
            Fields: p.StrmapFieldList("0xc4200f1f50",&ast.FieldList/*struct*/{
              Opening: 25934,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200eef80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2960",&ast.Ident/*struct*/{
                      NamePos: 25938,
                      Name: "Doc",
                      Obj: p.StrmapObject("0xc4200e39f0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Doc",
                        Decl: p.PtrmapField("0xc4200eef80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f29a0",&ast.StarExpr/*struct*/{
                    Star: 25946,
                    X: p.StrmapIdent("0xc4200f2980",&ast.Ident/*struct*/{
                      NamePos: 25947,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200eefc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f29c0",&ast.Ident/*struct*/{
                      NamePos: 25998,
                      Name: "Name",
                      Obj: p.StrmapObject("0xc4200e3a40",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Name",
                        Decl: p.PtrmapField("0xc4200eefc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f2a00",&ast.StarExpr/*struct*/{
                    Star: 26006,
                    X: p.StrmapIdent("0xc4200f29e0",&ast.Ident/*struct*/{
                      NamePos: 26007,
                      Name: "Ident",
                      Obj: p.PtrmapObject("0xc420051c70"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200ef000",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2a20",&ast.Ident/*struct*/{
                      NamePos: 26068,
                      Name: "Path",
                      Obj: p.StrmapObject("0xc4200e3a90",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Path",
                        Decl: p.PtrmapField("0xc4200ef000"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f2a60",&ast.StarExpr/*struct*/{
                    Star: 26076,
                    X: p.StrmapIdent("0xc4200f2a40",&ast.Ident/*struct*/{
                      NamePos: 26077,
                      Name: "BasicLit",
                      Obj: p.PtrmapObject("0xc420051ea0"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200ef040",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2aa0",&ast.Ident/*struct*/{
                      NamePos: 26107,
                      Name: "Comment",
                      Obj: p.StrmapObject("0xc4200e3ae0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Comment",
                        Decl: p.PtrmapField("0xc4200ef040"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f2ae0",&ast.StarExpr/*struct*/{
                    Star: 26115,
                    X: p.StrmapIdent("0xc4200f2ac0",&ast.Ident/*struct*/{
                      NamePos: 26116,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200ef080",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2b20",&ast.Ident/*struct*/{
                      NamePos: 26156,
                      Name: "EndPos",
                      Obj: p.StrmapObject("0xc4200e3b30",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "EndPos",
                        Decl: p.PtrmapField("0xc4200ef080"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200f2b80",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f2b40",&ast.Ident/*struct*/{
                      NamePos: 26164,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f2b60",&ast.Ident/*struct*/{
                      NamePos: 26170,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/}/*slice*/,
              Closing: 26226,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 1*/2: p.StrmapTypeSpec("0xc4200f1fb0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200f2be0",&ast.Ident/*struct*/{
            NamePos: 26340,
            Name: "ValueSpec",
            Obj: p.StrmapObject("0xc4200e3bd0",&ast.Object/*struct*/{
              Kind: "type",
              Name: "ValueSpec",
              Decl: p.PtrmapTypeSpec("0xc4200f1fb0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200f2e40",&ast.StructType/*struct*/{
            Struct: 26350,
            Fields: p.StrmapFieldList("0xc4200f4090",&ast.FieldList/*struct*/{
              Opening: 26357,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200ef100",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2c00",&ast.Ident/*struct*/{
                      NamePos: 26361,
                      Name: "Doc",
                      Obj: p.StrmapObject("0xc4200e3c20",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Doc",
                        Decl: p.PtrmapField("0xc4200ef100"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f2c40",&ast.StarExpr/*struct*/{
                    Star: 26369,
                    X: p.StrmapIdent("0xc4200f2c20",&ast.Ident/*struct*/{
                      NamePos: 26370,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200ef140",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2c60",&ast.Ident/*struct*/{
                      NamePos: 26421,
                      Name: "Names",
                      Obj: p.StrmapObject("0xc4200e3c70",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Names",
                        Decl: p.PtrmapField("0xc4200ef140"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200f4030",&ast.ArrayType/*struct*/{
                    Lbrack: 26429,
                    Elt: p.StrmapStarExpr("0xc4200f2ca0",&ast.StarExpr/*struct*/{
                      Star: 26431,
                      X: p.StrmapIdent("0xc4200f2c80",&ast.Ident/*struct*/{
                        NamePos: 26432,
                        Name: "Ident",
                        Obj: p.PtrmapObject("0xc420051c70"),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200ef180",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2ce0",&ast.Ident/*struct*/{
                      NamePos: 26477,
                      Name: "Type",
                      Obj: p.StrmapObject("0xc4200e3cc0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Type",
                        Decl: p.PtrmapField("0xc4200ef180"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200f2d00",&ast.Ident/*struct*/{
                    NamePos: 26485,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200ef1c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2d60",&ast.Ident/*struct*/{
                      NamePos: 26523,
                      Name: "Values",
                      Obj: p.StrmapObject("0xc4200e3d10",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Values",
                        Decl: p.PtrmapField("0xc4200ef1c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200f4060",&ast.ArrayType/*struct*/{
                    Lbrack: 26531,
                    Elt: p.StrmapIdent("0xc4200f2d80",&ast.Ident/*struct*/{
                      NamePos: 26533,
                      Name: "Expr",
                      Obj: p.PtrmapObject("0xc420050640"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200ef200",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2dc0",&ast.Ident/*struct*/{
                      NamePos: 26573,
                      Name: "Comment",
                      Obj: p.StrmapObject("0xc4200e3d60",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Comment",
                        Decl: p.PtrmapField("0xc4200ef200"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f2e00",&ast.StarExpr/*struct*/{
                    Star: 26581,
                    X: p.StrmapIdent("0xc4200f2de0",&ast.Ident/*struct*/{
                      NamePos: 26582,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/}/*slice*/,
              Closing: 26621,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 2*/3: p.StrmapTypeSpec("0xc4200f40c0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200f2e60",&ast.Ident/*struct*/{
            NamePos: 26698,
            Name: "TypeSpec",
            Obj: p.StrmapObject("0xc4200e3e00",&ast.Object/*struct*/{
              Kind: "type",
              Name: "TypeSpec",
              Decl: p.PtrmapTypeSpec("0xc4200f40c0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200f3020",&ast.StructType/*struct*/{
            Struct: 26707,
            Fields: p.StrmapFieldList("0xc4200f4120",&ast.FieldList/*struct*/{
              Opening: 26714,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200ef2c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2e80",&ast.Ident/*struct*/{
                      NamePos: 26718,
                      Name: "Doc",
                      Obj: p.StrmapObject("0xc4200e3e50",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Doc",
                        Decl: p.PtrmapField("0xc4200ef2c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f2ec0",&ast.StarExpr/*struct*/{
                    Star: 26726,
                    X: p.StrmapIdent("0xc4200f2ea0",&ast.Ident/*struct*/{
                      NamePos: 26727,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200ef300",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2ee0",&ast.Ident/*struct*/{
                      NamePos: 26778,
                      Name: "Name",
                      Obj: p.StrmapObject("0xc4200e3ea0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Name",
                        Decl: p.PtrmapField("0xc4200ef300"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f2f20",&ast.StarExpr/*struct*/{
                    Star: 26786,
                    X: p.StrmapIdent("0xc4200f2f00",&ast.Ident/*struct*/{
                      NamePos: 26787,
                      Name: "Ident",
                      Obj: p.PtrmapObject("0xc420051c70"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200ef340",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2f40",&ast.Ident/*struct*/{
                      NamePos: 26815,
                      Name: "Type",
                      Obj: p.StrmapObject("0xc4200e3f40",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Type",
                        Decl: p.PtrmapField("0xc4200ef340"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc4200f2f60",&ast.Ident/*struct*/{
                    NamePos: 26823,
                    Name: "Expr",
                    Obj: p.PtrmapObject("0xc420050640"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200ef380",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f2fa0",&ast.Ident/*struct*/{
                      NamePos: 26912,
                      Name: "Comment",
                      Obj: p.StrmapObject("0xc4200e3f90",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Comment",
                        Decl: p.PtrmapField("0xc4200ef380"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f2fe0",&ast.StarExpr/*struct*/{
                    Star: 26920,
                    X: p.StrmapIdent("0xc4200f2fc0",&ast.Ident/*struct*/{
                      NamePos: 26921,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 26960,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 3*/}/*slice*/,
      Rparen: 26962,
    }/*struct*/),/* slice_item: 159*/160: p.StrmapFuncDecl("0xc4200f4360",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f41b0",&ast.FieldList/*struct*/{
        Opening: 27018,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ef400",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f3040",&ast.Ident/*struct*/{
                NamePos: 27019,
                Name: "s",
                Obj: p.StrmapObject("0xc4200f6000",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ef400"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f3080",&ast.StarExpr/*struct*/{
              Star: 27021,
              X: p.StrmapIdent("0xc4200f3060",&ast.Ident/*struct*/{
                NamePos: 27022,
                Name: "ImportSpec",
                Obj: p.PtrmapObject("0xc4200e39a0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 27032,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f30a0",&ast.Ident/*struct*/{
        NamePos: 27034,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f3340",&ast.FuncType/*struct*/{
        Func: 27013,
        Params: p.StrmapFieldList("0xc4200f41e0",&ast.FieldList/*struct*/{
          Opening: 27037,
          Closing: 27038,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f4210",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ef440",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f3100",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f30c0",&ast.Ident/*struct*/{
                  NamePos: 27040,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f30e0",&ast.Ident/*struct*/{
                  NamePos: 27046,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f4330",&ast.BlockStmt/*struct*/{
        Lbrace: 27050,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200ef4c0",&ast.IfStmt/*struct*/{
            If: 27053,
            Cond: p.StrmapBinaryExpr("0xc4200f42a0",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200f3160",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f3120",&ast.Ident/*struct*/{
                  NamePos: 27056,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200f6000"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f3140",&ast.Ident/*struct*/{
                  NamePos: 27058,
                  Name: "Name",
                }/*struct*/),
              }/*struct*/),
              OpPos: 27063,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200f3180",&ast.Ident/*struct*/{
                NamePos: 27066,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200f4300",&ast.BlockStmt/*struct*/{
              Lbrace: 27070,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200f3240",&ast.ReturnStmt/*struct*/{
                  Return: 27074,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200ef480",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200f3220",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapSelectorExpr("0xc4200f31e0",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200f31a0",&ast.Ident/*struct*/{
                            NamePos: 27081,
                            Name: "s",
                            Obj: p.PtrmapObject("0xc4200f6000"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200f31c0",&ast.Ident/*struct*/{
                            NamePos: 27083,
                            Name: "Name",
                          }/*struct*/),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200f3200",&ast.Ident/*struct*/{
                          NamePos: 27088,
                          Name: "Pos",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 27091,
                      Ellipsis: 0,
                      Rparen: 27092,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 27095,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200f3300",&ast.ReturnStmt/*struct*/{
            Return: 27098,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ef500",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200f32e0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200f32a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f3260",&ast.Ident/*struct*/{
                      NamePos: 27105,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200f6000"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f3280",&ast.Ident/*struct*/{
                      NamePos: 27107,
                      Name: "Path",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f32c0",&ast.Ident/*struct*/{
                    NamePos: 27112,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 27115,
                Ellipsis: 0,
                Rparen: 27116,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 27118,
      }/*struct*/),
    }/*struct*/),/* slice_item: 160*/161: p.StrmapFuncDecl("0xc4200f44e0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f43c0",&ast.FieldList/*struct*/{
        Opening: 27125,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ef540",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f3360",&ast.Ident/*struct*/{
                NamePos: 27126,
                Name: "s",
                Obj: p.StrmapObject("0xc4200f6050",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ef540"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f33a0",&ast.StarExpr/*struct*/{
              Star: 27128,
              X: p.StrmapIdent("0xc4200f3380",&ast.Ident/*struct*/{
                NamePos: 27129,
                Name: "ValueSpec",
                Obj: p.PtrmapObject("0xc4200e3bd0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 27138,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f33c0",&ast.Ident/*struct*/{
        NamePos: 27140,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f3520",&ast.FuncType/*struct*/{
        Func: 27120,
        Params: p.StrmapFieldList("0xc4200f43f0",&ast.FieldList/*struct*/{
          Opening: 27143,
          Closing: 27144,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f4420",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ef580",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f3420",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f33e0",&ast.Ident/*struct*/{
                  NamePos: 27146,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f3400",&ast.Ident/*struct*/{
                  NamePos: 27152,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f44b0",&ast.BlockStmt/*struct*/{
        Lbrace: 27156,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200f3500",&ast.ReturnStmt/*struct*/{
            Return: 27158,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ef5c0",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200f34e0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIndexExpr("0xc4200f4480",&ast.IndexExpr/*struct*/{
                    X: p.StrmapSelectorExpr("0xc4200f3480",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200f3440",&ast.Ident/*struct*/{
                        NamePos: 27165,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200f6050"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200f3460",&ast.Ident/*struct*/{
                        NamePos: 27167,
                        Name: "Names",
                      }/*struct*/),
                    }/*struct*/),
                    Lbrack: 27172,
                    Index: p.StrmapBasicLit("0xc4200f34a0",&ast.BasicLit/*struct*/{
                      ValuePos: 27173,
                      Kind: token.Token(5)/*INT*/,
                      Value: "0",
                    }/*struct*/),
                    Rbrack: 27174,
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f34c0",&ast.Ident/*struct*/{
                    NamePos: 27176,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 27179,
                Ellipsis: 0,
                Rparen: 27180,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 27182,
      }/*struct*/),
    }/*struct*/),/* slice_item: 161*/162: p.StrmapFuncDecl("0xc4200f4630",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f4540",&ast.FieldList/*struct*/{
        Opening: 27189,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ef600",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f3540",&ast.Ident/*struct*/{
                NamePos: 27190,
                Name: "s",
                Obj: p.StrmapObject("0xc4200f60a0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ef600"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f3580",&ast.StarExpr/*struct*/{
              Star: 27192,
              X: p.StrmapIdent("0xc4200f3560",&ast.Ident/*struct*/{
                NamePos: 27193,
                Name: "TypeSpec",
                Obj: p.PtrmapObject("0xc4200e3e00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 27201,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f35a0",&ast.Ident/*struct*/{
        NamePos: 27203,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f36e0",&ast.FuncType/*struct*/{
        Func: 27184,
        Params: p.StrmapFieldList("0xc4200f4570",&ast.FieldList/*struct*/{
          Opening: 27206,
          Closing: 27207,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f45a0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ef640",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f3600",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f35c0",&ast.Ident/*struct*/{
                  NamePos: 27209,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f35e0",&ast.Ident/*struct*/{
                  NamePos: 27215,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f4600",&ast.BlockStmt/*struct*/{
        Lbrace: 27220,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200f36c0",&ast.ReturnStmt/*struct*/{
            Return: 27222,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ef680",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200f36a0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200f3660",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f3620",&ast.Ident/*struct*/{
                      NamePos: 27229,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200f60a0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f3640",&ast.Ident/*struct*/{
                      NamePos: 27231,
                      Name: "Name",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f3680",&ast.Ident/*struct*/{
                    NamePos: 27236,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 27239,
                Ellipsis: 0,
                Rparen: 27240,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 27242,
      }/*struct*/),
    }/*struct*/),/* slice_item: 162*/163: p.StrmapFuncDecl("0xc4200f4840",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f4690",&ast.FieldList/*struct*/{
        Opening: 27250,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ef6c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f3700",&ast.Ident/*struct*/{
                NamePos: 27251,
                Name: "s",
                Obj: p.StrmapObject("0xc4200f60f0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ef6c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f3740",&ast.StarExpr/*struct*/{
              Star: 27253,
              X: p.StrmapIdent("0xc4200f3720",&ast.Ident/*struct*/{
                NamePos: 27254,
                Name: "ImportSpec",
                Obj: p.PtrmapObject("0xc4200e39a0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 27264,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f3760",&ast.Ident/*struct*/{
        NamePos: 27266,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f39c0",&ast.FuncType/*struct*/{
        Func: 27245,
        Params: p.StrmapFieldList("0xc4200f46c0",&ast.FieldList/*struct*/{
          Opening: 27269,
          Closing: 27270,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f46f0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ef700",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f37c0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f3780",&ast.Ident/*struct*/{
                  NamePos: 27272,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f37a0",&ast.Ident/*struct*/{
                  NamePos: 27278,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f4810",&ast.BlockStmt/*struct*/{
        Lbrace: 27282,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200ef740",&ast.IfStmt/*struct*/{
            If: 27285,
            Cond: p.StrmapBinaryExpr("0xc4200f4780",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200f3820",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f37e0",&ast.Ident/*struct*/{
                  NamePos: 27288,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200f60f0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f3800",&ast.Ident/*struct*/{
                  NamePos: 27290,
                  Name: "EndPos",
                }/*struct*/),
              }/*struct*/),
              OpPos: 27297,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapBasicLit("0xc4200f3840",&ast.BasicLit/*struct*/{
                ValuePos: 27300,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200f47e0",&ast.BlockStmt/*struct*/{
              Lbrace: 27302,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200f38c0",&ast.ReturnStmt/*struct*/{
                  Return: 27306,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200f38a0",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200f3860",&ast.Ident/*struct*/{
                        NamePos: 27313,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200f60f0"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200f3880",&ast.Ident/*struct*/{
                        NamePos: 27315,
                        Name: "EndPos",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 27323,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200f3980",&ast.ReturnStmt/*struct*/{
            Return: 27326,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200ef780",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200f3960",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200f3920",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f38e0",&ast.Ident/*struct*/{
                      NamePos: 27333,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200f60f0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f3900",&ast.Ident/*struct*/{
                      NamePos: 27335,
                      Name: "Path",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f3940",&ast.Ident/*struct*/{
                    NamePos: 27340,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 27343,
                Ellipsis: 0,
                Rparen: 27344,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 27346,
      }/*struct*/),
    }/*struct*/),/* slice_item: 163*/164: p.StrmapFuncDecl("0xc4200f4bd0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f48a0",&ast.FieldList/*struct*/{
        Opening: 27354,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200ef7c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f39e0",&ast.Ident/*struct*/{
                NamePos: 27355,
                Name: "s",
                Obj: p.StrmapObject("0xc4200f6140",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200ef7c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f3a20",&ast.StarExpr/*struct*/{
              Star: 27357,
              X: p.StrmapIdent("0xc4200f3a00",&ast.Ident/*struct*/{
                NamePos: 27358,
                Name: "ValueSpec",
                Obj: p.PtrmapObject("0xc4200e3bd0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 27367,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f3a40",&ast.Ident/*struct*/{
        NamePos: 27369,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f3f60",&ast.FuncType/*struct*/{
        Func: 27349,
        Params: p.StrmapFieldList("0xc4200f48d0",&ast.FieldList/*struct*/{
          Opening: 27372,
          Closing: 27373,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f4900",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200ef800",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f3aa0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f3a60",&ast.Ident/*struct*/{
                  NamePos: 27375,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f3a80",&ast.Ident/*struct*/{
                  NamePos: 27381,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f4ba0",&ast.BlockStmt/*struct*/{
        Lbrace: 27385,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200ef900",&ast.IfStmt/*struct*/{
            If: 27388,
            Init: p.StrmapAssignStmt("0xc4200ef880",&ast.AssignStmt/*struct*/{
              Lhs: []ast.Expr /*Slice*/{
                0: p.StrmapIdent("0xc4200f3ac0",&ast.Ident/*struct*/{
                  NamePos: 27391,
                  Name: "n",
                  Obj: p.StrmapObject("0xc4200f6190",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "n",
                    Decl: p.PtrmapAssignStmt("0xc4200ef880"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              TokPos: 27393,
              Tok: token.Token(47)/*:=*/,
              Rhs: []ast.Expr /*Slice*/{
                0: p.StrmapCallExpr("0xc4200ef840",&ast.CallExpr/*struct*/{
                  Fun: p.StrmapIdent("0xc4200f3ae0",&ast.Ident/*struct*/{
                    NamePos: 27396,
                    Name: "len",
                  }/*struct*/),
                  Lparen: 27399,
                  Args: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200f3b40",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200f3b00",&ast.Ident/*struct*/{
                        NamePos: 27400,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200f6140"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200f3b20",&ast.Ident/*struct*/{
                        NamePos: 27402,
                        Name: "Values",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Ellipsis: 0,
                  Rparen: 27408,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
            }/*struct*/),
            Cond: p.StrmapBinaryExpr("0xc4200f4990",&ast.BinaryExpr/*struct*/{
              X: p.StrmapIdent("0xc4200f3b60",&ast.Ident/*struct*/{
                NamePos: 27411,
                Name: "n",
                Obj: p.PtrmapObject("0xc4200f6190"),
              }/*struct*/),
              OpPos: 27413,
              Op: token.Token(41)/*>*/,
              Y: p.StrmapBasicLit("0xc4200f3b80",&ast.BasicLit/*struct*/{
                ValuePos: 27415,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200f4a50",&ast.BlockStmt/*struct*/{
              Lbrace: 27417,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200f3c80",&ast.ReturnStmt/*struct*/{
                  Return: 27421,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200ef8c0",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200f3c60",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIndexExpr("0xc4200f4a20",&ast.IndexExpr/*struct*/{
                          X: p.StrmapSelectorExpr("0xc4200f3be0",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200f3ba0",&ast.Ident/*struct*/{
                              NamePos: 27428,
                              Name: "s",
                              Obj: p.PtrmapObject("0xc4200f6140"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200f3bc0",&ast.Ident/*struct*/{
                              NamePos: 27430,
                              Name: "Values",
                            }/*struct*/),
                          }/*struct*/),
                          Lbrack: 27436,
                          Index: p.StrmapBinaryExpr("0xc4200f49f0",&ast.BinaryExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200f3c00",&ast.Ident/*struct*/{
                              NamePos: 27437,
                              Name: "n",
                              Obj: p.PtrmapObject("0xc4200f6190"),
                            }/*struct*/),
                            OpPos: 27438,
                            Op: token.Token(13)/*-*/,
                            Y: p.StrmapBasicLit("0xc4200f3c20",&ast.BasicLit/*struct*/{
                              ValuePos: 27439,
                              Kind: token.Token(5)/*INT*/,
                              Value: "1",
                            }/*struct*/),
                          }/*struct*/),
                          Rbrack: 27440,
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200f3c40",&ast.Ident/*struct*/{
                          NamePos: 27442,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 27445,
                      Ellipsis: 0,
                      Rparen: 27446,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 27449,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapIfStmt("0xc4200ef980",&ast.IfStmt/*struct*/{
            If: 27452,
            Cond: p.StrmapBinaryExpr("0xc4200f4ab0",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200f3ce0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f3ca0",&ast.Ident/*struct*/{
                  NamePos: 27455,
                  Name: "s",
                  Obj: p.PtrmapObject("0xc4200f6140"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f3cc0",&ast.Ident/*struct*/{
                  NamePos: 27457,
                  Name: "Type",
                }/*struct*/),
              }/*struct*/),
              OpPos: 27462,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200f3d00",&ast.Ident/*struct*/{
                NamePos: 27465,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200f4b10",&ast.BlockStmt/*struct*/{
              Lbrace: 27469,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200f3dc0",&ast.ReturnStmt/*struct*/{
                  Return: 27473,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200ef940",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200f3da0",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapSelectorExpr("0xc4200f3d60",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200f3d20",&ast.Ident/*struct*/{
                            NamePos: 27480,
                            Name: "s",
                            Obj: p.PtrmapObject("0xc4200f6140"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200f3d40",&ast.Ident/*struct*/{
                            NamePos: 27482,
                            Name: "Type",
                          }/*struct*/),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200f3d80",&ast.Ident/*struct*/{
                          NamePos: 27487,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 27490,
                      Ellipsis: 0,
                      Rparen: 27491,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 27494,
            }/*struct*/),
          }/*struct*/),/* slice_item: 1*/2: p.StrmapReturnStmt("0xc4200f3f40",&ast.ReturnStmt/*struct*/{
            Return: 27497,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200efa00",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200f3f20",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIndexExpr("0xc4200f4b70",&ast.IndexExpr/*struct*/{
                    X: p.StrmapSelectorExpr("0xc4200f3e40",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200f3e00",&ast.Ident/*struct*/{
                        NamePos: 27504,
                        Name: "s",
                        Obj: p.PtrmapObject("0xc4200f6140"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200f3e20",&ast.Ident/*struct*/{
                        NamePos: 27506,
                        Name: "Names",
                      }/*struct*/),
                    }/*struct*/),
                    Lbrack: 27511,
                    Index: p.StrmapBinaryExpr("0xc4200f4b40",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapCallExpr("0xc4200ef9c0",&ast.CallExpr/*struct*/{
                        Fun: p.StrmapIdent("0xc4200f3e60",&ast.Ident/*struct*/{
                          NamePos: 27512,
                          Name: "len",
                        }/*struct*/),
                        Lparen: 27515,
                        Args: []ast.Expr /*Slice*/{
                          0: p.StrmapSelectorExpr("0xc4200f3ec0",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200f3e80",&ast.Ident/*struct*/{
                              NamePos: 27516,
                              Name: "s",
                              Obj: p.PtrmapObject("0xc4200f6140"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200f3ea0",&ast.Ident/*struct*/{
                              NamePos: 27518,
                              Name: "Names",
                            }/*struct*/),
                          }/*struct*/),/* slice_item: 0*/}/*slice*/,
                        Ellipsis: 0,
                        Rparen: 27523,
                      }/*struct*/),
                      OpPos: 27524,
                      Op: token.Token(13)/*-*/,
                      Y: p.StrmapBasicLit("0xc4200f3ee0",&ast.BasicLit/*struct*/{
                        ValuePos: 27525,
                        Kind: token.Token(5)/*INT*/,
                        Value: "1",
                      }/*struct*/),
                    }/*struct*/),
                    Rbrack: 27526,
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f3f00",&ast.Ident/*struct*/{
                    NamePos: 27528,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 27531,
                Ellipsis: 0,
                Rparen: 27532,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 2*/}/*slice*/,
        Rbrace: 27534,
      }/*struct*/),
    }/*struct*/),/* slice_item: 164*/165: p.StrmapFuncDecl("0xc4200f4d20",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f4c30",&ast.FieldList/*struct*/{
        Opening: 27541,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200efa80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f3f80",&ast.Ident/*struct*/{
                NamePos: 27542,
                Name: "s",
                Obj: p.StrmapObject("0xc4200f61e0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "s",
                  Decl: p.PtrmapField("0xc4200efa80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f3fc0",&ast.StarExpr/*struct*/{
              Star: 27544,
              X: p.StrmapIdent("0xc4200f3fa0",&ast.Ident/*struct*/{
                NamePos: 27545,
                Name: "TypeSpec",
                Obj: p.PtrmapObject("0xc4200e3e00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 27553,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f3fe0",&ast.Ident/*struct*/{
        NamePos: 27555,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f8140",&ast.FuncType/*struct*/{
        Func: 27536,
        Params: p.StrmapFieldList("0xc4200f4c60",&ast.FieldList/*struct*/{
          Opening: 27558,
          Closing: 27559,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f4c90",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200efac0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f8040",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f8000",&ast.Ident/*struct*/{
                  NamePos: 27561,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f8020",&ast.Ident/*struct*/{
                  NamePos: 27567,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f4cf0",&ast.BlockStmt/*struct*/{
        Lbrace: 27571,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200f8100",&ast.ReturnStmt/*struct*/{
            Return: 27573,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200efb00",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200f80e0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200f80a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f8060",&ast.Ident/*struct*/{
                      NamePos: 27580,
                      Name: "s",
                      Obj: p.PtrmapObject("0xc4200f61e0"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f8080",&ast.Ident/*struct*/{
                      NamePos: 27582,
                      Name: "Type",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f80c0",&ast.Ident/*struct*/{
                    NamePos: 27587,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 27590,
                Ellipsis: 0,
                Rparen: 27591,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 27593,
      }/*struct*/),
    }/*struct*/),/* slice_item: 165*/166: p.StrmapFuncDecl("0xc4200f4e40",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f4d80",&ast.FieldList/*struct*/{
        Opening: 27677,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200efb80",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f8180",&ast.StarExpr/*struct*/{
              Star: 27678,
              X: p.StrmapIdent("0xc4200f8160",&ast.Ident/*struct*/{
                NamePos: 27679,
                Name: "ImportSpec",
                Obj: p.PtrmapObject("0xc4200e39a0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 27689,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f81a0",&ast.Ident/*struct*/{
        NamePos: 27691,
        Name: "specNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f81c0",&ast.FuncType/*struct*/{
        Func: 27672,
        Params: p.StrmapFieldList("0xc4200f4db0",&ast.FieldList/*struct*/{
          Opening: 27699,
          Closing: 27700,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f4e10",&ast.BlockStmt/*struct*/{
        Lbrace: 27702,
        Rbrace: 27703,
      }/*struct*/),
    }/*struct*/),/* slice_item: 166*/167: p.StrmapFuncDecl("0xc4200f4f60",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f4ea0",&ast.FieldList/*struct*/{
        Opening: 27710,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200efbc0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f8200",&ast.StarExpr/*struct*/{
              Star: 27711,
              X: p.StrmapIdent("0xc4200f81e0",&ast.Ident/*struct*/{
                NamePos: 27712,
                Name: "ValueSpec",
                Obj: p.PtrmapObject("0xc4200e3bd0"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 27721,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f8220",&ast.Ident/*struct*/{
        NamePos: 27723,
        Name: "specNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f8240",&ast.FuncType/*struct*/{
        Func: 27705,
        Params: p.StrmapFieldList("0xc4200f4ed0",&ast.FieldList/*struct*/{
          Opening: 27731,
          Closing: 27732,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f4f30",&ast.BlockStmt/*struct*/{
        Lbrace: 27735,
        Rbrace: 27736,
      }/*struct*/),
    }/*struct*/),/* slice_item: 167*/168: p.StrmapFuncDecl("0xc4200f5080",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f4fc0",&ast.FieldList/*struct*/{
        Opening: 27743,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200efc00",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f8280",&ast.StarExpr/*struct*/{
              Star: 27744,
              X: p.StrmapIdent("0xc4200f8260",&ast.Ident/*struct*/{
                NamePos: 27745,
                Name: "TypeSpec",
                Obj: p.PtrmapObject("0xc4200e3e00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 27753,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f82a0",&ast.Ident/*struct*/{
        NamePos: 27755,
        Name: "specNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f82c0",&ast.FuncType/*struct*/{
        Func: 27738,
        Params: p.StrmapFieldList("0xc4200f4ff0",&ast.FieldList/*struct*/{
          Opening: 27763,
          Closing: 27764,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f5050",&ast.BlockStmt/*struct*/{
        Lbrace: 27768,
        Rbrace: 27769,
      }/*struct*/),
    }/*struct*/),/* slice_item: 168*/169: p.StrmapGenDecl("0xc4200fc180",&ast.GenDecl/*struct*/{
      TokPos: 27850,
      Tok: token.Token(84)/*type*/,
      Lparen: 27855,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200f50b0",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200f82e0",&ast.Ident/*struct*/{
            NamePos: 28003,
            Name: "BadDecl",
            Obj: p.StrmapObject("0xc4200f6280",&ast.Object/*struct*/{
              Kind: "type",
              Name: "BadDecl",
              Decl: p.PtrmapTypeSpec("0xc4200f50b0"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200f83c0",&ast.StructType/*struct*/{
            Struct: 28011,
            Fields: p.StrmapFieldList("0xc4200f5110",&ast.FieldList/*struct*/{
              Opening: 28018,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200efcc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f8300",&ast.Ident/*struct*/{
                      NamePos: 28022,
                      Name: "From",
                      Obj: p.StrmapObject("0xc4200f62d0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "From",
                        Decl: p.PtrmapField("0xc4200efcc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/1: p.StrmapIdent("0xc4200f8320",&ast.Ident/*struct*/{
                      NamePos: 28028,
                      Name: "To",
                      Obj: p.StrmapObject("0xc4200f6320",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "To",
                        Decl: p.PtrmapField("0xc4200efcc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 1*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200f83a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f8360",&ast.Ident/*struct*/{
                      NamePos: 28031,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f8380",&ast.Ident/*struct*/{
                      NamePos: 28037,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Closing: 28079,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/1: p.StrmapTypeSpec("0xc4200f5140",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200f8460",&ast.Ident/*struct*/{
            NamePos: 28467,
            Name: "GenDecl",
            Obj: p.StrmapObject("0xc4200f6410",&ast.Object/*struct*/{
              Kind: "type",
              Name: "GenDecl",
              Decl: p.PtrmapTypeSpec("0xc4200f5140"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200f87c0",&ast.StructType/*struct*/{
            Struct: 28475,
            Fields: p.StrmapFieldList("0xc4200f51d0",&ast.FieldList/*struct*/{
              Opening: 28482,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200efd80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f8480",&ast.Ident/*struct*/{
                      NamePos: 28486,
                      Name: "Doc",
                      Obj: p.StrmapObject("0xc4200f6460",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Doc",
                        Decl: p.PtrmapField("0xc4200efd80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f84c0",&ast.StarExpr/*struct*/{
                    Star: 28493,
                    X: p.StrmapIdent("0xc4200f84a0",&ast.Ident/*struct*/{
                      NamePos: 28494,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200efdc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f84e0",&ast.Ident/*struct*/{
                      NamePos: 28545,
                      Name: "TokPos",
                      Obj: p.StrmapObject("0xc4200f64b0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "TokPos",
                        Decl: p.PtrmapField("0xc4200efdc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200f8540",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f8500",&ast.Ident/*struct*/{
                      NamePos: 28552,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f8520",&ast.Ident/*struct*/{
                      NamePos: 28558,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200efe00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f8580",&ast.Ident/*struct*/{
                      NamePos: 28587,
                      Name: "Tok",
                      Obj: p.StrmapObject("0xc4200f6500",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Tok",
                        Decl: p.PtrmapField("0xc4200efe00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200f85e0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f85a0",&ast.Ident/*struct*/{
                      NamePos: 28594,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f85c0",&ast.Ident/*struct*/{
                      NamePos: 28600,
                      Name: "Token",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200efe40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f8640",&ast.Ident/*struct*/{
                      NamePos: 28638,
                      Name: "Lparen",
                      Obj: p.StrmapObject("0xc4200f6550",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Lparen",
                        Decl: p.PtrmapField("0xc4200efe40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200f86a0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f8660",&ast.Ident/*struct*/{
                      NamePos: 28645,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f8680",&ast.Ident/*struct*/{
                      NamePos: 28651,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200efe80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f86e0",&ast.Ident/*struct*/{
                      NamePos: 28688,
                      Name: "Specs",
                      Obj: p.StrmapObject("0xc4200f65a0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Specs",
                        Decl: p.PtrmapField("0xc4200efe80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200f51a0",&ast.ArrayType/*struct*/{
                    Lbrack: 28695,
                    Elt: p.StrmapIdent("0xc4200f8700",&ast.Ident/*struct*/{
                      NamePos: 28697,
                      Name: "Spec",
                      Obj: p.PtrmapObject("0xc4200e3900"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/5: p.StrmapField("0xc4200eff00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f8720",&ast.Ident/*struct*/{
                      NamePos: 28704,
                      Name: "Rparen",
                      Obj: p.StrmapObject("0xc4200f65f0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Rparen",
                        Decl: p.PtrmapField("0xc4200eff00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200f8780",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f8740",&ast.Ident/*struct*/{
                      NamePos: 28711,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f8760",&ast.Ident/*struct*/{
                      NamePos: 28717,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 5*/}/*slice*/,
              Closing: 28749,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 1*/2: p.StrmapTypeSpec("0xc4200f5200",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200f8800",&ast.Ident/*struct*/{
            NamePos: 28808,
            Name: "FuncDecl",
            Obj: p.StrmapObject("0xc4200f6640",&ast.Object/*struct*/{
              Kind: "type",
              Name: "FuncDecl",
              Decl: p.PtrmapTypeSpec("0xc4200f5200"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200f8a40",&ast.StructType/*struct*/{
            Struct: 28817,
            Fields: p.StrmapFieldList("0xc4200f52c0",&ast.FieldList/*struct*/{
              Opening: 28824,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200eff80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f8820",&ast.Ident/*struct*/{
                      NamePos: 28828,
                      Name: "Doc",
                      Obj: p.StrmapObject("0xc4200f6690",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Doc",
                        Decl: p.PtrmapField("0xc4200eff80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f8860",&ast.StarExpr/*struct*/{
                    Star: 28833,
                    X: p.StrmapIdent("0xc4200f8840",&ast.Ident/*struct*/{
                      NamePos: 28834,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200effc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f8880",&ast.Ident/*struct*/{
                      NamePos: 28885,
                      Name: "Recv",
                      Obj: p.StrmapObject("0xc4200f66e0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Recv",
                        Decl: p.PtrmapField("0xc4200effc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f88c0",&ast.StarExpr/*struct*/{
                    Star: 28890,
                    X: p.StrmapIdent("0xc4200f88a0",&ast.Ident/*struct*/{
                      NamePos: 28891,
                      Name: "FieldList",
                      Obj: p.PtrmapObject("0xc420051720"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200fc000",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f88e0",&ast.Ident/*struct*/{
                      NamePos: 28948,
                      Name: "Name",
                      Obj: p.StrmapObject("0xc4200f6730",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Name",
                        Decl: p.PtrmapField("0xc4200fc000"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f8920",&ast.StarExpr/*struct*/{
                    Star: 28953,
                    X: p.StrmapIdent("0xc4200f8900",&ast.Ident/*struct*/{
                      NamePos: 28954,
                      Name: "Ident",
                      Obj: p.PtrmapObject("0xc420051c70"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200fc040",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f8980",&ast.Ident/*struct*/{
                      NamePos: 28993,
                      Name: "Type",
                      Obj: p.StrmapObject("0xc4200f67d0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Type",
                        Decl: p.PtrmapField("0xc4200fc040"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f89c0",&ast.StarExpr/*struct*/{
                    Star: 28998,
                    X: p.StrmapIdent("0xc4200f89a0",&ast.Ident/*struct*/{
                      NamePos: 28999,
                      Name: "FuncType",
                      Obj: p.PtrmapObject("0xc4200c1770"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200fc080",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f89e0",&ast.Ident/*struct*/{
                      NamePos: 29089,
                      Name: "Body",
                      Obj: p.StrmapObject("0xc4200f6820",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Body",
                        Decl: p.PtrmapField("0xc4200fc080"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f8a20",&ast.StarExpr/*struct*/{
                    Star: 29094,
                    X: p.StrmapIdent("0xc4200f8a00",&ast.Ident/*struct*/{
                      NamePos: 29095,
                      Name: "BlockStmt",
                      Obj: p.PtrmapObject("0xc4200c9b30"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/}/*slice*/,
              Closing: 29156,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 2*/}/*slice*/,
      Rparen: 29158,
    }/*struct*/),/* slice_item: 169*/170: p.StrmapFuncDecl("0xc4200f5410",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f5320",&ast.FieldList/*struct*/{
        Opening: 29221,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fc1c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f8a60",&ast.Ident/*struct*/{
                NamePos: 29222,
                Name: "d",
                Obj: p.StrmapObject("0xc4200f6870",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "d",
                  Decl: p.PtrmapField("0xc4200fc1c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f8aa0",&ast.StarExpr/*struct*/{
              Star: 29224,
              X: p.StrmapIdent("0xc4200f8a80",&ast.Ident/*struct*/{
                NamePos: 29225,
                Name: "BadDecl",
                Obj: p.PtrmapObject("0xc4200f6280"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 29232,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f8ac0",&ast.Ident/*struct*/{
        NamePos: 29234,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f8bc0",&ast.FuncType/*struct*/{
        Func: 29216,
        Params: p.StrmapFieldList("0xc4200f5350",&ast.FieldList/*struct*/{
          Opening: 29237,
          Closing: 29238,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f5380",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fc200",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f8b20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f8ae0",&ast.Ident/*struct*/{
                  NamePos: 29240,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f8b00",&ast.Ident/*struct*/{
                  NamePos: 29246,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f53e0",&ast.BlockStmt/*struct*/{
        Lbrace: 29251,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200f8ba0",&ast.ReturnStmt/*struct*/{
            Return: 29253,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200f8b80",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f8b40",&ast.Ident/*struct*/{
                  NamePos: 29260,
                  Name: "d",
                  Obj: p.PtrmapObject("0xc4200f6870"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f8b60",&ast.Ident/*struct*/{
                  NamePos: 29262,
                  Name: "From",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 29267,
      }/*struct*/),
    }/*struct*/),/* slice_item: 170*/171: p.StrmapFuncDecl("0xc4200f5560",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f5470",&ast.FieldList/*struct*/{
        Opening: 29274,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fc240",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f8be0",&ast.Ident/*struct*/{
                NamePos: 29275,
                Name: "d",
                Obj: p.StrmapObject("0xc4200f68c0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "d",
                  Decl: p.PtrmapField("0xc4200fc240"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f8c20",&ast.StarExpr/*struct*/{
              Star: 29277,
              X: p.StrmapIdent("0xc4200f8c00",&ast.Ident/*struct*/{
                NamePos: 29278,
                Name: "GenDecl",
                Obj: p.PtrmapObject("0xc4200f6410"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 29285,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f8c40",&ast.Ident/*struct*/{
        NamePos: 29287,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f8d40",&ast.FuncType/*struct*/{
        Func: 29269,
        Params: p.StrmapFieldList("0xc4200f54a0",&ast.FieldList/*struct*/{
          Opening: 29290,
          Closing: 29291,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f54d0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fc280",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f8ca0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f8c60",&ast.Ident/*struct*/{
                  NamePos: 29293,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f8c80",&ast.Ident/*struct*/{
                  NamePos: 29299,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f5530",&ast.BlockStmt/*struct*/{
        Lbrace: 29304,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200f8d20",&ast.ReturnStmt/*struct*/{
            Return: 29306,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200f8d00",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f8cc0",&ast.Ident/*struct*/{
                  NamePos: 29313,
                  Name: "d",
                  Obj: p.PtrmapObject("0xc4200f68c0"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f8ce0",&ast.Ident/*struct*/{
                  NamePos: 29315,
                  Name: "TokPos",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 29322,
      }/*struct*/),
    }/*struct*/),/* slice_item: 171*/172: p.StrmapFuncDecl("0xc4200f56b0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f55c0",&ast.FieldList/*struct*/{
        Opening: 29329,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fc2c0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f8d60",&ast.Ident/*struct*/{
                NamePos: 29330,
                Name: "d",
                Obj: p.StrmapObject("0xc4200f6910",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "d",
                  Decl: p.PtrmapField("0xc4200fc2c0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f8da0",&ast.StarExpr/*struct*/{
              Star: 29332,
              X: p.StrmapIdent("0xc4200f8d80",&ast.Ident/*struct*/{
                NamePos: 29333,
                Name: "FuncDecl",
                Obj: p.PtrmapObject("0xc4200f6640"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 29341,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f8dc0",&ast.Ident/*struct*/{
        NamePos: 29343,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f8f00",&ast.FuncType/*struct*/{
        Func: 29324,
        Params: p.StrmapFieldList("0xc4200f55f0",&ast.FieldList/*struct*/{
          Opening: 29346,
          Closing: 29347,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f5620",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fc300",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f8e20",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f8de0",&ast.Ident/*struct*/{
                  NamePos: 29349,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f8e00",&ast.Ident/*struct*/{
                  NamePos: 29355,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f5680",&ast.BlockStmt/*struct*/{
        Lbrace: 29359,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200f8ee0",&ast.ReturnStmt/*struct*/{
            Return: 29361,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200fc340",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200f8ec0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200f8e80",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f8e40",&ast.Ident/*struct*/{
                      NamePos: 29368,
                      Name: "d",
                      Obj: p.PtrmapObject("0xc4200f6910"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f8e60",&ast.Ident/*struct*/{
                      NamePos: 29370,
                      Name: "Type",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f8ea0",&ast.Ident/*struct*/{
                    NamePos: 29375,
                    Name: "Pos",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 29378,
                Ellipsis: 0,
                Rparen: 29379,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 29381,
      }/*struct*/),
    }/*struct*/),/* slice_item: 172*/173: p.StrmapFuncDecl("0xc4200f5800",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f5710",&ast.FieldList/*struct*/{
        Opening: 29389,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fc380",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f8f20",&ast.Ident/*struct*/{
                NamePos: 29390,
                Name: "d",
                Obj: p.StrmapObject("0xc4200f6960",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "d",
                  Decl: p.PtrmapField("0xc4200fc380"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f8f60",&ast.StarExpr/*struct*/{
              Star: 29392,
              X: p.StrmapIdent("0xc4200f8f40",&ast.Ident/*struct*/{
                NamePos: 29393,
                Name: "BadDecl",
                Obj: p.PtrmapObject("0xc4200f6280"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 29400,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f8f80",&ast.Ident/*struct*/{
        NamePos: 29402,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f9080",&ast.FuncType/*struct*/{
        Func: 29384,
        Params: p.StrmapFieldList("0xc4200f5740",&ast.FieldList/*struct*/{
          Opening: 29405,
          Closing: 29406,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f5770",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fc3c0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f8fe0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f8fa0",&ast.Ident/*struct*/{
                  NamePos: 29408,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f8fc0",&ast.Ident/*struct*/{
                  NamePos: 29414,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f57d0",&ast.BlockStmt/*struct*/{
        Lbrace: 29418,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200f9060",&ast.ReturnStmt/*struct*/{
            Return: 29420,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200f9040",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f9000",&ast.Ident/*struct*/{
                  NamePos: 29427,
                  Name: "d",
                  Obj: p.PtrmapObject("0xc4200f6960"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f9020",&ast.Ident/*struct*/{
                  NamePos: 29429,
                  Name: "To",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 29432,
      }/*struct*/),
    }/*struct*/),/* slice_item: 173*/174: p.StrmapFuncDecl("0xc4200f5a40",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f5860",&ast.FieldList/*struct*/{
        Opening: 29439,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fc400",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f90a0",&ast.Ident/*struct*/{
                NamePos: 29440,
                Name: "d",
                Obj: p.StrmapObject("0xc4200f69b0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "d",
                  Decl: p.PtrmapField("0xc4200fc400"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f90e0",&ast.StarExpr/*struct*/{
              Star: 29442,
              X: p.StrmapIdent("0xc4200f90c0",&ast.Ident/*struct*/{
                NamePos: 29443,
                Name: "GenDecl",
                Obj: p.PtrmapObject("0xc4200f6410"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 29450,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f9100",&ast.Ident/*struct*/{
        NamePos: 29452,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f93c0",&ast.FuncType/*struct*/{
        Func: 29434,
        Params: p.StrmapFieldList("0xc4200f5890",&ast.FieldList/*struct*/{
          Opening: 29455,
          Closing: 29456,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f58c0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fc440",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f9160",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f9120",&ast.Ident/*struct*/{
                  NamePos: 29458,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f9140",&ast.Ident/*struct*/{
                  NamePos: 29464,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f5a10",&ast.BlockStmt/*struct*/{
        Lbrace: 29468,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200fc4c0",&ast.IfStmt/*struct*/{
            If: 29471,
            Cond: p.StrmapCallExpr("0xc4200fc480",&ast.CallExpr/*struct*/{
              Fun: p.StrmapSelectorExpr("0xc4200f9200",&ast.SelectorExpr/*struct*/{
                X: p.StrmapSelectorExpr("0xc4200f91c0",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIdent("0xc4200f9180",&ast.Ident/*struct*/{
                    NamePos: 29474,
                    Name: "d",
                    Obj: p.PtrmapObject("0xc4200f69b0"),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f91a0",&ast.Ident/*struct*/{
                    NamePos: 29476,
                    Name: "Rparen",
                  }/*struct*/),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f91e0",&ast.Ident/*struct*/{
                  NamePos: 29483,
                  Name: "IsValid",
                }/*struct*/),
              }/*struct*/),
              Lparen: 29490,
              Ellipsis: 0,
              Rparen: 29491,
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200f59b0",&ast.BlockStmt/*struct*/{
              Lbrace: 29493,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200f92a0",&ast.ReturnStmt/*struct*/{
                  Return: 29497,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapBinaryExpr("0xc4200f5980",&ast.BinaryExpr/*struct*/{
                      X: p.StrmapSelectorExpr("0xc4200f9260",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIdent("0xc4200f9220",&ast.Ident/*struct*/{
                          NamePos: 29504,
                          Name: "d",
                          Obj: p.PtrmapObject("0xc4200f69b0"),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200f9240",&ast.Ident/*struct*/{
                          NamePos: 29506,
                          Name: "Rparen",
                        }/*struct*/),
                      }/*struct*/),
                      OpPos: 29513,
                      Op: token.Token(12)/*+*/,
                      Y: p.StrmapBasicLit("0xc4200f9280",&ast.BasicLit/*struct*/{
                        ValuePos: 29515,
                        Kind: token.Token(5)/*INT*/,
                        Value: "1",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 29518,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200f9380",&ast.ReturnStmt/*struct*/{
            Return: 29521,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200fc500",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200f9360",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapIndexExpr("0xc4200f59e0",&ast.IndexExpr/*struct*/{
                    X: p.StrmapSelectorExpr("0xc4200f9300",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200f92c0",&ast.Ident/*struct*/{
                        NamePos: 29528,
                        Name: "d",
                        Obj: p.PtrmapObject("0xc4200f69b0"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200f92e0",&ast.Ident/*struct*/{
                        NamePos: 29530,
                        Name: "Specs",
                      }/*struct*/),
                    }/*struct*/),
                    Lbrack: 29535,
                    Index: p.StrmapBasicLit("0xc4200f9320",&ast.BasicLit/*struct*/{
                      ValuePos: 29536,
                      Kind: token.Token(5)/*INT*/,
                      Value: "0",
                    }/*struct*/),
                    Rbrack: 29537,
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f9340",&ast.Ident/*struct*/{
                    NamePos: 29539,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 29542,
                Ellipsis: 0,
                Rparen: 29543,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 29545,
      }/*struct*/),
    }/*struct*/),/* slice_item: 174*/175: p.StrmapFuncDecl("0xc4200f5c50",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f5aa0",&ast.FieldList/*struct*/{
        Opening: 29552,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fc540",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f93e0",&ast.Ident/*struct*/{
                NamePos: 29553,
                Name: "d",
                Obj: p.StrmapObject("0xc4200f6a00",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "d",
                  Decl: p.PtrmapField("0xc4200fc540"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f9420",&ast.StarExpr/*struct*/{
              Star: 29555,
              X: p.StrmapIdent("0xc4200f9400",&ast.Ident/*struct*/{
                NamePos: 29556,
                Name: "FuncDecl",
                Obj: p.PtrmapObject("0xc4200f6640"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 29564,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f9440",&ast.Ident/*struct*/{
        NamePos: 29566,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f9700",&ast.FuncType/*struct*/{
        Func: 29547,
        Params: p.StrmapFieldList("0xc4200f5ad0",&ast.FieldList/*struct*/{
          Opening: 29569,
          Closing: 29570,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200f5b00",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fc580",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f94a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f9460",&ast.Ident/*struct*/{
                  NamePos: 29572,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f9480",&ast.Ident/*struct*/{
                  NamePos: 29578,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f5c20",&ast.BlockStmt/*struct*/{
        Lbrace: 29582,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200fc600",&ast.IfStmt/*struct*/{
            If: 29585,
            Cond: p.StrmapBinaryExpr("0xc4200f5b90",&ast.BinaryExpr/*struct*/{
              X: p.StrmapSelectorExpr("0xc4200f9500",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f94c0",&ast.Ident/*struct*/{
                  NamePos: 29588,
                  Name: "d",
                  Obj: p.PtrmapObject("0xc4200f6a00"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f94e0",&ast.Ident/*struct*/{
                  NamePos: 29590,
                  Name: "Body",
                }/*struct*/),
              }/*struct*/),
              OpPos: 29595,
              Op: token.Token(44)/*!=*/,
              Y: p.StrmapIdent("0xc4200f9520",&ast.Ident/*struct*/{
                NamePos: 29598,
                Name: "nil",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200f5bf0",&ast.BlockStmt/*struct*/{
              Lbrace: 29602,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4200f95e0",&ast.ReturnStmt/*struct*/{
                  Return: 29606,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200fc5c0",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc4200f95c0",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapSelectorExpr("0xc4200f9580",&ast.SelectorExpr/*struct*/{
                          X: p.StrmapIdent("0xc4200f9540",&ast.Ident/*struct*/{
                            NamePos: 29613,
                            Name: "d",
                            Obj: p.PtrmapObject("0xc4200f6a00"),
                          }/*struct*/),
                          Sel: p.StrmapIdent("0xc4200f9560",&ast.Ident/*struct*/{
                            NamePos: 29615,
                            Name: "Body",
                          }/*struct*/),
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc4200f95a0",&ast.Ident/*struct*/{
                          NamePos: 29620,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 29623,
                      Ellipsis: 0,
                      Rparen: 29624,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 29627,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc4200f96a0",&ast.ReturnStmt/*struct*/{
            Return: 29630,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200fc640",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc4200f9680",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc4200f9640",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f9600",&ast.Ident/*struct*/{
                      NamePos: 29637,
                      Name: "d",
                      Obj: p.PtrmapObject("0xc4200f6a00"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f9620",&ast.Ident/*struct*/{
                      NamePos: 29639,
                      Name: "Type",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc4200f9660",&ast.Ident/*struct*/{
                    NamePos: 29644,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 29647,
                Ellipsis: 0,
                Rparen: 29648,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 29650,
      }/*struct*/),
    }/*struct*/),/* slice_item: 175*/176: p.StrmapFuncDecl("0xc4200f5d70",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f5cb0",&ast.FieldList/*struct*/{
        Opening: 29741,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fc6c0",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f9740",&ast.StarExpr/*struct*/{
              Star: 29742,
              X: p.StrmapIdent("0xc4200f9720",&ast.Ident/*struct*/{
                NamePos: 29743,
                Name: "BadDecl",
                Obj: p.PtrmapObject("0xc4200f6280"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 29750,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f9760",&ast.Ident/*struct*/{
        NamePos: 29752,
        Name: "declNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f9780",&ast.FuncType/*struct*/{
        Func: 29736,
        Params: p.StrmapFieldList("0xc4200f5ce0",&ast.FieldList/*struct*/{
          Opening: 29760,
          Closing: 29761,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f5d40",&ast.BlockStmt/*struct*/{
        Lbrace: 29764,
        Rbrace: 29765,
      }/*struct*/),
    }/*struct*/),/* slice_item: 176*/177: p.StrmapFuncDecl("0xc4200f5e90",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f5dd0",&ast.FieldList/*struct*/{
        Opening: 29772,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fc700",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f97c0",&ast.StarExpr/*struct*/{
              Star: 29773,
              X: p.StrmapIdent("0xc4200f97a0",&ast.Ident/*struct*/{
                NamePos: 29774,
                Name: "GenDecl",
                Obj: p.PtrmapObject("0xc4200f6410"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 29781,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f97e0",&ast.Ident/*struct*/{
        NamePos: 29783,
        Name: "declNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f9800",&ast.FuncType/*struct*/{
        Func: 29767,
        Params: p.StrmapFieldList("0xc4200f5e00",&ast.FieldList/*struct*/{
          Opening: 29791,
          Closing: 29792,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f5e60",&ast.BlockStmt/*struct*/{
        Lbrace: 29795,
        Rbrace: 29796,
      }/*struct*/),
    }/*struct*/),/* slice_item: 177*/178: p.StrmapFuncDecl("0xc4200fe000",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200f5ef0",&ast.FieldList/*struct*/{
        Opening: 29803,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fc740",&ast.Field/*struct*/{
            Type: p.StrmapStarExpr("0xc4200f9840",&ast.StarExpr/*struct*/{
              Star: 29804,
              X: p.StrmapIdent("0xc4200f9820",&ast.Ident/*struct*/{
                NamePos: 29805,
                Name: "FuncDecl",
                Obj: p.PtrmapObject("0xc4200f6640"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 29813,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f9860",&ast.Ident/*struct*/{
        NamePos: 29815,
        Name: "declNode",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f98c0",&ast.FuncType/*struct*/{
        Func: 29798,
        Params: p.StrmapFieldList("0xc4200f5f20",&ast.FieldList/*struct*/{
          Opening: 29823,
          Closing: 29824,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200f5f80",&ast.BlockStmt/*struct*/{
        Lbrace: 29826,
        Rbrace: 29827,
      }/*struct*/),
    }/*struct*/),/* slice_item: 178*/179: p.StrmapGenDecl("0xc4200fc9c0",&ast.GenDecl/*struct*/{
      TokPos: 30163,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200fe030",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4200f98e0",&ast.Ident/*struct*/{
            NamePos: 30168,
            Name: "File",
            Obj: p.StrmapObject("0xc4200f6b40",&ast.Object/*struct*/{
              Kind: "type",
              Name: "File",
              Decl: p.PtrmapTypeSpec("0xc4200fe030"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4200f9c60",&ast.StructType/*struct*/{
            Struct: 30173,
            Fields: p.StrmapFieldList("0xc4200fe210",&ast.FieldList/*struct*/{
              Opening: 30180,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200fc780",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f9900",&ast.Ident/*struct*/{
                      NamePos: 30183,
                      Name: "Doc",
                      Obj: p.StrmapObject("0xc4200f6b90",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Doc",
                        Decl: p.PtrmapField("0xc4200fc780"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f9940",&ast.StarExpr/*struct*/{
                    Star: 30194,
                    X: p.StrmapIdent("0xc4200f9920",&ast.Ident/*struct*/{
                      NamePos: 30195,
                      Name: "CommentGroup",
                      Obj: p.PtrmapObject("0xc420050a50"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200fc7c0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f9960",&ast.Ident/*struct*/{
                      NamePos: 30247,
                      Name: "Package",
                      Obj: p.StrmapObject("0xc4200f6be0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Package",
                        Decl: p.PtrmapField("0xc4200fc7c0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapSelectorExpr("0xc4200f99c0",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4200f9980",&ast.Ident/*struct*/{
                      NamePos: 30258,
                      Name: "token",
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4200f99a0",&ast.Ident/*struct*/{
                      NamePos: 30264,
                      Name: "Pos",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200fc800",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f9a00",&ast.Ident/*struct*/{
                      NamePos: 30308,
                      Name: "Name",
                      Obj: p.StrmapObject("0xc4200f6c30",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Name",
                        Decl: p.PtrmapField("0xc4200fc800"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f9a40",&ast.StarExpr/*struct*/{
                    Star: 30319,
                    X: p.StrmapIdent("0xc4200f9a20",&ast.Ident/*struct*/{
                      NamePos: 30320,
                      Name: "Ident",
                      Obj: p.PtrmapObject("0xc420051c70"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200fc840",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f9a80",&ast.Ident/*struct*/{
                      NamePos: 30352,
                      Name: "Decls",
                      Obj: p.StrmapObject("0xc4200f6c80",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Decls",
                        Decl: p.PtrmapField("0xc4200fc840"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200fe090",&ast.ArrayType/*struct*/{
                    Lbrack: 30363,
                    Elt: p.StrmapIdent("0xc4200f9aa0",&ast.Ident/*struct*/{
                      NamePos: 30365,
                      Name: "Decl",
                      Obj: p.PtrmapObject("0xc420050780"),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/4: p.StrmapField("0xc4200fc880",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f9ac0",&ast.Ident/*struct*/{
                      NamePos: 30414,
                      Name: "Scope",
                      Obj: p.StrmapObject("0xc4200f6cd0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Scope",
                        Decl: p.PtrmapField("0xc4200fc880"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc4200f9b00",&ast.StarExpr/*struct*/{
                    Star: 30425,
                    X: p.StrmapIdent("0xc4200f9ae0",&ast.Ident/*struct*/{
                      NamePos: 30426,
                      Name: "Scope",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 4*/5: p.StrmapField("0xc4200fc900",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f9b20",&ast.Ident/*struct*/{
                      NamePos: 30476,
                      Name: "Imports",
                      Obj: p.StrmapObject("0xc4200f6d20",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Imports",
                        Decl: p.PtrmapField("0xc4200fc900"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200fe120",&ast.ArrayType/*struct*/{
                    Lbrack: 30487,
                    Elt: p.StrmapStarExpr("0xc4200f9b60",&ast.StarExpr/*struct*/{
                      Star: 30489,
                      X: p.StrmapIdent("0xc4200f9b40",&ast.Ident/*struct*/{
                        NamePos: 30490,
                        Name: "ImportSpec",
                        Obj: p.PtrmapObject("0xc4200e39a0"),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 5*/6: p.StrmapField("0xc4200fc940",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f9ba0",&ast.Ident/*struct*/{
                      NamePos: 30528,
                      Name: "Unresolved",
                      Obj: p.StrmapObject("0xc4200f6d70",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Unresolved",
                        Decl: p.PtrmapField("0xc4200fc940"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200fe150",&ast.ArrayType/*struct*/{
                    Lbrack: 30539,
                    Elt: p.StrmapStarExpr("0xc4200f9be0",&ast.StarExpr/*struct*/{
                      Star: 30541,
                      X: p.StrmapIdent("0xc4200f9bc0",&ast.Ident/*struct*/{
                        NamePos: 30542,
                        Name: "Ident",
                        Obj: p.PtrmapObject("0xc420051c70"),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 6*/7: p.StrmapField("0xc4200fc980",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4200f9c00",&ast.Ident/*struct*/{
                      NamePos: 30595,
                      Name: "Comments",
                      Obj: p.StrmapObject("0xc4200f6dc0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Comments",
                        Decl: p.PtrmapField("0xc4200fc980"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapArrayType("0xc4200fe1b0",&ast.ArrayType/*struct*/{
                    Lbrack: 30606,
                    Elt: p.StrmapStarExpr("0xc4200f9c40",&ast.StarExpr/*struct*/{
                      Star: 30608,
                      X: p.StrmapIdent("0xc4200f9c20",&ast.Ident/*struct*/{
                        NamePos: 30609,
                        Name: "CommentGroup",
                        Obj: p.PtrmapObject("0xc420050a50"),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 7*/}/*slice*/,
              Closing: 30665,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 179*/180: p.StrmapFuncDecl("0xc4200fe360",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200fe270",&ast.FieldList/*struct*/{
        Opening: 30673,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fca00",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f9c80",&ast.Ident/*struct*/{
                NamePos: 30674,
                Name: "f",
                Obj: p.StrmapObject("0xc4200f6e10",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "f",
                  Decl: p.PtrmapField("0xc4200fca00"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f9cc0",&ast.StarExpr/*struct*/{
              Star: 30676,
              X: p.StrmapIdent("0xc4200f9ca0",&ast.Ident/*struct*/{
                NamePos: 30677,
                Name: "File",
                Obj: p.PtrmapObject("0xc4200f6b40"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 30681,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f9ce0",&ast.Ident/*struct*/{
        NamePos: 30683,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4200f9de0",&ast.FuncType/*struct*/{
        Func: 30668,
        Params: p.StrmapFieldList("0xc4200fe2a0",&ast.FieldList/*struct*/{
          Opening: 30686,
          Closing: 30687,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200fe2d0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fca40",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f9d40",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f9d00",&ast.Ident/*struct*/{
                  NamePos: 30689,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f9d20",&ast.Ident/*struct*/{
                  NamePos: 30695,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200fe330",&ast.BlockStmt/*struct*/{
        Lbrace: 30699,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4200f9dc0",&ast.ReturnStmt/*struct*/{
            Return: 30701,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc4200f9da0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f9d60",&ast.Ident/*struct*/{
                  NamePos: 30708,
                  Name: "f",
                  Obj: p.PtrmapObject("0xc4200f6e10"),
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f9d80",&ast.Ident/*struct*/{
                  NamePos: 30710,
                  Name: "Package",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 30718,
      }/*struct*/),
    }/*struct*/),/* slice_item: 180*/181: p.StrmapFuncDecl("0xc4200fe600",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200fe3c0",&ast.FieldList/*struct*/{
        Opening: 30725,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fca80",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4200f9e00",&ast.Ident/*struct*/{
                NamePos: 30726,
                Name: "f",
                Obj: p.StrmapObject("0xc4200f6e60",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "f",
                  Decl: p.PtrmapField("0xc4200fca80"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4200f9e40",&ast.StarExpr/*struct*/{
              Star: 30728,
              X: p.StrmapIdent("0xc4200f9e20",&ast.Ident/*struct*/{
                NamePos: 30729,
                Name: "File",
                Obj: p.PtrmapObject("0xc4200f6b40"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 30733,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4200f9e60",&ast.Ident/*struct*/{
        NamePos: 30735,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4201021a0",&ast.FuncType/*struct*/{
        Func: 30720,
        Params: p.StrmapFieldList("0xc4200fe3f0",&ast.FieldList/*struct*/{
          Opening: 30738,
          Closing: 30739,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200fe420",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fcac0",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4200f9ec0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4200f9e80",&ast.Ident/*struct*/{
                  NamePos: 30741,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4200f9ea0",&ast.Ident/*struct*/{
                  NamePos: 30747,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200fe5a0",&ast.BlockStmt/*struct*/{
        Lbrace: 30751,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapIfStmt("0xc4200fcbc0",&ast.IfStmt/*struct*/{
            If: 30754,
            Init: p.StrmapAssignStmt("0xc4200fcb40",&ast.AssignStmt/*struct*/{
              Lhs: []ast.Expr /*Slice*/{
                0: p.StrmapIdent("0xc4200f9ee0",&ast.Ident/*struct*/{
                  NamePos: 30757,
                  Name: "n",
                  Obj: p.StrmapObject("0xc4200f6eb0",&ast.Object/*struct*/{
                    Kind: "var",
                    Name: "n",
                    Decl: p.PtrmapAssignStmt("0xc4200fcb40"),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              TokPos: 30759,
              Tok: token.Token(47)/*:=*/,
              Rhs: []ast.Expr /*Slice*/{
                0: p.StrmapCallExpr("0xc4200fcb00",&ast.CallExpr/*struct*/{
                  Fun: p.StrmapIdent("0xc4200f9f00",&ast.Ident/*struct*/{
                    NamePos: 30762,
                    Name: "len",
                  }/*struct*/),
                  Lparen: 30765,
                  Args: []ast.Expr /*Slice*/{
                    0: p.StrmapSelectorExpr("0xc4200f9f60",&ast.SelectorExpr/*struct*/{
                      X: p.StrmapIdent("0xc4200f9f20",&ast.Ident/*struct*/{
                        NamePos: 30766,
                        Name: "f",
                        Obj: p.PtrmapObject("0xc4200f6e60"),
                      }/*struct*/),
                      Sel: p.StrmapIdent("0xc4200f9f40",&ast.Ident/*struct*/{
                        NamePos: 30768,
                        Name: "Decls",
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Ellipsis: 0,
                  Rparen: 30773,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
            }/*struct*/),
            Cond: p.StrmapBinaryExpr("0xc4200fe4b0",&ast.BinaryExpr/*struct*/{
              X: p.StrmapIdent("0xc4200f9f80",&ast.Ident/*struct*/{
                NamePos: 30776,
                Name: "n",
                Obj: p.PtrmapObject("0xc4200f6eb0"),
              }/*struct*/),
              OpPos: 30778,
              Op: token.Token(41)/*>*/,
              Y: p.StrmapBasicLit("0xc4200f9fa0",&ast.BasicLit/*struct*/{
                ValuePos: 30780,
                Kind: token.Token(5)/*INT*/,
                Value: "0",
              }/*struct*/),
            }/*struct*/),
            Body: p.StrmapBlockStmt("0xc4200fe570",&ast.BlockStmt/*struct*/{
              Lbrace: 30782,
              List: []ast.Stmt /*Slice*/{
                0: p.StrmapReturnStmt("0xc4201020a0",&ast.ReturnStmt/*struct*/{
                  Return: 30786,
                  Results: []ast.Expr /*Slice*/{
                    0: p.StrmapCallExpr("0xc4200fcb80",&ast.CallExpr/*struct*/{
                      Fun: p.StrmapSelectorExpr("0xc420102080",&ast.SelectorExpr/*struct*/{
                        X: p.StrmapIndexExpr("0xc4200fe540",&ast.IndexExpr/*struct*/{
                          X: p.StrmapSelectorExpr("0xc420102000",&ast.SelectorExpr/*struct*/{
                            X: p.StrmapIdent("0xc4200f9fc0",&ast.Ident/*struct*/{
                              NamePos: 30793,
                              Name: "f",
                              Obj: p.PtrmapObject("0xc4200f6e60"),
                            }/*struct*/),
                            Sel: p.StrmapIdent("0xc4200f9fe0",&ast.Ident/*struct*/{
                              NamePos: 30795,
                              Name: "Decls",
                            }/*struct*/),
                          }/*struct*/),
                          Lbrack: 30800,
                          Index: p.StrmapBinaryExpr("0xc4200fe510",&ast.BinaryExpr/*struct*/{
                            X: p.StrmapIdent("0xc420102020",&ast.Ident/*struct*/{
                              NamePos: 30801,
                              Name: "n",
                              Obj: p.PtrmapObject("0xc4200f6eb0"),
                            }/*struct*/),
                            OpPos: 30802,
                            Op: token.Token(13)/*-*/,
                            Y: p.StrmapBasicLit("0xc420102040",&ast.BasicLit/*struct*/{
                              ValuePos: 30803,
                              Kind: token.Token(5)/*INT*/,
                              Value: "1",
                            }/*struct*/),
                          }/*struct*/),
                          Rbrack: 30804,
                        }/*struct*/),
                        Sel: p.StrmapIdent("0xc420102060",&ast.Ident/*struct*/{
                          NamePos: 30806,
                          Name: "End",
                        }/*struct*/),
                      }/*struct*/),
                      Lparen: 30809,
                      Ellipsis: 0,
                      Rparen: 30810,
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                }/*struct*/),/* slice_item: 0*/}/*slice*/,
              Rbrace: 30813,
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/1: p.StrmapReturnStmt("0xc420102160",&ast.ReturnStmt/*struct*/{
            Return: 30816,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapCallExpr("0xc4200fcc00",&ast.CallExpr/*struct*/{
                Fun: p.StrmapSelectorExpr("0xc420102140",&ast.SelectorExpr/*struct*/{
                  X: p.StrmapSelectorExpr("0xc420102100",&ast.SelectorExpr/*struct*/{
                    X: p.StrmapIdent("0xc4201020c0",&ast.Ident/*struct*/{
                      NamePos: 30823,
                      Name: "f",
                      Obj: p.PtrmapObject("0xc4200f6e60"),
                    }/*struct*/),
                    Sel: p.StrmapIdent("0xc4201020e0",&ast.Ident/*struct*/{
                      NamePos: 30825,
                      Name: "Name",
                    }/*struct*/),
                  }/*struct*/),
                  Sel: p.StrmapIdent("0xc420102120",&ast.Ident/*struct*/{
                    NamePos: 30830,
                    Name: "End",
                  }/*struct*/),
                }/*struct*/),
                Lparen: 30833,
                Ellipsis: 0,
                Rparen: 30834,
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 1*/}/*slice*/,
        Rbrace: 30836,
      }/*struct*/),
    }/*struct*/),/* slice_item: 181*/182: p.StrmapGenDecl("0xc4200fcd80",&ast.GenDecl/*struct*/{
      TokPos: 30932,
      Tok: token.Token(84)/*type*/,
      Lparen: 0,
      Specs: []ast.Spec /*Slice*/{
        0: p.StrmapTypeSpec("0xc4200fe630",&ast.TypeSpec/*struct*/{
          Name: p.StrmapIdent("0xc4201021c0",&ast.Ident/*struct*/{
            NamePos: 30937,
            Name: "Package",
            Obj: p.StrmapObject("0xc4200f6f00",&ast.Object/*struct*/{
              Kind: "type",
              Name: "Package",
              Decl: p.PtrmapTypeSpec("0xc4200fe630"),
            }/*struct*/),
          }/*struct*/),
          Type: p.StrmapStructType("0xc4201023c0",&ast.StructType/*struct*/{
            Struct: 30945,
            Fields: p.StrmapFieldList("0xc4200fe720",&ast.FieldList/*struct*/{
              Opening: 30952,
              List: []*ast.Field /*Slice*/{
                0: p.StrmapField("0xc4200fcc80",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc4201021e0",&ast.Ident/*struct*/{
                      NamePos: 30955,
                      Name: "Name",
                      Obj: p.StrmapObject("0xc4200f6f50",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Name",
                        Decl: p.PtrmapField("0xc4200fcc80"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapIdent("0xc420102200",&ast.Ident/*struct*/{
                    NamePos: 30963,
                    Name: "string",
                  }/*struct*/),
                }/*struct*/),/* slice_item: 0*/1: p.StrmapField("0xc4200fccc0",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc420102220",&ast.Ident/*struct*/{
                      NamePos: 30999,
                      Name: "Scope",
                      Obj: p.StrmapObject("0xc4200f6fa0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Scope",
                        Decl: p.PtrmapField("0xc4200fccc0"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapStarExpr("0xc420102260",&ast.StarExpr/*struct*/{
                    Star: 31007,
                    X: p.StrmapIdent("0xc420102240",&ast.Ident/*struct*/{
                      NamePos: 31008,
                      Name: "Scope",
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 1*/2: p.StrmapField("0xc4200fcd00",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc420102280",&ast.Ident/*struct*/{
                      NamePos: 31061,
                      Name: "Imports",
                      Obj: p.StrmapObject("0xc4200f6ff0",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Imports",
                        Decl: p.PtrmapField("0xc4200fcd00"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapMapType("0xc4200fe690",&ast.MapType/*struct*/{
                    Map: 31069,
                    Key: p.StrmapIdent("0xc4201022a0",&ast.Ident/*struct*/{
                      NamePos: 31073,
                      Name: "string",
                    }/*struct*/),
                    Value: p.StrmapStarExpr("0xc4201022e0",&ast.StarExpr/*struct*/{
                      Star: 31080,
                      X: p.StrmapIdent("0xc4201022c0",&ast.Ident/*struct*/{
                        NamePos: 31081,
                        Name: "Object",
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 2*/3: p.StrmapField("0xc4200fcd40",&ast.Field/*struct*/{
                  Names: []*ast.Ident /*Slice*/{
                    0: p.StrmapIdent("0xc420102320",&ast.Ident/*struct*/{
                      NamePos: 31128,
                      Name: "Files",
                      Obj: p.StrmapObject("0xc4200f7040",&ast.Object/*struct*/{
                        Kind: "var",
                        Name: "Files",
                        Decl: p.PtrmapField("0xc4200fcd40"),
                      }/*struct*/),
                    }/*struct*/),/* slice_item: 0*/}/*slice*/,
                  Type: p.StrmapMapType("0xc4200fe6f0",&ast.MapType/*struct*/{
                    Map: 31136,
                    Key: p.StrmapIdent("0xc420102340",&ast.Ident/*struct*/{
                      NamePos: 31140,
                      Name: "string",
                    }/*struct*/),
                    Value: p.StrmapStarExpr("0xc420102380",&ast.StarExpr/*struct*/{
                      Star: 31147,
                      X: p.StrmapIdent("0xc420102360",&ast.Ident/*struct*/{
                        NamePos: 31148,
                        Name: "File",
                        Obj: p.PtrmapObject("0xc4200f6b40"),
                      }/*struct*/),
                    }/*struct*/),
                  }/*struct*/),
                }/*struct*/),/* slice_item: 3*/}/*slice*/,
              Closing: 31186,
            }/*struct*/),
            Incomplete: false,
          }/*struct*/),
        }/*struct*/),/* slice_item: 0*/}/*slice*/,
      Rparen: 0,
    }/*struct*/),/* slice_item: 182*/183: p.StrmapFuncDecl("0xc4200fe870",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200fe780",&ast.FieldList/*struct*/{
        Opening: 31194,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fcdc0",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc4201023e0",&ast.Ident/*struct*/{
                NamePos: 31195,
                Name: "p",
                Obj: p.StrmapObject("0xc4200f7090",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "p",
                  Decl: p.PtrmapField("0xc4200fcdc0"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc420102420",&ast.StarExpr/*struct*/{
              Star: 31197,
              X: p.StrmapIdent("0xc420102400",&ast.Ident/*struct*/{
                NamePos: 31198,
                Name: "Package",
                Obj: p.PtrmapObject("0xc4200f6f00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 31205,
      }/*struct*/),
      Name: p.StrmapIdent("0xc420102440",&ast.Ident/*struct*/{
        NamePos: 31207,
        Name: "Pos",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc420102540",&ast.FuncType/*struct*/{
        Func: 31189,
        Params: p.StrmapFieldList("0xc4200fe7b0",&ast.FieldList/*struct*/{
          Opening: 31210,
          Closing: 31211,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200fe7e0",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fce00",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc4201024a0",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc420102460",&ast.Ident/*struct*/{
                  NamePos: 31213,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc420102480",&ast.Ident/*struct*/{
                  NamePos: 31219,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200fe840",&ast.BlockStmt/*struct*/{
        Lbrace: 31223,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc420102520",&ast.ReturnStmt/*struct*/{
            Return: 31225,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc420102500",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4201024c0",&ast.Ident/*struct*/{
                  NamePos: 31232,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc4201024e0",&ast.Ident/*struct*/{
                  NamePos: 31238,
                  Name: "NoPos",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 31244,
      }/*struct*/),
    }/*struct*/),/* slice_item: 183*/184: p.StrmapFuncDecl("0xc4200fe9c0",&ast.FuncDecl/*struct*/{
      Recv: p.StrmapFieldList("0xc4200fe8d0",&ast.FieldList/*struct*/{
        Opening: 31251,
        List: []*ast.Field /*Slice*/{
          0: p.StrmapField("0xc4200fce40",&ast.Field/*struct*/{
            Names: []*ast.Ident /*Slice*/{
              0: p.StrmapIdent("0xc420102560",&ast.Ident/*struct*/{
                NamePos: 31252,
                Name: "p",
                Obj: p.StrmapObject("0xc4200f70e0",&ast.Object/*struct*/{
                  Kind: "var",
                  Name: "p",
                  Decl: p.PtrmapField("0xc4200fce40"),
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
            Type: p.StrmapStarExpr("0xc4201025a0",&ast.StarExpr/*struct*/{
              Star: 31254,
              X: p.StrmapIdent("0xc420102580",&ast.Ident/*struct*/{
                NamePos: 31255,
                Name: "Package",
                Obj: p.PtrmapObject("0xc4200f6f00"),
              }/*struct*/),
            }/*struct*/),
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Closing: 31262,
      }/*struct*/),
      Name: p.StrmapIdent("0xc4201025c0",&ast.Ident/*struct*/{
        NamePos: 31264,
        Name: "End",
      }/*struct*/),
      Type: p.StrmapFuncType("0xc4201026c0",&ast.FuncType/*struct*/{
        Func: 31246,
        Params: p.StrmapFieldList("0xc4200fe900",&ast.FieldList/*struct*/{
          Opening: 31267,
          Closing: 31268,
        }/*struct*/),
        Results: p.StrmapFieldList("0xc4200fe930",&ast.FieldList/*struct*/{
          Opening: 0,
          List: []*ast.Field /*Slice*/{
            0: p.StrmapField("0xc4200fce80",&ast.Field/*struct*/{
              Type: p.StrmapSelectorExpr("0xc420102620",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc4201025e0",&ast.Ident/*struct*/{
                  NamePos: 31270,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc420102600",&ast.Ident/*struct*/{
                  NamePos: 31276,
                  Name: "Pos",
                }/*struct*/),
              }/*struct*/),
            }/*struct*/),/* slice_item: 0*/}/*slice*/,
          Closing: 0,
        }/*struct*/),
      }/*struct*/),
      Body: p.StrmapBlockStmt("0xc4200fe990",&ast.BlockStmt/*struct*/{
        Lbrace: 31280,
        List: []ast.Stmt /*Slice*/{
          0: p.StrmapReturnStmt("0xc4201026a0",&ast.ReturnStmt/*struct*/{
            Return: 31282,
            Results: []ast.Expr /*Slice*/{
              0: p.StrmapSelectorExpr("0xc420102680",&ast.SelectorExpr/*struct*/{
                X: p.StrmapIdent("0xc420102640",&ast.Ident/*struct*/{
                  NamePos: 31289,
                  Name: "token",
                }/*struct*/),
                Sel: p.StrmapIdent("0xc420102660",&ast.Ident/*struct*/{
                  NamePos: 31295,
                  Name: "NoPos",
                }/*struct*/),
              }/*struct*/),/* slice_item: 0*/}/*slice*/,
          }/*struct*/),/* slice_item: 0*/}/*slice*/,
        Rbrace: 31301,
      }/*struct*/),
    }/*struct*/),/* slice_item: 184*/}/*slice*/,
  Scope: p.StrmapScope("0xc42000e850",&ast.Scope/*struct*/{
    Objects: map[string]*ast.Object /*Map*/ {
      "ForStmt"/*kv*/: p.PtrmapObject("0xc4200e25f0"),
      "ExprStmt"/*kv*/: p.PtrmapObject("0xc4200c9220"),
      "AssignStmt"/*kv*/: p.PtrmapObject("0xc4200c9590"),
      "BlockStmt"/*kv*/: p.PtrmapObject("0xc4200c9b30"),
      "SwitchStmt"/*kv*/: p.PtrmapObject("0xc4200e2050"),
      "Spec"/*kv*/: p.PtrmapObject("0xc4200e3900"),
      "InterfaceType"/*kv*/: p.PtrmapObject("0xc4200c1900"),
      "TypeAssertExpr"/*kv*/: p.PtrmapObject("0xc4200c0a00"),
      "StarExpr"/*kv*/: p.PtrmapObject("0xc4200c0e60"),
      "ChanDir"/*kv*/: p.PtrmapObject("0xc4200c1360"),
      "StructType"/*kv*/: p.PtrmapObject("0xc4200c1630"),
      "MapType"/*kv*/: p.PtrmapObject("0xc4200c1a40"),
      "BadStmt"/*kv*/: p.PtrmapObject("0xc4200c8e10"),
      "KeyValueExpr"/*kv*/: p.PtrmapObject("0xc4200c1220"),
      "RECV"/*kv*/: p.PtrmapObject("0xc4200c14a0"),
      "ArrayType"/*kv*/: p.PtrmapObject("0xc4200c14f0"),
      "Ellipsis"/*kv*/: p.PtrmapObject("0xc420051db0"),
      "SEND"/*kv*/: p.PtrmapObject("0xc4200c1400"),
      "IsExported"/*kv*/: p.PtrmapObject("0xc4200c8cd0"),
      "isWhitespace"/*kv*/: p.PtrmapObject("0xc420050be0"),
      "Ident"/*kv*/: p.PtrmapObject("0xc420051c70"),
      "ChanType"/*kv*/: p.PtrmapObject("0xc4200c1b80"),
      "GoStmt"/*kv*/: p.PtrmapObject("0xc4200c9720"),
      "IfStmt"/*kv*/: p.PtrmapObject("0xc4200c9c70"),
      "SelectStmt"/*kv*/: p.PtrmapObject("0xc4200e2500"),
      "GenDecl"/*kv*/: p.PtrmapObject("0xc4200f6410"),
      "FuncType"/*kv*/: p.PtrmapObject("0xc4200c1770"),
      "LabeledStmt"/*kv*/: p.PtrmapObject("0xc4200c90e0"),
      "SendStmt"/*kv*/: p.PtrmapObject("0xc4200c92c0"),
      "EmptyStmt"/*kv*/: p.PtrmapObject("0xc4200c8ff0"),
      "CaseClause"/*kv*/: p.PtrmapObject("0xc4200c9ea0"),
      "Expr"/*kv*/: p.PtrmapObject("0xc420050640"),
      "Comment"/*kv*/: p.PtrmapObject("0xc4200508c0"),
      "Field"/*kv*/: p.PtrmapObject("0xc420051450"),
      "BinaryExpr"/*kv*/: p.PtrmapObject("0xc4200c1090"),
      "TypeSpec"/*kv*/: p.PtrmapObject("0xc4200e3e00"),
      "Decl"/*kv*/: p.PtrmapObject("0xc420050780"),
      "BasicLit"/*kv*/: p.PtrmapObject("0xc420051ea0"),
      "CallExpr"/*kv*/: p.PtrmapObject("0xc4200c0be0"),
      "TypeSwitchStmt"/*kv*/: p.PtrmapObject("0xc4200e21e0"),
      "SelectorExpr"/*kv*/: p.PtrmapObject("0xc4200c0460"),
      "IndexExpr"/*kv*/: p.PtrmapObject("0xc4200c05a0"),
      "UnaryExpr"/*kv*/: p.PtrmapObject("0xc4200c0f50"),
      "RangeStmt"/*kv*/: p.PtrmapObject("0xc4200e27d0"),
      "ValueSpec"/*kv*/: p.PtrmapObject("0xc4200e3bd0"),
      "File"/*kv*/: p.PtrmapObject("0xc4200f6b40"),
      "Node"/*kv*/: p.PtrmapObject("0xc420050550"),
      "BadExpr"/*kv*/: p.PtrmapObject("0xc420051b80"),
      "CompositeLit"/*kv*/: p.PtrmapObject("0xc4200c0140"),
      "Package"/*kv*/: p.PtrmapObject("0xc4200f6f00"),
      "NewIdent"/*kv*/: p.PtrmapObject("0xc4200c8b90"),
      "DeferStmt"/*kv*/: p.PtrmapObject("0xc4200c9810"),
      "BranchStmt"/*kv*/: p.PtrmapObject("0xc4200c99f0"),
      "CommentGroup"/*kv*/: p.PtrmapObject("0xc420050a50"),
      "FieldList"/*kv*/: p.PtrmapObject("0xc420051720"),
      "ParenExpr"/*kv*/: p.PtrmapObject("0xc4200c02d0"),
      "ImportSpec"/*kv*/: p.PtrmapObject("0xc4200e39a0"),
      "BadDecl"/*kv*/: p.PtrmapObject("0xc4200f6280"),
      "FuncDecl"/*kv*/: p.PtrmapObject("0xc4200f6640"),
      "Stmt"/*kv*/: p.PtrmapObject("0xc4200506e0"),
      "FuncLit"/*kv*/: p.PtrmapObject("0xc4200c0050"),
      "ReturnStmt"/*kv*/: p.PtrmapObject("0xc4200c9900"),
      "IncDecStmt"/*kv*/: p.PtrmapObject("0xc4200c9450"),
      "CommClause"/*kv*/: p.PtrmapObject("0xc4200e2370"),
      "stripTrailingWhitespace"/*kv*/: p.PtrmapObject("0xc420050dc0"),
      "SliceExpr"/*kv*/: p.PtrmapObject("0xc4200c0780"),
      "DeclStmt"/*kv*/: p.PtrmapObject("0xc4200c8f00"),
    } /*hash end*/,
  }/*struct*/),
  Imports: []*ast.ImportSpec /*Slice*/{
    0: p.PtrmapImportSpec("0xc42000a5d0"),/* slice_item: 0*/1: p.PtrmapImportSpec("0xc42000a600"),/* slice_item: 1*/2: p.PtrmapImportSpec("0xc42000a630"),/* slice_item: 2*/3: p.PtrmapImportSpec("0xc42000a660"),/* slice_item: 3*/}/*slice*/,
  Unresolved: []*ast.Ident /*Slice*/{
    0: p.PtrmapIdent("0xc4200103c0"),/* slice_item: 0*/1: p.PtrmapIdent("0xc420010460"),/* slice_item: 1*/2: p.PtrmapIdent("0xc420010740"),/* slice_item: 2*/3: p.PtrmapIdent("0xc4200107e0"),/* slice_item: 3*/4: p.PtrmapIdent("0xc4200108a0"),/* slice_item: 4*/5: p.PtrmapIdent("0xc420010a40"),/* slice_item: 5*/6: p.PtrmapIdent("0xc420010aa0"),/* slice_item: 6*/7: p.PtrmapIdent("0xc420010b00"),/* slice_item: 7*/8: p.PtrmapIdent("0xc420010b80"),/* slice_item: 8*/9: p.PtrmapIdent("0xc420010d60"),/* slice_item: 9*/10: p.PtrmapIdent("0xc420010f40"),/* slice_item: 10*/11: p.PtrmapIdent("0xc420011000"),/* slice_item: 11*/12: p.PtrmapIdent("0xc420011160"),/* slice_item: 12*/13: p.PtrmapIdent("0xc420011180"),/* slice_item: 13*/14: p.PtrmapIdent("0xc420011340"),/* slice_item: 14*/15: p.PtrmapIdent("0xc420011360"),/* slice_item: 15*/16: p.PtrmapIdent("0xc4200113a0"),/* slice_item: 16*/17: p.PtrmapIdent("0xc420011620"),/* slice_item: 17*/18: p.PtrmapIdent("0xc420011660"),/* slice_item: 18*/19: p.PtrmapIdent("0xc4200116e0"),/* slice_item: 19*/20: p.PtrmapIdent("0xc420011700"),/* slice_item: 20*/21: p.PtrmapIdent("0xc420011720"),/* slice_item: 21*/22: p.PtrmapIdent("0xc420011980"),/* slice_item: 22*/23: p.PtrmapIdent("0xc4200119a0"),/* slice_item: 23*/24: p.PtrmapIdent("0xc420011ba0"),/* slice_item: 24*/25: p.PtrmapIdent("0xc420011d80"),/* slice_item: 25*/26: p.PtrmapIdent("0xc420011e40"),/* slice_item: 26*/27: p.PtrmapIdent("0xc420011fe0"),/* slice_item: 27*/28: p.PtrmapIdent("0xc4200b04a0"),/* slice_item: 28*/29: p.PtrmapIdent("0xc4200b0520"),/* slice_item: 29*/30: p.PtrmapIdent("0xc4200b0960"),/* slice_item: 30*/31: p.PtrmapIdent("0xc4200b09c0"),/* slice_item: 31*/32: p.PtrmapIdent("0xc4200b0cc0"),/* slice_item: 32*/33: p.PtrmapIdent("0xc4200b0d80"),/* slice_item: 33*/34: p.PtrmapIdent("0xc4200b0fa0"),/* slice_item: 34*/35: p.PtrmapIdent("0xc4200b10a0"),/* slice_item: 35*/36: p.PtrmapIdent("0xc4200b11c0"),/* slice_item: 36*/37: p.PtrmapIdent("0xc4200b1340"),/* slice_item: 37*/38: p.PtrmapIdent("0xc4200b14e0"),/* slice_item: 38*/39: p.PtrmapIdent("0xc4200b1600"),/* slice_item: 39*/40: p.PtrmapIdent("0xc4200b17c0"),/* slice_item: 40*/41: p.PtrmapIdent("0xc4200b19a0"),/* slice_item: 41*/42: p.PtrmapIdent("0xc4200b1ac0"),/* slice_item: 42*/43: p.PtrmapIdent("0xc4200b1b40"),/* slice_item: 43*/44: p.PtrmapIdent("0xc4200b1c60"),/* slice_item: 44*/45: p.PtrmapIdent("0xc4200b1ee0"),/* slice_item: 45*/46: p.PtrmapIdent("0xc4200b1fa0"),/* slice_item: 46*/47: p.PtrmapIdent("0xc4200ba040"),/* slice_item: 47*/48: p.PtrmapIdent("0xc4200ba0a0"),/* slice_item: 48*/49: p.PtrmapIdent("0xc4200ba1a0"),/* slice_item: 49*/50: p.PtrmapIdent("0xc4200ba2c0"),/* slice_item: 50*/51: p.PtrmapIdent("0xc4200ba360"),/* slice_item: 51*/52: p.PtrmapIdent("0xc4200ba3e0"),/* slice_item: 52*/53: p.PtrmapIdent("0xc4200ba5e0"),/* slice_item: 53*/54: p.PtrmapIdent("0xc4200ba6e0"),/* slice_item: 54*/55: p.PtrmapIdent("0xc4200ba7c0"),/* slice_item: 55*/56: p.PtrmapIdent("0xc4200ba8c0"),/* slice_item: 56*/57: p.PtrmapIdent("0xc4200bab00"),/* slice_item: 57*/58: p.PtrmapIdent("0xc4200bac20"),/* slice_item: 58*/59: p.PtrmapIdent("0xc4200bad40"),/* slice_item: 59*/60: p.PtrmapIdent("0xc4200baf00"),/* slice_item: 60*/61: p.PtrmapIdent("0xc4200baf40"),/* slice_item: 61*/62: p.PtrmapIdent("0xc4200bb080"),/* slice_item: 62*/63: p.PtrmapIdent("0xc4200bb180"),/* slice_item: 63*/64: p.PtrmapIdent("0xc4200bb2c0"),/* slice_item: 64*/65: p.PtrmapIdent("0xc4200bb3e0"),/* slice_item: 65*/66: p.PtrmapIdent("0xc4200bb480"),/* slice_item: 66*/67: p.PtrmapIdent("0xc4200bb560"),/* slice_item: 67*/68: p.PtrmapIdent("0xc4200bb680"),/* slice_item: 68*/69: p.PtrmapIdent("0xc4200bb720"),/* slice_item: 69*/70: p.PtrmapIdent("0xc4200bb880"),/* slice_item: 70*/71: p.PtrmapIdent("0xc4200bb920"),/* slice_item: 71*/72: p.PtrmapIdent("0xc4200bbaa0"),/* slice_item: 72*/73: p.PtrmapIdent("0xc4200bbbe0"),/* slice_item: 73*/74: p.PtrmapIdent("0xc4200bbc60"),/* slice_item: 74*/75: p.PtrmapIdent("0xc4200bbd00"),/* slice_item: 75*/76: p.PtrmapIdent("0xc4200bbe80"),/* slice_item: 76*/77: p.PtrmapIdent("0xc4200bbfa0"),/* slice_item: 77*/78: p.PtrmapIdent("0xc4200c4060"),/* slice_item: 78*/79: p.PtrmapIdent("0xc4200c4220"),/* slice_item: 79*/80: p.PtrmapIdent("0xc4200c4320"),/* slice_item: 80*/81: p.PtrmapIdent("0xc4200c43c0"),/* slice_item: 81*/82: p.PtrmapIdent("0xc4200c4540"),/* slice_item: 82*/83: p.PtrmapIdent("0xc4200c45c0"),/* slice_item: 83*/84: p.PtrmapIdent("0xc4200c4780"),/* slice_item: 84*/85: p.PtrmapIdent("0xc4200c4900"),/* slice_item: 85*/86: p.PtrmapIdent("0xc4200c4a80"),/* slice_item: 86*/87: p.PtrmapIdent("0xc4200c4c00"),/* slice_item: 87*/88: p.PtrmapIdent("0xc4200c4d80"),/* slice_item: 88*/89: p.PtrmapIdent("0xc4200c4f40"),/* slice_item: 89*/90: p.PtrmapIdent("0xc4200c5000"),/* slice_item: 90*/91: p.PtrmapIdent("0xc4200c5220"),/* slice_item: 91*/92: p.PtrmapIdent("0xc4200c53a0"),/* slice_item: 92*/93: p.PtrmapIdent("0xc4200c5560"),/* slice_item: 93*/94: p.PtrmapIdent("0xc4200c5720"),/* slice_item: 94*/95: p.PtrmapIdent("0xc4200c58e0"),/* slice_item: 95*/96: p.PtrmapIdent("0xc4200c5aa0"),/* slice_item: 96*/97: p.PtrmapIdent("0xc4200c5c60"),/* slice_item: 97*/98: p.PtrmapIdent("0xc4200c5de0"),/* slice_item: 98*/99: p.PtrmapIdent("0xc4200c5f60"),/* slice_item: 99*/100: p.PtrmapIdent("0xc4200ca120"),/* slice_item: 100*/101: p.PtrmapIdent("0xc4200ca2e0"),/* slice_item: 101*/102: p.PtrmapIdent("0xc4200ca460"),/* slice_item: 102*/103: p.PtrmapIdent("0xc4200ca5e0"),/* slice_item: 103*/104: p.PtrmapIdent("0xc4200ca740"),/* slice_item: 104*/105: p.PtrmapIdent("0xc4200ca980"),/* slice_item: 105*/106: p.PtrmapIdent("0xc4200cab00"),/* slice_item: 106*/107: p.PtrmapIdent("0xc4200cac80"),/* slice_item: 107*/108: p.PtrmapIdent("0xc4200cae00"),/* slice_item: 108*/109: p.PtrmapIdent("0xc4200caf80"),/* slice_item: 109*/110: p.PtrmapIdent("0xc4200cafe0"),/* slice_item: 110*/111: p.PtrmapIdent("0xc4200cb040"),/* slice_item: 111*/112: p.PtrmapIdent("0xc4200cb0c0"),/* slice_item: 112*/113: p.PtrmapIdent("0xc4200cb200"),/* slice_item: 113*/114: p.PtrmapIdent("0xc4200cb2c0"),/* slice_item: 114*/115: p.PtrmapIdent("0xc4200cb500"),/* slice_item: 115*/116: p.PtrmapIdent("0xc4200cb560"),/* slice_item: 116*/117: p.PtrmapIdent("0xc4200cb5c0"),/* slice_item: 117*/118: p.PtrmapIdent("0xc4200cb640"),/* slice_item: 118*/119: p.PtrmapIdent("0xc4200cb780"),/* slice_item: 119*/120: p.PtrmapIdent("0xc4200cb940"),/* slice_item: 120*/121: p.PtrmapIdent("0xc4200cbae0"),/* slice_item: 121*/122: p.PtrmapIdent("0xc4200cbc80"),/* slice_item: 122*/123: p.PtrmapIdent("0xc4200cbe40"),/* slice_item: 123*/124: p.PtrmapIdent("0xc4200cbfe0"),/* slice_item: 124*/125: p.PtrmapIdent("0xc4200ce180"),/* slice_item: 125*/126: p.PtrmapIdent("0xc4200ce320"),/* slice_item: 126*/127: p.PtrmapIdent("0xc4200ce4c0"),/* slice_item: 127*/128: p.PtrmapIdent("0xc4200ce680"),/* slice_item: 128*/129: p.PtrmapIdent("0xc4200ce840"),/* slice_item: 129*/130: p.PtrmapIdent("0xc4200cea00"),/* slice_item: 130*/131: p.PtrmapIdent("0xc4200cebc0"),/* slice_item: 131*/132: p.PtrmapIdent("0xc4200ced80"),/* slice_item: 132*/133: p.PtrmapIdent("0xc4200cef40"),/* slice_item: 133*/134: p.PtrmapIdent("0xc4200cf000"),/* slice_item: 134*/135: p.PtrmapIdent("0xc4200cf260"),/* slice_item: 135*/136: p.PtrmapIdent("0xc4200cf420"),/* slice_item: 136*/137: p.PtrmapIdent("0xc4200cf5e0"),/* slice_item: 137*/138: p.PtrmapIdent("0xc4200d8280"),/* slice_item: 138*/139: p.PtrmapIdent("0xc4200d8300"),/* slice_item: 139*/140: p.PtrmapIdent("0xc4200d83a0"),/* slice_item: 140*/141: p.PtrmapIdent("0xc4200d8460"),/* slice_item: 141*/142: p.PtrmapIdent("0xc4200d8480"),/* slice_item: 142*/143: p.PtrmapIdent("0xc4200d8500"),/* slice_item: 143*/144: p.PtrmapIdent("0xc4200d85a0"),/* slice_item: 144*/145: p.PtrmapIdent("0xc4200d8700"),/* slice_item: 145*/146: p.PtrmapIdent("0xc4200d8860"),/* slice_item: 146*/147: p.PtrmapIdent("0xc4200d88a0"),/* slice_item: 147*/148: p.PtrmapIdent("0xc4200d8a40"),/* slice_item: 148*/149: p.PtrmapIdent("0xc4200d8ba0"),/* slice_item: 149*/150: p.PtrmapIdent("0xc4200d8c40"),/* slice_item: 150*/151: p.PtrmapIdent("0xc4200d8d20"),/* slice_item: 151*/152: p.PtrmapIdent("0xc4200d8f40"),/* slice_item: 152*/153: p.PtrmapIdent("0xc4200d90c0"),/* slice_item: 153*/154: p.PtrmapIdent("0xc4200d9160"),/* slice_item: 154*/155: p.PtrmapIdent("0xc4200d92a0"),/* slice_item: 155*/156: p.PtrmapIdent("0xc4200d9340"),/* slice_item: 156*/157: p.PtrmapIdent("0xc4200d9480"),/* slice_item: 157*/158: p.PtrmapIdent("0xc4200d95c0"),/* slice_item: 158*/159: p.PtrmapIdent("0xc4200d9700"),/* slice_item: 159*/160: p.PtrmapIdent("0xc4200d9860"),/* slice_item: 160*/161: p.PtrmapIdent("0xc4200d9900"),/* slice_item: 161*/162: p.PtrmapIdent("0xc4200d9a60"),/* slice_item: 162*/163: p.PtrmapIdent("0xc4200d9b40"),/* slice_item: 163*/164: p.PtrmapIdent("0xc4200d9c40"),/* slice_item: 164*/165: p.PtrmapIdent("0xc4200d9e80"),/* slice_item: 165*/166: p.PtrmapIdent("0xc4200d9f40"),/* slice_item: 166*/167: p.PtrmapIdent("0xc4200e00a0"),/* slice_item: 167*/168: p.PtrmapIdent("0xc4200e02c0"),/* slice_item: 168*/169: p.PtrmapIdent("0xc4200e04e0"),/* slice_item: 169*/170: p.PtrmapIdent("0xc4200e05a0"),/* slice_item: 170*/171: p.PtrmapIdent("0xc4200e0700"),/* slice_item: 171*/172: p.PtrmapIdent("0xc4200e0860"),/* slice_item: 172*/173: p.PtrmapIdent("0xc4200e0aa0"),/* slice_item: 173*/174: p.PtrmapIdent("0xc4200e0be0"),/* slice_item: 174*/175: p.PtrmapIdent("0xc4200e0c80"),/* slice_item: 175*/176: p.PtrmapIdent("0xc4200e0e40"),/* slice_item: 176*/177: p.PtrmapIdent("0xc4200e0fc0"),/* slice_item: 177*/178: p.PtrmapIdent("0xc4200e1180"),/* slice_item: 178*/179: p.PtrmapIdent("0xc4200e1300"),/* slice_item: 179*/180: p.PtrmapIdent("0xc4200e14c0"),/* slice_item: 180*/181: p.PtrmapIdent("0xc4200e1680"),/* slice_item: 181*/182: p.PtrmapIdent("0xc4200e1840"),/* slice_item: 182*/183: p.PtrmapIdent("0xc4200e1a00"),/* slice_item: 183*/184: p.PtrmapIdent("0xc4200e1be0"),/* slice_item: 184*/185: p.PtrmapIdent("0xc4200e1d60"),/* slice_item: 185*/186: p.PtrmapIdent("0xc4200e1ee0"),/* slice_item: 186*/187: p.PtrmapIdent("0xc4200e4060"),/* slice_item: 187*/188: p.PtrmapIdent("0xc4200e41e0"),/* slice_item: 188*/189: p.PtrmapIdent("0xc4200e4360"),/* slice_item: 189*/190: p.PtrmapIdent("0xc4200e44e0"),/* slice_item: 190*/191: p.PtrmapIdent("0xc4200e4660"),/* slice_item: 191*/192: p.PtrmapIdent("0xc4200e47e0"),/* slice_item: 192*/193: p.PtrmapIdent("0xc4200e4960"),/* slice_item: 193*/194: p.PtrmapIdent("0xc4200e4ae0"),/* slice_item: 194*/195: p.PtrmapIdent("0xc4200e4c60"),/* slice_item: 195*/196: p.PtrmapIdent("0xc4200e4de0"),/* slice_item: 196*/197: p.PtrmapIdent("0xc4200e4f60"),/* slice_item: 197*/198: p.PtrmapIdent("0xc4200e50e0"),/* slice_item: 198*/199: p.PtrmapIdent("0xc4200e52a0"),/* slice_item: 199*/200: p.PtrmapIdent("0xc4200e5540"),/* slice_item: 200*/201: p.PtrmapIdent("0xc4200e5700"),/* slice_item: 201*/202: p.PtrmapIdent("0xc4200e58c0"),/* slice_item: 202*/203: p.PtrmapIdent("0xc4200e5a80"),/* slice_item: 203*/204: p.PtrmapIdent("0xc4200e5c20"),/* slice_item: 204*/205: p.PtrmapIdent("0xc4200e5ce0"),/* slice_item: 205*/206: p.PtrmapIdent("0xc4200e5e80"),/* slice_item: 206*/207: p.PtrmapIdent("0xc4200ec040"),/* slice_item: 207*/208: p.PtrmapIdent("0xc4200ec200"),/* slice_item: 208*/209: p.PtrmapIdent("0xc4200ec280"),/* slice_item: 209*/210: p.PtrmapIdent("0xc4200ec5a0"),/* slice_item: 210*/211: p.PtrmapIdent("0xc4200ec660"),/* slice_item: 211*/212: p.PtrmapIdent("0xc4200ec740"),/* slice_item: 212*/213: p.PtrmapIdent("0xc4200ec7a0"),/* slice_item: 213*/214: p.PtrmapIdent("0xc4200ec820"),/* slice_item: 214*/215: p.PtrmapIdent("0xc4200ec9c0"),/* slice_item: 215*/216: p.PtrmapIdent("0xc4200ecb60"),/* slice_item: 216*/217: p.PtrmapIdent("0xc4200ecc20"),/* slice_item: 217*/218: p.PtrmapIdent("0xc4200ece80"),/* slice_item: 218*/219: p.PtrmapIdent("0xc4200ecf00"),/* slice_item: 219*/220: p.PtrmapIdent("0xc4200ed220"),/* slice_item: 220*/221: p.PtrmapIdent("0xc4200ed3e0"),/* slice_item: 221*/222: p.PtrmapIdent("0xc4200ed5a0"),/* slice_item: 222*/223: p.PtrmapIdent("0xc4200ed620"),/* slice_item: 223*/224: p.PtrmapIdent("0xc4200ed940"),/* slice_item: 224*/225: p.PtrmapIdent("0xc4200edb00"),/* slice_item: 225*/226: p.PtrmapIdent("0xc4200edcc0"),/* slice_item: 226*/227: p.PtrmapIdent("0xc4200f2b40"),/* slice_item: 227*/228: p.PtrmapIdent("0xc4200f30c0"),/* slice_item: 228*/229: p.PtrmapIdent("0xc4200f3180"),/* slice_item: 229*/230: p.PtrmapIdent("0xc4200f33e0"),/* slice_item: 230*/231: p.PtrmapIdent("0xc4200f35c0"),/* slice_item: 231*/232: p.PtrmapIdent("0xc4200f3780"),/* slice_item: 232*/233: p.PtrmapIdent("0xc4200f3a60"),/* slice_item: 233*/234: p.PtrmapIdent("0xc4200f3ae0"),/* slice_item: 234*/235: p.PtrmapIdent("0xc4200f3d00"),/* slice_item: 235*/236: p.PtrmapIdent("0xc4200f3e60"),/* slice_item: 236*/237: p.PtrmapIdent("0xc4200f8000"),/* slice_item: 237*/238: p.PtrmapIdent("0xc4200f8360"),/* slice_item: 238*/239: p.PtrmapIdent("0xc4200f8500"),/* slice_item: 239*/240: p.PtrmapIdent("0xc4200f85a0"),/* slice_item: 240*/241: p.PtrmapIdent("0xc4200f8660"),/* slice_item: 241*/242: p.PtrmapIdent("0xc4200f8740"),/* slice_item: 242*/243: p.PtrmapIdent("0xc4200f8ae0"),/* slice_item: 243*/244: p.PtrmapIdent("0xc4200f8c60"),/* slice_item: 244*/245: p.PtrmapIdent("0xc4200f8de0"),/* slice_item: 245*/246: p.PtrmapIdent("0xc4200f8fa0"),/* slice_item: 246*/247: p.PtrmapIdent("0xc4200f9120"),/* slice_item: 247*/248: p.PtrmapIdent("0xc4200f9460"),/* slice_item: 248*/249: p.PtrmapIdent("0xc4200f9520"),/* slice_item: 249*/250: p.PtrmapIdent("0xc4200f9980"),/* slice_item: 250*/251: p.PtrmapIdent("0xc4200f9ae0"),/* slice_item: 251*/252: p.PtrmapIdent("0xc4200f9d00"),/* slice_item: 252*/253: p.PtrmapIdent("0xc4200f9e80"),/* slice_item: 253*/254: p.PtrmapIdent("0xc4200f9f00"),/* slice_item: 254*/255: p.PtrmapIdent("0xc420102200"),/* slice_item: 255*/256: p.PtrmapIdent("0xc420102240"),/* slice_item: 256*/257: p.PtrmapIdent("0xc4201022a0"),/* slice_item: 257*/258: p.PtrmapIdent("0xc4201022c0"),/* slice_item: 258*/259: p.PtrmapIdent("0xc420102340"),/* slice_item: 259*/260: p.PtrmapIdent("0xc420102460"),/* slice_item: 260*/261: p.PtrmapIdent("0xc4201024c0"),/* slice_item: 261*/262: p.PtrmapIdent("0xc4201025e0"),/* slice_item: 262*/263: p.PtrmapIdent("0xc420102640"),/* slice_item: 263*/}/*slice*/,
}/*struct*/)


}
