package process

import (
	"errors"
	"testing"

	. "github.com/babourine/x/pkg/testutil"
)

func TestUnitBail(t *testing.T) {

	f := func(e error, exp int) {
		oldOsExit := osExit
		defer func() { osExit = oldOsExit }()

		var got int
		myExit := func(code int) {
			got = code
		}

		osExit = myExit
		Bail(`to exit...`, e)
		if got != exp {
			F(t, &P{`exp`: exp, `got`: got})
		}
	}

	f(nil, 0)
	f(errors.New(`...or to not exit`), 1)

}
