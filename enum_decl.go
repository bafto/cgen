package cgen

type EnumValue struct {
	Name  string
	Value string
}

type EnumDecl struct {
	// name of the enum
	Name string
	// the values of the enum
	Values []EnumValue
}

func (decl EnumDecl) GetName() string {
	return decl.Name
}

func (decl EnumDecl) String() string {
	var result string
	result += "enum"
	if decl.Name != "" {
		result += " " + decl.Name
	}
	result += " {\n"

	for i, val := range decl.Values {
		result += "\t" + val.Name
		if val.Value != "" {
			result += " = " + val.Value
		}
		if i < len(decl.Values)-1 {
			result += ","
		}
		result += "\n"
	}
	return result + "}"
}

func (EnumDecl) needsSemicolon() bool {
	return true
}

func (decl EnumDecl) AsAnonymousType() Type {
	decl.Name = ""
	return Type(decl.String())
}

func (decl EnumDecl) GetTypeName() Type {
	return "enum " + Type(decl.Name)
}
