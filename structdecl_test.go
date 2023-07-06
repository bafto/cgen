package cgen

import "testing"

func TestStructDecl(t *testing.T) {
	tests := []struct {
		decl     StructDecl
		expected string
	}{
		{
			StructDecl{
				Name: "Point",
				Fields: []VarDecl{
					{
						Name: "x",
						Type: INT,
					},
					{
						Name: "y",
						Type: INT,
					},
					{
						Name: "z",
						Type: INT,
					},
				},
			},
			`struct Point {
	int x;
	int y;
	int z;
}`,
		},
		{
			StructDecl{
				Name: "HasFunc",
				Fields: []VarDecl{
					{
						Name: "func",
						Type: FuncPtr(INT, INT, INT),
					},
				},
			},
			`struct HasFunc {
	int (*func) (int, int);
}`,
		},
	}

	for _, test := range tests {
		if result := test.decl.String(); result != test.expected {
			t.Fatalf("%s != %s", result, test.expected)
		}
	}
}
