package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	filePointer := openFile(os.Args[1])
	printFile(filePointer)
}

func openFile(filename string) *os.File {
	filePointer, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error when opening file: ", err)
		os.Exit(1)
	}

	return filePointer
}

func printFile(filePointer *os.File) {
	lw := logWriter{}
	io.Copy(lw, filePointer)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("byteslice written: ", len(bs))
	return len(bs), nil
}
