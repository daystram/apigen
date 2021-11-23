package generator

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/daystram/apigen/internal/definition"
)

type FileGroup []File

type File struct {
	Name string
	Dir  string
	AST  *ast.File
}

func Generate(d definition.Service, pkg string) (FileGroup, error) {
	fg := make(FileGroup, 0)

	// TODO: generate go mod files

	main, err := generateMain(d, pkg)
	if err != nil {
		return nil, err
	}
	fg = append(fg, main...)

	// TODO: generate controllers

	return fg, nil
}

func generateMain(d definition.Service, pkg string) (FileGroup, error) {
	fg := make(FileGroup, 0)

	fg = append(fg, File{
		Name: "main.go",
		Dir:  "",
		AST: &ast.File{
			Package: 0,
			Name:    &ast.Ident{Name: "main"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.IMPORT,
					Specs: []ast.Spec{
						&ast.ImportSpec{
							Path: &ast.BasicLit{
								Kind:  token.STRING,
								Value: `"fmt"`,
							},
						},
						&ast.ImportSpec{
							Path: &ast.BasicLit{
								Kind:  token.STRING,
								Value: `"log"`,
							},
						},
						&ast.ImportSpec{
							Path: &ast.BasicLit{
								Kind:  token.STRING,
								Value: `"net/http"`,
							},
						},
						&ast.ImportSpec{
							Path: &ast.BasicLit{
								Kind:  token.STRING,
								Value: `"time"`,
							},
						},
						&ast.ImportSpec{
							Path: &ast.BasicLit{
								Kind:  token.STRING,
								Value: fmt.Sprintf(`"%s/controllers"`, pkg),
							},
						},
					},
				},
				&ast.FuncDecl{
					Name: &ast.Ident{Name: "main"},
					Type: &ast.FuncType{},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.AssignStmt{
								Lhs: []ast.Expr{
									&ast.Ident{Name: "s"},
								},
								Tok: token.DEFINE,
								Rhs: []ast.Expr{
									&ast.UnaryExpr{
										Op: token.AND,
										X: &ast.CompositeLit{
											Type: &ast.SelectorExpr{X: &ast.Ident{Name: "http"}, Sel: &ast.Ident{Name: "Server"}},
											Elts: []ast.Expr{
												&ast.KeyValueExpr{
													Key: &ast.Ident{Name: "Addr"},
													Value: &ast.BasicLit{
														Kind:  token.STRING,
														Value: fmt.Sprintf(`"%s:%d"`, d.Host, d.Port),
													},
												},
												&ast.KeyValueExpr{
													Key: &ast.Ident{Name: "Handler"},
													Value: &ast.CallExpr{
														Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "controllers"}, Sel: &ast.Ident{Name: "InitializeRouter"}},
													},
												},
												&ast.KeyValueExpr{
													Key: &ast.Ident{Name: "ReadTimeout"},
													Value: &ast.BinaryExpr{
														X:  &ast.BasicLit{Kind: token.INT, Value: "10"},
														Op: token.MUL,
														Y:  &ast.SelectorExpr{X: &ast.Ident{Name: "time"}, Sel: &ast.Ident{Name: "Second"}},
													},
												},
												&ast.KeyValueExpr{
													Key: &ast.Ident{Name: "WriteTimeout"},
													Value: &ast.BinaryExpr{
														X:  &ast.BasicLit{Kind: token.INT, Value: "10"},
														Op: token.MUL,
														Y:  &ast.SelectorExpr{X: &ast.Ident{Name: "time"}, Sel: &ast.Ident{Name: "Second"}},
													},
												},
												&ast.KeyValueExpr{
													Key: &ast.Ident{Name: "MaxHeaderBytes"},
													Value: &ast.BinaryExpr{
														X:  &ast.BasicLit{Kind: token.INT, Value: "1"},
														Op: token.SHL,
														Y:  &ast.BasicLit{Kind: token.INT, Value: "20"},
													},
												},
											},
										},
									},
								},
							},
							&ast.AssignStmt{
								Lhs: []ast.Expr{
									&ast.Ident{Name: "err"},
								},
								Tok: token.DEFINE,
								Rhs: []ast.Expr{
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "ListenAndServe"}},
									},
								},
							},
							&ast.IfStmt{
								Cond: &ast.BinaryExpr{
									X:  &ast.Ident{Name: "err"},
									Op: token.NEQ,
									Y:  &ast.Ident{Name: "nil"},
								},
								Body: &ast.BlockStmt{
									List: []ast.Stmt{
										&ast.ExprStmt{
											X: &ast.CallExpr{
												Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "log"}, Sel: &ast.Ident{Name: "Fatalln"}},
												Args: []ast.Expr{
													&ast.Ident{Name: "err"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	return fg, nil
}

