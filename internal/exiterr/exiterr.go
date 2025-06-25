package exiterr

import "strconv"

type ExitError struct {
	Code int
}

func (e ExitError) Error() string {
	return "exit status " + strconv.Itoa(e.Code)
}
