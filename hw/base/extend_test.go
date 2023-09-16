package base

import "testing"

func TestExec(t *testing.T) {
	ta := TaskA{}

	tb := TaskB{}
	Exec(&ta)
	Exec(&tb)

}
