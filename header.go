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
	incl_guard := strings.ToUpper(strings.TrimSuffix(h.Name, ".h") + "_H")
	builder.WriteString(fmt.Sprintf("#ifndef %s\n#define %s\n\n", incl_guard, incl_guard))

	builder.Grow(len(h.Includes) * (len(incl_fmt) + 5))

	for i := range h.Includes {
		builder.WriteString(fmt.Sprintf("#include %s\n", h.Includes[i]))
	}
	builder.WriteRune('\n')

	declWriter := func(decl Declaration) {
		builder.WriteString(decl.String())
		if decl.needsSemicolon() {
			builder.WriteRune(';')
		}
		builder.WriteRune('\n')
	}

	if ordered {
		// print the decls in a formatted order
		applyFiltered[Macro](h.Decls, declWriter)
		builder.WriteRune('\n')
		applyFiltered[StructDecl](h.Decls, declWriter)
		builder.WriteRune('\n')
		applyFiltered[Typedef](h.Decls, declWriter)
		builder.WriteRune('\n')
		applyFiltered[VarDecl](h.Decls, declWriter)
		builder.WriteRune('\n')
		applyFiltered[FuncDecl](h.Decls, declWriter)
		builder.WriteRune('\n')
	} else {
		applyFiltered[Declaration](h.Decls, declWriter)
	}

	builder.WriteString("\n#endif\n")
	return builder.String()
}

// writes h to path
func (h *Header) WriteToFile(path string, ordered bool) error {
	return os.WriteFile(path, []byte(h.AsString(ordered)), os.ModePerm)
}

// h.WriteToFile(h.Name)
func (h *Header) WriteFile(ordered bool) error {
	return h.WriteToFile(h.Name, ordered)
}

// adds incl to h.Includes
func (h *Header) AddInclude(incl string) {
	h.Includes = append(h.Includes, incl)
}

func (h *Header) Add(decl Declaration) {
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
