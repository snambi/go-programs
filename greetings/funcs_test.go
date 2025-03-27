package greetings

import "testing"

func TestFact2(t *testing.T) {
	x := Fact(2)

	t.Log("Fact(2) = ", x)
}

func TestFact1(t *testing.T) {
	x := Fact(1)

	t.Log("Fact(1) = ", x)
}
func TestFact0(t *testing.T) {
	x := Fact(0)

	if x != 1 {
		t.Errorf("Expected Fact(0) to be 1, but got %d", x)
	}
}

func TestFact3(t *testing.T) {
	x := Fact(3)

	if x != 6 {
		t.Errorf("Expected Fact(3) to be 6, but got %d", x)
	}
}

func TestFact5(t *testing.T) {
	x := Fact(5)

	if x != 120 {
		t.Errorf("Expected Fact(5) to be 120, but got %d", x)
	}
}

func TestFactNegative(t *testing.T) {
	x := Fact(-1)
	if x != -1 {
		t.Errorf("Expected Fact(-1) to be -1, but got %d", x)
	}
}

func TestFuncFunc(t *testing.T) {
	x := FuncFunc(1)

	if x != 12 {
		t.Errorf("Expected FuncFunc(1) to be 12, but got %d", x)
	}
}
