package cgen

type Declaration interface {
	// name of the declaration
	// e.g. the function name for functionDecls
	GetName() string
	// string representation (might be the full DeclString)
	String() string
	// the decl as C code without the semicolon
	// e.g.: void f(), extern int i, etc.
	DeclString() string
	// wether the declaration must be terminated with a semicolon
	needsSemicolon() bool
}
