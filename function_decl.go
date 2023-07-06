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
	// the parameters of the function
	// IsStatic, IsVolatile and IsExtern of them should be false
	Parameters []VarDecl
}

func (decl FuncDecl) GetName() string {
	return decl.Name
}

func (decl FuncDecl) String() string {
	result := string(decl.ReturnType) + " " + decl.Name + "("
	for i := range decl.Parameters {
		result += string(decl.Parameters[i].Type)
		if i < len(decl.Parameters)-1 {
			result += ", "
		}
	}
	return result + ")"
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
		// just to be sure
		decl.Parameters[i].IsExtern = false
		decl.Parameters[i].IsStatic = false
		decl.Parameters[i].IsVolatile = false
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
