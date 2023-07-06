package cgen

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
	_BOOL CType = "_Bool"
	// needs the <stdbool.h> include
	STDBOOL CType = "bool"
)

// appends the * suffix to typ
func Ptr(typ CType) CType {
	return typ + "*"
}

// preprends the const prefix to typ
func Const(typ CType) CType {
	return "const " + typ
}

// prepends the unsigned prefix to typ
// (should only be used on numeric types)
func Unsigned(typ CType) CType {
	return "unsigned " + typ
}
