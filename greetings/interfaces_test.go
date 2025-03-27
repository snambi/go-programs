package greetings

import "testing"

func TestMeasure(t *testing.T) {
	c := circle{radius: 5}
	Measure(c)
}

func TestDet1(t *testing.T) {
	c := circle{radius: 5}
	DetectCircle(c)
}

func TestDet2(t *testing.T) {
	c := rectangle{length: 5, width: 10}
	DetectCircle(c)
}
