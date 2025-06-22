package main

import "fmt"

type Logger interface {
	Log()
}

type File struct{}

func (f File) Log() {
	fmt.Println("Logging from value receiver")
}

func (f *File) Save() {
	fmt.Println("Saving from pointer receiver")
}
func main() {
	var l Logger
	f1 := File{}
	f2 := &File{}
	l = f1 // Try this
	l = f1 // Try this
	l.Log()
	l = f2 // Try this too
	l.Log()
}
