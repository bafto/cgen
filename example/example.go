package main

import "github.com/bafto/cgen"

func main() {
	my_header := cgen.Header{Name: "my_header.h"}

	my_header.AddInclude("<stdbool.h>")

	// typedef
	my_header.AddDecl(cgen.Typedef{
		Name: "Point",
		Type: cgen.StructType([]cgen.VarDecl{ // the struct type is anonymous
			{
				Name: "x",
				Type: cgen.INT,
			},
			{
				Name: "y",
				Type: cgen.INT,
			},
			{
				Name: "z",
				Type: cgen.INT,
			},
		}),
	})

	// struct declaration without typedef
	// this decl will appear above the typedef
	// because the header is sorted by default
	my_header.AddDecl(cgen.StructDecl{
		Name: "Point2",
		Fields: []cgen.VarDecl{
			{
				Name: "x",
				Type: cgen.DOUBLE,
			},
			{
				Name: "y",
				Type: cgen.DOUBLE,
			},
		},
	})

	// extern variable
	my_header.AddDecl(cgen.VarDecl{
		Name:     "my_point",
		Type:     "Point",
		IsExtern: true,
	})

	// func decl
	my_header.AddDecl(cgen.FuncDecl{
		Name: "foo",
		// the order is handled correctly, this results in a unsigned char *const, not a const unsigned char*
		ReturnType: cgen.Const(cgen.Ptr(cgen.Unsigned(cgen.CHAR))),
		Parameters: []cgen.VarDecl{
			{
				Name: "i",
				Type: cgen.INT,
			},
			{
				Name: "func",
				Type: cgen.FuncPtr(cgen.INT, cgen.INT), // generates a function pointer int(*func)(int)
			},
		},
	})

	// write the header file "my_header.h"
	// pass false to disable the sorting of declarations
	my_header.WriteFile(true)
	// my_header.WriteToFile("my_header.h", true)
}