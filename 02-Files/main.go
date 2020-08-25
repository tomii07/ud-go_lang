package main

import "fmt"
import "os"

func checkerror(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

f, err := os.Open("C:\\Users\\Lenovo\\Documents\\Cursos\\Go\\02-Files\\text.txt")
	checkerror(err)

	fileinfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	b1 := make([]byte, fileinfo.Size())
	f.Read(b1)
	checkerror(err)
	fmt.Printf(string(b1))

	f.Close()

}