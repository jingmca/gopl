package main

import (
	"fmt"
)

func main() {
	mess := make(chan int)

	for i := 0; i < 10; i++ {
		go func(x int) {
			for j := 0; j <= x; j++ {
				mess <- j
			}
			mess <- 0xfff
		}(i)
	}

	cnt := 0
	for m := range mess {
		fmt.Println(m)
		if m == 0xfff {
			cnt++
			if cnt == 10 {
				close(mess)
			}
		}
	}

}
