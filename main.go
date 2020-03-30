//This is called a package declaration.
//It informs the compiler whether to create an executable or a library.
//In this case, we're creating an executable named main.
package main

import (
	"fmt"
	t "time"
)

/*

import (
	You can import multiple packages using parentheses and refer to them using an alias.
	p1 "package1"
	p2 "package2"
)

*/

//func is a function declaration in Go
func main() {
	fmt.Println("Hello Go World!")
	fmt.Println(t.Now())
}
