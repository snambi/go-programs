package greetings

import "testing"

func TestCallError(t *testing.T) {
	_, e := CallError(-1)
	if e != nil {
		t.Log("TestCallError: ", e.Error())
	}
}
