package main

func main() {

	// Intro
	println("Hello")

	// Data Types
	var str string
	str = "Test 1"

	str1 := "Test 2"

	println(str + str1)

	var i int32
	i = 5
	println(i)

	var f float32
	f = 5.5
	println(f)

	// Math
	var f1 float32
	var f2 float32
	f1 = 2.3
	f2 = 0.1

	calc := f1 + f2

	println(calc)

	// Functions
	firstFunction(str, str1)

	returnation := secondFunction(str1, str)
	println(returnation)

	return1, return2 := thirdFunction(str1, str)
	println(return1, return2)
}

func firstFunction(str1 string, str2 string) {
	println(str1 + str2)
}

func secondFunction(str1 string, str2 string) string {
	return str1 + str2
}

func thirdFunction(str1 string, str2 string) (string, string) {
	return str1 + str2, " puto"
}