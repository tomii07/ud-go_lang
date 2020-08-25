package main

import "time"

func main() {

	a := make(chan string)
	b := make(chan string)
	c := make(chan string)
	d := make(chan string)

	in1 := "in1"
	in2 := "in2"
	in3 := "in3"
	in4 := "in4"

	go func(in string) {
		a <- testFunctionA(in)
	}(in1)

	go func(in string) {
		b <- testFunctionB(in)
	}(in2)

	go func(in string) {
		c <- testFunctionC(in)
	}(in3)

	go func(in string) {
		d <- testFunctionD(in)
	}(in4)

	finalterm := printAllFunction(<-a, <-b, <-c, <-d)

	println(finalterm)
}

func testFunctionA(input string) string {
	time.Sleep(9000)
	println(input)
	return input
}

func testFunctionB(input string) string {
	time.Sleep(6000)
	println(input)
	return input
}

func testFunctionC(input string) string {
	time.Sleep(3000)
	println(input)
	return input
}

func testFunctionD(input string) string {
	time.Sleep(1000)
	println(input)
	return input
}

func printAllFunction(a string, b string, c string, d string) string {
	println("-----")
	println(a)
	println(b)
	println(c)
	println(d)
	return "done"
}