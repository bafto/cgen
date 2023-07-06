package cgen

import (
	"fmt"
	"strings"
)

// a variable declaration
type VarDecl struct {
	// name of the variable
	Name string
	// type of the variable as string ("int", "long long",, "const int" etc.)
	Type Type
	// wether the variable is static
	IsStatic bool
	// wether the variable is extern
	IsExtern bool
	// wether the variable is volatile
	IsVolatile bool
}

func (decl VarDecl) GetName() string {
	return decl.Name
}

func (decl VarDecl) String() string {
	var result string
	if decl.IsStatic {
		result += "static "
	}
	if decl.IsExtern {
		result += "extern "
	}
	if decl.IsVolatile {
		result += "volatile "
	}
	result += string(decl.Type)
	if strings.Contains(result, "(*)") {
		return strings.Replace(result, "(*)", fmt.Sprintf("(*%s)", decl.Name), 1)
	}
	if decl.Name != "" {
		result += " "
	}
	return result + decl.Name
}

func (VarDecl) needsSemicolon() bool {
	return true
}
