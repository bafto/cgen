package cgen

import (
	"fmt"
	"strings"
)

type CType string

// some predefined C primitive types
const (
	VOID     CType = "void"
	CHAR     CType = "char"
	SHORT    CType = "short"
	INT      CType = "int"
	LONG     CType = "long"
	LONGLONG CType = "long long"
	FLOAT    CType = "float"
	DOUBLE   CType = "double"
	// _Bool primitive type (c99)
	BOOL CType = "_Bool"
	// needs the <stdbool.h> include
	STDBOOL CType = "bool"
)

// returns a type representing a pointer to typ
func Ptr(typ CType) CType {
	return typ + "*"
}

// returns a const version of typ
func Const(typ CType) CType {
	// const char* != char *const
	if strings.HasSuffix(string(typ), "*") {
		return typ[:len(typ)-1] + " *const"
	}
	return "const " + typ
}

// returns a unsigned version of typ
// (should only be used on numeric types)
func Unsigned(typ CType) CType {
	return "unsigned " + typ
}

// returns a function pointer type without name
// e.g.: FuncPtr(INT, INT, INT) == "int (*) (int, int)"
func FuncPtr(returnType CType, params ...CType) CType {
	result := CType(fmt.Sprintf("%s (*) (", returnType))

	for i, param := range params {
		result += param
		if i < len(params)-1 {
			result += ", "
		}
	}

	return result + ")"
}
