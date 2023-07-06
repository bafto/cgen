package cgen

// the basically anything that can be in a header file
// like function, variable or struct declarations
// typedefs, macros, etc.
type Declaration interface {
	// name of the declaration
	// e.g. the function name for functionDecls etc.
	GetName() string
	// string representation as it will appear in the generated code
	String() string
	// wether the declaration must be terminated with a semicolon
	needsSemicolon() bool
}
