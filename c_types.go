package cgen

import (
	"fmt"
	"strings"
)

// a Type name in C (can be any string like "int" or "struct Point" etc.)
type Type string

// some predefined C primitive types
const (
	VOID     Type = "void"
	CHAR     Type = "char"
	SHORT    Type = "short"
	INT      Type = "int"
	LONG     Type = "long"
	LONGLONG Type = "long long"
	FLOAT    Type = "float"
	DOUBLE   Type = "double"
	// _Bool primitive type (c99)
	BOOL Type = "_Bool"
	// needs the <stdbool.h> include
	STDBOOL Type = "bool"
)

// returns a type representing a pointer to typ
func Ptr(typ Type) Type {
	return typ + "*"
}

// returns a const version of typ
func Const(typ Type) Type {
	// const char* != char *const
	if strings.HasSuffix(string(typ), "*") {
		return typ[:len(typ)-1] + " *const"
	}
	return "const " + typ
}

// returns a unsigned version of typ
// (should only be used on numeric types)
func Unsigned(typ Type) Type {
	return "unsigned " + typ
}

// returns a function pointer type without name
// e.g.: FuncPtr(INT, INT, INT) == "int (*) (int, int)"
func FuncPtr(returnType Type, params ...Type) Type {
	result := Type(fmt.Sprintf("%s (*) (", returnType))

	for i, param := range params {
		result += param
		if i < len(params)-1 {
			result += ", "
		}
	}

	return result + ")"
}

// creates a anonymous struct type with the given fields
func StructType(fields []VarDecl) Type {
	return StructDecl{Fields: fields}.AsType()
}
