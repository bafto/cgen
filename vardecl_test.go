package cgen

import "testing"

func TestVarDecl(t *testing.T) {
	tests := []struct {
		decl     VarDecl
		expected string
	}{
		{
			VarDecl{
				Name:     "i",
				Type:     INT,
				IsExtern: true,
			},
			"extern int i",
		},
		{
			VarDecl{
				Name:     "j",
				Type:     "unsigned short",
				IsStatic: true,
			},
			"static unsigned short j",
		},
		{
			VarDecl{
				Name: "k",
				Type: "char*",
			},
			"char* k",
		},
		{
			VarDecl{
				Name: "foo",
				Type: FuncPtr(INT, INT, INT),
			},
			"int (*foo) (int, int)",
		},
	}

	for _, test := range tests {
		if result := test.decl.String(); result != test.expected {
			t.Fatalf("%s != %s", result, test.expected)
		}
	}
}
