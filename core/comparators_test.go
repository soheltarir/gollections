package core

import "testing"

func TestIntCompare(t *testing.T) {
	if IntCompare(1, 2) != Lesser {
		t.Errorf("1 should be less than 2")
	}
	if IntCompare(2, 1) != Greater {
		t.Errorf("2 should be greater than 1")
	}
	if IntCompare(10, 10) != Equal {
		t.Errorf("Equality check failed with value 10")
	}
}

func TestStringCompare(t *testing.T) {
	if StringCompare("app", "apple") != Lesser {
		t.Errorf("app should be less than apple lexicographically")
	}
	if StringCompare("apple", "app") != Greater {
		t.Errorf("apple should be greater than app lexicographically")
	}
	if StringCompare("app", "app") != Equal {
		t.Errorf("Equality check failed with value 'app'")
	}
}
