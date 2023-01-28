package base

import "fmt"

func Arg(parm ...interface{}) {
	fmt.Println("parm", parm)
}
