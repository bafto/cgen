package cgen

// a function declaration
type FuncDecl struct {
	// name of the function
	Name string
	// return type of the function as string ("int", "long long", etc.)
	ReturnType Type
	// wether the function is static
	IsStatic bool
	// wether the function is extern
	IsExtern bool
	// wether the function also takes variadic arguments
	IsVariadic bool
	// the parameters of the function
	// IsStatic, IsVolatile and IsExtern of them should be false
	Parameters []VarDecl
}

func (decl FuncDecl) GetName() string {
	return decl.Name
}

func (decl FuncDecl) String() string {
	var result string
	if decl.IsStatic {
		result += "static "
	}
	if decl.IsExtern {
		result += "extern "
	}
	result += string(decl.ReturnType) + " " + decl.Name + "("
	for i, param := range decl.Parameters {
		// just to be sure
		param.IsExtern = false
		param.IsStatic = false
		param.IsVolatile = false
		result += param.String()
		if i < len(decl.Parameters)-1 || decl.IsVariadic {
			result += ", "
		}
	}
	if decl.IsVariadic {
		result += "..."
	}
	result += ")"
	return result
}

func (FuncDecl) needsSemicolon() bool {
	return true
}
