package gopl

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//GFetch will start many goroutines for fetching url from cmd
func GFetch(f *string) {

	start := time.Now()
	ch := make(chan string)

	if *f == "" {
		for _, url := range flag.Args() {
			go fetch(url, ch)

			go fetch(url, ch)

		}

		for range flag.Args() {
			fmt.Println(<-ch)
		}
	} else {
		fio, err := os.Open(*f)
		if err != nil {
			fmt.Printf("while reading %s error %s", *f, err)
		}
		var i int
		input := bufio.NewScanner(fio)

		for ; input.Scan(); i++ {
			go fetch(input.Text(), ch)

			go fetch(input.Text(), ch)
		}

		fmt.Printf("%2d url spawned\n", i)
		for n := 2*i - 1; n >= 0; n-- {
			fmt.Println(<-ch)
		}

	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	if (!strings.HasPrefix(url, "http://")) && (!strings.HasPrefix(url, "https://")) {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- fmt.Sprintf("fetch %s status error : %d", url, resp.StatusCode)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("%s reading error: %s", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	// ch <- fmt.Sprintf("%s %.2fs %7dbytes \n-----\n%v\n-----", url, secs, nbytes, resp.Header)
	ch <- fmt.Sprintf("%s %5.2fs %7dbytes %s", url, secs, nbytes, resp.Status)
}
