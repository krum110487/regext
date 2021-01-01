package regext

import (
	"testing"
)

func TestFull(t *testing.T) {
	result := NewRegext([]byte(`seafood FoOl2 fool3 fool4 fool5 fool6   `)).PrintRaw().
		Trim().PrintRaw().
		ToLower().PrintRaw().
		FindAll(`foo..?`).PrintRaw().
		Trim().PrintRaw().
		FindAll(`(fool[4-6])`).Println("\tFall=", "").
		ReplaceAll(`ol`, `od`).Println("\t\tRall=", "").
		FilterOut(`food4`, `food.*`).Println("\t\tFout=", "").
		FilterAny(`food6`, `food5`).Println("\t\tFany=", "").
		FilterOutAny(`food3`, `food4`).Println("\t\tFony=", "").
		DeleteAny(`o`, `5`).Println("\t\tDany=", "").
		FilterOutByLen(2, 2).Println("\t\tFobl=", "").
		Split(`d`, `f`).Println("\t\tSplt=", "").
		String()

	if result != "6" {
		t.Errorf("TestFull failed, expected %v, got %v", "6", result)
	}
}

func TestTrim(t *testing.T) {
	result := NewRegext([]byte(`
	
	
	Trimmed    Value is Trim
	
	
	`)).
		Trim().PrintRaw().String()

	if result != "Trimmed    Value is Trim" {
		t.Errorf("TestFull failed, expected \"%v\", got \"%v\"", "Trimmed    Value is Trim", result)
	}
}

func TestJoin(t *testing.T) {
	result := NewRegext([]byte(`Test1 Test2 Test3 Test4 Test5`)).
		SplitIgnoreBlanks()

	if result != "Trimmed    Value is Trim" {
		t.Errorf("TestFull failed, expected \"%v\", got \"%v\"", "Trimmed    Value is Trim", result)
	}
}

//TODO: fill out test cases...
