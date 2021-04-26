package core

import "testing"

func TestIntChecker(t *testing.T) {
	if !IntChecker(1) {
		t.Errorf("Expected true but got false for '1'")
	}
	if IntChecker("one") {
		t.Errorf("Expected false but got true for 'one'")
	}
}

func TestStringChecker(t *testing.T) {
	if !StringChecker("one") {
		t.Errorf("Expected true but got false for 'one'")
	}
	if StringChecker(1) {
		t.Errorf("Expected false but got true for '1'")
	}
}
