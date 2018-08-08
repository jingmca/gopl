//echo
package main

import (
	"fmt"
	"os"

	gopl "github.com/jingmca/gopl"
)

func echo() {
	for index, arg := range os.Args {
		fmt.Printf("%d %s\n", index, arg)
	}
}

func main() {
	gopl.Dup()
}
