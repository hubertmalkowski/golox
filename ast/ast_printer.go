package ast

import (
	"fmt"
	"strings"
)

type AstPrinter struct {
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

func (a *AstPrinter) Print(expr Expr) (string, error) {
	r, _ := expr.Accept(a)
	return fmt.Sprintf("%v", r), nil
}

// visitBinaryExpr implements VisitorExpr.
func (a *AstPrinter) visitBinaryExpr(expr *Binary) (interface{}, error) {
	return a.paranthesize(expr.Operator.Lexeme, expr.Left, expr.Right), nil
}

// visitGroupingExpr implements VisitorExpr.
func (a *AstPrinter) visitGroupingExpr(expr *Grouping) (interface{}, error) {
	return a.paranthesize("group", expr.Expression), nil
}

// visitLiteralExpr implements VisitorExpr.
func (a *AstPrinter) visitLiteralExpr(expr *Literal) (interface{}, error) {
	if expr.Value == nil {
		return "nil", nil
	}

	return fmt.Sprintf("%v", expr.Value), nil
}

// visitUnaryExpr implements VisitorExpr.
func (a *AstPrinter) visitUnaryExpr(expr *Unary) (interface{}, error) {
	return a.paranthesize(expr.Operator.Lexeme, expr.Right), nil
}

func (a *AstPrinter) paranthesize(name string, exprs ...Expr) string {
	var sb strings.Builder

	sb.WriteString("(" + name)
	for _, expr := range exprs {
		eval_accept, _ := expr.Accept(a)
		sb.WriteString(" ")
		sb.WriteString(eval_accept.(string))
	}

	sb.WriteString(")")
	return sb.String()
}
