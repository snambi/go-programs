package greetings

import "testing"

func TestCall(t *testing.T) {
	Call(MONDAY)
	t.Log("TestCall")
}
