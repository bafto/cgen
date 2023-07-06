package cgen

type StructDecl struct {
	// name of the struct
	Name string
	// the Fields of the struct
	// IsStatic, IsVolatile and IsExtern of them should be false
	Fields []VarDecl
}

func (decl StructDecl) GetName() string {
	return decl.Name
}

func (decl StructDecl) String() string {
	return decl.DeclString()
}

func (decl StructDecl) DeclString() string {
	var result string
	result += "struct "
	if decl.Name != "" {
		result += decl.Name + " {\n"
	} else {
		result += "{\n"
	}

	for i := range decl.Fields {
		// just to be sure
		decl.Fields[i].IsExtern = false
		decl.Fields[i].IsStatic = false
		decl.Fields[i].IsVolatile = false
		result += "\t" + decl.Fields[i].DeclString() + ";\n"
	}
	return result + "}"
}

func (StructDecl) needsSemicolon() bool {
	return true
}

// returns the given decl as inline type (without the name)
func (decl StructDecl) AsType() Type {
	decl.Name = ""
	return Type(decl.DeclString())
}
