package cgen

import "testing"

func TestCTypes(t *testing.T) {
	tests := []struct {
		typ      CType
		expected string
	}{
		{
			INT,
			"int",
		},
		{
			Const(INT),
			"const int",
		},
		{
			Ptr(Const(INT)),
			"const int*",
		},
		{
			Const(Ptr(INT)),
			"int *const",
		},
	}

	for _, test := range tests {
		if string(test.typ) != test.expected {
			t.Fatalf("%s != %s", test.typ, test.expected)
		}
	}
}
