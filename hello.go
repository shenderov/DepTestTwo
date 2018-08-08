package deptesttwo // import "github.com/shenderov/DepTestTwo"

import (
	"fmt"
	"github.com/shenderov/DepTestOne"
)

func Hello() {
	fmt.Println("This is dep three")
	deptestone.Hello()
}
