package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"
)

func parseFile(code string) []byte {
	fst := token.NewFileSet()
	f, _ := parser.ParseFile(fst, "main.go", code, 0)
	m := walk(fst, f)
	res, _ := json.Marshal(m)
	return res
}

func walk(fst *token.FileSet, node interface{}) map[string]interface{} {
	if node == nil {
		return nil
	}

	m := make(map[string]interface{})

	if _, ok := node.(*ast.Scope); ok {
		return nil
	}

	if _, ok := node.(*ast.Object); ok {
		return nil
	}

	val := reflect.ValueOf(node)
	if val.IsNil() {
		return nil
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	ty := val.Type()
	m["_type"] = ty.Name()
	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)
		val := val.Field(i)
		if strings.HasSuffix(field.Name, "Pos") {
			continue
		}
		switch field.Type.Kind() {
		case reflect.Array, reflect.Slice:
			list := make([]interface{}, 0, val.Len())
			for i := 0; i < val.Len(); i++ {
				if item := walk(fst, val.Index(i).Interface()); item != nil {
					list = append(list, item)
				}
			}
			m[field.Name] = list
		case reflect.Ptr:
			if child := walk(fst, val.Interface()); child != nil {
				m[field.Name] = child
			}
		case reflect.Interface:
			if child := walk(fst, val.Interface()); child != nil {
				m[field.Name] = child
			}
		case reflect.String:
			m[field.Name] = val.String()
		case reflect.Int:
			if field.Type.Name() == "Token" {
				m[field.Name] = token.Token(val.Int()).String()
			} else {
				m[field.Name] = val.Int()
			}
		case reflect.Bool:
			m[field.Name] = val.Bool()
		default:
			fmt.Fprintln(os.Stderr, field)
		}
	}
	if n, ok := node.(ast.Node); ok {
		start := fst.Position(n.Pos())
		end := fst.Position(n.End())
		m["Loc"] = map[string]interface{}{"Start": start, "End": end}
	}
	return m
}
