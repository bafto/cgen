package cgen

type Macro struct {
	// Name of the macor (how it is called)
	Name string
	// parameters of the macro (may be nil)
	Parameters []string
	// the value of the macro
	Replacement string
}

func (decl Macro) GetName() string {
	return decl.Name
}

func (decl Macro) String() string {
	var result string
	result += "#define " + decl.Name

	if len(decl.Parameters) != 0 {
		for i := range decl.Parameters {
			result += decl.Parameters[i]
			if i < len(decl.Parameters)-1 {
				result += ", "
			}
		}
		result += ")"
	}

	return result + " " + decl.Replacement
}

func (Macro) needsSemicolon() bool {
	return false
}
