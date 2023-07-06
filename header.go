package cgen

import (
	"fmt"
	"os"
	"strings"
)

type Header struct {
	// name of the header file (header.h)
	Name string
	// includes of the header
	// e.g.: ["<stdlib.h>", "\"my_header.h\"", "<stdbool.h>"]
	Includes []string
	// declarations of this header
	// (structs, enums, (extern) variables, functions, typedefs, macros...)
	Decls []Declaration
}

const incl_fmt = "#include %s\n"

// returns the generated header file as a string
// if ordered is true, the declarations are ordered by their typ
func (h *Header) AsString(ordered bool) string {
	builder := strings.Builder{}
	builder.Grow(len(h.Includes) * (len(incl_fmt) + 5))

	for i := range h.Includes {
		builder.WriteString(fmt.Sprintf("#include %s\n", h.Includes[i]))
	}
	builder.WriteRune('\n')

	declWriter := func(decl Declaration) {
		builder.WriteString(decl.DeclString())
		if decl.needsSemicolon() {
			builder.WriteRune(';')
		}
		builder.WriteRune('\n')
	}

	if ordered {
		// print the decls in a formatted order
		applyFiltered[VarDecl](h.Decls, declWriter)
		applyFiltered[FuncDecl](h.Decls, declWriter)
	} else {
		applyFiltered[Declaration](h.Decls, declWriter)
	}

	return builder.String()
}

// writes h to path
func (h *Header) WriteToFile(path string) error {
	return os.WriteFile(path, []byte(h.AsString(true)), os.ModePerm)
}

// h.WriteToFile(h.Name)
func (h *Header) WriteFile() error {
	return h.WriteToFile(h.Name)
}

// appends incl to h.Includes
func (h *Header) AddRawInclude(incl string) {
	h.Includes = append(h.Includes, incl)
}

// like AddInclude but with <> instead of ""
func (h *Header) AddStdInclude(incl string) {
	h.Includes = append(h.Includes, fmt.Sprintf("<%s>", incl))
}

// adds "incl" to h.Includes
func (h *Header) AddInclude(incl string) {
	h.Includes = append(h.Includes, fmt.Sprintf(`"%s"`, incl))
}

func (h *Header) AddDecl(decl Declaration) {
	h.Decls = append(h.Decls, decl)
}

// applies fn to ever T in decls
func applyFiltered[T Declaration](decls []Declaration, fn func(Declaration)) {
	for _, decl := range decls {
		if _, ok := decl.(T); ok {
			fn(decl)
		}
	}
}
