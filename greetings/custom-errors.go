package greetings

import "fmt"

type cError struct {
	code int
	msg  string
}

func (e *cError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.code, e.msg)
}

func CallError(x int) (int, error) {
	if x < 0 {
		return 0, &cError{code: 404, msg: "Not Found"}
	}

	return x * 2, nil
}
