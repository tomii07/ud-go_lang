package main

import "fmt"
import "os"

func checkerror(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	
	f, err := os.OpenFile("C:\\Users\\Lenovo\\Documents\\Cursos\\Go\\02-Files\\newFile.txt", os.O_CREATE|os.O_RDWR, 0666)
	checkerror(err)
	fileinfo, err1 := f.Stat()
	checkerror(err1)
	f.Seek(fileinfo.Size(), 0)
	f.Write([]byte("test"))
	f.Sync()
	f.Seek(0, 0)
	fileinfo1, err2 := f.Stat()
	checkerror(err2)

	bt := make([]byte, fileinfo1.Size())
	_, err3 := f.Read(bt)
	checkerror(err3)
	fmt.Println(string(bt))

	f.Close()
}