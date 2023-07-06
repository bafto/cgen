package cgen

type Typedef struct {
	// the alias name
	Name string
	// the type that is typedefed
	Type Type
}

func (decl Typedef) GetName() string {
	return decl.Name
}

func (decl Typedef) String() string {
	var result string
	result += "typedef " + string(decl.Type) + " " + decl.Name
	return result
}

func (Typedef) needsSemicolon() bool {
	return true
}
