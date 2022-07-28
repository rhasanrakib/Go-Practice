package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println((a))

	f := "declare a variable and initial a value"
	fmt.Println(f)

	var b, c string = "multiple variable declare in same line", "these variables should be same type"

	fmt.Println(b, c)

	var d int
	fmt.Println("integer variables initialize with 0 = ", d)

	var e bool
	fmt.Println("boolen variables initialize with false = ", e)

	var g string
	fmt.Println("string variables initialize with empty= ", g)

}
