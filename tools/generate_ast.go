package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: generate_ast <output directory>\n")
		os.Exit(64)
	}
	outputDir := os.Args[1]

	defineAst(outputDir, "Expr", []string{
		"Binary   : left Expr, operator *Token, right Expr",
		"Grouping : expression Expr",
		"Literal  : value interface{}",
		"Unary    : operator *Token, right Expr",
	})

}

func defineAst(outputDir, baseName string, types []string) {
	path := outputDir + "/" + strings.ToLower(baseName) + ".go"
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writer.WriteString("package ast\n")
	writer.WriteString("type " + baseName + " interface {\n")
	writer.WriteString("Accept(VisitorExpr) (interface{}, error) \n")
	writer.WriteString("}\n")

	defineVisitor(writer, baseName, types)

	for _, typ := range types {
		splitTyp := strings.Split(typ, ":")
		structName := strings.TrimSpace(splitTyp[0])
		fields := strings.TrimSpace(splitTyp[1])
		defineType(writer, baseName, structName, fields)
	}

}

func defineVisitor(writer *bufio.Writer, baseName string, types []string) {
	writer.WriteString("type " + "Visitor" + baseName + " interface{\n")
	for _, typ := range types {
		typ := strings.TrimSpace(strings.Split(typ, ":")[0])
		writer.WriteString("\tvisit" + typ + "Expr(*" + typ + ") (interface{}, error) \n")
	}

	writer.WriteString("}\n")
}

func defineType(writer *bufio.Writer, baseName, structName, fields string) {
	writer.WriteString("type " + structName + " struct {\n")

	fieldList := strings.Split(fields, ",")
	for _, field := range fieldList {
		field := strings.TrimSpace(field)
		upperCased := strings.ToUpper(string(field[0])) + field[1:]
		writer.WriteString("\t" + upperCased + "\n")
	}

	writer.WriteString(" } \n\n")

	writer.WriteString("func New" + structName + "(" + fields + ") " + baseName + "{\n")
	writer.WriteString("return &" + structName + "{")
	args := make([]string, 0)
	for _, field := range fieldList {
		name := strings.Split(strings.TrimSpace(field), " ")[0]
		args = append(args, name)
	}
	writer.WriteString(strings.Join(args, ","))

	writer.WriteString("}\n")
	writer.WriteString("}\n\n")

	varName := string(strings.ToLower(structName)[0])
	writer.WriteString("func (" + varName + " *" + structName + ") Accept(visitor VisitorExpr) (interface{}, error) {\n\treturn visitor.visit" + structName + "Expr(" + varName + ")\n}\n\n\n")

}
