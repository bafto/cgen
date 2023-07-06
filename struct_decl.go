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
	result += "struct " + decl.Name + " {\n"
	for _, field := range decl.Fields {
		// just to be sure
		field.IsExtern = false
		field.IsStatic = false
		field.IsVolatile = false
		result += "\t" + field.DeclString() + ";\n"
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
