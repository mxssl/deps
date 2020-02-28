package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	D := addPackage("D")
	E := addPackage("E")
	C := addPackage("C", D, E)
	B := addPackage("B", C, E)
	A := addPackage("A", D, B)

	expected := []string{"D", "B", "C", "E"}
	output := getDeps(A)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}
}

func TestCase2(t *testing.T) {
	D := addPackage("D")
	E := addPackage("E")
	C := addPackage("C", D, E)
	B := addPackage("B", C, E)

	expected := []string{"C", "D", "E"}
	output := getDeps(B)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}
}

func TestCase3(t *testing.T) {
	D := addPackage("D")
	E := addPackage("E")
	C := addPackage("C", D, E)

	expected := []string{"D", "E"}
	output := getDeps(C)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}
}
