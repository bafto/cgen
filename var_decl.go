package cgen

// a variable declaration
type VarDecl struct {
	// name of the variable
	Name string
	// type of the variable as string ("int", "long long",, "const int" etc.)
	Type CType
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
	return decl.DeclString()
}

func (decl VarDecl) DeclString() string {
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
	result += string(decl.Type) + " " + decl.Name
	return result
}

func (VarDecl) needsSemicolon() bool {
	return true
}
