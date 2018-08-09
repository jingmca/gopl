//Package gopl dup lines
package gopl

import (
	"bufio"
	"fmt"
	"os"
)

type record struct {
	filename string
	linenu   []int
}

type recordList []record

//Dup func() will lookup duplicate lines from input
func Dup() {

	counts := make(map[string]int)
	records := make(map[string]recordList)

	for _, file := range os.Args[1:] {
		if len(file) == 0 {
			fmt.Println("dup: no files")
			continue
		}
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %s\n", err)
			continue
		}
		input := bufio.NewScanner(f)

		for nu := 0; input.Scan(); {
			line := input.Text()
			counts[line]++
			if len(records[line]) == 0 {
				records[line] = recordList{record{file, []int{nu + 1}}}
			} else {
				addRecord(&records, &line, &file, nu)
			}

		}
	}

	for line, n := range counts {
		fmt.Printf("------\nNUMS[%d]\tTEXT:<%s>\n", n, line)
		for _, list := range records[line] {
			fmt.Printf("FILE: <%s>\tLINE: %v\n", list.filename, list.linenu)
		}
	}
}

func addRecord(list *map[string]recordList, line *string, filename *string, linenu int) bool {

	ret := false
	alist := (*list)[*line]
	for i, item := range alist {
		if item.filename == *filename {
			(*list)[*line][i].linenu = append(item.linenu, linenu+1)
			ret = true
			break
		}
	}
	if ret == false {
		var tmp = record{*filename, []int{linenu + 1}}
		(*list)[*line] = append(alist, tmp)
	}

	return ret
}
