package main

import "github.com/bafto/cgen"

func main() {
	my_header := cgen.Header{Name: "my_header.h"}

	my_header.AddInclude("<stdbool.h>")

	my_header.Add(cgen.Macro{
		Name:        "_POSIX_C_SOURCE",
		Replacement: "1",
	})

	my_header.Add(cgen.Macro{
		Name: "NOMINMAX",
	})

	enum := cgen.EnumDecl{
		Name: "Day",
		Values: []cgen.EnumValue{
			{
				Name: "Saturday",
			},
			{
				Name:  "Sunday",
				Value: "0",
			},
			{
				Name: "Monday",
			},
			{
				Name: "Wednesday",
			},
		},
	}

	my_header.Add(enum)

	my_header.Add(cgen.Typedef{Name: "WeekDay", Type: enum.GetTypeName()})

	my_header.Add(cgen.Typedef{
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
	my_header.Add(cgen.StructDecl{
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
	my_header.Add(cgen.VarDecl{
		Name:     "my_point",
		Type:     "Point",
		IsExtern: true,
	})

	// func decl
	my_header.Add(cgen.FuncDecl{
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

	my_header.Add(cgen.FuncDecl{
		Name:       "bar",
		ReturnType: cgen.VOID,
		Parameters: []cgen.VarDecl{
			{
				Type: cgen.STRING, // const char*
			},
		},
		IsVariadic: true,
	})

	my_header.Add(cgen.FuncDecl{
		Name:       "baz",
		ReturnType: cgen.VOID,
		IsVariadic: true,
	})

	// write the header file "my_header.h"
	// pass false to disable the sorting of declarations
	my_header.WriteFile(true)
	// my_header.WriteToFile("my_header.h", true)
}
