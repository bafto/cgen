package cgen

// a function declaration
type FuncDecl struct {
	// name of the function
	Name string
	// return type of the function as string ("int", "long long", etc.)
	ReturnType CType
	// wether the function is static
	IsStatic bool
	// wether the function is extern
	IsExtern bool
	// the parameters of the function
	// IsStatic, IsVolatile and IsExtern of them should be false
	Parameters []VarDecl
}

func (decl FuncDecl) GetName() string {
	return decl.Name
}

func (decl FuncDecl) String() string {
	return string(decl.ReturnType) + " "
}

func (decl FuncDecl) DeclString() string {
	var result string
	if decl.IsStatic {
		result += "static "
	}
	if decl.IsExtern {
		result += "extern "
	}
	result += string(decl.ReturnType) + " " + decl.Name + "("
	for i := range decl.Parameters {
		result += decl.Parameters[i].DeclString()
		if i < len(decl.Parameters)-1 {
			result += ", "
		}
	}
	result += ")"
	return result
}

func (FuncDecl) needsSemicolon() bool {
	return true
}
