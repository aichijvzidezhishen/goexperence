package base

import (
	"fmt"
	"testing"
)

func TestIfaceExtend(t *testing.T) {

	IfaceExtend()

}

func TestJudge(t *testing.T) {
	// var i interface{}
	// var i interface{} = new(Stu)
	var i interface{} = (*Stu)(nil) // nil pointer to struct	type
	fmt.Printf("%p %v \n", &i, i)

	Judge(i)
}

//
// FuncA()
