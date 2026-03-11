package base

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	name string
	age  int
}

func Test_getStructTag(t *testing.T) {
	user := &User{
		name: "test",
		age:  18,
	}

	filed, ok := reflect.TypeOf(user).Elem().FieldByName("name")
	if !ok {
		t.Fatal("field not found")
	}
	fmt.Println("filed: ", filed.Name, filed)
	tags := getStructTag(filed)

	fmt.Println(tags)
}
