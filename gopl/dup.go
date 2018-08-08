//Package gopl dup lines
package gopl

import (
	"fmt"
	"os"
)

//Dup func() will lookup duplicate lines from input
func Dup() {
	for _, files := range os.Args[1:] {
		if len(files) == 0 {
			fmt.Println("dup: no files")
			return
		}
		fmt.Printf("%T, %s", files, files)
		for _, file := range files {
			fmt.Println(file)
		}
	}
}
