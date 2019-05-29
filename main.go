// main.go
package main

import (
	"fmt"
	"github.com/sbinet/go-python"
)

func main() {
	python.Initialize()
	defer python.Finalize()

	helloModule := python.PyImport_ImportModule("hello")
	if helloModule == nil {
		panic("Error importing module")
	}

	helloFunc := helloModule.GetAttrString("hello_function")
	if helloFunc == nil {
		panic("Error importing function")
	}

	// The Python function takes no params but when using the C api
	// we're required to send (empty) *args and **kwargs anyways.
	res := helloFunc.Call(python.PyTuple_New(0), python.PyDict_New())
	str := python.PyString_AS_STRING(res.Str())

	fmt.Println(str)
}
