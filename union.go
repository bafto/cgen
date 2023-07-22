package cgen

type UnionDecl struct {
	// name of the union
	Name string
	// the Fields of the union
	// IsStatic, IsVolatile and IsExtern of them should be false
	Fields []VarDecl
}

func (decl UnionDecl) GetName() string {
	return decl.Name
}

func (decl UnionDecl) String() string {
	var result string
	result += "union "
	if decl.Name != "" {
		result += decl.Name + " {\n"
	} else {
		result += "{\n"
	}

	for _, field := range decl.Fields {
		// just to be sure
		field.IsExtern = false
		field.IsStatic = false
		field.IsVolatile = false
		result += "\t" + field.String() + ";\n"
	}
	return result + "}"
}

func (UnionDecl) needsSemicolon() bool {
	return true
}

// returns the given decl as inline type (without the name)
func (decl UnionDecl) AsAnonymousType() Type {
	decl.Name = ""
	return Type(decl.String())
}

func (decl UnionDecl) GetTypeName() Type {
	return "union " + Type(decl.Name)
}
