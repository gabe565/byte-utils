package exiterr

import "strconv"

type ExitErr struct {
	Code int
}

func (e ExitErr) Error() string {
	return "exit status " + strconv.Itoa(e.Code)
}
