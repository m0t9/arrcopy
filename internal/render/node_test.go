package render_test

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/m0t9/goperflint/internal/render"
)

func TestNode_ValidExpression(t *testing.T) {
	fset := token.NewFileSet()
	expr := &ast.BasicLit{
		ValuePos: 1,
		Kind:     token.INT,
		Value:    "42",
	}

	result := render.Node(fset, expr)
	expected := "42" // The expected output for a basic literal

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestNode_InvalidInput(t *testing.T) {
	fset := token.NewFileSet()
	invalidInput := struct{}{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for invalid input, but did not panic")
		}
	}()

	render.Node(fset, invalidInput)
}

func TestNode_EmptyExpression(t *testing.T) {
	fset := token.NewFileSet()
	expr := &ast.CompositeLit{
		Type: &ast.Ident{Name: "struct{}"},
	}

	result := render.Node(fset, expr)
	expected := "struct{}{}"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestNode_SimpleFunction(t *testing.T) {
	fset := token.NewFileSet()
	funcDecl := &ast.FuncDecl{
		Name: &ast.Ident{Name: "MyFunc"},
		Type: &ast.FuncType{},
	}

	result := render.Node(fset, funcDecl)
	expected := "func MyFunc()"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestNode_ComplexStruct(t *testing.T) {
	fset := token.NewFileSet()
	structDecl := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Type: &ast.Ident{Name: "int"},
					Names: []*ast.Ident{
						{Name: "Field1"},
					},
				},
			},
		},
	}

	result := render.Node(fset, structDecl)
	expected := "struct {\n\tField1 int\n}"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
