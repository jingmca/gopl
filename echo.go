//echo
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/jingmca/gopl"
)

func echo() {
	for index, arg := range os.Args {
		fmt.Printf("%d %s\n", index, arg)
	}
}

func main() {
	var f = flag.String("c", "dup", "which command to execute")
	filename := flag.String("f", "", "url list file")

	flag.Parse()
	switch *f {
	case "dup":
		gopl.Dup()
	case "lisa":
		gopl.Lisass(os.Stdout)
	case "fetch":
		gopl.GFetch(filename)
	case "httpd":
		gopl.App()
	case "lcs":
		gopl.LCS([]rune("你吗"), []rune("你好家庭伟大吗"))
	case "maxsum":
		gopl.MaxSum([]int{1, -1, 3})
		gopl.MaxSum([]int{-9, -2, -3, 0, 1, 3})
	case "lss":
		gopl.LSS([]int{-9, -2, -3, 9, 1, 3}, 2)
	default:
		fmt.Println(*f)
	}

	//gopl.Dup()
	//gopl.Lisass(os.Stdout)
}

func cpuProfile() {
	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("CPU Profile started")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	time.Sleep(60 * time.Second)
	fmt.Println("CPU Profile stopped")
}

func heapProfile() {
	f, err := os.OpenFile("heap.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	time.Sleep(30 * time.Second)

	pprof.WriteHeapProfile(f)
	fmt.Println("Heap Profile generated")
}
